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

            excep