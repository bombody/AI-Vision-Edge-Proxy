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
    redis_conn.delete(RedisInMemoryIFrameListPrefix+device_id) # all keys for stored i-frames
    redis_conn.delete(RedisInMemoryDecodedImagesPrefix+device_id) # all decoded in-memory buffer images

def setCodecInfo(redis_conn, in_av_container,deviceId):
    '''
    Sets the current streams codec info at the same time clean out the in memory redis queues
    '''
    streams = in_av_container.streams
    if len(streams) > 0:
        for stream in streams:
            if stream.type == "video":

                codec_ctx = stream.codec_context
                vc = video_streaming_pb2.VideoCodec()
                vc.name = codec_ctx.name
                vc.long_name = codec_ctx.codec.long_name
                vc.width = codec_ctx.width
                vc.height = codec_ctx.height
                vc.pix_fmt = codec_ctx.pix_fmt
                vc.extradata = codec_ctx.extradata
                vc.extradata_size = codec_ctx.extradata_size

                vcData = vc.SerializeToString()
                redis_conn.set(RedisCodecVideoInfo+deviceId, vcData)


def getCodecInfo(redis_conn, deviceId):
    '''
    Reading the current video stream codec info from redis
    '''
    info = redis_conn.get(RedisCodecVideoInfo+deviceId)
    if info is not None:
        vc = video_streaming_pb2.VideoCodec()
        vc.ParseFromString(info)
        return vc
    return None

def packetToInMemoryBuffer(redis_conn,memory_buffer_size, device_id,in_av_container, packet):
    if memory_buffer_size > 0:
        
        redisStreamName = RedisInMemoryQueuePrefix + device_id
        redisIFrameList = RedisInMemoryIFrameListPrefix + device_id

        for stream in in_av_container.streams:
            if stream.type == "video":
                codec_ctx = stream.codec_context
                video_height = codec_ctx.height
                video_width = codec_ctx.width
                is_keyframe = packet.is_keyframe
                packetBytes = packet.to_bytes()
                codec_name = codec_ctx.name
                pix_fmt = codec_ctx.pix_fmt

                vf = video_streaming_pb2.VideoFrame()
                vf.data = packetBytes
                vf.width = video_width
                vf.height = video_height
                vf.timestamp = int(packet.pts * float(packet.time_base))
                vf.pts = packet.pts
                vf.dts = packet.dts
                vf.keyframe = is_keyframe
                vf.time_base = float(packet.time_base)
                vf.is_keyframe = packet.is_keyframe
                vf.is_corrupt = packet.is_corrupt
                vf.codec_name = codec_name
                vf.pix_fmt = pix_fmt

                vfData = vf.SerializeToString()
   