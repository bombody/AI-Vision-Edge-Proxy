
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: video_streaming.proto

package chrys_cloud_videostreaming_v1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Annotation messages
type AnnotateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceName       string        `protobuf:"bytes,1,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`                          // required: device name (required) identity of device
	RemoteStreamId   string        `protobuf:"bytes,2,opt,name=remote_stream_id,json=remoteStreamId,proto3" json:"remote_stream_id,omitempty"`            //optional: if associated with storage, the ID of Chrysalis Cloud deviceID
	Type             string        `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`                                                        // required: event type: e.g. moving, exit, entry, stopped, parked, ...
	StartTimestamp   int64         `protobuf:"varint,4,opt,name=start_timestamp,json=startTimestamp,proto3" json:"start_timestamp,omitempty"`             //required: start of the event
	EndTimestamp     int64         `protobuf:"varint,5,opt,name=end_timestamp,json=endTimestamp,proto3" json:"end_timestamp,omitempty"`                   // optional: event of the event
	ObjectType       string        `protobuf:"bytes,6,opt,name=object_type,json=objectType,proto3" json:"object_type,omitempty"`                          // optional: e.g. person, car, face, bag, roadsign,...
	ObjectId         string        `protobuf:"bytes,7,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`                                // optional: e.g. object id from the ML model
	ObjectTrackingId string        `protobuf:"bytes,8,opt,name=object_tracking_id,json=objectTrackingId,proto3" json:"object_tracking_id,omitempty"`      // optional: tracking id of the object
	Confidence       float64       `protobuf:"fixed64,9,opt,name=confidence,proto3" json:"confidence,omitempty"`                                          // confidence of inference [0-1.0]
	ObjectBoudingBox *BoudingBox   `protobuf:"bytes,10,opt,name=object_bouding_box,json=objectBoudingBox,proto3" json:"object_bouding_box,omitempty"`     // optional: object bounding box
	Location         *Location     `protobuf:"bytes,11,opt,name=location,proto3" json:"location,omitempty"`                                               // optional: object GEO location
	ObjectCoordinate *Coordinate   `protobuf:"bytes,12,opt,name=object_coordinate,json=objectCoordinate,proto3" json:"object_coordinate,omitempty"`       // optional: object coordinates within the image
	Mask             []*Coordinate `protobuf:"bytes,13,rep,name=mask,proto3" json:"mask,omitempty"`                                                       // optional" object mask (polygon)
	ObjectSignature  []float64     `protobuf:"fixed64,14,rep,packed,name=object_signature,json=objectSignature,proto3" json:"object_signature,omitempty"` // optional: signature of the detected item
	MlModel          string        `protobuf:"bytes,15,opt,name=ml_model,json=mlModel,proto3" json:"ml_model,omitempty"`                                  // optional: description of the module that generated this event
	MlModelVersion   string        `protobuf:"bytes,16,opt,name=ml_model_version,json=mlModelVersion,proto3" json:"ml_model_version,omitempty"`           // optional: version of the ML model
	Width            int32         `protobuf:"varint,17,opt,name=width,proto3" json:"width,omitempty"`                                                    // optional: image width
	Height           int32         `protobuf:"varint,18,opt,name=height,proto3" json:"height,omitempty"`                                                  // optional: image height
	IsKeyframe       bool          `protobuf:"varint,19,opt,name=is_keyframe,json=isKeyframe,proto3" json:"is_keyframe,omitempty"`                        // optional: true/false if this annotation is from keyframe
	VideoType        string        `protobuf:"bytes,20,opt,name=video_type,json=videoType,proto3" json:"video_type,omitempty"`                            // optional: e.g. mp4 filename, live stream, ...
	OffsetTimestamp  int64         `protobuf:"varint,21,opt,name=offset_timestamp,json=offsetTimestamp,proto3" json:"offset_timestamp,omitempty"`         // optional: offset from the beginning
	OffsetDuration   int64         `protobuf:"varint,22,opt,name=offset_duration,json=offsetDuration,proto3" json:"offset_duration,omitempty"`            // optional: duration from the offset
	OffsetFrameId    int64         `protobuf:"varint,23,opt,name=offset_frame_id,json=offsetFrameId,proto3" json:"offset_frame_id,omitempty"`             // optional: frame id of the
	OffsetPacketId   int64         `protobuf:"varint,24,opt,name=offset_packet_id,json=offsetPacketId,proto3" json:"offset_packet_id,omitempty"`          // optional: offset of the packet
	// extending the event message meta data (optional)
	CustomMeta_1 string `protobuf:"bytes,25,opt,name=custom_meta_1,json=customMeta1,proto3" json:"custom_meta_1,omitempty"` // e.g. gender, hair, car model, ...
	CustomMeta_2 string `protobuf:"bytes,26,opt,name=custom_meta_2,json=customMeta2,proto3" json:"custom_meta_2,omitempty"`
	CustomMeta_3 string `protobuf:"bytes,27,opt,name=custom_meta_3,json=customMeta3,proto3" json:"custom_meta_3,omitempty"`
	CustomMeta_4 string `protobuf:"bytes,28,opt,name=custom_meta_4,json=customMeta4,proto3" json:"custom_meta_4,omitempty"`
	CustomMeta_5 string `protobuf:"bytes,29,opt,name=custom_meta_5,json=customMeta5,proto3" json:"custom_meta_5,omitempty"`
}

func (x *AnnotateRequest) Reset() {
	*x = AnnotateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_streaming_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnotateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotateRequest) ProtoMessage() {}

func (x *AnnotateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_streaming_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnotateRequest.ProtoReflect.Descriptor instead.
func (*AnnotateRequest) Descriptor() ([]byte, []int) {
	return file_video_streaming_proto_rawDescGZIP(), []int{0}
}

func (x *AnnotateRequest) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *AnnotateRequest) GetRemoteStreamId() string {
	if x != nil {
		return x.RemoteStreamId
	}
	return ""
}

func (x *AnnotateRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AnnotateRequest) GetStartTimestamp() int64 {
	if x != nil {
		return x.StartTimestamp
	}
	return 0
}

func (x *AnnotateRequest) GetEndTimestamp() int64 {
	if x != nil {
		return x.EndTimestamp
	}
	return 0
}

func (x *AnnotateRequest) GetObjectType() string {
	if x != nil {
		return x.ObjectType
	}
	return ""
}

func (x *AnnotateRequest) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *AnnotateRequest) GetObjectTrackingId() string {
	if x != nil {
		return x.ObjectTrackingId
	}
	return ""
}

func (x *AnnotateRequest) GetConfidence() float64 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

func (x *AnnotateRequest) GetObjectBoudingBox() *BoudingBox {
	if x != nil {
		return x.ObjectBoudingBox
	}
	return nil
}

func (x *AnnotateRequest) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *AnnotateRequest) GetObjectCoordinate() *Coordinate {
	if x != nil {
		return x.ObjectCoordinate
	}
	return nil
}

func (x *AnnotateRequest) GetMask() []*Coordinate {
	if x != nil {
		return x.Mask
	}
	return nil
}

func (x *AnnotateRequest) GetObjectSignature() []float64 {
	if x != nil {
		return x.ObjectSignature
	}
	return nil
}

func (x *AnnotateRequest) GetMlModel() string {
	if x != nil {
		return x.MlModel
	}
	return ""
}

func (x *AnnotateRequest) GetMlModelVersion() string {
	if x != nil {
		return x.MlModelVersion
	}
	return ""
}

func (x *AnnotateRequest) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *AnnotateRequest) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *AnnotateRequest) GetIsKeyframe() bool {
	if x != nil {
		return x.IsKeyframe
	}
	return false
}

func (x *AnnotateRequest) GetVideoType() string {
	if x != nil {
		return x.VideoType
	}
	return ""
}

func (x *AnnotateRequest) GetOffsetTimestamp() int64 {
	if x != nil {
		return x.OffsetTimestamp
	}
	return 0
}

func (x *AnnotateRequest) GetOffsetDuration() int64 {
	if x != nil {
		return x.OffsetDuration
	}
	return 0
}

func (x *AnnotateRequest) GetOffsetFrameId() int64 {
	if x != nil {
		return x.OffsetFrameId
	}
	return 0
}

func (x *AnnotateRequest) GetOffsetPacketId() int64 {
	if x != nil {
		return x.OffsetPacketId
	}
	return 0
}

func (x *AnnotateRequest) GetCustomMeta_1() string {
	if x != nil {
		return x.CustomMeta_1
	}
	return ""
}

func (x *AnnotateRequest) GetCustomMeta_2() string {
	if x != nil {
		return x.CustomMeta_2
	}
	return ""
}

func (x *AnnotateRequest) GetCustomMeta_3() string {
	if x != nil {
		return x.CustomMeta_3
	}
	return ""
}

func (x *AnnotateRequest) GetCustomMeta_4() string {
	if x != nil {
		return x.CustomMeta_4
	}
	return ""
}

func (x *AnnotateRequest) GetCustomMeta_5() string {
	if x != nil {
		return x.CustomMeta_5
	}
	return ""
}

type AnnotateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceName     string `protobuf:"bytes,1,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	RemoteStreamId string `protobuf:"bytes,2,opt,name=remote_stream_id,json=remoteStreamId,proto3" json:"remote_stream_id,omitempty"`
	Type           string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	StartTimestamp int64  `protobuf:"varint,4,opt,name=start_timestamp,json=startTimestamp,proto3" json:"start_timestamp,omitempty"`
}

func (x *AnnotateResponse) Reset() {
	*x = AnnotateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_streaming_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnotateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotateResponse) ProtoMessage() {}

func (x *AnnotateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_streaming_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnotateResponse.ProtoReflect.Descriptor instead.
func (*AnnotateResponse) Descriptor() ([]byte, []int) {
	return file_video_streaming_proto_rawDescGZIP(), []int{1}
}

func (x *AnnotateResponse) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *AnnotateResponse) GetRemoteStreamId() string {
	if x != nil {
		return x.RemoteStreamId
	}
	return ""
}

func (x *AnnotateResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AnnotateResponse) GetStartTimestamp() int64 {
	if x != nil {
		return x.StartTimestamp
	}
	return 0
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat float64 `protobuf:"fixed64,1,opt,name=lat,proto3" json:"lat,omitempty"` // latitude
	Lon float64 `protobuf:"fixed64,2,opt,name=lon,proto3" json:"lon,omitempty"` // longitude
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_streaming_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_video_streaming_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_video_streaming_proto_rawDescGZIP(), []int{2}
}

func (x *Location) GetLat() float64 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Location) GetLon() float64 {
	if x != nil {
		return x.Lon
	}
	return 0
}

type Coordinate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float64 `protobuf:"fixed64,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Coordinate) Reset() {
	*x = Coordinate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_streaming_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinate) ProtoMessage() {}

func (x *Coordinate) ProtoReflect() protoreflect.Message {
	mi := &file_video_streaming_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coordinate.ProtoReflect.Descriptor instead.
func (*Coordinate) Descriptor() ([]byte, []int) {
	return file_video_streaming_proto_rawDescGZIP(), []int{3}
}

func (x *Coordinate) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Coordinate) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Coordinate) GetZ() float64 {
	if x != nil {
		return x.Z
	}
	return 0
}

type BoudingBox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Top    int32 `protobuf:"varint,1,opt,name=top,proto3" json:"top,omitempty"`
	Left   int32 `protobuf:"varint,2,opt,name=left,proto3" json:"left,omitempty"`
	Width  int32 `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	Height int32 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *BoudingBox) Reset() {
	*x = BoudingBox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_streaming_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoudingBox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoudingBox) ProtoMessage() {}

func (x *BoudingBox) ProtoReflect() protoreflect.Message {
	mi := &file_video_streaming_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoudingBox.ProtoReflect.Descriptor instead.
func (*BoudingBox) Descriptor() ([]byte, []int) {
	return file_video_streaming_proto_rawDescGZIP(), []int{4}
}

func (x *BoudingBox) GetTop() int32 {
	if x != nil {
		return x.Top
	}
	return 0
}

func (x *BoudingBox) GetLeft() int32 {
	if x != nil {
		return x.Left
	}
	return 0
}

func (x *BoudingBox) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *BoudingBox) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

// Video Streaming messages
type ShapeProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dim []*ShapeProto_Dim `protobuf:"bytes,2,rep,name=dim,proto3" json:"dim,omitempty"`
}

func (x *ShapeProto) Reset() {
	*x = ShapeProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_streaming_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShapeProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShapeProto) ProtoMessage() {}

func (x *ShapeProto) ProtoReflect() protoreflect.Message {
	mi := &file_video_streaming_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShapeProto.ProtoReflect.Descriptor instead.
func (*ShapeProto) Descriptor() ([]byte, []int) {
	return file_video_streaming_proto_rawDescGZIP(), []int{5}
}

func (x *ShapeProto) GetDim() []*ShapeProto_Dim {
	if x != nil {
		return x.Dim
	}
	return nil
}

type VideoFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width      int64       `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height     int64       `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Data       []byte      `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Timestamp  int64       `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	IsKeyframe bool        `protobuf:"varint,5,opt,name=is_keyframe,json=isKeyframe,proto3" json:"is_keyframe,omitempty"`
	Pts        int64       `protobuf:"varint,6,opt,name=pts,proto3" json:"pts,omitempty"`
	Dts        int64       `protobuf:"varint,7,opt,name=dts,proto3" json:"dts,omitempty"`
	FrameType  string      `protobuf:"bytes,8,opt,name=frame_type,json=frameType,proto3" json:"frame_type,omitempty"`
	IsCorrupt  bool        `protobuf:"varint,9,opt,name=is_corrupt,json=isCorrupt,proto3" json:"is_corrupt,omitempty"`
	TimeBase   float64     `protobuf:"fixed64,10,opt,name=time_base,json=timeBase,proto3" json:"time_base,omitempty"`
	Shape      *ShapeProto `protobuf:"bytes,11,opt,name=shape,proto3" json:"shape,omitempty"`
	DeviceId   string      `protobuf:"bytes,12,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Packet     int64       `protobuf:"varint,13,opt,name=packet,proto3" json:"packet,omitempty"`
	Keyframe   int64       `protobuf:"varint,14,opt,name=keyframe,proto3" json:"keyframe,omitempty"`
	Extradata  []byte      `protobuf:"bytes,15,opt,name=extradata,proto3" json:"extradata,omitempty"`
	CodecName  string      `protobuf:"bytes,16,opt,name=codec_name,json=codecName,proto3" json:"codec_name,omitempty"`
	PixFmt     string      `protobuf:"bytes,17,opt,name=pix_fmt,json=pixFmt,proto3" json:"pix_fmt,omitempty"`
}

func (x *VideoFrame) Reset() {
	*x = VideoFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_streaming_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoFrame) ProtoMessage() {}

func (x *VideoFrame) ProtoReflect() protoreflect.Message {
	mi := &file_video_streaming_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoFrame.ProtoReflect.Descriptor instead.
func (*VideoFrame) Descriptor() ([]byte, []int) {
	return file_video_streaming_proto_rawDescGZIP(), []int{6}
}

func (x *VideoFrame) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *VideoFrame) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *VideoFrame) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *VideoFrame) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *VideoFrame) GetIsKeyframe() bool {
	if x != nil {
		return x.IsKeyframe
	}
	return false
}

func (x *VideoFrame) GetPts() int64 {
	if x != nil {
		return x.Pts
	}
	return 0
}

func (x *VideoFrame) GetDts() int64 {
	if x != nil {
		return x.Dts
	}
	return 0
}

func (x *VideoFrame) GetFrameType() string {
	if x != nil {
		return x.FrameType
	}
	return ""
}

func (x *VideoFrame) GetIsCorrupt() bool {
	if x != nil {
		return x.IsCorrupt
	}
	return false
}

func (x *VideoFrame) GetTimeBase() float64 {
	if x != nil {
		return x.TimeBase
	}
	return 0
}

func (x *VideoFrame) GetShape() *ShapeProto {
	if x != nil {
		return x.Shape
	}
	return nil
}

func (x *VideoFrame) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *VideoFrame) GetPacket() int64 {
	if x != nil {
		return x.Packet
	}
	return 0
}

func (x *VideoFrame) GetKeyframe() int64 {
	if x != nil {
		return x.Keyframe
	}
	return 0
}

func (x *VideoFrame) GetExtradata() []byte {
	if x != nil {
		return x.Extradata
	}
	return nil
}

func (x *VideoFrame) GetCodecName() string {
	if x != nil {
		return x.CodecName
	}
	return ""
}

func (x *VideoFrame) GetPixFmt() string {
	if x != nil {
		return x.PixFmt
	}
	return ""
}

type VideoFrameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyFrameOnly bool   `protobuf:"varint,1,opt,name=key_frame_only,json=keyFrameOnly,proto3" json:"key_frame_only,omitempty"`
	DeviceId     string `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *VideoFrameRequest) Reset() {
	*x = VideoFrameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_streaming_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoFrameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoFrameRequest) ProtoMessage() {}

func (x *VideoFrameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_streaming_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoFrameRequest.ProtoReflect.Descriptor instead.
func (*VideoFrameRequest) Descriptor() ([]byte, []int) {
	return file_video_streaming_proto_rawDescGZIP(), []int{7}
}

func (x *VideoFrameRequest) GetKeyFrameOnly() bool {
	if x != nil {
		return x.KeyFrameOnly
	}
	return false
}

func (x *VideoFrameRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type VideoFrameBufferedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId      string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	TimestampFrom int64  `protobuf:"varint,2,opt,name=timestamp_from,json=timestampFrom,proto3" json:"timestamp_from,omitempty"`
	TimestampTo   int64  `protobuf:"varint,3,opt,name=timestamp_to,json=timestampTo,proto3" json:"timestamp_to,omitempty"`
}

func (x *VideoFrameBufferedRequest) Reset() {
	*x = VideoFrameBufferedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_streaming_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoFrameBufferedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoFrameBufferedRequest) ProtoMessage() {}

func (x *VideoFrameBufferedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_streaming_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoFrameBufferedRequest.ProtoReflect.Descriptor instead.
func (*VideoFrameBufferedRequest) Descriptor() ([]byte, []int) {
	return file_video_streaming_proto_rawDescGZIP(), []int{8}
}

func (x *VideoFrameBufferedRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *VideoFrameBufferedRequest) GetTimestampFrom() int64 {
	if x != nil {
		return x.TimestampFrom
	}
	return 0
}

func (x *VideoFrameBufferedRequest) GetTimestampTo() int64 {
	if x != nil {
		return x.TimestampTo
	}
	return 0
}

// ListStream messages
type ListStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status        string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	FailingStreak int64  `protobuf:"varint,3,opt,name=failing_streak,json=failingStreak,proto3" json:"failing_streak,omitempty"`
	HealthStatus  string `protobuf:"bytes,4,opt,name=health_status,json=healthStatus,proto3" json:"health_status,omitempty"`
	Dead          bool   `protobuf:"varint,5,opt,name=dead,proto3" json:"dead,omitempty"`
	ExitCode      int64  `protobuf:"varint,6,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	Pid           int32  `protobuf:"varint,7,opt,name=pid,proto3" json:"pid,omitempty"`
	Running       bool   `protobuf:"varint,8,opt,name=running,proto3" json:"running,omitempty"`
	Paused        bool   `protobuf:"varint,9,opt,name=paused,proto3" json:"paused,omitempty"`
	Restarting    bool   `protobuf:"varint,10,opt,name=restarting,proto3" json:"restarting,omitempty"`
	Oomkilled     bool   `protobuf:"varint,11,opt,name=oomkilled,proto3" json:"oomkilled,omitempty"`
	Error         string `protobuf:"bytes,12,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ListStream) Reset() {
	*x = ListStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_streaming_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStream) ProtoMessage() {}

func (x *ListStream) ProtoReflect() protoreflect.Message {
	mi := &file_video_streaming_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStream.ProtoReflect.Descriptor instead.
func (*ListStream) Descriptor() ([]byte, []int) {
	return file_video_streaming_proto_rawDescGZIP(), []int{9}
}

func (x *ListStream) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListStream) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ListStream) GetFailingStreak() int64 {
	if x != nil {
		return x.FailingStreak
	}
	return 0
}

func (x *ListStream) GetHealthStatus() string {
	if x != nil {
		return x.HealthStatus
	}
	return ""
}

func (x *ListStream) GetDead() bool {
	if x != nil {
		return x.Dead
	}
	return false
}

func (x *ListStream) GetExitCode() int64 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *ListStream) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *ListStream) GetRunning() bool {
	if x != nil {
		return x.Running
	}
	return false
}

func (x *ListStream) GetPaused() bool {
	if x != nil {
		return x.Paused
	}
	return false
}

func (x *ListStream) GetRestarting() bool {
	if x != nil {
		return x.Restarting
	}
	return false
}

func (x *ListStream) GetOomkilled() bool {
	if x != nil {
		return x.Oomkilled
	}
	return false
}

func (x *ListStream) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ListStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListStreamRequest) Reset() {
	*x = ListStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_streaming_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStreamRequest) ProtoMessage() {}

func (x *ListStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_streaming_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStreamRequest.ProtoReflect.Descriptor instead.
func (*ListStreamRequest) Descriptor() ([]byte, []int) {
	return file_video_streaming_proto_rawDescGZIP(), []int{10}
}

// Proxy messages
type ProxyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId    string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Passthrough bool   `protobuf:"varint,2,opt,name=passthrough,proto3" json:"passthrough,omitempty"` // true = passthrough streaming, false = stop passthrough streaming
}

func (x *ProxyRequest) Reset() {
	*x = ProxyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_streaming_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyRequest) ProtoMessage() {}

func (x *ProxyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_streaming_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyRequest.ProtoReflect.Descriptor instead.
func (*ProxyRequest) Descriptor() ([]byte, []int) {
	return file_video_streaming_proto_rawDescGZIP(), []int{11}
}

func (x *ProxyRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *ProxyRequest) GetPassthrough() bool {
	if x != nil {
		return x.Passthrough
	}
	return false
}

type ProxyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId    string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Passthrough bool   `protobuf:"varint,2,opt,name=passthrough,proto3" json:"passthrough,omitempty"`
}

func (x *ProxyResponse) Reset() {
	*x = ProxyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_streaming_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyResponse) ProtoMessage() {}

func (x *ProxyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_streaming_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyResponse.ProtoReflect.Descriptor instead.
func (*ProxyResponse) Descriptor() ([]byte, []int) {
	return file_video_streaming_proto_rawDescGZIP(), []int{12}
}

func (x *ProxyResponse) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *ProxyResponse) GetPassthrough() bool {
	if x != nil {
		return x.Passthrough
	}
	return false
}

// Storage messages
type StorageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Start    bool   `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
}
