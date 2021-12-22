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

import threading
from typing import MutableSequence
import av
import base64
import redis
import json
import sys
import io
import numpy as np
import time
from proto import video_streaming_pb2
import multiprocessing

# constants from global vars
from global_vars import RedisInMemoryBufferChannel,RedisInMemoryDecodedImagesPrefix, RedisInMemoryIFrameListPrefix,RedisCodecVideoInfo,RedisInMemoryQueuePrefix

def memoryCleanup(redis_conn, device_id):
    '''
    Cleanup redis memory
    '''
    redis_conn.delete(RedisInMemoryQueuePrefix+device_id) # the complete memory buffer of compressed stream
    redis_conn.delete(RedisInMemoryIFrameLi