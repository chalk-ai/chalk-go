// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc        (unknown)
// source: chalk/server/v1/billing.proto

package serverv1

import (
	_ "github.com/chalk-ai/chalk-go/gen/chalk/auth/v1"
	v1 "github.com/chalk-ai/chalk-go/gen/chalk/pubsub/v1"
	v11 "github.com/chalk-ai/chalk-go/gen/chalk/usage/v1"
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

type UsageChartPeriod int32

const (
	UsageChartPeriod_USAGE_CHART_PERIOD_UNSPECIFIED UsageChartPeriod = 0
	UsageChartPeriod_USAGE_CHART_PERIOD_DAILY       UsageChartPeriod = 1
	UsageChartPeriod_USAGE_CHART_PERIOD_MONTHLY     UsageChartPeriod = 2
)

// Enum value maps for UsageChartPeriod.
var (
	UsageChartPeriod_name = map[int32]string{
		0: "USAGE_CHART_PERIOD_UNSPECIFIED",
		1: "USAGE_CHART_PERIOD_DAILY",
		2: "USAGE_CHART_PERIOD_MONTHLY",
	}
	UsageChartPeriod_value = map[string]int32{
		"USAGE_CHART_PERIOD_UNSPECIFIED": 0,
		"USAGE_CHART_PERIOD_DAILY":       1,
		"USAGE_CHART_PERIOD_MONTHLY":     2,
	}
)

func (x UsageChartPeriod) Enum() *UsageChartPeriod {
	p := new(UsageChartPeriod)
	*p = x
	return p
}

func (x UsageChartPeriod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UsageChartPeriod) Descriptor() protoreflect.EnumDescriptor {
	return file_chalk_server_v1_billing_proto_enumTypes[0].Descriptor()
}

func (UsageChartPeriod) Type() protoreflect.EnumType {
	return &file_chalk_server_v1_billing_proto_enumTypes[0]
}

func (x UsageChartPeriod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UsageChartPeriod.Descriptor instead.
func (UsageChartPeriod) EnumDescriptor() ([]byte, []int) {
	return file_chalk_server_v1_billing_proto_rawDescGZIP(), []int{0}
}

type GetNodesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetNodesRequest) Reset() {
	*x = GetNodesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_billing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodesRequest) ProtoMessage() {}

func (x *GetNodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_billing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodesRequest.ProtoReflect.Descriptor instead.
func (*GetNodesRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_billing_proto_rawDescGZIP(), []int{0}
}

type GetNodesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeStatuses []*v1.NodeStatusPubSub `protobuf:"bytes,1,rep,name=node_statuses,json=nodeStatuses,proto3" json:"node_statuses,omitempty"`
	PodStatuses  []*v1.PodStatusPubSub  `protobuf:"bytes,2,rep,name=pod_statuses,json=podStatuses,proto3" json:"pod_statuses,omitempty"`
}

func (x *GetNodesResponse) Reset() {
	*x = GetNodesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_billing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodesResponse) ProtoMessage() {}

func (x *GetNodesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_billing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodesResponse.ProtoReflect.Descriptor instead.
func (*GetNodesResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_billing_proto_rawDescGZIP(), []int{1}
}

func (x *GetNodesResponse) GetNodeStatuses() []*v1.NodeStatusPubSub {
	if x != nil {
		return x.NodeStatuses
	}
	return nil
}

func (x *GetNodesResponse) GetPodStatuses() []*v1.PodStatusPubSub {
	if x != nil {
		return x.PodStatuses
	}
	return nil
}

type GetUsageChartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartMs *int64            `protobuf:"varint,1,opt,name=start_ms,json=startMs,proto3,oneof" json:"start_ms,omitempty"`
	EndMs   *int64            `protobuf:"varint,2,opt,name=end_ms,json=endMs,proto3,oneof" json:"end_ms,omitempty"`
	Period  *UsageChartPeriod `protobuf:"varint,3,opt,name=period,proto3,enum=chalk.server.v1.UsageChartPeriod,oneof" json:"period,omitempty"`
}

func (x *GetUsageChartRequest) Reset() {
	*x = GetUsageChartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_billing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsageChartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsageChartRequest) ProtoMessage() {}

func (x *GetUsageChartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_billing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsageChartRequest.ProtoReflect.Descriptor instead.
func (*GetUsageChartRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_billing_proto_rawDescGZIP(), []int{2}
}

func (x *GetUsageChartRequest) GetStartMs() int64 {
	if x != nil && x.StartMs != nil {
		return *x.StartMs
	}
	return 0
}

func (x *GetUsageChartRequest) GetEndMs() int64 {
	if x != nil && x.EndMs != nil {
		return *x.EndMs
	}
	return 0
}

func (x *GetUsageChartRequest) GetPeriod() UsageChartPeriod {
	if x != nil && x.Period != nil {
		return *x.Period
	}
	return UsageChartPeriod_USAGE_CHART_PERIOD_UNSPECIFIED
}

type GetUsageChartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chart *Chart `protobuf:"bytes,1,opt,name=chart,proto3" json:"chart,omitempty"`
}

func (x *GetUsageChartResponse) Reset() {
	*x = GetUsageChartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_billing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsageChartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsageChartResponse) ProtoMessage() {}

func (x *GetUsageChartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_billing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsageChartResponse.ProtoReflect.Descriptor instead.
func (*GetUsageChartResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_billing_proto_rawDescGZIP(), []int{3}
}

func (x *GetUsageChartResponse) GetChart() *Chart {
	if x != nil {
		return x.Chart
	}
	return nil
}

type GetUtilizationRatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUtilizationRatesRequest) Reset() {
	*x = GetUtilizationRatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_billing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUtilizationRatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUtilizationRatesRequest) ProtoMessage() {}

func (x *GetUtilizationRatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_billing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUtilizationRatesRequest.ProtoReflect.Descriptor instead.
func (*GetUtilizationRatesRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_billing_proto_rawDescGZIP(), []int{4}
}

type GetUtilizationRatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rates []*v11.MachineRate `protobuf:"bytes,1,rep,name=rates,proto3" json:"rates,omitempty"`
}

func (x *GetUtilizationRatesResponse) Reset() {
	*x = GetUtilizationRatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_billing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUtilizationRatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUtilizationRatesResponse) ProtoMessage() {}

func (x *GetUtilizationRatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_billing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUtilizationRatesResponse.ProtoReflect.Descriptor instead.
func (*GetUtilizationRatesResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_billing_proto_rawDescGZIP(), []int{5}
}

func (x *GetUtilizationRatesResponse) GetRates() []*v11.MachineRate {
	if x != nil {
		return x.Rates
	}
	return nil
}

var File_chalk_server_v1_billing_proto protoreflect.FileDescriptor

var file_chalk_server_v1_billing_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x21, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2f,
	0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x70, 0x75, 0x62, 0x73,
	0x75, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x75, 0x73, 0x61, 0x67, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x9f, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x75, 0x62, 0x53, 0x75, 0x62,
	0x52, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x12, 0x43,
	0x0a, 0x0c, 0x70, 0x6f, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x70, 0x75, 0x62,
	0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x52, 0x0b, 0x70, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x65, 0x73, 0x22, 0xb5, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x08,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00,
	0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x4d, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x06,
	0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x05,
	0x65, 0x6e, 0x64, 0x4d, 0x73, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x68, 0x61, 0x72, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x48, 0x02, 0x52, 0x06, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x6d, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x73,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22, 0x45, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x05, 0x63, 0x68, 0x61,
	0x72, 0x74, 0x22, 0x1c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x50, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x05, 0x72, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x75, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x05, 0x72, 0x61, 0x74,
	0x65, 0x73, 0x2a, 0x74, 0x0a, 0x10, 0x55, 0x73, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x22, 0x0a, 0x1e, 0x55, 0x53, 0x41, 0x47, 0x45, 0x5f,
	0x43, 0x48, 0x41, 0x52, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x49, 0x4f, 0x44, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x55, 0x53,
	0x41, 0x47, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x49, 0x4f, 0x44,
	0x5f, 0x44, 0x41, 0x49, 0x4c, 0x59, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x55, 0x53, 0x41, 0x47,
	0x45, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x54, 0x5f, 0x50, 0x45, 0x52, 0x49, 0x4f, 0x44, 0x5f, 0x4d,
	0x4f, 0x4e, 0x54, 0x48, 0x4c, 0x59, 0x10, 0x02, 0x32, 0xcb, 0x02, 0x0a, 0x0e, 0x42, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x80, 0x7d,
	0x0b, 0x90, 0x02, 0x01, 0x12, 0x66, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x25, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x80, 0x7d, 0x06, 0x90, 0x02, 0x01, 0x12, 0x78, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x2b, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2c, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x74, 0x69, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06,
	0x80, 0x7d, 0x01, 0x90, 0x02, 0x01, 0x42, 0xbc, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c,
	0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa,
	0x02, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x11, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_server_v1_billing_proto_rawDescOnce sync.Once
	file_chalk_server_v1_billing_proto_rawDescData = file_chalk_server_v1_billing_proto_rawDesc
)

func file_chalk_server_v1_billing_proto_rawDescGZIP() []byte {
	file_chalk_server_v1_billing_proto_rawDescOnce.Do(func() {
		file_chalk_server_v1_billing_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_server_v1_billing_proto_rawDescData)
	})
	return file_chalk_server_v1_billing_proto_rawDescData
}

var file_chalk_server_v1_billing_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_chalk_server_v1_billing_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_chalk_server_v1_billing_proto_goTypes = []interface{}{
	(UsageChartPeriod)(0),               // 0: chalk.server.v1.UsageChartPeriod
	(*GetNodesRequest)(nil),             // 1: chalk.server.v1.GetNodesRequest
	(*GetNodesResponse)(nil),            // 2: chalk.server.v1.GetNodesResponse
	(*GetUsageChartRequest)(nil),        // 3: chalk.server.v1.GetUsageChartRequest
	(*GetUsageChartResponse)(nil),       // 4: chalk.server.v1.GetUsageChartResponse
	(*GetUtilizationRatesRequest)(nil),  // 5: chalk.server.v1.GetUtilizationRatesRequest
	(*GetUtilizationRatesResponse)(nil), // 6: chalk.server.v1.GetUtilizationRatesResponse
	(*v1.NodeStatusPubSub)(nil),         // 7: chalk.pubsub.v1.NodeStatusPubSub
	(*v1.PodStatusPubSub)(nil),          // 8: chalk.pubsub.v1.PodStatusPubSub
	(*Chart)(nil),                       // 9: chalk.server.v1.Chart
	(*v11.MachineRate)(nil),             // 10: chalk.usage.v1.MachineRate
}
var file_chalk_server_v1_billing_proto_depIdxs = []int32{
	7,  // 0: chalk.server.v1.GetNodesResponse.node_statuses:type_name -> chalk.pubsub.v1.NodeStatusPubSub
	8,  // 1: chalk.server.v1.GetNodesResponse.pod_statuses:type_name -> chalk.pubsub.v1.PodStatusPubSub
	0,  // 2: chalk.server.v1.GetUsageChartRequest.period:type_name -> chalk.server.v1.UsageChartPeriod
	9,  // 3: chalk.server.v1.GetUsageChartResponse.chart:type_name -> chalk.server.v1.Chart
	10, // 4: chalk.server.v1.GetUtilizationRatesResponse.rates:type_name -> chalk.usage.v1.MachineRate
	1,  // 5: chalk.server.v1.BillingService.GetNodes:input_type -> chalk.server.v1.GetNodesRequest
	3,  // 6: chalk.server.v1.BillingService.GetUsageChart:input_type -> chalk.server.v1.GetUsageChartRequest
	5,  // 7: chalk.server.v1.BillingService.GetUtilizationRates:input_type -> chalk.server.v1.GetUtilizationRatesRequest
	2,  // 8: chalk.server.v1.BillingService.GetNodes:output_type -> chalk.server.v1.GetNodesResponse
	4,  // 9: chalk.server.v1.BillingService.GetUsageChart:output_type -> chalk.server.v1.GetUsageChartResponse
	6,  // 10: chalk.server.v1.BillingService.GetUtilizationRates:output_type -> chalk.server.v1.GetUtilizationRatesResponse
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_chalk_server_v1_billing_proto_init() }
func file_chalk_server_v1_billing_proto_init() {
	if File_chalk_server_v1_billing_proto != nil {
		return
	}
	file_chalk_server_v1_chart_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chalk_server_v1_billing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodesRequest); i {
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
		file_chalk_server_v1_billing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodesResponse); i {
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
		file_chalk_server_v1_billing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsageChartRequest); i {
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
		file_chalk_server_v1_billing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsageChartResponse); i {
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
		file_chalk_server_v1_billing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUtilizationRatesRequest); i {
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
		file_chalk_server_v1_billing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUtilizationRatesResponse); i {
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
	file_chalk_server_v1_billing_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_server_v1_billing_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chalk_server_v1_billing_proto_goTypes,
		DependencyIndexes: file_chalk_server_v1_billing_proto_depIdxs,
		EnumInfos:         file_chalk_server_v1_billing_proto_enumTypes,
		MessageInfos:      file_chalk_server_v1_billing_proto_msgTypes,
	}.Build()
	File_chalk_server_v1_billing_proto = out.File
	file_chalk_server_v1_billing_proto_rawDesc = nil
	file_chalk_server_v1_billing_proto_goTypes = nil
	file_chalk_server_v1_billing_proto_depIdxs = nil
}
