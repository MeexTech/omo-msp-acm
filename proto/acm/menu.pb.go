// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.11.2
// source: proto/acm/menu.proto

package omo_msp_acm

import (
	proto "github.com/golang/protobuf/proto"
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

type PermissionAction int32

const (
	PermissionAction_Unknown PermissionAction = 0
	PermissionAction_Get     PermissionAction = 1
	PermissionAction_Post    PermissionAction = 2
	PermissionAction_Put     PermissionAction = 3
	PermissionAction_Delete  PermissionAction = 4
)

// Enum value maps for PermissionAction.
var (
	PermissionAction_name = map[int32]string{
		0: "Unknown",
		1: "Get",
		2: "Post",
		3: "Put",
		4: "Delete",
	}
	PermissionAction_value = map[string]int32{
		"Unknown": 0,
		"Get":     1,
		"Post":    2,
		"Put":     3,
		"Delete":  4,
	}
)

func (x PermissionAction) Enum() *PermissionAction {
	p := new(PermissionAction)
	*p = x
	return p
}

func (x PermissionAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PermissionAction) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_acm_menu_proto_enumTypes[0].Descriptor()
}

func (PermissionAction) Type() protoreflect.EnumType {
	return &file_proto_acm_menu_proto_enumTypes[0]
}

func (x PermissionAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PermissionAction.Descriptor instead.
func (PermissionAction) EnumDescriptor() ([]byte, []int) {
	return file_proto_acm_menu_proto_rawDescGZIP(), []int{0}
}

type MenuInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid     string           `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id      uint64           `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created int64            `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated int64            `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Name    string           `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Type    string           `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	Path    string           `protobuf:"bytes,7,opt,name=path,proto3" json:"path,omitempty"`
	Method  PermissionAction `protobuf:"varint,8,opt,name=method,proto3,enum=omo.msp.acm.PermissionAction" json:"method,omitempty"`
}

func (x *MenuInfo) Reset() {
	*x = MenuInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acm_menu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MenuInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MenuInfo) ProtoMessage() {}

func (x *MenuInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acm_menu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MenuInfo.ProtoReflect.Descriptor instead.
func (*MenuInfo) Descriptor() ([]byte, []int) {
	return file_proto_acm_menu_proto_rawDescGZIP(), []int{0}
}

func (x *MenuInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *MenuInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MenuInfo) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *MenuInfo) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *MenuInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MenuInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MenuInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *MenuInfo) GetMethod() PermissionAction {
	if x != nil {
		return x.Method
	}
	return PermissionAction_Unknown
}

type ReqMenuAdd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type   string           `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Path   string           `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Method PermissionAction `protobuf:"varint,4,opt,name=method,proto3,enum=omo.msp.acm.PermissionAction" json:"method,omitempty"`
}

func (x *ReqMenuAdd) Reset() {
	*x = ReqMenuAdd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acm_menu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqMenuAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqMenuAdd) ProtoMessage() {}

func (x *ReqMenuAdd) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acm_menu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqMenuAdd.ProtoReflect.Descriptor instead.
func (*ReqMenuAdd) Descriptor() ([]byte, []int) {
	return file_proto_acm_menu_proto_rawDescGZIP(), []int{1}
}

func (x *ReqMenuAdd) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReqMenuAdd) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ReqMenuAdd) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ReqMenuAdd) GetMethod() PermissionAction {
	if x != nil {
		return x.Method
	}
	return PermissionAction_Unknown
}

type ReplyMenuInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info   *MenuInfo    `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status ResultStatus `protobuf:"varint,2,opt,name=status,proto3,enum=omo.msp.acm.ResultStatus" json:"status,omitempty"`
}

func (x *ReplyMenuInfo) Reset() {
	*x = ReplyMenuInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acm_menu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyMenuInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyMenuInfo) ProtoMessage() {}

func (x *ReplyMenuInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acm_menu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyMenuInfo.ProtoReflect.Descriptor instead.
func (*ReplyMenuInfo) Descriptor() ([]byte, []int) {
	return file_proto_acm_menu_proto_rawDescGZIP(), []int{2}
}

func (x *ReplyMenuInfo) GetInfo() *MenuInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ReplyMenuInfo) GetStatus() ResultStatus {
	if x != nil {
		return x.Status
	}
	return ResultStatus_Success
}

type ReplyMenuList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32       `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	List   []*MenuInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ReplyMenuList) Reset() {
	*x = ReplyMenuList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acm_menu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyMenuList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyMenuList) ProtoMessage() {}

func (x *ReplyMenuList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acm_menu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyMenuList.ProtoReflect.Descriptor instead.
func (*ReplyMenuList) Descriptor() ([]byte, []int) {
	return file_proto_acm_menu_proto_rawDescGZIP(), []int{3}
}

func (x *ReplyMenuList) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *ReplyMenuList) GetList() []*MenuInfo {
	if x != nil {
		return x.List
	}
	return nil
}

type ReqMenuUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    string           `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name   string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type   string           `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Path   string           `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Method PermissionAction `protobuf:"varint,5,opt,name=method,proto3,enum=omo.msp.acm.PermissionAction" json:"method,omitempty"`
}

func (x *ReqMenuUpdate) Reset() {
	*x = ReqMenuUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acm_menu_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqMenuUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqMenuUpdate) ProtoMessage() {}

func (x *ReqMenuUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acm_menu_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqMenuUpdate.ProtoReflect.Descriptor instead.
func (*ReqMenuUpdate) Descriptor() ([]byte, []int) {
	return file_proto_acm_menu_proto_rawDescGZIP(), []int{4}
}

func (x *ReqMenuUpdate) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ReqMenuUpdate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReqMenuUpdate) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ReqMenuUpdate) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ReqMenuUpdate) GetMethod() PermissionAction {
	if x != nil {
		return x.Method
	}
	return PermissionAction_Unknown
}

var File_proto_acm_menu_proto protoreflect.FileDescriptor

var file_proto_acm_menu_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x6d, 0x2f, 0x6d, 0x65, 0x6e, 0x75,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e,
	0x61, 0x63, 0x6d, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x6d, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x08,
	0x4d, 0x65, 0x6e, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x35, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6f, 0x6d, 0x6f,
	0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x22, 0x7f, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x4d, 0x65, 0x6e, 0x75, 0x41, 0x64, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x35, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6f, 0x6d,
	0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x22, 0x6d, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x6e, 0x75, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x29, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e,
	0x4d, 0x65, 0x6e, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x31,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19,
	0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x52, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d,
	0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x94, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x4d, 0x65, 0x6e,
	0x75, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x35, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e,
	0x61, 0x63, 0x6d, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2a, 0x47, 0x0a, 0x10,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x10, 0x02,
	0x12, 0x07, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x10, 0x04, 0x32, 0xdb, 0x02, 0x0a, 0x0b, 0x4d, 0x65, 0x6e, 0x75, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x4f, 0x6e, 0x65, 0x12,
	0x17, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65,
	0x71, 0x4d, 0x65, 0x6e, 0x75, 0x41, 0x64, 0x64, 0x1a, 0x1a, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d,
	0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x6e, 0x75,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65,
	0x12, 0x18, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1a, 0x2e, 0x6f, 0x6d, 0x6f,
	0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65,
	0x6e, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x09, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x4f, 0x6e, 0x65, 0x12, 0x18, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e,
	0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a,
	0x16, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x12, 0x18, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63,
	0x6d, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1a, 0x2e,
	0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x4d, 0x65, 0x6e, 0x75, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x73, 0x65, 0x12, 0x1a, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e,
	0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x71, 0x4d, 0x65, 0x6e, 0x75, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x1a, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e,
	0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61,
	0x63, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_acm_menu_proto_rawDescOnce sync.Once
	file_proto_acm_menu_proto_rawDescData = file_proto_acm_menu_proto_rawDesc
)

func file_proto_acm_menu_proto_rawDescGZIP() []byte {
	file_proto_acm_menu_proto_rawDescOnce.Do(func() {
		file_proto_acm_menu_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_acm_menu_proto_rawDescData)
	})
	return file_proto_acm_menu_proto_rawDescData
}

var file_proto_acm_menu_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_acm_menu_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_acm_menu_proto_goTypes = []interface{}{
	(PermissionAction)(0), // 0: omo.msp.acm.PermissionAction
	(*MenuInfo)(nil),      // 1: omo.msp.acm.MenuInfo
	(*ReqMenuAdd)(nil),    // 2: omo.msp.acm.ReqMenuAdd
	(*ReplyMenuInfo)(nil), // 3: omo.msp.acm.ReplyMenuInfo
	(*ReplyMenuList)(nil), // 4: omo.msp.acm.ReplyMenuList
	(*ReqMenuUpdate)(nil), // 5: omo.msp.acm.ReqMenuUpdate
	(ResultStatus)(0),     // 6: omo.msp.acm.ResultStatus
	(*RequestInfo)(nil),   // 7: omo.msp.acm.RequestInfo
	(*ReplyInfo)(nil),     // 8: omo.msp.acm.ReplyInfo
}
var file_proto_acm_menu_proto_depIdxs = []int32{
	0,  // 0: omo.msp.acm.MenuInfo.method:type_name -> omo.msp.acm.PermissionAction
	0,  // 1: omo.msp.acm.ReqMenuAdd.method:type_name -> omo.msp.acm.PermissionAction
	1,  // 2: omo.msp.acm.ReplyMenuInfo.info:type_name -> omo.msp.acm.MenuInfo
	6,  // 3: omo.msp.acm.ReplyMenuInfo.status:type_name -> omo.msp.acm.ResultStatus
	1,  // 4: omo.msp.acm.ReplyMenuList.list:type_name -> omo.msp.acm.MenuInfo
	0,  // 5: omo.msp.acm.ReqMenuUpdate.method:type_name -> omo.msp.acm.PermissionAction
	2,  // 6: omo.msp.acm.MenuService.AddOne:input_type -> omo.msp.acm.ReqMenuAdd
	7,  // 7: omo.msp.acm.MenuService.GetOne:input_type -> omo.msp.acm.RequestInfo
	7,  // 8: omo.msp.acm.MenuService.RemoveOne:input_type -> omo.msp.acm.RequestInfo
	7,  // 9: omo.msp.acm.MenuService.GetAll:input_type -> omo.msp.acm.RequestInfo
	5,  // 10: omo.msp.acm.MenuService.UpdateBase:input_type -> omo.msp.acm.ReqMenuUpdate
	3,  // 11: omo.msp.acm.MenuService.AddOne:output_type -> omo.msp.acm.ReplyMenuInfo
	3,  // 12: omo.msp.acm.MenuService.GetOne:output_type -> omo.msp.acm.ReplyMenuInfo
	8,  // 13: omo.msp.acm.MenuService.RemoveOne:output_type -> omo.msp.acm.ReplyInfo
	4,  // 14: omo.msp.acm.MenuService.GetAll:output_type -> omo.msp.acm.ReplyMenuList
	3,  // 15: omo.msp.acm.MenuService.UpdateBase:output_type -> omo.msp.acm.ReplyMenuInfo
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_acm_menu_proto_init() }
func file_proto_acm_menu_proto_init() {
	if File_proto_acm_menu_proto != nil {
		return
	}
	file_proto_acm_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_acm_menu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MenuInfo); i {
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
		file_proto_acm_menu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqMenuAdd); i {
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
		file_proto_acm_menu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyMenuInfo); i {
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
		file_proto_acm_menu_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyMenuList); i {
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
		file_proto_acm_menu_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqMenuUpdate); i {
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
			RawDescriptor: file_proto_acm_menu_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_acm_menu_proto_goTypes,
		DependencyIndexes: file_proto_acm_menu_proto_depIdxs,
		EnumInfos:         file_proto_acm_menu_proto_enumTypes,
		MessageInfos:      file_proto_acm_menu_proto_msgTypes,
	}.Build()
	File_proto_acm_menu_proto = out.File
	file_proto_acm_menu_proto_rawDesc = nil
	file_proto_acm_menu_proto_goTypes = nil
	file_proto_acm_menu_proto_depIdxs = nil
}
