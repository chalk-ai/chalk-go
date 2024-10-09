// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: chalk/server/v1/builder.proto

package serverv1

import (
	_ "github.com/chalk-ai/chalk-go/gen/chalk/auth/v1"
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

type ActivateDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExistingDeploymentId string `protobuf:"bytes,1,opt,name=existing_deployment_id,json=existingDeploymentId,proto3" json:"existing_deployment_id,omitempty"`
}

func (x *ActivateDeploymentRequest) Reset() {
	*x = ActivateDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_builder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateDeploymentRequest) ProtoMessage() {}

func (x *ActivateDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_builder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateDeploymentRequest.ProtoReflect.Descriptor instead.
func (*ActivateDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_builder_proto_rawDescGZIP(), []int{0}
}

func (x *ActivateDeploymentRequest) GetExistingDeploymentId() string {
	if x != nil {
		return x.ExistingDeploymentId
	}
	return ""
}

type ActivateDeploymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ActivateDeploymentResponse) Reset() {
	*x = ActivateDeploymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_builder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateDeploymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateDeploymentResponse) ProtoMessage() {}

func (x *ActivateDeploymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_builder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateDeploymentResponse.ProtoReflect.Descriptor instead.
func (*ActivateDeploymentResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_builder_proto_rawDescGZIP(), []int{1}
}

type IndexDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExistingDeploymentId string `protobuf:"bytes,1,opt,name=existing_deployment_id,json=existingDeploymentId,proto3" json:"existing_deployment_id,omitempty"`
}

func (x *IndexDeploymentRequest) Reset() {
	*x = IndexDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_builder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexDeploymentRequest) ProtoMessage() {}

func (x *IndexDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_builder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexDeploymentRequest.ProtoReflect.Descriptor instead.
func (*IndexDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_builder_proto_rawDescGZIP(), []int{2}
}

func (x *IndexDeploymentRequest) GetExistingDeploymentId() string {
	if x != nil {
		return x.ExistingDeploymentId
	}
	return ""
}

type IndexDeploymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *IndexDeploymentResponse) Reset() {
	*x = IndexDeploymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_builder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexDeploymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexDeploymentResponse) ProtoMessage() {}

func (x *IndexDeploymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_builder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexDeploymentResponse.ProtoReflect.Descriptor instead.
func (*IndexDeploymentResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_builder_proto_rawDescGZIP(), []int{3}
}

type DeployKubeComponentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExistingDeploymentId string `protobuf:"bytes,1,opt,name=existing_deployment_id,json=existingDeploymentId,proto3" json:"existing_deployment_id,omitempty"`
}

func (x *DeployKubeComponentsRequest) Reset() {
	*x = DeployKubeComponentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_builder_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployKubeComponentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployKubeComponentsRequest) ProtoMessage() {}

func (x *DeployKubeComponentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_builder_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployKubeComponentsRequest.ProtoReflect.Descriptor instead.
func (*DeployKubeComponentsRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_builder_proto_rawDescGZIP(), []int{4}
}

func (x *DeployKubeComponentsRequest) GetExistingDeploymentId() string {
	if x != nil {
		return x.ExistingDeploymentId
	}
	return ""
}

type DeployKubeComponentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeployKubeComponentsResponse) Reset() {
	*x = DeployKubeComponentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_builder_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployKubeComponentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployKubeComponentsResponse) ProtoMessage() {}

func (x *DeployKubeComponentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_builder_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployKubeComponentsResponse.ProtoReflect.Descriptor instead.
func (*DeployKubeComponentsResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_builder_proto_rawDescGZIP(), []int{5}
}

var File_chalk_server_v1_builder_proto protoreflect.FileDescriptor

var file_chalk_server_v1_builder_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x51, 0x0a, 0x19, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34,
	0x0a, 0x16, 0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x22, 0x1c, 0x0a, 0x1a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x4e, 0x0a, 0x16, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x16,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x65, 0x78,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x0a,
	0x1b, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x4b, 0x75, 0x62, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x16,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x65, 0x78,
	0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x22, 0x1e, 0x0a, 0x1c, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x4b, 0x75, 0x62, 0x65,
	0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x32, 0xe9, 0x02, 0x0a, 0x0e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x72, 0x0a, 0x12, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x2e, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x80, 0x7d, 0x0e, 0x12, 0x69, 0x0a, 0x0f, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x2e, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x03, 0x80, 0x7d, 0x0e, 0x12, 0x78, 0x0a, 0x14, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x4b, 0x75,
	0x62, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2c, 0x2e, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x4b, 0x75, 0x62, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x4b, 0x75, 0x62, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x80, 0x7d, 0x0c, 0x42, 0xbc,
	0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x43, 0x68, 0x61, 0x6c,
	0x6b, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x43, 0x68,
	0x61, 0x6c, 0x6b, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x43, 0x68, 0x61, 0x6c,
	0x6b, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_server_v1_builder_proto_rawDescOnce sync.Once
	file_chalk_server_v1_builder_proto_rawDescData = file_chalk_server_v1_builder_proto_rawDesc
)

func file_chalk_server_v1_builder_proto_rawDescGZIP() []byte {
	file_chalk_server_v1_builder_proto_rawDescOnce.Do(func() {
		file_chalk_server_v1_builder_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_server_v1_builder_proto_rawDescData)
	})
	return file_chalk_server_v1_builder_proto_rawDescData
}

var file_chalk_server_v1_builder_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_chalk_server_v1_builder_proto_goTypes = []any{
	(*ActivateDeploymentRequest)(nil),    // 0: chalk.server.v1.ActivateDeploymentRequest
	(*ActivateDeploymentResponse)(nil),   // 1: chalk.server.v1.ActivateDeploymentResponse
	(*IndexDeploymentRequest)(nil),       // 2: chalk.server.v1.IndexDeploymentRequest
	(*IndexDeploymentResponse)(nil),      // 3: chalk.server.v1.IndexDeploymentResponse
	(*DeployKubeComponentsRequest)(nil),  // 4: chalk.server.v1.DeployKubeComponentsRequest
	(*DeployKubeComponentsResponse)(nil), // 5: chalk.server.v1.DeployKubeComponentsResponse
}
var file_chalk_server_v1_builder_proto_depIdxs = []int32{
	0, // 0: chalk.server.v1.BuilderService.ActivateDeployment:input_type -> chalk.server.v1.ActivateDeploymentRequest
	2, // 1: chalk.server.v1.BuilderService.IndexDeployment:input_type -> chalk.server.v1.IndexDeploymentRequest
	4, // 2: chalk.server.v1.BuilderService.DeployKubeComponents:input_type -> chalk.server.v1.DeployKubeComponentsRequest
	1, // 3: chalk.server.v1.BuilderService.ActivateDeployment:output_type -> chalk.server.v1.ActivateDeploymentResponse
	3, // 4: chalk.server.v1.BuilderService.IndexDeployment:output_type -> chalk.server.v1.IndexDeploymentResponse
	5, // 5: chalk.server.v1.BuilderService.DeployKubeComponents:output_type -> chalk.server.v1.DeployKubeComponentsResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chalk_server_v1_builder_proto_init() }
func file_chalk_server_v1_builder_proto_init() {
	if File_chalk_server_v1_builder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chalk_server_v1_builder_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ActivateDeploymentRequest); i {
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
		file_chalk_server_v1_builder_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ActivateDeploymentResponse); i {
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
		file_chalk_server_v1_builder_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*IndexDeploymentRequest); i {
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
		file_chalk_server_v1_builder_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*IndexDeploymentResponse); i {
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
		file_chalk_server_v1_builder_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeployKubeComponentsRequest); i {
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
		file_chalk_server_v1_builder_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeployKubeComponentsResponse); i {
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
			RawDescriptor: file_chalk_server_v1_builder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chalk_server_v1_builder_proto_goTypes,
		DependencyIndexes: file_chalk_server_v1_builder_proto_depIdxs,
		MessageInfos:      file_chalk_server_v1_builder_proto_msgTypes,
	}.Build()
	File_chalk_server_v1_builder_proto = out.File
	file_chalk_server_v1_builder_proto_rawDesc = nil
	file_chalk_server_v1_builder_proto_goTypes = nil
	file_chalk_server_v1_builder_proto_depIdxs = nil
}
