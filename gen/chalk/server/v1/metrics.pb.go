// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: chalk/server/v1/metrics.proto

package serverv1

import (
	_ "github.com/chalk-ai/chalk-go/v2/gen/chalk/auth/v1"
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

type GetOverviewSummaryMetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RangeStart *string `protobuf:"bytes,1,opt,name=range_start,json=rangeStart,proto3,oneof" json:"range_start,omitempty"`
	RangeEnd   *string `protobuf:"bytes,2,opt,name=range_end,json=rangeEnd,proto3,oneof" json:"range_end,omitempty"`
}

func (x *GetOverviewSummaryMetricsRequest) Reset() {
	*x = GetOverviewSummaryMetricsRequest{}
	mi := &file_chalk_server_v1_metrics_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOverviewSummaryMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOverviewSummaryMetricsRequest) ProtoMessage() {}

func (x *GetOverviewSummaryMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_metrics_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOverviewSummaryMetricsRequest.ProtoReflect.Descriptor instead.
func (*GetOverviewSummaryMetricsRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *GetOverviewSummaryMetricsRequest) GetRangeStart() string {
	if x != nil && x.RangeStart != nil {
		return *x.RangeStart
	}
	return ""
}

func (x *GetOverviewSummaryMetricsRequest) GetRangeEnd() string {
	if x != nil && x.RangeEnd != nil {
		return *x.RangeEnd
	}
	return ""
}

type OverviewSummaryMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value float32 `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *OverviewSummaryMetric) Reset() {
	*x = OverviewSummaryMetric{}
	mi := &file_chalk_server_v1_metrics_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OverviewSummaryMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OverviewSummaryMetric) ProtoMessage() {}

func (x *OverviewSummaryMetric) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_metrics_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OverviewSummaryMetric.ProtoReflect.Descriptor instead.
func (*OverviewSummaryMetric) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_metrics_proto_rawDescGZIP(), []int{1}
}

func (x *OverviewSummaryMetric) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OverviewSummaryMetric) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type GetOverviewSummaryMetricsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*OverviewSummaryMetric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *GetOverviewSummaryMetricsResponse) Reset() {
	*x = GetOverviewSummaryMetricsResponse{}
	mi := &file_chalk_server_v1_metrics_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOverviewSummaryMetricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOverviewSummaryMetricsResponse) ProtoMessage() {}

func (x *GetOverviewSummaryMetricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_metrics_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOverviewSummaryMetricsResponse.ProtoReflect.Descriptor instead.
func (*GetOverviewSummaryMetricsResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_metrics_proto_rawDescGZIP(), []int{2}
}

func (x *GetOverviewSummaryMetricsResponse) GetMetrics() []*OverviewSummaryMetric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

var File_chalk_server_v1_metrics_proto protoreflect.FileDescriptor

var file_chalk_server_v1_metrics_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x88, 0x01, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65,
	0x77, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x08, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x22, 0x41, 0x0a, 0x15,
	0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x65, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x07, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x32, 0x9a, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x87, 0x01, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x31, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x76, 0x65,
	0x72, 0x76, 0x69, 0x65, 0x77, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03,
	0x80, 0x7d, 0x06, 0x42, 0xbc, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69,
	0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02, 0x0f, 0x43,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x1b, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x11, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_server_v1_metrics_proto_rawDescOnce sync.Once
	file_chalk_server_v1_metrics_proto_rawDescData = file_chalk_server_v1_metrics_proto_rawDesc
)

func file_chalk_server_v1_metrics_proto_rawDescGZIP() []byte {
	file_chalk_server_v1_metrics_proto_rawDescOnce.Do(func() {
		file_chalk_server_v1_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_server_v1_metrics_proto_rawDescData)
	})
	return file_chalk_server_v1_metrics_proto_rawDescData
}

var file_chalk_server_v1_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chalk_server_v1_metrics_proto_goTypes = []any{
	(*GetOverviewSummaryMetricsRequest)(nil),  // 0: chalk.server.v1.GetOverviewSummaryMetricsRequest
	(*OverviewSummaryMetric)(nil),             // 1: chalk.server.v1.OverviewSummaryMetric
	(*GetOverviewSummaryMetricsResponse)(nil), // 2: chalk.server.v1.GetOverviewSummaryMetricsResponse
}
var file_chalk_server_v1_metrics_proto_depIdxs = []int32{
	1, // 0: chalk.server.v1.GetOverviewSummaryMetricsResponse.metrics:type_name -> chalk.server.v1.OverviewSummaryMetric
	0, // 1: chalk.server.v1.MetricsService.GetOverviewSummaryMetrics:input_type -> chalk.server.v1.GetOverviewSummaryMetricsRequest
	2, // 2: chalk.server.v1.MetricsService.GetOverviewSummaryMetrics:output_type -> chalk.server.v1.GetOverviewSummaryMetricsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chalk_server_v1_metrics_proto_init() }
func file_chalk_server_v1_metrics_proto_init() {
	if File_chalk_server_v1_metrics_proto != nil {
		return
	}
	file_chalk_server_v1_metrics_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_server_v1_metrics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chalk_server_v1_metrics_proto_goTypes,
		DependencyIndexes: file_chalk_server_v1_metrics_proto_depIdxs,
		MessageInfos:      file_chalk_server_v1_metrics_proto_msgTypes,
	}.Build()
	File_chalk_server_v1_metrics_proto = out.File
	file_chalk_server_v1_metrics_proto_rawDesc = nil
	file_chalk_server_v1_metrics_proto_goTypes = nil
	file_chalk_server_v1_metrics_proto_depIdxs = nil
}
