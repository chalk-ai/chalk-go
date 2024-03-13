// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: chalk/server/v1/team.proto

package serverv1

import (
	v1 "github.com/chalk-ai/chalk-go/gen/chalk/auth/v1"
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

type GetEnvRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetEnvRequest) Reset() {
	*x = GetEnvRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_team_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEnvRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEnvRequest) ProtoMessage() {}

func (x *GetEnvRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_team_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEnvRequest.ProtoReflect.Descriptor instead.
func (*GetEnvRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_team_proto_rawDescGZIP(), []int{0}
}

type GetEnvResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Environment *Environment `protobuf:"bytes,1,opt,name=environment,proto3" json:"environment,omitempty"`
}

func (x *GetEnvResponse) Reset() {
	*x = GetEnvResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_team_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEnvResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEnvResponse) ProtoMessage() {}

func (x *GetEnvResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_team_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEnvResponse.ProtoReflect.Descriptor instead.
func (*GetEnvResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_team_proto_rawDescGZIP(), []int{1}
}

func (x *GetEnvResponse) GetEnvironment() *Environment {
	if x != nil {
		return x.Environment
	}
	return nil
}

type GetEnvironmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
}

func (x *GetEnvironmentsRequest) Reset() {
	*x = GetEnvironmentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_team_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEnvironmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEnvironmentsRequest) ProtoMessage() {}

func (x *GetEnvironmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_team_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEnvironmentsRequest.ProtoReflect.Descriptor instead.
func (*GetEnvironmentsRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_team_proto_rawDescGZIP(), []int{2}
}

func (x *GetEnvironmentsRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

type GetEnvironmentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Environments []*Environment `protobuf:"bytes,2,rep,name=environments,proto3" json:"environments,omitempty"`
}

func (x *GetEnvironmentsResponse) Reset() {
	*x = GetEnvironmentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_team_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEnvironmentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEnvironmentsResponse) ProtoMessage() {}

func (x *GetEnvironmentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_team_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEnvironmentsResponse.ProtoReflect.Descriptor instead.
func (*GetEnvironmentsResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_team_proto_rawDescGZIP(), []int{3}
}

func (x *GetEnvironmentsResponse) GetEnvironments() []*Environment {
	if x != nil {
		return x.Environments
	}
	return nil
}

type GetAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAgentRequest) Reset() {
	*x = GetAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_team_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAgentRequest) ProtoMessage() {}

func (x *GetAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_team_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAgentRequest.ProtoReflect.Descriptor instead.
func (*GetAgentRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_team_proto_rawDescGZIP(), []int{4}
}

type GetAgentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Agent *v1.Agent `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent,omitempty"`
}

func (x *GetAgentResponse) Reset() {
	*x = GetAgentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_team_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAgentResponse) ProtoMessage() {}

func (x *GetAgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_team_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAgentResponse.ProtoReflect.Descriptor instead.
func (*GetAgentResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_team_proto_rawDescGZIP(), []int{5}
}

func (x *GetAgentResponse) GetAgent() *v1.Agent {
	if x != nil {
		return x.Agent
	}
	return nil
}

type GetDisplayAgentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDisplayAgentRequest) Reset() {
	*x = GetDisplayAgentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_team_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDisplayAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDisplayAgentRequest) ProtoMessage() {}

func (x *GetDisplayAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_team_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDisplayAgentRequest.ProtoReflect.Descriptor instead.
func (*GetDisplayAgentRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_team_proto_rawDescGZIP(), []int{6}
}

type GetDisplayAgentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Agent *v1.DisplayAgent `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent,omitempty"`
}

func (x *GetDisplayAgentResponse) Reset() {
	*x = GetDisplayAgentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_team_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDisplayAgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDisplayAgentResponse) ProtoMessage() {}

func (x *GetDisplayAgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_team_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDisplayAgentResponse.ProtoReflect.Descriptor instead.
func (*GetDisplayAgentResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_team_proto_rawDescGZIP(), []int{7}
}

func (x *GetDisplayAgentResponse) GetAgent() *v1.DisplayAgent {
	if x != nil {
		return x.Agent
	}
	return nil
}

var File_chalk_server_v1_team_proto protoreflect.FileDescriptor

var file_chalk_server_v1_team_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x50, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0x32, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x5b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x40, 0x0a, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x05, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x05,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x4c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x32, 0x95, 0x03,
	0x0a, 0x0b, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a,
	0x06, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x76, 0x12, 0x1e, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x76,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x76,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x80, 0x7d, 0x0b, 0x90, 0x02, 0x01,
	0x12, 0x6c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x27, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x80, 0x7d, 0x02, 0x90, 0x02, 0x01, 0x12, 0x57,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x06, 0x80, 0x7d, 0x02, 0x90, 0x02, 0x01, 0x12, 0x6c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x80,
	0x7d, 0x02, 0x90, 0x02, 0x01, 0x42, 0xb9, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x54,
	0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02, 0x0f, 0x43, 0x68,
	0x61, 0x6c, 0x6b, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f,
	0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1b, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11,
	0x43, 0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_server_v1_team_proto_rawDescOnce sync.Once
	file_chalk_server_v1_team_proto_rawDescData = file_chalk_server_v1_team_proto_rawDesc
)

func file_chalk_server_v1_team_proto_rawDescGZIP() []byte {
	file_chalk_server_v1_team_proto_rawDescOnce.Do(func() {
		file_chalk_server_v1_team_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_server_v1_team_proto_rawDescData)
	})
	return file_chalk_server_v1_team_proto_rawDescData
}

var file_chalk_server_v1_team_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_chalk_server_v1_team_proto_goTypes = []interface{}{
	(*GetEnvRequest)(nil),           // 0: chalk.server.v1.GetEnvRequest
	(*GetEnvResponse)(nil),          // 1: chalk.server.v1.GetEnvResponse
	(*GetEnvironmentsRequest)(nil),  // 2: chalk.server.v1.GetEnvironmentsRequest
	(*GetEnvironmentsResponse)(nil), // 3: chalk.server.v1.GetEnvironmentsResponse
	(*GetAgentRequest)(nil),         // 4: chalk.server.v1.GetAgentRequest
	(*GetAgentResponse)(nil),        // 5: chalk.server.v1.GetAgentResponse
	(*GetDisplayAgentRequest)(nil),  // 6: chalk.server.v1.GetDisplayAgentRequest
	(*GetDisplayAgentResponse)(nil), // 7: chalk.server.v1.GetDisplayAgentResponse
	(*Environment)(nil),             // 8: chalk.server.v1.Environment
	(*v1.Agent)(nil),                // 9: chalk.auth.v1.Agent
	(*v1.DisplayAgent)(nil),         // 10: chalk.auth.v1.DisplayAgent
}
var file_chalk_server_v1_team_proto_depIdxs = []int32{
	8,  // 0: chalk.server.v1.GetEnvResponse.environment:type_name -> chalk.server.v1.Environment
	8,  // 1: chalk.server.v1.GetEnvironmentsResponse.environments:type_name -> chalk.server.v1.Environment
	9,  // 2: chalk.server.v1.GetAgentResponse.agent:type_name -> chalk.auth.v1.Agent
	10, // 3: chalk.server.v1.GetDisplayAgentResponse.agent:type_name -> chalk.auth.v1.DisplayAgent
	0,  // 4: chalk.server.v1.TeamService.GetEnv:input_type -> chalk.server.v1.GetEnvRequest
	2,  // 5: chalk.server.v1.TeamService.GetEnvironments:input_type -> chalk.server.v1.GetEnvironmentsRequest
	4,  // 6: chalk.server.v1.TeamService.GetAgent:input_type -> chalk.server.v1.GetAgentRequest
	6,  // 7: chalk.server.v1.TeamService.GetDisplayAgent:input_type -> chalk.server.v1.GetDisplayAgentRequest
	1,  // 8: chalk.server.v1.TeamService.GetEnv:output_type -> chalk.server.v1.GetEnvResponse
	3,  // 9: chalk.server.v1.TeamService.GetEnvironments:output_type -> chalk.server.v1.GetEnvironmentsResponse
	5,  // 10: chalk.server.v1.TeamService.GetAgent:output_type -> chalk.server.v1.GetAgentResponse
	7,  // 11: chalk.server.v1.TeamService.GetDisplayAgent:output_type -> chalk.server.v1.GetDisplayAgentResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_chalk_server_v1_team_proto_init() }
func file_chalk_server_v1_team_proto_init() {
	if File_chalk_server_v1_team_proto != nil {
		return
	}
	file_chalk_server_v1_environment_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chalk_server_v1_team_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEnvRequest); i {
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
		file_chalk_server_v1_team_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEnvResponse); i {
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
		file_chalk_server_v1_team_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEnvironmentsRequest); i {
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
		file_chalk_server_v1_team_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEnvironmentsResponse); i {
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
		file_chalk_server_v1_team_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAgentRequest); i {
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
		file_chalk_server_v1_team_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAgentResponse); i {
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
		file_chalk_server_v1_team_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDisplayAgentRequest); i {
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
		file_chalk_server_v1_team_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDisplayAgentResponse); i {
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
			RawDescriptor: file_chalk_server_v1_team_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chalk_server_v1_team_proto_goTypes,
		DependencyIndexes: file_chalk_server_v1_team_proto_depIdxs,
		MessageInfos:      file_chalk_server_v1_team_proto_msgTypes,
	}.Build()
	File_chalk_server_v1_team_proto = out.File
	file_chalk_server_v1_team_proto_rawDesc = nil
	file_chalk_server_v1_team_proto_goTypes = nil
	file_chalk_server_v1_team_proto_depIdxs = nil
}
