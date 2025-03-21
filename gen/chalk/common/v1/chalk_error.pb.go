// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: chalk/common/v1/chalk_error.proto

package commonv1

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

type ErrorCode int32

const (
	// An unspecified error occurred.
	ErrorCode_ERROR_CODE_INTERNAL_SERVER_ERROR_UNSPECIFIED ErrorCode = 0
	// The query contained features that do not exist
	ErrorCode_ERROR_CODE_PARSE_FAILED ErrorCode = 1
	// A resolver was required as part of running the dependency
	// graph that could not be found.
	ErrorCode_ERROR_CODE_RESOLVER_NOT_FOUND ErrorCode = 2
	// The query is invalid. All supplied features need to be
	// rooted in the same top-level entity.
	ErrorCode_ERROR_CODE_INVALID_QUERY ErrorCode = 3
	// A feature value did not match the expected schema
	// (e.g. `incompatible type "int"; expected "str"`)
	ErrorCode_ERROR_CODE_VALIDATION_FAILED ErrorCode = 4
	// The resolver for a feature errored.
	ErrorCode_ERROR_CODE_RESOLVER_FAILED ErrorCode = 5
	// The resolver for a feature timed out.
	ErrorCode_ERROR_CODE_RESOLVER_TIMED_OUT ErrorCode = 6
	// A crash in a resolver that was to produce an input for
	// the resolver crashed, and so the resolver could not run
	// crashed, and so the resolver could not run.
	ErrorCode_ERROR_CODE_UPSTREAM_FAILED ErrorCode = 7
	// The request was submitted with an invalid authentication header.
	ErrorCode_ERROR_CODE_UNAUTHENTICATED ErrorCode = 8
	// The supplied credentials do not provide the right authorization to execute the request.
	ErrorCode_ERROR_CODE_UNAUTHORIZED ErrorCode = 9
	// The operation was cancelled, typically by the caller.
	ErrorCode_ERROR_CODE_CANCELLED ErrorCode = 10
	// The deadline expired before the operation could complete.
	ErrorCode_ERROR_CODE_DEADLINE_EXCEEDED ErrorCode = 11
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:  "ERROR_CODE_INTERNAL_SERVER_ERROR_UNSPECIFIED",
		1:  "ERROR_CODE_PARSE_FAILED",
		2:  "ERROR_CODE_RESOLVER_NOT_FOUND",
		3:  "ERROR_CODE_INVALID_QUERY",
		4:  "ERROR_CODE_VALIDATION_FAILED",
		5:  "ERROR_CODE_RESOLVER_FAILED",
		6:  "ERROR_CODE_RESOLVER_TIMED_OUT",
		7:  "ERROR_CODE_UPSTREAM_FAILED",
		8:  "ERROR_CODE_UNAUTHENTICATED",
		9:  "ERROR_CODE_UNAUTHORIZED",
		10: "ERROR_CODE_CANCELLED",
		11: "ERROR_CODE_DEADLINE_EXCEEDED",
	}
	ErrorCode_value = map[string]int32{
		"ERROR_CODE_INTERNAL_SERVER_ERROR_UNSPECIFIED": 0,
		"ERROR_CODE_PARSE_FAILED":                      1,
		"ERROR_CODE_RESOLVER_NOT_FOUND":                2,
		"ERROR_CODE_INVALID_QUERY":                     3,
		"ERROR_CODE_VALIDATION_FAILED":                 4,
		"ERROR_CODE_RESOLVER_FAILED":                   5,
		"ERROR_CODE_RESOLVER_TIMED_OUT":                6,
		"ERROR_CODE_UPSTREAM_FAILED":                   7,
		"ERROR_CODE_UNAUTHENTICATED":                   8,
		"ERROR_CODE_UNAUTHORIZED":                      9,
		"ERROR_CODE_CANCELLED":                         10,
		"ERROR_CODE_DEADLINE_EXCEEDED":                 11,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_chalk_common_v1_chalk_error_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_chalk_common_v1_chalk_error_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_chalk_common_v1_chalk_error_proto_rawDescGZIP(), []int{0}
}

type ErrorCodeCategory int32

const (
	// -- DEFAULT VALUE --
	ErrorCodeCategory_ERROR_CODE_CATEGORY_NETWORK_UNSPECIFIED ErrorCodeCategory = 0
	// Request errors are raised before execution of your
	// resolver code. They may occur due to invalid feature
	// names in the input or a request that cannot be satisfied
	// by the resolvers you have defined.
	ErrorCodeCategory_ERROR_CODE_CATEGORY_REQUEST ErrorCodeCategory = 1
	// Field errors are raised while running a feature resolver
	// for a particular field. For this type of error, you'll
	// find a feature and resolver attribute in the error type.
	// When a feature resolver crashes, you will receive null
	// value in the response. To differentiate from a resolver
	// returning a null value and a failure in the resolver,
	// you need to check the error schema.
	ErrorCodeCategory_ERROR_CODE_CATEGORY_FIELD ErrorCodeCategory = 2
)

// Enum value maps for ErrorCodeCategory.
var (
	ErrorCodeCategory_name = map[int32]string{
		0: "ERROR_CODE_CATEGORY_NETWORK_UNSPECIFIED",
		1: "ERROR_CODE_CATEGORY_REQUEST",
		2: "ERROR_CODE_CATEGORY_FIELD",
	}
	ErrorCodeCategory_value = map[string]int32{
		"ERROR_CODE_CATEGORY_NETWORK_UNSPECIFIED": 0,
		"ERROR_CODE_CATEGORY_REQUEST":             1,
		"ERROR_CODE_CATEGORY_FIELD":               2,
	}
)

func (x ErrorCodeCategory) Enum() *ErrorCodeCategory {
	p := new(ErrorCodeCategory)
	*p = x
	return p
}

func (x ErrorCodeCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCodeCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_chalk_common_v1_chalk_error_proto_enumTypes[1].Descriptor()
}

func (ErrorCodeCategory) Type() protoreflect.EnumType {
	return &file_chalk_common_v1_chalk_error_proto_enumTypes[1]
}

func (x ErrorCodeCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCodeCategory.Descriptor instead.
func (ErrorCodeCategory) EnumDescriptor() ([]byte, []int) {
	return file_chalk_common_v1_chalk_error_proto_rawDescGZIP(), []int{1}
}

type ChalkException struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the class of the exception.
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// The message taken from the exception.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// The stacktrace produced by the code.
	Stacktrace string `protobuf:"bytes,3,opt,name=stacktrace,proto3" json:"stacktrace,omitempty"`
	// The stacktrace produced by the code, full detail.
	InternalStacktrace string `protobuf:"bytes,4,opt,name=internal_stacktrace,json=internalStacktrace,proto3" json:"internal_stacktrace,omitempty"`
}

func (x *ChalkException) Reset() {
	*x = ChalkException{}
	mi := &file_chalk_common_v1_chalk_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChalkException) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChalkException) ProtoMessage() {}

func (x *ChalkException) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_common_v1_chalk_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChalkException.ProtoReflect.Descriptor instead.
func (*ChalkException) Descriptor() ([]byte, []int) {
	return file_chalk_common_v1_chalk_error_proto_rawDescGZIP(), []int{0}
}

func (x *ChalkException) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ChalkException) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ChalkException) GetStacktrace() string {
	if x != nil {
		return x.Stacktrace
	}
	return ""
}

func (x *ChalkException) GetInternalStacktrace() string {
	if x != nil {
		return x.InternalStacktrace
	}
	return ""
}

type ChalkError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     ErrorCode         `protobuf:"varint,1,opt,name=code,proto3,enum=chalk.common.v1.ErrorCode" json:"code,omitempty"`
	Category ErrorCodeCategory `protobuf:"varint,2,opt,name=category,proto3,enum=chalk.common.v1.ErrorCodeCategory" json:"category,omitempty"`
	// A readable description of the error message.
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// A human-readable hint that can be used to identify the entity that this error is associated with.
	DisplayPrimaryKey *string `protobuf:"bytes,101,opt,name=display_primary_key,json=displayPrimaryKey,proto3,oneof" json:"display_primary_key,omitempty"`
	// If provided, can be used to add additional context to 'display_primary_key'.
	DisplayPrimaryKeyFqn *string `protobuf:"bytes,102,opt,name=display_primary_key_fqn,json=displayPrimaryKeyFqn,proto3,oneof" json:"display_primary_key_fqn,omitempty"`
	// The exception that caused the failure, if applicable.
	Exception *ChalkException `protobuf:"bytes,103,opt,name=exception,proto3,oneof" json:"exception,omitempty"`
	// The fully qualified name of the failing feature, e.g. `user.identity.has_voip_phone`.
	Feature *string `protobuf:"bytes,104,opt,name=feature,proto3,oneof" json:"feature,omitempty"`
	// The fully qualified name of the failing resolver, e.g. `my.project.get_fraud_score`.
	Resolver *string `protobuf:"bytes,105,opt,name=resolver,proto3,oneof" json:"resolver,omitempty"`
}

func (x *ChalkError) Reset() {
	*x = ChalkError{}
	mi := &file_chalk_common_v1_chalk_error_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChalkError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChalkError) ProtoMessage() {}

func (x *ChalkError) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_common_v1_chalk_error_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChalkError.ProtoReflect.Descriptor instead.
func (*ChalkError) Descriptor() ([]byte, []int) {
	return file_chalk_common_v1_chalk_error_proto_rawDescGZIP(), []int{1}
}

func (x *ChalkError) GetCode() ErrorCode {
	if x != nil {
		return x.Code
	}
	return ErrorCode_ERROR_CODE_INTERNAL_SERVER_ERROR_UNSPECIFIED
}

func (x *ChalkError) GetCategory() ErrorCodeCategory {
	if x != nil {
		return x.Category
	}
	return ErrorCodeCategory_ERROR_CODE_CATEGORY_NETWORK_UNSPECIFIED
}

func (x *ChalkError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ChalkError) GetDisplayPrimaryKey() string {
	if x != nil && x.DisplayPrimaryKey != nil {
		return *x.DisplayPrimaryKey
	}
	return ""
}

func (x *ChalkError) GetDisplayPrimaryKeyFqn() string {
	if x != nil && x.DisplayPrimaryKeyFqn != nil {
		return *x.DisplayPrimaryKeyFqn
	}
	return ""
}

func (x *ChalkError) GetException() *ChalkException {
	if x != nil {
		return x.Exception
	}
	return nil
}

func (x *ChalkError) GetFeature() string {
	if x != nil && x.Feature != nil {
		return *x.Feature
	}
	return ""
}

func (x *ChalkError) GetResolver() string {
	if x != nil && x.Resolver != nil {
		return *x.Resolver
	}
	return ""
}

var File_chalk_common_v1_chalk_error_proto protoreflect.FileDescriptor

var file_chalk_common_v1_chalk_error_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x22, 0x8f, 0x01, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x45, 0x78,
	0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x5f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x63,
	0x6b, 0x74, 0x72, 0x61, 0x63, 0x65, 0x22, 0xe6, 0x03, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6c, 0x6b,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x33, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x11,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65,
	0x79, 0x88, 0x01, 0x01, 0x12, 0x3a, 0x0a, 0x17, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x66, 0x71, 0x6e, 0x18,
	0x66, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x14, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x46, 0x71, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x42, 0x0a, 0x09, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x67, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x45, 0x78, 0x63, 0x65, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x02, 0x52, 0x09, 0x65, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x68, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x18,
	0x69, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x72, 0x88, 0x01, 0x01, 0x42, 0x16, 0x0a, 0x14, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x42, 0x1a, 0x0a, 0x18,
	0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x66, 0x71, 0x6e, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x65, 0x78, 0x63,
	0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2a,
	0x99, 0x03, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x0a,
	0x2c, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45,
	0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x1b, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x41,
	0x52, 0x53, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x4c,
	0x56, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x02, 0x12,
	0x1c, 0x0a, 0x18, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x10, 0x03, 0x12, 0x20, 0x0a,
	0x1c, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12,
	0x1e, 0x0a, 0x1a, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x45,
	0x53, 0x4f, 0x4c, 0x56, 0x45, 0x52, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x12,
	0x21, 0x0a, 0x1d, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x45,
	0x53, 0x4f, 0x4c, 0x56, 0x45, 0x52, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x44, 0x5f, 0x4f, 0x55, 0x54,
	0x10, 0x06, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x55, 0x50, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0x07, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e, 0x54, 0x49, 0x43, 0x41, 0x54, 0x45, 0x44,
	0x10, 0x08, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x5f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x09, 0x12,
	0x18, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x41,
	0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x20, 0x0a, 0x1c, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x44, 0x45, 0x41, 0x44, 0x4c, 0x49, 0x4e, 0x45,
	0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x0b, 0x2a, 0x80, 0x01, 0x0a, 0x11,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x2b, 0x0a, 0x27, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f,
	0x0a, 0x1b, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x41, 0x54,
	0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12,
	0x1d, 0x0a, 0x19, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x41,
	0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x02, 0x42, 0xbf,
	0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x61, 0x69, 0x2f, 0x63,
	0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x6c,
	0x6b, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x58, 0xaa, 0x02, 0x0f, 0x43, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x43,
	0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1b, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x43,
	0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_common_v1_chalk_error_proto_rawDescOnce sync.Once
	file_chalk_common_v1_chalk_error_proto_rawDescData = file_chalk_common_v1_chalk_error_proto_rawDesc
)

func file_chalk_common_v1_chalk_error_proto_rawDescGZIP() []byte {
	file_chalk_common_v1_chalk_error_proto_rawDescOnce.Do(func() {
		file_chalk_common_v1_chalk_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_common_v1_chalk_error_proto_rawDescData)
	})
	return file_chalk_common_v1_chalk_error_proto_rawDescData
}

var file_chalk_common_v1_chalk_error_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_chalk_common_v1_chalk_error_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_chalk_common_v1_chalk_error_proto_goTypes = []any{
	(ErrorCode)(0),         // 0: chalk.common.v1.ErrorCode
	(ErrorCodeCategory)(0), // 1: chalk.common.v1.ErrorCodeCategory
	(*ChalkException)(nil), // 2: chalk.common.v1.ChalkException
	(*ChalkError)(nil),     // 3: chalk.common.v1.ChalkError
}
var file_chalk_common_v1_chalk_error_proto_depIdxs = []int32{
	0, // 0: chalk.common.v1.ChalkError.code:type_name -> chalk.common.v1.ErrorCode
	1, // 1: chalk.common.v1.ChalkError.category:type_name -> chalk.common.v1.ErrorCodeCategory
	2, // 2: chalk.common.v1.ChalkError.exception:type_name -> chalk.common.v1.ChalkException
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_chalk_common_v1_chalk_error_proto_init() }
func file_chalk_common_v1_chalk_error_proto_init() {
	if File_chalk_common_v1_chalk_error_proto != nil {
		return
	}
	file_chalk_common_v1_chalk_error_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_common_v1_chalk_error_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chalk_common_v1_chalk_error_proto_goTypes,
		DependencyIndexes: file_chalk_common_v1_chalk_error_proto_depIdxs,
		EnumInfos:         file_chalk_common_v1_chalk_error_proto_enumTypes,
		MessageInfos:      file_chalk_common_v1_chalk_error_proto_msgTypes,
	}.Build()
	File_chalk_common_v1_chalk_error_proto = out.File
	file_chalk_common_v1_chalk_error_proto_rawDesc = nil
	file_chalk_common_v1_chalk_error_proto_goTypes = nil
	file_chalk_common_v1_chalk_error_proto_depIdxs = nil
}
