// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: chalk/pubsub/v1/queries.proto

package pubsubv1

import (
	_ "github.com/chalk-ai/chalk-go/gen/gen_bq_schema"
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

// QueryMessagePubSub corresponds exactly to the QueryStorageMessage pydantic model,
// which in turn mostly corresponds to the QueryStorageRequest pydantic model.
// It captures information about an online query from the header, plan, parsed metadata, and more.
type QueryMessagePubSub struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Operation unique identifier
	OperationId string `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	// Environment query belongs to
	EnvironmentId string `protobuf:"bytes,2,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	// maps to FeatherRequestHeader.meta which can be specified via the client
	QueryMeta map[string]string `protobuf:"bytes,3,rep,name=query_meta,json=queryMeta,proto3" json:"query_meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// maps to FeatherRequestHeader.query_name which can be specified via the client
	QueryName *string `protobuf:"bytes,4,opt,name=query_name,json=queryName,proto3,oneof" json:"query_name,omitempty"`
	// maps to FeatherRequestHeader.query_name_version which can be specified via the client
	QueryNameVersion *string `protobuf:"bytes,5,opt,name=query_name_version,json=queryNameVersion,proto3,oneof" json:"query_name_version,omitempty"`
	// maps to FeatherRequestHeader.correlation_id which can be specified via the client
	CorrelationId *string `protobuf:"bytes,6,opt,name=correlation_id,json=correlationId,proto3,oneof" json:"correlation_id,omitempty"`
	// parsed input feature fqns
	InputFeatures []string `protobuf:"bytes,7,rep,name=input_features,json=inputFeatures,proto3" json:"input_features,omitempty"`
	// parsed output feature fqns
	OutputFeatures []string `protobuf:"bytes,8,rep,name=output_features,json=outputFeatures,proto3" json:"output_features,omitempty"`
	// parsed output feature root fqns
	OutputRootFqns []string `protobuf:"bytes,9,rep,name=output_root_fqns,json=outputRootFqns,proto3" json:"output_root_fqns,omitempty"`
	// intermediate feature fqns currently empty
	IntermediateFeatures []string `protobuf:"bytes,10,rep,name=intermediate_features,json=intermediateFeatures,proto3" json:"intermediate_features,omitempty"`
	// resolver fqns used in plan
	Resolvers []string `protobuf:"bytes,11,rep,name=resolvers,proto3" json:"resolvers,omitempty"`
	// id of final plan
	QueryPlanId *string `protobuf:"bytes,12,opt,name=query_plan_id,json=queryPlanId,proto3,oneof" json:"query_plan_id,omitempty"`
	// timestamp
	CreatedAt int64 `protobuf:"varint,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// if query had errors
	HasErrors *bool `protobuf:"varint,14,opt,name=has_errors,json=hasErrors,proto3,oneof" json:"has_errors,omitempty"`
	// aka as requester_id from run context
	AgentId *string `protobuf:"bytes,15,opt,name=agent_id,json=agentId,proto3,oneof" json:"agent_id,omitempty"`
	// branch query was run on
	BranchName *string `protobuf:"bytes,16,opt,name=branch_name,json=branchName,proto3,oneof" json:"branch_name,omitempty"`
	// deployment query was run on
	DeploymentId *string `protobuf:"bytes,17,opt,name=deployment_id,json=deploymentId,proto3,oneof" json:"deployment_id,omitempty"`
	// if query was run with plan_stages
	HasPlanStages *bool `protobuf:"varint,18,opt,name=has_plan_stages,json=hasPlanStages,proto3,oneof" json:"has_plan_stages,omitempty"`
}

func (x *QueryMessagePubSub) Reset() {
	*x = QueryMessagePubSub{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_pubsub_v1_queries_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryMessagePubSub) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryMessagePubSub) ProtoMessage() {}

func (x *QueryMessagePubSub) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_pubsub_v1_queries_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryMessagePubSub.ProtoReflect.Descriptor instead.
func (*QueryMessagePubSub) Descriptor() ([]byte, []int) {
	return file_chalk_pubsub_v1_queries_proto_rawDescGZIP(), []int{0}
}

func (x *QueryMessagePubSub) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

func (x *QueryMessagePubSub) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *QueryMessagePubSub) GetQueryMeta() map[string]string {
	if x != nil {
		return x.QueryMeta
	}
	return nil
}

func (x *QueryMessagePubSub) GetQueryName() string {
	if x != nil && x.QueryName != nil {
		return *x.QueryName
	}
	return ""
}

func (x *QueryMessagePubSub) GetQueryNameVersion() string {
	if x != nil && x.QueryNameVersion != nil {
		return *x.QueryNameVersion
	}
	return ""
}

func (x *QueryMessagePubSub) GetCorrelationId() string {
	if x != nil && x.CorrelationId != nil {
		return *x.CorrelationId
	}
	return ""
}

func (x *QueryMessagePubSub) GetInputFeatures() []string {
	if x != nil {
		return x.InputFeatures
	}
	return nil
}

func (x *QueryMessagePubSub) GetOutputFeatures() []string {
	if x != nil {
		return x.OutputFeatures
	}
	return nil
}

func (x *QueryMessagePubSub) GetOutputRootFqns() []string {
	if x != nil {
		return x.OutputRootFqns
	}
	return nil
}

func (x *QueryMessagePubSub) GetIntermediateFeatures() []string {
	if x != nil {
		return x.IntermediateFeatures
	}
	return nil
}

func (x *QueryMessagePubSub) GetResolvers() []string {
	if x != nil {
		return x.Resolvers
	}
	return nil
}

func (x *QueryMessagePubSub) GetQueryPlanId() string {
	if x != nil && x.QueryPlanId != nil {
		return *x.QueryPlanId
	}
	return ""
}

func (x *QueryMessagePubSub) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *QueryMessagePubSub) GetHasErrors() bool {
	if x != nil && x.HasErrors != nil {
		return *x.HasErrors
	}
	return false
}

func (x *QueryMessagePubSub) GetAgentId() string {
	if x != nil && x.AgentId != nil {
		return *x.AgentId
	}
	return ""
}

func (x *QueryMessagePubSub) GetBranchName() string {
	if x != nil && x.BranchName != nil {
		return *x.BranchName
	}
	return ""
}

func (x *QueryMessagePubSub) GetDeploymentId() string {
	if x != nil && x.DeploymentId != nil {
		return *x.DeploymentId
	}
	return ""
}

func (x *QueryMessagePubSub) GetHasPlanStages() bool {
	if x != nil && x.HasPlanStages != nil {
		return *x.HasPlanStages
	}
	return false
}

var File_chalk_pubsub_v1_queries_proto protoreflect.FileDescriptor

var file_chalk_pubsub_v1_queries_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2f, 0x76,
	0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x65, 0x6e, 0x5f, 0x62, 0x71, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f,
	0x62, 0x71, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x65, 0x6e, 0x5f, 0x62, 0x71, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x62, 0x71,
	0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x08, 0x0a,
	0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x75, 0x62,
	0x53, 0x75, 0x62, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x51, 0x0a,
	0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x32, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x61,
	0x12, 0x22, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x12, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x10, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x72, 0x6f,
	0x6f, 0x74, 0x5f, 0x66, 0x71, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x6f, 0x6f, 0x74, 0x46, 0x71, 0x6e, 0x73, 0x12, 0x33, 0x0a,
	0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x73, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x73,
	0x12, 0x27, 0x0a, 0x0d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x6c, 0x61, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0e, 0xea,
	0x3f, 0x0b, 0x12, 0x09, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x0a, 0x68, 0x61, 0x73, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x48, 0x04, 0x52, 0x09,
	0x68, 0x61, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a, 0x08,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05,
	0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x06, 0x52, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x0c, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2b, 0x0a, 0x0f,
	0x68, 0x61, 0x73, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x08, 0x48, 0x08, 0x52, 0x0d, 0x68, 0x61, 0x73, 0x50, 0x6c, 0x61, 0x6e,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x73, 0x88, 0x01, 0x01, 0x1a, 0x3c, 0x0a, 0x0e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x0c, 0xea, 0x3f, 0x09, 0x0a, 0x07, 0x71, 0x75,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x11, 0x0a, 0x0f, 0x5f,
	0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x42, 0x10,
	0x0a, 0x0e, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x68, 0x61, 0x73, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x10, 0x0a, 0x0e,
	0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x12,
	0x0a, 0x10, 0x5f, 0x68, 0x61, 0x73, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x73, 0x42, 0xbc, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x51, 0x75, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x75, 0x62,
	0x73, 0x75, 0x62, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x50, 0x58, 0xaa, 0x02, 0x0f, 0x43, 0x68,
	0x61, 0x6c, 0x6b, 0x2e, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f,
	0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1b, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0x5c, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11,
	0x43, 0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_pubsub_v1_queries_proto_rawDescOnce sync.Once
	file_chalk_pubsub_v1_queries_proto_rawDescData = file_chalk_pubsub_v1_queries_proto_rawDesc
)

func file_chalk_pubsub_v1_queries_proto_rawDescGZIP() []byte {
	file_chalk_pubsub_v1_queries_proto_rawDescOnce.Do(func() {
		file_chalk_pubsub_v1_queries_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_pubsub_v1_queries_proto_rawDescData)
	})
	return file_chalk_pubsub_v1_queries_proto_rawDescData
}

var file_chalk_pubsub_v1_queries_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chalk_pubsub_v1_queries_proto_goTypes = []interface{}{
	(*QueryMessagePubSub)(nil), // 0: chalk.pubsub.v1.QueryMessagePubSub
	nil,                        // 1: chalk.pubsub.v1.QueryMessagePubSub.QueryMetaEntry
}
var file_chalk_pubsub_v1_queries_proto_depIdxs = []int32{
	1, // 0: chalk.pubsub.v1.QueryMessagePubSub.query_meta:type_name -> chalk.pubsub.v1.QueryMessagePubSub.QueryMetaEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chalk_pubsub_v1_queries_proto_init() }
func file_chalk_pubsub_v1_queries_proto_init() {
	if File_chalk_pubsub_v1_queries_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chalk_pubsub_v1_queries_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryMessagePubSub); i {
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
	file_chalk_pubsub_v1_queries_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_pubsub_v1_queries_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chalk_pubsub_v1_queries_proto_goTypes,
		DependencyIndexes: file_chalk_pubsub_v1_queries_proto_depIdxs,
		MessageInfos:      file_chalk_pubsub_v1_queries_proto_msgTypes,
	}.Build()
	File_chalk_pubsub_v1_queries_proto = out.File
	file_chalk_pubsub_v1_queries_proto_rawDesc = nil
	file_chalk_pubsub_v1_queries_proto_goTypes = nil
	file_chalk_pubsub_v1_queries_proto_depIdxs = nil
}
