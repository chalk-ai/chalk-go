// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: chalk/server/v1/topicpush.proto

package serverv1

import (
	_ "github.com/chalk-ai/chalk-go/gen/chalk/auth/v1"
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

type ScheduledJob struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Payload     string            `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Schedule    string            `protobuf:"bytes,4,opt,name=schedule,proto3" json:"schedule,omitempty"`
	Topic       *Topic            `protobuf:"bytes,5,opt,name=topic,proto3" json:"topic,omitempty"`
	Attributes  map[string]string `protobuf:"bytes,6,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Environment *string           `protobuf:"bytes,7,opt,name=environment,proto3,oneof" json:"environment,omitempty"`
	Tags        []string          `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *ScheduledJob) Reset() {
	*x = ScheduledJob{}
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScheduledJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduledJob) ProtoMessage() {}

func (x *ScheduledJob) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduledJob.ProtoReflect.Descriptor instead.
func (*ScheduledJob) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_topicpush_proto_rawDescGZIP(), []int{0}
}

func (x *ScheduledJob) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ScheduledJob) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ScheduledJob) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *ScheduledJob) GetSchedule() string {
	if x != nil {
		return x.Schedule
	}
	return ""
}

func (x *ScheduledJob) GetTopic() *Topic {
	if x != nil {
		return x.Topic
	}
	return nil
}

func (x *ScheduledJob) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *ScheduledJob) GetEnvironment() string {
	if x != nil && x.Environment != nil {
		return *x.Environment
	}
	return ""
}

func (x *ScheduledJob) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type UpdateJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Payload     string            `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Schedule    string            `protobuf:"bytes,4,opt,name=schedule,proto3" json:"schedule,omitempty"`
	Topic       *Topic            `protobuf:"bytes,5,opt,name=topic,proto3" json:"topic,omitempty"`
	Attributes  map[string]string `protobuf:"bytes,6,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Environment *string           `protobuf:"bytes,7,opt,name=environment,proto3,oneof" json:"environment,omitempty"`
	Tags        []string          `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *UpdateJobRequest) Reset() {
	*x = UpdateJobRequest{}
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobRequest) ProtoMessage() {}

func (x *UpdateJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobRequest.ProtoReflect.Descriptor instead.
func (*UpdateJobRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_topicpush_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateJobRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateJobRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateJobRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *UpdateJobRequest) GetSchedule() string {
	if x != nil {
		return x.Schedule
	}
	return ""
}

func (x *UpdateJobRequest) GetTopic() *Topic {
	if x != nil {
		return x.Topic
	}
	return nil
}

func (x *UpdateJobRequest) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *UpdateJobRequest) GetEnvironment() string {
	if x != nil && x.Environment != nil {
		return *x.Environment
	}
	return ""
}

func (x *UpdateJobRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type UpdateJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job *ScheduledJob `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *UpdateJobResponse) Reset() {
	*x = UpdateJobResponse{}
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobResponse) ProtoMessage() {}

func (x *UpdateJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateJobResponse.ProtoReflect.Descriptor instead.
func (*UpdateJobResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_topicpush_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateJobResponse) GetJob() *ScheduledJob {
	if x != nil {
		return x.Job
	}
	return nil
}

type ListJobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListJobsRequest) Reset() {
	*x = ListJobsRequest{}
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsRequest) ProtoMessage() {}

func (x *ListJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsRequest.ProtoReflect.Descriptor instead.
func (*ListJobsRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_topicpush_proto_rawDescGZIP(), []int{3}
}

type ListJobsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Jobs []*ScheduledJob `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
}

func (x *ListJobsResponse) Reset() {
	*x = ListJobsResponse{}
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsResponse) ProtoMessage() {}

func (x *ListJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListJobsResponse.ProtoReflect.Descriptor instead.
func (*ListJobsResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_topicpush_proto_rawDescGZIP(), []int{4}
}

func (x *ListJobsResponse) GetJobs() []*ScheduledJob {
	if x != nil {
		return x.Jobs
	}
	return nil
}

type CreateJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Payload     string            `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Schedule    string            `protobuf:"bytes,3,opt,name=schedule,proto3" json:"schedule,omitempty"`
	Topic       *Topic            `protobuf:"bytes,4,opt,name=topic,proto3" json:"topic,omitempty"`
	Attributes  map[string]string `protobuf:"bytes,5,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Environment *string           `protobuf:"bytes,7,opt,name=environment,proto3,oneof" json:"environment,omitempty"`
	Tags        []string          `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *CreateJobRequest) Reset() {
	*x = CreateJobRequest{}
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobRequest) ProtoMessage() {}

func (x *CreateJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobRequest.ProtoReflect.Descriptor instead.
func (*CreateJobRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_topicpush_proto_rawDescGZIP(), []int{5}
}

func (x *CreateJobRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateJobRequest) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *CreateJobRequest) GetSchedule() string {
	if x != nil {
		return x.Schedule
	}
	return ""
}

func (x *CreateJobRequest) GetTopic() *Topic {
	if x != nil {
		return x.Topic
	}
	return nil
}

func (x *CreateJobRequest) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *CreateJobRequest) GetEnvironment() string {
	if x != nil && x.Environment != nil {
		return *x.Environment
	}
	return ""
}

func (x *CreateJobRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type CreateJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job *ScheduledJob `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *CreateJobResponse) Reset() {
	*x = CreateJobResponse{}
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobResponse) ProtoMessage() {}

func (x *CreateJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJobResponse.ProtoReflect.Descriptor instead.
func (*CreateJobResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_topicpush_proto_rawDescGZIP(), []int{6}
}

func (x *CreateJobResponse) GetJob() *ScheduledJob {
	if x != nil {
		return x.Job
	}
	return nil
}

type DeleteJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteJobRequest) Reset() {
	*x = DeleteJobRequest{}
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobRequest) ProtoMessage() {}

func (x *DeleteJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJobRequest.ProtoReflect.Descriptor instead.
func (*DeleteJobRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_topicpush_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteJobRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteJobResponse) Reset() {
	*x = DeleteJobResponse{}
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobResponse) ProtoMessage() {}

func (x *DeleteJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteJobResponse.ProtoReflect.Descriptor instead.
func (*DeleteJobResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_topicpush_proto_rawDescGZIP(), []int{8}
}

type GetJobByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetJobByNameRequest) Reset() {
	*x = GetJobByNameRequest{}
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetJobByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobByNameRequest) ProtoMessage() {}

func (x *GetJobByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobByNameRequest.ProtoReflect.Descriptor instead.
func (*GetJobByNameRequest) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_topicpush_proto_rawDescGZIP(), []int{9}
}

func (x *GetJobByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetJobByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job *ScheduledJob `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *GetJobByNameResponse) Reset() {
	*x = GetJobByNameResponse{}
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetJobByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobByNameResponse) ProtoMessage() {}

func (x *GetJobByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobByNameResponse.ProtoReflect.Descriptor instead.
func (*GetJobByNameResponse) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_topicpush_proto_rawDescGZIP(), []int{10}
}

func (x *GetJobByNameResponse) GetJob() *ScheduledJob {
	if x != nil {
		return x.Job
	}
	return nil
}

var File_chalk_server_v1_topicpush_proto protoreflect.FileDescriptor

var file_chalk_server_v1_topicpush_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xef, 0x02, 0x0a, 0x0c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x4a, 0x6f,
	0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x4d, 0x0a, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x2e, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69,
	0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0xf7, 0x02, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x12, 0x2c, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x51, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x12, 0x25, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x3d, 0x0a,
	0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0e, 0x0a, 0x0c,
	0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x44, 0x0a, 0x11,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2f, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a,
	0x6f, 0x62, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x45, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x6a, 0x6f, 0x62,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22, 0xe7, 0x02, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x51, 0x0a, 0x0a, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0b,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x65, 0x6e, 0x76, 0x69, 0x72,
	0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x44, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x03, 0x6a,
	0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x22, 0x22, 0x0a, 0x10,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x47, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x64, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x32, 0xde, 0x03, 0x0a, 0x10, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x50, 0x75, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57,
	0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x20, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x06, 0x80, 0x7d, 0x1b, 0x90, 0x02, 0x01, 0x12, 0x57, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4a, 0x6f, 0x62, 0x12, 0x21, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x80, 0x7d, 0x1b,
	0x12, 0x5a, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x21, 0x2e,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x80, 0x7d, 0x1b, 0x90, 0x02, 0x02, 0x12, 0x57, 0x0a, 0x09,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x21, 0x2e, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x03, 0x80, 0x7d, 0x1b, 0x12, 0x63, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4a, 0x6f, 0x62, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x06, 0x80, 0x7d, 0x1b, 0x90, 0x02, 0x02, 0x42, 0xbe, 0x01, 0x0a, 0x13, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x42, 0x0e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x70, 0x75, 0x73, 0x68, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d,
	0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72,
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
	file_chalk_server_v1_topicpush_proto_rawDescOnce sync.Once
	file_chalk_server_v1_topicpush_proto_rawDescData = file_chalk_server_v1_topicpush_proto_rawDesc
)

func file_chalk_server_v1_topicpush_proto_rawDescGZIP() []byte {
	file_chalk_server_v1_topicpush_proto_rawDescOnce.Do(func() {
		file_chalk_server_v1_topicpush_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_server_v1_topicpush_proto_rawDescData)
	})
	return file_chalk_server_v1_topicpush_proto_rawDescData
}

var file_chalk_server_v1_topicpush_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_chalk_server_v1_topicpush_proto_goTypes = []any{
	(*ScheduledJob)(nil),         // 0: chalk.server.v1.ScheduledJob
	(*UpdateJobRequest)(nil),     // 1: chalk.server.v1.UpdateJobRequest
	(*UpdateJobResponse)(nil),    // 2: chalk.server.v1.UpdateJobResponse
	(*ListJobsRequest)(nil),      // 3: chalk.server.v1.ListJobsRequest
	(*ListJobsResponse)(nil),     // 4: chalk.server.v1.ListJobsResponse
	(*CreateJobRequest)(nil),     // 5: chalk.server.v1.CreateJobRequest
	(*CreateJobResponse)(nil),    // 6: chalk.server.v1.CreateJobResponse
	(*DeleteJobRequest)(nil),     // 7: chalk.server.v1.DeleteJobRequest
	(*DeleteJobResponse)(nil),    // 8: chalk.server.v1.DeleteJobResponse
	(*GetJobByNameRequest)(nil),  // 9: chalk.server.v1.GetJobByNameRequest
	(*GetJobByNameResponse)(nil), // 10: chalk.server.v1.GetJobByNameResponse
	nil,                          // 11: chalk.server.v1.ScheduledJob.AttributesEntry
	nil,                          // 12: chalk.server.v1.UpdateJobRequest.AttributesEntry
	nil,                          // 13: chalk.server.v1.CreateJobRequest.AttributesEntry
	(*Topic)(nil),                // 14: chalk.server.v1.Topic
}
var file_chalk_server_v1_topicpush_proto_depIdxs = []int32{
	14, // 0: chalk.server.v1.ScheduledJob.topic:type_name -> chalk.server.v1.Topic
	11, // 1: chalk.server.v1.ScheduledJob.attributes:type_name -> chalk.server.v1.ScheduledJob.AttributesEntry
	14, // 2: chalk.server.v1.UpdateJobRequest.topic:type_name -> chalk.server.v1.Topic
	12, // 3: chalk.server.v1.UpdateJobRequest.attributes:type_name -> chalk.server.v1.UpdateJobRequest.AttributesEntry
	0,  // 4: chalk.server.v1.UpdateJobResponse.job:type_name -> chalk.server.v1.ScheduledJob
	0,  // 5: chalk.server.v1.ListJobsResponse.jobs:type_name -> chalk.server.v1.ScheduledJob
	14, // 6: chalk.server.v1.CreateJobRequest.topic:type_name -> chalk.server.v1.Topic
	13, // 7: chalk.server.v1.CreateJobRequest.attributes:type_name -> chalk.server.v1.CreateJobRequest.AttributesEntry
	0,  // 8: chalk.server.v1.CreateJobResponse.job:type_name -> chalk.server.v1.ScheduledJob
	0,  // 9: chalk.server.v1.GetJobByNameResponse.job:type_name -> chalk.server.v1.ScheduledJob
	3,  // 10: chalk.server.v1.TopicPushService.ListJobs:input_type -> chalk.server.v1.ListJobsRequest
	5,  // 11: chalk.server.v1.TopicPushService.CreateJob:input_type -> chalk.server.v1.CreateJobRequest
	1,  // 12: chalk.server.v1.TopicPushService.UpdateJob:input_type -> chalk.server.v1.UpdateJobRequest
	7,  // 13: chalk.server.v1.TopicPushService.DeleteJob:input_type -> chalk.server.v1.DeleteJobRequest
	9,  // 14: chalk.server.v1.TopicPushService.GetJobByName:input_type -> chalk.server.v1.GetJobByNameRequest
	4,  // 15: chalk.server.v1.TopicPushService.ListJobs:output_type -> chalk.server.v1.ListJobsResponse
	6,  // 16: chalk.server.v1.TopicPushService.CreateJob:output_type -> chalk.server.v1.CreateJobResponse
	2,  // 17: chalk.server.v1.TopicPushService.UpdateJob:output_type -> chalk.server.v1.UpdateJobResponse
	8,  // 18: chalk.server.v1.TopicPushService.DeleteJob:output_type -> chalk.server.v1.DeleteJobResponse
	10, // 19: chalk.server.v1.TopicPushService.GetJobByName:output_type -> chalk.server.v1.GetJobByNameResponse
	15, // [15:20] is the sub-list for method output_type
	10, // [10:15] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_chalk_server_v1_topicpush_proto_init() }
func file_chalk_server_v1_topicpush_proto_init() {
	if File_chalk_server_v1_topicpush_proto != nil {
		return
	}
	file_chalk_server_v1_topic_proto_init()
	file_chalk_server_v1_topicpush_proto_msgTypes[0].OneofWrappers = []any{}
	file_chalk_server_v1_topicpush_proto_msgTypes[1].OneofWrappers = []any{}
	file_chalk_server_v1_topicpush_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_server_v1_topicpush_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chalk_server_v1_topicpush_proto_goTypes,
		DependencyIndexes: file_chalk_server_v1_topicpush_proto_depIdxs,
		MessageInfos:      file_chalk_server_v1_topicpush_proto_msgTypes,
	}.Build()
	File_chalk_server_v1_topicpush_proto = out.File
	file_chalk_server_v1_topicpush_proto_rawDesc = nil
	file_chalk_server_v1_topicpush_proto_goTypes = nil
	file_chalk_server_v1_topicpush_proto_depIdxs = nil
}
