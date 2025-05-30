// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: chalk/server/v1/deployment.proto

package serverv1

import (
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

type DeploymentStatus int32

const (
	DeploymentStatus_DEPLOYMENT_STATUS_UNSPECIFIED     DeploymentStatus = 0
	DeploymentStatus_DEPLOYMENT_STATUS_UNKNOWN         DeploymentStatus = 1
	DeploymentStatus_DEPLOYMENT_STATUS_PENDING         DeploymentStatus = 2
	DeploymentStatus_DEPLOYMENT_STATUS_QUEUED          DeploymentStatus = 3
	DeploymentStatus_DEPLOYMENT_STATUS_WORKING         DeploymentStatus = 4
	DeploymentStatus_DEPLOYMENT_STATUS_SUCCESS         DeploymentStatus = 5
	DeploymentStatus_DEPLOYMENT_STATUS_FAILURE         DeploymentStatus = 6
	DeploymentStatus_DEPLOYMENT_STATUS_INTERNAL_ERROR  DeploymentStatus = 7
	DeploymentStatus_DEPLOYMENT_STATUS_TIMEOUT         DeploymentStatus = 8
	DeploymentStatus_DEPLOYMENT_STATUS_CANCELLED       DeploymentStatus = 9
	DeploymentStatus_DEPLOYMENT_STATUS_EXPIRED         DeploymentStatus = 10
	DeploymentStatus_DEPLOYMENT_STATUS_BOOT_ERRORS     DeploymentStatus = 11
	DeploymentStatus_DEPLOYMENT_STATUS_AWAITING_SOURCE DeploymentStatus = 12
	DeploymentStatus_DEPLOYMENT_STATUS_DEPLOYING       DeploymentStatus = 13
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
		12: "DEPLOYMENT_STATUS_AWAITING_SOURCE",
		13: "DEPLOYMENT_STATUS_DEPLOYING",
	}
	DeploymentStatus_value = map[string]int32{
		"DEPLOYMENT_STATUS_UNSPECIFIED":     0,
		"DEPLOYMENT_STATUS_UNKNOWN":         1,
		"DEPLOYMENT_STATUS_PENDING":         2,
		"DEPLOYMENT_STATUS_QUEUED":          3,
		"DEPLOYMENT_STATUS_WORKING":         4,
		"DEPLOYMENT_STATUS_SUCCESS":         5,
		"DEPLOYMENT_STATUS_FAILURE":         6,
		"DEPLOYMENT_STATUS_INTERNAL_ERROR":  7,
		"DEPLOYMENT_STATUS_TIMEOUT":         8,
		"DEPLOYMENT_STATUS_CANCELLED":       9,
		"DEPLOYMENT_STATUS_EXPIRED":         10,
		"DEPLOYMENT_STATUS_BOOT_ERRORS":     11,
		"DEPLOYMENT_STATUS_AWAITING_SOURCE": 12,
		"DEPLOYMENT_STATUS_DEPLOYING":       13,
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

type DeploymentProfilingMode int32

const (
	DeploymentProfilingMode_DEPLOYMENT_PROFILING_MODE_UNSPECIFIED DeploymentProfilingMode = 0
	DeploymentProfilingMode_DEPLOYMENT_PROFILING_MODE_NONE        DeploymentProfilingMode = 1
	DeploymentProfilingMode_DEPLOYMENT_PROFILING_MODE_O2          DeploymentProfilingMode = 2
)

// Enum value maps for DeploymentProfilingMode.
var (
	DeploymentProfilingMode_name = map[int32]string{
		0: "DEPLOYMENT_PROFILING_MODE_UNSPECIFIED",
		1: "DEPLOYMENT_PROFILING_MODE_NONE",
		2: "DEPLOYMENT_PROFILING_MODE_O2",
	}
	DeploymentProfilingMode_value = map[string]int32{
		"DEPLOYMENT_PROFILING_MODE_UNSPECIFIED": 0,
		"DEPLOYMENT_PROFILING_MODE_NONE":        1,
		"DEPLOYMENT_PROFILING_MODE_O2":          2,
	}
)

func (x DeploymentProfilingMode) Enum() *DeploymentProfilingMode {
	p := new(DeploymentProfilingMode)
	*p = x
	return p
}

func (x DeploymentProfilingMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeploymentProfilingMode) Descriptor() protoreflect.EnumDescriptor {
	return file_chalk_server_v1_deployment_proto_enumTypes[1].Descriptor()
}

func (DeploymentProfilingMode) Type() protoreflect.EnumType {
	return &file_chalk_server_v1_deployment_proto_enumTypes[1]
}

func (x DeploymentProfilingMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeploymentProfilingMode.Descriptor instead.
func (DeploymentProfilingMode) EnumDescriptor() ([]byte, []int) {
	return file_chalk_server_v1_deployment_proto_rawDescGZIP(), []int{1}
}

type InstanceSizing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinInstances *uint32 `protobuf:"varint,1,opt,name=min_instances,json=minInstances,proto3,oneof" json:"min_instances,omitempty"`
	MaxInstances *uint32 `protobuf:"varint,2,opt,name=max_instances,json=maxInstances,proto3,oneof" json:"max_instances,omitempty"`
}

func (x *InstanceSizing) Reset() {
	*x = InstanceSizing{}
	mi := &file_chalk_server_v1_deployment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InstanceSizing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceSizing) ProtoMessage() {}

func (x *InstanceSizing) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_deployment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceSizing.ProtoReflect.Descriptor instead.
func (*InstanceSizing) Descriptor() ([]byte, []int) {
	return file_chalk_server_v1_deployment_proto_rawDescGZIP(), []int{0}
}

func (x *InstanceSizing) GetMinInstances() uint32 {
	if x != nil && x.MinInstances != nil {
		return *x.MinInstances
	}
	return 0
}

func (x *InstanceSizing) GetMaxInstances() uint32 {
	if x != nil && x.MaxInstances != nil {
		return *x.MaxInstances
	}
	return 0
}

type Deployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EnvironmentId         string                   `protobuf:"bytes,2,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	Status                DeploymentStatus         `protobuf:"varint,3,opt,name=status,proto3,enum=chalk.server.v1.DeploymentStatus" json:"status,omitempty"`
	DeploymentTags        []string                 `protobuf:"bytes,4,rep,name=deployment_tags,json=deploymentTags,proto3" json:"deployment_tags,omitempty"`
	CloudBuildId          string                   `protobuf:"bytes,5,opt,name=cloud_build_id,json=cloudBuildId,proto3" json:"cloud_build_id,omitempty"`
	TriggeredBy           string                   `protobuf:"bytes,6,opt,name=triggered_by,json=triggeredBy,proto3" json:"triggered_by,omitempty"`
	RequirementsFilepath  *string                  `protobuf:"bytes,7,opt,name=requirements_filepath,json=requirementsFilepath,proto3,oneof" json:"requirements_filepath,omitempty"`
	DockerfileFilepath    *string                  `protobuf:"bytes,8,opt,name=dockerfile_filepath,json=dockerfileFilepath,proto3,oneof" json:"dockerfile_filepath,omitempty"`
	Runtime               *string                  `protobuf:"bytes,9,opt,name=runtime,proto3,oneof" json:"runtime,omitempty"`
	ChalkpyVersion        string                   `protobuf:"bytes,10,opt,name=chalkpy_version,json=chalkpyVersion,proto3" json:"chalkpy_version,omitempty"`
	RawDependencyHash     string                   `protobuf:"bytes,11,opt,name=raw_dependency_hash,json=rawDependencyHash,proto3" json:"raw_dependency_hash,omitempty"`
	FinalDependencyHash   *string                  `protobuf:"bytes,12,opt,name=final_dependency_hash,json=finalDependencyHash,proto3,oneof" json:"final_dependency_hash,omitempty"`
	IsPreviewDeployment   *bool                    `protobuf:"varint,13,opt,name=is_preview_deployment,json=isPreviewDeployment,proto3,oneof" json:"is_preview_deployment,omitempty"`
	CreatedAt             *timestamppb.Timestamp   `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt             *timestamppb.Timestamp   `protobuf:"bytes,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	GitCommit             string                   `protobuf:"bytes,16,opt,name=git_commit,json=gitCommit,proto3" json:"git_commit,omitempty"`
	GitPr                 string                   `protobuf:"bytes,17,opt,name=git_pr,json=gitPr,proto3" json:"git_pr,omitempty"`
	GitBranch             string                   `protobuf:"bytes,18,opt,name=git_branch,json=gitBranch,proto3" json:"git_branch,omitempty"`
	GitAuthorEmail        string                   `protobuf:"bytes,19,opt,name=git_author_email,json=gitAuthorEmail,proto3" json:"git_author_email,omitempty"`
	Branch                string                   `protobuf:"bytes,20,opt,name=branch,proto3" json:"branch,omitempty"`
	ProjectSettings       string                   `protobuf:"bytes,21,opt,name=project_settings,json=projectSettings,proto3" json:"project_settings,omitempty"`
	RequirementsFiles     *string                  `protobuf:"bytes,22,opt,name=requirements_files,json=requirementsFiles,proto3,oneof" json:"requirements_files,omitempty"`
	GitTag                string                   `protobuf:"bytes,23,opt,name=git_tag,json=gitTag,proto3" json:"git_tag,omitempty"`
	BaseImageSha          string                   `protobuf:"bytes,24,opt,name=base_image_sha,json=baseImageSha,proto3" json:"base_image_sha,omitempty"`
	StatusChangedAt       *timestamppb.Timestamp   `protobuf:"bytes,25,opt,name=status_changed_at,json=statusChangedAt,proto3" json:"status_changed_at,omitempty"`
	PinnedPlatformVersion *string                  `protobuf:"bytes,26,opt,name=pinned_platform_version,json=pinnedPlatformVersion,proto3,oneof" json:"pinned_platform_version,omitempty"`
	PreviewDeploymentTag  *string                  `protobuf:"bytes,27,opt,name=preview_deployment_tag,json=previewDeploymentTag,proto3,oneof" json:"preview_deployment_tag,omitempty"`
	ProfilingMode         *DeploymentProfilingMode `protobuf:"varint,28,opt,name=profiling_mode,json=profilingMode,proto3,enum=chalk.server.v1.DeploymentProfilingMode,oneof" json:"profiling_mode,omitempty"`
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	mi := &file_chalk_server_v1_deployment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_server_v1_deployment_proto_msgTypes[1]
	if x != nil {
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
	return file_chalk_server_v1_deployment_proto_rawDescGZIP(), []int{1}
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

func (x *Deployment) GetDeploymentTags() []string {
	if x != nil {
		return x.DeploymentTags
	}
	return nil
}

func (x *Deployment) GetCloudBuildId() string {
	if x != nil {
		return x.CloudBuildId
	}
	return ""
}

func (x *Deployment) GetTriggeredBy() string {
	if x != nil {
		return x.TriggeredBy
	}
	return ""
}

func (x *Deployment) GetRequirementsFilepath() string {
	if x != nil && x.RequirementsFilepath != nil {
		return *x.RequirementsFilepath
	}
	return ""
}

func (x *Deployment) GetDockerfileFilepath() string {
	if x != nil && x.DockerfileFilepath != nil {
		return *x.DockerfileFilepath
	}
	return ""
}

func (x *Deployment) GetRuntime() string {
	if x != nil && x.Runtime != nil {
		return *x.Runtime
	}
	return ""
}

func (x *Deployment) GetChalkpyVersion() string {
	if x != nil {
		return x.ChalkpyVersion
	}
	return ""
}

func (x *Deployment) GetRawDependencyHash() string {
	if x != nil {
		return x.RawDependencyHash
	}
	return ""
}

func (x *Deployment) GetFinalDependencyHash() string {
	if x != nil && x.FinalDependencyHash != nil {
		return *x.FinalDependencyHash
	}
	return ""
}

func (x *Deployment) GetIsPreviewDeployment() bool {
	if x != nil && x.IsPreviewDeployment != nil {
		return *x.IsPreviewDeployment
	}
	return false
}

func (x *Deployment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Deployment) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Deployment) GetGitCommit() string {
	if x != nil {
		return x.GitCommit
	}
	return ""
}

func (x *Deployment) GetGitPr() string {
	if x != nil {
		return x.GitPr
	}
	return ""
}

func (x *Deployment) GetGitBranch() string {
	if x != nil {
		return x.GitBranch
	}
	return ""
}

func (x *Deployment) GetGitAuthorEmail() string {
	if x != nil {
		return x.GitAuthorEmail
	}
	return ""
}

func (x *Deployment) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *Deployment) GetProjectSettings() string {
	if x != nil {
		return x.ProjectSettings
	}
	return ""
}

func (x *Deployment) GetRequirementsFiles() string {
	if x != nil && x.RequirementsFiles != nil {
		return *x.RequirementsFiles
	}
	return ""
}

func (x *Deployment) GetGitTag() string {
	if x != nil {
		return x.GitTag
	}
	return ""
}

func (x *Deployment) GetBaseImageSha() string {
	if x != nil {
		return x.BaseImageSha
	}
	return ""
}

func (x *Deployment) GetStatusChangedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StatusChangedAt
	}
	return nil
}

func (x *Deployment) GetPinnedPlatformVersion() string {
	if x != nil && x.PinnedPlatformVersion != nil {
		return *x.PinnedPlatformVersion
	}
	return ""
}

func (x *Deployment) GetPreviewDeploymentTag() string {
	if x != nil && x.PreviewDeploymentTag != nil {
		return *x.PreviewDeploymentTag
	}
	return ""
}

func (x *Deployment) GetProfilingMode() DeploymentProfilingMode {
	if x != nil && x.ProfilingMode != nil {
		return *x.ProfilingMode
	}
	return DeploymentProfilingMode_DEPLOYMENT_PROFILING_MODE_UNSPECIFIED
}

var File_chalk_server_v1_deployment_proto protoreflect.FileDescriptor

var file_chalk_server_v1_deployment_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x53, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x12, 0x28, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x5f, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00,
	0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x88, 0x01,
	0x01, 0x12, 0x28, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f,
	0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x42, 0x10, 0x0a,
	0x0e, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x22,
	0xde, 0x0b, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74,
	0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x67, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x5f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x65, 0x64,
	0x42, 0x79, 0x12, 0x38, 0x0a, 0x15, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x14, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x46, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x13,
	0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x12, 0x64, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x88,
	0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x70, 0x79, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x70, 0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x61,
	0x77, 0x5f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x61, 0x77, 0x44, 0x65, 0x70, 0x65,
	0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x48, 0x61, 0x73, 0x68, 0x12, 0x37, 0x0a, 0x15, 0x66, 0x69,
	0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x13, 0x66, 0x69, 0x6e,
	0x61, 0x6c, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x48, 0x61, 0x73, 0x68,
	0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x15, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x04, 0x52, 0x13, 0x69, 0x73, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x12, 0x15, 0x0a, 0x06, 0x67, 0x69, 0x74, 0x5f, 0x70, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x67, 0x69, 0x74, 0x50, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x69, 0x74, 0x5f,
	0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x69,
	0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x28, 0x0a, 0x10, 0x67, 0x69, 0x74, 0x5f, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x67, 0x69, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x32, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x05, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x69, 0x74, 0x5f,
	0x74, 0x61, 0x67, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x69, 0x74, 0x54, 0x61,
	0x67, 0x12, 0x24, 0x0a, 0x0e, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x68, 0x61, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x53, 0x68, 0x61, 0x12, 0x46, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x3b, 0x0a, 0x17, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x06, 0x52, 0x15, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x39, 0x0a, 0x16,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07, 0x52, 0x14,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x61, 0x67, 0x88, 0x01, 0x01, 0x12, 0x54, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x28, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x48, 0x08, 0x52, 0x0d, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x18, 0x0a,
	0x16, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x64, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x18, 0x0a, 0x16, 0x5f,
	0x66, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x69, 0x73, 0x5f, 0x70, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x42,
	0x15, 0x0a, 0x13, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x42, 0x1a, 0x0a, 0x18, 0x5f, 0x70, 0x69, 0x6e, 0x6e, 0x65,
	0x64, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x42, 0x19, 0x0a, 0x17, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x42, 0x11, 0x0a,
	0x0f, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x2a, 0xde, 0x03, 0x0a, 0x10, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x1d, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d,
	0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x45, 0x50, 0x4c,
	0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x45, 0x50, 0x4c, 0x4f,
	0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x51, 0x55, 0x45, 0x55,
	0x45, 0x44, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x49, 0x4e,
	0x47, 0x10, 0x04, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53,
	0x10, 0x05, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10,
	0x06, 0x12, 0x24, 0x0a, 0x20, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x07, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x45, 0x50, 0x4c, 0x4f,
	0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x4f, 0x55, 0x54, 0x10, 0x08, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x41, 0x4e, 0x43,
	0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x09, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x45, 0x50, 0x4c, 0x4f,
	0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x58, 0x50,
	0x49, 0x52, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x21, 0x0a, 0x1d, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x42, 0x4f, 0x4f, 0x54,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x53, 0x10, 0x0b, 0x12, 0x25, 0x0a, 0x21, 0x44, 0x45, 0x50,
	0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41,
	0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x0c,
	0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x49, 0x4e, 0x47, 0x10,
	0x0d, 0x2a, 0x8a, 0x01, 0x0a, 0x17, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a,
	0x25, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x46,
	0x49, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x22, 0x0a, 0x1e, 0x44, 0x45, 0x50, 0x4c,
	0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x49, 0x4e, 0x47,
	0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c,
	0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x50, 0x52, 0x4f, 0x46, 0x49,
	0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x4f, 0x32, 0x10, 0x02, 0x42, 0xbf,
	0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02, 0x0f, 0x43, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x43,
	0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1b, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x43,
	0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_chalk_server_v1_deployment_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_chalk_server_v1_deployment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chalk_server_v1_deployment_proto_goTypes = []any{
	(DeploymentStatus)(0),         // 0: chalk.server.v1.DeploymentStatus
	(DeploymentProfilingMode)(0),  // 1: chalk.server.v1.DeploymentProfilingMode
	(*InstanceSizing)(nil),        // 2: chalk.server.v1.InstanceSizing
	(*Deployment)(nil),            // 3: chalk.server.v1.Deployment
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_chalk_server_v1_deployment_proto_depIdxs = []int32{
	0, // 0: chalk.server.v1.Deployment.status:type_name -> chalk.server.v1.DeploymentStatus
	4, // 1: chalk.server.v1.Deployment.created_at:type_name -> google.protobuf.Timestamp
	4, // 2: chalk.server.v1.Deployment.updated_at:type_name -> google.protobuf.Timestamp
	4, // 3: chalk.server.v1.Deployment.status_changed_at:type_name -> google.protobuf.Timestamp
	1, // 4: chalk.server.v1.Deployment.profiling_mode:type_name -> chalk.server.v1.DeploymentProfilingMode
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_chalk_server_v1_deployment_proto_init() }
func file_chalk_server_v1_deployment_proto_init() {
	if File_chalk_server_v1_deployment_proto != nil {
		return
	}
	file_chalk_server_v1_deployment_proto_msgTypes[0].OneofWrappers = []any{}
	file_chalk_server_v1_deployment_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_server_v1_deployment_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
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
