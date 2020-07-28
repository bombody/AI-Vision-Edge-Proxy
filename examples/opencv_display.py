
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


import grpc
import video_streaming_pb2_grpc, video_streaming_pb2
import argparse
import cv2
import numpy as np
import time
import os

def gen_buffer_probe_request(device_name):
    """ Create GRPC request to get in memory probe info """

    req = video_streaming_pb2.VideoProbeRequest()
    req.device_id = device_name

    return req

def gen_image_request(device_name, keyframe_only):
    """ Create an object to request a video frame """


    req = video_streaming_pb2.VideoFrameRequest()
    req.device_id = device_name
    req.key_frame_only = keyframe_only
    return req

if __name__ == "__main__":
    
    parser = argparse.ArgumentParser(description='Chrysalis Edge Proxy Basic Example')
    parser.add_argument("--device", type=str, default=None, required=True)
    parser.add_argument("--keyframe", action='store_true')
    args = parser.parse_args()

    # grpc connection to video-edge-ai-proxy
    options = [('grpc.max_receive_message_length', 50 * 1024 * 1024)]
    channel = grpc.insecure_channel('127.0.0.1:50001', options=options)
    stub = video_streaming_pb2_grpc.ImageStub(channel)

    # displaying video info (if exists)
    probe = stub.VideoProbe(gen_buffer_probe_request(device_name=args.device))
    aproxFps = 30
    if probe.buffer:
        if probe.buffer.approximate_fps > 0:
            aproxFps = probe.buffer.approximate_fps

    print(probe)
    
    # requesting video frames
    req = gen_image_request(device_name=args.device,keyframe_only=args.keyframe)
    while True:
        frame = stub.VideoLatestImage(req)
        # read raw frame data and convert to numpy array
        if frame is not None:
            img_bytes = frame.data 
            re_img = np.frombuffer(img_bytes, dtype=np.uint8)

            # reshape image back into original dimensions
            if len(frame.shape.dim) > 0:
                reshape = tuple([int(dim.size) for dim in frame.shape.dim])
                re_img = np.reshape(re_img, reshape)

                # # display image
                cv2.imshow('box', re_img)
                
                if cv2.waitKey(1) & 0xFF == ord('q'):
                    break
        
        # delay by assumed fps rate
        delay = 1 / aproxFps
        time.sleep(delay)