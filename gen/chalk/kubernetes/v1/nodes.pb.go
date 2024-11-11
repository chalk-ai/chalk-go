// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: chalk/kubernetes/v1/nodes.proto

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

// InstanceUsage is a message that represents the usage of a single instance.
type KubernetesNodeData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Chalk team name that incurred the usage.
	Team string `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	// node.Name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// node.UID
	Uid string `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	// node.kubernetes.io/instance-type
	InstanceType string `protobuf:"bytes,4,opt,name=instance_type,json=instanceType,proto3" json:"instance_type,omitempty"`
	// topology.kubernetes.io/region
	Region string `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	// topology.kubernetes.io/zone
	Zone string `protobuf:"bytes,6,opt,name=zone,proto3" json:"zone,omitempty"`
	// The time that the instance was created.
	// node.CreationTimestamp.Unix()
	CreationTimestamp int64 `protobuf:"varint,8,opt,name=creation_timestamp,json=creationTimestamp,proto3" json:"creation_timestamp,omitempty"`
	// node.DeletionTimestamp.Unix()
	// The time that the instance was deleted. May be 0 if the instance is still running.
	DeletionTimestamp int64 `protobuf:"varint,9,opt,name=deletion_timestamp,json=deletionTimestamp,proto3" json:"deletion_timestamp,omitempty"`
	// The time that we polled the instance for usage.
	ObservedTimestamp int64 `protobuf:"varint,10,opt,name=observed_timestamp,json=observedTimestamp,proto3" json:"observed_timestamp,omitempty"`
	// node.Labels
	Labels map[string]string `protobuf:"bytes,11,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// node.Annotations
	Annotations   map[string]string `protobuf:"bytes,12,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MachineId     string            `protobuf:"bytes,13,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	SystemUuid    string            `protobuf:"bytes,14,opt,name=system_uuid,json=systemUuid,proto3" json:"system_uuid,omitempty"`
	BootId        string            `protobuf:"bytes,15,opt,name=boot_id,json=bootId,proto3" json:"boot_id,omitempty"`
	Unschedulable bool              `protobuf:"varint,16,opt,name=unschedulable,proto3" json:"unschedulable,omitempty"`
	Namespace     string            `protobuf:"bytes,17,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// container.googleapis.com/instance_id
	InstanceId        string `protobuf:"bytes,18,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	Cluster           string `protobuf:"bytes,19,opt,name=cluster,proto3" json:"cluster,omitempty"`
	TotalCpu          string `protobuf:"bytes,20,opt,name=total_cpu,json=totalCpu,proto3" json:"total_cpu,omitempty"`
	TotalMemory       string `protobuf:"bytes,21,opt,name=total_memory,json=totalMemory,proto3" json:"total_memory,omitempty"`
	AllocatableCpu    string `protobuf:"bytes,22,opt,name=allocatable_cpu,json=allocatableCpu,proto3" json:"allocatable_cpu,omitempty"`
	AllocatableMemory string `protobuf:"bytes,23,opt,name=allocatable_memory,json=allocatableMemory,proto3" json:"allocatable_memory,omitempty"`
}

func (x *KubernetesNodeData) Reset() {
	*x = KubernetesNodeData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_kubernetes_v1_nodes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesNodeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesNodeData) ProtoMessage() {}

func (x *KubernetesNodeData) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_kubernetes_v1_nodes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesNodeData.ProtoReflect.Descriptor instead.
func (*KubernetesNodeData) Descriptor() ([]byte, []int) {
	return file_chalk_kubernetes_v1_nodes_proto_rawDescGZIP(), []int{0}
}

func (x *KubernetesNodeData) GetTeam() string {
	if x != nil {
		return x.Team
	}
	return ""
}

func (x *KubernetesNodeData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KubernetesNodeData) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *KubernetesNodeData) GetInstanceType() string {
	if x != nil {
		return x.InstanceType
	}
	return ""
}

func (x *KubernetesNodeData) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *KubernetesNodeData) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *KubernetesNodeData) GetCreationTimestamp() int64 {
	if x != nil {
		return x.CreationTimestamp
	}
	return 0
}

func (x *KubernetesNodeData) GetDeletionTimestamp() int64 {
	if x != nil {
		return x.DeletionTimestamp
	}
	return 0
}

func (x *KubernetesNodeData) GetObservedTimestamp() int64 {
	if x != nil {
		return x.ObservedTimestamp
	}
	return 0
}

func (x *KubernetesNodeData) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *KubernetesNodeData) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *KubernetesNodeData) GetMachineId() string {
	if x != nil {
		return x.MachineId
	}
	return ""
}

func (x *KubernetesNodeData) GetSystemUuid() string {
	if x != nil {
		return x.SystemUuid
	}
	return ""
}

func (x *KubernetesNodeData) GetBootId() string {
	if x != nil {
		return x.BootId
	}
	return ""
}

func (x *KubernetesNodeData) GetUnschedulable() bool {
	if x != nil {
		return x.Unschedulable
	}
	return false
}

func (x *KubernetesNodeData) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *KubernetesNodeData) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *KubernetesNodeData) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *KubernetesNodeData) GetTotalCpu() string {
	if x != nil {
		return x.TotalCpu
	}
	return ""
}

func (x *KubernetesNodeData) GetTotalMemory() string {
	if x != nil {
		return x.TotalMemory
	}
	return ""
}

func (x *KubernetesNodeData) GetAllocatableCpu() string {
	if x != nil {
		return x.AllocatableCpu
	}
	return ""
}

func (x *KubernetesNodeData) GetAllocatableMemory() string {
	if x != nil {
		return x.AllocatableMemory
	}
	return ""
}

var File_chalk_kubernetes_v1_nodes_proto protoreflect.FileDescriptor

var file_chalk_kubernetes_v1_nodes_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x22, 0xc0, 0x07, 0x0a, 0x12, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x61,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2d, 0x0a, 0x12, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x11, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2d, 0x0a, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x11, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x4b, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x12, 0x5a, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x75, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x62, 0x6f, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x62, 0x6f, 0x6f, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x75, 0x6e, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x75, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x63, 0x70, 0x75, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x70, 0x75, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x70, 0x75, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x70, 0x75, 0x12,
	0x2d, 0x0a, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x6c, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x1a, 0x39,
	0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0xd6, 0x01, 0x0a, 0x17, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x67,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x4b, 0x58, 0xaa, 0x02, 0x13, 0x43,
	0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x13, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1f, 0x43, 0x68, 0x61, 0x6c, 0x6b,
	0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x43, 0x68, 0x61,
	0x6c, 0x6b, 0x3a, 0x3a, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_kubernetes_v1_nodes_proto_rawDescOnce sync.Once
	file_chalk_kubernetes_v1_nodes_proto_rawDescData = file_chalk_kubernetes_v1_nodes_proto_rawDesc
)

func file_chalk_kubernetes_v1_nodes_proto_rawDescGZIP() []byte {
	file_chalk_kubernetes_v1_nodes_proto_rawDescOnce.Do(func() {
		file_chalk_kubernetes_v1_nodes_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_kubernetes_v1_nodes_proto_rawDescData)
	})
	return file_chalk_kubernetes_v1_nodes_proto_rawDescData
}

var file_chalk_kubernetes_v1_nodes_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chalk_kubernetes_v1_nodes_proto_goTypes = []any{
	(*KubernetesNodeData)(nil), // 0: chalk.kubernetes.v1.KubernetesNodeData
	nil,                        // 1: chalk.kubernetes.v1.KubernetesNodeData.LabelsEntry
	nil,                        // 2: chalk.kubernetes.v1.KubernetesNodeData.AnnotationsEntry
}
var file_chalk_kubernetes_v1_nodes_proto_depIdxs = []int32{
	1, // 0: chalk.kubernetes.v1.KubernetesNodeData.labels:type_name -> chalk.kubernetes.v1.KubernetesNodeData.LabelsEntry
	2, // 1: chalk.kubernetes.v1.KubernetesNodeData.annotations:type_name -> chalk.kubernetes.v1.KubernetesNodeData.AnnotationsEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_chalk_kubernetes_v1_nodes_proto_init() }
func file_chalk_kubernetes_v1_nodes_proto_init() {
	if File_chalk_kubernetes_v1_nodes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chalk_kubernetes_v1_nodes_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*KubernetesNodeData); i {
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
			RawDescriptor: file_chalk_kubernetes_v1_nodes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chalk_kubernetes_v1_nodes_proto_goTypes,
		DependencyIndexes: file_chalk_kubernetes_v1_nodes_proto_depIdxs,
		MessageInfos:      file_chalk_kubernetes_v1_nodes_proto_msgTypes,
	}.Build()
	File_chalk_kubernetes_v1_nodes_proto = out.File
	file_chalk_kubernetes_v1_nodes_proto_rawDesc = nil
	file_chalk_kubernetes_v1_nodes_proto_goTypes = nil
	file_chalk_kubernetes_v1_nodes_proto_depIdxs = nil
}
