// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: chalk/aggregate/v1/backfill.proto

package aggregatev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type AggregateBackfillCostEstimate struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	MaxBuckets          int64                  `protobuf:"varint,1,opt,name=max_buckets,json=maxBuckets,proto3" json:"max_buckets,omitempty"`
	ExpectedBuckets     int64                  `protobuf:"varint,2,opt,name=expected_buckets,json=expectedBuckets,proto3" json:"expected_buckets,omitempty"`
	ExpectedBytes       int64                  `protobuf:"varint,3,opt,name=expected_bytes,json=expectedBytes,proto3" json:"expected_bytes,omitempty"`
	ExpectedStorageCost float64                `protobuf:"fixed64,4,opt,name=expected_storage_cost,json=expectedStorageCost,proto3" json:"expected_storage_cost,omitempty"`
	ExpectedRuntime     *durationpb.Duration   `protobuf:"bytes,5,opt,name=expected_runtime,json=expectedRuntime,proto3" json:"expected_runtime,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *AggregateBackfillCostEstimate) Reset() {
	*x = AggregateBackfillCostEstimate{}
	mi := &file_chalk_aggregate_v1_backfill_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AggregateBackfillCostEstimate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateBackfillCostEstimate) ProtoMessage() {}

func (x *AggregateBackfillCostEstimate) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_aggregate_v1_backfill_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateBackfillCostEstimate.ProtoReflect.Descriptor instead.
func (*AggregateBackfillCostEstimate) Descriptor() ([]byte, []int) {
	return file_chalk_aggregate_v1_backfill_proto_rawDescGZIP(), []int{0}
}

func (x *AggregateBackfillCostEstimate) GetMaxBuckets() int64 {
	if x != nil {
		return x.MaxBuckets
	}
	return 0
}

func (x *AggregateBackfillCostEstimate) GetExpectedBuckets() int64 {
	if x != nil {
		return x.ExpectedBuckets
	}
	return 0
}

func (x *AggregateBackfillCostEstimate) GetExpectedBytes() int64 {
	if x != nil {
		return x.ExpectedBytes
	}
	return 0
}

func (x *AggregateBackfillCostEstimate) GetExpectedStorageCost() float64 {
	if x != nil {
		return x.ExpectedStorageCost
	}
	return 0
}

func (x *AggregateBackfillCostEstimate) GetExpectedRuntime() *durationpb.Duration {
	if x != nil {
		return x.ExpectedRuntime
	}
	return nil
}

type AggregateBackfillUserParams struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	Features []string               `protobuf:"bytes,1,rep,name=features,proto3" json:"features,omitempty"`
	Resolver *string                `protobuf:"bytes,2,opt,name=resolver,proto3,oneof" json:"resolver,omitempty"`
	// Deprecated: Marked as deprecated in chalk/aggregate/v1/backfill.proto.
	TimestampColumnName *string                `protobuf:"bytes,3,opt,name=timestamp_column_name,json=timestampColumnName,proto3,oneof" json:"timestamp_column_name,omitempty"`
	LowerBound          *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=lower_bound,json=lowerBound,proto3,oneof" json:"lower_bound,omitempty"`
	UpperBound          *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=upper_bound,json=upperBound,proto3,oneof" json:"upper_bound,omitempty"`
	Exact               bool                   `protobuf:"varint,6,opt,name=exact,proto3" json:"exact,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *AggregateBackfillUserParams) Reset() {
	*x = AggregateBackfillUserParams{}
	mi := &file_chalk_aggregate_v1_backfill_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AggregateBackfillUserParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateBackfillUserParams) ProtoMessage() {}

func (x *AggregateBackfillUserParams) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_aggregate_v1_backfill_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateBackfillUserParams.ProtoReflect.Descriptor instead.
func (*AggregateBackfillUserParams) Descriptor() ([]byte, []int) {
	return file_chalk_aggregate_v1_backfill_proto_rawDescGZIP(), []int{1}
}

func (x *AggregateBackfillUserParams) GetFeatures() []string {
	if x != nil {
		return x.Features
	}
	return nil
}

func (x *AggregateBackfillUserParams) GetResolver() string {
	if x != nil && x.Resolver != nil {
		return *x.Resolver
	}
	return ""
}

// Deprecated: Marked as deprecated in chalk/aggregate/v1/backfill.proto.
func (x *AggregateBackfillUserParams) GetTimestampColumnName() string {
	if x != nil && x.TimestampColumnName != nil {
		return *x.TimestampColumnName
	}
	return ""
}

func (x *AggregateBackfillUserParams) GetLowerBound() *timestamppb.Timestamp {
	if x != nil {
		return x.LowerBound
	}
	return nil
}

func (x *AggregateBackfillUserParams) GetUpperBound() *timestamppb.Timestamp {
	if x != nil {
		return x.UpperBound
	}
	return nil
}

func (x *AggregateBackfillUserParams) GetExact() bool {
	if x != nil {
		return x.Exact
	}
	return false
}

type AggregateBackfill struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	Series             []*AggregateTimeSeries `protobuf:"bytes,1,rep,name=series,proto3" json:"series,omitempty"`
	Resolver           string                 `protobuf:"bytes,2,opt,name=resolver,proto3" json:"resolver,omitempty"`
	DatetimeFeature    string                 `protobuf:"bytes,3,opt,name=datetime_feature,json=datetimeFeature,proto3" json:"datetime_feature,omitempty"`
	BucketDuration     *durationpb.Duration   `protobuf:"bytes,4,opt,name=bucket_duration,json=bucketDuration,proto3" json:"bucket_duration,omitempty"`
	FiltersDescription string                 `protobuf:"bytes,5,opt,name=filters_description,json=filtersDescription,proto3" json:"filters_description,omitempty"`
	GroupBy            []string               `protobuf:"bytes,6,rep,name=group_by,json=groupBy,proto3" json:"group_by,omitempty"`
	MaxRetention       *durationpb.Duration   `protobuf:"bytes,7,opt,name=max_retention,json=maxRetention,proto3" json:"max_retention,omitempty"`
	LowerBound         *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=lower_bound,json=lowerBound,proto3" json:"lower_bound,omitempty"`
	UpperBound         *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=upper_bound,json=upperBound,proto3" json:"upper_bound,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *AggregateBackfill) Reset() {
	*x = AggregateBackfill{}
	mi := &file_chalk_aggregate_v1_backfill_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AggregateBackfill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateBackfill) ProtoMessage() {}

func (x *AggregateBackfill) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_aggregate_v1_backfill_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateBackfill.ProtoReflect.Descriptor instead.
func (*AggregateBackfill) Descriptor() ([]byte, []int) {
	return file_chalk_aggregate_v1_backfill_proto_rawDescGZIP(), []int{2}
}

func (x *AggregateBackfill) GetSeries() []*AggregateTimeSeries {
	if x != nil {
		return x.Series
	}
	return nil
}

func (x *AggregateBackfill) GetResolver() string {
	if x != nil {
		return x.Resolver
	}
	return ""
}

func (x *AggregateBackfill) GetDatetimeFeature() string {
	if x != nil {
		return x.DatetimeFeature
	}
	return ""
}

func (x *AggregateBackfill) GetBucketDuration() *durationpb.Duration {
	if x != nil {
		return x.BucketDuration
	}
	return nil
}

func (x *AggregateBackfill) GetFiltersDescription() string {
	if x != nil {
		return x.FiltersDescription
	}
	return ""
}

func (x *AggregateBackfill) GetGroupBy() []string {
	if x != nil {
		return x.GroupBy
	}
	return nil
}

func (x *AggregateBackfill) GetMaxRetention() *durationpb.Duration {
	if x != nil {
		return x.MaxRetention
	}
	return nil
}

func (x *AggregateBackfill) GetLowerBound() *timestamppb.Timestamp {
	if x != nil {
		return x.LowerBound
	}
	return nil
}

func (x *AggregateBackfill) GetUpperBound() *timestamppb.Timestamp {
	if x != nil {
		return x.UpperBound
	}
	return nil
}

type AggregateBackfillWithCostEstimate struct {
	state         protoimpl.MessageState         `protogen:"open.v1"`
	Backfill      *AggregateBackfill             `protobuf:"bytes,1,opt,name=backfill,proto3" json:"backfill,omitempty"`
	Estimate      *AggregateBackfillCostEstimate `protobuf:"bytes,2,opt,name=estimate,proto3" json:"estimate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AggregateBackfillWithCostEstimate) Reset() {
	*x = AggregateBackfillWithCostEstimate{}
	mi := &file_chalk_aggregate_v1_backfill_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AggregateBackfillWithCostEstimate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateBackfillWithCostEstimate) ProtoMessage() {}

func (x *AggregateBackfillWithCostEstimate) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_aggregate_v1_backfill_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateBackfillWithCostEstimate.ProtoReflect.Descriptor instead.
func (*AggregateBackfillWithCostEstimate) Descriptor() ([]byte, []int) {
	return file_chalk_aggregate_v1_backfill_proto_rawDescGZIP(), []int{3}
}

func (x *AggregateBackfillWithCostEstimate) GetBackfill() *AggregateBackfill {
	if x != nil {
		return x.Backfill
	}
	return nil
}

func (x *AggregateBackfillWithCostEstimate) GetEstimate() *AggregateBackfillCostEstimate {
	if x != nil {
		return x.Estimate
	}
	return nil
}

type AggregateBackfillJob struct {
	state                   protoimpl.MessageState `protogen:"open.v1"`
	Id                      string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EnvironmentId           string                 `protobuf:"bytes,2,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	Resolver                *string                `protobuf:"bytes,3,opt,name=resolver,proto3,oneof" json:"resolver,omitempty"`
	Features                []string               `protobuf:"bytes,4,rep,name=features,proto3" json:"features,omitempty"`
	AgentId                 *string                `protobuf:"bytes,5,opt,name=agent_id,json=agentId,proto3,oneof" json:"agent_id,omitempty"`
	DeploymentId            *string                `protobuf:"bytes,6,opt,name=deployment_id,json=deploymentId,proto3,oneof" json:"deployment_id,omitempty"`
	CreatedAt               *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt               *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Resolvers               []string               `protobuf:"bytes,9,rep,name=resolvers,proto3" json:"resolvers,omitempty"`
	CronAggregateBackfillId *string                `protobuf:"bytes,10,opt,name=cron_aggregate_backfill_id,json=cronAggregateBackfillId,proto3,oneof" json:"cron_aggregate_backfill_id,omitempty"`
	PlanHash                *string                `protobuf:"bytes,11,opt,name=plan_hash,json=planHash,proto3,oneof" json:"plan_hash,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *AggregateBackfillJob) Reset() {
	*x = AggregateBackfillJob{}
	mi := &file_chalk_aggregate_v1_backfill_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AggregateBackfillJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateBackfillJob) ProtoMessage() {}

func (x *AggregateBackfillJob) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_aggregate_v1_backfill_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateBackfillJob.ProtoReflect.Descriptor instead.
func (*AggregateBackfillJob) Descriptor() ([]byte, []int) {
	return file_chalk_aggregate_v1_backfill_proto_rawDescGZIP(), []int{4}
}

func (x *AggregateBackfillJob) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AggregateBackfillJob) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *AggregateBackfillJob) GetResolver() string {
	if x != nil && x.Resolver != nil {
		return *x.Resolver
	}
	return ""
}

func (x *AggregateBackfillJob) GetFeatures() []string {
	if x != nil {
		return x.Features
	}
	return nil
}

func (x *AggregateBackfillJob) GetAgentId() string {
	if x != nil && x.AgentId != nil {
		return *x.AgentId
	}
	return ""
}

func (x *AggregateBackfillJob) GetDeploymentId() string {
	if x != nil && x.DeploymentId != nil {
		return *x.DeploymentId
	}
	return ""
}

func (x *AggregateBackfillJob) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *AggregateBackfillJob) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *AggregateBackfillJob) GetResolvers() []string {
	if x != nil {
		return x.Resolvers
	}
	return nil
}

func (x *AggregateBackfillJob) GetCronAggregateBackfillId() string {
	if x != nil && x.CronAggregateBackfillId != nil {
		return *x.CronAggregateBackfillId
	}
	return ""
}

func (x *AggregateBackfillJob) GetPlanHash() string {
	if x != nil && x.PlanHash != nil {
		return *x.PlanHash
	}
	return ""
}

type CronAggregateBackfill struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EnvironmentId string                 `protobuf:"bytes,2,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	DeploymentId  string                 `protobuf:"bytes,3,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	Schedule      string                 `protobuf:"bytes,4,opt,name=schedule,proto3" json:"schedule,omitempty"`
	PlanHash      string                 `protobuf:"bytes,5,opt,name=plan_hash,json=planHash,proto3" json:"plan_hash,omitempty"`
	Features      []string               `protobuf:"bytes,8,rep,name=features,proto3" json:"features,omitempty"`
	Resolvers     []string               `protobuf:"bytes,9,rep,name=resolvers,proto3" json:"resolvers,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CronAggregateBackfill) Reset() {
	*x = CronAggregateBackfill{}
	mi := &file_chalk_aggregate_v1_backfill_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CronAggregateBackfill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CronAggregateBackfill) ProtoMessage() {}

func (x *CronAggregateBackfill) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_aggregate_v1_backfill_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CronAggregateBackfill.ProtoReflect.Descriptor instead.
func (*CronAggregateBackfill) Descriptor() ([]byte, []int) {
	return file_chalk_aggregate_v1_backfill_proto_rawDescGZIP(), []int{5}
}

func (x *CronAggregateBackfill) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CronAggregateBackfill) GetEnvironmentId() string {
	if x != nil {
		return x.EnvironmentId
	}
	return ""
}

func (x *CronAggregateBackfill) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *CronAggregateBackfill) GetSchedule() string {
	if x != nil {
		return x.Schedule
	}
	return ""
}

func (x *CronAggregateBackfill) GetPlanHash() string {
	if x != nil {
		return x.PlanHash
	}
	return ""
}

func (x *CronAggregateBackfill) GetFeatures() []string {
	if x != nil {
		return x.Features
	}
	return nil
}

func (x *CronAggregateBackfill) GetResolvers() []string {
	if x != nil {
		return x.Resolvers
	}
	return nil
}

func (x *CronAggregateBackfill) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CronAggregateBackfill) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_chalk_aggregate_v1_backfill_proto protoreflect.FileDescriptor

var file_chalk_aggregate_v1_backfill_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x23, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x02,
	0x0a, 0x1d, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x66,
	0x69, 0x6c, 0x6c, 0x43, 0x6f, 0x73, 0x74, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x12, 0x29, 0x0a, 0x10, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x65,
	0x63, 0x74, 0x65, 0x64, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x65,
	0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x13, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x10, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x5f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x65, 0x78, 0x70,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x22, 0xf8, 0x02, 0x0a,
	0x1b, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69,
	0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x15, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x48, 0x01, 0x52, 0x13,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e,
	0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x0b, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52, 0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x42, 0x6f, 0x75, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x70, 0x65,
	0x72, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x03, 0x52, 0x0a, 0x75, 0x70, 0x70,
	0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78,
	0x61, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x78, 0x61, 0x63, 0x74,
	0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x42, 0x18, 0x0a,
	0x16, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6c, 0x6f, 0x77, 0x65,
	0x72, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x75, 0x70, 0x70, 0x65,
	0x72, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0xe5, 0x03, 0x0a, 0x11, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x12, 0x3f, 0x0a,
	0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x61,
	0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x42, 0x0a, 0x0f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x62, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x13, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x42, 0x79, 0x12, 0x3e, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x74,
	0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x74, 0x65,
	0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x42, 0x6f, 0x75,
	0x6e, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x70, 0x65, 0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x22,
	0xb5, 0x01, 0x0a, 0x21, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63,
	0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x73, 0x74, 0x45, 0x73, 0x74,
	0x69, 0x6d, 0x61, 0x74, 0x65, 0x12, 0x41, 0x0a, 0x08, 0x62, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x52, 0x08,
	0x62, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x12, 0x4d, 0x0a, 0x08, 0x65, 0x73, 0x74, 0x69,
	0x6d, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c,
	0x6c, 0x43, 0x6f, 0x73, 0x74, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x52, 0x08, 0x65,
	0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x22, 0xa5, 0x04, 0x0a, 0x14, 0x41, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x4a, 0x6f, 0x62,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x28, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0c, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x72, 0x73, 0x12, 0x40, 0x0a, 0x1a, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x17, 0x63, 0x72, 0x6f, 0x6e, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x6e, 0x48,
	0x61, 0x73, 0x68, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x72, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x42, 0x1d, 0x0a, 0x1b, 0x5f, 0x63, 0x72, 0x6f, 0x6e, 0x5f, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x69,
	0x64, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x22,
	0xdc, 0x02, 0x0a, 0x15, 0x43, 0x72, 0x6f, 0x6e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x42, 0x61, 0x63, 0x6b, 0x66, 0x69, 0x6c, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x76,
	0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0xd2,
	0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x42, 0x61, 0x63, 0x6b, 0x66,
	0x69, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x41,
	0x58, 0xaa, 0x02, 0x12, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x41,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x43, 0x68,
	0x61, 0x6c, 0x6b, 0x5c, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x43,
	0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_aggregate_v1_backfill_proto_rawDescOnce sync.Once
	file_chalk_aggregate_v1_backfill_proto_rawDescData = file_chalk_aggregate_v1_backfill_proto_rawDesc
)

func file_chalk_aggregate_v1_backfill_proto_rawDescGZIP() []byte {
	file_chalk_aggregate_v1_backfill_proto_rawDescOnce.Do(func() {
		file_chalk_aggregate_v1_backfill_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_aggregate_v1_backfill_proto_rawDescData)
	})
	return file_chalk_aggregate_v1_backfill_proto_rawDescData
}

var file_chalk_aggregate_v1_backfill_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_chalk_aggregate_v1_backfill_proto_goTypes = []any{
	(*AggregateBackfillCostEstimate)(nil),     // 0: chalk.aggregate.v1.AggregateBackfillCostEstimate
	(*AggregateBackfillUserParams)(nil),       // 1: chalk.aggregate.v1.AggregateBackfillUserParams
	(*AggregateBackfill)(nil),                 // 2: chalk.aggregate.v1.AggregateBackfill
	(*AggregateBackfillWithCostEstimate)(nil), // 3: chalk.aggregate.v1.AggregateBackfillWithCostEstimate
	(*AggregateBackfillJob)(nil),              // 4: chalk.aggregate.v1.AggregateBackfillJob
	(*CronAggregateBackfill)(nil),             // 5: chalk.aggregate.v1.CronAggregateBackfill
	(*durationpb.Duration)(nil),               // 6: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil),             // 7: google.protobuf.Timestamp
	(*AggregateTimeSeries)(nil),               // 8: chalk.aggregate.v1.AggregateTimeSeries
}
var file_chalk_aggregate_v1_backfill_proto_depIdxs = []int32{
	6,  // 0: chalk.aggregate.v1.AggregateBackfillCostEstimate.expected_runtime:type_name -> google.protobuf.Duration
	7,  // 1: chalk.aggregate.v1.AggregateBackfillUserParams.lower_bound:type_name -> google.protobuf.Timestamp
	7,  // 2: chalk.aggregate.v1.AggregateBackfillUserParams.upper_bound:type_name -> google.protobuf.Timestamp
	8,  // 3: chalk.aggregate.v1.AggregateBackfill.series:type_name -> chalk.aggregate.v1.AggregateTimeSeries
	6,  // 4: chalk.aggregate.v1.AggregateBackfill.bucket_duration:type_name -> google.protobuf.Duration
	6,  // 5: chalk.aggregate.v1.AggregateBackfill.max_retention:type_name -> google.protobuf.Duration
	7,  // 6: chalk.aggregate.v1.AggregateBackfill.lower_bound:type_name -> google.protobuf.Timestamp
	7,  // 7: chalk.aggregate.v1.AggregateBackfill.upper_bound:type_name -> google.protobuf.Timestamp
	2,  // 8: chalk.aggregate.v1.AggregateBackfillWithCostEstimate.backfill:type_name -> chalk.aggregate.v1.AggregateBackfill
	0,  // 9: chalk.aggregate.v1.AggregateBackfillWithCostEstimate.estimate:type_name -> chalk.aggregate.v1.AggregateBackfillCostEstimate
	7,  // 10: chalk.aggregate.v1.AggregateBackfillJob.created_at:type_name -> google.protobuf.Timestamp
	7,  // 11: chalk.aggregate.v1.AggregateBackfillJob.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 12: chalk.aggregate.v1.CronAggregateBackfill.created_at:type_name -> google.protobuf.Timestamp
	7,  // 13: chalk.aggregate.v1.CronAggregateBackfill.updated_at:type_name -> google.protobuf.Timestamp
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_chalk_aggregate_v1_backfill_proto_init() }
func file_chalk_aggregate_v1_backfill_proto_init() {
	if File_chalk_aggregate_v1_backfill_proto != nil {
		return
	}
	file_chalk_aggregate_v1_timeseries_proto_init()
	file_chalk_aggregate_v1_backfill_proto_msgTypes[1].OneofWrappers = []any{}
	file_chalk_aggregate_v1_backfill_proto_msgTypes[4].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_aggregate_v1_backfill_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chalk_aggregate_v1_backfill_proto_goTypes,
		DependencyIndexes: file_chalk_aggregate_v1_backfill_proto_depIdxs,
		MessageInfos:      file_chalk_aggregate_v1_backfill_proto_msgTypes,
	}.Build()
	File_chalk_aggregate_v1_backfill_proto = out.File
	file_chalk_aggregate_v1_backfill_proto_rawDesc = nil
	file_chalk_aggregate_v1_backfill_proto_goTypes = nil
	file_chalk_aggregate_v1_backfill_proto_depIdxs = nil
}
