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


class RTSPtoRTMP(thr