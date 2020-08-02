import grpc
import video_streaming_pb2_grpc, video_streaming_pb2
import argparse

def gen_buffer_probe_request(device_name):
    """ Create GRPC request to get in memory probe info """

    req = video_streaming_pb2.VideoProbeRequest()
    req.device_id = device_name

    return req

def gen_system_time_request():

    return video_streaming_pb2.SystemTimeRequest()

if __name__ == "__main__":

    parser = argparse.ArgumentParser(description='Chrysalis Edge buffered images example')
 