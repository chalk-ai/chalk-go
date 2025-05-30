// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: chalk/server/v1/chart.proto

package serverv1

import (
	v1 "github.com/chalk-ai/chalk-go/gen/chalk/artifacts/v1"
	_ "github.com/chalk-ai/chalk-go/gen/chalk/auth/v1"
	v11 "github.com/chalk-ai/chalk-go/gen/chalk/chart/v1"
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

type Series struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Points []float64 `protobuf:"fixed64,1,rep,packed,name=points,proto3" json:"points,omitempty"`
	Label  string    `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Units  string    `protobuf:"bytes,3,opt,name=units,proto3" json:"units,omitempty"`
}

func (x *Series) Reset() {
	*x = Series{}
	mi := &file_chalk_server_v1_chart_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Series) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Series) ProtoMessage() {}

func (x *Series) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_chart_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Series.ProtoReflect.Descriptor instead.
func (*Series) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_chart_proto_rawDescGZIP(), []int{0}
}

func (x *Series) GetPoints() []float64 {
	if x != nil {
		return x.Points
	}
	return nil
}

func (x *Series) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Series) GetUnits() string {
	if x != nil {
		return x.Units
	}
	return ""
}

type Chart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string    `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Series       []*Series `protobuf:"bytes,2,rep,name=series,proto3" json:"series,omitempty"`
	XTimestampMs []int64   `protobuf:"varint,4,rep,packed,name=x_timestamp_ms,json=xTimestampMs,proto3" json:"x_timestamp_ms,omitempty"`
}

func (x *Chart) Reset() {
	*x = Chart{}
	mi := &file_chalk_server_v1_chart_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Chart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chart) ProtoMessage() {}

func (x *Chart) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_chart_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chart.ProtoReflect.Descriptor instead.
func (*Chart) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_chart_proto_rawDescGZIP(), []int{1}
}

func (x *Chart) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Chart) GetSeries() []*Series {
	if x != nil {
		return x.Series
	}
	return nil
}

func (x *Chart) GetXTimestampMs() []int64 {
	if x != nil {
		return x.XTimestampMs
	}
	return nil
}

// Deprecated: Marked as deprecated in chalk/server/v1/chart.proto.
type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// wrapped in a Point to allow for optional (empty space in time series)
	Value *int64 `protobuf:"varint,1,opt,name=value,proto3,oneof" json:"value,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	mi := &file_chalk_server_v1_chart_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_chart_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_chart_proto_rawDescGZIP(), []int{2}
}

func (x *Point) GetValue() int64 {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return 0
}

// Deprecated: Marked as deprecated in chalk/server/v1/chart.proto.
type TimeSeries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Points []*Point `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
	Label  string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Units  string   `protobuf:"bytes,3,opt,name=units,proto3" json:"units,omitempty"`
}

func (x *TimeSeries) Reset() {
	*x = TimeSeries{}
	mi := &file_chalk_server_v1_chart_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeSeries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSeries) ProtoMessage() {}

func (x *TimeSeries) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_chart_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSeries.ProtoReflect.Descriptor instead.
func (*TimeSeries) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_chart_proto_rawDescGZIP(), []int{3}
}

func (x *TimeSeries) GetPoints() []*Point {
	if x != nil {
		return x.Points
	}
	return nil
}

func (x *TimeSeries) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *TimeSeries) GetUnits() string {
	if x != nil {
		return x.Units
	}
	return ""
}

// Deprecated: Marked as deprecated in chalk/server/v1/chart.proto.
type TimeSeriesChart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string                   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Series       []*TimeSeries            `protobuf:"bytes,2,rep,name=series,proto3" json:"series,omitempty"`
	XSeries      []*timestamppb.Timestamp `protobuf:"bytes,3,rep,name=x_series,json=xSeries,proto3" json:"x_series,omitempty"`
	WindowPeriod *durationpb.Duration     `protobuf:"bytes,4,opt,name=window_period,json=windowPeriod,proto3" json:"window_period,omitempty"`
}

func (x *TimeSeriesChart) Reset() {
	*x = TimeSeriesChart{}
	mi := &file_chalk_server_v1_chart_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeSeriesChart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSeriesChart) ProtoMessage() {}

func (x *TimeSeriesChart) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_chart_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSeriesChart.ProtoReflect.Descriptor instead.
func (*TimeSeriesChart) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_chart_proto_rawDescGZIP(), []int{4}
}

func (x *TimeSeriesChart) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *TimeSeriesChart) GetSeries() []*TimeSeries {
	if x != nil {
		return x.Series
	}
	return nil
}

func (x *TimeSeriesChart) GetXSeries() []*timestamppb.Timestamp {
	if x != nil {
		return x.XSeries
	}
	return nil
}

func (x *TimeSeriesChart) GetWindowPeriod() *durationpb.Duration {
	if x != nil {
		return x.WindowPeriod
	}
	return nil
}

type ListChartsFilters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LinkEntityKind *v1.ChartLinkKind `protobuf:"varint,1,opt,name=link_entity_kind,json=linkEntityKind,proto3,enum=chalk.artifacts.v1.ChartLinkKind,oneof" json:"link_entity_kind,omitempty"`
	LinkedEntityId *string           `protobuf:"bytes,2,opt,name=linked_entity_id,json=linkedEntityId,proto3,oneof" json:"linked_entity_id,omitempty"`
}

func (x *ListChartsFilters) Reset() {
	*x = ListChartsFilters{}
	mi := &file_chalk_server_v1_chart_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListChartsFilters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChartsFilters) ProtoMessage() {}

func (x *ListChartsFilters) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_chart_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChartsFilters.ProtoReflect.Descriptor instead.
func (*ListChartsFilters) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_chart_proto_rawDescGZIP(), []int{5}
}

func (x *ListChartsFilters) GetLinkEntityKind() v1.ChartLinkKind {
	if x != nil && x.LinkEntityKind != nil {
		return *x.LinkEntityKind
	}
	return v1.ChartLinkKind(0)
}

func (x *ListChartsFilters) GetLinkedEntityId() string {
	if x != nil && x.LinkedEntityId != nil {
		return *x.LinkedEntityId
	}
	return ""
}

type ListChartPageToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Charts are sorted first by creation time, descending
	CreatedAtHwm *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=created_at_hwm,json=createdAtHwm,proto3" json:"created_at_hwm,omitempty"`
	// Then by id, descending.
	// Note: this is the chart link id, not the metric config id hwm!
	IdHwm string `protobuf:"bytes,2,opt,name=id_hwm,json=idHwm,proto3" json:"id_hwm,omitempty"`
}

func (x *ListChartPageToken) Reset() {
	*x = ListChartPageToken{}
	mi := &file_chalk_server_v1_chart_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListChartPageToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChartPageToken) ProtoMessage() {}

func (x *ListChartPageToken) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_chart_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChartPageToken.ProtoReflect.Descriptor instead.
func (*ListChartPageToken) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_chart_proto_rawDescGZIP(), []int{6}
}

func (x *ListChartPageToken) GetCreatedAtHwm() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAtHwm
	}
	return nil
}

func (x *ListChartPageToken) GetIdHwm() string {
	if x != nil {
		return x.IdHwm
	}
	return ""
}

type ListChartsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filters *ListChartsFilters `protobuf:"bytes,1,opt,name=filters,proto3,oneof" json:"filters,omitempty"`
	Limit   *int32             `protobuf:"varint,2,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	// Must be encoded ListChartPageToken
	// see https://protobuf.dev/best-practices/api/#define-pagination-api
	PageToken *string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3,oneof" json:"page_token,omitempty"`
}

func (x *ListChartsRequest) Reset() {
	*x = ListChartsRequest{}
	mi := &file_chalk_server_v1_chart_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListChartsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChartsRequest) ProtoMessage() {}

func (x *ListChartsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_chart_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChartsRequest.ProtoReflect.Descriptor instead.
func (*ListChartsRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_chart_proto_rawDescGZIP(), []int{7}
}

func (x *ListChartsRequest) GetFilters() *ListChartsFilters {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *ListChartsRequest) GetLimit() int32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *ListChartsRequest) GetPageToken() string {
	if x != nil && x.PageToken != nil {
		return *x.PageToken
	}
	return ""
}

type ListChartsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// While still supported, heavily suggest using the second option charts_with_link
	// which includes metadata about other related entities to the chart
	//
	// Deprecated: Marked as deprecated in chalk/server/v1/chart.proto.
	Charts          []*v1.MetricConfig `protobuf:"bytes,1,rep,name=charts,proto3" json:"charts,omitempty"`
	ChartsWithLinks []*v1.Chart        `protobuf:"bytes,2,rep,name=charts_with_links,json=chartsWithLinks,proto3" json:"charts_with_links,omitempty"`
	// encoded ListChartPageToken
	// see https://protobuf.dev/best-practices/api/#define-pagination-api
	NextPageToken *string `protobuf:"bytes,3,opt,name=next_page_token,json=nextPageToken,proto3,oneof" json:"next_page_token,omitempty"`
}

func (x *ListChartsResponse) Reset() {
	*x = ListChartsResponse{}
	mi := &file_chalk_server_v1_chart_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListChartsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChartsResponse) ProtoMessage() {}

func (x *ListChartsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_chart_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChartsResponse.ProtoReflect.Descriptor instead.
func (*ListChartsResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_chart_proto_rawDescGZIP(), []int{8}
}

// Deprecated: Marked as deprecated in chalk/server/v1/chart.proto.
func (x *ListChartsResponse) GetCharts() []*v1.MetricConfig {
	if x != nil {
		return x.Charts
	}
	return nil
}

func (x *ListChartsResponse) GetChartsWithLinks() []*v1.Chart {
	if x != nil {
		return x.ChartsWithLinks
	}
	return nil
}

func (x *ListChartsResponse) GetNextPageToken() string {
	if x != nil && x.NextPageToken != nil {
		return *x.NextPageToken
	}
	return ""
}

type GetChartSnapshotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetricConfig *v1.MetricConfig       `protobuf:"bytes,1,opt,name=metric_config,json=metricConfig,proto3" json:"metric_config,omitempty"`
	StartTime    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime      *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *GetChartSnapshotRequest) Reset() {
	*x = GetChartSnapshotRequest{}
	mi := &file_chalk_server_v1_chart_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetChartSnapshotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChartSnapshotRequest) ProtoMessage() {}

func (x *GetChartSnapshotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_chart_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChartSnapshotRequest.ProtoReflect.Descriptor instead.
func (*GetChartSnapshotRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_chart_proto_rawDescGZIP(), []int{9}
}

func (x *GetChartSnapshotRequest) GetMetricConfig() *v1.MetricConfig {
	if x != nil {
		return x.MetricConfig
	}
	return nil
}

func (x *GetChartSnapshotRequest) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *GetChartSnapshotRequest) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

type GetChartSnapshotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Charts       []*v11.DenseTimeSeriesChart `protobuf:"bytes,1,rep,name=charts,proto3" json:"charts,omitempty"`
	XSeries      []*timestamppb.Timestamp    `protobuf:"bytes,2,rep,name=x_series,json=xSeries,proto3" json:"x_series,omitempty"`
	WindowPeriod *durationpb.Duration        `protobuf:"bytes,3,opt,name=window_period,json=windowPeriod,proto3" json:"window_period,omitempty"`
}

func (x *GetChartSnapshotResponse) Reset() {
	*x = GetChartSnapshotResponse{}
	mi := &file_chalk_server_v1_chart_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetChartSnapshotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChartSnapshotResponse) ProtoMessage() {}

func (x *GetChartSnapshotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_chart_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChartSnapshotResponse.ProtoReflect.Descriptor instead.
func (*GetChartSnapshotResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_chart_proto_rawDescGZIP(), []int{10}
}

func (x *GetChartSnapshotResponse) GetCharts() []*v11.DenseTimeSeriesChart {
	if x != nil {
		return x.Charts
	}
	return nil
}

func (x *GetChartSnapshotResponse) GetXSeries() []*timestamppb.Timestamp {
	if x != nil {
		return x.XSeries
	}
	return nil
}

func (x *GetChartSnapshotResponse) GetWindowPeriod() *durationpb.Duration {
	if x != nil {
		return x.WindowPeriod
	}
	return nil
}

var File_chalk_server_v1_chart_proto protoreflect.FileDescriptor

var file_chalk_server_v1_chart_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1e,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x29, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x64, 0x65, 0x6e, 0x73, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x63,
	0x68, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x06, 0x53,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x01, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x22, 0x74, 0x0a, 0x05, 0x43, 0x68, 0x61,
	0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x78, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x0c, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4d, 0x73, 0x22,
	0x30, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x88, 0x01, 0x01, 0x3a, 0x02, 0x18, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x6c, 0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12,
	0x2e, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x3a, 0x02, 0x18, 0x01, 0x22,
	0xd7, 0x01, 0x0a, 0x0f, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x43, 0x68,
	0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x35,
	0x0a, 0x08, 0x78, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x78, 0x53,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x0d, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x5f,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x77, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x3a, 0x02, 0x18, 0x01, 0x22, 0xbe, 0x01, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12,
	0x50, 0x0a, 0x10, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6b,
	0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x68, 0x61, 0x72, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x4b, 0x69, 0x6e, 0x64, 0x48, 0x00, 0x52, 0x0e,
	0x6c, 0x69, 0x6e, 0x6b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4b, 0x69, 0x6e, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x2d, 0x0a, 0x10, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0e, 0x6c,
	0x69, 0x6e, 0x6b, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x42, 0x13, 0x0a, 0x11, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64,
	0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x22, 0x6d, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x40, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x68,
	0x77, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x48,
	0x77, 0x6d, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x64, 0x5f, 0x68, 0x77, 0x6d, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x64, 0x48, 0x77, 0x6d, 0x22, 0xba, 0x01, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x41, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x48, 0x00, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01,
	0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x42, 0x08, 0x0a,
	0x06, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xda, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x68, 0x61, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a,
	0x06, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x06, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x12, 0x45, 0x0a, 0x11, 0x63,
	0x68, 0x61, 0x72, 0x74, 0x73, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x52, 0x0f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x57, 0x69, 0x74, 0x68, 0x4c, 0x69, 0x6e,
	0x6b, 0x73, 0x12, 0x2b, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42,
	0x12, 0x0a, 0x10, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0xd2, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x45, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xcf, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x61, 0x72, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x68,
	0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6e, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x06, 0x63, 0x68, 0x61,
	0x72, 0x74, 0x73, 0x12, 0x35, 0x0a, 0x08, 0x78, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x07, 0x78, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x0d, 0x77, 0x69,
	0x6e, 0x64, 0x6f, 0x77, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x77, 0x69,
	0x6e, 0x64, 0x6f, 0x77, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x32, 0xd9, 0x01, 0x0a, 0x0d, 0x43,
	0x68, 0x61, 0x72, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0a,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x12, 0x22, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x03, 0x80, 0x7d, 0x06, 0x12, 0x6c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x72, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x28, 0x2e, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x03, 0x80, 0x7d, 0x06, 0x42, 0xba, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0a,
	0x43, 0x68, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61,
	0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02, 0x0f,
	0x43, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x1b, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x11, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_server_v1_chart_proto_rawDescOnce sync.Once
	file_chalk_server_v1_chart_proto_rawDescData = file_chalk_server_v1_chart_proto_rawDesc
)

func file_chalk_server_v1_chart_proto_rawDescGZIP() []byte {
	file_chalk_server_v1_chart_proto_rawDescOnce.Do(func() {
		file_chalk_server_v1_chart_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_server_v1_chart_proto_rawDescData)
	})
	return file_chalk_server_v1_chart_proto_rawDescData
}

var file_chalk_server_v1_chart_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_chalk_server_v1_chart_proto_goTypes = []any{
	(*Series)(nil),                   // 0: chalk.server.v1.Series
	(*Chart)(nil),                    // 1: chalk.server.v1.Chart
	(*Point)(nil),                    // 2: chalk.server.v1.Point
	(*TimeSeries)(nil),               // 3: chalk.server.v1.TimeSeries
	(*TimeSeriesChart)(nil),          // 4: chalk.server.v1.TimeSeriesChart
	(*ListChartsFilters)(nil),        // 5: chalk.server.v1.ListChartsFilters
	(*ListChartPageToken)(nil),       // 6: chalk.server.v1.ListChartPageToken
	(*ListChartsRequest)(nil),        // 7: chalk.server.v1.ListChartsRequest
	(*ListChartsResponse)(nil),       // 8: chalk.server.v1.ListChartsResponse
	(*GetChartSnapshotRequest)(nil),  // 9: chalk.server.v1.GetChartSnapshotRequest
	(*GetChartSnapshotResponse)(nil), // 10: chalk.server.v1.GetChartSnapshotResponse
	(*timestamppb.Timestamp)(nil),    // 11: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),      // 12: google.protobuf.Duration
	(v1.ChartLinkKind)(0),            // 13: chalk.artifacts.v1.ChartLinkKind
	(*v1.MetricConfig)(nil),          // 14: chalk.artifacts.v1.MetricConfig
	(*v1.Chart)(nil),                 // 15: chalk.artifacts.v1.Chart
	(*v11.DenseTimeSeriesChart)(nil), // 16: chalk.chart.v1.DenseTimeSeriesChart
}
var file_chalk_server_v1_chart_proto_depIdxs = []int32{
	0,  // 0: chalk.server.v1.Chart.series:type_name -> chalk.server.v1.Series
	2,  // 1: chalk.server.v1.TimeSeries.points:type_name -> chalk.server.v1.Point
	3,  // 2: chalk.server.v1.TimeSeriesChart.series:type_name -> chalk.server.v1.TimeSeries
	11, // 3: chalk.server.v1.TimeSeriesChart.x_series:type_name -> google.protobuf.Timestamp
	12, // 4: chalk.server.v1.TimeSeriesChart.window_period:type_name -> google.protobuf.Duration
	13, // 5: chalk.server.v1.ListChartsFilters.link_entity_kind:type_name -> chalk.artifacts.v1.ChartLinkKind
	11, // 6: chalk.server.v1.ListChartPageToken.created_at_hwm:type_name -> google.protobuf.Timestamp
	5,  // 7: chalk.server.v1.ListChartsRequest.filters:type_name -> chalk.server.v1.ListChartsFilters
	14, // 8: chalk.server.v1.ListChartsResponse.charts:type_name -> chalk.artifacts.v1.MetricConfig
	15, // 9: chalk.server.v1.ListChartsResponse.charts_with_links:type_name -> chalk.artifacts.v1.Chart
	14, // 10: chalk.server.v1.GetChartSnapshotRequest.metric_config:type_name -> chalk.artifacts.v1.MetricConfig
	11, // 11: chalk.server.v1.GetChartSnapshotRequest.start_time:type_name -> google.protobuf.Timestamp
	11, // 12: chalk.server.v1.GetChartSnapshotRequest.end_time:type_name -> google.protobuf.Timestamp
	16, // 13: chalk.server.v1.GetChartSnapshotResponse.charts:type_name -> chalk.chart.v1.DenseTimeSeriesChart
	11, // 14: chalk.server.v1.GetChartSnapshotResponse.x_series:type_name -> google.protobuf.Timestamp
	12, // 15: chalk.server.v1.GetChartSnapshotResponse.window_period:type_name -> google.protobuf.Duration
	7,  // 16: chalk.server.v1.ChartsService.ListCharts:input_type -> chalk.server.v1.ListChartsRequest
	9,  // 17: chalk.server.v1.ChartsService.GetChartSnapshot:input_type -> chalk.server.v1.GetChartSnapshotRequest
	8,  // 18: chalk.server.v1.ChartsService.ListCharts:output_type -> chalk.server.v1.ListChartsResponse
	10, // 19: chalk.server.v1.ChartsService.GetChartSnapshot:output_type -> chalk.server.v1.GetChartSnapshotResponse
	18, // [18:20] is the sub-list for method output_type
	16, // [16:18] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_chalk_server_v1_chart_proto_init() }
func file_chalk_server_v1_chart_proto_init() {
	if File_chalk_server_v1_chart_proto != nil {
		return
	}
	file_chalk_server_v1_chart_proto_msgTypes[2].OneofWrappers = []any{}
	file_chalk_server_v1_chart_proto_msgTypes[5].OneofWrappers = []any{}
	file_chalk_server_v1_chart_proto_msgTypes[7].OneofWrappers = []any{}
	file_chalk_server_v1_chart_proto_msgTypes[8].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_server_v1_chart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chalk_server_v1_chart_proto_goTypes,
		DependencyIndexes: file_chalk_server_v1_chart_proto_depIdxs,
		MessageInfos:      file_chalk_server_v1_chart_proto_msgTypes,
	}.Build()
	File_chalk_server_v1_chart_proto = out.File
	file_chalk_server_v1_chart_proto_rawDesc = nil
	file_chalk_server_v1_chart_proto_goTypes = nil
	file_chalk_server_v1_chart_proto_depIdxs = nil
}
