// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: chalk/aggregate/v1/timeseries.proto

package aggregatev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AggregateTimeSeriesRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Aggregation    string               `protobuf:"bytes,1,opt,name=aggregation,proto3" json:"aggregation,omitempty"`
	BucketDuration *durationpb.Duration `protobuf:"bytes,2,opt,name=bucket_duration,json=bucketDuration,proto3" json:"bucket_duration,omitempty"`
	// The features which depend on this rule.
	DependentFeatures []string             `protobuf:"bytes,3,rep,name=dependent_features,json=dependentFeatures,proto3" json:"dependent_features,omitempty"`
	Retention         *durationpb.Duration `protobuf:"bytes,4,opt,name=retention,proto3" json:"retention,omitempty"`
	DatetimeFeature   string               `protobuf:"bytes,5,opt,name=datetime_feature,json=datetimeFeature,proto3" json:"datetime_feature,omitempty"`
}

func (x *AggregateTimeSeriesRule) Reset() {
	*x = AggregateTimeSeriesRule{}
	mi := &file_chalk_aggregate_v1_timeseries_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AggregateTimeSeriesRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateTimeSeriesRule) ProtoMessage() {}

func (x *AggregateTimeSeriesRule) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_aggregate_v1_timeseries_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateTimeSeriesRule.ProtoReflect.Descriptor instead.
func (*AggregateTimeSeriesRule) Descriptor() ([]byte, []int) {
	return file_chalk_aggregate_v1_timeseries_proto_rawDescGZIP(), []int{0}
}

func (x *AggregateTimeSeriesRule) GetAggregation() string {
	if x != nil {
		return x.Aggregation
	}
	return ""
}

func (x *AggregateTimeSeriesRule) GetBucketDuration() *durationpb.Duration {
	if x != nil {
		return x.BucketDuration
	}
	return nil
}

func (x *AggregateTimeSeriesRule) GetDependentFeatures() []string {
	if x != nil {
		return x.DependentFeatures
	}
	return nil
}

func (x *AggregateTimeSeriesRule) GetRetention() *durationpb.Duration {
	if x != nil {
		return x.Retention
	}
	return nil
}

func (x *AggregateTimeSeriesRule) GetDatetimeFeature() string {
	if x != nil {
		return x.DatetimeFeature
	}
	return ""
}

type AggregateTimeSeries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace          string                     `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	AggregateOn        string                     `protobuf:"bytes,2,opt,name=aggregate_on,json=aggregateOn,proto3" json:"aggregate_on,omitempty"`
	GroupBy            []string                   `protobuf:"bytes,3,rep,name=group_by,json=groupBy,proto3" json:"group_by,omitempty"`
	Rules              []*AggregateTimeSeriesRule `protobuf:"bytes,5,rep,name=rules,proto3" json:"rules,omitempty"`
	FiltersDescription string                     `protobuf:"bytes,6,opt,name=filters_description,json=filtersDescription,proto3" json:"filters_description,omitempty"`
	BucketFeature      string                     `protobuf:"bytes,7,opt,name=bucket_feature,json=bucketFeature,proto3" json:"bucket_feature,omitempty"`
}

func (x *AggregateTimeSeries) Reset() {
	*x = AggregateTimeSeries{}
	mi := &file_chalk_aggregate_v1_timeseries_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AggregateTimeSeries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateTimeSeries) ProtoMessage() {}

func (x *AggregateTimeSeries) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_aggregate_v1_timeseries_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateTimeSeries.ProtoReflect.Descriptor instead.
func (*AggregateTimeSeries) Descriptor() ([]byte, []int) {
	return file_chalk_aggregate_v1_timeseries_proto_rawDescGZIP(), []int{1}
}

func (x *AggregateTimeSeries) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *AggregateTimeSeries) GetAggregateOn() string {
	if x != nil {
		return x.AggregateOn
	}
	return ""
}

func (x *AggregateTimeSeries) GetGroupBy() []string {
	if x != nil {
		return x.GroupBy
	}
	return nil
}

func (x *AggregateTimeSeries) GetRules() []*AggregateTimeSeriesRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *AggregateTimeSeries) GetFiltersDescription() string {
	if x != nil {
		return x.FiltersDescription
	}
	return ""
}

func (x *AggregateTimeSeries) GetBucketFeature() string {
	if x != nil {
		return x.BucketFeature
	}
	return ""
}

var File_chalk_aggregate_v1_timeseries_proto protoreflect.FileDescriptor

var file_chalk_aggregate_v1_timeseries_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x02, 0x0a, 0x17, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x0f, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x12, 0x64,
	0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65,
	0x6e, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x09, 0x72, 0x65,
	0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64,
	0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x8c,
	0x02, 0x0a, 0x13, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x5f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x62, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x42, 0x79, 0x12, 0x41, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0xd4, 0x01,
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69,
	0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43,
	0x41, 0x58, 0xaa, 0x02, 0x12, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x43,
	0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14,
	0x43, 0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_aggregate_v1_timeseries_proto_rawDescOnce sync.Once
	file_chalk_aggregate_v1_timeseries_proto_rawDescData = file_chalk_aggregate_v1_timeseries_proto_rawDesc
)

func file_chalk_aggregate_v1_timeseries_proto_rawDescGZIP() []byte {
	file_chalk_aggregate_v1_timeseries_proto_rawDescOnce.Do(func() {
		file_chalk_aggregate_v1_timeseries_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_aggregate_v1_timeseries_proto_rawDescData)
	})
	return file_chalk_aggregate_v1_timeseries_proto_rawDescData
}

var file_chalk_aggregate_v1_timeseries_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chalk_aggregate_v1_timeseries_proto_goTypes = []any{
	(*AggregateTimeSeriesRule)(nil), // 0: chalk.aggregate.v1.AggregateTimeSeriesRule
	(*AggregateTimeSeries)(nil),     // 1: chalk.aggregate.v1.AggregateTimeSeries
	(*durationpb.Duration)(nil),     // 2: google.protobuf.Duration
}
var file_chalk_aggregate_v1_timeseries_proto_depIdxs = []int32{
	2, // 0: chalk.aggregate.v1.AggregateTimeSeriesRule.bucket_duration:type_name -> google.protobuf.Duration
	2, // 1: chalk.aggregate.v1.AggregateTimeSeriesRule.retention:type_name -> google.protobuf.Duration
	0, // 2: chalk.aggregate.v1.AggregateTimeSeries.rules:type_name -> chalk.aggregate.v1.AggregateTimeSeriesRule
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_chalk_aggregate_v1_timeseries_proto_init() }
func file_chalk_aggregate_v1_timeseries_proto_init() {
	if File_chalk_aggregate_v1_timeseries_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_aggregate_v1_timeseries_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chalk_aggregate_v1_timeseries_proto_goTypes,
		DependencyIndexes: file_chalk_aggregate_v1_timeseries_proto_depIdxs,
		MessageInfos:      file_chalk_aggregate_v1_timeseries_proto_msgTypes,
	}.Build()
	File_chalk_aggregate_v1_timeseries_proto = out.File
	file_chalk_aggregate_v1_timeseries_proto_rawDesc = nil
	file_chalk_aggregate_v1_timeseries_proto_goTypes = nil
	file_chalk_aggregate_v1_timeseries_proto_depIdxs = nil
}
