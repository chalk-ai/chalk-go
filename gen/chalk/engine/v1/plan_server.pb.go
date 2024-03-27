// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc        (unknown)
// source: chalk/engine/v1/plan_server.proto

package enginev1

import (
	v1 "github.com/chalk-ai/chalk-private/go-api-server/gen/chalk/common/v1"
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

type GetPlanRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OnlineQueryRequest *v1.OnlineQueryRequest `protobuf:"bytes,1,opt,name=online_query_request,json=onlineQueryRequest,proto3" json:"online_query_request,omitempty"`
}

func (x *GetPlanRequest) Reset() {
	*x = GetPlanRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_engine_v1_plan_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlanRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlanRequest) ProtoMessage() {}

func (x *GetPlanRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_engine_v1_plan_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlanRequest.ProtoReflect.Descriptor instead.
func (*GetPlanRequest) Descriptor() ([]byte, []int) {
	return file_chalk_engine_v1_plan_server_proto_rawDescGZIP(), []int{0}
}

func (x *GetPlanRequest) GetOnlineQueryRequest() *v1.OnlineQueryRequest {
	if x != nil {
		return x.OnlineQueryRequest
	}
	return nil
}

type GetPlanResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plan *Plan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan,omitempty"`
}

func (x *GetPlanResponse) Reset() {
	*x = GetPlanResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_engine_v1_plan_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlanResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlanResponse) ProtoMessage() {}

func (x *GetPlanResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_engine_v1_plan_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlanResponse.ProtoReflect.Descriptor instead.
func (*GetPlanResponse) Descriptor() ([]byte, []int) {
	return file_chalk_engine_v1_plan_server_proto_rawDescGZIP(), []int{1}
}

func (x *GetPlanResponse) GetPlan() *Plan {
	if x != nil {
		return x.Plan
	}
	return nil
}

type ExecuteQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OnlineQueryRequest *v1.OnlineQueryRequest `protobuf:"bytes,1,opt,name=online_query_request,json=onlineQueryRequest,proto3" json:"online_query_request,omitempty"`
}

func (x *ExecuteQueryRequest) Reset() {
	*x = ExecuteQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_engine_v1_plan_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteQueryRequest) ProtoMessage() {}

func (x *ExecuteQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_engine_v1_plan_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteQueryRequest.ProtoReflect.Descriptor instead.
func (*ExecuteQueryRequest) Descriptor() ([]byte, []int) {
	return file_chalk_engine_v1_plan_server_proto_rawDescGZIP(), []int{2}
}

func (x *ExecuteQueryRequest) GetOnlineQueryRequest() *v1.OnlineQueryRequest {
	if x != nil {
		return x.OnlineQueryRequest
	}
	return nil
}

type ExecuteQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OnlineQueryResponse *v1.OnlineQueryResponse `protobuf:"bytes,1,opt,name=online_query_response,json=onlineQueryResponse,proto3" json:"online_query_response,omitempty"`
}

func (x *ExecuteQueryResponse) Reset() {
	*x = ExecuteQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_engine_v1_plan_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteQueryResponse) ProtoMessage() {}

func (x *ExecuteQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_engine_v1_plan_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteQueryResponse.ProtoReflect.Descriptor instead.
func (*ExecuteQueryResponse) Descriptor() ([]byte, []int) {
	return file_chalk_engine_v1_plan_server_proto_rawDescGZIP(), []int{3}
}

func (x *ExecuteQueryResponse) GetOnlineQueryResponse() *v1.OnlineQueryResponse {
	if x != nil {
		return x.OnlineQueryResponse
	}
	return nil
}

var File_chalk_engine_v1_plan_server_proto protoreflect.FileDescriptor

var file_chalk_engine_v1_plan_server_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x22, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x55, 0x0a, 0x14, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x12, 0x6f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x22, 0x6c, 0x0a, 0x13, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x55, 0x0a, 0x14, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x12, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x70, 0x0a, 0x14, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x58, 0x0a, 0x15, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x13, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xbc, 0x01, 0x0a, 0x0b,
	0x50, 0x6c, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x1f, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0c, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x24, 0x2e, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xd2, 0x01, 0x0a, 0x13, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x42, 0x0f, 0x50, 0x6c, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2d, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x45, 0x58, 0xaa, 0x02, 0x0f, 0x43, 0x68, 0x61, 0x6c,
	0x6b, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x43, 0x68,
	0x61, 0x6c, 0x6b, 0x5c, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b,
	0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x43, 0x68,
	0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_engine_v1_plan_server_proto_rawDescOnce sync.Once
	file_chalk_engine_v1_plan_server_proto_rawDescData = file_chalk_engine_v1_plan_server_proto_rawDesc
)

func file_chalk_engine_v1_plan_server_proto_rawDescGZIP() []byte {
	file_chalk_engine_v1_plan_server_proto_rawDescOnce.Do(func() {
		file_chalk_engine_v1_plan_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_engine_v1_plan_server_proto_rawDescData)
	})
	return file_chalk_engine_v1_plan_server_proto_rawDescData
}

var file_chalk_engine_v1_plan_server_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chalk_engine_v1_plan_server_proto_goTypes = []interface{}{
	(*GetPlanRequest)(nil),         // 0: chalk.engine.v1.GetPlanRequest
	(*GetPlanResponse)(nil),        // 1: chalk.engine.v1.GetPlanResponse
	(*ExecuteQueryRequest)(nil),    // 2: chalk.engine.v1.ExecuteQueryRequest
	(*ExecuteQueryResponse)(nil),   // 3: chalk.engine.v1.ExecuteQueryResponse
	(*v1.OnlineQueryRequest)(nil),  // 4: chalk.common.v1.OnlineQueryRequest
	(*Plan)(nil),                   // 5: chalk.engine.v1.Plan
	(*v1.OnlineQueryResponse)(nil), // 6: chalk.common.v1.OnlineQueryResponse
}
var file_chalk_engine_v1_plan_server_proto_depIdxs = []int32{
	4, // 0: chalk.engine.v1.GetPlanRequest.online_query_request:type_name -> chalk.common.v1.OnlineQueryRequest
	5, // 1: chalk.engine.v1.GetPlanResponse.plan:type_name -> chalk.engine.v1.Plan
	4, // 2: chalk.engine.v1.ExecuteQueryRequest.online_query_request:type_name -> chalk.common.v1.OnlineQueryRequest
	6, // 3: chalk.engine.v1.ExecuteQueryResponse.online_query_response:type_name -> chalk.common.v1.OnlineQueryResponse
	0, // 4: chalk.engine.v1.PlanService.GetPlan:input_type -> chalk.engine.v1.GetPlanRequest
	2, // 5: chalk.engine.v1.PlanService.ExecuteQuery:input_type -> chalk.engine.v1.ExecuteQueryRequest
	1, // 6: chalk.engine.v1.PlanService.GetPlan:output_type -> chalk.engine.v1.GetPlanResponse
	3, // 7: chalk.engine.v1.PlanService.ExecuteQuery:output_type -> chalk.engine.v1.ExecuteQueryResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_chalk_engine_v1_plan_server_proto_init() }
func file_chalk_engine_v1_plan_server_proto_init() {
	if File_chalk_engine_v1_plan_server_proto != nil {
		return
	}
	file_chalk_engine_v1_plan_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chalk_engine_v1_plan_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlanRequest); i {
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
		file_chalk_engine_v1_plan_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlanResponse); i {
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
		file_chalk_engine_v1_plan_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteQueryRequest); i {
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
		file_chalk_engine_v1_plan_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteQueryResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_engine_v1_plan_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chalk_engine_v1_plan_server_proto_goTypes,
		DependencyIndexes: file_chalk_engine_v1_plan_server_proto_depIdxs,
		MessageInfos:      file_chalk_engine_v1_plan_server_proto_msgTypes,
	}.Build()
	File_chalk_engine_v1_plan_server_proto = out.File
	file_chalk_engine_v1_plan_server_proto_rawDesc = nil
	file_chalk_engine_v1_plan_server_proto_goTypes = nil
	file_chalk_engine_v1_plan_server_proto_depIdxs = nil
}
