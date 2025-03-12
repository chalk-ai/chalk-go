// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: chalk/engine/v2/feature_values_chart.proto

package enginev2

import (
	v1 "github.com/chalk-ai/chalk-go/v2/gen/chalk/chart/v1"
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

type FeatureValueBaseWindowFunction int32

const (
	FeatureValueBaseWindowFunction_FEATURE_VALUE_BASE_WINDOW_FUNCTION_UNSPECIFIED        FeatureValueBaseWindowFunction = 0
	FeatureValueBaseWindowFunction_FEATURE_VALUE_BASE_WINDOW_FUNCTION_UNIQUE_VALUES      FeatureValueBaseWindowFunction = 1
	FeatureValueBaseWindowFunction_FEATURE_VALUE_BASE_WINDOW_FUNCTION_TOTAL_OBSERVATIONS FeatureValueBaseWindowFunction = 2
	FeatureValueBaseWindowFunction_FEATURE_VALUE_BASE_WINDOW_FUNCTION_NULL_PERCENTAGE    FeatureValueBaseWindowFunction = 3
	FeatureValueBaseWindowFunction_FEATURE_VALUE_BASE_WINDOW_FUNCTION_MAX_VALUE          FeatureValueBaseWindowFunction = 4
	FeatureValueBaseWindowFunction_FEATURE_VALUE_BASE_WINDOW_FUNCTION_MIN_VALUE          FeatureValueBaseWindowFunction = 5
	FeatureValueBaseWindowFunction_FEATURE_VALUE_BASE_WINDOW_FUNCTION_AVERAGE            FeatureValueBaseWindowFunction = 6
	FeatureValueBaseWindowFunction_FEATURE_VALUE_BASE_WINDOW_FUNCTION_UNIQUE_PKEYS       FeatureValueBaseWindowFunction = 7
)

// Enum value maps for FeatureValueBaseWindowFunction.
var (
	FeatureValueBaseWindowFunction_name = map[int32]string{
		0: "FEATURE_VALUE_BASE_WINDOW_FUNCTION_UNSPECIFIED",
		1: "FEATURE_VALUE_BASE_WINDOW_FUNCTION_UNIQUE_VALUES",
		2: "FEATURE_VALUE_BASE_WINDOW_FUNCTION_TOTAL_OBSERVATIONS",
		3: "FEATURE_VALUE_BASE_WINDOW_FUNCTION_NULL_PERCENTAGE",
		4: "FEATURE_VALUE_BASE_WINDOW_FUNCTION_MAX_VALUE",
		5: "FEATURE_VALUE_BASE_WINDOW_FUNCTION_MIN_VALUE",
		6: "FEATURE_VALUE_BASE_WINDOW_FUNCTION_AVERAGE",
		7: "FEATURE_VALUE_BASE_WINDOW_FUNCTION_UNIQUE_PKEYS",
	}
	FeatureValueBaseWindowFunction_value = map[string]int32{
		"FEATURE_VALUE_BASE_WINDOW_FUNCTION_UNSPECIFIED":        0,
		"FEATURE_VALUE_BASE_WINDOW_FUNCTION_UNIQUE_VALUES":      1,
		"FEATURE_VALUE_BASE_WINDOW_FUNCTION_TOTAL_OBSERVATIONS": 2,
		"FEATURE_VALUE_BASE_WINDOW_FUNCTION_NULL_PERCENTAGE":    3,
		"FEATURE_VALUE_BASE_WINDOW_FUNCTION_MAX_VALUE":          4,
		"FEATURE_VALUE_BASE_WINDOW_FUNCTION_MIN_VALUE":          5,
		"FEATURE_VALUE_BASE_WINDOW_FUNCTION_AVERAGE":            6,
		"FEATURE_VALUE_BASE_WINDOW_FUNCTION_UNIQUE_PKEYS":       7,
	}
)

func (x FeatureValueBaseWindowFunction) Enum() *FeatureValueBaseWindowFunction {
	p := new(FeatureValueBaseWindowFunction)
	*p = x
	return p
}

func (x FeatureValueBaseWindowFunction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeatureValueBaseWindowFunction) Descriptor() protoreflect.EnumDescriptor {
	return file_chalk_engine_v2_feature_values_chart_proto_enumTypes[0].Descriptor()
}

func (FeatureValueBaseWindowFunction) Type() protoreflect.EnumType {
	return &file_chalk_engine_v2_feature_values_chart_proto_enumTypes[0]
}

func (x FeatureValueBaseWindowFunction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeatureValueBaseWindowFunction.Descriptor instead.
func (FeatureValueBaseWindowFunction) EnumDescriptor() ([]byte, []int) {
	return file_chalk_engine_v2_feature_values_chart_proto_rawDescGZIP(), []int{0}
}

type FeatureValueGroupBy int32

const (
	FeatureValueGroupBy_FEATURE_VALUE_GROUP_BY_UNSPECIFIED FeatureValueGroupBy = 0
	// Approximately ordered by cardinality / specificity
	FeatureValueGroupBy_FEATURE_VALUE_GROUP_BY_CATEGORICAL_VALUE FeatureValueGroupBy = 1
	FeatureValueGroupBy_FEATURE_VALUE_GROUP_BY_PRIMARY_KEY       FeatureValueGroupBy = 2
	FeatureValueGroupBy_FEATURE_VALUE_GROUP_BY_DEPLOYMENT_ID     FeatureValueGroupBy = 3
	FeatureValueGroupBy_FEATURE_VALUE_GROUP_BY_RESOLVER          FeatureValueGroupBy = 4
	FeatureValueGroupBy_FEATURE_VALUE_GROUP_BY_OPERATION_KIND    FeatureValueGroupBy = 5
	FeatureValueGroupBy_FEATURE_VALUE_GROUP_BY_OPERATION_ID      FeatureValueGroupBy = 6
	FeatureValueGroupBy_FEATURE_VALUE_GROUP_BY_FEATURE_VERSION   FeatureValueGroupBy = 7
)

// Enum value maps for FeatureValueGroupBy.
var (
	FeatureValueGroupBy_name = map[int32]string{
		0: "FEATURE_VALUE_GROUP_BY_UNSPECIFIED",
		1: "FEATURE_VALUE_GROUP_BY_CATEGORICAL_VALUE",
		2: "FEATURE_VALUE_GROUP_BY_PRIMARY_KEY",
		3: "FEATURE_VALUE_GROUP_BY_DEPLOYMENT_ID",
		4: "FEATURE_VALUE_GROUP_BY_RESOLVER",
		5: "FEATURE_VALUE_GROUP_BY_OPERATION_KIND",
		6: "FEATURE_VALUE_GROUP_BY_OPERATION_ID",
		7: "FEATURE_VALUE_GROUP_BY_FEATURE_VERSION",
	}
	FeatureValueGroupBy_value = map[string]int32{
		"FEATURE_VALUE_GROUP_BY_UNSPECIFIED":       0,
		"FEATURE_VALUE_GROUP_BY_CATEGORICAL_VALUE": 1,
		"FEATURE_VALUE_GROUP_BY_PRIMARY_KEY":       2,
		"FEATURE_VALUE_GROUP_BY_DEPLOYMENT_ID":     3,
		"FEATURE_VALUE_GROUP_BY_RESOLVER":          4,
		"FEATURE_VALUE_GROUP_BY_OPERATION_KIND":    5,
		"FEATURE_VALUE_GROUP_BY_OPERATION_ID":      6,
		"FEATURE_VALUE_GROUP_BY_FEATURE_VERSION":   7,
	}
)

func (x FeatureValueGroupBy) Enum() *FeatureValueGroupBy {
	p := new(FeatureValueGroupBy)
	*p = x
	return p
}

func (x FeatureValueGroupBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeatureValueGroupBy) Descriptor() protoreflect.EnumDescriptor {
	return file_chalk_engine_v2_feature_values_chart_proto_enumTypes[1].Descriptor()
}

func (FeatureValueGroupBy) Type() protoreflect.EnumType {
	return &file_chalk_engine_v2_feature_values_chart_proto_enumTypes[1]
}

func (x FeatureValueGroupBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeatureValueGroupBy.Descriptor instead.
func (FeatureValueGroupBy) EnumDescriptor() ([]byte, []int) {
	return file_chalk_engine_v2_feature_values_chart_proto_rawDescGZIP(), []int{1}
}

type FeatureValueTimeSeries int32

const (
	FeatureValueTimeSeries_FEATURE_VALUE_TIME_SERIES_UNSPECIFIED FeatureValueTimeSeries = 0
	FeatureValueTimeSeries_FEATURE_VALUE_TIME_SERIES_INSERTED_AT FeatureValueTimeSeries = 1
	FeatureValueTimeSeries_FEATURE_VALUE_TIME_SERIES_OBSERVED_AT FeatureValueTimeSeries = 2
)

// Enum value maps for FeatureValueTimeSeries.
var (
	FeatureValueTimeSeries_name = map[int32]string{
		0: "FEATURE_VALUE_TIME_SERIES_UNSPECIFIED",
		1: "FEATURE_VALUE_TIME_SERIES_INSERTED_AT",
		2: "FEATURE_VALUE_TIME_SERIES_OBSERVED_AT",
	}
	FeatureValueTimeSeries_value = map[string]int32{
		"FEATURE_VALUE_TIME_SERIES_UNSPECIFIED": 0,
		"FEATURE_VALUE_TIME_SERIES_INSERTED_AT": 1,
		"FEATURE_VALUE_TIME_SERIES_OBSERVED_AT": 2,
	}
)

func (x FeatureValueTimeSeries) Enum() *FeatureValueTimeSeries {
	p := new(FeatureValueTimeSeries)
	*p = x
	return p
}

func (x FeatureValueTimeSeries) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeatureValueTimeSeries) Descriptor() protoreflect.EnumDescriptor {
	return file_chalk_engine_v2_feature_values_chart_proto_enumTypes[2].Descriptor()
}

func (FeatureValueTimeSeries) Type() protoreflect.EnumType {
	return &file_chalk_engine_v2_feature_values_chart_proto_enumTypes[2]
}

func (x FeatureValueTimeSeries) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeatureValueTimeSeries.Descriptor instead.
func (FeatureValueTimeSeries) EnumDescriptor() ([]byte, []int) {
	return file_chalk_engine_v2_feature_values_chart_proto_rawDescGZIP(), []int{2}
}

type FeatureValuePercentileWindowFunction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Must be a value between 0 and 100.
	Percentile float64 `protobuf:"fixed64,1,opt,name=percentile,proto3" json:"percentile,omitempty"`
}

func (x *FeatureValuePercentileWindowFunction) Reset() {
	*x = FeatureValuePercentileWindowFunction{}
	mi := &file_chalk_engine_v2_feature_values_chart_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeatureValuePercentileWindowFunction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureValuePercentileWindowFunction) ProtoMessage() {}

func (x *FeatureValuePercentileWindowFunction) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_engine_v2_feature_values_chart_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureValuePercentileWindowFunction.ProtoReflect.Descriptor instead.
func (*FeatureValuePercentileWindowFunction) Descriptor() ([]byte, []int) {
	return file_chalk_engine_v2_feature_values_chart_proto_rawDescGZIP(), []int{0}
}

func (x *FeatureValuePercentileWindowFunction) GetPercentile() float64 {
	if x != nil {
		return x.Percentile
	}
	return 0
}

type FeatureValueSeries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeatureFqn string `protobuf:"bytes,1,opt,name=feature_fqn,json=featureFqn,proto3" json:"feature_fqn,omitempty"`
	// If the user passes in a series title, uses that. Otherwise, will compose a title based on the fqn and the function used
	SeriesTitle *string `protobuf:"bytes,2,opt,name=series_title,json=seriesTitle,proto3,oneof" json:"series_title,omitempty"`
	// Types that are assignable to WindowFunction:
	//
	//	*FeatureValueSeries_BaseWindowFunction
	//	*FeatureValueSeries_PercentileWindowFunction
	WindowFunction isFeatureValueSeries_WindowFunction `protobuf_oneof:"window_function"`
}

func (x *FeatureValueSeries) Reset() {
	*x = FeatureValueSeries{}
	mi := &file_chalk_engine_v2_feature_values_chart_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeatureValueSeries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureValueSeries) ProtoMessage() {}

func (x *FeatureValueSeries) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_engine_v2_feature_values_chart_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureValueSeries.ProtoReflect.Descriptor instead.
func (*FeatureValueSeries) Descriptor() ([]byte, []int) {
	return file_chalk_engine_v2_feature_values_chart_proto_rawDescGZIP(), []int{1}
}

func (x *FeatureValueSeries) GetFeatureFqn() string {
	if x != nil {
		return x.FeatureFqn
	}
	return ""
}

func (x *FeatureValueSeries) GetSeriesTitle() string {
	if x != nil && x.SeriesTitle != nil {
		return *x.SeriesTitle
	}
	return ""
}

func (m *FeatureValueSeries) GetWindowFunction() isFeatureValueSeries_WindowFunction {
	if m != nil {
		return m.WindowFunction
	}
	return nil
}

func (x *FeatureValueSeries) GetBaseWindowFunction() FeatureValueBaseWindowFunction {
	if x, ok := x.GetWindowFunction().(*FeatureValueSeries_BaseWindowFunction); ok {
		return x.BaseWindowFunction
	}
	return FeatureValueBaseWindowFunction_FEATURE_VALUE_BASE_WINDOW_FUNCTION_UNSPECIFIED
}

func (x *FeatureValueSeries) GetPercentileWindowFunction() *FeatureValuePercentileWindowFunction {
	if x, ok := x.GetWindowFunction().(*FeatureValueSeries_PercentileWindowFunction); ok {
		return x.PercentileWindowFunction
	}
	return nil
}

type isFeatureValueSeries_WindowFunction interface {
	isFeatureValueSeries_WindowFunction()
}

type FeatureValueSeries_BaseWindowFunction struct {
	BaseWindowFunction FeatureValueBaseWindowFunction `protobuf:"varint,3,opt,name=base_window_function,json=baseWindowFunction,proto3,enum=chalk.engine.v2.FeatureValueBaseWindowFunction,oneof"`
}

type FeatureValueSeries_PercentileWindowFunction struct {
	PercentileWindowFunction *FeatureValuePercentileWindowFunction `protobuf:"bytes,4,opt,name=percentile_window_function,json=percentileWindowFunction,proto3,oneof"`
}

func (*FeatureValueSeries_BaseWindowFunction) isFeatureValueSeries_WindowFunction() {}

func (*FeatureValueSeries_PercentileWindowFunction) isFeatureValueSeries_WindowFunction() {}

type GetFeatureValuesTimeSeriesChartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  *string               `protobuf:"bytes,1,opt,name=title,proto3,oneof" json:"title,omitempty"`
	Series []*FeatureValueSeries `protobuf:"bytes,2,rep,name=series,proto3" json:"series,omitempty"`
	// The size of each time bucket + how far apart time points are in the chart
	// Ideally a divisor of the total window between end and start ms
	// If not, the bucket aligned with end_timestamp_exclusive will be of size total_window % window_period
	// For JSON format, this is a string `${duration_in_seconds}s`
	WindowPeriod            *durationpb.Duration   `protobuf:"bytes,3,opt,name=window_period,json=windowPeriod,proto3" json:"window_period,omitempty"`
	StartTimestampInclusive *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=start_timestamp_inclusive,json=startTimestampInclusive,proto3" json:"start_timestamp_inclusive,omitempty"`
	// If not provided, will assume current time is the ending time
	EndTimestampExclusive *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=end_timestamp_exclusive,json=endTimestampExclusive,proto3,oneof" json:"end_timestamp_exclusive,omitempty"`
	// The dimension to use as the time series axis. Defaults to inserted_at if not specified
	TimeSeries *FeatureValueTimeSeries `protobuf:"varint,6,opt,name=time_series,json=timeSeries,proto3,enum=chalk.engine.v2.FeatureValueTimeSeries,oneof" json:"time_series,omitempty"`
	// The group bys to apply to this chart. Multiple group bys may produce high #'s of axes.
	GroupBy []FeatureValueGroupBy `protobuf:"varint,7,rep,packed,name=group_by,json=groupBy,proto3,enum=chalk.engine.v2.FeatureValueGroupBy" json:"group_by,omitempty"`
	// The maximum number of series produced by the group by.
	// Important because some group bys are high cardinality (ex. value)
	GroupBySeriesLimit *int32 `protobuf:"varint,8,opt,name=group_by_series_limit,json=groupBySeriesLimit,proto3,oneof" json:"group_by_series_limit,omitempty"`
}

func (x *GetFeatureValuesTimeSeriesChartRequest) Reset() {
	*x = GetFeatureValuesTimeSeriesChartRequest{}
	mi := &file_chalk_engine_v2_feature_values_chart_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFeatureValuesTimeSeriesChartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeatureValuesTimeSeriesChartRequest) ProtoMessage() {}

func (x *GetFeatureValuesTimeSeriesChartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_engine_v2_feature_values_chart_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeatureValuesTimeSeriesChartRequest.ProtoReflect.Descriptor instead.
func (*GetFeatureValuesTimeSeriesChartRequest) Descriptor() ([]byte, []int) {
	return file_chalk_engine_v2_feature_values_chart_proto_rawDescGZIP(), []int{2}
}

func (x *GetFeatureValuesTimeSeriesChartRequest) GetTitle() string {
	if x != nil && x.Title != nil {
		return *x.Title
	}
	return ""
}

func (x *GetFeatureValuesTimeSeriesChartRequest) GetSeries() []*FeatureValueSeries {
	if x != nil {
		return x.Series
	}
	return nil
}

func (x *GetFeatureValuesTimeSeriesChartRequest) GetWindowPeriod() *durationpb.Duration {
	if x != nil {
		return x.WindowPeriod
	}
	return nil
}

func (x *GetFeatureValuesTimeSeriesChartRequest) GetStartTimestampInclusive() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTimestampInclusive
	}
	return nil
}

func (x *GetFeatureValuesTimeSeriesChartRequest) GetEndTimestampExclusive() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTimestampExclusive
	}
	return nil
}

func (x *GetFeatureValuesTimeSeriesChartRequest) GetTimeSeries() FeatureValueTimeSeries {
	if x != nil && x.TimeSeries != nil {
		return *x.TimeSeries
	}
	return FeatureValueTimeSeries_FEATURE_VALUE_TIME_SERIES_UNSPECIFIED
}

func (x *GetFeatureValuesTimeSeriesChartRequest) GetGroupBy() []FeatureValueGroupBy {
	if x != nil {
		return x.GroupBy
	}
	return nil
}

func (x *GetFeatureValuesTimeSeriesChartRequest) GetGroupBySeriesLimit() int32 {
	if x != nil && x.GroupBySeriesLimit != nil {
		return *x.GroupBySeriesLimit
	}
	return 0
}

type GetFeatureValuesTimeSeriesChartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chart *v1.DenseTimeSeriesChart `protobuf:"bytes,1,opt,name=chart,proto3" json:"chart,omitempty"`
}

func (x *GetFeatureValuesTimeSeriesChartResponse) Reset() {
	*x = GetFeatureValuesTimeSeriesChartResponse{}
	mi := &file_chalk_engine_v2_feature_values_chart_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFeatureValuesTimeSeriesChartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeatureValuesTimeSeriesChartResponse) ProtoMessage() {}

func (x *GetFeatureValuesTimeSeriesChartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_engine_v2_feature_values_chart_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeatureValuesTimeSeriesChartResponse.ProtoReflect.Descriptor instead.
func (*GetFeatureValuesTimeSeriesChartResponse) Descriptor() ([]byte, []int) {
	return file_chalk_engine_v2_feature_values_chart_proto_rawDescGZIP(), []int{3}
}

func (x *GetFeatureValuesTimeSeriesChartResponse) GetChart() *v1.DenseTimeSeriesChart {
	if x != nil {
		return x.Chart
	}
	return nil
}

var File_chalk_engine_v2_feature_values_chart_proto protoreflect.FileDescriptor

var file_chalk_engine_v2_feature_values_chart_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x32, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x32, 0x1a, 0x29, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65,
	0x6e, 0x73, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x63, 0x68, 0x61,
	0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x24, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x69, 0x6c, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c,
	0x65, 0x22, 0xdd, 0x02, 0x0a, 0x12, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x5f, 0x66, 0x71, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x71, 0x6e, 0x12, 0x26, 0x0a, 0x0c, 0x73, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x63, 0x0a, 0x14, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77,
	0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2f, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x61,
	0x73, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x12, 0x62, 0x61, 0x73, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x75, 0x0a, 0x1a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x69, 0x6c, 0x65, 0x5f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x69, 0x6c, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x18, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x69, 0x6c, 0x65, 0x57,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x11, 0x0a,
	0x0f, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x22, 0x89, 0x05, 0x0a, 0x26, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x06, 0x73, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x0d, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x12, 0x56, 0x0a, 0x19, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x17, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x12, 0x57, 0x0a, 0x17,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x15, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69,
	0x76, 0x65, 0x88, 0x01, 0x01, 0x12, 0x4d, 0x0a, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x48, 0x02, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x3f, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62, 0x79,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x42, 0x79, 0x12, 0x36, 0x0a, 0x15, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62,
	0x79, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x12, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x79, 0x53,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x42, 0x1a, 0x0a, 0x18, 0x5f, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73,
	0x69, 0x76, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x62, 0x79,
	0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x65, 0x0a,
	0x27, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x43, 0x68, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x72,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6e, 0x73, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x05, 0x63,
	0x68, 0x61, 0x72, 0x74, 0x2a, 0xc6, 0x03, 0x0a, 0x1e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x61, 0x73, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x2e, 0x46, 0x45, 0x41, 0x54, 0x55,
	0x52, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x57, 0x49,
	0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x34, 0x0a, 0x30, 0x46,
	0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x42, 0x41, 0x53,
	0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x55, 0x4e, 0x49, 0x51, 0x55, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x53, 0x10,
	0x01, 0x12, 0x39, 0x0a, 0x35, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x56, 0x41, 0x4c,
	0x55, 0x45, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x46,
	0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x54, 0x41, 0x4c, 0x5f, 0x4f, 0x42,
	0x53, 0x45, 0x52, 0x56, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x10, 0x02, 0x12, 0x36, 0x0a, 0x32,
	0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x42, 0x41,
	0x53, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x4e, 0x55, 0x4c, 0x4c, 0x5f, 0x50, 0x45, 0x52, 0x43, 0x45, 0x4e, 0x54, 0x41,
	0x47, 0x45, 0x10, 0x03, 0x12, 0x30, 0x0a, 0x2c, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f,
	0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x44, 0x4f,
	0x57, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x41, 0x58, 0x5f, 0x56,
	0x41, 0x4c, 0x55, 0x45, 0x10, 0x04, 0x12, 0x30, 0x0a, 0x2c, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52,
	0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x57, 0x49, 0x4e,
	0x44, 0x4f, 0x57, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x49, 0x4e,
	0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x05, 0x12, 0x2e, 0x0a, 0x2a, 0x46, 0x45, 0x41, 0x54,
	0x55, 0x52, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x57,
	0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41,
	0x56, 0x45, 0x52, 0x41, 0x47, 0x45, 0x10, 0x06, 0x12, 0x33, 0x0a, 0x2f, 0x46, 0x45, 0x41, 0x54,
	0x55, 0x52, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x5f, 0x57,
	0x49, 0x4e, 0x44, 0x4f, 0x57, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x4e, 0x49, 0x51, 0x55, 0x45, 0x5f, 0x50, 0x4b, 0x45, 0x59, 0x53, 0x10, 0x07, 0x2a, 0xe2, 0x02,
	0x0a, 0x13, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x42, 0x79, 0x12, 0x26, 0x0a, 0x22, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45,
	0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x42, 0x59, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x2c, 0x0a,
	0x28, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x47,
	0x52, 0x4f, 0x55, 0x50, 0x5f, 0x42, 0x59, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x49,
	0x43, 0x41, 0x4c, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x01, 0x12, 0x26, 0x0a, 0x22, 0x46,
	0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x47, 0x52, 0x4f,
	0x55, 0x50, 0x5f, 0x42, 0x59, 0x5f, 0x50, 0x52, 0x49, 0x4d, 0x41, 0x52, 0x59, 0x5f, 0x4b, 0x45,
	0x59, 0x10, 0x02, 0x12, 0x28, 0x0a, 0x24, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x56,
	0x41, 0x4c, 0x55, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x42, 0x59, 0x5f, 0x44, 0x45,
	0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x49, 0x44, 0x10, 0x03, 0x12, 0x23, 0x0a,
	0x1f, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x47,
	0x52, 0x4f, 0x55, 0x50, 0x5f, 0x42, 0x59, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x4c, 0x56, 0x45, 0x52,
	0x10, 0x04, 0x12, 0x29, 0x0a, 0x25, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x56, 0x41,
	0x4c, 0x55, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x42, 0x59, 0x5f, 0x4f, 0x50, 0x45,
	0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x10, 0x05, 0x12, 0x27, 0x0a,
	0x23, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x47,
	0x52, 0x4f, 0x55, 0x50, 0x5f, 0x42, 0x59, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f,
	0x4e, 0x5f, 0x49, 0x44, 0x10, 0x06, 0x12, 0x2a, 0x0a, 0x26, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52,
	0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x42, 0x59,
	0x5f, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e,
	0x10, 0x07, 0x2a, 0x99, 0x01, 0x0a, 0x16, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x29, 0x0a,
	0x25, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54,
	0x49, 0x4d, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x49, 0x45, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x29, 0x0a, 0x25, 0x46, 0x45, 0x41, 0x54,
	0x55, 0x52, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x53,
	0x45, 0x52, 0x49, 0x45, 0x53, 0x5f, 0x49, 0x4e, 0x53, 0x45, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x41,
	0x54, 0x10, 0x01, 0x12, 0x29, 0x0a, 0x25, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x56,
	0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x49, 0x45, 0x53,
	0x5f, 0x4f, 0x42, 0x53, 0x45, 0x52, 0x56, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x10, 0x02, 0x42, 0xc7,
	0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x32, 0x42, 0x17, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x43, 0x68, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2f, 0x76, 0x32, 0x3b, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x43,
	0x45, 0x58, 0xaa, 0x02, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x5c, 0x56, 0x32, 0xe2, 0x02, 0x1b, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_engine_v2_feature_values_chart_proto_rawDescOnce sync.Once
	file_chalk_engine_v2_feature_values_chart_proto_rawDescData = file_chalk_engine_v2_feature_values_chart_proto_rawDesc
)

func file_chalk_engine_v2_feature_values_chart_proto_rawDescGZIP() []byte {
	file_chalk_engine_v2_feature_values_chart_proto_rawDescOnce.Do(func() {
		file_chalk_engine_v2_feature_values_chart_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_engine_v2_feature_values_chart_proto_rawDescData)
	})
	return file_chalk_engine_v2_feature_values_chart_proto_rawDescData
}

var file_chalk_engine_v2_feature_values_chart_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_chalk_engine_v2_feature_values_chart_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chalk_engine_v2_feature_values_chart_proto_goTypes = []any{
	(FeatureValueBaseWindowFunction)(0),             // 0: chalk.engine.v2.FeatureValueBaseWindowFunction
	(FeatureValueGroupBy)(0),                        // 1: chalk.engine.v2.FeatureValueGroupBy
	(FeatureValueTimeSeries)(0),                     // 2: chalk.engine.v2.FeatureValueTimeSeries
	(*FeatureValuePercentileWindowFunction)(nil),    // 3: chalk.engine.v2.FeatureValuePercentileWindowFunction
	(*FeatureValueSeries)(nil),                      // 4: chalk.engine.v2.FeatureValueSeries
	(*GetFeatureValuesTimeSeriesChartRequest)(nil),  // 5: chalk.engine.v2.GetFeatureValuesTimeSeriesChartRequest
	(*GetFeatureValuesTimeSeriesChartResponse)(nil), // 6: chalk.engine.v2.GetFeatureValuesTimeSeriesChartResponse
	(*durationpb.Duration)(nil),                     // 7: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil),                   // 8: google.protobuf.Timestamp
	(*v1.DenseTimeSeriesChart)(nil),                 // 9: chalk.chart.v1.DenseTimeSeriesChart
}
var file_chalk_engine_v2_feature_values_chart_proto_depIdxs = []int32{
	0, // 0: chalk.engine.v2.FeatureValueSeries.base_window_function:type_name -> chalk.engine.v2.FeatureValueBaseWindowFunction
	3, // 1: chalk.engine.v2.FeatureValueSeries.percentile_window_function:type_name -> chalk.engine.v2.FeatureValuePercentileWindowFunction
	4, // 2: chalk.engine.v2.GetFeatureValuesTimeSeriesChartRequest.series:type_name -> chalk.engine.v2.FeatureValueSeries
	7, // 3: chalk.engine.v2.GetFeatureValuesTimeSeriesChartRequest.window_period:type_name -> google.protobuf.Duration
	8, // 4: chalk.engine.v2.GetFeatureValuesTimeSeriesChartRequest.start_timestamp_inclusive:type_name -> google.protobuf.Timestamp
	8, // 5: chalk.engine.v2.GetFeatureValuesTimeSeriesChartRequest.end_timestamp_exclusive:type_name -> google.protobuf.Timestamp
	2, // 6: chalk.engine.v2.GetFeatureValuesTimeSeriesChartRequest.time_series:type_name -> chalk.engine.v2.FeatureValueTimeSeries
	1, // 7: chalk.engine.v2.GetFeatureValuesTimeSeriesChartRequest.group_by:type_name -> chalk.engine.v2.FeatureValueGroupBy
	9, // 8: chalk.engine.v2.GetFeatureValuesTimeSeriesChartResponse.chart:type_name -> chalk.chart.v1.DenseTimeSeriesChart
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_chalk_engine_v2_feature_values_chart_proto_init() }
func file_chalk_engine_v2_feature_values_chart_proto_init() {
	if File_chalk_engine_v2_feature_values_chart_proto != nil {
		return
	}
	file_chalk_engine_v2_feature_values_chart_proto_msgTypes[1].OneofWrappers = []any{
		(*FeatureValueSeries_BaseWindowFunction)(nil),
		(*FeatureValueSeries_PercentileWindowFunction)(nil),
	}
	file_chalk_engine_v2_feature_values_chart_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_engine_v2_feature_values_chart_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chalk_engine_v2_feature_values_chart_proto_goTypes,
		DependencyIndexes: file_chalk_engine_v2_feature_values_chart_proto_depIdxs,
		EnumInfos:         file_chalk_engine_v2_feature_values_chart_proto_enumTypes,
		MessageInfos:      file_chalk_engine_v2_feature_values_chart_proto_msgTypes,
	}.Build()
	File_chalk_engine_v2_feature_values_chart_proto = out.File
	file_chalk_engine_v2_feature_values_chart_proto_rawDesc = nil
	file_chalk_engine_v2_feature_values_chart_proto_goTypes = nil
	file_chalk_engine_v2_feature_values_chart_proto_depIdxs = nil
}
