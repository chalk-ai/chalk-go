// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc        (unknown)
// source: chalk/server/v1/cli.proto

package serverv1

import (
	_ "github.com/chalk-ai/chalk-private/go-api-server/gen/chalk/auth/v1"
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

type CommandLineInterfaceVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	DownloadUrl string `protobuf:"bytes,2,opt,name=download_url,json=downloadUrl,proto3" json:"download_url,omitempty"`
	Os          string `protobuf:"bytes,3,opt,name=os,proto3" json:"os,omitempty"`
	Arch        string `protobuf:"bytes,4,opt,name=arch,proto3" json:"arch,omitempty"`
	Generation  int64  `protobuf:"varint,5,opt,name=generation,proto3" json:"generation,omitempty"`
}

func (x *CommandLineInterfaceVersion) Reset() {
	*x = CommandLineInterfaceVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_cli_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandLineInterfaceVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandLineInterfaceVersion) ProtoMessage() {}

func (x *CommandLineInterfaceVersion) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_cli_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandLineInterfaceVersion.ProtoReflect.Descriptor instead.
func (*CommandLineInterfaceVersion) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_cli_proto_rawDescGZIP(), []int{0}
}

func (x *CommandLineInterfaceVersion) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *CommandLineInterfaceVersion) GetDownloadUrl() string {
	if x != nil {
		return x.DownloadUrl
	}
	return ""
}

func (x *CommandLineInterfaceVersion) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *CommandLineInterfaceVersion) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

func (x *CommandLineInterfaceVersion) GetGeneration() int64 {
	if x != nil {
		return x.Generation
	}
	return 0
}

type GetVersionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Os   *string `protobuf:"bytes,1,opt,name=os,proto3,oneof" json:"os,omitempty"`
	Arch *string `protobuf:"bytes,2,opt,name=arch,proto3,oneof" json:"arch,omitempty"`
}

func (x *GetVersionsRequest) Reset() {
	*x = GetVersionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_cli_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVersionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVersionsRequest) ProtoMessage() {}

func (x *GetVersionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_cli_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVersionsRequest.ProtoReflect.Descriptor instead.
func (*GetVersionsRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_cli_proto_rawDescGZIP(), []int{1}
}

func (x *GetVersionsRequest) GetOs() string {
	if x != nil && x.Os != nil {
		return *x.Os
	}
	return ""
}

func (x *GetVersionsRequest) GetArch() string {
	if x != nil && x.Arch != nil {
		return *x.Arch
	}
	return ""
}

type GetVersionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Versions []*CommandLineInterfaceVersion `protobuf:"bytes,1,rep,name=versions,proto3" json:"versions,omitempty"`
	Latest   *CommandLineInterfaceVersion   `protobuf:"bytes,2,opt,name=latest,proto3" json:"latest,omitempty"`
	Nightly  *CommandLineInterfaceVersion   `protobuf:"bytes,3,opt,name=nightly,proto3" json:"nightly,omitempty"`
	Minimum  string                         `protobuf:"bytes,4,opt,name=minimum,proto3" json:"minimum,omitempty"`
}

func (x *GetVersionsResponse) Reset() {
	*x = GetVersionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_cli_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVersionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVersionsResponse) ProtoMessage() {}

func (x *GetVersionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_cli_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVersionsResponse.ProtoReflect.Descriptor instead.
func (*GetVersionsResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_cli_proto_rawDescGZIP(), []int{2}
}

func (x *GetVersionsResponse) GetVersions() []*CommandLineInterfaceVersion {
	if x != nil {
		return x.Versions
	}
	return nil
}

func (x *GetVersionsResponse) GetLatest() *CommandLineInterfaceVersion {
	if x != nil {
		return x.Latest
	}
	return nil
}

func (x *GetVersionsResponse) GetNightly() *CommandLineInterfaceVersion {
	if x != nil {
		return x.Nightly
	}
	return nil
}

func (x *GetVersionsResponse) GetMinimum() string {
	if x != nil {
		return x.Minimum
	}
	return ""
}

var File_chalk_server_v1_cli_proto protoreflect.FileDescriptor

var file_chalk_server_v1_cli_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6c, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01,
	0x0a, 0x1b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72,
	0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1e,
	0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x52,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x02, 0x6f, 0x73, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x61, 0x72, 0x63,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x61, 0x72, 0x63, 0x68, 0x88,
	0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x6f, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x61, 0x72,
	0x63, 0x68, 0x22, 0x87, 0x02, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x08, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x44, 0x0a, 0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69,
	0x6e, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x07, 0x6e, 0x69,
	0x67, 0x68, 0x74, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6e, 0x69, 0x67, 0x68, 0x74,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x32, 0x7f, 0x0a, 0x1b,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x24, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x80, 0x7d, 0x01, 0x90, 0x02, 0x01, 0x42, 0xcb, 0x01,
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x43, 0x6c, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72,
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
	file_chalk_server_v1_cli_proto_rawDescOnce sync.Once
	file_chalk_server_v1_cli_proto_rawDescData = file_chalk_server_v1_cli_proto_rawDesc
)

func file_chalk_server_v1_cli_proto_rawDescGZIP() []byte {
	file_chalk_server_v1_cli_proto_rawDescOnce.Do(func() {
		file_chalk_server_v1_cli_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_server_v1_cli_proto_rawDescData)
	})
	return file_chalk_server_v1_cli_proto_rawDescData
}

var file_chalk_server_v1_cli_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chalk_server_v1_cli_proto_goTypes = []interface{}{
	(*CommandLineInterfaceVersion)(nil), // 0: chalk.server.v1.CommandLineInterfaceVersion
	(*GetVersionsRequest)(nil),          // 1: chalk.server.v1.GetVersionsRequest
	(*GetVersionsResponse)(nil),         // 2: chalk.server.v1.GetVersionsResponse
}
var file_chalk_server_v1_cli_proto_depIdxs = []int32{
	0, // 0: chalk.server.v1.GetVersionsResponse.versions:type_name -> chalk.server.v1.CommandLineInterfaceVersion
	0, // 1: chalk.server.v1.GetVersionsResponse.latest:type_name -> chalk.server.v1.CommandLineInterfaceVersion
	0, // 2: chalk.server.v1.GetVersionsResponse.nightly:type_name -> chalk.server.v1.CommandLineInterfaceVersion
	1, // 3: chalk.server.v1.CommandLineInterfaceService.GetVersions:input_type -> chalk.server.v1.GetVersionsRequest
	2, // 4: chalk.server.v1.CommandLineInterfaceService.GetVersions:output_type -> chalk.server.v1.GetVersionsResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_chalk_server_v1_cli_proto_init() }
func file_chalk_server_v1_cli_proto_init() {
	if File_chalk_server_v1_cli_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chalk_server_v1_cli_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandLineInterfaceVersion); i {
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
		file_chalk_server_v1_cli_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVersionsRequest); i {
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
		file_chalk_server_v1_cli_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVersionsResponse); i {
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
	file_chalk_server_v1_cli_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_server_v1_cli_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chalk_server_v1_cli_proto_goTypes,
		DependencyIndexes: file_chalk_server_v1_cli_proto_depIdxs,
		MessageInfos:      file_chalk_server_v1_cli_proto_msgTypes,
	}.Build()
	File_chalk_server_v1_cli_proto = out.File
	file_chalk_server_v1_cli_proto_rawDesc = nil
	file_chalk_server_v1_cli_proto_goTypes = nil
	file_chalk_server_v1_cli_proto_depIdxs = nil
}
