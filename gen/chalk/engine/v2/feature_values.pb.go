// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: chalk/engine/v2/feature_values.proto

package enginev2

import (
	v11 "github.com/chalk-ai/chalk-go/gen/chalk/arrow/v1"
	v1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type FeatureValuesTimestampType int32

const (
	FeatureValuesTimestampType_FEATURE_VALUES_TIMESTAMP_TYPE_UNSPECIFIED FeatureValuesTimestampType = 0
	FeatureValuesTimestampType_FEATURE_VALUES_TIMESTAMP_TYPE_INSERTED_AT FeatureValuesTimestampType = 1
	FeatureValuesTimestampType_FEATURE_VALUES_TIMESTAMP_TYPE_OBSERVED_AT FeatureValuesTimestampType = 2
)

// Enum value maps for FeatureValuesTimestampType.
var (
	FeatureValuesTimestampType_name = map[int32]string{
		0: "FEATURE_VALUES_TIMESTAMP_TYPE_UNSPECIFIED",
		1: "FEATURE_VALUES_TIMESTAMP_TYPE_INSERTED_AT",
		2: "FEATURE_VALUES_TIMESTAMP_TYPE_OBSERVED_AT",
	}
	FeatureValuesTimestampType_value = map[string]int32{
		"FEATURE_VALUES_TIMESTAMP_TYPE_UNSPECIFIED": 0,
		"FEATURE_VALUES_TIMESTAMP_TYPE_INSERTED_AT": 1,
		"FEATURE_VALUES_TIMESTAMP_TYPE_OBSERVED_AT": 2,
	}
)

func (x FeatureValuesTimestampType) Enum() *FeatureValuesTimestampType {
	p := new(FeatureValuesTimestampType)
	*p = x
	return p
}

func (x FeatureValuesTimestampType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeatureValuesTimestampType) Descriptor() protoreflect.EnumDescriptor {
	return file_chalk_engine_v2_feature_values_proto_enumTypes[0].Descriptor()
}

func (FeatureValuesTimestampType) Type() protoreflect.EnumType {
	return &file_chalk_engine_v2_feature_values_proto_enumTypes[0]
}

func (x FeatureValuesTimestampType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeatureValuesTimestampType.Descriptor instead.
func (FeatureValuesTimestampType) EnumDescriptor() ([]byte, []int) {
	return file_chalk_engine_v2_feature_values_proto_rawDescGZIP(), []int{0}
}

type FeatureValueFilters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResolverFqn   []string           `protobuf:"bytes,1,rep,name=resolver_fqn,json=resolverFqn,proto3" json:"resolver_fqn,omitempty"`
	DeploymentId  []string           `protobuf:"bytes,2,rep,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	OperationId   []string           `protobuf:"bytes,3,rep,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	OperationKind []v1.OperationKind `protobuf:"varint,4,rep,packed,name=operation_kind,json=operationKind,proto3,enum=chalk.common.v1.OperationKind" json:"operation_kind,omitempty"`
	PrimaryKey    []*v11.ScalarValue `protobuf:"bytes,5,rep,name=primary_key,json=primaryKey,proto3" json:"primary_key,omitempty"`
}

func (x *FeatureValueFilters) Reset() {
	*x = FeatureValueFilters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_engine_v2_feature_values_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureValueFilters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureValueFilters) ProtoMessage() {}

func (x *FeatureValueFilters) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_engine_v2_feature_values_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureValueFilters.ProtoReflect.Descriptor instead.
func (*FeatureValueFilters) Descriptor() ([]byte, []int) {
	return file_chalk_engine_v2_feature_values_proto_rawDescGZIP(), []int{0}
}

func (x *FeatureValueFilters) GetResolverFqn() []string {
	if x != nil {
		return x.ResolverFqn
	}
	return nil
}

func (x *FeatureValueFilters) GetDeploymentId() []string {
	if x != nil {
		return x.DeploymentId
	}
	return nil
}

func (x *FeatureValueFilters) GetOperationId() []string {
	if x != nil {
		return x.OperationId
	}
	return nil
}

func (x *FeatureValueFilters) GetOperationKind() []v1.OperationKind {
	if x != nil {
		return x.OperationKind
	}
	return nil
}

func (x *FeatureValueFilters) GetPrimaryKey() []*v11.ScalarValue {
	if x != nil {
		return x.PrimaryKey
	}
	return nil
}

// Internal protobuf representing a next page token. Contains the operation id and the query timestamp for the last row
// in the previous batch. Results are sorted query timestamp, then by operation id, then observation id lexicographically,
// so this is all we need to know where the next page begins
type GetFeatureValuesPageToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimestampHwm     *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp_hwm,json=timestampHwm,proto3" json:"timestamp_hwm,omitempty"`
	OperationIdHwm   string                 `protobuf:"bytes,2,opt,name=operation_id_hwm,json=operationIdHwm,proto3" json:"operation_id_hwm,omitempty"`
	ObservationIdHwm string                 `protobuf:"bytes,3,opt,name=observation_id_hwm,json=observationIdHwm,proto3" json:"observation_id_hwm,omitempty"`
}

func (x *GetFeatureValuesPageToken) Reset() {
	*x = GetFeatureValuesPageToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_engine_v2_feature_values_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeatureValuesPageToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeatureValuesPageToken) ProtoMessage() {}

func (x *GetFeatureValuesPageToken) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_engine_v2_feature_values_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeatureValuesPageToken.ProtoReflect.Descriptor instead.
func (*GetFeatureValuesPageToken) Descriptor() ([]byte, []int) {
	return file_chalk_engine_v2_feature_values_proto_rawDescGZIP(), []int{1}
}

func (x *GetFeatureValuesPageToken) GetTimestampHwm() *timestamppb.Timestamp {
	if x != nil {
		return x.TimestampHwm
	}
	return nil
}

func (x *GetFeatureValuesPageToken) GetOperationIdHwm() string {
	if x != nil {
		return x.OperationIdHwm
	}
	return ""
}

func (x *GetFeatureValuesPageToken) GetObservationIdHwm() string {
	if x != nil {
		return x.ObservationIdHwm
	}
	return ""
}

type GetFeatureValuesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeatureFqn string `protobuf:"bytes,1,opt,name=feature_fqn,json=featureFqn,proto3" json:"feature_fqn,omitempty"`
	// Along which time dimension to query features by.
	// Time ranges can either be specified along inserted_at or observed_at
	// Optional. If unspecified, default to inserted_at
	TimestampType       FeatureValuesTimestampType `protobuf:"varint,2,opt,name=timestamp_type,json=timestampType,proto3,enum=chalk.engine.v2.FeatureValuesTimestampType" json:"timestamp_type,omitempty"`
	LowerBoundInclusive *timestamppb.Timestamp     `protobuf:"bytes,3,opt,name=lower_bound_inclusive,json=lowerBoundInclusive,proto3" json:"lower_bound_inclusive,omitempty"`
	UpperBoundExclusive *timestamppb.Timestamp     `protobuf:"bytes,4,opt,name=upper_bound_exclusive,json=upperBoundExclusive,proto3" json:"upper_bound_exclusive,omitempty"`
	// optional filters that can be specified to filter down feature rows returned
	Filters *FeatureValueFilters `protobuf:"bytes,5,opt,name=filters,proto3" json:"filters,omitempty"`
	// The (maximum) page size for results. If zero, then the server picks.
	PageSize int32 `protobuf:"varint,6,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// When dealing with paginated responses, the next token can be provided to get the next page of responses
	// The query parameters above must be exactly the same when using a next token
	// This is opaque to the clients, but in practice, it is an encoding of the GetFeatureValuesPageToken
	// if unspecified, retrieve the zeroth page.
	PageToken string `protobuf:"bytes,7,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *GetFeatureValuesRequest) Reset() {
	*x = GetFeatureValuesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_engine_v2_feature_values_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeatureValuesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeatureValuesRequest) ProtoMessage() {}

func (x *GetFeatureValuesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_engine_v2_feature_values_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeatureValuesRequest.ProtoReflect.Descriptor instead.
func (*GetFeatureValuesRequest) Descriptor() ([]byte, []int) {
	return file_chalk_engine_v2_feature_values_proto_rawDescGZIP(), []int{2}
}

func (x *GetFeatureValuesRequest) GetFeatureFqn() string {
	if x != nil {
		return x.FeatureFqn
	}
	return ""
}

func (x *GetFeatureValuesRequest) GetTimestampType() FeatureValuesTimestampType {
	if x != nil {
		return x.TimestampType
	}
	return FeatureValuesTimestampType_FEATURE_VALUES_TIMESTAMP_TYPE_UNSPECIFIED
}

func (x *GetFeatureValuesRequest) GetLowerBoundInclusive() *timestamppb.Timestamp {
	if x != nil {
		return x.LowerBoundInclusive
	}
	return nil
}

func (x *GetFeatureValuesRequest) GetUpperBoundExclusive() *timestamppb.Timestamp {
	if x != nil {
		return x.UpperBoundExclusive
	}
	return nil
}

func (x *GetFeatureValuesRequest) GetFilters() *FeatureValueFilters {
	if x != nil {
		return x.Filters
	}
	return nil
}

func (x *GetFeatureValuesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetFeatureValuesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type GetFeatureValuesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If non-empty, call this endpoint again, with this next token to get the next page of responses.
	NextPageToken string `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// The number of results returned
	TotalSize int32 `protobuf:"varint,2,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	// The response payload. Defining as a one-of to future-proof it should we decide to support multiple encodings (parquet, feather, ...)
	//
	// Types that are assignable to Payload:
	//
	//	*GetFeatureValuesResponse_Parquet
	Payload isGetFeatureValuesResponse_Payload `protobuf_oneof:"payload"`
}

func (x *GetFeatureValuesResponse) Reset() {
	*x = GetFeatureValuesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_engine_v2_feature_values_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeatureValuesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeatureValuesResponse) ProtoMessage() {}

func (x *GetFeatureValuesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_engine_v2_feature_values_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeatureValuesResponse.ProtoReflect.Descriptor instead.
func (*GetFeatureValuesResponse) Descriptor() ([]byte, []int) {
	return file_chalk_engine_v2_feature_values_proto_rawDescGZIP(), []int{3}
}

func (x *GetFeatureValuesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (x *GetFeatureValuesResponse) GetTotalSize() int32 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

func (m *GetFeatureValuesResponse) GetPayload() isGetFeatureValuesResponse_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *GetFeatureValuesResponse) GetParquet() []byte {
	if x, ok := x.GetPayload().(*GetFeatureValuesResponse_Parquet); ok {
		return x.Parquet
	}
	return nil
}

type isGetFeatureValuesResponse_Payload interface {
	isGetFeatureValuesResponse_Payload()
}

type GetFeatureValuesResponse_Parquet struct {
	Parquet []byte `protobuf:"bytes,3,opt,name=parquet,proto3,oneof"`
}

func (*GetFeatureValuesResponse_Parquet) isGetFeatureValuesResponse_Payload() {}

var File_chalk_engine_v2_feature_values_proto protoreflect.FileDescriptor

var file_chalk_engine_v2_feature_values_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x32, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x32, 0x1a, 0x1a, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61,
	0x72, 0x72, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b,
	0x69, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x02, 0x0a, 0x13, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x5f, 0x66,
	0x71, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x72, 0x46, 0x71, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x45, 0x0a,
	0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4b, 0x69, 0x6e, 0x64, 0x12, 0x3c, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2e, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61,
	0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b,
	0x65, 0x79, 0x22, 0xb4, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x3f, 0x0a, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x68, 0x77,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x77,
	0x6d, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x5f, 0x68, 0x77, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x48, 0x77, 0x6d, 0x12, 0x2c, 0x0a, 0x12, 0x6f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x68, 0x77,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x48, 0x77, 0x6d, 0x22, 0xaa, 0x03, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x5f, 0x66, 0x71, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x46, 0x71, 0x6e, 0x12, 0x52, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b,
	0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x32,
	0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4e, 0x0a, 0x15, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73,
	0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x13, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e,
	0x64, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x12, 0x4e, 0x0a, 0x15, 0x75, 0x70,
	0x70, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73,
	0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x13, 0x75, 0x70, 0x70, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e,
	0x64, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x88, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65,
	0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x07, 0x70, 0x61,
	0x72, 0x71, 0x75, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x70,
	0x61, 0x72, 0x71, 0x75, 0x65, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x2a, 0xa9, 0x01, 0x0a, 0x1a, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x2d, 0x0a, 0x29, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55,
	0x45, 0x53, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x2d, 0x0a, 0x29, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x53, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x49, 0x4e, 0x53, 0x45, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x10, 0x01, 0x12, 0x2d,
	0x0a, 0x29, 0x46, 0x45, 0x41, 0x54, 0x55, 0x52, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x53,
	0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4f, 0x42, 0x53, 0x45, 0x52, 0x56, 0x45, 0x44, 0x5f, 0x41, 0x54, 0x10, 0x02, 0x42, 0xc2, 0x01,
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x32, 0x42, 0x12, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69,
	0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x32, 0x3b, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x43, 0x45, 0x58, 0xaa, 0x02, 0x0f, 0x43,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x32, 0xca, 0x02,
	0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x32,
	0xe2, 0x02, 0x1b, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c,
	0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x11, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a,
	0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_engine_v2_feature_values_proto_rawDescOnce sync.Once
	file_chalk_engine_v2_feature_values_proto_rawDescData = file_chalk_engine_v2_feature_values_proto_rawDesc
)

func file_chalk_engine_v2_feature_values_proto_rawDescGZIP() []byte {
	file_chalk_engine_v2_feature_values_proto_rawDescOnce.Do(func() {
		file_chalk_engine_v2_feature_values_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_engine_v2_feature_values_proto_rawDescData)
	})
	return file_chalk_engine_v2_feature_values_proto_rawDescData
}

var file_chalk_engine_v2_feature_values_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_chalk_engine_v2_feature_values_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chalk_engine_v2_feature_values_proto_goTypes = []any{
	(FeatureValuesTimestampType)(0),   // 0: chalk.engine.v2.FeatureValuesTimestampType
	(*FeatureValueFilters)(nil),       // 1: chalk.engine.v2.FeatureValueFilters
	(*GetFeatureValuesPageToken)(nil), // 2: chalk.engine.v2.GetFeatureValuesPageToken
	(*GetFeatureValuesRequest)(nil),   // 3: chalk.engine.v2.GetFeatureValuesRequest
	(*GetFeatureValuesResponse)(nil),  // 4: chalk.engine.v2.GetFeatureValuesResponse
	(v1.OperationKind)(0),             // 5: chalk.common.v1.OperationKind
	(*v11.ScalarValue)(nil),           // 6: chalk.arrow.v1.ScalarValue
	(*timestamppb.Timestamp)(nil),     // 7: google.protobuf.Timestamp
}
var file_chalk_engine_v2_feature_values_proto_depIdxs = []int32{
	5, // 0: chalk.engine.v2.FeatureValueFilters.operation_kind:type_name -> chalk.common.v1.OperationKind
	6, // 1: chalk.engine.v2.FeatureValueFilters.primary_key:type_name -> chalk.arrow.v1.ScalarValue
	7, // 2: chalk.engine.v2.GetFeatureValuesPageToken.timestamp_hwm:type_name -> google.protobuf.Timestamp
	0, // 3: chalk.engine.v2.GetFeatureValuesRequest.timestamp_type:type_name -> chalk.engine.v2.FeatureValuesTimestampType
	7, // 4: chalk.engine.v2.GetFeatureValuesRequest.lower_bound_inclusive:type_name -> google.protobuf.Timestamp
	7, // 5: chalk.engine.v2.GetFeatureValuesRequest.upper_bound_exclusive:type_name -> google.protobuf.Timestamp
	1, // 6: chalk.engine.v2.GetFeatureValuesRequest.filters:type_name -> chalk.engine.v2.FeatureValueFilters
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_chalk_engine_v2_feature_values_proto_init() }
func file_chalk_engine_v2_feature_values_proto_init() {
	if File_chalk_engine_v2_feature_values_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chalk_engine_v2_feature_values_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*FeatureValueFilters); i {
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
		file_chalk_engine_v2_feature_values_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetFeatureValuesPageToken); i {
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
		file_chalk_engine_v2_feature_values_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetFeatureValuesRequest); i {
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
		file_chalk_engine_v2_feature_values_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetFeatureValuesResponse); i {
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
	file_chalk_engine_v2_feature_values_proto_msgTypes[3].OneofWrappers = []any{
		(*GetFeatureValuesResponse_Parquet)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_engine_v2_feature_values_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chalk_engine_v2_feature_values_proto_goTypes,
		DependencyIndexes: file_chalk_engine_v2_feature_values_proto_depIdxs,
		EnumInfos:         file_chalk_engine_v2_feature_values_proto_enumTypes,
		MessageInfos:      file_chalk_engine_v2_feature_values_proto_msgTypes,
	}.Build()
	File_chalk_engine_v2_feature_values_proto = out.File
	file_chalk_engine_v2_feature_values_proto_rawDesc = nil
	file_chalk_engine_v2_feature_values_proto_goTypes = nil
	file_chalk_engine_v2_feature_values_proto_depIdxs = nil
}
