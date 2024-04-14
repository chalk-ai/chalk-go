// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: chalk/server/v1/authtesting.proto

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

type GetUnauthedTestEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUnauthedTestEndpointRequest) Reset() {
	*x = GetUnauthedTestEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_authtesting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUnauthedTestEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUnauthedTestEndpointRequest) ProtoMessage() {}

func (x *GetUnauthedTestEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_authtesting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUnauthedTestEndpointRequest.ProtoReflect.Descriptor instead.
func (*GetUnauthedTestEndpointRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_authtesting_proto_rawDescGZIP(), []int{0}
}

type GetAuthedTestEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAuthedTestEndpointRequest) Reset() {
	*x = GetAuthedTestEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_authtesting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthedTestEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthedTestEndpointRequest) ProtoMessage() {}

func (x *GetAuthedTestEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_authtesting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthedTestEndpointRequest.ProtoReflect.Descriptor instead.
func (*GetAuthedTestEndpointRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_authtesting_proto_rawDescGZIP(), []int{1}
}

type GetViewerTestEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetViewerTestEndpointRequest) Reset() {
	*x = GetViewerTestEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_authtesting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetViewerTestEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetViewerTestEndpointRequest) ProtoMessage() {}

func (x *GetViewerTestEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_authtesting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetViewerTestEndpointRequest.ProtoReflect.Descriptor instead.
func (*GetViewerTestEndpointRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_authtesting_proto_rawDescGZIP(), []int{2}
}

type GetDataScientistTestEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDataScientistTestEndpointRequest) Reset() {
	*x = GetDataScientistTestEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_authtesting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataScientistTestEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataScientistTestEndpointRequest) ProtoMessage() {}

func (x *GetDataScientistTestEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_authtesting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataScientistTestEndpointRequest.ProtoReflect.Descriptor instead.
func (*GetDataScientistTestEndpointRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_authtesting_proto_rawDescGZIP(), []int{3}
}

type GetDeveloperTestEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDeveloperTestEndpointRequest) Reset() {
	*x = GetDeveloperTestEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_authtesting_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeveloperTestEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeveloperTestEndpointRequest) ProtoMessage() {}

func (x *GetDeveloperTestEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_authtesting_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeveloperTestEndpointRequest.ProtoReflect.Descriptor instead.
func (*GetDeveloperTestEndpointRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_authtesting_proto_rawDescGZIP(), []int{4}
}

type GetAdminTestEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAdminTestEndpointRequest) Reset() {
	*x = GetAdminTestEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_authtesting_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminTestEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminTestEndpointRequest) ProtoMessage() {}

func (x *GetAdminTestEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_authtesting_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminTestEndpointRequest.ProtoReflect.Descriptor instead.
func (*GetAdminTestEndpointRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_authtesting_proto_rawDescGZIP(), []int{5}
}

type GetOwnerTestEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOwnerTestEndpointRequest) Reset() {
	*x = GetOwnerTestEndpointRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_authtesting_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOwnerTestEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOwnerTestEndpointRequest) ProtoMessage() {}

func (x *GetOwnerTestEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_authtesting_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOwnerTestEndpointRequest.ProtoReflect.Descriptor instead.
func (*GetOwnerTestEndpointRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_authtesting_proto_rawDescGZIP(), []int{6}
}

type GetUnauthedTestEndpointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUnauthedTestEndpointResponse) Reset() {
	*x = GetUnauthedTestEndpointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_authtesting_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUnauthedTestEndpointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUnauthedTestEndpointResponse) ProtoMessage() {}

func (x *GetUnauthedTestEndpointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_authtesting_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUnauthedTestEndpointResponse.ProtoReflect.Descriptor instead.
func (*GetUnauthedTestEndpointResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_authtesting_proto_rawDescGZIP(), []int{7}
}

type GetAuthedTestEndpointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAuthedTestEndpointResponse) Reset() {
	*x = GetAuthedTestEndpointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_authtesting_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthedTestEndpointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthedTestEndpointResponse) ProtoMessage() {}

func (x *GetAuthedTestEndpointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_authtesting_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthedTestEndpointResponse.ProtoReflect.Descriptor instead.
func (*GetAuthedTestEndpointResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_authtesting_proto_rawDescGZIP(), []int{8}
}

type GetViewerTestEndpointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetViewerTestEndpointResponse) Reset() {
	*x = GetViewerTestEndpointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_authtesting_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetViewerTestEndpointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetViewerTestEndpointResponse) ProtoMessage() {}

func (x *GetViewerTestEndpointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_authtesting_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetViewerTestEndpointResponse.ProtoReflect.Descriptor instead.
func (*GetViewerTestEndpointResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_authtesting_proto_rawDescGZIP(), []int{9}
}

type GetDataScientistTestEndpointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDataScientistTestEndpointResponse) Reset() {
	*x = GetDataScientistTestEndpointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_authtesting_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataScientistTestEndpointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataScientistTestEndpointResponse) ProtoMessage() {}

func (x *GetDataScientistTestEndpointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_authtesting_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataScientistTestEndpointResponse.ProtoReflect.Descriptor instead.
func (*GetDataScientistTestEndpointResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_authtesting_proto_rawDescGZIP(), []int{10}
}

type GetDeveloperTestEndpointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDeveloperTestEndpointResponse) Reset() {
	*x = GetDeveloperTestEndpointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_authtesting_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeveloperTestEndpointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeveloperTestEndpointResponse) ProtoMessage() {}

func (x *GetDeveloperTestEndpointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_authtesting_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeveloperTestEndpointResponse.ProtoReflect.Descriptor instead.
func (*GetDeveloperTestEndpointResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_authtesting_proto_rawDescGZIP(), []int{11}
}

type GetAdminTestEndpointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAdminTestEndpointResponse) Reset() {
	*x = GetAdminTestEndpointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_authtesting_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminTestEndpointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminTestEndpointResponse) ProtoMessage() {}

func (x *GetAdminTestEndpointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_authtesting_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminTestEndpointResponse.ProtoReflect.Descriptor instead.
func (*GetAdminTestEndpointResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_authtesting_proto_rawDescGZIP(), []int{12}
}

type GetOwnerTestEndpointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOwnerTestEndpointResponse) Reset() {
	*x = GetOwnerTestEndpointResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_authtesting_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOwnerTestEndpointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOwnerTestEndpointResponse) ProtoMessage() {}

func (x *GetOwnerTestEndpointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_authtesting_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOwnerTestEndpointResponse.ProtoReflect.Descriptor instead.
func (*GetOwnerTestEndpointResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_authtesting_proto_rawDescGZIP(), []int{13}
}

var File_chalk_server_v1_authtesting_proto protoreflect.FileDescriptor

var file_chalk_server_v1_authtesting_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x56, 0x69,
	0x65, 0x77, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x25, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x21,
	0x0a, 0x1f, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x54, 0x65,
	0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x1d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x54, 0x65, 0x73,
	0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x1d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x21, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x64, 0x54, 0x65,
	0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1f, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x64, 0x54,
	0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1f, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72,
	0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x0a, 0x24, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x63, 0x69, 0x65, 0x6e, 0x74, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x0a, 0x20,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x54, 0x65, 0x73, 0x74,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xb5, 0x07, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55,
	0x6e, 0x61, 0x75, 0x74, 0x68, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x2f, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x65,
	0x64, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x80, 0x7d, 0x01, 0x90, 0x02, 0x01, 0x12, 0x7e,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x80, 0x7d, 0x02, 0x90, 0x02, 0x01, 0x12, 0x7e,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x65,
	0x77, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77,
	0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x80, 0x7d, 0x0b, 0x90, 0x02, 0x01, 0x12, 0x93,
	0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x69, 0x65, 0x6e, 0x74,
	0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x34, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x69,
	0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x63, 0x69, 0x65, 0x6e, 0x74, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x80, 0x7d,
	0x03, 0x90, 0x02, 0x01, 0x12, 0x87, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x12, 0x30, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72,
	0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x80, 0x7d, 0x11, 0x90, 0x02, 0x01, 0x12, 0x7b,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2c, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x54,
	0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x06, 0x80, 0x7d, 0x0c, 0x90, 0x02, 0x01, 0x12, 0x7b, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x2c, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x65,
	0x73, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x65, 0x73, 0x74,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x06, 0x80, 0x7d, 0x0a, 0x90, 0x02, 0x01, 0x42, 0xc0, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d,
	0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x42, 0x10, 0x41, 0x75, 0x74, 0x68, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d,
	0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x43, 0x68, 0x61, 0x6c,
	0x6b, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x3a,
	0x3a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_chalk_server_v1_authtesting_proto_rawDescOnce sync.Once
	file_chalk_server_v1_authtesting_proto_rawDescData = file_chalk_server_v1_authtesting_proto_rawDesc
)

func file_chalk_server_v1_authtesting_proto_rawDescGZIP() []byte {
	file_chalk_server_v1_authtesting_proto_rawDescOnce.Do(func() {
		file_chalk_server_v1_authtesting_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_server_v1_authtesting_proto_rawDescData)
	})
	return file_chalk_server_v1_authtesting_proto_rawDescData
}

var file_chalk_server_v1_authtesting_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_chalk_server_v1_authtesting_proto_goTypes = []interface{}{
	(*GetUnauthedTestEndpointRequest)(nil),       // 0: chalk.server.v1.GetUnauthedTestEndpointRequest
	(*GetAuthedTestEndpointRequest)(nil),         // 1: chalk.server.v1.GetAuthedTestEndpointRequest
	(*GetViewerTestEndpointRequest)(nil),         // 2: chalk.server.v1.GetViewerTestEndpointRequest
	(*GetDataScientistTestEndpointRequest)(nil),  // 3: chalk.server.v1.GetDataScientistTestEndpointRequest
	(*GetDeveloperTestEndpointRequest)(nil),      // 4: chalk.server.v1.GetDeveloperTestEndpointRequest
	(*GetAdminTestEndpointRequest)(nil),          // 5: chalk.server.v1.GetAdminTestEndpointRequest
	(*GetOwnerTestEndpointRequest)(nil),          // 6: chalk.server.v1.GetOwnerTestEndpointRequest
	(*GetUnauthedTestEndpointResponse)(nil),      // 7: chalk.server.v1.GetUnauthedTestEndpointResponse
	(*GetAuthedTestEndpointResponse)(nil),        // 8: chalk.server.v1.GetAuthedTestEndpointResponse
	(*GetViewerTestEndpointResponse)(nil),        // 9: chalk.server.v1.GetViewerTestEndpointResponse
	(*GetDataScientistTestEndpointResponse)(nil), // 10: chalk.server.v1.GetDataScientistTestEndpointResponse
	(*GetDeveloperTestEndpointResponse)(nil),     // 11: chalk.server.v1.GetDeveloperTestEndpointResponse
	(*GetAdminTestEndpointResponse)(nil),         // 12: chalk.server.v1.GetAdminTestEndpointResponse
	(*GetOwnerTestEndpointResponse)(nil),         // 13: chalk.server.v1.GetOwnerTestEndpointResponse
}
var file_chalk_server_v1_authtesting_proto_depIdxs = []int32{
	0,  // 0: chalk.server.v1.AuthTestingService.GetUnauthedTestEndpoint:input_type -> chalk.server.v1.GetUnauthedTestEndpointRequest
	1,  // 1: chalk.server.v1.AuthTestingService.GetAuthedTestEndpoint:input_type -> chalk.server.v1.GetAuthedTestEndpointRequest
	2,  // 2: chalk.server.v1.AuthTestingService.GetViewerTestEndpoint:input_type -> chalk.server.v1.GetViewerTestEndpointRequest
	3,  // 3: chalk.server.v1.AuthTestingService.GetDataScientistTestEndpoint:input_type -> chalk.server.v1.GetDataScientistTestEndpointRequest
	4,  // 4: chalk.server.v1.AuthTestingService.GetDeveloperTestEndpoint:input_type -> chalk.server.v1.GetDeveloperTestEndpointRequest
	5,  // 5: chalk.server.v1.AuthTestingService.GetAdminTestEndpoint:input_type -> chalk.server.v1.GetAdminTestEndpointRequest
	6,  // 6: chalk.server.v1.AuthTestingService.GetOwnerTestEndpoint:input_type -> chalk.server.v1.GetOwnerTestEndpointRequest
	7,  // 7: chalk.server.v1.AuthTestingService.GetUnauthedTestEndpoint:output_type -> chalk.server.v1.GetUnauthedTestEndpointResponse
	8,  // 8: chalk.server.v1.AuthTestingService.GetAuthedTestEndpoint:output_type -> chalk.server.v1.GetAuthedTestEndpointResponse
	9,  // 9: chalk.server.v1.AuthTestingService.GetViewerTestEndpoint:output_type -> chalk.server.v1.GetViewerTestEndpointResponse
	10, // 10: chalk.server.v1.AuthTestingService.GetDataScientistTestEndpoint:output_type -> chalk.server.v1.GetDataScientistTestEndpointResponse
	11, // 11: chalk.server.v1.AuthTestingService.GetDeveloperTestEndpoint:output_type -> chalk.server.v1.GetDeveloperTestEndpointResponse
	12, // 12: chalk.server.v1.AuthTestingService.GetAdminTestEndpoint:output_type -> chalk.server.v1.GetAdminTestEndpointResponse
	13, // 13: chalk.server.v1.AuthTestingService.GetOwnerTestEndpoint:output_type -> chalk.server.v1.GetOwnerTestEndpointResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_chalk_server_v1_authtesting_proto_init() }
func file_chalk_server_v1_authtesting_proto_init() {
	if File_chalk_server_v1_authtesting_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chalk_server_v1_authtesting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUnauthedTestEndpointRequest); i {
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
		file_chalk_server_v1_authtesting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthedTestEndpointRequest); i {
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
		file_chalk_server_v1_authtesting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetViewerTestEndpointRequest); i {
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
		file_chalk_server_v1_authtesting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataScientistTestEndpointRequest); i {
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
		file_chalk_server_v1_authtesting_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeveloperTestEndpointRequest); i {
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
		file_chalk_server_v1_authtesting_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdminTestEndpointRequest); i {
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
		file_chalk_server_v1_authtesting_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOwnerTestEndpointRequest); i {
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
		file_chalk_server_v1_authtesting_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUnauthedTestEndpointResponse); i {
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
		file_chalk_server_v1_authtesting_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthedTestEndpointResponse); i {
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
		file_chalk_server_v1_authtesting_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetViewerTestEndpointResponse); i {
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
		file_chalk_server_v1_authtesting_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataScientistTestEndpointResponse); i {
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
		file_chalk_server_v1_authtesting_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeveloperTestEndpointResponse); i {
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
		file_chalk_server_v1_authtesting_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdminTestEndpointResponse); i {
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
		file_chalk_server_v1_authtesting_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOwnerTestEndpointResponse); i {
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
			RawDescriptor: file_chalk_server_v1_authtesting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chalk_server_v1_authtesting_proto_goTypes,
		DependencyIndexes: file_chalk_server_v1_authtesting_proto_depIdxs,
		MessageInfos:      file_chalk_server_v1_authtesting_proto_msgTypes,
	}.Build()
	File_chalk_server_v1_authtesting_proto = out.File
	file_chalk_server_v1_authtesting_proto_rawDesc = nil
	file_chalk_server_v1_authtesting_proto_goTypes = nil
	file_chalk_server_v1_authtesting_proto_depIdxs = nil
}