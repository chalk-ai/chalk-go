// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: chalk/server/v1/flag.proto

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

type FlagScope int32

const (
	FlagScope_FLAG_SCOPE_UNSPECIFIED FlagScope = 0
	FlagScope_FLAG_SCOPE_TEAM        FlagScope = 1
	FlagScope_FLAG_SCOPE_ENVIRONMENT FlagScope = 2
)

// Enum value maps for FlagScope.
var (
	FlagScope_name = map[int32]string{
		0: "FLAG_SCOPE_UNSPECIFIED",
		1: "FLAG_SCOPE_TEAM",
		2: "FLAG_SCOPE_ENVIRONMENT",
	}
	FlagScope_value = map[string]int32{
		"FLAG_SCOPE_UNSPECIFIED": 0,
		"FLAG_SCOPE_TEAM":        1,
		"FLAG_SCOPE_ENVIRONMENT": 2,
	}
)

func (x FlagScope) Enum() *FlagScope {
	p := new(FlagScope)
	*p = x
	return p
}

func (x FlagScope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FlagScope) Descriptor() protoreflect.EnumDescriptor {
	return file_chalk_server_v1_flag_proto_enumTypes[0].Descriptor()
}

func (FlagScope) Type() protoreflect.EnumType {
	return &file_chalk_server_v1_flag_proto_enumTypes[0]
}

func (x FlagScope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FlagScope.Descriptor instead.
func (FlagScope) EnumDescriptor() ([]byte, []int) {
	return file_chalk_server_v1_flag_proto_rawDescGZIP(), []int{0}
}

type FeatureFlagValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag  string `protobuf:"bytes,1,opt,name=flag,proto3" json:"flag,omitempty"`
	Value bool   `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *FeatureFlagValue) Reset() {
	*x = FeatureFlagValue{}
	mi := &file_chalk_server_v1_flag_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeatureFlagValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureFlagValue) ProtoMessage() {}

func (x *FeatureFlagValue) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_flag_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureFlagValue.ProtoReflect.Descriptor instead.
func (*FeatureFlagValue) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_flag_proto_rawDescGZIP(), []int{0}
}

func (x *FeatureFlagValue) GetFlag() string {
	if x != nil {
		return x.Flag
	}
	return ""
}

func (x *FeatureFlagValue) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type GetFeatureFlagsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetFeatureFlagsRequest) Reset() {
	*x = GetFeatureFlagsRequest{}
	mi := &file_chalk_server_v1_flag_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFeatureFlagsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeatureFlagsRequest) ProtoMessage() {}

func (x *GetFeatureFlagsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_flag_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeatureFlagsRequest.ProtoReflect.Descriptor instead.
func (*GetFeatureFlagsRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_flag_proto_rawDescGZIP(), []int{1}
}

type GetFeatureFlagsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flags []*FeatureFlagValue `protobuf:"bytes,1,rep,name=flags,proto3" json:"flags,omitempty"`
}

func (x *GetFeatureFlagsResponse) Reset() {
	*x = GetFeatureFlagsResponse{}
	mi := &file_chalk_server_v1_flag_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetFeatureFlagsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeatureFlagsResponse) ProtoMessage() {}

func (x *GetFeatureFlagsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_flag_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeatureFlagsResponse.ProtoReflect.Descriptor instead.
func (*GetFeatureFlagsResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_flag_proto_rawDescGZIP(), []int{2}
}

func (x *GetFeatureFlagsResponse) GetFlags() []*FeatureFlagValue {
	if x != nil {
		return x.Flags
	}
	return nil
}

type SetFeatureFlagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag  string    `protobuf:"bytes,1,opt,name=flag,proto3" json:"flag,omitempty"`
	Value bool      `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	Scope FlagScope `protobuf:"varint,3,opt,name=scope,proto3,enum=chalk.server.v1.FlagScope" json:"scope,omitempty"`
}

func (x *SetFeatureFlagRequest) Reset() {
	*x = SetFeatureFlagRequest{}
	mi := &file_chalk_server_v1_flag_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetFeatureFlagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetFeatureFlagRequest) ProtoMessage() {}

func (x *SetFeatureFlagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_flag_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetFeatureFlagRequest.ProtoReflect.Descriptor instead.
func (*SetFeatureFlagRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_flag_proto_rawDescGZIP(), []int{3}
}

func (x *SetFeatureFlagRequest) GetFlag() string {
	if x != nil {
		return x.Flag
	}
	return ""
}

func (x *SetFeatureFlagRequest) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

func (x *SetFeatureFlagRequest) GetScope() FlagScope {
	if x != nil {
		return x.Scope
	}
	return FlagScope_FLAG_SCOPE_UNSPECIFIED
}

type SetFeatureFlagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetFeatureFlagResponse) Reset() {
	*x = SetFeatureFlagResponse{}
	mi := &file_chalk_server_v1_flag_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetFeatureFlagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetFeatureFlagResponse) ProtoMessage() {}

func (x *SetFeatureFlagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_flag_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetFeatureFlagResponse.ProtoReflect.Descriptor instead.
func (*SetFeatureFlagResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_flag_proto_rawDescGZIP(), []int{4}
}

var File_chalk_server_v1_flag_proto protoreflect.FileDescriptor

var file_chalk_server_v1_flag_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x66, 0x6c, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c,
	0x0a, 0x10, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x18, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x52, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x37, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x22, 0x73, 0x0a, 0x15, 0x53, 0x65,
	0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x30, 0x0a,
	0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x6c, 0x61, 0x67, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x22,
	0x18, 0x0a, 0x16, 0x53, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x58, 0x0a, 0x09, 0x46, 0x6c, 0x61,
	0x67, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x46, 0x4c, 0x41, 0x47, 0x5f, 0x53,
	0x43, 0x4f, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x46, 0x4c, 0x41, 0x47, 0x5f, 0x53, 0x43, 0x4f, 0x50, 0x45,
	0x5f, 0x54, 0x45, 0x41, 0x4d, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x46, 0x4c, 0x41, 0x47, 0x5f,
	0x53, 0x43, 0x4f, 0x50, 0x45, 0x5f, 0x45, 0x4e, 0x56, 0x49, 0x52, 0x4f, 0x4e, 0x4d, 0x45, 0x4e,
	0x54, 0x10, 0x02, 0x32, 0xed, 0x01, 0x0a, 0x12, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46,
	0x6c, 0x61, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x27, 0x2e,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x06, 0x80, 0x7d, 0x02, 0x90, 0x02, 0x01, 0x12, 0x69, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x26, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46,
	0x6c, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x80, 0x7d, 0x1b,
	0x90, 0x02, 0x02, 0x42, 0xb9, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x46, 0x6c, 0x61,
	0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02, 0x0f, 0x43, 0x68, 0x61, 0x6c,
	0x6b, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x43, 0x68,
	0x61, 0x6c, 0x6b, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b,
	0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x43, 0x68,
	0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_server_v1_flag_proto_rawDescOnce sync.Once
	file_chalk_server_v1_flag_proto_rawDescData = file_chalk_server_v1_flag_proto_rawDesc
)

func file_chalk_server_v1_flag_proto_rawDescGZIP() []byte {
	file_chalk_server_v1_flag_proto_rawDescOnce.Do(func() {
		file_chalk_server_v1_flag_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_server_v1_flag_proto_rawDescData)
	})
	return file_chalk_server_v1_flag_proto_rawDescData
}

var file_chalk_server_v1_flag_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_chalk_server_v1_flag_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_chalk_server_v1_flag_proto_goTypes = []any{
	(FlagScope)(0),                  // 0: chalk.server.v1.FlagScope
	(*FeatureFlagValue)(nil),        // 1: chalk.server.v1.FeatureFlagValue
	(*GetFeatureFlagsRequest)(nil),  // 2: chalk.server.v1.GetFeatureFlagsRequest
	(*GetFeatureFlagsResponse)(nil), // 3: chalk.server.v1.GetFeatureFlagsResponse
	(*SetFeatureFlagRequest)(nil),   // 4: chalk.server.v1.SetFeatureFlagRequest
	(*SetFeatureFlagResponse)(nil),  // 5: chalk.server.v1.SetFeatureFlagResponse
}
var file_chalk_server_v1_flag_proto_depIdxs = []int32{
	1, // 0: chalk.server.v1.GetFeatureFlagsResponse.flags:type_name -> chalk.server.v1.FeatureFlagValue
	0, // 1: chalk.server.v1.SetFeatureFlagRequest.scope:type_name -> chalk.server.v1.FlagScope
	2, // 2: chalk.server.v1.FeatureFlagService.GetFeatureFlags:input_type -> chalk.server.v1.GetFeatureFlagsRequest
	4, // 3: chalk.server.v1.FeatureFlagService.SetFeatureFlag:input_type -> chalk.server.v1.SetFeatureFlagRequest
	3, // 4: chalk.server.v1.FeatureFlagService.GetFeatureFlags:output_type -> chalk.server.v1.GetFeatureFlagsResponse
	5, // 5: chalk.server.v1.FeatureFlagService.SetFeatureFlag:output_type -> chalk.server.v1.SetFeatureFlagResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_chalk_server_v1_flag_proto_init() }
func file_chalk_server_v1_flag_proto_init() {
	if File_chalk_server_v1_flag_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_server_v1_flag_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chalk_server_v1_flag_proto_goTypes,
		DependencyIndexes: file_chalk_server_v1_flag_proto_depIdxs,
		EnumInfos:         file_chalk_server_v1_flag_proto_enumTypes,
		MessageInfos:      file_chalk_server_v1_flag_proto_msgTypes,
	}.Build()
	File_chalk_server_v1_flag_proto = out.File
	file_chalk_server_v1_flag_proto_rawDesc = nil
	file_chalk_server_v1_flag_proto_goTypes = nil
	file_chalk_server_v1_flag_proto_depIdxs = nil
}
