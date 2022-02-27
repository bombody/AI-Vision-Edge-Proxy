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
        for c, n in zip(nodes, nodes[1