// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: chalk/server/v1/bootstrap.proto

package serverv1

import (
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

type BootstrapExtraSettingsEnvironment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Settings map[string]bool `protobuf:"bytes,1,rep,name=settings,proto3" json:"settings,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *BootstrapExtraSettingsEnvironment) Reset() {
	*x = BootstrapExtraSettingsEnvironment{}
	mi := &file_chalk_server_v1_bootstrap_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BootstrapExtraSettingsEnvironment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrapExtraSettingsEnvironment) ProtoMessage() {}

func (x *BootstrapExtraSettingsEnvironment) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_bootstrap_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrapExtraSettingsEnvironment.ProtoReflect.Descriptor instead.
func (*BootstrapExtraSettingsEnvironment) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_bootstrap_proto_rawDescGZIP(), []int{0}
}

func (x *BootstrapExtraSettingsEnvironment) GetSettings() map[string]bool {
	if x != nil {
		return x.Settings
	}
	return nil
}

type BootstrapExtraSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Global       map[string]bool                               `protobuf:"bytes,1,rep,name=global,proto3" json:"global,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Environments map[string]*BootstrapExtraSettingsEnvironment `protobuf:"bytes,2,rep,name=environments,proto3" json:"environments,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BootstrapExtraSettings) Reset() {
	*x = BootstrapExtraSettings{}
	mi := &file_chalk_server_v1_bootstrap_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BootstrapExtraSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrapExtraSettings) ProtoMessage() {}

func (x *BootstrapExtraSettings) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_bootstrap_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrapExtraSettings.ProtoReflect.Descriptor instead.
func (*BootstrapExtraSettings) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_bootstrap_proto_rawDescGZIP(), []int{1}
}

func (x *BootstrapExtraSettings) GetGlobal() map[string]bool {
	if x != nil {
		return x.Global
	}
	return nil
}

func (x *BootstrapExtraSettings) GetEnvironments() map[string]*BootstrapExtraSettingsEnvironment {
	if x != nil {
		return x.Environments
	}
	return nil
}

type ParsedBootstrapConfigs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Teams                 []*Team                 `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
	Projects              []*Project              `protobuf:"bytes,2,rep,name=projects,proto3" json:"projects,omitempty"`
	Environments          []*Environment          `protobuf:"bytes,3,rep,name=environments,proto3" json:"environments,omitempty"`
	TeamInvites           []*TeamInvite           `protobuf:"bytes,4,rep,name=team_invites,json=teamInvites,proto3" json:"team_invites,omitempty"`
	ExtraSettings         *BootstrapExtraSettings `protobuf:"bytes,5,opt,name=extra_settings,json=extraSettings,proto3" json:"extra_settings,omitempty"`
	GlobalPinnedBaseImage *string                 `protobuf:"bytes,6,opt,name=global_pinned_base_image,json=globalPinnedBaseImage,proto3,oneof" json:"global_pinned_base_image,omitempty"`
}

func (x *ParsedBootstrapConfigs) Reset() {
	*x = ParsedBootstrapConfigs{}
	mi := &file_chalk_server_v1_bootstrap_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParsedBootstrapConfigs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParsedBootstrapConfigs) ProtoMessage() {}

func (x *ParsedBootstrapConfigs) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_bootstrap_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParsedBootstrapConfigs.ProtoReflect.Descriptor instead.
func (*ParsedBootstrapConfigs) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_bootstrap_proto_rawDescGZIP(), []int{2}
}

func (x *ParsedBootstrapConfigs) GetTeams() []*Team {
	if x != nil {
		return x.Teams
	}
	return nil
}

func (x *ParsedBootstrapConfigs) GetProjects() []*Project {
	if x != nil {
		return x.Projects
	}
	return nil
}

func (x *ParsedBootstrapConfigs) GetEnvironments() []*Environment {
	if x != nil {
		return x.Environments
	}
	return nil
}

func (x *ParsedBootstrapConfigs) GetTeamInvites() []*TeamInvite {
	if x != nil {
		return x.TeamInvites
	}
	return nil
}

func (x *ParsedBootstrapConfigs) GetExtraSettings() *BootstrapExtraSettings {
	if x != nil {
		return x.ExtraSettings
	}
	return nil
}

func (x *ParsedBootstrapConfigs) GetGlobalPinnedBaseImage() string {
	if x != nil && x.GlobalPinnedBaseImage != nil {
		return *x.GlobalPinnedBaseImage
	}
	return ""
}

var File_chalk_server_v1_bootstrap_proto protoreflect.FileDescriptor

var file_chalk_server_v1_bootstrap_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x21, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x21, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x5c, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x74,
	0x73, 0x74, 0x72, 0x61, 0x70, 0x45, 0x78, 0x74, 0x72, 0x61, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x1a, 0x3b, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xf4, 0x02, 0x0a, 0x16, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x4b, 0x0a,
	0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x45, 0x78, 0x74, 0x72, 0x61, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x12, 0x5d, 0x0a, 0x0c, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x39, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x47, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x73, 0x0a, 0x11, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x48, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f,
	0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x45, 0x78, 0x74, 0x72, 0x61, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa8, 0x03, 0x0a, 0x16, 0x50, 0x61,
	0x72, 0x73, 0x65, 0x64, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x12, 0x2b, 0x0a, 0x05, 0x74, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x05, 0x74, 0x65, 0x61, 0x6d,
	0x73, 0x12, 0x34, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x40, 0x0a, 0x0c, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0c, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3e, 0x0a, 0x0c, 0x74, 0x65, 0x61,
	0x6d, 0x5f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x65, 0x61, 0x6d, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x52, 0x0b, 0x74, 0x65,
	0x61, 0x6d, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x0e, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3c, 0x0a, 0x18, 0x67, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x5f, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x15, 0x67,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x50, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x42, 0x61, 0x73, 0x65, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x1b, 0x0a, 0x19, 0x5f, 0x67, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x5f, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x42, 0xbe, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x42, 0x6f,
	0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39,
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
	file_chalk_server_v1_bootstrap_proto_rawDescOnce sync.Once
	file_chalk_server_v1_bootstrap_proto_rawDescData = file_chalk_server_v1_bootstrap_proto_rawDesc
)

func file_chalk_server_v1_bootstrap_proto_rawDescGZIP() []byte {
	file_chalk_server_v1_bootstrap_proto_rawDescOnce.Do(func() {
		file_chalk_server_v1_bootstrap_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_server_v1_bootstrap_proto_rawDescData)
	})
	return file_chalk_server_v1_bootstrap_proto_rawDescData
}

var file_chalk_server_v1_bootstrap_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_chalk_server_v1_bootstrap_proto_goTypes = []any{
	(*BootstrapExtraSettingsEnvironment)(nil), // 0: chalk.server.v1.BootstrapExtraSettingsEnvironment
	(*BootstrapExtraSettings)(nil),            // 1: chalk.server.v1.BootstrapExtraSettings
	(*ParsedBootstrapConfigs)(nil),            // 2: chalk.server.v1.ParsedBootstrapConfigs
	nil,                                       // 3: chalk.server.v1.BootstrapExtraSettingsEnvironment.SettingsEntry
	nil,                                       // 4: chalk.server.v1.BootstrapExtraSettings.GlobalEntry
	nil,                                       // 5: chalk.server.v1.BootstrapExtraSettings.EnvironmentsEntry
	(*Team)(nil),                              // 6: chalk.server.v1.Team
	(*Project)(nil),                           // 7: chalk.server.v1.Project
	(*Environment)(nil),                       // 8: chalk.server.v1.Environment
	(*TeamInvite)(nil),                        // 9: chalk.server.v1.TeamInvite
}
var file_chalk_server_v1_bootstrap_proto_depIdxs = []int32{
	3, // 0: chalk.server.v1.BootstrapExtraSettingsEnvironment.settings:type_name -> chalk.server.v1.BootstrapExtraSettingsEnvironment.SettingsEntry
	4, // 1: chalk.server.v1.BootstrapExtraSettings.global:type_name -> chalk.server.v1.BootstrapExtraSettings.GlobalEntry
	5, // 2: chalk.server.v1.BootstrapExtraSettings.environments:type_name -> chalk.server.v1.BootstrapExtraSettings.EnvironmentsEntry
	6, // 3: chalk.server.v1.ParsedBootstrapConfigs.teams:type_name -> chalk.server.v1.Team
	7, // 4: chalk.server.v1.ParsedBootstrapConfigs.projects:type_name -> chalk.server.v1.Project
	8, // 5: chalk.server.v1.ParsedBootstrapConfigs.environments:type_name -> chalk.server.v1.Environment
	9, // 6: chalk.server.v1.ParsedBootstrapConfigs.team_invites:type_name -> chalk.server.v1.TeamInvite
	1, // 7: chalk.server.v1.ParsedBootstrapConfigs.extra_settings:type_name -> chalk.server.v1.BootstrapExtraSettings
	0, // 8: chalk.server.v1.BootstrapExtraSettings.EnvironmentsEntry.value:type_name -> chalk.server.v1.BootstrapExtraSettingsEnvironment
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_chalk_server_v1_bootstrap_proto_init() }
func file_chalk_server_v1_bootstrap_proto_init() {
	if File_chalk_server_v1_bootstrap_proto != nil {
		return
	}
	file_chalk_server_v1_environment_proto_init()
	file_chalk_server_v1_team_proto_init()
	file_chalk_server_v1_bootstrap_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_server_v1_bootstrap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chalk_server_v1_bootstrap_proto_goTypes,
		DependencyIndexes: file_chalk_server_v1_bootstrap_proto_depIdxs,
		MessageInfos:      file_chalk_server_v1_bootstrap_proto_msgTypes,
	}.Build()
	File_chalk_server_v1_bootstrap_proto = out.File
	file_chalk_server_v1_bootstrap_proto_rawDesc = nil
	file_chalk_server_v1_bootstrap_proto_goTypes = nil
	file_chalk_server_v1_bootstrap_proto_depIdxs = nil
}
