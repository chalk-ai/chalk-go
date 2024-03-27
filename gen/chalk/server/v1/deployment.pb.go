// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: chalk/server/v1/deployment.proto

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

type DeploymentStatus int32

const (
	DeploymentStatus_DEPLOYMENT_STATUS_UNSPECIFIED    DeploymentStatus = 0
	DeploymentStatus_DEPLOYMENT_STATUS_UNKNOWN        DeploymentStatus = 1
	DeploymentStatus_DEPLOYMENT_STATUS_PENDING        DeploymentStatus = 2
	DeploymentStatus_DEPLOYMENT_STATUS_QUEUED         DeploymentStatus = 3
	DeploymentStatus_DEPLOYMENT_STATUS_WORKING        DeploymentStatus = 4
	DeploymentStatus_DEPLOYMENT_STATUS_SUCCESS        DeploymentStatus = 5
	DeploymentStatus_DEPLOYMENT_STATUS_FAILURE        DeploymentStatus = 6
	DeploymentStatus_DEPLOYMENT_STATUS_INTERNAL_ERROR DeploymentStatus = 7
	DeploymentStatus_DEPLOYMENT_STATUS_TIMEOUT        DeploymentStatus = 8
	DeploymentStatus_DEPLOYMENT_STATUS_CANCELLED      DeploymentStatus = 9
	DeploymentStatus_DEPLOYMENT_STATUS_EXPIRED        DeploymentStatus = 10
	DeploymentStatus_DEPLOYMENT_STATUS_BOOT_ERRORS    DeploymentStatus = 11
)

// Enum value maps for DeploymentStatus.
var (
	DeploymentStatus_name = map[int32]string{
		0:  "DEPLOYMENT_STATUS_UNSPECIFIED",
		1:  "DEPLOYMENT_STATUS_UNKNOWN",
		2:  "DEPLOYMENT_STATUS_PENDING",
		3:  "DEPLOYMENT_STATUS_QUEUED",
		4:  "DEPLOYMENT_STATUS_WORKING",
		5:  "DEPLOYMENT_STATUS_SUCCESS",
		6:  "DEPLOYMENT_STATUS_FAILURE",
		7:  "DEPLOYMENT_STATUS_INTERNAL_ERROR",
		8:  "DEPLOYMENT_STATUS_TIMEOUT",
		9:  "DEPLOYMENT_STATUS_CANCELLED",
		10: "DEPLOYMENT_STATUS_EXPIRED",
		11: "DEPLOYMENT_STATUS_BOOT_ERRORS",
	}
	DeploymentStatus_value = map[string]int32{
		"DEPLOYMENT_STATUS_UNSPECIFIED":    0,
		"DEPLOYMENT_STATUS_UNKNOWN":        1,
		"DEPLOYMENT_STATUS_PENDING":        2,
		"DEPLOYMENT_STATUS_QUEUED":         3,
		"DEPLOYMENT_STATUS_WORKING":        4,
		"DEPLOYMENT_STATUS_SUCCESS":        5,
		"DEPLOYMENT_STATUS_FAILURE":        6,
		"DEPLOYMENT_STATUS_INTERNAL_ERROR": 7,
		"DEPLOYMENT_STATUS_TIMEOUT":        8,
		"DEPLOYMENT_STATUS_CANCELLED":      9,
		"DEPLOYMENT_STATUS_EXPIRED":        10,
		"DEPLOYMENT_STATUS_BOOT_ERRORS":    11,
	}
)

func (x DeploymentStatus) Enum() *DeploymentStatus {
	p := new(DeploymentStatus)
	*p = x
	return p
}

func (x DeploymentStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeploymentStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_chalk_server_v1_deployment_proto_enumTypes[0].Descriptor()
}

func (DeploymentStatus) Type() protoreflect.EnumType {
	return &file_chalk_server_v1_deployment_proto_enumTypes[0]
}

func (x DeploymentStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeploymentStatus.Descriptor instead.
func (DeploymentStatus) EnumDescriptor() ([]byte, []int) {
	return file_chalk_server_v1_deployment_proto_rawDescGZIP(), []int{0}
}

type Deployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EnvironmentId string           `protobuf:"bytes,2,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	Status        DeploymentStatus `protobuf:"varint,3,opt,name=status,proto3,enum=chalk.server.v1.DeploymentStatus" json:"status,omitempty"`
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_deployment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_deployment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_deployment_proto_rawDescGZIP(), []int{0}
}

func (x *Deployment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Deployment) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *Deployment) GetStatus() DeploymentStatus {
	if x != nil {
		return x.Status
	}
	return DeploymentStatus_DEPLOYMENT_STATUS_UNSPECIFIED
}

var File_chalk_server_v1_deployment_proto protoreflect.FileDescriptor

var file_chalk_server_v1_deployment_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x22, 0x7e, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2a, 0x96, 0x03, 0x0a, 0x10, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x1d, 0x44, 0x45, 0x50, 0x4c,
	0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x44,
	0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x45,
	0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x45, 0x50,
	0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x51,
	0x55, 0x45, 0x55, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x45, 0x50, 0x4c, 0x4f,
	0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x57, 0x4f, 0x52,
	0x4b, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43,
	0x45, 0x53, 0x53, 0x10, 0x05, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55,
	0x52, 0x45, 0x10, 0x06, 0x12, 0x24, 0x0a, 0x20, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e,
	0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x07, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x45,
	0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10, 0x08, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x45, 0x50,
	0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43,
	0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x09, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x45,
	0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x21, 0x0a, 0x1d, 0x44, 0x45, 0x50,
	0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x42,
	0x4f, 0x4f, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x53, 0x10, 0x0b, 0x42, 0xbf, 0x01, 0x0a,
	0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x43, 0x68, 0x61,
	0x6c, 0x6b, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x43,
	0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x43, 0x68, 0x61,
	0x6c, 0x6b, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_server_v1_deployment_proto_rawDescOnce sync.Once
	file_chalk_server_v1_deployment_proto_rawDescData = file_chalk_server_v1_deployment_proto_rawDesc
)

func file_chalk_server_v1_deployment_proto_rawDescGZIP() []byte {
	file_chalk_server_v1_deployment_proto_rawDescOnce.Do(func() {
		file_chalk_server_v1_deployment_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_server_v1_deployment_proto_rawDescData)
	})
	return file_chalk_server_v1_deployment_proto_rawDescData
}

var file_chalk_server_v1_deployment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_chalk_server_v1_deployment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chalk_server_v1_deployment_proto_goTypes = []interface{}{
	(DeploymentStatus)(0), // 0: chalk.server.v1.DeploymentStatus
	(*Deployment)(nil),    // 1: chalk.server.v1.Deployment
}
var file_chalk_server_v1_deployment_proto_depIdxs = []int32{
	0, // 0: chalk.server.v1.Deployment.status:type_name -> chalk.server.v1.DeploymentStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chalk_server_v1_deployment_proto_init() }
func file_chalk_server_v1_deployment_proto_init() {
	if File_chalk_server_v1_deployment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chalk_server_v1_deployment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployment); i {
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
			RawDescriptor: file_chalk_server_v1_deployment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chalk_server_v1_deployment_proto_goTypes,
		DependencyIndexes: file_chalk_server_v1_deployment_proto_depIdxs,
		EnumInfos:         file_chalk_server_v1_deployment_proto_enumTypes,
		MessageInfos:      file_chalk_server_v1_deployment_proto_msgTypes,
	}.Build()
	File_chalk_server_v1_deployment_proto = out.File
	file_chalk_server_v1_deployment_proto_rawDesc = nil
	file_chalk_server_v1_deployment_proto_goTypes = nil
	file_chalk_server_v1_deployment_proto_depIdxs = nil
}
