// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/warehouse/warehouse.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetWHTablesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UploadId  int64  `protobuf:"varint,1,opt,name=upload_id,json=uploadId,proto3" json:"upload_id,omitempty"`
	TableName string `protobuf:"bytes,2,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
}

func (x *GetWHTablesRequest) Reset() {
	*x = GetWHTablesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_warehouse_warehouse_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWHTablesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWHTablesRequest) ProtoMessage() {}

func (x *GetWHTablesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_warehouse_warehouse_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWHTablesRequest.ProtoReflect.Descriptor instead.
func (*GetWHTablesRequest) Descriptor() ([]byte, []int) {
	return file_proto_warehouse_warehouse_proto_rawDescGZIP(), []int{0}
}

func (x *GetWHTablesRequest) GetUploadId() int64 {
	if x != nil {
		return x.UploadId
	}
	return 0
}

func (x *GetWHTablesRequest) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

type GetWHUploadsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceId           string `protobuf:"bytes,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	DestinationId      string `protobuf:"bytes,2,opt,name=destination_id,json=destinationId,proto3" json:"destination_id,omitempty"`
	DestinationType    string `protobuf:"bytes,3,opt,name=destination_type,json=destinationType,proto3" json:"destination_type,omitempty"`
	Status             string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	IncludeTablesInRes bool   `protobuf:"varint,5,opt,name=include_tables_in_res,json=includeTablesInRes,proto3" json:"include_tables_in_res,omitempty"`
	Limit              int32  `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset             int32  `protobuf:"varint,7,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *GetWHUploadsRequest) Reset() {
	*x = GetWHUploadsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_warehouse_warehouse_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWHUploadsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWHUploadsRequest) ProtoMessage() {}

func (x *GetWHUploadsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_warehouse_warehouse_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWHUploadsRequest.ProtoReflect.Descriptor instead.
func (*GetWHUploadsRequest) Descriptor() ([]byte, []int) {
	return file_proto_warehouse_warehouse_proto_rawDescGZIP(), []int{1}
}

func (x *GetWHUploadsRequest) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *GetWHUploadsRequest) GetDestinationId() string {
	if x != nil {
		return x.DestinationId
	}
	return ""
}

func (x *GetWHUploadsRequest) GetDestinationType() string {
	if x != nil {
		return x.DestinationType
	}
	return ""
}

func (x *GetWHUploadsRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetWHUploadsRequest) GetIncludeTablesInRes() bool {
	if x != nil {
		return x.IncludeTablesInRes
	}
	return false
}

func (x *GetWHUploadsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetWHUploadsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetWHUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UploadId int64 `protobuf:"varint,1,opt,name=upload_id,json=uploadId,proto3" json:"upload_id,omitempty"`
}

func (x *GetWHUploadRequest) Reset() {
	*x = GetWHUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_warehouse_warehouse_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWHUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWHUploadRequest) ProtoMessage() {}

func (x *GetWHUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_warehouse_warehouse_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWHUploadRequest.ProtoReflect.Descriptor instead.
func (*GetWHUploadRequest) Descriptor() ([]byte, []int) {
	return file_proto_warehouse_warehouse_proto_rawDescGZIP(), []int{2}
}

func (x *GetWHUploadRequest) GetUploadId() int64 {
	if x != nil {
		return x.UploadId
	}
	return 0
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total  int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_warehouse_warehouse_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_proto_warehouse_warehouse_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_proto_warehouse_warehouse_proto_rawDescGZIP(), []int{3}
}

func (x *Pagination) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Pagination) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Pagination) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type GetWHUploadsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uploads    []*GetWHUploadResponse `protobuf:"bytes,1,rep,name=uploads,proto3" json:"uploads,omitempty"`
	Pagination *Pagination            `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *GetWHUploadsResponse) Reset() {
	*x = GetWHUploadsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_warehouse_warehouse_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWHUploadsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWHUploadsResponse) ProtoMessage() {}

func (x *GetWHUploadsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_warehouse_warehouse_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWHUploadsResponse.ProtoReflect.Descriptor instead.
func (*GetWHUploadsResponse) Descriptor() ([]byte, []int) {
	return file_proto_warehouse_warehouse_proto_rawDescGZIP(), []int{4}
}

func (x *GetWHUploadsResponse) GetUploads() []*GetWHUploadResponse {
	if x != nil {
		return x.Uploads
	}
	return nil
}

func (x *GetWHUploadsResponse) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type GetWHTablesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tables []*WHTable `protobuf:"bytes,1,rep,name=tables,proto3" json:"tables,omitempty"`
}

func (x *GetWHTablesResponse) Reset() {
	*x = GetWHTablesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_warehouse_warehouse_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWHTablesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWHTablesResponse) ProtoMessage() {}

func (x *GetWHTablesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_warehouse_warehouse_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWHTablesResponse.ProtoReflect.Descriptor instead.
func (*GetWHTablesResponse) Descriptor() ([]byte, []int) {
	return file_proto_warehouse_warehouse_proto_rawDescGZIP(), []int{5}
}

func (x *GetWHTablesResponse) GetTables() []*WHTable {
	if x != nil {
		return x.Tables
	}
	return nil
}

type GetWHUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SourceId        string                 `protobuf:"bytes,2,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	DestinationId   string                 `protobuf:"bytes,3,opt,name=destination_id,json=destinationId,proto3" json:"destination_id,omitempty"`
	DestinationType string                 `protobuf:"bytes,4,opt,name=destination_type,json=destinationType,proto3" json:"destination_type,omitempty"`
	Namespace       string                 `protobuf:"bytes,5,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Error           string                 `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
	Attempt         int32                  `protobuf:"varint,7,opt,name=attempt,proto3" json:"attempt,omitempty"`
	Status          string                 `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	FirstEventAt    *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=first_event_at,json=firstEventAt,proto3" json:"first_event_at,omitempty"`
	LastEventAt     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=last_event_at,json=lastEventAt,proto3" json:"last_event_at,omitempty"`
	NextRetryTime   string                 `protobuf:"bytes,11,opt,name=next_retry_time,json=nextRetryTime,proto3" json:"next_retry_time,omitempty"`
	Duration        int32                  `protobuf:"varint,12,opt,name=duration,proto3" json:"duration,omitempty"`
	Tables          []*WHTable             `protobuf:"bytes,13,rep,name=tables,proto3" json:"tables,omitempty"`
}

func (x *GetWHUploadResponse) Reset() {
	*x = GetWHUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_warehouse_warehouse_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWHUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWHUploadResponse) ProtoMessage() {}

func (x *GetWHUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_warehouse_warehouse_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWHUploadResponse.ProtoReflect.Descriptor instead.
func (*GetWHUploadResponse) Descriptor() ([]byte, []int) {
	return file_proto_warehouse_warehouse_proto_rawDescGZIP(), []int{6}
}

func (x *GetWHUploadResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetWHUploadResponse) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *GetWHUploadResponse) GetDestinationId() string {
	if x != nil {
		return x.DestinationId
	}
	return ""
}

func (x *GetWHUploadResponse) GetDestinationType() string {
	if x != nil {
		return x.DestinationType
	}
	return ""
}

func (x *GetWHUploadResponse) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *GetWHUploadResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *GetWHUploadResponse) GetAttempt() int32 {
	if x != nil {
		return x.Attempt
	}
	return 0
}

func (x *GetWHUploadResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetWHUploadResponse) GetFirstEventAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FirstEventAt
	}
	return nil
}

func (x *GetWHUploadResponse) GetLastEventAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastEventAt
	}
	return nil
}

func (x *GetWHUploadResponse) GetNextRetryTime() string {
	if x != nil {
		return x.NextRetryTime
	}
	return ""
}

func (x *GetWHUploadResponse) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *GetWHUploadResponse) GetTables() []*WHTable {
	if x != nil {
		return x.Tables
	}
	return nil
}

type WHTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UploadId           int64                  `protobuf:"varint,2,opt,name=upload_id,json=uploadId,proto3" json:"upload_id,omitempty"`
	Name               string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Status             string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Error              string                 `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	LastExecAt         *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=last_exec_at,json=lastExecAt,proto3" json:"last_exec_at,omitempty"`
	Count              int32                  `protobuf:"varint,7,opt,name=count,proto3" json:"count,omitempty"`
	FirstScheduledTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=first_scheduled_time,json=firstScheduledTime,proto3" json:"first_scheduled_time,omitempty"`
}

func (x *WHTable) Reset() {
	*x = WHTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_warehouse_warehouse_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WHTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WHTable) ProtoMessage() {}

func (x *WHTable) ProtoReflect() protoreflect.Message {
	mi := &file_proto_warehouse_warehouse_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WHTable.ProtoReflect.Descriptor instead.
func (*WHTable) Descriptor() ([]byte, []int) {
	return file_proto_warehouse_warehouse_proto_rawDescGZIP(), []int{7}
}

func (x *WHTable) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WHTable) GetUploadId() int64 {
	if x != nil {
		return x.UploadId
	}
	return 0
}

func (x *WHTable) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WHTable) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *WHTable) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *WHTable) GetLastExecAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastExecAt
	}
	return nil
}

func (x *WHTable) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *WHTable) GetFirstScheduledTime() *timestamppb.Timestamp {
	if x != nil {
		return x.FirstScheduledTime
	}
	return nil
}

var File_proto_warehouse_warehouse_proto protoreflect.FileDescriptor

var file_proto_warehouse_warehouse_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73,
	0x65, 0x2f, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x57, 0x48, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xfd, 0x01, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x57, 0x48, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x15, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x5f,
	0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x31, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x57, 0x48, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x22, 0x50,
	0x0a, 0x0a, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x22, 0x7f, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x57, 0x48, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x48, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x31,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x3d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x57, 0x48, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x57, 0x48, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73,
	0x22, 0xe8, 0x03, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x57, 0x48, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x61,
	0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x74,
	0x74, 0x65, 0x6d, 0x70, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x40, 0x0a,
	0x0e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x12,
	0x3e, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x52, 0x65,
	0x74, 0x72, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x0d, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x48, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x22, 0x9a, 0x02, 0x0a, 0x07,
	0x57, 0x48, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x3c, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65,
	0x78, 0x65, 0x63, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x78,
	0x65, 0x63, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x4c, 0x0a, 0x14, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x66, 0x69, 0x72, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xe0, 0x01, 0x0a, 0x09, 0x57, 0x61, 0x72,
	0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x57, 0x48, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x57, 0x48, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x48,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x44, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x57, 0x48, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x48, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x48, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x57, 0x48, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x57, 0x48, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x48, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_warehouse_warehouse_proto_rawDescOnce sync.Once
	file_proto_warehouse_warehouse_proto_rawDescData = file_proto_warehouse_warehouse_proto_rawDesc
)

func file_proto_warehouse_warehouse_proto_rawDescGZIP() []byte {
	file_proto_warehouse_warehouse_proto_rawDescOnce.Do(func() {
		file_proto_warehouse_warehouse_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_warehouse_warehouse_proto_rawDescData)
	})
	return file_proto_warehouse_warehouse_proto_rawDescData
}

var file_proto_warehouse_warehouse_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_warehouse_warehouse_proto_goTypes = []interface{}{
	(*GetWHTablesRequest)(nil),    // 0: proto.GetWHTablesRequest
	(*GetWHUploadsRequest)(nil),   // 1: proto.GetWHUploadsRequest
	(*GetWHUploadRequest)(nil),    // 2: proto.GetWHUploadRequest
	(*Pagination)(nil),            // 3: proto.Pagination
	(*GetWHUploadsResponse)(nil),  // 4: proto.GetWHUploadsResponse
	(*GetWHTablesResponse)(nil),   // 5: proto.GetWHTablesResponse
	(*GetWHUploadResponse)(nil),   // 6: proto.GetWHUploadResponse
	(*WHTable)(nil),               // 7: proto.WHTable
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_proto_warehouse_warehouse_proto_depIdxs = []int32{
	6,  // 0: proto.GetWHUploadsResponse.uploads:type_name -> proto.GetWHUploadResponse
	3,  // 1: proto.GetWHUploadsResponse.pagination:type_name -> proto.Pagination
	7,  // 2: proto.GetWHTablesResponse.tables:type_name -> proto.WHTable
	8,  // 3: proto.GetWHUploadResponse.first_event_at:type_name -> google.protobuf.Timestamp
	8,  // 4: proto.GetWHUploadResponse.last_event_at:type_name -> google.protobuf.Timestamp
	7,  // 5: proto.GetWHUploadResponse.tables:type_name -> proto.WHTable
	8,  // 6: proto.WHTable.last_exec_at:type_name -> google.protobuf.Timestamp
	8,  // 7: proto.WHTable.first_scheduled_time:type_name -> google.protobuf.Timestamp
	1,  // 8: proto.Warehouse.GetWHUploads:input_type -> proto.GetWHUploadsRequest
	2,  // 9: proto.Warehouse.GetWHUpload:input_type -> proto.GetWHUploadRequest
	0,  // 10: proto.Warehouse.GetWHTables:input_type -> proto.GetWHTablesRequest
	4,  // 11: proto.Warehouse.GetWHUploads:output_type -> proto.GetWHUploadsResponse
	6,  // 12: proto.Warehouse.GetWHUpload:output_type -> proto.GetWHUploadResponse
	5,  // 13: proto.Warehouse.GetWHTables:output_type -> proto.GetWHTablesResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_warehouse_warehouse_proto_init() }
func file_proto_warehouse_warehouse_proto_init() {
	if File_proto_warehouse_warehouse_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_warehouse_warehouse_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWHTablesRequest); i {
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
		file_proto_warehouse_warehouse_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWHUploadsRequest); i {
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
		file_proto_warehouse_warehouse_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWHUploadRequest); i {
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
		file_proto_warehouse_warehouse_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_proto_warehouse_warehouse_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWHUploadsResponse); i {
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
		file_proto_warehouse_warehouse_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWHTablesResponse); i {
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
		file_proto_warehouse_warehouse_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWHUploadResponse); i {
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
		file_proto_warehouse_warehouse_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WHTable); i {
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
			RawDescriptor: file_proto_warehouse_warehouse_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_warehouse_warehouse_proto_goTypes,
		DependencyIndexes: file_proto_warehouse_warehouse_proto_depIdxs,
		MessageInfos:      file_proto_warehouse_warehouse_proto_msgTypes,
	}.Build()
	File_proto_warehouse_warehouse_proto = out.File
	file_proto_warehouse_warehouse_proto_rawDesc = nil
	file_proto_warehouse_warehouse_proto_goTypes = nil
	file_proto_warehouse_warehouse_proto_depIdxs = nil
}