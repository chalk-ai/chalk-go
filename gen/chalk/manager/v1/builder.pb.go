// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: chalk/manager/v1/builder.proto

package managerv1

import (
	_ "github.com/chalk-ai/chalk-go/v2/gen/chalk/auth/v1"
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

type KafkaTopic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Partitions  int32  `protobuf:"varint,2,opt,name=partitions,proto3" json:"partitions,omitempty"`
	Replication *int32 `protobuf:"varint,3,opt,name=replication,proto3,oneof" json:"replication,omitempty"`
	RetentionMs int32  `protobuf:"varint,4,opt,name=retention_ms,json=retentionMs,proto3" json:"retention_ms,omitempty"`
}

func (x *KafkaTopic) Reset() {
	*x = KafkaTopic{}
	mi := &file_chalk_manager_v1_builder_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KafkaTopic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KafkaTopic) ProtoMessage() {}

func (x *KafkaTopic) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_manager_v1_builder_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KafkaTopic.ProtoReflect.Descriptor instead.
func (*KafkaTopic) Descriptor() ([]byte, []int) {
	return file_chalk_manager_v1_builder_proto_rawDescGZIP(), []int{0}
}

func (x *KafkaTopic) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KafkaTopic) GetPartitions() int32 {
	if x != nil {
		return x.Partitions
	}
	return 0
}

func (x *KafkaTopic) GetReplication() int32 {
	if x != nil && x.Replication != nil {
		return *x.Replication
	}
	return 0
}

func (x *KafkaTopic) GetRetentionMs() int32 {
	if x != nil {
		return x.RetentionMs
	}
	return 0
}

type CreateKafkaTopicsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topics []*KafkaTopic `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *CreateKafkaTopicsRequest) Reset() {
	*x = CreateKafkaTopicsRequest{}
	mi := &file_chalk_manager_v1_builder_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateKafkaTopicsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKafkaTopicsRequest) ProtoMessage() {}

func (x *CreateKafkaTopicsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_manager_v1_builder_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKafkaTopicsRequest.ProtoReflect.Descriptor instead.
func (*CreateKafkaTopicsRequest) Descriptor() ([]byte, []int) {
	return file_chalk_manager_v1_builder_proto_rawDescGZIP(), []int{1}
}

func (x *CreateKafkaTopicsRequest) GetTopics() []*KafkaTopic {
	if x != nil {
		return x.Topics
	}
	return nil
}

type CreateKafkaTopicsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateKafkaTopicsResponse) Reset() {
	*x = CreateKafkaTopicsResponse{}
	mi := &file_chalk_manager_v1_builder_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateKafkaTopicsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKafkaTopicsResponse) ProtoMessage() {}

func (x *CreateKafkaTopicsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_manager_v1_builder_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKafkaTopicsResponse.ProtoReflect.Descriptor instead.
func (*CreateKafkaTopicsResponse) Descriptor() ([]byte, []int) {
	return file_chalk_manager_v1_builder_proto_rawDescGZIP(), []int{2}
}

type GetKafkaTopicsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetKafkaTopicsRequest) Reset() {
	*x = GetKafkaTopicsRequest{}
	mi := &file_chalk_manager_v1_builder_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetKafkaTopicsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKafkaTopicsRequest) ProtoMessage() {}

func (x *GetKafkaTopicsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_manager_v1_builder_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKafkaTopicsRequest.ProtoReflect.Descriptor instead.
func (*GetKafkaTopicsRequest) Descriptor() ([]byte, []int) {
	return file_chalk_manager_v1_builder_proto_rawDescGZIP(), []int{3}
}

type GetKafkaTopicsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topics []*KafkaTopic `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *GetKafkaTopicsResponse) Reset() {
	*x = GetKafkaTopicsResponse{}
	mi := &file_chalk_manager_v1_builder_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetKafkaTopicsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKafkaTopicsResponse) ProtoMessage() {}

func (x *GetKafkaTopicsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_manager_v1_builder_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKafkaTopicsResponse.ProtoReflect.Descriptor instead.
func (*GetKafkaTopicsResponse) Descriptor() ([]byte, []int) {
	return file_chalk_manager_v1_builder_proto_rawDescGZIP(), []int{4}
}

func (x *GetKafkaTopicsResponse) GetTopics() []*KafkaTopic {
	if x != nil {
		return x.Topics
	}
	return nil
}

var File_chalk_manager_v1_builder_proto protoreflect.FileDescriptor

var file_chalk_manager_v1_builder_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x0a, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0b, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x50, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x06,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4b, 0x61, 0x66, 0x6b, 0x61, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x73, 0x22, 0x1b, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b,
	0x61, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4b,
	0x61, 0x66, 0x6b, 0x61, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x32, 0xf4, 0x01, 0x0a, 0x15, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x71, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x61, 0x66, 0x6b,
	0x61, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x2a, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x61, 0x66,
	0x6b, 0x61, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x03, 0x80, 0x7d, 0x0a, 0x12, 0x68, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b,
	0x61, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x27, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x61,
	0x66, 0x6b, 0x61, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x80, 0x7d, 0x0a, 0x42,
	0xc3, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x4d, 0x58, 0xaa, 0x02, 0x10, 0x43, 0x68,
	0x61, 0x6c, 0x6b, 0x2e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x10, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x1c, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x12, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_manager_v1_builder_proto_rawDescOnce sync.Once
	file_chalk_manager_v1_builder_proto_rawDescData = file_chalk_manager_v1_builder_proto_rawDesc
)

func file_chalk_manager_v1_builder_proto_rawDescGZIP() []byte {
	file_chalk_manager_v1_builder_proto_rawDescOnce.Do(func() {
		file_chalk_manager_v1_builder_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_manager_v1_builder_proto_rawDescData)
	})
	return file_chalk_manager_v1_builder_proto_rawDescData
}

var file_chalk_manager_v1_builder_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_chalk_manager_v1_builder_proto_goTypes = []any{
	(*KafkaTopic)(nil),                // 0: chalk.manager.v1.KafkaTopic
	(*CreateKafkaTopicsRequest)(nil),  // 1: chalk.manager.v1.CreateKafkaTopicsRequest
	(*CreateKafkaTopicsResponse)(nil), // 2: chalk.manager.v1.CreateKafkaTopicsResponse
	(*GetKafkaTopicsRequest)(nil),     // 3: chalk.manager.v1.GetKafkaTopicsRequest
	(*GetKafkaTopicsResponse)(nil),    // 4: chalk.manager.v1.GetKafkaTopicsResponse
}
var file_chalk_manager_v1_builder_proto_depIdxs = []int32{
	0, // 0: chalk.manager.v1.CreateKafkaTopicsRequest.topics:type_name -> chalk.manager.v1.KafkaTopic
	0, // 1: chalk.manager.v1.GetKafkaTopicsResponse.topics:type_name -> chalk.manager.v1.KafkaTopic
	1, // 2: chalk.manager.v1.ClusterBuilderService.CreateKafkaTopics:input_type -> chalk.manager.v1.CreateKafkaTopicsRequest
	3, // 3: chalk.manager.v1.ClusterBuilderService.GetKafkaTopics:input_type -> chalk.manager.v1.GetKafkaTopicsRequest
	2, // 4: chalk.manager.v1.ClusterBuilderService.CreateKafkaTopics:output_type -> chalk.manager.v1.CreateKafkaTopicsResponse
	4, // 5: chalk.manager.v1.ClusterBuilderService.GetKafkaTopics:output_type -> chalk.manager.v1.GetKafkaTopicsResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_chalk_manager_v1_builder_proto_init() }
func file_chalk_manager_v1_builder_proto_init() {
	if File_chalk_manager_v1_builder_proto != nil {
		return
	}
	file_chalk_manager_v1_builder_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_manager_v1_builder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chalk_manager_v1_builder_proto_goTypes,
		DependencyIndexes: file_chalk_manager_v1_builder_proto_depIdxs,
		MessageInfos:      file_chalk_manager_v1_builder_proto_msgTypes,
	}.Build()
	File_chalk_manager_v1_builder_proto = out.File
	file_chalk_manager_v1_builder_proto_rawDesc = nil
	file_chalk_manager_v1_builder_proto_goTypes = nil
	file_chalk_manager_v1_builder_proto_depIdxs = nil
}
