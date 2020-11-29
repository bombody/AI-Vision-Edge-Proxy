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
  serialized_pb=b'\n\x15video_streaming.proto\x12\"chrys.cloud.videostreaming.v1beta1\"\xd2\x06\n\x0f\x41nnotateRequest\x12\x13\n\x0b\x64\x65vice_name\x18\x01 \x01(\t\x12\x18\n\x10remote_stream_id\x18\x02 \x01(\t\x12\x0c\n\x04type\x18\x03 \x01(\t\x12\x17\n\x0fstart_timestamp\x18\x04 \x01(\x03\x12\x15\n\rend_timestamp\x18\x05 \x01(\x03\x12\x13\n\x0bobject_type\x18\x06 \x01(\t\x12\x11\n\tobject_id\x18\x07 \x01(\t\x12\x1a\n\x12object_tracking_id\x18\x08 \x01(\t\x12\x12\n\nconfidence\x18\t \x01(\x01\x12J\n\x12object_bouding_box\x18\n \x01(\x0b\x32..chrys.cloud.videostreaming.v1beta1.BoudingBox\x12>\n\x08location\x18\x0b \x01(\x0b\x32,.chrys.cloud.videostreaming.v1beta1.Location\x12I\n\x11object_coordinate\x18\x0c \x01(\x0b\x32..chrys.cloud.videostreaming.v1beta1.Coordinate\x12<\n\x04mask\x18\r \x03(\x0b\x32..chrys.cloud.videostreaming.v1beta1.Coordinate\x12\x18\n\x10object_signature\x18\x0e \x03(\x01\x12\x10\n\x08ml_model\x18\x0f \x01(\t\x12\x18\n\x10ml_model_version\x18\x10 \x01(\t\x12\r\n\x05width\x18\x11 \x01(\x05\x12\x0e\n\x06height\x18\x12 \x01(\x05\x12\x13\n\x0bis_keyframe\x18\x13 \x01(\x08\x12\x12\n\nvideo_type\x18\x14 \x01(\t\x12\x18\n\x10offset_timestamp\x18\x15 \x01(\x03\x12\x17\n\x0foffset_duration\x18\x16 \x01(\x03\x12\x17\n\x0foffset_frame_id\x18\x17 \x01(\x03\x12\x18\n\x10offset_packet_id\x18\x18 \x01(\x03\x12\x15\n\rcustom_meta_1\x18\x19 \x01(\t\x12\x15\n\rcustom_meta_2\x18\x1a \x01(\t\x12\x15\n\rcustom_meta_3\x18\x1b \x01(\t\x12\x15\n\rcustom_meta_4\x18\x1c \x01(\t\x12\x15\n\rcustom_meta_5\x18\x1d \x01(\t\"h\n\x10\x41nnotateResponse\x12\x13\n\x0b\x64\x65vice_name\x18\x01 \x01(\t\x12\x18\n\x10remote_stream_id\x18\x02 \x01(\t\x12\x0c\n\x04type\x18\x03 \x01(\t\x12\x17\n\x0fstart_timestamp\x18\x04 \x01(\x03\"$\n\x08Location\x12\x0b\n\x03lat\x18\x01 \x01(\x01\x12\x0b\n\x03lon\x18\x02 \x01(\x01\"-\n\nCoordinate\x12\t\n\x01x\x18\x01 \x01(\x01\x12\t\n\x01y\x18\x02 \x01(\x01\x12\t\n\x01z\x18\x03 \x01(\x01\"F\n\nBoudingBox\x12\x0b\n\x03top\x18\x01 \x01(\x05\x12\x0c\n\x04left\x18\x02 \x01(\x05\x12\r\n\x05width\x18\x03 \x01(\x05\x12\x0e\n\x06height\x18\x04 \x01(\x05\"p\n\nShapeProto\x12?\n\x03\x64im\x18\x02 \x03(\x0b\x32\x32.chrys.cloud.videostreaming.v1beta1.ShapeProto.Dim\x1a!\n\x03\x44im\x12\x0c\n\x04size\x18\x01 \x01(\x03\x12\x0c\n\x04name\x18\x02 \x01(\t\"\xe2\x02\n\nVideoFrame\x12\r\n\x05width\x18\x01 \x01(\x03\x12\x0e\n\x06height\x18\x02 \x01(\x03\x12\x0c\n\x04\x64\x61ta\x18\x03 \x01(\x0c\x12\x11\n\ttimestamp\x18\x04 \x01(\x03\x12\x13\n\x0bis_keyframe\x18\x05 \x01(\x08\x12\x0b\n\x03pts\x18\x06 \x01(\x03\x12\x0b\n\x03\x64ts\x18\x07 \x01(\x03\x12\x12\n\nframe_type\x18\x08 \x01(\t\x12\x12\n\nis_corrupt\x18\t \x01(\x08\x12\x11\n\ttime_base\x18\n \x01(\x01\x12=\n\x05shape\x18\x0b \x01(\x0b\x32..chrys.cloud.videostreaming.v1beta1.ShapeProto\x12\x11\n\tdevice_id\x18\x0c \x01(\t\x12\x0e\n\x06packet\x18\r \x01(\x03\x12\x10\n\x08keyframe\x18\x0e \x01(\x03\x12\x11\n\textradata\x18\x0f \x01(\x0c\x12\x12\n\ncodec_name\x18\x10 \x01(\t\x12\x0f\n\x07pix_fmt\x18\x11 \x01(\t\">\n\x11VideoFrameRequest\x12\x16\n\x0ekey_frame_only\x18\x01 \x01(\x08\x12\x11\n\tdevice_id\x18\x02 \x01(\t\"\\\n\x19VideoFrameBufferedRequest\x12\x11\n\tdevice_id\x18\x01 \x01(\t\x12\x16\n\x0etimestamp_from\x18\x02 \x01(\x03\x12\x14\n\x0ctimestamp_to\x18\x03 \x01(\x03\"\xde\x01\n\nListStream\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0e\n\x06status\x18\x02 \x01(\t\x12\x16\n\x0e\x66\x61iling_streak\x18\x03 \x01(\x03\x12\x15\n\rhealth_status\x18\x04 \x01(\t\x12\x0c\n\x04\x64\x65\x61\x64\x18\x05 \x01(\x08\x12\x11\n\texit_code\x18\x06 \x01(\x03\x12\x0b\n\x03pid\x18\x07 \x01(\x05\x12\x0f\n\x07running\x18\x08 \x01(\x08\x12\x0e\n\x06paused\x18\t \x01(\x08\x12\x12\n\nrestarting\x18\n \x01(\x08\x12\x11\n\toomkilled\x18\x0b \x01(\x08\x12\r\n\x05\x65rror\x18\x0c \x01(\t\"\x13\n\x11ListStreamRequest\"6\n\x0cProxyRequest\x12\x11\n\tdevice_id\x18\x01 \x01(\t\x12\x13\n\x0bpassthrough\x18\x02 \x01(\x08\"7\n\rProxyResponse\x12\x11\n\tdevice_id\x18\x01 \x01(\t\x12\x13\n\x0bpassthrough\x18\x02 \x01(\x08\"2\n\x0eStorageRequest\x12\x11\n\tdevice_id\x18\x01 \x01(\t\x12\r\n\x05start\x18\x02 \x01(\x08\"3\n\x0fStorageResponse\x12\x11\n\tdevice_id\x18\x01 \x01(\t\x12\r\n\x05start\x18\x02 \x01(\x08\"\x88\x01\n\nVideoCodec\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05width\x18\x02 \x01(\x05\x12\x0e\n\x06height\x18\x03 \x01(\x05\x12\x0f\n\x07pix_fmt\x18\x04 \x01(\t\x12\x11\n\textradata\x18\x05 \x01(\x0c\x12\x16\n\x0e\x65xtradata_size\x18\x06 \x01(\x05\x12\x11\n\tlong_name\x18\x07 \x01(\t\"&\n\x11VideoProbeRequest\x12\x11\n\tdevice_id\x18\x01 \x01(\t\"\x9a\x01\n\x12VideoProbeResponse\x12\x43\n\x0bvideo_codec\x18\x01 \x01(\x0b\x32..chrys.cloud.videostreaming.v1beta1.VideoCodec\x12?\n\x06\x62uffer\x18\x02 \x01(\x0b\x32/.chrys.cloud.videostreaming.v1beta1.VideoBuffer\"v\n\x0bVideoBuffer\x12\x12\n\nstart_time\x18\x01 \x01(\x03\x12\x10\n\x08\x65nd_time\x18\x02 \x01(\x03\x12\x18\n\x10\x64uration_seconds\x18\x03 \x01(\x03\x12\x17\n\x0f\x61pproximate_fps\x18\x04 \x01(\x05\x12\x0e\n\x06\x66rames\x18\x05 \x01(\x03\"*\n\x12SystemTimeResponse\x12\x14\n\x0c\x63urrent_time\x18\x01 \x01(\x03\"\x13\n\x11SystemTimeRequest2\xe5\x07\n\x05Image\x12{\n\x10VideoLatestImage\x12\x35.chrys.cloud.videostreaming.v1beta1.VideoFrameRequest\x1a..chrys.cloud.videostreaming.v1beta1.VideoFrame\"\x00\x12\x87\x01\n\x12VideoBufferedImage\x12=.chrys.cloud.videostreaming.v1beta1.VideoFrameBufferedRequest\x1a..chrys.cloud.videostreaming.v1beta1.VideoFrame\"\x00\x30\x01\x12}\n\nVideoProbe\x12\x35.chrys.cloud.videostreaming.v1beta1.VideoProbeRequest\x1a\x36.chrys.cloud.videostreaming.v1beta1.VideoProbeResponse\"\x00\x12x\n\x0bListStreams\x12\x35.chrys.cloud.videostreaming.v1beta1.ListStreamRequest\x1a..chrys.cloud.videostreaming.v1beta1.ListStream\"\x00\x30\x01\x12w\n\x08\x41nnotate\x12\x33.chrys.cloud.videostreaming.v1beta1.AnnotateRequest\x1a\x34.chrys.cloud.videostreaming.v1beta1.AnnotateResponse\"\x00\x12n\n\x05Proxy\x12\x30.chrys.cloud.videostreaming.v1beta1.ProxyRequest\x1a\x31.chrys.cloud.videostreaming.v1beta1.ProxyResponse\"\x00\x12t\n\x07Storage\x12\x32.chrys.cloud.videostreaming.v1beta1.StorageRequest\x1a\x33.chrys.cloud.videostreaming.v1beta1.StorageResponse\"\x00\x12}\n\nSystemTime\x12\x35.chrys.cloud.videostreaming.v1beta1.SystemTimeRequest\x1a\x36.chrys.cloud.videostreaming.v1beta1.SystemTimeResponse\"\x00\x62\x06proto3'
)




_ANNOTATEREQUEST = _descriptor.Descriptor(
  name='AnnotateRequest',
  full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='device_name', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.device_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='remote_stream_id', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.remote_stream_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='type', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.type', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='start_timestamp', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.start_timestamp', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='end_timestamp', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.end_timestamp', index=4,
      number=5, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='object_type', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.object_type', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='object_id', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.object_id', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='object_tracking_id', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.object_tracking_id', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='confidence', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.confidence', index=8,
      number=9, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='object_bouding_box', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.object_bouding_box', index=9,
      number=10, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='location', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.location', index=10,
      number=11, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='object_coordinate', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.object_coordinate', index=11,
      number=12, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='mask', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.mask', index=12,
      number=13, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='object_signature', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.object_signature', index=13,
      number=14, type=1, cpp_type=5, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ml_model', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.ml_model', index=14,
      number=15, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ml_model_version', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.ml_model_version', index=15,
      number=16, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='width', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.width', index=16,
      number=17, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='height', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.height', index=17,
      number=18, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='is_keyframe', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.is_keyframe', index=18,
      number=19, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='video_type', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.video_type', index=19,
      number=20, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='offset_timestamp', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.offset_timestamp', index=20,
      number=21, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='offset_duration', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.offset_duration', index=21,
      number=22, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='offset_frame_id', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.offset_frame_id', index=22,
      number=23, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='offset_packet_id', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.offset_packet_id', index=23,
      number=24, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='custom_meta_1', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.custom_meta_1', index=24,
      number=25, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='custom_meta_2', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.custom_meta_2', index=25,
      number=26, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='custom_meta_3', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.custom_meta_3', index=26,
      number=27, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='custom_meta_4', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.custom_meta_4', index=27,
      number=28, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='custom_meta_5', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateRequest.custom_meta_5', index=28,
      number=29, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=62,
  serialized_end=912,
)


_ANNOTATERESPONSE = _descriptor.Descriptor(
  name='AnnotateResponse',
  full_name='chrys.cloud.videostreaming.v1beta1.AnnotateResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='device_name', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateResponse.device_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='remote_stream_id', full_name='chrys.cloud.videostreaming.v1beta1.AnnotateResponse.remote_stream_id', index=1,
      num