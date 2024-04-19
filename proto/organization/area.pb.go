// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/organization/area.proto

package organization

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

type ReqAreaAdd struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Parent               string   `protobuf:"bytes,4,opt,name=parent,proto3" json:"parent,omitempty"`
	Operator             string   `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Aspect               string   `protobuf:"bytes,6,opt,name=aspect,proto3" json:"aspect,omitempty"`
	Assets               []string `protobuf:"bytes,15,rep,name=assets,proto3" json:"assets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqAreaAdd) Reset()         { *m = ReqAreaAdd{} }
func (m *ReqAreaAdd) String() string { return proto.CompactTextString(m) }
func (*ReqAreaAdd) ProtoMessage()    {}
func (*ReqAreaAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77d3e3dd5be64be, []int{0}
}

func (m *ReqAreaAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqAreaAdd.Unmarshal(m, b)
}
func (m *ReqAreaAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqAreaAdd.Marshal(b, m, deterministic)
}
func (m *ReqAreaAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqAreaAdd.Merge(m, src)
}
func (m *ReqAreaAdd) XXX_Size() int {
	return xxx_messageInfo_ReqAreaAdd.Size(m)
}
func (m *ReqAreaAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqAreaAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqAreaAdd proto.InternalMessageInfo

func (m *ReqAreaAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqAreaAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqAreaAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqAreaAdd) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ReqAreaAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqAreaAdd) GetAspect() string {
	if m != nil {
		return m.Aspect
	}
	return ""
}

func (m *ReqAreaAdd) GetAssets() []string {
	if m != nil {
		return m.Assets
	}
	return nil
}

type ReqAreaBase struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqAreaBase) Reset()         { *m = ReqAreaBase{} }
func (m *ReqAreaBase) String() string { return proto.CompactTextString(m) }
func (*ReqAreaBase) ProtoMessage()    {}
func (*ReqAreaBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77d3e3dd5be64be, []int{1}
}

func (m *ReqAreaBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqAreaBase.Unmarshal(m, b)
}
func (m *ReqAreaBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqAreaBase.Marshal(b, m, deterministic)
}
func (m *ReqAreaBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqAreaBase.Merge(m, src)
}
func (m *ReqAreaBase) XXX_Size() int {
	return xxx_messageInfo_ReqAreaBase.Size(m)
}
func (m *ReqAreaBase) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqAreaBase.DiscardUnknown(m)
}

var xxx_messageInfo_ReqAreaBase proto.InternalMessageInfo

func (m *ReqAreaBase) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqAreaBase) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqAreaBase) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqAreaBase) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyAreaInfo struct {
	Info                 *AreaInfo    `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyAreaInfo) Reset()         { *m = ReplyAreaInfo{} }
func (m *ReplyAreaInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyAreaInfo) ProtoMessage()    {}
func (*ReplyAreaInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77d3e3dd5be64be, []int{2}
}

func (m *ReplyAreaInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyAreaInfo.Unmarshal(m, b)
}
func (m *ReplyAreaInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyAreaInfo.Marshal(b, m, deterministic)
}
func (m *ReplyAreaInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyAreaInfo.Merge(m, src)
}
func (m *ReplyAreaInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyAreaInfo.Size(m)
}
func (m *ReplyAreaInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyAreaInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyAreaInfo proto.InternalMessageInfo

func (m *ReplyAreaInfo) GetInfo() *AreaInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyAreaInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyAreaList struct {
	Total                uint32       `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Page                 uint32       `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Pages                uint32       `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Number               uint32       `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
	List                 []*AreaInfo  `protobuf:"bytes,5,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyAreaList) Reset()         { *m = ReplyAreaList{} }
func (m *ReplyAreaList) String() string { return proto.CompactTextString(m) }
func (*ReplyAreaList) ProtoMessage()    {}
func (*ReplyAreaList) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77d3e3dd5be64be, []int{3}
}

func (m *ReplyAreaList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyAreaList.Unmarshal(m, b)
}
func (m *ReplyAreaList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyAreaList.Marshal(b, m, deterministic)
}
func (m *ReplyAreaList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyAreaList.Merge(m, src)
}
func (m *ReplyAreaList) XXX_Size() int {
	return xxx_messageInfo_ReplyAreaList.Size(m)
}
func (m *ReplyAreaList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyAreaList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyAreaList proto.InternalMessageInfo

func (m *ReplyAreaList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyAreaList) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyAreaList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyAreaList) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyAreaList) GetList() []*AreaInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyAreaList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*ReqAreaAdd)(nil), "omo.msp.scene.ReqAreaAdd")
	proto.RegisterType((*ReqAreaBase)(nil), "omo.msp.scene.ReqAreaBase")
	proto.RegisterType((*ReplyAreaInfo)(nil), "omo.msp.scene.ReplyAreaInfo")
	proto.RegisterType((*ReplyAreaList)(nil), "omo.msp.scene.ReplyAreaList")
}

func init() {
	proto.RegisterFile("proto/organization/area.proto", fileDescriptor_c77d3e3dd5be64be)
}

var fileDescriptor_c77d3e3dd5be64be = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x5d, 0x69, 0x16, 0xb1, 0x5b, 0xc2, 0x90, 0x35, 0x41, 0x88, 0x36, 0xa8, 0xf2, 0x84, 0x84,
	0x94, 0x49, 0xe5, 0x0b, 0xda, 0x21, 0x2a, 0xa4, 0xa1, 0x49, 0xae, 0x78, 0xe1, 0xcd, 0x4b, 0xee,
	0x8a, 0x45, 0x63, 0x67, 0xb6, 0x3b, 0x34, 0x3e, 0x8b, 0xaf, 0xe0, 0x0f, 0xf8, 0x1d, 0x74, 0xed,
	0x14, 0x3a, 0x1a, 0x5a, 0xe0, 0xa9, 0x3e, 0xc7, 0xc7, 0xc7, 0xe7, 0xf8, 0xb6, 0x85, 0x93, 0xc6,
	0x68, 0xa7, 0x4f, 0xb5, 0x99, 0x0b, 0x25, 0xbf, 0x08, 0x27, 0xb5, 0x3a, 0x15, 0x06, 0x45, 0xe1,
	0x79, 0x96, 0xe8, 0x5a, 0x17, 0xb5, 0x6d, 0x0a, 0x5b, 0xa2, 0xc2, 0xec, 0x79, 0x87, 0xba, 0xd4,
	0x75, 0xad, 0x55, 0xd0, 0xe7, 0x5f, 0x7b, 0x00, 0x1c, 0xaf, 0xc7, 0x06, 0xc5, 0xb8, 0xaa, 0xd8,
	0x11, 0xec, 0xeb, 0xcf, 0x0a, 0x4d, 0xda, 0x1b, 0xf6, 0x5e, 0x1c, 0xf0, 0x00, 0x18, 0x83, 0x48,
	0x89, 0x1a, 0xd3, 0x7b, 0x9e, 0xf4, 0x6b, 0xf6, 0x18, 0x62, 0x83, 0xb5, 0x30, 0x9f, 0xd2, 0xbe,
	0x67, 0x5b, 0x44, 0x7c, 0x23, 0x0c, 0x2a, 0x97, 0x46, 0x81, 0x0f, 0x88, 0x65, 0x70, 0x5f, 0x37,
	0x68, 0x84, 0xd3, 0x26, 0xdd, 0xf7, 0x3b, 0x3f, 0x31, 0x9d, 0x11, 0xb6, 0xc1, 0xd2, 0xa5, 0x71,
	0x38, 0x13, 0x50, 0xe0, 0x2d, 0x3a, 0x9b, 0x1e, 0x0e, 0xfb, 0x81, 0x27, 0x94, 0xcf, 0x61, 0xd0,
	0x66, 0x9e, 0x08, 0x8b, 0xec, 0x11, 0xf4, 0x97, 0xb2, 0x6a, 0x23, 0xd3, 0xf2, 0x9f, 0x02, 0xaf,
	0x07, 0x8b, 0xee, 0x06, 0xcb, 0x1b, 0x48, 0x38, 0x36, 0x8b, 0x5b, 0xba, 0xea, 0xad, 0xba, 0xd2,
	0xec, 0x25, 0x44, 0x52, 0x5d, 0x69, 0x7f, 0xd7, 0x60, 0xf4, 0xa4, 0xb8, 0xf3, 0xda, 0xc5, 0x4a,
	0xc6, 0xbd, 0x88, 0x8d, 0x20, 0xb6, 0x4e, 0xb8, 0xa5, 0xf5, 0x39, 0x06, 0xa3, 0xec, 0x37, 0xb9,
	0xb7, 0x9e, 0x79, 0x05, 0x6f, 0x95, 0xf9, 0xb7, 0xde, 0xda, 0x95, 0xe7, 0xd2, 0x3a, 0x1a, 0x89,
	0xd3, 0x4e, 0x2c, 0xfc, 0x9d, 0x09, 0x0f, 0x80, 0x1a, 0x36, 0x62, 0x1e, 0x1a, 0x26, 0xdc, 0xaf,
	0x49, 0x49, 0x9f, 0xd6, 0x17, 0x4c, 0x78, 0x00, 0xd4, 0x5b, 0x2d, 0xeb, 0x4b, 0x0c, 0xed, 0x12,
	0xde, 0x22, 0xaa, 0xb2, 0x90, 0xd6, 0xa5, 0xfb, 0xc3, 0xfe, 0xd6, 0x2a, 0x24, 0x5a, 0xab, 0x12,
	0xff, 0x6d, 0x95, 0xd1, 0xf7, 0x08, 0x06, 0x64, 0x33, 0x43, 0x73, 0x23, 0x4b, 0x64, 0x67, 0x10,
	0x8f, 0xab, 0xea, 0x42, 0x21, 0x7b, 0xba, 0x71, 0x7a, 0xf5, 0x05, 0xcc, 0x8e, 0xbb, 0x8c, 0x57,
	0x61, 0xf2, 0x3d, 0xf6, 0x1a, 0xe2, 0x29, 0x3a, 0x32, 0xd9, 0x8c, 0x70, 0xbd, 0x44, 0xeb, 0x48,
	0xb7, 0xd3, 0xe5, 0x0c, 0x0e, 0x38, 0xd6, 0xfa, 0x06, 0x77, 0x19, 0xa5, 0x5d, 0x46, 0xbf, 0xa2,
	0xcc, 0x50, 0x98, 0xf2, 0xe3, 0xff, 0x45, 0xa1, 0xe1, 0xe6, 0x7b, 0xec, 0x02, 0x0e, 0xa7, 0xe8,
	0x08, 0x4c, 0x6e, 0xdf, 0xc8, 0x85, 0x43, 0xc3, 0x8e, 0xbb, 0xed, 0xc2, 0xee, 0x4e, 0xc3, 0x77,
	0xf0, 0x60, 0x8a, 0x8e, 0x66, 0x21, 0xad, 0x93, 0xe5, 0x0e, 0xb7, 0x93, 0x3f, 0x0d, 0xd2, 0x1f,
	0xf6, 0x2d, 0xe1, 0x7d, 0x53, 0x09, 0x87, 0xfe, 0xa7, 0x96, 0x75, 0x4f, 0x8e, 0xf6, 0xb6, 0xbe,
	0xd5, 0x39, 0x3c, 0x6c, 0x5d, 0x56, 0x25, 0x9f, 0x6d, 0x3a, 0x05, 0x45, 0x1b, 0x6c, 0x8b, 0xdb,
	0xe4, 0xe8, 0x03, 0xdb, 0xfc, 0x5f, 0xbb, 0x8c, 0x3d, 0xf7, 0xea, 0x47, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x18, 0x2d, 0x9a, 0x8e, 0x22, 0x05, 0x00, 0x00,
}
