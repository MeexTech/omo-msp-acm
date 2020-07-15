// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.11.2
// source: proto/acm/user.proto

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

type UserLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid   string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	User  string   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Roles []string `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *UserLink) Reset() {
	*x = UserLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acm_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLink) ProtoMessage() {}

func (x *UserLink) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acm_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLink.ProtoReflect.Descriptor instead.
func (*UserLink) Descriptor() ([]byte, []int) {
	return file_proto_acm_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserLink) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *UserLink) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *UserLink) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type ReqUserAdd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Operator string   `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	Roles    []string `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *ReqUserAdd) Reset() {
	*x = ReqUserAdd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acm_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqUserAdd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqUserAdd) ProtoMessage() {}

func (x *ReqUserAdd) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acm_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqUserAdd.ProtoReflect.Descriptor instead.
func (*ReqUserAdd) Descriptor() ([]byte, []int) {
	return file_proto_acm_user_proto_rawDescGZIP(), []int{1}
}

func (x *ReqUserAdd) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ReqUserAdd) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *ReqUserAdd) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type ReplyUserLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info   *UserLink    `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status ResultStatus `protobuf:"varint,2,opt,name=status,proto3,enum=omo.msp.acm.ResultStatus" json:"status,omitempty"`
}

func (x *ReplyUserLink) Reset() {
	*x = ReplyUserLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acm_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyUserLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyUserLink) ProtoMessage() {}

func (x *ReplyUserLink) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acm_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyUserLink.ProtoReflect.Descriptor instead.
func (*ReplyUserLink) Descriptor() ([]byte, []int) {
	return file_proto_acm_user_proto_rawDescGZIP(), []int{2}
}

func (x *ReplyUserLink) GetInfo() *UserLink {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ReplyUserLink) GetStatus() ResultStatus {
	if x != nil {
		return x.Status
	}
	return ResultStatus_Success
}

type ReqUserList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Number int32 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *ReqUserList) Reset() {
	*x = ReqUserList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acm_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqUserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqUserList) ProtoMessage() {}

func (x *ReqUserList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acm_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqUserList.ProtoReflect.Descriptor instead.
func (*ReqUserList) Descriptor() ([]byte, []int) {
	return file_proto_acm_user_proto_rawDescGZIP(), []int{3}
}

func (x *ReqUserList) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ReqUserList) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type ReplyUserList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserLink `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *ReplyUserList) Reset() {
	*x = ReplyUserList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acm_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyUserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyUserList) ProtoMessage() {}

func (x *ReplyUserList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acm_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyUserList.ProtoReflect.Descriptor instead.
func (*ReplyUserList) Descriptor() ([]byte, []int) {
	return file_proto_acm_user_proto_rawDescGZIP(), []int{4}
}

func (x *ReplyUserList) GetUsers() []*UserLink {
	if x != nil {
		return x.Users
	}
	return nil
}

type ReplyUserPermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User       string       `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Permission bool         `protobuf:"varint,2,opt,name=permission,proto3" json:"permission,omitempty"`
	Status     ResultStatus `protobuf:"varint,3,opt,name=status,proto3,enum=omo.msp.acm.ResultStatus" json:"status,omitempty"`
}

func (x *ReplyUserPermission) Reset() {
	*x = ReplyUserPermission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acm_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyUserPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyUserPermission) ProtoMessage() {}

func (x *ReplyUserPermission) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acm_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyUserPermission.ProtoReflect.Descriptor instead.
func (*ReplyUserPermission) Descriptor() ([]byte, []int) {
	return file_proto_acm_user_proto_rawDescGZIP(), []int{5}
}

func (x *ReplyUserPermission) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ReplyUserPermission) GetPermission() bool {
	if x != nil {
		return x.Permission
	}
	return false
}

func (x *ReplyUserPermission) GetStatus() ResultStatus {
	if x != nil {
		return x.Status
	}
	return ResultStatus_Success
}

type ReqUserPermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Path   string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Action string `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *ReqUserPermission) Reset() {
	*x = ReqUserPermission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acm_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqUserPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqUserPermission) ProtoMessage() {}

func (x *ReqUserPermission) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acm_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqUserPermission.ProtoReflect.Descriptor instead.
func (*ReqUserPermission) Descriptor() ([]byte, []int) {
	return file_proto_acm_user_proto_rawDescGZIP(), []int{6}
}

func (x *ReqUserPermission) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ReqUserPermission) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ReqUserPermission) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type ReqLinkRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Role     string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Operator string `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
}

func (x *ReqLinkRole) Reset() {
	*x = ReqLinkRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acm_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqLinkRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqLinkRole) ProtoMessage() {}

func (x *ReqLinkRole) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acm_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqLinkRole.ProtoReflect.Descriptor instead.
func (*ReqLinkRole) Descriptor() ([]byte, []int) {
	return file_proto_acm_user_proto_rawDescGZIP(), []int{7}
}

func (x *ReqLinkRole) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ReqLinkRole) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *ReqLinkRole) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

type ReplyLinkRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   string       `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Roles  []string     `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	Status ResultStatus `protobuf:"varint,3,opt,name=status,proto3,enum=omo.msp.acm.ResultStatus" json:"status,omitempty"`
}

func (x *ReplyLinkRole) Reset() {
	*x = ReplyLinkRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_acm_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyLinkRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyLinkRole) ProtoMessage() {}

func (x *ReplyLinkRole) ProtoReflect() protoreflect.Message {
	mi := &file_proto_acm_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyLinkRole.ProtoReflect.Descriptor instead.
func (*ReplyLinkRole) Descriptor() ([]byte, []int) {
	return file_proto_acm_user_proto_rawDescGZIP(), []int{8}
}

func (x *ReplyLinkRole) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ReplyLinkRole) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *ReplyLinkRole) GetStatus() ResultStatus {
	if x != nil {
		return x.Status
	}
	return ResultStatus_Success
}

var File_proto_acm_user_proto protoreflect.FileDescriptor

var file_proto_acm_user_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x6d, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e,
	0x61, 0x63, 0x6d, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x6d, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x22, 0x52, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0x6d, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x29, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70,
	0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x04, 0x69,
	0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63,
	0x6d, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x39, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x3c, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x2b, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22,
	0x7c, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6f, 0x6d, 0x6f,
	0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x53, 0x0a,
	0x11, 0x52, 0x65, 0x71, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x51, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x6f, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x6c, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4c, 0x69,
	0x6e, 0x6b, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x12, 0x31, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x32, 0xf6, 0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x4f, 0x6e, 0x65, 0x12, 0x17, 0x2e,
	0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x71, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x1a, 0x1a, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70,
	0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69,
	0x6e, 0x6b, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x12, 0x18,
	0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1a, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d,
	0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x09, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x4f, 0x6e, 0x65, 0x12, 0x18, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63,
	0x6d, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e,
	0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x18, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d,
	0x2e, 0x52, 0x65, 0x71, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6f,
	0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0c, 0x49, 0x73,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x6f, 0x6d, 0x6f,
	0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x71, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x20, 0x2e, 0x6f, 0x6d, 0x6f,
	0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x44,
	0x0a, 0x0a, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x6f,
	0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x71, 0x4c, 0x69,
	0x6e, 0x6b, 0x52, 0x6f, 0x6c, 0x65, 0x1a, 0x1a, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70,
	0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x6f,
	0x6c, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61,
	0x63, 0x6d, 0x2e, 0x52, 0x65, 0x71, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x6f, 0x6c, 0x65, 0x1a, 0x1a,
	0x2e, 0x6f, 0x6d, 0x6f, 0x2e, 0x6d, 0x73, 0x70, 0x2e, 0x61, 0x63, 0x6d, 0x2e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x6d, 0x3b, 0x6f, 0x6d, 0x6f, 0x5f, 0x6d, 0x73,
	0x70, 0x5f, 0x61, 0x63, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_acm_user_proto_rawDescOnce sync.Once
	file_proto_acm_user_proto_rawDescData = file_proto_acm_user_proto_rawDesc
)

func file_proto_acm_user_proto_rawDescGZIP() []byte {
	file_proto_acm_user_proto_rawDescOnce.Do(func() {
		file_proto_acm_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_acm_user_proto_rawDescData)
	})
	return file_proto_acm_user_proto_rawDescData
}

var file_proto_acm_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_acm_user_proto_goTypes = []interface{}{
	(*UserLink)(nil),            // 0: omo.msp.acm.UserLink
	(*ReqUserAdd)(nil),          // 1: omo.msp.acm.ReqUserAdd
	(*ReplyUserLink)(nil),       // 2: omo.msp.acm.ReplyUserLink
	(*ReqUserList)(nil),         // 3: omo.msp.acm.ReqUserList
	(*ReplyUserList)(nil),       // 4: omo.msp.acm.ReplyUserList
	(*ReplyUserPermission)(nil), // 5: omo.msp.acm.ReplyUserPermission
	(*ReqUserPermission)(nil),   // 6: omo.msp.acm.ReqUserPermission
	(*ReqLinkRole)(nil),         // 7: omo.msp.acm.ReqLinkRole
	(*ReplyLinkRole)(nil),       // 8: omo.msp.acm.ReplyLinkRole
	(ResultStatus)(0),           // 9: omo.msp.acm.ResultStatus
	(*RequestInfo)(nil),         // 10: omo.msp.acm.RequestInfo
	(*ReplyInfo)(nil),           // 11: omo.msp.acm.ReplyInfo
}
var file_proto_acm_user_proto_depIdxs = []int32{
	0,  // 0: omo.msp.acm.ReplyUserLink.info:type_name -> omo.msp.acm.UserLink
	9,  // 1: omo.msp.acm.ReplyUserLink.status:type_name -> omo.msp.acm.ResultStatus
	0,  // 2: omo.msp.acm.ReplyUserList.users:type_name -> omo.msp.acm.UserLink
	9,  // 3: omo.msp.acm.ReplyUserPermission.status:type_name -> omo.msp.acm.ResultStatus
	9,  // 4: omo.msp.acm.ReplyLinkRole.status:type_name -> omo.msp.acm.ResultStatus
	1,  // 5: omo.msp.acm.UserService.AddOne:input_type -> omo.msp.acm.ReqUserAdd
	10, // 6: omo.msp.acm.UserService.GetOne:input_type -> omo.msp.acm.RequestInfo
	10, // 7: omo.msp.acm.UserService.RemoveOne:input_type -> omo.msp.acm.RequestInfo
	3,  // 8: omo.msp.acm.UserService.GetList:input_type -> omo.msp.acm.ReqUserList
	6,  // 9: omo.msp.acm.UserService.IsPermission:input_type -> omo.msp.acm.ReqUserPermission
	7,  // 10: omo.msp.acm.UserService.AppendRole:input_type -> omo.msp.acm.ReqLinkRole
	7,  // 11: omo.msp.acm.UserService.SubtractRole:input_type -> omo.msp.acm.ReqLinkRole
	2,  // 12: omo.msp.acm.UserService.AddOne:output_type -> omo.msp.acm.ReplyUserLink
	2,  // 13: omo.msp.acm.UserService.GetOne:output_type -> omo.msp.acm.ReplyUserLink
	11, // 14: omo.msp.acm.UserService.RemoveOne:output_type -> omo.msp.acm.ReplyInfo
	4,  // 15: omo.msp.acm.UserService.GetList:output_type -> omo.msp.acm.ReplyUserList
	5,  // 16: omo.msp.acm.UserService.IsPermission:output_type -> omo.msp.acm.ReplyUserPermission
	8,  // 17: omo.msp.acm.UserService.AppendRole:output_type -> omo.msp.acm.ReplyLinkRole
	8,  // 18: omo.msp.acm.UserService.SubtractRole:output_type -> omo.msp.acm.ReplyLinkRole
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_acm_user_proto_init() }
func file_proto_acm_user_proto_init() {
	if File_proto_acm_user_proto != nil {
		return
	}
	file_proto_acm_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_acm_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLink); i {
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
		file_proto_acm_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqUserAdd); i {
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
		file_proto_acm_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyUserLink); i {
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
		file_proto_acm_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqUserList); i {
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
		file_proto_acm_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyUserList); i {
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
		file_proto_acm_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyUserPermission); i {
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
		file_proto_acm_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqUserPermission); i {
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
		file_proto_acm_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqLinkRole); i {
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
		file_proto_acm_user_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyLinkRole); i {
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
			RawDescriptor: file_proto_acm_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_acm_user_proto_goTypes,
		DependencyIndexes: file_proto_acm_user_proto_depIdxs,
		MessageInfos:      file_proto_acm_user_proto_msgTypes,
	}.Build()
	File_proto_acm_user_proto = out.File
	file_proto_acm_user_proto_rawDesc = nil
	file_proto_acm_user_proto_goTypes = nil
	file_proto_acm_user_proto_depIdxs = nil
}
