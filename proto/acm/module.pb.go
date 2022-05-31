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
	Type                 uint32   `protobuf:"varint,9,opt,name=type,proto3" json:"type,omitempty"`
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

func (m *ModuleInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
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
	Type                 uint32   `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
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

func (m *ReqModuleAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
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
	0x10, 0xc7, 0xeb, 0x8f, 0xb8, 0xcd, 0x84, 0x02, 0x5a, 0x41, 0xba, 0x44, 0x3d, 0x58, 0x3e, 0x59,
	0x42, 0x72, 0x51, 0x38, 0x72, 0x40, 0x09, 0x48, 0x08, 0x89, 0x0a, 0xb4, 0x15, 0x17, 0x2e, 0x95,
	0x6b, 0x4f, 0x91, 0xa9, 0xd7, 0xeb, 0xfa, 0xa3, 0x52, 0xdf, 0x80, 0x27, 0xe1, 0xf1, 0x78, 0x06,
	0xb4, 0x63, 0x27, 0x8e, 0x1b, 0x27, 0xa8, 0xb7, 0x9d, 0x9d, 0xff, 0xfe, 0xe7, 0x37, 0xb3, 0x6b,
	0xc3, 0x34, 0x2f, 0x54, 0xa5, 0xce, 0xc2, 0x48, 0x9e, 0x49, 0x15, 0xd7, 0x29, 0x06, 0xb4, 0xc1,
	0x26, 0x4a, 0xaa, 0x40, 0x96, 0x79, 0x10, 0x46, 0x72, 0xb6, 0x21, 0x8a, 0x94, 0x94, 0x2a, 0x6b,
	0x44, 0xde, 0x5f, 0x03, 0xe0, 0x9c, 0x4e, 0x7d, 0xce, 0xae, 0x15, 0x7b, 0x0e, 0x56, 0x9d, 0xc4,
	0xdc, 0x70, 0x0d, 0x7f, 0x2c, 0xf4, 0x92, 0x3d, 0x05, 0x33, 0x89, 0xb9, 0xe9, 0x1a, 0xbe, 0x2d,
	0xcc, 0x24, 0x66, 0x1c, 0x0e, 0xa3, 0x02, 0xc3, 0x0a, 0x63, 0x6e, 0xb9, 0x86, 0x6f, 0x89, 0x55,
	0xa8, 0x33, 0x75, 0x1e, 0x53, 0xc6, 0x6e, 0x32, 0x6d, 0xc8, 0x18, 0xd8, 0x59, 0x28, 0x91, 0x8f,
	0xc8, 0x96, 0xd6, 0x6c, 0x0a, 0x4e, 0x81, 0x32, 0x2c, 0x6e, 0xb8, 0x43, 0xbb, 0x6d, 0xb4, 0xf6,
	0x57, 0x05, 0x3f, 0xa4, 0xc4, 0x2a, 0x64, 0x33, 0x38, 0x52, 0x39, 0x16, 0x94, 0x3a, 0xa2, 0xd4,
	0x3a, 0xd6, 0x15, 0xaa, 0xfb, 0x1c, 0xf9, 0xd8, 0x35, 0xfc, 0x63, 0x41, 0x6b, 0xf6, 0x02, 0x46,
	0x12, 0xb3, 0xba, 0xe4, 0xe0, 0x5a, 0xfe, 0x58, 0x34, 0x81, 0xf7, 0x0b, 0x9e, 0x08, 0xbc, 0x6d,
	0x5a, 0x5e, 0xc4, 0x1d, 0x9b, 0x31, 0xc8, 0x66, 0xf6, 0xd8, 0x36, 0x09, 0xac, 0x1d, 0x04, 0x76,
	0x47, 0xe0, 0xe5, 0xf0, 0x4c, 0x60, 0x9e, 0xde, 0x6f, 0x0c, 0xf8, 0x35, 0xd8, 0x49, 0x76, 0xad,
	0xa8, 0xdc, 0x64, 0x7e, 0x12, 0x6c, 0xdc, 0x51, 0xd0, 0xc9, 0x04, 0x89, 0xd8, 0x1b, 0x70, 0xca,
	0x2a, 0xac, 0xea, 0x92, 0x38, 0x26, 0x73, 0xde, 0x93, 0x93, 0xf5, 0x05, 0xe5, 0x45, 0xab, 0xf3,
	0x7e, 0x1b, 0xbd, 0x92, 0x5f, 0x92, 0xb2, 0xd2, 0xdd, 0x64, 0xb5, 0xbc, 0xc2, 0x82, 0x8a, 0x8e,
	0x44, 0x1b, 0x3d, 0xde, 0x5d, 0xc3, 0xa7, 0x49, 0x59, 0x71, 0xcb, 0xb5, 0xf6, 0xc2, 0x6b, 0x91,
	0x77, 0xa3, 0x49, 0xda, 0x41, 0x7f, 0xa7, 0x87, 0x30, 0xf0, 0xba, 0x56, 0xd3, 0x37, 0x07, 0xa7,
	0x6f, 0xed, 0x9c, 0xbe, 0xdd, 0x9f, 0xfe, 0xfc, 0x8f, 0x05, 0xc7, 0x4d, 0xa9, 0x0b, 0x2c, 0xee,
	0x92, 0x08, 0xd9, 0x07, 0x70, 0x16, 0x71, 0xfc, 0x35, 0x43, 0xf6, 0xea, 0x41, 0x5f, 0xdd, 0xe5,
	0xcf, 0x4e, 0xb7, 0x5b, 0xee, 0xfa, 0xf0, 0x0e, 0xd8, 0x12, 0x9c, 0x4f, 0x58, 0x69, 0x93, 0x87,
	0xc3, 0xb9, 0xad, 0xb1, 0xac, 0xb4, 0xea, 0xbf, 0x1e, 0xef, 0x61, 0x2c, 0x50, 0xaa, 0x3b, 0xdc,
	0x6f, 0x33, 0xdd, 0xb6, 0xe9, 0x41, 0x2c, 0xd2, 0x74, 0xf8, 0xf4, 0xb7, 0xf0, 0x27, 0xee, 0x86,
	0xd0, 0x2f, 0xc0, 0x3b, 0x60, 0x1f, 0x01, 0x9a, 0x3b, 0x58, 0x86, 0x25, 0xb2, 0xd3, 0xe1, 0x89,
	0x34, 0x8a, 0x3d, 0x24, 0x0b, 0x98, 0x34, 0x9a, 0x73, 0xfd, 0x29, 0x0d, 0xe3, 0xe8, 0x82, 0x43,
	0x16, 0x0d, 0xc8, 0xf2, 0xe4, 0xc7, 0xcb, 0xf5, 0x9f, 0xe8, 0x9d, 0x92, 0xea, 0x52, 0x96, 0xf9,
	0x65, 0x18, 0xc9, 0x2b, 0x87, 0xb6, 0xdf, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xf3, 0xee,
	0xa6, 0xce, 0x04, 0x00, 0x00,
}
