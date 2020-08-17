// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: cron.proto

package cron

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 列表页 请求参数
type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *SearchPage `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cron_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cron_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_cron_proto_rawDescGZIP(), []int{0}
}

func (x *ListRequest) GetPage() *SearchPage {
	if x != nil {
		return x.Page
	}
	return nil
}

// 列表页返回参数
type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*CronList `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cron_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cron_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_cron_proto_rawDescGZIP(), []int{1}
}

func (x *ListResponse) GetItems() []*CronList {
	if x != nil {
		return x.Items
	}
	return nil
}

// 添加的请求参数
type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	R *CronBase `protobuf:"bytes,1,opt,name=r,proto3" json:"r,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cron_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cron_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_cron_proto_rawDescGZIP(), []int{2}
}

func (x *AddRequest) GetR() *CronBase {
	if x != nil {
		return x.R
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	R  *CronBase `protobuf:"bytes,2,opt,name=r,proto3" json:"r,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cron_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cron_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_cron_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateRequest) GetR() *CronBase {
	if x != nil {
		return x.R
	}
	return nil
}

type KillerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *KillerRequest) Reset() {
	*x = KillerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cron_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KillerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KillerRequest) ProtoMessage() {}

func (x *KillerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cron_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KillerRequest.ProtoReflect.Descriptor instead.
func (*KillerRequest) Descriptor() ([]byte, []int) {
	return file_cron_proto_rawDescGZIP(), []int{4}
}

func (x *KillerRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdRequest) Reset() {
	*x = IdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cron_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cron_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdRequest.ProtoReflect.Descriptor instead.
func (*IdRequest) Descriptor() ([]byte, []int) {
	return file_cron_proto_rawDescGZIP(), []int{5}
}

func (x *IdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 请求的基础数据
type CronBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Exp       string `protobuf:"bytes,2,opt,name=exp,proto3" json:"exp,omitempty"`
	Command   string `protobuf:"bytes,3,opt,name=command,proto3" json:"command,omitempty"`
	Desc      string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	IsKiller  int32  `protobuf:"varint,5,opt,name=is_killer,json=isKiller,proto3" json:"is_killer,omitempty"`
	RunAt     string `protobuf:"bytes,6,opt,name=run_at,json=runAt,proto3" json:"run_at,omitempty"`
	CreatedAt string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Id        int32  `protobuf:"varint,9,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CronBase) Reset() {
	*x = CronBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cron_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CronBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CronBase) ProtoMessage() {}

func (x *CronBase) ProtoReflect() protoreflect.Message {
	mi := &file_cron_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CronBase.ProtoReflect.Descriptor instead.
func (*CronBase) Descriptor() ([]byte, []int) {
	return file_cron_proto_rawDescGZIP(), []int{6}
}

func (x *CronBase) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CronBase) GetExp() string {
	if x != nil {
		return x.Exp
	}
	return ""
}

func (x *CronBase) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *CronBase) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *CronBase) GetIsKiller() int32 {
	if x != nil {
		return x.IsKiller
	}
	return 0
}

func (x *CronBase) GetRunAt() string {
	if x != nil {
		return x.RunAt
	}
	return ""
}

func (x *CronBase) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CronBase) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *CronBase) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

//  返回数据的基础数据
type CronList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Exp       string `protobuf:"bytes,3,opt,name=exp,proto3" json:"exp,omitempty"`
	Command   string `protobuf:"bytes,4,opt,name=command,proto3" json:"command,omitempty"`
	Desc      string `protobuf:"bytes,5,opt,name=desc,proto3" json:"desc,omitempty"`
	CreatedAt string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	IsKiller  int32  `protobuf:"varint,8,opt,name=is_killer,json=isKiller,proto3" json:"is_killer,omitempty"`
}

func (x *CronList) Reset() {
	*x = CronList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cron_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CronList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CronList) ProtoMessage() {}

func (x *CronList) ProtoReflect() protoreflect.Message {
	mi := &file_cron_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CronList.ProtoReflect.Descriptor instead.
func (*CronList) Descriptor() ([]byte, []int) {
	return file_cron_proto_rawDescGZIP(), []int{7}
}

func (x *CronList) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CronList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CronList) GetExp() string {
	if x != nil {
		return x.Exp
	}
	return ""
}

func (x *CronList) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *CronList) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *CronList) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CronList) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *CronList) GetIsKiller() int32 {
	if x != nil {
		return x.IsKiller
	}
	return 0
}

// 请求分页
type SearchPage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (x *SearchPage) Reset() {
	*x = SearchPage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cron_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchPage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPage) ProtoMessage() {}

func (x *SearchPage) ProtoReflect() protoreflect.Message {
	mi := &file_cron_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPage.ProtoReflect.Descriptor instead.
func (*SearchPage) Descriptor() ([]byte, []int) {
	return file_cron_proto_rawDescGZIP(), []int{8}
}

func (x *SearchPage) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *SearchPage) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// 不需要返回参数的数据
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cron_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_cron_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_cron_proto_rawDescGZIP(), []int{9}
}

var File_cron_proto protoreflect.FileDescriptor

var file_cron_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x72,
	0x6f, 0x6e, 0x22, 0x33, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x67,
	0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x34, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x43, 0x72,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x2a, 0x0a,
	0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x01, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x43, 0x72,
	0x6f, 0x6e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x01, 0x72, 0x22, 0x3d, 0x0a, 0x0d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x01, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x6f,
	0x6e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x01, 0x72, 0x22, 0x1f, 0x0a, 0x0d, 0x4b, 0x69, 0x6c, 0x6c,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1b, 0x0a, 0x09, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xe0, 0x01, 0x0a, 0x08, 0x43, 0x72, 0x6f, 0x6e, 0x42,
	0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6b, 0x69,
	0x6c, 0x6c, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x73, 0x4b, 0x69,
	0x6c, 0x6c, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f, 0x61, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xc9, 0x01, 0x0a, 0x08, 0x43, 0x72,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6b,
	0x69, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x73, 0x4b,
	0x69, 0x6c, 0x6c, 0x65, 0x72, 0x22, 0x3d, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x91, 0x02,
	0x0a, 0x04, 0x43, 0x72, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x05, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12,
	0x11, 0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0f, 0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x43,
	0x72, 0x6f, 0x6e, 0x42, 0x61, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x10,
	0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0b, 0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2a, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x63,
	0x72, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x06, 0x4b, 0x69, 0x6c,
	0x6c, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x4b, 0x69, 0x6c, 0x6c, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x06, 0x4f, 0x6e, 0x4c, 0x69, 0x6e, 0x65, 0x12,
	0x13, 0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x4b, 0x69, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x63, 0x72, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cron_proto_rawDescOnce sync.Once
	file_cron_proto_rawDescData = file_cron_proto_rawDesc
)

func file_cron_proto_rawDescGZIP() []byte {
	file_cron_proto_rawDescOnce.Do(func() {
		file_cron_proto_rawDescData = protoimpl.X.CompressGZIP(file_cron_proto_rawDescData)
	})
	return file_cron_proto_rawDescData
}

var file_cron_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_cron_proto_goTypes = []interface{}{
	(*ListRequest)(nil),   // 0: cron.ListRequest
	(*ListResponse)(nil),  // 1: cron.ListResponse
	(*AddRequest)(nil),    // 2: cron.AddRequest
	(*UpdateRequest)(nil), // 3: cron.UpdateRequest
	(*KillerRequest)(nil), // 4: cron.KillerRequest
	(*IdRequest)(nil),     // 5: cron.IdRequest
	(*CronBase)(nil),      // 6: cron.CronBase
	(*CronList)(nil),      // 7: cron.CronList
	(*SearchPage)(nil),    // 8: cron.SearchPage
	(*Error)(nil),         // 9: cron.Error
}
var file_cron_proto_depIdxs = []int32{
	8,  // 0: cron.ListRequest.page:type_name -> cron.SearchPage
	7,  // 1: cron.ListResponse.items:type_name -> cron.CronList
	6,  // 2: cron.AddRequest.r:type_name -> cron.CronBase
	6,  // 3: cron.UpdateRequest.r:type_name -> cron.CronBase
	0,  // 4: cron.Cron.Lists:input_type -> cron.ListRequest
	5,  // 5: cron.Cron.GetFirstById:input_type -> cron.IdRequest
	2,  // 6: cron.Cron.Add:input_type -> cron.AddRequest
	3,  // 7: cron.Cron.Update:input_type -> cron.UpdateRequest
	4,  // 8: cron.Cron.Killer:input_type -> cron.KillerRequest
	4,  // 9: cron.Cron.OnLine:input_type -> cron.KillerRequest
	1,  // 10: cron.Cron.Lists:output_type -> cron.ListResponse
	6,  // 11: cron.Cron.GetFirstById:output_type -> cron.CronBase
	9,  // 12: cron.Cron.Add:output_type -> cron.Error
	9,  // 13: cron.Cron.Update:output_type -> cron.Error
	9,  // 14: cron.Cron.Killer:output_type -> cron.Error
	9,  // 15: cron.Cron.OnLine:output_type -> cron.Error
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_cron_proto_init() }
func file_cron_proto_init() {
	if File_cron_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cron_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_cron_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_cron_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_cron_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_cron_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KillerRequest); i {
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
		file_cron_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdRequest); i {
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
		file_cron_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CronBase); i {
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
		file_cron_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CronList); i {
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
		file_cron_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchPage); i {
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
		file_cron_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_cron_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cron_proto_goTypes,
		DependencyIndexes: file_cron_proto_depIdxs,
		MessageInfos:      file_cron_proto_msgTypes,
	}.Build()
	File_cron_proto = out.File
	file_cron_proto_rawDesc = nil
	file_cron_proto_goTypes = nil
	file_cron_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CronClient is the client API for Cron service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CronClient interface {
	// 获取列表页
	Lists(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	GetFirstById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*CronBase, error)
	// 添加数据
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*Error, error)
	// 更改数据
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Error, error)
	// 杀死数据
	Killer(ctx context.Context, in *KillerRequest, opts ...grpc.CallOption) (*Error, error)
	// 启动杀死的数据
	OnLine(ctx context.Context, in *KillerRequest, opts ...grpc.CallOption) (*Error, error)
}

type cronClient struct {
	cc grpc.ClientConnInterface
}

func NewCronClient(cc grpc.ClientConnInterface) CronClient {
	return &cronClient{cc}
}

func (c *cronClient) Lists(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/cron.Cron/Lists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronClient) GetFirstById(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*CronBase, error) {
	out := new(CronBase)
	err := c.cc.Invoke(ctx, "/cron.Cron/GetFirstById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/cron.Cron/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/cron.Cron/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronClient) Killer(ctx context.Context, in *KillerRequest, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/cron.Cron/Killer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cronClient) OnLine(ctx context.Context, in *KillerRequest, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/cron.Cron/OnLine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CronServer is the server API for Cron service.
type CronServer interface {
	// 获取列表页
	Lists(context.Context, *ListRequest) (*ListResponse, error)
	GetFirstById(context.Context, *IdRequest) (*CronBase, error)
	// 添加数据
	Add(context.Context, *AddRequest) (*Error, error)
	// 更改数据
	Update(context.Context, *UpdateRequest) (*Error, error)
	// 杀死数据
	Killer(context.Context, *KillerRequest) (*Error, error)
	// 启动杀死的数据
	OnLine(context.Context, *KillerRequest) (*Error, error)
}

// UnimplementedCronServer can be embedded to have forward compatible implementations.
type UnimplementedCronServer struct {
}

func (*UnimplementedCronServer) Lists(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lists not implemented")
}
func (*UnimplementedCronServer) GetFirstById(context.Context, *IdRequest) (*CronBase, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFirstById not implemented")
}
func (*UnimplementedCronServer) Add(context.Context, *AddRequest) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedCronServer) Update(context.Context, *UpdateRequest) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedCronServer) Killer(context.Context, *KillerRequest) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Killer not implemented")
}
func (*UnimplementedCronServer) OnLine(context.Context, *KillerRequest) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnLine not implemented")
}

func RegisterCronServer(s *grpc.Server, srv CronServer) {
	s.RegisterService(&_Cron_serviceDesc, srv)
}

func _Cron_Lists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronServer).Lists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cron.Cron/Lists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronServer).Lists(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cron_GetFirstById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronServer).GetFirstById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cron.Cron/GetFirstById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronServer).GetFirstById(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cron_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cron.Cron/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cron_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cron.Cron/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cron_Killer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KillerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronServer).Killer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cron.Cron/Killer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronServer).Killer(ctx, req.(*KillerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cron_OnLine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KillerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CronServer).OnLine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cron.Cron/OnLine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CronServer).OnLine(ctx, req.(*KillerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cron_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cron.Cron",
	HandlerType: (*CronServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Lists",
			Handler:    _Cron_Lists_Handler,
		},
		{
			MethodName: "GetFirstById",
			Handler:    _Cron_GetFirstById_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Cron_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Cron_Update_Handler,
		},
		{
			MethodName: "Killer",
			Handler:    _Cron_Killer_Handler,
		},
		{
			MethodName: "OnLine",
			Handler:    _Cron_OnLine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cron.proto",
}
