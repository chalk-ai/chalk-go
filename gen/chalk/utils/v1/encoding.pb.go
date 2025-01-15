// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: chalk/utils/v1/encoding.proto

package utilsv1

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

type StringEncoding struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Mapping       map[int32]string       `protobuf:"bytes,1,rep,name=mapping,proto3" json:"mapping,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StringEncoding) Reset() {
	*x = StringEncoding{}
	mi := &file_chalk_utils_v1_encoding_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StringEncoding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringEncoding) ProtoMessage() {}

func (x *StringEncoding) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_utils_v1_encoding_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringEncoding.ProtoReflect.Descriptor instead.
func (*StringEncoding) Descriptor() ([]byte, []int) {
	return file_chalk_utils_v1_encoding_proto_rawDescGZIP(), []int{0}
}

func (x *StringEncoding) GetMapping() map[int32]string {
	if x != nil {
		return x.Mapping
	}
	return nil
}

var file_chalk_utils_v1_encoding_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*StringEncoding)(nil),
		Field:         80412,
		Name:          "chalk.utils.v1.encoding",
		Tag:           "bytes,80412,opt,name=encoding",
		Filename:      "chalk/utils/v1/encoding.proto",
	},
}

// Extension fields to descriptorpb.EnumOptions.
var (
	// optional chalk.utils.v1.StringEncoding encoding = 80412;
	E_Encoding = &file_chalk_utils_v1_encoding_proto_extTypes[0]
)

var File_chalk_utils_v1_encoding_proto protoreflect.FileDescriptor

var file_chalk_utils_v1_encoding_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x93, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x63, 0x6f,
	0x64, 0x69, 0x6e, 0x67, 0x12, 0x45, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x75, 0x74,
	0x69, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x1a, 0x3a, 0x0a, 0x0c, 0x4d,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x5d, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x9c, 0xf4, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x42, 0xb6, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x45,
	0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x3b,
	0x75, 0x74, 0x69, 0x6c, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x55, 0x58, 0xaa, 0x02, 0x0e,
	0x43, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x55, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x0e, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x55, 0x74, 0x69, 0x6c, 0x73, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1a, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x55, 0x74, 0x69, 0x6c, 0x73, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x43,
	0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x55, 0x74, 0x69, 0x6c, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_utils_v1_encoding_proto_rawDescOnce sync.Once
	file_chalk_utils_v1_encoding_proto_rawDescData = file_chalk_utils_v1_encoding_proto_rawDesc
)

func file_chalk_utils_v1_encoding_proto_rawDescGZIP() []byte {
	file_chalk_utils_v1_encoding_proto_rawDescOnce.Do(func() {
		file_chalk_utils_v1_encoding_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_utils_v1_encoding_proto_rawDescData)
	})
	return file_chalk_utils_v1_encoding_proto_rawDescData
}

var file_chalk_utils_v1_encoding_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chalk_utils_v1_encoding_proto_goTypes = []any{
	(*StringEncoding)(nil),           // 0: chalk.utils.v1.StringEncoding
	nil,                              // 1: chalk.utils.v1.StringEncoding.MappingEntry
	(*descriptorpb.EnumOptions)(nil), // 2: google.protobuf.EnumOptions
}
var file_chalk_utils_v1_encoding_proto_depIdxs = []int32{
	1, // 0: chalk.utils.v1.StringEncoding.mapping:type_name -> chalk.utils.v1.StringEncoding.MappingEntry
	2, // 1: chalk.utils.v1.encoding:extendee -> google.protobuf.EnumOptions
	0, // 2: chalk.utils.v1.encoding:type_name -> chalk.utils.v1.StringEncoding
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	2, // [2:3] is the sub-list for extension type_name
	1, // [1:2] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chalk_utils_v1_encoding_proto_init() }
func file_chalk_utils_v1_encoding_proto_init() {
	if File_chalk_utils_v1_encoding_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_utils_v1_encoding_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_chalk_utils_v1_encoding_proto_goTypes,
		DependencyIndexes: file_chalk_utils_v1_encoding_proto_depIdxs,
		MessageInfos:      file_chalk_utils_v1_encoding_proto_msgTypes,
		ExtensionInfos:    file_chalk_utils_v1_encoding_proto_extTypes,
	}.Build()
	File_chalk_utils_v1_encoding_proto = out.File
	file_chalk_utils_v1_encoding_proto_rawDesc = nil
	file_chalk_utils_v1_encoding_proto_goTypes = nil
	file_chalk_utils_v1_encoding_proto_depIdxs = nil
}
