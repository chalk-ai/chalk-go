// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: chalk/metrics/v4/metrics.proto

package metricsv4

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

type MetricType int32

const (
	MetricType_METRIC_TYPE_UNSPECIFIED     MetricType = 0
	MetricType_METRIC_TYPE_COUNTER         MetricType = 1
	MetricType_METRIC_TYPE_UP_DOWN_COUNTER MetricType = 2
	MetricType_METRIC_TYPE_HISTOGRAM       MetricType = 3
)

// Enum value maps for MetricType.
var (
	MetricType_name = map[int32]string{
		0: "METRIC_TYPE_UNSPECIFIED",
		1: "METRIC_TYPE_COUNTER",
		2: "METRIC_TYPE_UP_DOWN_COUNTER",
		3: "METRIC_TYPE_HISTOGRAM",
	}
	MetricType_value = map[string]int32{
		"METRIC_TYPE_UNSPECIFIED":     0,
		"METRIC_TYPE_COUNTER":         1,
		"METRIC_TYPE_UP_DOWN_COUNTER": 2,
		"METRIC_TYPE_HISTOGRAM":       3,
	}
)

func (x MetricType) Enum() *MetricType {
	p := new(MetricType)
	*p = x
	return p
}

func (x MetricType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricType) Descriptor() protoreflect.EnumDescriptor {
	return file_chalk_metrics_v4_metrics_proto_enumTypes[0].Descriptor()
}

func (MetricType) Type() protoreflect.EnumType {
	return &file_chalk_metrics_v4_metrics_proto_enumTypes[0]
}

func (x MetricType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricType.Descriptor instead.
func (MetricType) EnumDescriptor() ([]byte, []int) {
	return file_chalk_metrics_v4_metrics_proto_rawDescGZIP(), []int{0}
}

type MetricNamespace int32

const (
	MetricNamespace_METRIC_NAMESPACE_UNSPECIFIED MetricNamespace = 0
	MetricNamespace_METRIC_NAMESPACE_FEATURE     MetricNamespace = 1
	MetricNamespace_METRIC_NAMESPACE_SYSTEM      MetricNamespace = 2
)

// Enum value maps for MetricNamespace.
var (
	MetricNamespace_name = map[int32]string{
		0: "METRIC_NAMESPACE_UNSPECIFIED",
		1: "METRIC_NAMESPACE_FEATURE",
		2: "METRIC_NAMESPACE_SYSTEM",
	}
	MetricNamespace_value = map[string]int32{
		"METRIC_NAMESPACE_UNSPECIFIED": 0,
		"METRIC_NAMESPACE_FEATURE":     1,
		"METRIC_NAMESPACE_SYSTEM":      2,
	}
)

func (x MetricNamespace) Enum() *MetricNamespace {
	p := new(MetricNamespace)
	*p = x
	return p
}

func (x MetricNamespace) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricNamespace) Descriptor() protoreflect.EnumDescriptor {
	return file_chalk_metrics_v4_metrics_proto_enumTypes[1].Descriptor()
}

func (MetricNamespace) Type() protoreflect.EnumType {
	return &file_chalk_metrics_v4_metrics_proto_enumTypes[1]
}

func (x MetricNamespace) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricNamespace.Descriptor instead.
func (MetricNamespace) EnumDescriptor() ([]byte, []int) {
	return file_chalk_metrics_v4_metrics_proto_rawDescGZIP(), []int{1}
}

type MetricName int32

const (
	MetricName_METRIC_NAME_UNSPECIFIED               MetricName = 0
	MetricName_METRIC_NAME_DEFAULT_VALUES_USED       MetricName = 1
	MetricName_METRIC_NAME_GIVENS_VALUES_USED        MetricName = 2
	MetricName_METRIC_NAME_OFFLINE_STORE_VALUES_USED MetricName = 3
	MetricName_METRIC_NAME_ONLINE_STORE_VALUES_USED  MetricName = 4
	MetricName_METRIC_NAME_RESOLVER_VALUES_USED      MetricName = 5
	MetricName_METRIC_NAME_NULL_COUNT                MetricName = 6
	MetricName_METRIC_NAME_NON_NULL_COUNT            MetricName = 7
	MetricName_METRIC_NAME_MISSING_COUNT             MetricName = 8
	MetricName_METRIC_NAME_ZERO_COUNT                MetricName = 9
	MetricName_METRIC_NAME_VALUES_DISTRIBUTION       MetricName = 10
	MetricName_METRIC_NAME_RESOLVER_LATENCY          MetricName = 11
	MetricName_METRIC_NAME_BYTES_DOWNLOADED          MetricName = 12
	MetricName_METRIC_NAME_NUM_ROWS                  MetricName = 13
	MetricName_METRIC_NAME_NUM_FEATURES              MetricName = 14
)

// Enum value maps for MetricName.
var (
	MetricName_name = map[int32]string{
		0:  "METRIC_NAME_UNSPECIFIED",
		1:  "METRIC_NAME_DEFAULT_VALUES_USED",
		2:  "METRIC_NAME_GIVENS_VALUES_USED",
		3:  "METRIC_NAME_OFFLINE_STORE_VALUES_USED",
		4:  "METRIC_NAME_ONLINE_STORE_VALUES_USED",
		5:  "METRIC_NAME_RESOLVER_VALUES_USED",
		6:  "METRIC_NAME_NULL_COUNT",
		7:  "METRIC_NAME_NON_NULL_COUNT",
		8:  "METRIC_NAME_MISSING_COUNT",
		9:  "METRIC_NAME_ZERO_COUNT",
		10: "METRIC_NAME_VALUES_DISTRIBUTION",
		11: "METRIC_NAME_RESOLVER_LATENCY",
		12: "METRIC_NAME_BYTES_DOWNLOADED",
		13: "METRIC_NAME_NUM_ROWS",
		14: "METRIC_NAME_NUM_FEATURES",
	}
	MetricName_value = map[string]int32{
		"METRIC_NAME_UNSPECIFIED":               0,
		"METRIC_NAME_DEFAULT_VALUES_USED":       1,
		"METRIC_NAME_GIVENS_VALUES_USED":        2,
		"METRIC_NAME_OFFLINE_STORE_VALUES_USED": 3,
		"METRIC_NAME_ONLINE_STORE_VALUES_USED":  4,
		"METRIC_NAME_RESOLVER_VALUES_USED":      5,
		"METRIC_NAME_NULL_COUNT":                6,
		"METRIC_NAME_NON_NULL_COUNT":            7,
		"METRIC_NAME_MISSING_COUNT":             8,
		"METRIC_NAME_ZERO_COUNT":                9,
		"METRIC_NAME_VALUES_DISTRIBUTION":       10,
		"METRIC_NAME_RESOLVER_LATENCY":          11,
		"METRIC_NAME_BYTES_DOWNLOADED":          12,
		"METRIC_NAME_NUM_ROWS":                  13,
		"METRIC_NAME_NUM_FEATURES":              14,
	}
)

func (x MetricName) Enum() *MetricName {
	p := new(MetricName)
	*p = x
	return p
}

func (x MetricName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricName) Descriptor() protoreflect.EnumDescriptor {
	return file_chalk_metrics_v4_metrics_proto_enumTypes[2].Descriptor()
}

func (MetricName) Type() protoreflect.EnumType {
	return &file_chalk_metrics_v4_metrics_proto_enumTypes[2]
}

func (x MetricName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricName.Descriptor instead.
func (MetricName) EnumDescriptor() ([]byte, []int) {
	return file_chalk_metrics_v4_metrics_proto_rawDescGZIP(), []int{2}
}

type MetricBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*Metric `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *MetricBatch) Reset() {
	*x = MetricBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_metrics_v4_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricBatch) ProtoMessage() {}

func (x *MetricBatch) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_metrics_v4_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricBatch.ProtoReflect.Descriptor instead.
func (*MetricBatch) Descriptor() ([]byte, []int) {
	return file_chalk_metrics_v4_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *MetricBatch) GetMetrics() []*Metric {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type Metric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// optional because we may not have an id when we're doing the initial publish
	Id            *int64          `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	EnvironmentId string          `protobuf:"bytes,2,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	Namespace     MetricNamespace `protobuf:"varint,3,opt,name=namespace,proto3,enum=chalk.metrics.v4.MetricNamespace" json:"namespace,omitempty"`
	Source        string          `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Name          MetricName      `protobuf:"varint,5,opt,name=name,proto3,enum=chalk.metrics.v4.MetricName" json:"name,omitempty"`
	Type          MetricType      `protobuf:"varint,6,opt,name=type,proto3,enum=chalk.metrics.v4.MetricType" json:"type,omitempty"`
	DeploymentId  string          `protobuf:"bytes,7,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	// values
	//
	// Types that are assignable to Value:
	//
	//	*Metric_Sketch
	//	*Metric_Numeric
	Value          isMetric_Value         `protobuf_oneof:"value"`
	Statistics     *MetricStatistics      `protobuf:"bytes,12,opt,name=statistics,proto3,oneof" json:"statistics,omitempty"`
	ObservedAt     *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=observed_at,json=observedAt,proto3,oneof" json:"observed_at,omitempty"`
	AdditionalTags map[string]string      `protobuf:"bytes,30,rep,name=additional_tags,json=additionalTags,proto3" json:"additional_tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Metric) Reset() {
	*x = Metric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_metrics_v4_metrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metric) ProtoMessage() {}

func (x *Metric) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_metrics_v4_metrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metric.ProtoReflect.Descriptor instead.
func (*Metric) Descriptor() ([]byte, []int) {
	return file_chalk_metrics_v4_metrics_proto_rawDescGZIP(), []int{1}
}

func (x *Metric) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *Metric) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *Metric) GetNamespace() MetricNamespace {
	if x != nil {
		return x.Namespace
	}
	return MetricNamespace_METRIC_NAMESPACE_UNSPECIFIED
}

func (x *Metric) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Metric) GetName() MetricName {
	if x != nil {
		return x.Name
	}
	return MetricName_METRIC_NAME_UNSPECIFIED
}

func (x *Metric) GetType() MetricType {
	if x != nil {
		return x.Type
	}
	return MetricType_METRIC_TYPE_UNSPECIFIED
}

func (x *Metric) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (m *Metric) GetValue() isMetric_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Metric) GetSketch() string {
	if x, ok := x.GetValue().(*Metric_Sketch); ok {
		return x.Sketch
	}
	return ""
}

func (x *Metric) GetNumeric() float64 {
	if x, ok := x.GetValue().(*Metric_Numeric); ok {
		return x.Numeric
	}
	return 0
}

func (x *Metric) GetStatistics() *MetricStatistics {
	if x != nil {
		return x.Statistics
	}
	return nil
}

func (x *Metric) GetObservedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ObservedAt
	}
	return nil
}

func (x *Metric) GetAdditionalTags() map[string]string {
	if x != nil {
		return x.AdditionalTags
	}
	return nil
}

type isMetric_Value interface {
	isMetric_Value()
}

type Metric_Sketch struct {
	Sketch string `protobuf:"bytes,10,opt,name=sketch,proto3,oneof"`
}

type Metric_Numeric struct {
	Numeric float64 `protobuf:"fixed64,11,opt,name=numeric,proto3,oneof"`
}

func (*Metric_Sketch) isMetric_Value() {}

func (*Metric_Numeric) isMetric_Value() {}

type MetricStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min       *float64 `protobuf:"fixed64,1,opt,name=min,proto3,oneof" json:"min,omitempty"`
	Max       *float64 `protobuf:"fixed64,2,opt,name=max,proto3,oneof" json:"max,omitempty"`
	Count     *uint64  `protobuf:"varint,3,opt,name=count,proto3,oneof" json:"count,omitempty"`
	NullCount *uint64  `protobuf:"varint,4,opt,name=null_count,json=nullCount,proto3,oneof" json:"null_count,omitempty"`
	Mean      *float64 `protobuf:"fixed64,5,opt,name=mean,proto3,oneof" json:"mean,omitempty"`
	Variance  *float64 `protobuf:"fixed64,6,opt,name=variance,proto3,oneof" json:"variance,omitempty"`
}

func (x *MetricStatistics) Reset() {
	*x = MetricStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_metrics_v4_metrics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricStatistics) ProtoMessage() {}

func (x *MetricStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_metrics_v4_metrics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricStatistics.ProtoReflect.Descriptor instead.
func (*MetricStatistics) Descriptor() ([]byte, []int) {
	return file_chalk_metrics_v4_metrics_proto_rawDescGZIP(), []int{2}
}

func (x *MetricStatistics) GetMin() float64 {
	if x != nil && x.Min != nil {
		return *x.Min
	}
	return 0
}

func (x *MetricStatistics) GetMax() float64 {
	if x != nil && x.Max != nil {
		return *x.Max
	}
	return 0
}

func (x *MetricStatistics) GetCount() uint64 {
	if x != nil && x.Count != nil {
		return *x.Count
	}
	return 0
}

func (x *MetricStatistics) GetNullCount() uint64 {
	if x != nil && x.NullCount != nil {
		return *x.NullCount
	}
	return 0
}

func (x *MetricStatistics) GetMean() float64 {
	if x != nil && x.Mean != nil {
		return *x.Mean
	}
	return 0
}

func (x *MetricStatistics) GetVariance() float64 {
	if x != nil && x.Variance != nil {
		return *x.Variance
	}
	return 0
}

var File_chalk_metrics_v4_metrics_proto protoreflect.FileDescriptor

var file_chalk_metrics_v4_metrics_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f,
	0x76, 0x34, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x76, 0x34, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x12, 0x32, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x07, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0xb0, 0x05, 0x0a, 0x06, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52,
	0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3f, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x21, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x2e, 0x76, 0x34, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x06, 0x73, 0x6b, 0x65, 0x74, 0x63, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x06, 0x73, 0x6b, 0x65, 0x74, 0x63, 0x68, 0x12, 0x1a, 0x0a, 0x07, 0x6e, 0x75, 0x6d,
	0x65, 0x72, 0x69, 0x63, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x07, 0x6e, 0x75,
	0x6d, 0x65, 0x72, 0x69, 0x63, 0x12, 0x47, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x48, 0x02, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x88, 0x01, 0x01, 0x12, 0x40,
	0x0a, 0x0b, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48,
	0x03, 0x52, 0x0a, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x55, 0x0a, 0x0f, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x1e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x34, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x61,
	0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x54, 0x61, 0x67, 0x73, 0x1a, 0x41, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0xf8, 0x01, 0x0a, 0x10, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x15,
	0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x03, 0x6d,
	0x69, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x48, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x48, 0x02, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x6e, 0x75, 0x6c, 0x6c, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x48, 0x03, 0x52, 0x09, 0x6e,
	0x75, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6d,
	0x65, 0x61, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x48, 0x04, 0x52, 0x04, 0x6d, 0x65, 0x61,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x48, 0x05, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e,
	0x63, 0x65, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x69, 0x6e, 0x42, 0x06, 0x0a,
	0x04, 0x5f, 0x6d, 0x61, 0x78, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x6e, 0x75, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x6d, 0x65, 0x61, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x76, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0x2a, 0x7e, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x17, 0x0a, 0x13, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43,
	0x4f, 0x55, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x4d, 0x45, 0x54, 0x52,
	0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x50, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x5f,
	0x43, 0x4f, 0x55, 0x4e, 0x54, 0x45, 0x52, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x45, 0x54,
	0x52, 0x49, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x47, 0x52,
	0x41, 0x4d, 0x10, 0x03, 0x2a, 0x6e, 0x0a, 0x0f, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x1c, 0x4d, 0x45, 0x54, 0x52, 0x49,
	0x43, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x4d, 0x45, 0x54,
	0x52, 0x49, 0x43, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x46, 0x45,
	0x41, 0x54, 0x55, 0x52, 0x45, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x45, 0x54, 0x52, 0x49,
	0x43, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50, 0x41, 0x43, 0x45, 0x5f, 0x53, 0x59, 0x53, 0x54,
	0x45, 0x4d, 0x10, 0x02, 0x2a, 0x85, 0x04, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x4e, 0x41,
	0x4d, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x23, 0x0a, 0x1f, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f,
	0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x53, 0x5f, 0x55,
	0x53, 0x45, 0x44, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f,
	0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x47, 0x49, 0x56, 0x45, 0x4e, 0x53, 0x5f, 0x56, 0x41, 0x4c, 0x55,
	0x45, 0x53, 0x5f, 0x55, 0x53, 0x45, 0x44, 0x10, 0x02, 0x12, 0x29, 0x0a, 0x25, 0x4d, 0x45, 0x54,
	0x52, 0x49, 0x43, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45,
	0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x53, 0x5f, 0x55, 0x53,
	0x45, 0x44, 0x10, 0x03, 0x12, 0x28, 0x0a, 0x24, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x4e,
	0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45,
	0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x53, 0x5f, 0x55, 0x53, 0x45, 0x44, 0x10, 0x04, 0x12, 0x24,
	0x0a, 0x20, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x52, 0x45,
	0x53, 0x4f, 0x4c, 0x56, 0x45, 0x52, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x53, 0x5f, 0x55, 0x53,
	0x45, 0x44, 0x10, 0x05, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x4e,
	0x41, 0x4d, 0x45, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x06,
	0x12, 0x1e, 0x0a, 0x1a, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f,
	0x4e, 0x4f, 0x4e, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x07,
	0x12, 0x1d, 0x0a, 0x19, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f,
	0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x08, 0x12,
	0x1a, 0x0a, 0x16, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x5a,
	0x45, 0x52, 0x4f, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x09, 0x12, 0x23, 0x0a, 0x1f, 0x4d,
	0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x53, 0x5f, 0x44, 0x49, 0x53, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a,
	0x12, 0x20, 0x0a, 0x1c, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f,
	0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x52, 0x5f, 0x4c, 0x41, 0x54, 0x45, 0x4e, 0x43, 0x59,
	0x10, 0x0b, 0x12, 0x20, 0x0a, 0x1c, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x4e, 0x41, 0x4d,
	0x45, 0x5f, 0x42, 0x59, 0x54, 0x45, 0x53, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44,
	0x45, 0x44, 0x10, 0x0c, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x4e,
	0x41, 0x4d, 0x45, 0x5f, 0x4e, 0x55, 0x4d, 0x5f, 0x52, 0x4f, 0x57, 0x53, 0x10, 0x0d, 0x12, 0x1c,
	0x0a, 0x18, 0x4d, 0x45, 0x54, 0x52, 0x49, 0x43, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4e, 0x55,
	0x4d, 0x5f, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x53, 0x10, 0x0e, 0x42, 0xc3, 0x01, 0x0a,
	0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2e, 0x76, 0x34, 0x42, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x34, 0x3b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x76, 0x34, 0xa2, 0x02, 0x03, 0x43, 0x4d, 0x58, 0xaa, 0x02, 0x10, 0x43, 0x68, 0x61, 0x6c, 0x6b,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x56, 0x34, 0xca, 0x02, 0x10, 0x43, 0x68,
	0x61, 0x6c, 0x6b, 0x5c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5c, 0x56, 0x34, 0xe2, 0x02,
	0x1c, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5c, 0x56,
	0x34, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12,
	0x43, 0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x3a, 0x3a,
	0x56, 0x34, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_metrics_v4_metrics_proto_rawDescOnce sync.Once
	file_chalk_metrics_v4_metrics_proto_rawDescData = file_chalk_metrics_v4_metrics_proto_rawDesc
)

func file_chalk_metrics_v4_metrics_proto_rawDescGZIP() []byte {
	file_chalk_metrics_v4_metrics_proto_rawDescOnce.Do(func() {
		file_chalk_metrics_v4_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_metrics_v4_metrics_proto_rawDescData)
	})
	return file_chalk_metrics_v4_metrics_proto_rawDescData
}

var file_chalk_metrics_v4_metrics_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_chalk_metrics_v4_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chalk_metrics_v4_metrics_proto_goTypes = []any{
	(MetricType)(0),               // 0: chalk.metrics.v4.MetricType
	(MetricNamespace)(0),          // 1: chalk.metrics.v4.MetricNamespace
	(MetricName)(0),               // 2: chalk.metrics.v4.MetricName
	(*MetricBatch)(nil),           // 3: chalk.metrics.v4.MetricBatch
	(*Metric)(nil),                // 4: chalk.metrics.v4.Metric
	(*MetricStatistics)(nil),      // 5: chalk.metrics.v4.MetricStatistics
	nil,                           // 6: chalk.metrics.v4.Metric.AdditionalTagsEntry
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_chalk_metrics_v4_metrics_proto_depIdxs = []int32{
	4, // 0: chalk.metrics.v4.MetricBatch.metrics:type_name -> chalk.metrics.v4.Metric
	1, // 1: chalk.metrics.v4.Metric.namespace:type_name -> chalk.metrics.v4.MetricNamespace
	2, // 2: chalk.metrics.v4.Metric.name:type_name -> chalk.metrics.v4.MetricName
	0, // 3: chalk.metrics.v4.Metric.type:type_name -> chalk.metrics.v4.MetricType
	5, // 4: chalk.metrics.v4.Metric.statistics:type_name -> chalk.metrics.v4.MetricStatistics
	7, // 5: chalk.metrics.v4.Metric.observed_at:type_name -> google.protobuf.Timestamp
	6, // 6: chalk.metrics.v4.Metric.additional_tags:type_name -> chalk.metrics.v4.Metric.AdditionalTagsEntry
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_chalk_metrics_v4_metrics_proto_init() }
func file_chalk_metrics_v4_metrics_proto_init() {
	if File_chalk_metrics_v4_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chalk_metrics_v4_metrics_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MetricBatch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chalk_metrics_v4_metrics_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Metric); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chalk_metrics_v4_metrics_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*MetricStatistics); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_chalk_metrics_v4_metrics_proto_msgTypes[1].OneofWrappers = []any{
		(*Metric_Sketch)(nil),
		(*Metric_Numeric)(nil),
	}
	file_chalk_metrics_v4_metrics_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_metrics_v4_metrics_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chalk_metrics_v4_metrics_proto_goTypes,
		DependencyIndexes: file_chalk_metrics_v4_metrics_proto_depIdxs,
		EnumInfos:         file_chalk_metrics_v4_metrics_proto_enumTypes,
		MessageInfos:      file_chalk_metrics_v4_metrics_proto_msgTypes,
	}.Build()
	File_chalk_metrics_v4_metrics_proto = out.File
	file_chalk_metrics_v4_metrics_proto_rawDesc = nil
	file_chalk_metrics_v4_metrics_proto_goTypes = nil
	file_chalk_metrics_v4_metrics_proto_depIdxs = nil
}
