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

import threading, queue
from proto import video_streaming_pb2
import numpy as np
import time
import os
import redis
from global_vars import query_timestamp, RedisLastAccessPrefix, RedisIsKeyFrameOnlyPrefix, ChrysException
from datetime import datetime

class ReadImage(threading.Thread):

    def __init__(self, packet_queue, device_id, memory_buffer, redis_conn, is_decode_packets_event, lock_condition):
        threading.Thread.__init__(self)
        self._packet_queue = packet_queue
        self.device_id = device_id
        self._memory_buffer = memory_buffer
        self.redis_conn = redis_conn
        self.is_decode_packets_event = is_decode_packets_event
        self.lock_condition = lock_condition
        self.last_query_timestamp = 0
        self.packet_group = []

    # checks if only keyframes requested
    def check_decode_only_keyframes(self):
        global RedisIsKeyFrameOnlyPrefix
        decode_only_keyframes = False
        decodeOnlyKeyFramesKey = RedisIsKeyFrameOnlyPrefix + self.device_id
        only_keyframes = self.redis_conn.get(decodeOnlyKeyFramesKey)
        if only_ke