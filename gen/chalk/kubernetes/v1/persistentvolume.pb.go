// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: chalk/kubernetes/v1/persistentvolume.proto

package kubernetesv1

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

type ChalkKubernetesPersistentVolume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec    *ChalkKubernetesPersistentVolumeSpec    `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	Metrics *ChalkKubernetesPersistentVolumeMetrics `protobuf:"bytes,2,opt,name=metrics,proto3" json:"metrics,omitempty"`
}

func (x *ChalkKubernetesPersistentVolume) Reset() {
	*x = ChalkKubernetesPersistentVolume{}
	mi := &file_chalk_kubernetes_v1_persistentvolume_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChalkKubernetesPersistentVolume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChalkKubernetesPersistentVolume) ProtoMessage() {}

func (x *ChalkKubernetesPersistentVolume) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_kubernetes_v1_persistentvolume_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChalkKubernetesPersistentVolume.ProtoReflect.Descriptor instead.
func (*ChalkKubernetesPersistentVolume) Descriptor() ([]byte, []int) {
	return file_chalk_kubernetes_v1_persistentvolume_proto_rawDescGZIP(), []int{0}
}

func (x *ChalkKubernetesPersistentVolume) GetSpec() *ChalkKubernetesPersistentVolumeSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *ChalkKubernetesPersistentVolume) GetMetrics() *ChalkKubernetesPersistentVolumeMetrics {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type ChalkKubernetesPersistentVolumeSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageClass  string   `protobuf:"bytes,1,opt,name=storage_class,json=storageClass,proto3" json:"storage_class,omitempty"`
	Name          string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AccessModes   []string `protobuf:"bytes,3,rep,name=access_modes,json=accessModes,proto3" json:"access_modes,omitempty"`
	Capacity      string   `protobuf:"bytes,4,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Status        string   `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	ReclaimPolicy string   `protobuf:"bytes,6,opt,name=reclaim_policy,json=reclaimPolicy,proto3" json:"reclaim_policy,omitempty"`
	Claim         string   `protobuf:"bytes,7,opt,name=claim,proto3" json:"claim,omitempty"`
}

func (x *ChalkKubernetesPersistentVolumeSpec) Reset() {
	*x = ChalkKubernetesPersistentVolumeSpec{}
	mi := &file_chalk_kubernetes_v1_persistentvolume_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChalkKubernetesPersistentVolumeSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChalkKubernetesPersistentVolumeSpec) ProtoMessage() {}

func (x *ChalkKubernetesPersistentVolumeSpec) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_kubernetes_v1_persistentvolume_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChalkKubernetesPersistentVolumeSpec.ProtoReflect.Descriptor instead.
func (*ChalkKubernetesPersistentVolumeSpec) Descriptor() ([]byte, []int) {
	return file_chalk_kubernetes_v1_persistentvolume_proto_rawDescGZIP(), []int{1}
}

func (x *ChalkKubernetesPersistentVolumeSpec) GetStorageClass() string {
	if x != nil {
		return x.StorageClass
	}
	return ""
}

func (x *ChalkKubernetesPersistentVolumeSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChalkKubernetesPersistentVolumeSpec) GetAccessModes() []string {
	if x != nil {
		return x.AccessModes
	}
	return nil
}

func (x *ChalkKubernetesPersistentVolumeSpec) GetCapacity() string {
	if x != nil {
		return x.Capacity
	}
	return ""
}

func (x *ChalkKubernetesPersistentVolumeSpec) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ChalkKubernetesPersistentVolumeSpec) GetReclaimPolicy() string {
	if x != nil {
		return x.ReclaimPolicy
	}
	return ""
}

func (x *ChalkKubernetesPersistentVolumeSpec) GetClaim() string {
	if x != nil {
		return x.Claim
	}
	return ""
}

type ChalkKubernetesPersistentVolumeMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CapacityBytes  float64 `protobuf:"fixed64,1,opt,name=capacity_bytes,json=capacityBytes,proto3" json:"capacity_bytes,omitempty"` // Double to match prometheus type
	UsedBytes      float64 `protobuf:"fixed64,2,opt,name=used_bytes,json=usedBytes,proto3" json:"used_bytes,omitempty"`
	AvailableBytes float64 `protobuf:"fixed64,3,opt,name=available_bytes,json=availableBytes,proto3" json:"available_bytes,omitempty"`
}

func (x *ChalkKubernetesPersistentVolumeMetrics) Reset() {
	*x = ChalkKubernetesPersistentVolumeMetrics{}
	mi := &file_chalk_kubernetes_v1_persistentvolume_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChalkKubernetesPersistentVolumeMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChalkKubernetesPersistentVolumeMetrics) ProtoMessage() {}

func (x *ChalkKubernetesPersistentVolumeMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_kubernetes_v1_persistentvolume_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChalkKubernetesPersistentVolumeMetrics.ProtoReflect.Descriptor instead.
func (*ChalkKubernetesPersistentVolumeMetrics) Descriptor() ([]byte, []int) {
	return file_chalk_kubernetes_v1_persistentvolume_proto_rawDescGZIP(), []int{2}
}

func (x *ChalkKubernetesPersistentVolumeMetrics) GetCapacityBytes() float64 {
	if x != nil {
		return x.CapacityBytes
	}
	return 0
}

func (x *ChalkKubernetesPersistentVolumeMetrics) GetUsedBytes() float64 {
	if x != nil {
		return x.UsedBytes
	}
	return 0
}

func (x *ChalkKubernetesPersistentVolumeMetrics) GetAvailableBytes() float64 {
	if x != nil {
		return x.AvailableBytes
	}
	return 0
}

var File_chalk_kubernetes_v1_persistentvolume_proto protoreflect.FileDescriptor

var file_chalk_kubernetes_v1_persistentvolume_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x22, 0xc6, 0x01, 0x0a, 0x1f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x12, 0x55, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6b,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0xf2, 0x01, 0x0a, 0x23, 0x43,
	0x68, 0x61, 0x6c, 0x6b, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x50, 0x65,
	0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x6c,
	0x61, 0x69, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x61,
	0x69, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x22,
	0x97, 0x01, 0x0a, 0x26, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61,
	0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0d, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x75, 0x73, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x12, 0x27, 0x0a, 0x0f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x42, 0xe1, 0x01, 0x0a, 0x17, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x15, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x74, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x43, 0x4b, 0x58, 0xaa, 0x02, 0x13, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x13,
	0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_kubernetes_v1_persistentvolume_proto_rawDescOnce sync.Once
	file_chalk_kubernetes_v1_persistentvolume_proto_rawDescData = file_chalk_kubernetes_v1_persistentvolume_proto_rawDesc
)

func file_chalk_kubernetes_v1_persistentvolume_proto_rawDescGZIP() []byte {
	file_chalk_kubernetes_v1_persistentvolume_proto_rawDescOnce.Do(func() {
		file_chalk_kubernetes_v1_persistentvolume_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_kubernetes_v1_persistentvolume_proto_rawDescData)
	})
	return file_chalk_kubernetes_v1_persistentvolume_proto_rawDescData
}

var file_chalk_kubernetes_v1_persistentvolume_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chalk_kubernetes_v1_persistentvolume_proto_goTypes = []any{
	(*ChalkKubernetesPersistentVolume)(nil),        // 0: chalk.kubernetes.v1.ChalkKubernetesPersistentVolume
	(*ChalkKubernetesPersistentVolumeSpec)(nil),    // 1: chalk.kubernetes.v1.ChalkKubernetesPersistentVolumeSpec
	(*ChalkKubernetesPersistentVolumeMetrics)(nil), // 2: chalk.kubernetes.v1.ChalkKubernetesPersistentVolumeMetrics
}
var file_chalk_kubernetes_v1_persistentvolume_proto_depIdxs = []int32{
	1, // 0: chalk.kubernetes.v1.ChalkKubernetesPersistentVolume.spec:type_name -> chalk.kubernetes.v1.ChalkKubernetesPersistentVolumeSpec
	2, // 1: chalk.kubernetes.v1.ChalkKubernetesPersistentVolume.metrics:type_name -> chalk.kubernetes.v1.ChalkKubernetesPersistentVolumeMetrics
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_chalk_kubernetes_v1_persistentvolume_proto_init() }
func file_chalk_kubernetes_v1_persistentvolume_proto_init() {
	if File_chalk_kubernetes_v1_persistentvolume_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_kubernetes_v1_persistentvolume_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chalk_kubernetes_v1_persistentvolume_proto_goTypes,
		DependencyIndexes: file_chalk_kubernetes_v1_persistentvolume_proto_depIdxs,
		MessageInfos:      file_chalk_kubernetes_v1_persistentvolume_proto_msgTypes,
	}.Build()
	File_chalk_kubernetes_v1_persistentvolume_proto = out.File
	file_chalk_kubernetes_v1_persistentvolume_proto_rawDesc = nil
	file_chalk_kubernetes_v1_persistentvolume_proto_goTypes = nil
	file_chalk_kubernetes_v1_persistentvolume_proto_depIdxs = nil
}
