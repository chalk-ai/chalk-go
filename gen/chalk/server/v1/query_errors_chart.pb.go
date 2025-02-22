// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: chalk/server/v1/query_errors_chart.proto

package serverv1

import (
	v1 "github.com/chalk-ai/chalk-go/gen/chalk/chart/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetQueryErrorsChartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartTimestampInclusive *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=start_timestamp_inclusive,json=startTimestampInclusive,proto3" json:"start_timestamp_inclusive,omitempty"`
	// If not specified, assumes the current timestamp
	// To avoid awkward small window buckets at the end, specify this.
	EndTimestampExclusive *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=end_timestamp_exclusive,json=endTimestampExclusive,proto3,oneof" json:"end_timestamp_exclusive,omitempty"`
	// The size of each time bucket + how far apart time points are in the chart
	// Ideally a divisor of the total window between end and start ms
	// If not, the bucket start-aligned with end_timestamp_exclusive will be of size total_window % window_period
	WindowPeriod *durationpb.Duration `protobuf:"bytes,3,opt,name=window_period,json=windowPeriod,proto3" json:"window_period,omitempty"`
	Filters      *QueryErrorFilters   `protobuf:"bytes,4,opt,name=filters,proto3" json:"filters,omitempty"`
}

func (x *GetQueryErrorsChartRequest) Reset() {
	*x = GetQueryErrorsChartRequest{}
	mi := &file_chalk_server_v1_query_errors_chart_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetQueryErrorsChartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQueryErrorsChartRequest) ProtoMessage() {}

func (x *GetQueryErrorsChartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_query_errors_chart_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQueryErrorsChartRequest.ProtoReflect.Descriptor instead.
func (*GetQueryErrorsChartRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_query_errors_chart_proto_rawDescGZIP(), []int{0}
}

func (x *GetQueryErrorsChartRequest) GetStartTimestampInclusive() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTimestampInclusive
	}
	return nil
}

func (x *GetQueryErrorsChartRequest) GetEndTimestampExclusive() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTimestampExclusive
	}
	return nil
}

func (x *GetQueryErrorsChartRequest) GetWindowPeriod() *durationpb.Duration {
	if x != nil {
		return x.WindowPeriod
	}
	return nil
}

func (x *GetQueryErrorsChartRequest) GetFilters() *QueryErrorFilters {
	if x != nil {
		return x.Filters
	}
	return nil
}

type GetQueryErrorsChartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chart *v1.DenseTimeSeriesChart `protobuf:"bytes,1,opt,name=chart,proto3" json:"chart,omitempty"`
}

func (x *GetQueryErrorsChartResponse) Reset() {
	*x = GetQueryErrorsChartResponse{}
	mi := &file_chalk_server_v1_query_errors_chart_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetQueryErrorsChartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQueryErrorsChartResponse) ProtoMessage() {}

func (x *GetQueryErrorsChartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_query_errors_chart_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQueryErrorsChartResponse.ProtoReflect.Descriptor instead.
func (*GetQueryErrorsChartResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_query_errors_chart_proto_rawDescGZIP(), []int{1}
}

func (x *GetQueryErrorsChartResponse) GetChart() *v1.DenseTimeSeriesChart {
	if x != nil {
		return x.Chart
	}
	return nil
}

var File_chalk_server_v1_query_errors_chart_proto protoreflect.FileDescriptor

var file_chalk_server_v1_query_errors_chart_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x5f, 0x63,
	0x68, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x29, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x6e, 0x73,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x63, 0x68, 0x61, 0x72, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x02, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x43, 0x68, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x56, 0x0a, 0x19, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x17, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76,
	0x65, 0x12, 0x57, 0x0a, 0x17, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x5f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00,
	0x52, 0x15, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x45, 0x78,
	0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x0d, 0x77, 0x69,
	0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x77, 0x69,
	0x6e, 0x64, 0x6f, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x3c, 0x0a, 0x07, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52,
	0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x42, 0x1a, 0x0a, 0x18, 0x5f, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x65, 0x78, 0x63, 0x6c, 0x75,
	0x73, 0x69, 0x76, 0x65, 0x22, 0x59, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x05, 0x63, 0x68, 0x61, 0x72, 0x74, 0x42,
	0xc5, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x51, 0x75, 0x65, 0x72, 0x79, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x43, 0x68, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53,
	0x58, 0xaa, 0x02, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_server_v1_query_errors_chart_proto_rawDescOnce sync.Once
	file_chalk_server_v1_query_errors_chart_proto_rawDescData = file_chalk_server_v1_query_errors_chart_proto_rawDesc
)

func file_chalk_server_v1_query_errors_chart_proto_rawDescGZIP() []byte {
	file_chalk_server_v1_query_errors_chart_proto_rawDescOnce.Do(func() {
		file_chalk_server_v1_query_errors_chart_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_server_v1_query_errors_chart_proto_rawDescData)
	})
	return file_chalk_server_v1_query_errors_chart_proto_rawDescData
}

var file_chalk_server_v1_query_errors_chart_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chalk_server_v1_query_errors_chart_proto_goTypes = []any{
	(*GetQueryErrorsChartRequest)(nil),  // 0: chalk.server.v1.GetQueryErrorsChartRequest
	(*GetQueryErrorsChartResponse)(nil), // 1: chalk.server.v1.GetQueryErrorsChartResponse
	(*timestamppb.Timestamp)(nil),       // 2: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),         // 3: google.protobuf.Duration
	(*QueryErrorFilters)(nil),           // 4: chalk.server.v1.QueryErrorFilters
	(*v1.DenseTimeSeriesChart)(nil),     // 5: chalk.chart.v1.DenseTimeSeriesChart
}
var file_chalk_server_v1_query_errors_chart_proto_depIdxs = []int32{
	2, // 0: chalk.server.v1.GetQueryErrorsChartRequest.start_timestamp_inclusive:type_name -> google.protobuf.Timestamp
	2, // 1: chalk.server.v1.GetQueryErrorsChartRequest.end_timestamp_exclusive:type_name -> google.protobuf.Timestamp
	3, // 2: chalk.server.v1.GetQueryErrorsChartRequest.window_period:type_name -> google.protobuf.Duration
	4, // 3: chalk.server.v1.GetQueryErrorsChartRequest.filters:type_name -> chalk.server.v1.QueryErrorFilters
	5, // 4: chalk.server.v1.GetQueryErrorsChartResponse.chart:type_name -> chalk.chart.v1.DenseTimeSeriesChart
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_chalk_server_v1_query_errors_chart_proto_init() }
func file_chalk_server_v1_query_errors_chart_proto_init() {
	if File_chalk_server_v1_query_errors_chart_proto != nil {
		return
	}
	file_chalk_server_v1_query_error_proto_init()
	file_chalk_server_v1_query_errors_chart_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_server_v1_query_errors_chart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chalk_server_v1_query_errors_chart_proto_goTypes,
		DependencyIndexes: file_chalk_server_v1_query_errors_chart_proto_depIdxs,
		MessageInfos:      file_chalk_server_v1_query_errors_chart_proto_msgTypes,
	}.Build()
	File_chalk_server_v1_query_errors_chart_proto = out.File
	file_chalk_server_v1_query_errors_chart_proto_rawDesc = nil
	file_chalk_server_v1_query_errors_chart_proto_goTypes = nil
	file_chalk_server_v1_query_errors_chart_proto_depIdxs = nil
}
