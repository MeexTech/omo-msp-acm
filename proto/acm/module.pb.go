// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/acm/module.proto

package omo_msp_acm

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ModuleInfo struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64    `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64    `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty"`
	Creator              string   `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,8,opt,name=operator,proto3" json:"operator,omitempty"`
	Owner                string   `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
	Menus                []string `protobuf:"bytes,10,rep,name=menus,proto3" json:"menus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModuleInfo) Reset()         { *m = ModuleInfo{} }
func (m *ModuleInfo) String() string { return proto.CompactTextString(m) }
func (*ModuleInfo) ProtoMessage()    {}
func (*ModuleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef308d0e5e1a8b70, []int{0}
}

func (m *ModuleInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModuleInfo.Unmarshal(m, b)
}
func (m *ModuleInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModuleInfo.Marshal(b, m, deterministic)
}
func (m *ModuleInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModuleInfo.Merge(m, src)
}
func (m *ModuleInfo) XXX_Size() int {
	return xxx_messageInfo_ModuleInfo.Size(m)
}
func (m *ModuleInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ModuleInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ModuleInfo proto.InternalMessageInfo

func (m *ModuleInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ModuleInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ModuleInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ModuleInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *ModuleInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModuleInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ModuleInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ModuleInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ModuleInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ModuleInfo) GetMenus() []string {
	if m != nil {
		return m.Menus
	}
	return nil
}

type ReqModuleAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	Owner                string   `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Menus                []string `protobuf:"bytes,5,rep,name=menus,proto3" json:"menus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqModuleAdd) Reset()         { *m = ReqModuleAdd{} }
func (m *ReqModuleAdd) String() string { return proto.CompactTextString(m) }
func (*ReqModuleAdd) ProtoMessage()    {}
func (*ReqModuleAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef308d0e5e1a8b70, []int{1}
}

func (m *ReqModuleAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqModuleAdd.Unmarshal(m, b)
}
func (m *ReqModuleAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqModuleAdd.Marshal(b, m, deterministic)
}
func (m *ReqModuleAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqModuleAdd.Merge(m, src)
}
func (m *ReqModuleAdd) XXX_Size() int {
	return xxx_messageInfo_ReqModuleAdd.Size(m)
}
func (m *ReqModuleAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqModuleAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqModuleAdd proto.InternalMessageInfo

func (m *ReqModuleAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqModuleAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqModuleAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqModuleAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqModuleAdd) GetMenus() []string {
	if m != nil {
		return m.Menus
	}
	return nil
}

type ReplyModuleInfo struct {
	Info                 *ModuleInfo  `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyModuleInfo) Reset()         { *m = ReplyModuleInfo{} }
func (m *ReplyModuleInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyModuleInfo) ProtoMessage()    {}
func (*ReplyModuleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef308d0e5e1a8b70, []int{2}
}

func (m *ReplyModuleInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyModuleInfo.Unmarshal(m, b)
}
func (m *ReplyModuleInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyModuleInfo.Marshal(b, m, deterministic)
}
func (m *ReplyModuleInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyModuleInfo.Merge(m, src)
}
func (m *ReplyModuleInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyModuleInfo.Size(m)
}
func (m *ReplyModuleInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyModuleInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyModuleInfo proto.InternalMessageInfo

func (m *ReplyModuleInfo) GetInfo() *ModuleInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyModuleInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyModuleList struct {
	Number               int32         `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Status               *ReplyStatus  `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*ModuleInfo `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReplyModuleList) Reset()         { *m = ReplyModuleList{} }
func (m *ReplyModuleList) String() string { return proto.CompactTextString(m) }
func (*ReplyModuleList) ProtoMessage()    {}
func (*ReplyModuleList) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef308d0e5e1a8b70, []int{3}
}

func (m *ReplyModuleList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyModuleList.Unmarshal(m, b)
}
func (m *ReplyModuleList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyModuleList.Marshal(b, m, deterministic)
}
func (m *ReplyModuleList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyModuleList.Merge(m, src)
}
func (m *ReplyModuleList) XXX_Size() int {
	return xxx_messageInfo_ReplyModuleList.Size(m)
}
func (m *ReplyModuleList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyModuleList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyModuleList proto.InternalMessageInfo

func (m *ReplyModuleList) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyModuleList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyModuleList) GetList() []*ModuleInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type ReqModuleUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqModuleUpdate) Reset()         { *m = ReqModuleUpdate{} }
func (m *ReqModuleUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqModuleUpdate) ProtoMessage()    {}
func (*ReqModuleUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef308d0e5e1a8b70, []int{4}
}

func (m *ReqModuleUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqModuleUpdate.Unmarshal(m, b)
}
func (m *ReqModuleUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqModuleUpdate.Marshal(b, m, deterministic)
}
func (m *ReqModuleUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqModuleUpdate.Merge(m, src)
}
func (m *ReqModuleUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqModuleUpdate.Size(m)
}
func (m *ReqModuleUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqModuleUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqModuleUpdate proto.InternalMessageInfo

func (m *ReqModuleUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqModuleUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqModuleUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqModuleUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func init() {
	proto.RegisterType((*ModuleInfo)(nil), "omo.msp.acm.ModuleInfo")
	proto.RegisterType((*ReqModuleAdd)(nil), "omo.msp.acm.ReqModuleAdd")
	proto.RegisterType((*ReplyModuleInfo)(nil), "omo.msp.acm.ReplyModuleInfo")
	proto.RegisterType((*ReplyModuleList)(nil), "omo.msp.acm.ReplyModuleList")
	proto.RegisterType((*ReqModuleUpdate)(nil), "omo.msp.acm.ReqModuleUpdate")
}

func init() {
	proto.RegisterFile("proto/acm/module.proto", fileDescriptor_ef308d0e5e1a8b70)
}

var fileDescriptor_ef308d0e5e1a8b70 = []byte{
	// 501 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xeb, 0x8f, 0xb8, 0xcd, 0x84, 0x2f, 0xad, 0x20, 0x5d, 0xa2, 0x1e, 0x2c, 0x9f, 0x22,
	0x21, 0xb9, 0x28, 0x1c, 0x39, 0xa0, 0x04, 0x24, 0x84, 0x44, 0x05, 0x72, 0xc5, 0x85, 0x4b, 0xe5,
	0xda, 0x53, 0x64, 0xd5, 0xeb, 0x75, 0xbd, 0x76, 0x11, 0x57, 0x4e, 0x3c, 0x09, 0xcf, 0xc7, 0x23,
	0xa0, 0x9d, 0x75, 0x12, 0xbb, 0x75, 0x82, 0xb8, 0x79, 0x66, 0xfe, 0xfb, 0xdf, 0xdf, 0xec, 0x8c,
	0x0c, 0xd3, 0xb2, 0x92, 0xb5, 0x3c, 0x8d, 0x13, 0x71, 0x2a, 0x64, 0xda, 0xe4, 0x18, 0x52, 0x82,
	0x4d, 0xa4, 0x90, 0xa1, 0x50, 0x65, 0x18, 0x27, 0x62, 0xd6, 0x11, 0x25, 0x52, 0x08, 0x59, 0x18,
	0x51, 0xf0, 0xc7, 0x02, 0x38, 0xa3, 0x53, 0x1f, 0x8a, 0x2b, 0xc9, 0x9e, 0x80, 0xd3, 0x64, 0x29,
	0xb7, 0x7c, 0x6b, 0x3e, 0x8e, 0xf4, 0x27, 0x7b, 0x04, 0x76, 0x96, 0x72, 0xdb, 0xb7, 0xe6, 0x6e,
	0x64, 0x67, 0x29, 0xe3, 0x70, 0x98, 0x54, 0x18, 0xd7, 0x98, 0x72, 0xc7, 0xb7, 0xe6, 0x4e, 0xb4,
	0x0e, 0x75, 0xa5, 0x29, 0x53, 0xaa, 0xb8, 0xa6, 0xd2, 0x86, 0x8c, 0x81, 0x5b, 0xc4, 0x02, 0xf9,
	0x88, 0x6c, 0xe9, 0x9b, 0x4d, 0xc1, 0xab, 0x50, 0xc4, 0xd5, 0x35, 0xf7, 0x28, 0xdb, 0x46, 0x1b,
	0x7f, 0x59, 0xf1, 0x43, 0x2a, 0xac, 0x43, 0x36, 0x83, 0x23, 0x59, 0x62, 0x45, 0xa5, 0x23, 0x2a,
	0x6d, 0x62, 0xf6, 0x14, 0x46, 0xf2, 0x7b, 0x81, 0x15, 0x1f, 0x53, 0xc1, 0x04, 0x3a, 0x2b, 0xb0,
	0x68, 0x14, 0x07, 0xdf, 0xd1, 0x59, 0x0a, 0x82, 0x9f, 0x16, 0x3c, 0x88, 0xf0, 0xc6, 0x74, 0xbd,
	0x4c, 0xb7, 0x78, 0xd6, 0x20, 0x9e, 0xdd, 0xc3, 0xeb, 0x42, 0x38, 0xbb, 0x20, 0xdc, 0x41, 0x88,
	0x51, 0x17, 0xa2, 0x84, 0xc7, 0x11, 0x96, 0xf9, 0x8f, 0xce, 0xdb, 0xbf, 0x00, 0x37, 0x2b, 0xae,
	0x24, 0x61, 0x4c, 0x16, 0xc7, 0x61, 0x67, 0x7c, 0xe1, 0x56, 0x16, 0x91, 0x88, 0xbd, 0x04, 0x4f,
	0xd5, 0x71, 0xdd, 0x28, 0xe2, 0x9b, 0x2c, 0x78, 0x4f, 0x4e, 0xd6, 0xe7, 0x54, 0x8f, 0x5a, 0x5d,
	0xf0, 0xcb, 0xea, 0x5d, 0xf9, 0x31, 0x53, 0xb5, 0xee, 0xb2, 0x68, 0xc4, 0x25, 0x56, 0x74, 0xe9,
	0x28, 0x6a, 0xa3, 0xff, 0x77, 0xd7, 0xf0, 0x79, 0xa6, 0x6a, 0xee, 0xf8, 0xce, 0x5e, 0x78, 0x2d,
	0x0a, 0xae, 0x35, 0x49, 0x3b, 0x80, 0x2f, 0xb4, 0x23, 0x03, 0x8b, 0xb7, 0x9e, 0x8a, 0x3d, 0x38,
	0x15, 0x67, 0xe7, 0x54, 0xdc, 0xfe, 0x54, 0x16, 0xbf, 0x1d, 0x78, 0x68, 0xae, 0x3a, 0xc7, 0xea,
	0x36, 0x4b, 0x90, 0xbd, 0x05, 0x6f, 0x99, 0xa6, 0x9f, 0x0a, 0x64, 0xcf, 0xef, 0xf4, 0xb5, 0x5d,
	0x8a, 0xd9, 0xc9, 0xfd, 0x96, 0xb7, 0x7d, 0x04, 0x07, 0x6c, 0x05, 0xde, 0x7b, 0xac, 0xb5, 0xc9,
	0xdd, 0xc7, 0xb9, 0x69, 0x50, 0xd5, 0x5a, 0xf5, 0x4f, 0x8f, 0x37, 0x30, 0x8e, 0x50, 0xc8, 0x5b,
	0xdc, 0x6f, 0x33, 0xbd, 0x6f, 0xd3, 0x83, 0x58, 0xe6, 0xf9, 0xf0, 0xe9, 0xcf, 0xf1, 0x37, 0xdc,
	0x0d, 0xa1, 0x37, 0x20, 0x38, 0x60, 0xef, 0x00, 0xcc, 0x0c, 0x56, 0xb1, 0x42, 0x76, 0x32, 0xfc,
	0x22, 0x46, 0xb1, 0x87, 0x64, 0x09, 0x13, 0xa3, 0x39, 0xd3, 0xeb, 0x3d, 0x8c, 0xa3, 0x2f, 0x1c,
	0xb2, 0x30, 0x20, 0xab, 0xe3, 0xaf, 0xcf, 0x36, 0x3f, 0xa9, 0xd7, 0x52, 0xc8, 0x0b, 0xa1, 0xca,
	0x8b, 0x38, 0x11, 0x97, 0x1e, 0xa5, 0x5f, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x5a, 0xa5, 0x3b,
	0xe7, 0xe9, 0x04, 0x00, 0x00,
}
