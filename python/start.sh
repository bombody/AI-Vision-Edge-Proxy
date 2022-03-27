#!/bin/bash

# stop bash script on error
# set -o errexit
set -o noclobber # enable >|
set -e

rtsp_endpoint=${rtsp_endpoint}
device_id=${device_id}
rtmp_endpoint=${rtmp_endpoint}
in_memory_buffer=${in_memory_buffer}
in_memory_scale=${memory_scale}
disk_buffer_path=${disk_buffer_path}
disk_cleanup_rate=${disk_cleanup_rate}
redis_host=${redis_host}
redis_port=${redis_port}

 if [ -z "$rtsp_endpoint" ]; then 
    echo "rtsp endpoint must be defined in environment variables"
    exit 1 
fi
if [ -z "$device_id" ]; then 
    echo "device_id endpoint must be defined in environment variables"
    exit 1
fi


echo "Connection to rtsp camera"
source activate chrysedgeai

cmd=" -u rtsp_to_rtmp.py --rtsp $rtsp_endpoint --device_id $device_i