// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: chalk/engine/v1/query_server.proto

package enginev1

import (
	v11 "github.com/chalk-ai/chalk-go/gen/chalk/aggregate/v1"
	_ "github.com/chalk-ai/chalk-go/gen/chalk/auth/v1"
	v1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
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

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	mi := &file_chalk_engine_v1_query_server_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_engine_v1_query_server_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_chalk_engine_v1_query_server_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type PingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num int32 `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	mi := &file_chalk_engine_v1_query_server_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_engine_v1_query_server_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_chalk_engine_v1_query_server_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_chalk_engine_v1_query_server_proto protoreflect.FileDescriptor

var file_chalk_engine_v1_query_server_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1f, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x6e, 0x75, 0x6d, 0x22, 0x20, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x32, 0xe4, 0x06, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12,
	0x1c, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x80, 0x7d,
	0x01, 0x90, 0x02, 0x01, 0x12, 0x5d, 0x0a, 0x0b, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x23, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03,
	0x80, 0x7d, 0x03, 0x12, 0x69, 0x0a, 0x0f, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x42, 0x75, 0x6c, 0x6b, 0x12, 0x27, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x42, 0x75, 0x6c, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x75, 0x6c,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x80, 0x7d, 0x03, 0x12, 0x6c,
	0x0a, 0x10, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x12, 0x28, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x80, 0x7d, 0x03, 0x12, 0x72, 0x0a, 0x12,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x42, 0x75,
	0x6c, 0x6b, 0x12, 0x2a, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x42, 0x75, 0x6c, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b,
	0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x42,
	0x75, 0x6c, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x80, 0x7d, 0x03,
	0x12, 0x66, 0x0a, 0x0e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x12, 0x26, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x03, 0x80, 0x7d, 0x03, 0x12, 0x84, 0x01, 0x0a, 0x15, 0x50, 0x6c, 0x61,
	0x6e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69,
	0x6c, 0x6c, 0x12, 0x30, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x80, 0x7d, 0x0c, 0x90, 0x02, 0x01, 0x12,
	0x6c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73,
	0x12, 0x28, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x80, 0x7d, 0x0b, 0x90, 0x02, 0x01, 0x42, 0xc0, 0x01,
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x45, 0x58, 0xaa, 0x02, 0x0f, 0x43, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x43,
	0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1b, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x43,
	0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_engine_v1_query_server_proto_rawDescOnce sync.Once
	file_chalk_engine_v1_query_server_proto_rawDescData = file_chalk_engine_v1_query_server_proto_rawDesc
)

func file_chalk_engine_v1_query_server_proto_rawDescGZIP() []byte {
	file_chalk_engine_v1_query_server_proto_rawDescOnce.Do(func() {
		file_chalk_engine_v1_query_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_engine_v1_query_server_proto_rawDescData)
	})
	return file_chalk_engine_v1_query_server_proto_rawDescData
}

var file_chalk_engine_v1_query_server_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chalk_engine_v1_query_server_proto_goTypes = []any{
	(*PingRequest)(nil),                       // 0: chalk.engine.v1.PingRequest
	(*PingResponse)(nil),                      // 1: chalk.engine.v1.PingResponse
	(*v1.OnlineQueryRequest)(nil),             // 2: chalk.common.v1.OnlineQueryRequest
	(*v1.OnlineQueryBulkRequest)(nil),         // 3: chalk.common.v1.OnlineQueryBulkRequest
	(*v1.OnlineQueryMultiRequest)(nil),        // 4: chalk.common.v1.OnlineQueryMultiRequest
	(*v1.UploadFeaturesBulkRequest)(nil),      // 5: chalk.common.v1.UploadFeaturesBulkRequest
	(*v1.UploadFeaturesRequest)(nil),          // 6: chalk.common.v1.UploadFeaturesRequest
	(*v11.PlanAggregateBackfillRequest)(nil),  // 7: chalk.aggregate.v1.PlanAggregateBackfillRequest
	(*v11.GetAggregatesRequest)(nil),          // 8: chalk.aggregate.v1.GetAggregatesRequest
	(*v1.OnlineQueryResponse)(nil),            // 9: chalk.common.v1.OnlineQueryResponse
	(*v1.OnlineQueryBulkResponse)(nil),        // 10: chalk.common.v1.OnlineQueryBulkResponse
	(*v1.OnlineQueryMultiResponse)(nil),       // 11: chalk.common.v1.OnlineQueryMultiResponse
	(*v1.UploadFeaturesBulkResponse)(nil),     // 12: chalk.common.v1.UploadFeaturesBulkResponse
	(*v1.UploadFeaturesResponse)(nil),         // 13: chalk.common.v1.UploadFeaturesResponse
	(*v11.PlanAggregateBackfillResponse)(nil), // 14: chalk.aggregate.v1.PlanAggregateBackfillResponse
	(*v11.GetAggregatesResponse)(nil),         // 15: chalk.aggregate.v1.GetAggregatesResponse
}
var file_chalk_engine_v1_query_server_proto_depIdxs = []int32{
	0,  // 0: chalk.engine.v1.QueryService.Ping:input_type -> chalk.engine.v1.PingRequest
	2,  // 1: chalk.engine.v1.QueryService.OnlineQuery:input_type -> chalk.common.v1.OnlineQueryRequest
	3,  // 2: chalk.engine.v1.QueryService.OnlineQueryBulk:input_type -> chalk.common.v1.OnlineQueryBulkRequest
	4,  // 3: chalk.engine.v1.QueryService.OnlineQueryMulti:input_type -> chalk.common.v1.OnlineQueryMultiRequest
	5,  // 4: chalk.engine.v1.QueryService.UploadFeaturesBulk:input_type -> chalk.common.v1.UploadFeaturesBulkRequest
	6,  // 5: chalk.engine.v1.QueryService.UploadFeatures:input_type -> chalk.common.v1.UploadFeaturesRequest
	7,  // 6: chalk.engine.v1.QueryService.PlanAggregateBackfill:input_type -> chalk.aggregate.v1.PlanAggregateBackfillRequest
	8,  // 7: chalk.engine.v1.QueryService.GetAggregates:input_type -> chalk.aggregate.v1.GetAggregatesRequest
	1,  // 8: chalk.engine.v1.QueryService.Ping:output_type -> chalk.engine.v1.PingResponse
	9,  // 9: chalk.engine.v1.QueryService.OnlineQuery:output_type -> chalk.common.v1.OnlineQueryResponse
	10, // 10: chalk.engine.v1.QueryService.OnlineQueryBulk:output_type -> chalk.common.v1.OnlineQueryBulkResponse
	11, // 11: chalk.engine.v1.QueryService.OnlineQueryMulti:output_type -> chalk.common.v1.OnlineQueryMultiResponse
	12, // 12: chalk.engine.v1.QueryService.UploadFeaturesBulk:output_type -> chalk.common.v1.UploadFeaturesBulkResponse
	13, // 13: chalk.engine.v1.QueryService.UploadFeatures:output_type -> chalk.common.v1.UploadFeaturesResponse
	14, // 14: chalk.engine.v1.QueryService.PlanAggregateBackfill:output_type -> chalk.aggregate.v1.PlanAggregateBackfillResponse
	15, // 15: chalk.engine.v1.QueryService.GetAggregates:output_type -> chalk.aggregate.v1.GetAggregatesResponse
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_chalk_engine_v1_query_server_proto_init() }
func file_chalk_engine_v1_query_server_proto_init() {
	if File_chalk_engine_v1_query_server_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_engine_v1_query_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chalk_engine_v1_query_server_proto_goTypes,
		DependencyIndexes: file_chalk_engine_v1_query_server_proto_depIdxs,
		MessageInfos:      file_chalk_engine_v1_query_server_proto_msgTypes,
	}.Build()
	File_chalk_engine_v1_query_server_proto = out.File
	file_chalk_engine_v1_query_server_proto_rawDesc = nil
	file_chalk_engine_v1_query_server_proto_goTypes = nil
	file_chalk_engine_v1_query_server_proto_depIdxs = nil
}
