// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: chalk/server/v1/link.proto

package serverv1

import (
	_ "github.com/chalk-ai/chalk-go/gen/chalk/utils/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
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

type LinkSessionStatus int32

const (
	LinkSessionStatus_LINK_SESSION_STATUS_UNSPECIFIED LinkSessionStatus = 0
	LinkSessionStatus_LINK_SESSION_STATUS_PENDING     LinkSessionStatus = 1
	LinkSessionStatus_LINK_SESSION_STATUS_SUCCESS     LinkSessionStatus = 2
	LinkSessionStatus_LINK_SESSION_STATUS_FAILED      LinkSessionStatus = 3
	LinkSessionStatus_LINK_SESSION_STATUS_NOT_FOUND   LinkSessionStatus = 4
	LinkSessionStatus_LINK_SESSION_STATUS_FORBIDDEN   LinkSessionStatus = 5
)

// Enum value maps for LinkSessionStatus.
var (
	LinkSessionStatus_name = map[int32]string{
		0: "LINK_SESSION_STATUS_UNSPECIFIED",
		1: "LINK_SESSION_STATUS_PENDING",
		2: "LINK_SESSION_STATUS_SUCCESS",
		3: "LINK_SESSION_STATUS_FAILED",
		4: "LINK_SESSION_STATUS_NOT_FOUND",
		5: "LINK_SESSION_STATUS_FORBIDDEN",
	}
	LinkSessionStatus_value = map[string]int32{
		"LINK_SESSION_STATUS_UNSPECIFIED": 0,
		"LINK_SESSION_STATUS_PENDING":     1,
		"LINK_SESSION_STATUS_SUCCESS":     2,
		"LINK_SESSION_STATUS_FAILED":      3,
		"LINK_SESSION_STATUS_NOT_FOUND":   4,
		"LINK_SESSION_STATUS_FORBIDDEN":   5,
	}
)

func (x LinkSessionStatus) Enum() *LinkSessionStatus {
	p := new(LinkSessionStatus)
	*p = x
	return p
}

func (x LinkSessionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LinkSessionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_chalk_server_v1_link_proto_enumTypes[0].Descriptor()
}

func (LinkSessionStatus) Type() protoreflect.EnumType {
	return &file_chalk_server_v1_link_proto_enumTypes[0]
}

func (x LinkSessionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LinkSessionStatus.Descriptor instead.
func (LinkSessionStatus) EnumDescriptor() ([]byte, []int) {
	return file_chalk_server_v1_link_proto_rawDescGZIP(), []int{0}
}

type LinkToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name              string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ClientId          string                 `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret      string                 `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	ApiServer         string                 `protobuf:"bytes,4,opt,name=api_server,json=apiServer,proto3" json:"api_server,omitempty"`
	ActiveEnvironment *string                `protobuf:"bytes,5,opt,name=active_environment,json=activeEnvironment,proto3,oneof" json:"active_environment,omitempty"`
	ValidUntil        *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=valid_until,json=validUntil,proto3" json:"valid_until,omitempty"`
}

func (x *LinkToken) Reset() {
	*x = LinkToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkToken) ProtoMessage() {}

func (x *LinkToken) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkToken.ProtoReflect.Descriptor instead.
func (*LinkToken) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_link_proto_rawDescGZIP(), []int{0}
}

func (x *LinkToken) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LinkToken) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *LinkToken) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *LinkToken) GetApiServer() string {
	if x != nil {
		return x.ApiServer
	}
	return ""
}

func (x *LinkToken) GetActiveEnvironment() string {
	if x != nil && x.ActiveEnvironment != nil {
		return *x.ActiveEnvironment
	}
	return ""
}

func (x *LinkToken) GetValidUntil() *timestamppb.Timestamp {
	if x != nil {
		return x.ValidUntil
	}
	return nil
}

// GET LINK
type GetLinkSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LinkCode    string `protobuf:"bytes,1,opt,name=link_code,json=linkCode,proto3" json:"link_code,omitempty"`
	ProjectName string `protobuf:"bytes,2,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
}

func (x *GetLinkSessionRequest) Reset() {
	*x = GetLinkSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_link_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLinkSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLinkSessionRequest) ProtoMessage() {}

func (x *GetLinkSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_link_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLinkSessionRequest.ProtoReflect.Descriptor instead.
func (*GetLinkSessionRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_link_proto_rawDescGZIP(), []int{1}
}

func (x *GetLinkSessionRequest) GetLinkCode() string {
	if x != nil {
		return x.LinkCode
	}
	return ""
}

func (x *GetLinkSessionRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

type GetLinkSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  LinkSessionStatus `protobuf:"varint,1,opt,name=status,proto3,enum=chalk.server.v1.LinkSessionStatus" json:"status,omitempty"`
	Message string            `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Token   *LinkToken        `protobuf:"bytes,3,opt,name=token,proto3,oneof" json:"token,omitempty"`
}

func (x *GetLinkSessionResponse) Reset() {
	*x = GetLinkSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_link_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLinkSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLinkSessionResponse) ProtoMessage() {}

func (x *GetLinkSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_link_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLinkSessionResponse.ProtoReflect.Descriptor instead.
func (*GetLinkSessionResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_link_proto_rawDescGZIP(), []int{2}
}

func (x *GetLinkSessionResponse) GetStatus() LinkSessionStatus {
	if x != nil {
		return x.Status
	}
	return LinkSessionStatus_LINK_SESSION_STATUS_UNSPECIFIED
}

func (x *GetLinkSessionResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetLinkSessionResponse) GetToken() *LinkToken {
	if x != nil {
		return x.Token
	}
	return nil
}

// CREATE LINK
type CreateLinkSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateLinkSessionRequest) Reset() {
	*x = CreateLinkSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_link_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLinkSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLinkSessionRequest) ProtoMessage() {}

func (x *CreateLinkSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_link_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLinkSessionRequest.ProtoReflect.Descriptor instead.
func (*CreateLinkSessionRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_link_proto_rawDescGZIP(), []int{3}
}

type CreateLinkSessionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LinkCode  string                 `protobuf:"bytes,1,opt,name=link_code,json=linkCode,proto3" json:"link_code,omitempty"`
	AuthLink  string                 `protobuf:"bytes,2,opt,name=auth_link,json=authLink,proto3" json:"auth_link,omitempty"`
	ExpiresAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *CreateLinkSessionResponse) Reset() {
	*x = CreateLinkSessionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_link_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLinkSessionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLinkSessionResponse) ProtoMessage() {}

func (x *CreateLinkSessionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_link_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLinkSessionResponse.ProtoReflect.Descriptor instead.
func (*CreateLinkSessionResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_link_proto_rawDescGZIP(), []int{4}
}

func (x *CreateLinkSessionResponse) GetLinkCode() string {
	if x != nil {
		return x.LinkCode
	}
	return ""
}

func (x *CreateLinkSessionResponse) GetAuthLink() string {
	if x != nil {
		return x.AuthLink
	}
	return ""
}

func (x *CreateLinkSessionResponse) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

var File_chalk_server_v1_link_proto protoreflect.FileDescriptor

var file_chalk_server_v1_link_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8e, 0x02, 0x0a, 0x09, 0x4c, 0x69, 0x6e, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x29, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xd8, 0xa1, 0x27, 0x01, 0x52, 0x0c, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70,
	0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x12, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x11, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x45,
	0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a,
	0x0b, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x55, 0x6e, 0x74, 0x69, 0x6c, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0x5d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x6c, 0x69,
	0x6e, 0x6b, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xd8,
	0xa1, 0x27, 0x01, 0x52, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0xaf, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x6e, 0x6b, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x35, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x48, 0x00, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x1a, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x9c,
	0x01, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x09,
	0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x04, 0xd8, 0xa1, 0x27, 0x01, 0x52, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x21, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x04, 0xd8, 0xa1, 0x27, 0x01, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x2a, 0xe0, 0x01,
	0x0a, 0x11, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x1f, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x53, 0x45, 0x53, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x4c, 0x49, 0x4e, 0x4b,
	0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x4c, 0x49, 0x4e,
	0x4b, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x4c, 0x49,
	0x4e, 0x4b, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x21, 0x0a, 0x1d, 0x4c, 0x49,
	0x4e, 0x4b, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x04, 0x12, 0x21, 0x0a,
	0x1d, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x53, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x05,
	0x42, 0xb9, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b,
	0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x43, 0x68, 0x61,
	0x6c, 0x6b, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x43, 0x68, 0x61, 0x6c, 0x6b,
	0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_server_v1_link_proto_rawDescOnce sync.Once
	file_chalk_server_v1_link_proto_rawDescData = file_chalk_server_v1_link_proto_rawDesc
)

func file_chalk_server_v1_link_proto_rawDescGZIP() []byte {
	file_chalk_server_v1_link_proto_rawDescOnce.Do(func() {
		file_chalk_server_v1_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_server_v1_link_proto_rawDescData)
	})
	return file_chalk_server_v1_link_proto_rawDescData
}

var file_chalk_server_v1_link_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_chalk_server_v1_link_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_chalk_server_v1_link_proto_goTypes = []interface{}{
	(LinkSessionStatus)(0),            // 0: chalk.server.v1.LinkSessionStatus
	(*LinkToken)(nil),                 // 1: chalk.server.v1.LinkToken
	(*GetLinkSessionRequest)(nil),     // 2: chalk.server.v1.GetLinkSessionRequest
	(*GetLinkSessionResponse)(nil),    // 3: chalk.server.v1.GetLinkSessionResponse
	(*CreateLinkSessionRequest)(nil),  // 4: chalk.server.v1.CreateLinkSessionRequest
	(*CreateLinkSessionResponse)(nil), // 5: chalk.server.v1.CreateLinkSessionResponse
	(*timestamppb.Timestamp)(nil),     // 6: google.protobuf.Timestamp
}
var file_chalk_server_v1_link_proto_depIdxs = []int32{
	6, // 0: chalk.server.v1.LinkToken.valid_until:type_name -> google.protobuf.Timestamp
	0, // 1: chalk.server.v1.GetLinkSessionResponse.status:type_name -> chalk.server.v1.LinkSessionStatus
	1, // 2: chalk.server.v1.GetLinkSessionResponse.token:type_name -> chalk.server.v1.LinkToken
	6, // 3: chalk.server.v1.CreateLinkSessionResponse.expires_at:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_chalk_server_v1_link_proto_init() }
func file_chalk_server_v1_link_proto_init() {
	if File_chalk_server_v1_link_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chalk_server_v1_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkToken); i {
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
		file_chalk_server_v1_link_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLinkSessionRequest); i {
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
		file_chalk_server_v1_link_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLinkSessionResponse); i {
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
		file_chalk_server_v1_link_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLinkSessionRequest); i {
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
		file_chalk_server_v1_link_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLinkSessionResponse); i {
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
	file_chalk_server_v1_link_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_chalk_server_v1_link_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_server_v1_link_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chalk_server_v1_link_proto_goTypes,
		DependencyIndexes: file_chalk_server_v1_link_proto_depIdxs,
		EnumInfos:         file_chalk_server_v1_link_proto_enumTypes,
		MessageInfos:      file_chalk_server_v1_link_proto_msgTypes,
	}.Build()
	File_chalk_server_v1_link_proto = out.File
	file_chalk_server_v1_link_proto_rawDesc = nil
	file_chalk_server_v1_link_proto_goTypes = nil
	file_chalk_server_v1_link_proto_depIdxs = nil
}
