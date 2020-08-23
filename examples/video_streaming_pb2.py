# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: video_streaming.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='video_streaming.proto',
  package='chrys.cloud.videostreaming.v1beta1',
  syntax='proto3',
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x15video_streaming.proto\x12\"chrys.cloud.videostreaming.v1beta1\"\xd2\x06\n\x0f\x41nnotateRequest\x12\x13\n\x0b\x64\x65vice_name\x18\x01 \x01(\t\x12\x18\n\x10remote_stream_id\x18\x02 \x01(\t\x12\x0c\n\x04type\x18\x03 \x01(\t\x12\x17\n\x0fstart_timestamp\x18\x04 \x01(\x03\x12\x15\n\rend_timestamp\x18\x05 \x01(\x03\x12\x13\n\x0bobject_type\x18\x06 \x01(\t\x12\x11\n\tobject_id\x18\x07 \x01(\t\x12\x1a\n\x12object_tracking_id\x18\x08 \x01(\t\x12\x12\n\nconfidence\x18\t \x01(\x01\x12J\n\x12object_bouding_box\x18\n \x01(\x0b\x32..chrys.cloud.videostreaming.v1beta1.BoudingBox\x12>\n\x08location\x18\x0b \x01(\x0b\x32,.chrys.cloud.videostreaming.v1beta1.Location\x12I\n\x11object_coordinate\x18\x0c \x01(\x0b\x32..chrys.cloud.videostreaming.v1beta1.Coordinate\x12<\n\x04mask\x18\r \x03(\x0b\x32..chrys.cloud.videostreaming.v1beta1.Coordinate\x12\x18\n\x10object_signature\x18\x0e \x03(\x01\x12\x10\n\x08ml_model\x18\x0f \x01(\t\x12\x18\n\x10ml_model_version\x18\x10 \x01(\t\x12\r\n\x05width\x18\x11 \x01(\x05\x12\x0e\n\x06height\x18\x12 \x01(\x05\x12\x13\n\x0bis_keyframe\x18\x13 \x01(\x08\x12\x12\n\nvideo_type\x18\x14 \x01(\t\x12\x18\n\x10offset_timestamp\x18\x15 \x01(\x03\x12\x17\n\x0foffset_duration\x18\x16 \x01(\x03\x12\x17\n\x0foffset_frame_id\x18\x17 \x01(\x03\x12\x18\n\x10offset_packet_id\x18\x18 \x01(\x03\x12\x15\n\rcustom_meta_1\x18\x19 \x01(\t\x12\x15\n\rcustom_meta_2\x18\x1a \x01(\t\x12\x15\n\rcustom_meta_3\x18\x1b \x01(\t\x12\x15\n\rcustom_meta_4\x18\x1c \x01(\t\x12\x15\n\rcustom_meta_5\x18\x1d \x01(\t\"h\n\x10\x41nnotateResponse\x12\x13\n\x0b\x64\x65vice_name\x18\x01 \x01(\t\x12\x18\n\x10remote_stream_id\x18\x02 \x01(\t\x12\x0c\n\x04type\x18\x03 \x01(\t\x12\x17\n\x0fstart_timestamp\x18\x04 \x01(\x03\"$\n\x08Location\x12\x0b\n\x03lat\x18\x01 \x01(\x01\x12\x0b\n\x03lon\x18\x02 \x01(\x01\"-\n\nCoordinate\x12\t\n\x01x\x18\x01 \x01(\x01\x12\t\n\x01y\x18\x02 \x01(\x01\x12\t\n\x01z\x18\x03 \x01(\x01\"F\n\nBoudingBox\x12\x0b\n\x03top\x18\x01 \x01(\x05\x12\x0c\n\x04left\x18\x02 \x01(\x05\x12\r\n\x05width\x18\x03 \x01(\x05\x12\x0e\n\x06height\x18\x04 \x01(\x05\"p\n\nShapeProto\x12?\n\x03\x64im\x18\x02 \x03(\x0b\x32\x32.chrys.cloud.videostreaming.v1beta1.ShapeProto.Dim\x1a!\n\x03\x44im\x12\x0c\n\x04size\x18\x01 \x01(\x03\x12\x0c\n\x04name\x18\x02 \x01(\t\"\xe2\x02\n\nVideoFrame\x12\r\n\x05width\x18\x01 \x01(\x03\x12\x0e\n\x06height\x18\x02 \x01(\x03\x12\x0c\n\x04\x64\x61ta\x18\x03 \x01(\x0c\x12\x11\n\ttimestamp\x18\x04 \x01(\x03\x12\x13\n\x0bis_keyframe\x18\x05 \x01(\x08\x12\x0b\n\x03pts\x18\x06 \x01(\x03\x12\x0b\n\x03\x64ts\x18\x07 \x01(\x03\x12\x12\n\nframe_type\x18\x08 \x01(\t\x12\x12\n\nis_corrupt\x18\t \x01(\x08\x12\x11\n\ttime_base\x18\n \x01(\x01\x12=\n\x05shape\x18\x0b \x01(\x0b\x32..chrys.cloud.videostreaming.v1beta1.ShapeProto\x12\x11\n\tdevice_id\x18\x0c \x01(\t\x12\x0e\n\x06packet\x18\r \x01(\x03\x12\x10\n\x08keyframe\x18\x0e \x01(\x03\x12\x11\n\textradata\x18\x0f \x01(\x0c\x12\x12\n\ncodec_name\x18\x10 \x01(\t\x12\x0f\n\x07pix_fmt\x18\x11 \x01(\t\">\n\x11VideoFrameRequest\x12\x16\n\x0ekey_frame_only\x18\x01 \x01(\x08\x12\x11\n\tdevice_id\x18\x02 \x01(\t\"\\\n\x19VideoFrameBufferedRequest\x12\x11\n\tdevice_id\x18\x01 \x01(\t\x12\x16\n\x0etimestamp_from\x18\x02 \x01(\x03\x12\x14\n\x0ctimestamp_to\x18\x03 \x01(\x03\"\xde\x01\n\nListStream\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0e\n\x06status\x18\x02 \x01(\t\x12\x16\n\x0e\x66\x61iling_streak\x18\x03 \x01(\x03\x12\x15\n\rhealth_status\x18\x04 \x0