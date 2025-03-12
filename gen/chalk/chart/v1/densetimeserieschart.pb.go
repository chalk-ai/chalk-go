// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: chalk/chart/v1/densetimeserieschart.proto

package chartv1

import (
	v1 "github.com/chalk-ai/chalk-go/v2/gen/chalk/arrow/v1"
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

type DensePoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// wrapped in a Point to allow for optional (empty space in time series)
	Value *float64 `protobuf:"fixed64,1,opt,name=value,proto3,oneof" json:"value,omitempty"`
}

func (x *DensePoint) Reset() {
	*x = DensePoint{}
	mi := &file_chalk_chart_v1_densetimeserieschart_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DensePoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DensePoint) ProtoMessage() {}

func (x *DensePoint) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_chart_v1_densetimeserieschart_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DensePoint.ProtoReflect.Descriptor instead.
func (*DensePoint) Descriptor() ([]byte, []int) {
	return file_chalk_chart_v1_densetimeserieschart_proto_rawDescGZIP(), []int{0}
}

func (x *DensePoint) GetValue() float64 {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return 0
}

// If this series was created as part of a group-by(s)
// This stores extra information about which ones and what value it pertains to
type GroupTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupKey string          `protobuf:"bytes,1,opt,name=group_key,json=groupKey,proto3" json:"group_key,omitempty"`
	Value    *v1.ScalarValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GroupTag) Reset() {
	*x = GroupTag{}
	mi := &file_chalk_chart_v1_densetimeserieschart_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GroupTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupTag) ProtoMessage() {}

func (x *GroupTag) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_chart_v1_densetimeserieschart_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupTag.ProtoReflect.Descriptor instead.
func (*GroupTag) Descriptor() ([]byte, []int) {
	return file_chalk_chart_v1_densetimeserieschart_proto_rawDescGZIP(), []int{1}
}

func (x *GroupTag) GetGroupKey() string {
	if x != nil {
		return x.GroupKey
	}
	return ""
}

func (x *GroupTag) GetValue() *v1.ScalarValue {
	if x != nil {
		return x.Value
	}
	return nil
}

type DenseTimeSeries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Points    []*DensePoint `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
	Label     string        `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Unit      string        `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
	GroupTags []*GroupTag   `protobuf:"bytes,4,rep,name=group_tags,json=groupTags,proto3" json:"group_tags,omitempty"`
}

func (x *DenseTimeSeries) Reset() {
	*x = DenseTimeSeries{}
	mi := &file_chalk_chart_v1_densetimeserieschart_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DenseTimeSeries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DenseTimeSeries) ProtoMessage() {}

func (x *DenseTimeSeries) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_chart_v1_densetimeserieschart_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DenseTimeSeries.ProtoReflect.Descriptor instead.
func (*DenseTimeSeries) Descriptor() ([]byte, []int) {
	return file_chalk_chart_v1_densetimeserieschart_proto_rawDescGZIP(), []int{2}
}

func (x *DenseTimeSeries) GetPoints() []*DensePoint {
	if x != nil {
		return x.Points
	}
	return nil
}

func (x *DenseTimeSeries) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *DenseTimeSeries) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *DenseTimeSeries) GetGroupTags() []*GroupTag {
	if x != nil {
		return x.GroupTags
	}
	return nil
}

// *
// A representation of a densely encoded chart, suitable for time series where the series
// have mostly non-null values.
// Consider implementing a sparse chart if not every series has data at every tick
type DenseTimeSeriesChart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string                   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Series       []*DenseTimeSeries       `protobuf:"bytes,2,rep,name=series,proto3" json:"series,omitempty"`
	XSeries      []*timestamppb.Timestamp `protobuf:"bytes,3,rep,name=x_series,json=xSeries,proto3" json:"x_series,omitempty"`
	WindowPeriod *durationpb.Duration     `protobuf:"bytes,4,opt,name=window_period,json=windowPeriod,proto3" json:"window_period,omitempty"`
}

func (x *DenseTimeSeriesChart) Reset() {
	*x = DenseTimeSeriesChart{}
	mi := &file_chalk_chart_v1_densetimeserieschart_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DenseTimeSeriesChart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DenseTimeSeriesChart) ProtoMessage() {}

func (x *DenseTimeSeriesChart) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_chart_v1_densetimeserieschart_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DenseTimeSeriesChart.ProtoReflect.Descriptor instead.
func (*DenseTimeSeriesChart) Descriptor() ([]byte, []int) {
	return file_chalk_chart_v1_densetimeserieschart_proto_rawDescGZIP(), []int{3}
}

func (x *DenseTimeSeriesChart) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *DenseTimeSeriesChart) GetSeries() []*DenseTimeSeries {
	if x != nil {
		return x.Series
	}
	return nil
}

func (x *DenseTimeSeriesChart) GetXSeries() []*timestamppb.Timestamp {
	if x != nil {
		return x.XSeries
	}
	return nil
}

func (x *DenseTimeSeriesChart) GetWindowPeriod() *durationpb.Duration {
	if x != nil {
		return x.WindowPeriod
	}
	return nil
}

var File_chalk_chart_v1_densetimeserieschart_proto protoreflect.FileDescriptor

var file_chalk_chart_v1_densetimeserieschart_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x65, 0x6e, 0x73, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2f, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x72, 0x6f,
	0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x0a, 0x44, 0x65, 0x6e, 0x73,
	0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5a, 0x0a, 0x08, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x54, 0x61, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x4b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x72, 0x72, 0x6f,
	0x77, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa8, 0x01, 0x0a, 0x0f, 0x44, 0x65, 0x6e, 0x73,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x06, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6e,
	0x73, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x37, 0x0a, 0x0a, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x54, 0x61, 0x67, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x54, 0x61,
	0x67, 0x73, 0x22, 0xdc, 0x01, 0x0a, 0x14, 0x44, 0x65, 0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x78, 0x5f,
	0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x78, 0x53, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x3e, 0x0a, 0x0d, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x42, 0xc2, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x19, 0x44, 0x65, 0x6e, 0x73, 0x65, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x63, 0x68, 0x61, 0x72, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x63, 0x68,
	0x61, 0x72, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x68, 0x61, 0x72, 0x74, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x43, 0x43, 0x58, 0xaa, 0x02, 0x0e, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x43, 0x68, 0x61,
	0x72, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x43, 0x68,
	0x61, 0x72, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x43,
	0x68, 0x61, 0x72, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x43, 0x68, 0x61,
	0x72, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_chart_v1_densetimeserieschart_proto_rawDescOnce sync.Once
	file_chalk_chart_v1_densetimeserieschart_proto_rawDescData = file_chalk_chart_v1_densetimeserieschart_proto_rawDesc
)

func file_chalk_chart_v1_densetimeserieschart_proto_rawDescGZIP() []byte {
	file_chalk_chart_v1_densetimeserieschart_proto_rawDescOnce.Do(func() {
		file_chalk_chart_v1_densetimeserieschart_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_chart_v1_densetimeserieschart_proto_rawDescData)
	})
	return file_chalk_chart_v1_densetimeserieschart_proto_rawDescData
}

var file_chalk_chart_v1_densetimeserieschart_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chalk_chart_v1_densetimeserieschart_proto_goTypes = []any{
	(*DensePoint)(nil),            // 0: chalk.chart.v1.DensePoint
	(*GroupTag)(nil),              // 1: chalk.chart.v1.GroupTag
	(*DenseTimeSeries)(nil),       // 2: chalk.chart.v1.DenseTimeSeries
	(*DenseTimeSeriesChart)(nil),  // 3: chalk.chart.v1.DenseTimeSeriesChart
	(*v1.ScalarValue)(nil),        // 4: chalk.arrow.v1.ScalarValue
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 6: google.protobuf.Duration
}
var file_chalk_chart_v1_densetimeserieschart_proto_depIdxs = []int32{
	4, // 0: chalk.chart.v1.GroupTag.value:type_name -> chalk.arrow.v1.ScalarValue
	0, // 1: chalk.chart.v1.DenseTimeSeries.points:type_name -> chalk.chart.v1.DensePoint
	1, // 2: chalk.chart.v1.DenseTimeSeries.group_tags:type_name -> chalk.chart.v1.GroupTag
	2, // 3: chalk.chart.v1.DenseTimeSeriesChart.series:type_name -> chalk.chart.v1.DenseTimeSeries
	5, // 4: chalk.chart.v1.DenseTimeSeriesChart.x_series:type_name -> google.protobuf.Timestamp
	6, // 5: chalk.chart.v1.DenseTimeSeriesChart.window_period:type_name -> google.protobuf.Duration
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_chalk_chart_v1_densetimeserieschart_proto_init() }
func file_chalk_chart_v1_densetimeserieschart_proto_init() {
	if File_chalk_chart_v1_densetimeserieschart_proto != nil {
		return
	}
	file_chalk_chart_v1_densetimeserieschart_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_chart_v1_densetimeserieschart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chalk_chart_v1_densetimeserieschart_proto_goTypes,
		DependencyIndexes: file_chalk_chart_v1_densetimeserieschart_proto_depIdxs,
		MessageInfos:      file_chalk_chart_v1_densetimeserieschart_proto_msgTypes,
	}.Build()
	File_chalk_chart_v1_densetimeserieschart_proto = out.File
	file_chalk_chart_v1_densetimeserieschart_proto_rawDesc = nil
	file_chalk_chart_v1_densetimeserieschart_proto_goTypes = nil
	file_chalk_chart_v1_densetimeserieschart_proto_depIdxs = nil
}
