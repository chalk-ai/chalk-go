// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: chalk/auth/v1/audit.proto

package authv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuditLevel int32

const (
	AuditLevel_AUDIT_LEVEL_UNSPECIFIED AuditLevel = 0
	AuditLevel_AUDIT_LEVEL_ERRORS      AuditLevel = 1
	AuditLevel_AUDIT_LEVEL_ALL         AuditLevel = 2
)

// Enum value maps for AuditLevel.
var (
	AuditLevel_name = map[int32]string{
		0: "AUDIT_LEVEL_UNSPECIFIED",
		1: "AUDIT_LEVEL_ERRORS",
		2: "AUDIT_LEVEL_ALL",
	}
	AuditLevel_value = map[string]int32{
		"AUDIT_LEVEL_UNSPECIFIED": 0,
		"AUDIT_LEVEL_ERRORS":      1,
		"AUDIT_LEVEL_ALL":         2,
	}
)

func (x AuditLevel) Enum() *AuditLevel {
	p := new(AuditLevel)
	*p = x
	return p
}

func (x AuditLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuditLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_chalk_auth_v1_audit_proto_enumTypes[0].Descriptor()
}

func (AuditLevel) Type() protoreflect.EnumType {
	return &file_chalk_auth_v1_audit_proto_enumTypes[0]
}

func (x AuditLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuditLevel.Descriptor instead.
func (AuditLevel) EnumDescriptor() ([]byte, []int) {
	return file_chalk_auth_v1_audit_proto_rawDescGZIP(), []int{0}
}

type AuditOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level       AuditLevel `protobuf:"varint,1,opt,name=level,proto3,enum=chalk.auth.v1.AuditLevel" json:"level,omitempty"`
	Description string     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *AuditOptions) Reset() {
	*x = AuditOptions{}
	mi := &file_chalk_auth_v1_audit_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuditOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditOptions) ProtoMessage() {}

func (x *AuditOptions) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_auth_v1_audit_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditOptions.ProtoReflect.Descriptor instead.
func (*AuditOptions) Descriptor() ([]byte, []int) {
	return file_chalk_auth_v1_audit_proto_rawDescGZIP(), []int{0}
}

func (x *AuditOptions) GetLevel() AuditLevel {
	if x != nil {
		return x.Level
	}
	return AuditLevel_AUDIT_LEVEL_UNSPECIFIED
}

func (x *AuditOptions) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var file_chalk_auth_v1_audit_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*AuditOptions)(nil),
		Field:         30001,
		Name:          "chalk.auth.v1.audit",
		Tag:           "bytes,30001,opt,name=audit",
		Filename:      "chalk/auth/v1/audit.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional chalk.auth.v1.AuditOptions audit = 30001;
	E_Audit = &file_chalk_auth_v1_audit_proto_extTypes[0]
)

var File_chalk_auth_v1_audit_proto protoreflect.FileDescriptor

var file_chalk_auth_v1_audit_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x0c,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2f, 0x0a, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a,
	0x56, 0x0a, 0x0a, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a,
	0x17, 0x41, 0x55, 0x44, 0x49, 0x54, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x55,
	0x44, 0x49, 0x54, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x53,
	0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x55, 0x44, 0x49, 0x54, 0x5f, 0x4c, 0x45, 0x56, 0x45,
	0x4c, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x02, 0x3a, 0x56, 0x0a, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xb1, 0xea, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x61, 0x75, 0x64, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42,
	0xac, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x41, 0x75, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x67,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x41, 0x58,
	0xaa, 0x02, 0x0d, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0d, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x19, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x43,
	0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_auth_v1_audit_proto_rawDescOnce sync.Once
	file_chalk_auth_v1_audit_proto_rawDescData = file_chalk_auth_v1_audit_proto_rawDesc
)

func file_chalk_auth_v1_audit_proto_rawDescGZIP() []byte {
	file_chalk_auth_v1_audit_proto_rawDescOnce.Do(func() {
		file_chalk_auth_v1_audit_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_auth_v1_audit_proto_rawDescData)
	})
	return file_chalk_auth_v1_audit_proto_rawDescData
}

var file_chalk_auth_v1_audit_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_chalk_auth_v1_audit_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chalk_auth_v1_audit_proto_goTypes = []any{
	(AuditLevel)(0),                    // 0: chalk.auth.v1.AuditLevel
	(*AuditOptions)(nil),               // 1: chalk.auth.v1.AuditOptions
	(*descriptorpb.MethodOptions)(nil), // 2: google.protobuf.MethodOptions
}
var file_chalk_auth_v1_audit_proto_depIdxs = []int32{
	0, // 0: chalk.auth.v1.AuditOptions.level:type_name -> chalk.auth.v1.AuditLevel
	2, // 1: chalk.auth.v1.audit:extendee -> google.protobuf.MethodOptions
	1, // 2: chalk.auth.v1.audit:type_name -> chalk.auth.v1.AuditOptions
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	1, // [1:2] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chalk_auth_v1_audit_proto_init() }
func file_chalk_auth_v1_audit_proto_init() {
	if File_chalk_auth_v1_audit_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_auth_v1_audit_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_chalk_auth_v1_audit_proto_goTypes,
		DependencyIndexes: file_chalk_auth_v1_audit_proto_depIdxs,
		EnumInfos:         file_chalk_auth_v1_audit_proto_enumTypes,
		MessageInfos:      file_chalk_auth_v1_audit_proto_msgTypes,
		ExtensionInfos:    file_chalk_auth_v1_audit_proto_extTypes,
	}.Build()
	File_chalk_auth_v1_audit_proto = out.File
	file_chalk_auth_v1_audit_proto_rawDesc = nil
	file_chalk_auth_v1_audit_proto_goTypes = nil
	file_chalk_auth_v1_audit_proto_depIdxs = nil
}
