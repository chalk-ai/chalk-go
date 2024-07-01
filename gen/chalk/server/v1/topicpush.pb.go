// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
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

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Payload  string `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Schedule string `protobuf:"bytes,4,opt,name=schedule,proto3" json:"schedule,omitempty"`
	Topic    *Topic `protobuf:"bytes,5,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *ScheduledJob) Reset() {
	*x = ScheduledJob{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_topicpush_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduledJob) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduledJob) ProtoMessage() {}

func (x *ScheduledJob) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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

type UpdateJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Payload  string `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Schedule string `protobuf:"bytes,4,opt,name=schedule,proto3" json:"schedule,omitempty"`
	Topic    *Topic `protobuf:"bytes,5,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *UpdateJobRequest) Reset() {
	*x = UpdateJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_topicpush_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobRequest) ProtoMessage() {}

func (x *UpdateJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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

type UpdateJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job *ScheduledJob `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *UpdateJobResponse) Reset() {
	*x = UpdateJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_topicpush_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateJobResponse) ProtoMessage() {}

func (x *UpdateJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_topicpush_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsRequest) ProtoMessage() {}

func (x *ListJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_topicpush_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListJobsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListJobsResponse) ProtoMessage() {}

func (x *ListJobsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
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

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Payload  string `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Schedule string `protobuf:"bytes,3,opt,name=schedule,proto3" json:"schedule,omitempty"`
	Topic    *Topic `protobuf:"bytes,4,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *CreateJobRequest) Reset() {
	*x = CreateJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_topicpush_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobRequest) ProtoMessage() {}

func (x *CreateJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
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

type CreateJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job *ScheduledJob `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
}

func (x *CreateJobResponse) Reset() {
	*x = CreateJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_topicpush_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJobResponse) ProtoMessage() {}

func (x *CreateJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
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
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_topicpush_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobRequest) ProtoMessage() {}

func (x *DeleteJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
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
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_topicpush_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteJobResponse) ProtoMessage() {}

func (x *DeleteJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
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
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_topicpush_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobByNameRequest) ProtoMessage() {}

func (x *GetJobByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
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
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_server_v1_topicpush_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobByNameResponse) ProtoMessage() {}

func (x *GetJobByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_topicpush_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
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
	0x22, 0x96, 0x01, 0x0a, 0x0c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x4a, 0x6f,
	0x62, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x9a, 0x01, 0x0a, 0x10, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52,
	0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x44, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x03, 0x6a,
	0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x64, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x22, 0x11, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x45, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x4a, 0x6f, 0x62,
	0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x22, 0x44, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x64, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x13, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x29, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x47, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x4a, 0x6f,
	0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62, 0x32, 0xde, 0x03, 0x0a, 0x10, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x50, 0x75, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x08, 0x4c,
	0x69, 0x73, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x20, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4a, 0x6f,
	0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x06, 0x80, 0x7d,
	0x1b, 0x90, 0x02, 0x01, 0x12, 0x57, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f,
	0x62, 0x12, 0x21, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x80, 0x7d, 0x1b, 0x12, 0x5a, 0x0a,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x21, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x06, 0x80, 0x7d, 0x1b, 0x90, 0x02, 0x02, 0x12, 0x57, 0x0a, 0x09, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4a, 0x6f, 0x62, 0x12, 0x21, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x03, 0x80,
	0x7d, 0x1b, 0x12, 0x63, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x24, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f,
	0x62, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x06, 0x80, 0x7d, 0x1b, 0x90, 0x02, 0x02, 0x42, 0xbe, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42,
	0x0e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x70, 0x75, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43,
	0x53, 0x58, 0xaa, 0x02, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1b, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_chalk_server_v1_topicpush_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_chalk_server_v1_topicpush_proto_goTypes = []interface{}{
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
	(*Topic)(nil),                // 11: chalk.server.v1.Topic
}
var file_chalk_server_v1_topicpush_proto_depIdxs = []int32{
	11, // 0: chalk.server.v1.ScheduledJob.topic:type_name -> chalk.server.v1.Topic
	11, // 1: chalk.server.v1.UpdateJobRequest.topic:type_name -> chalk.server.v1.Topic
	0,  // 2: chalk.server.v1.UpdateJobResponse.job:type_name -> chalk.server.v1.ScheduledJob
	0,  // 3: chalk.server.v1.ListJobsResponse.jobs:type_name -> chalk.server.v1.ScheduledJob
	11, // 4: chalk.server.v1.CreateJobRequest.topic:type_name -> chalk.server.v1.Topic
	0,  // 5: chalk.server.v1.CreateJobResponse.job:type_name -> chalk.server.v1.ScheduledJob
	0,  // 6: chalk.server.v1.GetJobByNameResponse.job:type_name -> chalk.server.v1.ScheduledJob
	3,  // 7: chalk.server.v1.TopicPushService.ListJobs:input_type -> chalk.server.v1.ListJobsRequest
	5,  // 8: chalk.server.v1.TopicPushService.CreateJob:input_type -> chalk.server.v1.CreateJobRequest
	1,  // 9: chalk.server.v1.TopicPushService.UpdateJob:input_type -> chalk.server.v1.UpdateJobRequest
	7,  // 10: chalk.server.v1.TopicPushService.DeleteJob:input_type -> chalk.server.v1.DeleteJobRequest
	9,  // 11: chalk.server.v1.TopicPushService.GetJobByName:input_type -> chalk.server.v1.GetJobByNameRequest
	4,  // 12: chalk.server.v1.TopicPushService.ListJobs:output_type -> chalk.server.v1.ListJobsResponse
	6,  // 13: chalk.server.v1.TopicPushService.CreateJob:output_type -> chalk.server.v1.CreateJobResponse
	2,  // 14: chalk.server.v1.TopicPushService.UpdateJob:output_type -> chalk.server.v1.UpdateJobResponse
	8,  // 15: chalk.server.v1.TopicPushService.DeleteJob:output_type -> chalk.server.v1.DeleteJobResponse
	10, // 16: chalk.server.v1.TopicPushService.GetJobByName:output_type -> chalk.server.v1.GetJobByNameResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_chalk_server_v1_topicpush_proto_init() }
func file_chalk_server_v1_topicpush_proto_init() {
	if File_chalk_server_v1_topicpush_proto != nil {
		return
	}
	file_chalk_server_v1_topic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chalk_server_v1_topicpush_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduledJob); i {
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
		file_chalk_server_v1_topicpush_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJobRequest); i {
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
		file_chalk_server_v1_topicpush_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateJobResponse); i {
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
		file_chalk_server_v1_topicpush_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJobsRequest); i {
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
		file_chalk_server_v1_topicpush_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListJobsResponse); i {
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
		file_chalk_server_v1_topicpush_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJobRequest); i {
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
		file_chalk_server_v1_topicpush_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJobResponse); i {
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
		file_chalk_server_v1_topicpush_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteJobRequest); i {
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
		file_chalk_server_v1_topicpush_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteJobResponse); i {
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
		file_chalk_server_v1_topicpush_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobByNameRequest); i {
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
		file_chalk_server_v1_topicpush_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobByNameResponse); i {
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
			RawDescriptor: file_chalk_server_v1_topicpush_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
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
