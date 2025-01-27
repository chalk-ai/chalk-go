// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: chalk/manager/v1/status.proto

package managerv1

import (
	_ "github.com/chalk-ai/chalk-go/gen/chalk/auth/v1"
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

type GetClusterMetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetClusterMetricsRequest) Reset() {
	*x = GetClusterMetricsRequest{}
	mi := &file_chalk_manager_v1_status_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetClusterMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterMetricsRequest) ProtoMessage() {}

func (x *GetClusterMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_manager_v1_status_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterMetricsRequest.ProtoReflect.Descriptor instead.
func (*GetClusterMetricsRequest) Descriptor() ([]byte, []int) {
	return file_chalk_manager_v1_status_proto_rawDescGZIP(), []int{0}
}

type GetClusterMetricsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics string `protobuf:"bytes,1,opt,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *GetClusterMetricsResponse) Reset() {
	*x = GetClusterMetricsResponse{}
	mi := &file_chalk_manager_v1_status_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetClusterMetricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterMetricsResponse) ProtoMessage() {}

func (x *GetClusterMetricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_manager_v1_status_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterMetricsResponse.ProtoReflect.Descriptor instead.
func (*GetClusterMetricsResponse) Descriptor() ([]byte, []int) {
	return file_chalk_manager_v1_status_proto_rawDescGZIP(), []int{1}
}

func (x *GetClusterMetricsResponse) GetMetrics() string {
	if x != nil {
		return x.Metrics
	}
	return ""
}

var File_chalk_manager_v1_status_proto protoreflect.FileDescriptor

var file_chalk_manager_v1_status_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x10, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x35,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x32, 0x85, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x74, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x2a, 0x2e, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x80, 0x7d, 0x0a, 0x90, 0x02, 0x01, 0x42, 0xc2, 0x01,
	0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x4d, 0x58, 0xaa, 0x02, 0x10, 0x43, 0x68, 0x61, 0x6c, 0x6b,
	0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x43, 0x68,
	0x61, 0x6c, 0x6b, 0x5c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1c, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12,
	0x43, 0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_manager_v1_status_proto_rawDescOnce sync.Once
	file_chalk_manager_v1_status_proto_rawDescData = file_chalk_manager_v1_status_proto_rawDesc
)

func file_chalk_manager_v1_status_proto_rawDescGZIP() []byte {
	file_chalk_manager_v1_status_proto_rawDescOnce.Do(func() {
		file_chalk_manager_v1_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_manager_v1_status_proto_rawDescData)
	})
	return file_chalk_manager_v1_status_proto_rawDescData
}

var file_chalk_manager_v1_status_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chalk_manager_v1_status_proto_goTypes = []any{
	(*GetClusterMetricsRequest)(nil),  // 0: chalk.manager.v1.GetClusterMetricsRequest
	(*GetClusterMetricsResponse)(nil), // 1: chalk.manager.v1.GetClusterMetricsResponse
}
var file_chalk_manager_v1_status_proto_depIdxs = []int32{
	0, // 0: chalk.manager.v1.StatusService.GetClusterMetrics:input_type -> chalk.manager.v1.GetClusterMetricsRequest
	1, // 1: chalk.manager.v1.StatusService.GetClusterMetrics:output_type -> chalk.manager.v1.GetClusterMetricsResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chalk_manager_v1_status_proto_init() }
func file_chalk_manager_v1_status_proto_init() {
	if File_chalk_manager_v1_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_manager_v1_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chalk_manager_v1_status_proto_goTypes,
		DependencyIndexes: file_chalk_manager_v1_status_proto_depIdxs,
		MessageInfos:      file_chalk_manager_v1_status_proto_msgTypes,
	}.Build()
	File_chalk_manager_v1_status_proto = out.File
	file_chalk_manager_v1_status_proto_rawDesc = nil
	file_chalk_manager_v1_status_proto_goTypes = nil
	file_chalk_manager_v1_status_proto_depIdxs = nil
}
