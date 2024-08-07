// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: chalk/server/v1/named_query.proto

package serverv1

import (
	_ "github.com/chalk-ai/chalk-go/gen/chalk/auth/v1"
	v1 "github.com/chalk-ai/chalk-go/gen/chalk/graph/v1"
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

type GetAllNamedQueriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeploymentId string `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
}

func (x *GetAllNamedQueriesRequest) Reset() {
	*x = GetAllNamedQueriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_named_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllNamedQueriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllNamedQueriesRequest) ProtoMessage() {}

func (x *GetAllNamedQueriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_named_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllNamedQueriesRequest.ProtoReflect.Descriptor instead.
func (*GetAllNamedQueriesRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_named_query_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllNamedQueriesRequest) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

type GetAllNamedQueriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NamedQueries []*v1.NamedQuery `protobuf:"bytes,1,rep,name=named_queries,json=namedQueries,proto3" json:"named_queries,omitempty"`
}

func (x *GetAllNamedQueriesResponse) Reset() {
	*x = GetAllNamedQueriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_named_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllNamedQueriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllNamedQueriesResponse) ProtoMessage() {}

func (x *GetAllNamedQueriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_named_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllNamedQueriesResponse.ProtoReflect.Descriptor instead.
func (*GetAllNamedQueriesResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_named_query_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllNamedQueriesResponse) GetNamedQueries() []*v1.NamedQuery {
	if x != nil {
		return x.NamedQueries
	}
	return nil
}

type GetAllNamedQueriesActiveDeploymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllNamedQueriesActiveDeploymentRequest) Reset() {
	*x = GetAllNamedQueriesActiveDeploymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_named_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllNamedQueriesActiveDeploymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllNamedQueriesActiveDeploymentRequest) ProtoMessage() {}

func (x *GetAllNamedQueriesActiveDeploymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_named_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllNamedQueriesActiveDeploymentRequest.ProtoReflect.Descriptor instead.
func (*GetAllNamedQueriesActiveDeploymentRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_named_query_proto_rawDescGZIP(), []int{2}
}

type GetAllNamedQueriesActiveDeploymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NamedQueries []*v1.NamedQuery `protobuf:"bytes,1,rep,name=named_queries,json=namedQueries,proto3" json:"named_queries,omitempty"`
}

func (x *GetAllNamedQueriesActiveDeploymentResponse) Reset() {
	*x = GetAllNamedQueriesActiveDeploymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_named_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllNamedQueriesActiveDeploymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllNamedQueriesActiveDeploymentResponse) ProtoMessage() {}

func (x *GetAllNamedQueriesActiveDeploymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_named_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllNamedQueriesActiveDeploymentResponse.ProtoReflect.Descriptor instead.
func (*GetAllNamedQueriesActiveDeploymentResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_named_query_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllNamedQueriesActiveDeploymentResponse) GetNamedQueries() []*v1.NamedQuery {
	if x != nil {
		return x.NamedQueries
	}
	return nil
}

var File_chalk_server_v1_named_query_proto protoreflect.FileDescriptor

var file_chalk_server_v1_named_query_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x40, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x64,
	0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x61, 0x6d,
	0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x22, 0x2b, 0x0a, 0x29, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x64, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x6d, 0x0a, 0x2a, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x51, 0x75,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a,
	0x0d, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x52, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x32, 0xb2,
	0x02, 0x0a, 0x11, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x75, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2a, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x06, 0x80, 0x7d, 0x06, 0x90, 0x02, 0x01, 0x12, 0xa5, 0x01, 0x0a, 0x22,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x3a, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x64,
	0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3b,
	0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x51, 0x75, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x80, 0x7d, 0x06,
	0x90, 0x02, 0x01, 0x42, 0xbf, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x4e, 0x61, 0x6d,
	0x65, 0x64, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39,
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
	file_chalk_server_v1_named_query_proto_rawDescOnce sync.Once
	file_chalk_server_v1_named_query_proto_rawDescData = file_chalk_server_v1_named_query_proto_rawDesc
)

func file_chalk_server_v1_named_query_proto_rawDescGZIP() []byte {
	file_chalk_server_v1_named_query_proto_rawDescOnce.Do(func() {
		file_chalk_server_v1_named_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_server_v1_named_query_proto_rawDescData)
	})
	return file_chalk_server_v1_named_query_proto_rawDescData
}

var file_chalk_server_v1_named_query_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chalk_server_v1_named_query_proto_goTypes = []any{
	(*GetAllNamedQueriesRequest)(nil),                  // 0: chalk.server.v1.GetAllNamedQueriesRequest
	(*GetAllNamedQueriesResponse)(nil),                 // 1: chalk.server.v1.GetAllNamedQueriesResponse
	(*GetAllNamedQueriesActiveDeploymentRequest)(nil),  // 2: chalk.server.v1.GetAllNamedQueriesActiveDeploymentRequest
	(*GetAllNamedQueriesActiveDeploymentResponse)(nil), // 3: chalk.server.v1.GetAllNamedQueriesActiveDeploymentResponse
	(*v1.NamedQuery)(nil),                              // 4: chalk.graph.v1.NamedQuery
}
var file_chalk_server_v1_named_query_proto_depIdxs = []int32{
	4, // 0: chalk.server.v1.GetAllNamedQueriesResponse.named_queries:type_name -> chalk.graph.v1.NamedQuery
	4, // 1: chalk.server.v1.GetAllNamedQueriesActiveDeploymentResponse.named_queries:type_name -> chalk.graph.v1.NamedQuery
	0, // 2: chalk.server.v1.NamedQueryService.GetAllNamedQueries:input_type -> chalk.server.v1.GetAllNamedQueriesRequest
	2, // 3: chalk.server.v1.NamedQueryService.GetAllNamedQueriesActiveDeployment:input_type -> chalk.server.v1.GetAllNamedQueriesActiveDeploymentRequest
	1, // 4: chalk.server.v1.NamedQueryService.GetAllNamedQueries:output_type -> chalk.server.v1.GetAllNamedQueriesResponse
	3, // 5: chalk.server.v1.NamedQueryService.GetAllNamedQueriesActiveDeployment:output_type -> chalk.server.v1.GetAllNamedQueriesActiveDeploymentResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_chalk_server_v1_named_query_proto_init() }
func file_chalk_server_v1_named_query_proto_init() {
	if File_chalk_server_v1_named_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chalk_server_v1_named_query_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllNamedQueriesRequest); i {
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
		file_chalk_server_v1_named_query_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllNamedQueriesResponse); i {
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
		file_chalk_server_v1_named_query_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllNamedQueriesActiveDeploymentRequest); i {
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
		file_chalk_server_v1_named_query_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllNamedQueriesActiveDeploymentResponse); i {
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
			RawDescriptor: file_chalk_server_v1_named_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chalk_server_v1_named_query_proto_goTypes,
		DependencyIndexes: file_chalk_server_v1_named_query_proto_depIdxs,
		MessageInfos:      file_chalk_server_v1_named_query_proto_msgTypes,
	}.Build()
	File_chalk_server_v1_named_query_proto = out.File
	file_chalk_server_v1_named_query_proto_rawDesc = nil
	file_chalk_server_v1_named_query_proto_goTypes = nil
	file_chalk_server_v1_named_query_proto_depIdxs = nil
}
