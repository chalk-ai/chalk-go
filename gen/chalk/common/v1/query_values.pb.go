// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: chalk/common/v1/query_values.proto

package commonv1

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

type OperationIdTableIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If you want the results for a particular operation id, no need to look up the value tables separately.
	// The engine will do that for you
	OperationId string `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
}

func (x *OperationIdTableIdentifier) Reset() {
	*x = OperationIdTableIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_common_v1_query_values_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationIdTableIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationIdTableIdentifier) ProtoMessage() {}

func (x *OperationIdTableIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_common_v1_query_values_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationIdTableIdentifier.ProtoReflect.Descriptor instead.
func (*OperationIdTableIdentifier) Descriptor() ([]byte, []int) {
	return file_chalk_common_v1_query_values_proto_rawDescGZIP(), []int{0}
}

func (x *OperationIdTableIdentifier) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

type TableNameTableIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableName string `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	// Optionally, you can filter. Specifying any filters will result in a join against the query log table.
	Filters *QueryLogFilters `protobuf:"bytes,2,opt,name=filters,proto3" json:"filters,omitempty"`
}

func (x *TableNameTableIdentifier) Reset() {
	*x = TableNameTableIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_common_v1_query_values_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableNameTableIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableNameTableIdentifier) ProtoMessage() {}

func (x *TableNameTableIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_common_v1_query_values_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableNameTableIdentifier.ProtoReflect.Descriptor instead.
func (*TableNameTableIdentifier) Descriptor() ([]byte, []int) {
	return file_chalk_common_v1_query_values_proto_rawDescGZIP(), []int{1}
}

func (x *TableNameTableIdentifier) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *TableNameTableIdentifier) GetFilters() *QueryLogFilters {
	if x != nil {
		return x.Filters
	}
	return nil
}

// Internal protobuf representing a next page token. Contains the operation id and the query timestamp for the last row in the pervious batch. Results are sorted query timestamp
// then by operation id lexagraphically then by row id, so this is all we need to know where the next page begins
type GetQueryValuesPageToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryTimestampHwm *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=query_timestamp_hwm,json=queryTimestampHwm,proto3" json:"query_timestamp_hwm,omitempty"`
	OperationIdHwm    string                 `protobuf:"bytes,2,opt,name=operation_id_hwm,json=operationIdHwm,proto3" json:"operation_id_hwm,omitempty"`
	RowIdHwm          int64                  `protobuf:"varint,3,opt,name=row_id_hwm,json=rowIdHwm,proto3" json:"row_id_hwm,omitempty"`
}

func (x *GetQueryValuesPageToken) Reset() {
	*x = GetQueryValuesPageToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_common_v1_query_values_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQueryValuesPageToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQueryValuesPageToken) ProtoMessage() {}

func (x *GetQueryValuesPageToken) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_common_v1_query_values_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQueryValuesPageToken.ProtoReflect.Descriptor instead.
func (*GetQueryValuesPageToken) Descriptor() ([]byte, []int) {
	return file_chalk_common_v1_query_values_proto_rawDescGZIP(), []int{2}
}

func (x *GetQueryValuesPageToken) GetQueryTimestampHwm() *timestamppb.Timestamp {
	if x != nil {
		return x.QueryTimestampHwm
	}
	return nil
}

func (x *GetQueryValuesPageToken) GetOperationIdHwm() string {
	if x != nil {
		return x.OperationIdHwm
	}
	return ""
}

func (x *GetQueryValuesPageToken) GetRowIdHwm() int64 {
	if x != nil {
		return x.RowIdHwm
	}
	return 0
}

type GetQueryValuesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to TableIdentifier:
	//
	//	*GetQueryValuesRequest_OperationIdIdentifier
	//	*GetQueryValuesRequest_TableNameIdentifier
	TableIdentifier isGetQueryValuesRequest_TableIdentifier `protobuf_oneof:"table_identifier"`
	// The query log table is partitioned / sorted by timestamp, so we must provide these
	// for the queries to be efficient, even if we are querying for a single operation id
	// This is always required.
	// If you know the operation id, then its feasible that you know the exact query timestamp, too.
	QueryTimestampLowerBoundInclusive *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=query_timestamp_lower_bound_inclusive,json=queryTimestampLowerBoundInclusive,proto3" json:"query_timestamp_lower_bound_inclusive,omitempty"`
	// If the upper bound is ommitted, then the lower bound will be used as an exact (equality) filter
	QueryTimestampUpperBoundExclusive *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=query_timestamp_upper_bound_exclusive,json=queryTimestampUpperBoundExclusive,proto3,oneof" json:"query_timestamp_upper_bound_exclusive,omitempty"`
	// If you're insterested in a subset of features, specify those here. Other columns won't be selected from the database, which will help reduce query costs.
	// If empty, all features will be returned
	Features []string `protobuf:"bytes,5,rep,name=features,proto3" json:"features,omitempty"`
	// The (maximum) page size for results. If zero, then the server picks.
	PageSize int32 `protobuf:"varint,7,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// When dealing with paginated responses, specify the next token to resume where you left off. The subsequent request must be identicial to the original (except for the value of the next_token)
	// Leave empty if querying for the zeroth page.
	PageToken string `protobuf:"bytes,8,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *GetQueryValuesRequest) Reset() {
	*x = GetQueryValuesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_common_v1_query_values_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQueryValuesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQueryValuesRequest) ProtoMessage() {}

func (x *GetQueryValuesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_common_v1_query_values_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQueryValuesRequest.ProtoReflect.Descriptor instead.
func (*GetQueryValuesRequest) Descriptor() ([]byte, []int) {
	return file_chalk_common_v1_query_values_proto_rawDescGZIP(), []int{3}
}

func (m *GetQueryValuesRequest) GetTableIdentifier() isGetQueryValuesRequest_TableIdentifier {
	if m != nil {
		return m.TableIdentifier
	}
	return nil
}

func (x *GetQueryValuesRequest) GetOperationIdIdentifier() *OperationIdTableIdentifier {
	if x, ok := x.GetTableIdentifier().(*GetQueryValuesRequest_OperationIdIdentifier); ok {
		return x.OperationIdIdentifier
	}
	return nil
}

func (x *GetQueryValuesRequest) GetTableNameIdentifier() *TableNameTableIdentifier {
	if x, ok := x.GetTableIdentifier().(*GetQueryValuesRequest_TableNameIdentifier); ok {
		return x.TableNameIdentifier
	}
	return nil
}

func (x *GetQueryValuesRequest) GetQueryTimestampLowerBoundInclusive() *timestamppb.Timestamp {
	if x != nil {
		return x.QueryTimestampLowerBoundInclusive
	}
	return nil
}

func (x *GetQueryValuesRequest) GetQueryTimestampUpperBoundExclusive() *timestamppb.Timestamp {
	if x != nil {
		return x.QueryTimestampUpperBoundExclusive
	}
	return nil
}

func (x *GetQueryValuesRequest) GetFeatures() []string {
	if x != nil {
		return x.Features
	}
	return nil
}

func (x *GetQueryValuesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetQueryValuesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type isGetQueryValuesRequest_TableIdentifier interface {
	isGetQueryValuesRequest_TableIdentifier()
}

type GetQueryValuesRequest_OperationIdIdentifier struct {
	// Forcing the client to specify the table name can be a bit narly. Instead, for use case 1), it can be easier to allow the client to specify the operation id,
	// and the engine can figure out what table to query.
	OperationIdIdentifier *OperationIdTableIdentifier `protobuf:"bytes,1,opt,name=operation_id_identifier,json=operationIdIdentifier,proto3,oneof"`
}

type GetQueryValuesRequest_TableNameIdentifier struct {
	// For use case 2, you have to specify which table name to look up, if you want to query across multiple queries
	TableNameIdentifier *TableNameTableIdentifier `protobuf:"bytes,2,opt,name=table_name_identifier,json=tableNameIdentifier,proto3,oneof"`
}

func (*GetQueryValuesRequest_OperationIdIdentifier) isGetQueryValuesRequest_TableIdentifier() {}

func (*GetQueryValuesRequest_TableNameIdentifier) isGetQueryValuesRequest_TableIdentifier() {}

type GetQueryValuesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If non-empty, call this endpoint again, with this next token to get the next page of responses.
	NextPageToken string `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	// The response payload. Defining as a one-of to future-proof it should we decide to support multiple encodings (parquet, feather, ...)
	//
	// Types that are assignable to Payload:
	//
	//	*GetQueryValuesResponse_Parquet
	Payload isGetQueryValuesResponse_Payload `protobuf_oneof:"payload"`
}

func (x *GetQueryValuesResponse) Reset() {
	*x = GetQueryValuesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chalk_common_v1_query_values_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQueryValuesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQueryValuesResponse) ProtoMessage() {}

func (x *GetQueryValuesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chalk_common_v1_query_values_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQueryValuesResponse.ProtoReflect.Descriptor instead.
func (*GetQueryValuesResponse) Descriptor() ([]byte, []int) {
	return file_chalk_common_v1_query_values_proto_rawDescGZIP(), []int{4}
}

func (x *GetQueryValuesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

func (m *GetQueryValuesResponse) GetPayload() isGetQueryValuesResponse_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *GetQueryValuesResponse) GetParquet() []byte {
	if x, ok := x.GetPayload().(*GetQueryValuesResponse_Parquet); ok {
		return x.Parquet
	}
	return nil
}

type isGetQueryValuesResponse_Payload interface {
	isGetQueryValuesResponse_Payload()
}

type GetQueryValuesResponse_Parquet struct {
	Parquet []byte `protobuf:"bytes,2,opt,name=parquet,proto3,oneof"`
}

func (*GetQueryValuesResponse_Parquet) isGetQueryValuesResponse_Payload() {}

var File_chalk_common_v1_query_values_proto protoreflect.FileDescriptor

var file_chalk_common_v1_query_values_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76,
	0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6c, 0x6f, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x1a, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x75, 0x0a, 0x18, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4c, 0x6f, 0x67, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22,
	0xad, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x4a, 0x0a, 0x13, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x68,
	0x77, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x71, 0x75, 0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x48, 0x77, 0x6d, 0x12, 0x28, 0x0a, 0x10, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x68, 0x77, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x48, 0x77,
	0x6d, 0x12, 0x1c, 0x0a, 0x0a, 0x72, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x5f, 0x68, 0x77, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x6f, 0x77, 0x49, 0x64, 0x48, 0x77, 0x6d, 0x22,
	0xd6, 0x04, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x65, 0x0a, 0x17, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63, 0x68, 0x61,
	0x6c, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x48, 0x00, 0x52, 0x15, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x12, 0x5f, 0x0a, 0x15, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x48, 0x00, 0x52, 0x13, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x6c, 0x0a, 0x25, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x21, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4c, 0x6f, 0x77, 0x65,
	0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x12,
	0x71, 0x0a, 0x25, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x5f, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x65,
	0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x21, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x55, 0x70, 0x70, 0x65,
	0x72, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x12, 0x0a, 0x10, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x28,
	0x0a, 0x26, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x5f, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x65,
	0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x22, 0x67, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x07, 0x70, 0x61,
	0x72, 0x71, 0x75, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x07, 0x70,
	0x61, 0x72, 0x71, 0x75, 0x65, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x42, 0xc0, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d,
	0x61, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x63, 0x68, 0x61, 0x6c, 0x6b, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x58, 0xaa, 0x02,
	0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x1b, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x11, 0x43, 0x68, 0x61, 0x6c, 0x6b, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chalk_common_v1_query_values_proto_rawDescOnce sync.Once
	file_chalk_common_v1_query_values_proto_rawDescData = file_chalk_common_v1_query_values_proto_rawDesc
)

func file_chalk_common_v1_query_values_proto_rawDescGZIP() []byte {
	file_chalk_common_v1_query_values_proto_rawDescOnce.Do(func() {
		file_chalk_common_v1_query_values_proto_rawDescData = protoimpl.X.CompressGZIP(file_chalk_common_v1_query_values_proto_rawDescData)
	})
	return file_chalk_common_v1_query_values_proto_rawDescData
}

var file_chalk_common_v1_query_values_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_chalk_common_v1_query_values_proto_goTypes = []interface{}{
	(*OperationIdTableIdentifier)(nil), // 0: chalk.common.v1.OperationIdTableIdentifier
	(*TableNameTableIdentifier)(nil),   // 1: chalk.common.v1.TableNameTableIdentifier
	(*GetQueryValuesPageToken)(nil),    // 2: chalk.common.v1.GetQueryValuesPageToken
	(*GetQueryValuesRequest)(nil),      // 3: chalk.common.v1.GetQueryValuesRequest
	(*GetQueryValuesResponse)(nil),     // 4: chalk.common.v1.GetQueryValuesResponse
	(*QueryLogFilters)(nil),            // 5: chalk.common.v1.QueryLogFilters
	(*timestamppb.Timestamp)(nil),      // 6: google.protobuf.Timestamp
}
var file_chalk_common_v1_query_values_proto_depIdxs = []int32{
	5, // 0: chalk.common.v1.TableNameTableIdentifier.filters:type_name -> chalk.common.v1.QueryLogFilters
	6, // 1: chalk.common.v1.GetQueryValuesPageToken.query_timestamp_hwm:type_name -> google.protobuf.Timestamp
	0, // 2: chalk.common.v1.GetQueryValuesRequest.operation_id_identifier:type_name -> chalk.common.v1.OperationIdTableIdentifier
	1, // 3: chalk.common.v1.GetQueryValuesRequest.table_name_identifier:type_name -> chalk.common.v1.TableNameTableIdentifier
	6, // 4: chalk.common.v1.GetQueryValuesRequest.query_timestamp_lower_bound_inclusive:type_name -> google.protobuf.Timestamp
	6, // 5: chalk.common.v1.GetQueryValuesRequest.query_timestamp_upper_bound_exclusive:type_name -> google.protobuf.Timestamp
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_chalk_common_v1_query_values_proto_init() }
func file_chalk_common_v1_query_values_proto_init() {
	if File_chalk_common_v1_query_values_proto != nil {
		return
	}
	file_chalk_common_v1_query_log_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chalk_common_v1_query_values_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationIdTableIdentifier); i {
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
		file_chalk_common_v1_query_values_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableNameTableIdentifier); i {
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
		file_chalk_common_v1_query_values_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQueryValuesPageToken); i {
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
		file_chalk_common_v1_query_values_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQueryValuesRequest); i {
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
		file_chalk_common_v1_query_values_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQueryValuesResponse); i {
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
	file_chalk_common_v1_query_values_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*GetQueryValuesRequest_OperationIdIdentifier)(nil),
		(*GetQueryValuesRequest_TableNameIdentifier)(nil),
	}
	file_chalk_common_v1_query_values_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*GetQueryValuesResponse_Parquet)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chalk_common_v1_query_values_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chalk_common_v1_query_values_proto_goTypes,
		DependencyIndexes: file_chalk_common_v1_query_values_proto_depIdxs,
		MessageInfos:      file_chalk_common_v1_query_values_proto_msgTypes,
	}.Build()
	File_chalk_common_v1_query_values_proto = out.File
	file_chalk_common_v1_query_values_proto_rawDesc = nil
	file_chalk_common_v1_query_values_proto_goTypes = nil
	file_chalk_common_v1_query_values_proto_depIdxs = nil
}