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
redis_port