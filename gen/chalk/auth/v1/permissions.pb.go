// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: chalk/auth/v1/permissions.proto

package authv1

import (
	_ "github.com/chalk-ai/chalk-go/gen/chalk/utils/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Permission int32

const (
	// Default value -- should never be set.
	Permission_PERMISSION_UNSPECIFIED Permission = 0
	// Only used for creating a new token
	Permission_PERMISSION_INSECURE_UNAUTHENTICATED Permission = 1
	// User is authenticated FOR THE ENVIRONMENT. This permissions should be used sparingly
	// (e.g. for ping checks that shouldn't be exposed to the public)
	Permission_PERMISSION_AUTHENTICATED Permission = 2
	// Query online features.
	Permission_PERMISSION_QUERY_ONLINE Permission = 3
	// Query offline features.
	Permission_PERMISSION_QUERY_OFFLINE Permission = 4
	// Create a new chart or alert
	Permission_PERMISSION_MONITORING_CREATE Permission = 5
	// See charts and alerts
	Permission_PERMISSION_MONITORING_READ Permission = 6
	// Add team members to your organization
	Permission_PERMISSION_TEAM_ADD Permission = 7
	// Remove a team member
	Permission_PERMISSION_TEAM_DELETE Permission = 8
	// See the team members in your organization
	Permission_PERMISSION_TEAM_LIST Permission = 9
	// Configure authentication options for your organization
	Permission_PERMISSION_TEAM_ADMIN Permission = 10
	// Read information about deployments.
	Permission_PERMISSION_DEPLOY_READ Permission = 11
	// Create a new deployment.
	Permission_PERMISSION_DEPLOY_CREATE Permission = 12
	// Create a new preview deployment.
	Permission_PERMISSION_DEPLOY_PREVIEW Permission = 13
	// Redeploy an existing deployment.
	Permission_PERMISSION_DEPLOY_REDEPLOY Permission = 14
	// Read logs from resolvers.
	Permission_PERMISSION_LOGS_LIST Permission = 15
	// Read the scheduled runs.
	Permission_PERMISSION_CRON_READ Permission = 16
	// Trigger a new scheduled run.
	Permission_PERMISSION_CRON_CREATE Permission = 17
	// Create, modify, or delete secret values.
	Permission_PERMISSION_SECRETS_WRITE Permission = 18
	// Decrypt secret values.
	Permission_PERMISSION_SECRETS_DECRYPT Permission = 19
	// See the list of available secrets. Reading secrets is not allowed with this permission.
	Permission_PERMISSION_SECRETS_LIST Permission = 20
	// Create, modify, or delete service tokens.
	Permission_PERMISSION_TOKENS_WRITE Permission = 21
	// List the service tokens and see client ids, but not client secrets.
	Permission_PERMISSION_TOKENS_LIST Permission = 22
	// View information about migrations.
	Permission_PERMISSION_MIGRATE_READ Permission = 23
	// Create a migration plan.
	Permission_PERMISSION_MIGRATE_PLAN Permission = 24
	// Execute a migration plan.
	Permission_PERMISSION_MIGRATE_EXECUTE Permission = 25
	// Create a new project
	Permission_PERMISSION_PROJECT_CREATE Permission = 26
	// Administer Chalk
	Permission_PERMISSION_CHALK_ADMIN Permission = 27
	// Read billing information
	Permission_PERMISSION_BILLING_READ Permission = 28
	// Read prompts information
	Permission_PERMISSION_PROMPTS_READ Permission = 29
	// Create, modify, or delete prompts
	Permission_PERMISSION_PROMPTS_WRITE Permission = 30
	// Run evaluations on prompts
	Permission_PERMISSION_PROMPTS_EVALUATE Permission = 31
)

// Enum value maps for Permission.
var (
	Permission_name = map[int32]string{
		0:  "PERMISSION_UNSPECIFIED",
		1:  "PERMISSION_INSECURE_UNAUTHENTICATED",
		2:  "PERMISSION_AUTHENTICATED",
		3:  "PERMISSION_QUERY_ONLINE",
		4:  "PERMISSION_QUERY_OFFLINE",
		5:  "PERMISSION_MONITORING_CREATE",
		6:  "PERMISSION_MONITORING_READ",
		7:  "PERMISSION_TEAM_ADD",
		8:  "PERMISSION_TEAM_DELETE",
		9:  "PERMISSION_TEAM_LIST",
		10: "PERMISSION_TEAM_ADMIN",
		11: "PERMISSION_DEPLOY_READ",
		12: "PERMISSION_DEPLOY_CREATE",
		13: "PERMISSION_DEPLOY_PREVIEW",
		14: "PERMISSION_DEPLOY_REDEPLOY",
		15: "PERMISSION_LOGS_LIST",
		16: "PERMISSION_CRON_READ",
		17: "PERMISSION_CRON_CREATE",
		18: "PERMISSION_SECRETS_WRITE",
		19: "PERMISSION_SECRETS_DECRYPT",
		20: "PERMISSION_SECRETS_LIST",
		21: "PERMISSION_TOKENS_WRITE",
		22: "PERMISSION_TOKENS_LIST",
		23: "PERMISSION_MIGRATE_READ",
		24: "PERMISSION_MIGRATE_PLAN",
		25: "PERMISSION_MIGRATE_EXECUTE",
		26: "PERMISSION_PROJECT_CREATE",
		27: "PERMISSION_CHALK_ADMIN",
		28: "PERMISSION_BILLING_READ",
		29: "PERMISSION_PROMPTS_READ",
		30: "PERMISSION_PROMPTS_WRITE",
		31: "PERMISSION_PROMPTS_EVALUATE",
	}
	Permission_value = map[string]int32{
		"PERMISSION_UNSPECIFIED":              0,
		"PERMISSION_INSECURE_UNAUTHENTICATED": 1,
		"PERMISSION_AUTHENTICATED":            2,
		"PERMISSION_QUERY_ONLINE":             3,
		"PERMISSION_QUERY_OFFLINE":            4,
		"PERMISSION_MONITORING_CREATE":        5,
		"PERMISSION_MONITORING_READ":          6,
		"PERMISSION_TEAM_ADD":                 7,
		"PERMISSION_TEAM_DELETE":              8,
		"PERMISSION_TEAM_LIST":                9,
		"PERMISSION_TEAM_ADMIN":               10,
		"PERMISSION_DEPLOY_READ":              11,
		"PERMISSION_DEPLOY_CREATE":            12,
		"PERMISSION_DEPLOY_PREVIEW":           13,
		"PERMISSION_DEPLOY_REDEPLOY":          14,
		"PERMISSION_LOGS_LIST":                15,
		"PERMISSION_CRON_READ":                16,
		"PERMISSION_CRON_CREATE":              17,
		"PERMISSION_SECRETS_WRITE":            18,
		"PERMISSION_SECRETS_DECRYPT":          19,
		"PERMISSION_SECRETS_LIST":             20,
		"PERMISSION_TOKENS_WRITE":             21,
		"PERMISSION_TOKENS_LIST":              22,
		"PERMISSION_MIGRATE_READ":             23,
		"PERMISSION_MIGRATE_PLAN":             24,
		"PERMISSION_MIGRATE_EXECUTE":          25,
		"PERMISSION_PROJECT_CREATE":           26,
		"PERMISSION_CHALK_ADMIN":              27,
		"PERMISSION_BILLING_READ":             28,
		"PERMISSION_PROMPTS_READ":             29,
		"PERMISSION_PROMPTS_WRITE":            30,
		"PERMISSION_PROMPTS_EVALUATE":         31,
	}
)

func (x Permission) Enum() *Permission {
	p := new(Permission)
	*p = x
	return p
}

func (x Permission) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Permission) Descriptor() protoreflect.EnumDescriptor {
	return file_chalk_auth_v1_permissions_proto_enumTypes[0].Descriptor()
}

func (Permission) Type() protoreflect.EnumType {
	return &file_chalk_auth_v1_permissions_proto_enumTypes[0]
}

func (x Permission) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Permission.Descriptor instead.
func (Permission) EnumDescriptor() ([]byte, []int) {
	return file_chalk_auth_v1_permissions_proto_rawDescGZIP(), []int{0}
}

var file_chalk_auth_v1_permissions_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Permission)(nil),
		Field:         2000,
		Name:          "chalk.auth.v1.permission",
		Tag:           "varint,2000,opt,name=permission,enum=chalk.auth.v1.Permission",
		Filename:      "chalk/auth/v1/permissions.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional chalk.auth.v1.Permission permission = 2000;
	E_Permission = &file_chalk_auth_v1_permissions_proto_extTypes[0]
)

var File_chalk_auth_v1_permissions_proto protoreflect.FileDescriptor

var file_chalk_auth_v1_permissions_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0d, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31,
	0x1a, 0x1d, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0x8d, 0x0c, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x16, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x27, 0x0a, 0x23,
	0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x53, 0x45, 0x43,
	0x55, 0x52, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41,
	0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x03,
	0x12, 0x1c, 0x0a, 0x18, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x51,
	0x55, 0x45, 0x52, 0x59, 0x5f, 0x4f, 0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x04, 0x12, 0x20,
	0x0a, 0x1c, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x4f, 0x4e,
	0x49, 0x54, 0x4f, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x05,
	0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4d,
	0x4f, 0x4e, 0x49, 0x54, 0x4f, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x06,
	0x12, 0x17, 0x0a, 0x13, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x45, 0x41, 0x4d, 0x5f, 0x41, 0x44, 0x44, 0x10, 0x07, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x45, 0x52,
	0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x44, 0x45, 0x4c,
	0x45, 0x54, 0x45, 0x10, 0x08, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x45, 0x41, 0x4d, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x09, 0x12,
	0x19, 0x0a, 0x15, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x45,
	0x41, 0x4d, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x45,
	0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x5f,
	0x52, 0x45, 0x41, 0x44, 0x10, 0x0b, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x5f, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x10, 0x0c, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x5f, 0x50, 0x52, 0x45, 0x56, 0x49, 0x45,
	0x57, 0x10, 0x0d, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x5f, 0x52, 0x45, 0x44, 0x45, 0x50, 0x4c, 0x4f,
	0x59, 0x10, 0x0e, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x4c, 0x4f, 0x47, 0x53, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x0f, 0x12, 0x18, 0x0a,
	0x14, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x52, 0x4f, 0x4e,
	0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x10, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x45, 0x52, 0x4d, 0x49,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x52, 0x4f, 0x4e, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x10, 0x11, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x53, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10,
	0x12, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x53, 0x5f, 0x44, 0x45, 0x43, 0x52, 0x59, 0x50, 0x54, 0x10,
	0x13, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x53, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x14, 0x12, 0x1b,
	0x0a, 0x17, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x4b,
	0x45, 0x4e, 0x53, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x10, 0x15, 0x12, 0x1a, 0x0a, 0x16, 0x50,
	0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x53,
	0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x16, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x45, 0x52, 0x4d, 0x49,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x49, 0x47, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45,
	0x41, 0x44, 0x10, 0x17, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x4d, 0x49, 0x47, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x4c, 0x41, 0x4e, 0x10,
	0x18, 0x12, 0x1e, 0x0a, 0x1a, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x4d, 0x49, 0x47, 0x52, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x45, 0x10,
	0x19, 0x12, 0x1d, 0x0a, 0x19, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x1a,
	0x12, 0x1a, 0x0a, 0x16, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x43,
	0x48, 0x41, 0x4c, 0x4b, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x1b, 0x12, 0x1b, 0x0a, 0x17,
	0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x42, 0x49, 0x4c, 0x4c, 0x49,
	0x4e, 0x47, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x10, 0x1c, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x45, 0x52,
	0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x4d, 0x50, 0x54, 0x53, 0x5f,
	0x52, 0x45, 0x41, 0x44, 0x10, 0x1d, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x4d, 0x50, 0x54, 0x53, 0x5f, 0x57, 0x52, 0x49,
	0x54, 0x45, 0x10, 0x1e, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x4d, 0x50, 0x54, 0x53, 0x5f, 0x45, 0x56, 0x41, 0x4c, 0x55,
	0x41, 0x54, 0x45, 0x10, 0x1f, 0x1a, 0xc9, 0x04, 0xe2, 0xa1, 0x27, 0xc4, 0x04, 0x0a, 0x1c, 0x08,
	0x01, 0x12, 0x18, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x5f, 0x75, 0x6e, 0x61, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x0a, 0x11, 0x08, 0x02, 0x12,
	0x0d, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x0a, 0x10,
	0x08, 0x03, 0x12, 0x0c, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x0a, 0x11, 0x08, 0x04, 0x12, 0x0d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x6f, 0x66, 0x66, 0x6c,
	0x69, 0x6e, 0x65, 0x0a, 0x15, 0x08, 0x05, 0x12, 0x11, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72,
	0x69, 0x6e, 0x67, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x0a, 0x13, 0x08, 0x06, 0x12, 0x0f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x0a,
	0x0c, 0x08, 0x07, 0x12, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x61, 0x64, 0x64, 0x0a, 0x0f, 0x08,
	0x08, 0x12, 0x0b, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x0a, 0x0d,
	0x08, 0x09, 0x12, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x0a, 0x0e, 0x08,
	0x0a, 0x12, 0x0a, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x0a, 0x0f, 0x08,
	0x0b, 0x12, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x0a, 0x11,
	0x08, 0x0c, 0x12, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x0a, 0x12, 0x08, 0x0d, 0x12, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x2e, 0x70, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x0a, 0x13, 0x08, 0x0e, 0x12, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x2e, 0x72, 0x65, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x0a, 0x0d, 0x08, 0x0f, 0x12, 0x09,
	0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x0a, 0x0d, 0x08, 0x10, 0x12, 0x09, 0x63,
	0x72, 0x6f, 0x6e, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x0a, 0x0f, 0x08, 0x11, 0x12, 0x0b, 0x63, 0x72,
	0x6f, 0x6e, 0x2e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x0a, 0x11, 0x08, 0x12, 0x12, 0x0d, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x65, 0x0a, 0x13, 0x08, 0x13,
	0x12, 0x0f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x0a, 0x10, 0x08, 0x14, 0x12, 0x0c, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x2e, 0x6c,
	0x69, 0x73, 0x74, 0x0a, 0x10, 0x08, 0x15, 0x12, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x0a, 0x0f, 0x08, 0x16, 0x12, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x0a, 0x10, 0x08, 0x17, 0x12, 0x0c, 0x6d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x65, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x0a, 0x10, 0x08, 0x18, 0x12, 0x0c, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x0a, 0x13, 0x08, 0x19, 0x12, 0x0f,
	0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x0a,
	0x12, 0x08, 0x1a, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x0a, 0x0f, 0x08, 0x1b, 0x12, 0x0b, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x0a, 0x10, 0x08, 0x1c, 0x12, 0x0c, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x0a, 0x10, 0x08, 0x1d, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x73, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x0a, 0x11, 0x08, 0x1e, 0x12, 0x0d, 0x70, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x73, 0x2e, 0x77, 0x72, 0x69, 0x74, 0x65, 0x0a, 0x14, 0x08, 0x1f, 0x12,
	0x10, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x73, 0x2e, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74,
	0x65, 0x3a, 0x5a, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xd0, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0xb2, 0x01,
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x42, 0x10, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x43, 0x41, 0x58, 0xaa, 0x02, 0x0d, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0d, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x41, 0x75, 0x74,
	0x68, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x19, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x41, 0x75, 0x74,
	0x68, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x68, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_auth_v1_permissions_proto_rawDescOnce sync.Once
	file_chalk_auth_v1_permissions_proto_rawDescData = file_chalk_auth_v1_permissions_proto_rawDesc
)

func file_chalk_auth_v1_permissions_proto_rawDescGZIP() []byte {
	file_chalk_auth_v1_permissions_proto_rawDescOnce.Do(func() {
		file_chalk_auth_v1_permissions_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_auth_v1_permissions_proto_rawDescData)
	})
	return file_chalk_auth_v1_permissions_proto_rawDescData
}

var file_chalk_auth_v1_permissions_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_chalk_auth_v1_permissions_proto_goTypes = []any{
	(Permission)(0),                    // 0: chalk.auth.v1.Permission
	(*descriptorpb.MethodOptions)(nil), // 1: google.protobuf.MethodOptions
}
var file_chalk_auth_v1_permissions_proto_depIdxs = []int32{
	1, // 0: chalk.auth.v1.permission:extendee -> google.protobuf.MethodOptions
	0, // 1: chalk.auth.v1.permission:type_name -> chalk.auth.v1.Permission
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chalk_auth_v1_permissions_proto_init() }
func file_chalk_auth_v1_permissions_proto_init() {
	if File_chalk_auth_v1_permissions_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_auth_v1_permissions_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_chalk_auth_v1_permissions_proto_goTypes,
		DependencyIndexes: file_chalk_auth_v1_permissions_proto_depIdxs,
		EnumInfos:         file_chalk_auth_v1_permissions_proto_enumTypes,
		ExtensionInfos:    file_chalk_auth_v1_permissions_proto_extTypes,
	}.Build()
	File_chalk_auth_v1_permissions_proto = out.File
	file_chalk_auth_v1_permissions_proto_rawDesc = nil
	file_chalk_auth_v1_permissions_proto_goTypes = nil
	file_chalk_auth_v1_permissions_proto_depIdxs = nil
}
