// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: chalk/artifacts/v1/cron_query.proto

package artifactsv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type RecomputeSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeatureFqns []string `protobuf:"bytes,1,rep,name=feature_fqns,json=featureFqns,proto3" json:"feature_fqns,omitempty"`
	AllFeatures bool     `protobuf:"varint,2,opt,name=all_features,json=allFeatures,proto3" json:"all_features,omitempty"`
}

func (x *RecomputeSettings) Reset() {
	*x = RecomputeSettings{}
	mi := &file_chalk_artifacts_v1_cron_query_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecomputeSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecomputeSettings) ProtoMessage() {}

func (x *RecomputeSettings) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_artifacts_v1_cron_query_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecomputeSettings.ProtoReflect.Descriptor instead.
func (*RecomputeSettings) Descriptor() ([]byte, []int) {
	return file_chalk_artifacts_v1_cron_query_proto_rawDescGZIP(), []int{0}
}

func (x *RecomputeSettings) GetFeatureFqns() []string {
	if x != nil {
		return x.FeatureFqns
	}
	return nil
}

func (x *RecomputeSettings) GetAllFeatures() bool {
	if x != nil {
		return x.AllFeatures
	}
	return false
}

type CronQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                 string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cron                 string                 `protobuf:"bytes,2,opt,name=cron,proto3" json:"cron,omitempty"`
	FileName             string                 `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Output               []string               `protobuf:"bytes,4,rep,name=output,proto3" json:"output,omitempty"`
	MaxSamples           *int64                 `protobuf:"varint,5,opt,name=max_samples,json=maxSamples,proto3,oneof" json:"max_samples,omitempty"`
	Recompute            *RecomputeSettings     `protobuf:"bytes,6,opt,name=recompute,proto3" json:"recompute,omitempty"`
	LowerBound           *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=lower_bound,json=lowerBound,proto3" json:"lower_bound,omitempty"`
	UpperBound           *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=upper_bound,json=upperBound,proto3" json:"upper_bound,omitempty"`
	Tags                 []string               `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"`
	RequiredResolverTags []string               `protobuf:"bytes,10,rep,name=required_resolver_tags,json=requiredResolverTags,proto3" json:"required_resolver_tags,omitempty"`
	StoreOnline          bool                   `protobuf:"varint,11,opt,name=store_online,json=storeOnline,proto3" json:"store_online,omitempty"`
	StoreOffline         bool                   `protobuf:"varint,12,opt,name=store_offline,json=storeOffline,proto3" json:"store_offline,omitempty"`
	IncrementalSources   []string               `protobuf:"bytes,13,rep,name=incremental_sources,json=incrementalSources,proto3" json:"incremental_sources,omitempty"`
}

func (x *CronQuery) Reset() {
	*x = CronQuery{}
	mi := &file_chalk_artifacts_v1_cron_query_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CronQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CronQuery) ProtoMessage() {}

func (x *CronQuery) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_artifacts_v1_cron_query_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CronQuery.ProtoReflect.Descriptor instead.
func (*CronQuery) Descriptor() ([]byte, []int) {
	return file_chalk_artifacts_v1_cron_query_proto_rawDescGZIP(), []int{1}
}

func (x *CronQuery) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CronQuery) GetCron() string {
	if x != nil {
		return x.Cron
	}
	return ""
}

func (x *CronQuery) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *CronQuery) GetOutput() []string {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *CronQuery) GetMaxSamples() int64 {
	if x != nil && x.MaxSamples != nil {
		return *x.MaxSamples
	}
	return 0
}

func (x *CronQuery) GetRecompute() *RecomputeSettings {
	if x != nil {
		return x.Recompute
	}
	return nil
}

func (x *CronQuery) GetLowerBound() *timestamppb.Timestamp {
	if x != nil {
		return x.LowerBound
	}
	return nil
}

func (x *CronQuery) GetUpperBound() *timestamppb.Timestamp {
	if x != nil {
		return x.UpperBound
	}
	return nil
}

func (x *CronQuery) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *CronQuery) GetRequiredResolverTags() []string {
	if x != nil {
		return x.RequiredResolverTags
	}
	return nil
}

func (x *CronQuery) GetStoreOnline() bool {
	if x != nil {
		return x.StoreOnline
	}
	return false
}

func (x *CronQuery) GetStoreOffline() bool {
	if x != nil {
		return x.StoreOffline
	}
	return false
}

func (x *CronQuery) GetIncrementalSources() []string {
	if x != nil {
		return x.IncrementalSources
	}
	return nil
}

var File_chalk_artifacts_v1_cron_query_proto protoreflect.FileDescriptor

var file_chalk_artifacts_v1_cron_query_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x11, 0x52, 0x65,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x66, 0x71, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x71,
	0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x6c, 0x6c, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x61, 0x6c, 0x6c, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0xa0, 0x04, 0x0a, 0x09, 0x43, 0x72, 0x6f, 0x6e, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x72, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x72, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x24, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x43, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x70, 0x65,
	0x72, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x70, 0x65, 0x72,
	0x42, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x5f, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x54, 0x61, 0x67, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x6e, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6f, 0x66, 0x66, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x4f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x69, 0x6e, 0x63, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x0d,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x69, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6d, 0x61, 0x78,
	0x5f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x42, 0xd3, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x42, 0x0e, 0x43, 0x72, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x41, 0x58, 0xaa, 0x02, 0x12, 0x43,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x12, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x41,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x3a,
	0x3a, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_artifacts_v1_cron_query_proto_rawDescOnce sync.Once
	file_chalk_artifacts_v1_cron_query_proto_rawDescData = file_chalk_artifacts_v1_cron_query_proto_rawDesc
)

func file_chalk_artifacts_v1_cron_query_proto_rawDescGZIP() []byte {
	file_chalk_artifacts_v1_cron_query_proto_rawDescOnce.Do(func() {
		file_chalk_artifacts_v1_cron_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_artifacts_v1_cron_query_proto_rawDescData)
	})
	return file_chalk_artifacts_v1_cron_query_proto_rawDescData
}

var file_chalk_artifacts_v1_cron_query_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chalk_artifacts_v1_cron_query_proto_goTypes = []any{
	(*RecomputeSettings)(nil),     // 0: chalk.artifacts.v1.RecomputeSettings
	(*CronQuery)(nil),             // 1: chalk.artifacts.v1.CronQuery
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_chalk_artifacts_v1_cron_query_proto_depIdxs = []int32{
	0, // 0: chalk.artifacts.v1.CronQuery.recompute:type_name -> chalk.artifacts.v1.RecomputeSettings
	2, // 1: chalk.artifacts.v1.CronQuery.lower_bound:type_name -> google.protobuf.Timestamp
	2, // 2: chalk.artifacts.v1.CronQuery.upper_bound:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_chalk_artifacts_v1_cron_query_proto_init() }
func file_chalk_artifacts_v1_cron_query_proto_init() {
	if File_chalk_artifacts_v1_cron_query_proto != nil {
		return
	}
	file_chalk_artifacts_v1_cron_query_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_artifacts_v1_cron_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chalk_artifacts_v1_cron_query_proto_goTypes,
		DependencyIndexes: file_chalk_artifacts_v1_cron_query_proto_depIdxs,
		MessageInfos:      file_chalk_artifacts_v1_cron_query_proto_msgTypes,
	}.Build()
	File_chalk_artifacts_v1_cron_query_proto = out.File
	file_chalk_artifacts_v1_cron_query_proto_rawDesc = nil
	file_chalk_artifacts_v1_cron_query_proto_goTypes = nil
	file_chalk_artifacts_v1_cron_query_proto_depIdxs = nil
}
