// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        (unknown)
// source: chalk/artifacts/v1/deployment.proto

package artifactsv1

import (
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

type DeploymentArtifacts struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Graph         *v1.Graph              `protobuf:"bytes,1,opt,name=graph,proto3" json:"graph,omitempty"`
	Crons         []*CronQuery           `protobuf:"bytes,2,rep,name=crons,proto3" json:"crons,omitempty"`
	Charts        []*Chart               `protobuf:"bytes,3,rep,name=charts,proto3" json:"charts,omitempty"`
	CdcSources    []*CDCSource           `protobuf:"bytes,4,rep,name=cdc_sources,json=cdcSources,proto3" json:"cdc_sources,omitempty"`
	Config        *ProjectSettings       `protobuf:"bytes,5,opt,name=config,proto3" json:"config,omitempty"`
	Chalkpy       *ChalkpyInfo           `protobuf:"bytes,6,opt,name=chalkpy,proto3" json:"chalkpy,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeploymentArtifacts) Reset() {
	*x = DeploymentArtifacts{}
	mi := &file_chalk_artifacts_v1_deployment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeploymentArtifacts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentArtifacts) ProtoMessage() {}

func (x *DeploymentArtifacts) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_artifacts_v1_deployment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentArtifacts.ProtoReflect.Descriptor instead.
func (*DeploymentArtifacts) Descriptor() ([]byte, []int) {
	return file_chalk_artifacts_v1_deployment_proto_rawDescGZIP(), []int{0}
}

func (x *DeploymentArtifacts) GetGraph() *v1.Graph {
	if x != nil {
		return x.Graph
	}
	return nil
}

func (x *DeploymentArtifacts) GetCrons() []*CronQuery {
	if x != nil {
		return x.Crons
	}
	return nil
}

func (x *DeploymentArtifacts) GetCharts() []*Chart {
	if x != nil {
		return x.Charts
	}
	return nil
}

func (x *DeploymentArtifacts) GetCdcSources() []*CDCSource {
	if x != nil {
		return x.CdcSources
	}
	return nil
}

func (x *DeploymentArtifacts) GetConfig() *ProjectSettings {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *DeploymentArtifacts) GetChalkpy() *ChalkpyInfo {
	if x != nil {
		return x.Chalkpy
	}
	return nil
}

var File_chalk_artifacts_v1_deployment_proto protoreflect.FileDescriptor

var file_chalk_artifacts_v1_deployment_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x63, 0x68, 0x61, 0x6c, 0x6b,
	0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x64,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x68, 0x61, 0x72,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x6f, 0x6e,
	0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x02, 0x0a, 0x13, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x73, 0x12, 0x2b, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x12, 0x33,
	0x0a, 0x05, 0x63, 0x72, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x6f, 0x6e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x05, 0x63, 0x72,
	0x6f, 0x6e, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52, 0x06,
	0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x12, 0x3e, 0x0a, 0x0b, 0x63, 0x64, 0x63, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x68,
	0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x44, 0x43, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0a, 0x63, 0x64, 0x63, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x3b, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x39, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x70, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x70,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x70, 0x79, 0x42, 0xd4,
	0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61,
	0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x43, 0x41, 0x58, 0xaa, 0x02, 0x12, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x43, 0x68, 0x61, 0x6c, 0x6b,
	0x5c, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e,
	0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x14, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_artifacts_v1_deployment_proto_rawDescOnce sync.Once
	file_chalk_artifacts_v1_deployment_proto_rawDescData = file_chalk_artifacts_v1_deployment_proto_rawDesc
)

func file_chalk_artifacts_v1_deployment_proto_rawDescGZIP() []byte {
	file_chalk_artifacts_v1_deployment_proto_rawDescOnce.Do(func() {
		file_chalk_artifacts_v1_deployment_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_artifacts_v1_deployment_proto_rawDescData)
	})
	return file_chalk_artifacts_v1_deployment_proto_rawDescData
}

var file_chalk_artifacts_v1_deployment_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chalk_artifacts_v1_deployment_proto_goTypes = []any{
	(*DeploymentArtifacts)(nil), // 0: chalk.artifacts.v1.DeploymentArtifacts
	(*v1.Graph)(nil),            // 1: chalk.graph.v1.Graph
	(*CronQuery)(nil),           // 2: chalk.artifacts.v1.CronQuery
	(*Chart)(nil),               // 3: chalk.artifacts.v1.Chart
	(*CDCSource)(nil),           // 4: chalk.artifacts.v1.CDCSource
	(*ProjectSettings)(nil),     // 5: chalk.artifacts.v1.ProjectSettings
	(*ChalkpyInfo)(nil),         // 6: chalk.artifacts.v1.ChalkpyInfo
}
var file_chalk_artifacts_v1_deployment_proto_depIdxs = []int32{
	1, // 0: chalk.artifacts.v1.DeploymentArtifacts.graph:type_name -> chalk.graph.v1.Graph
	2, // 1: chalk.artifacts.v1.DeploymentArtifacts.crons:type_name -> chalk.artifacts.v1.CronQuery
	3, // 2: chalk.artifacts.v1.DeploymentArtifacts.charts:type_name -> chalk.artifacts.v1.Chart
	4, // 3: chalk.artifacts.v1.DeploymentArtifacts.cdc_sources:type_name -> chalk.artifacts.v1.CDCSource
	5, // 4: chalk.artifacts.v1.DeploymentArtifacts.config:type_name -> chalk.artifacts.v1.ProjectSettings
	6, // 5: chalk.artifacts.v1.DeploymentArtifacts.chalkpy:type_name -> chalk.artifacts.v1.ChalkpyInfo
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_chalk_artifacts_v1_deployment_proto_init() }
func file_chalk_artifacts_v1_deployment_proto_init() {
	if File_chalk_artifacts_v1_deployment_proto != nil {
		return
	}
	file_chalk_artifacts_v1_cdc_proto_init()
	file_chalk_artifacts_v1_chart_proto_init()
	file_chalk_artifacts_v1_cron_query_proto_init()
	file_chalk_artifacts_v1_export_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_artifacts_v1_deployment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chalk_artifacts_v1_deployment_proto_goTypes,
		DependencyIndexes: file_chalk_artifacts_v1_deployment_proto_depIdxs,
		MessageInfos:      file_chalk_artifacts_v1_deployment_proto_msgTypes,
	}.Build()
	File_chalk_artifacts_v1_deployment_proto = out.File
	file_chalk_artifacts_v1_deployment_proto_rawDesc = nil
	file_chalk_artifacts_v1_deployment_proto_goTypes = nil
	file_chalk_artifacts_v1_deployment_proto_depIdxs = nil
}
