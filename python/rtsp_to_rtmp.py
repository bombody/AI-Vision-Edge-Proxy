# Copyright 2020 Wearless Tech Inc All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import av
import time
import os
import io
from av.filter import Filter, Graph
from av.codec import CodecContext
import redis
import threading, queue
from read_image import ReadImage
import random
from argparse import ArgumentParser
import sys
from archive import StoreMP4VideoChunks
from disk_cleanup import CleanupScheduler
from inmemory_buffer import InMemoryBuffer, packetToInMemoryBuffer, setCodecInfo,getCodecInfo, memoryCleanup
from global_vars import query_timestamp, RedisIsKeyFrameOnlyPrefix, RedisLastAccessPrefix, ArchivePacketGroup
import datetime


class RTSPtoRTMP(threading.Thread):

    def __init__(self, rtsp_endpoint, rtmp_endpoint, packet_queue, device_id, disk_path, redis_conn, memory_buffer, is_decode_packets_event, lock_condition):
        threading.Thread.__init__(self) 
        self._packet_queue = packet_queue
        self._disk_path = disk_path
        self.rtsp_endpoint = rtsp_endpoint
        self.rtmp_endpoint = rtmp_endpoint
        self.redis_conn = redis_conn
        self.device_id = device_id
        self.__memory_buffer_size = memory_buffer
        self.is_decode_packets_event = is_decode_packets_event
        self.lock_condition = lock_condition
        self.query_timestamp = query_timestamp

    def link_nodes(self,*nodes):
        for c, n in zip(nodes, nodes[1:]):
            c.link_to(n)

    def join(self):
        threading.Thread.join(self)
        if self.exc:
            raise self.exc

    def run(self):
        global RedisLastAccessPrefix    

        # cleanup all redis memory
        try:
            memoryCleanup(self.redis_conn, self.device_id)
        except Exception as ex:
            self.exc = ex
            os._exit(1)

        current_packet_group = []
        flush_current_packet_group = False

        # init archiving 
        iframe_start_timestamp = 0
        packet_group_queue = queue.Queue()

        apg:ArchivePacketGroup = None

        should_mux = False

        last_loop_run = int(time.time() * 1000)

        while True:
            try:
                options = {'rtsp_transport': 'tcp','stimeout': '5000000', 'max_delay': '5000000', 'use_wallclock_as_timestamps':"1", "fflags":"+genpts", 'acodec':'aac'}
                self.in_container = av.open(self.rtsp_endpoint, options=options)
                self.in_video_stream = self.in_container.streams.video[0]
                self.in_audio_stream = None
                if len(self.in_container.streams.audio) > 0:
                    for c in self.in_container.streams.audio:
                        print(c)
                    # self.in_audio_stream = self.in_container.streams.audio[0]

                # set codec context in redis
                setCodecInfo(self.redis_conn, self.in_container, self.device_id)

                # init mp4 local archive
                if self._disk_path is not None:
                    self._mp4archive = StoreMP4VideoChunks(queue=packet_group_queue, path=self._disk_path, device_id=self.device_id, video_stream=self.in_video_stream, audio_stream=self.in_audio_stream)
                    self._mp4archive.daemon = True
                    self._mp4archive.start()

            except Exception as ex:
                print("failed to connect to RTSP camera", ex)
                self.exc = ex
                os._exit(1)
            
            keyframe_found = False
            global query_timestamp

            if self.rtmp_endpoint is not None:
                output = av.open(self.rtmp_endpoint, format="flv", mode='w')
                output_video_stream = output.add_stream(template=self.in_video_stream)  

            output_audio_stream = None
            if self.in_audio_stream is not None and self.rtmp_endpoint is not None:
                output_audio_stream = output.add_stream(template=self.in_audio_stream)


            for packet in self.in_container.demux(self.in_video_stream):

                if packet.dts is None:
                    continue
                
                if packet.is_keyframe:
                    # if we already found a keyframe previously, archive what we have

                    if len(current_packet_group) > 0:
                        packet_group = current_packet_group.copy()
                        
                        # send to archiver! (packet_group, iframe_start_timestamp)
                        if self._disk_path is not None:
                            apg = ArchivePacketGroup(packet_group, iframe_start_timestamp)
                            packet_group_queue.put(apg)

                    keyframe_found = True
                    current_packet_group = []
                    iframe_start_timestamp = int(round(time.time() * 1000))

                if keyframe_found == False:
                    print("skipping, since not a keyframe")
                    continue

                if keyframe_found == False:
                    print("skipping, since not a keyframe")
                    continue

                # method to push packet to the redis in a in-memory buffer
                try:
                    packetToInMemoryBuffer(self.redis_conn, self.__memory_buffer_size, self.device_id, self.in_container, packet)
                except Exception as ex:
                    self.exc = ex
                    print(ex)
                    os._exit(1)

                '''
                Live Redis Settings
                -------------------
                This should be invoked only every 500 ms, This If needs to moved to it's own method
                '''
                # shouldn't be a problem for redis but maybe every 200ms to query for latest timestamp only
                settings_dict = self.redis_conn.hgetall(RedisLastAccessPrefix + device_id)

                if settings_dict is not None and len(settings_dict) > 0:
                    settings_dict