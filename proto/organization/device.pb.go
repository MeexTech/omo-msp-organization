// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/organization/device.proto

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

type DeviceInfo struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Created              uint64   `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated              uint64   `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string   `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string   `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Remark               string   `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	Owner                string   `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
	Sn                   string   `protobuf:"bytes,10,opt,name=sn,proto3" json:"sn,omitempty"`
	Quote                string   `protobuf:"bytes,11,opt,name=quote,proto3" json:"quote,omitempty"`
	Os                   string   `protobuf:"bytes,12,opt,name=os,proto3" json:"os,omitempty"`
	Running              uint32   `protobuf:"varint,13,opt,name=running,proto3" json:"running,omitempty"`
	Certificate          string   `protobuf:"bytes,14,opt,name=certificate,proto3" json:"certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceInfo) Reset()         { *m = DeviceInfo{} }
func (m *DeviceInfo) String() string { return proto.CompactTextString(m) }
func (*DeviceInfo) ProtoMessage()    {}
func (*DeviceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b6e6e2e28953de, []int{0}
}

func (m *DeviceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceInfo.Unmarshal(m, b)
}
func (m *DeviceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceInfo.Marshal(b, m, deterministic)
}
func (m *DeviceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceInfo.Merge(m, src)
}
func (m *DeviceInfo) XXX_Size() int {
	return xxx_messageInfo_DeviceInfo.Size(m)
}
func (m *DeviceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceInfo proto.InternalMessageInfo

func (m *DeviceInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DeviceInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *DeviceInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceInfo) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *DeviceInfo) GetUpdated() uint64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *DeviceInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *DeviceInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *DeviceInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *DeviceInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *DeviceInfo) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *DeviceInfo) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *DeviceInfo) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *DeviceInfo) GetRunning() uint32 {
	if m != nil {
		return m.Running
	}
	return 0
}

func (m *DeviceInfo) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

type ReqDeviceAdd struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Sn                   string   `protobuf:"bytes,4,opt,name=sn,proto3" json:"sn,omitempty"`
	Quote                string   `protobuf:"bytes,5,opt,name=quote,proto3" json:"quote,omitempty"`
	Os                   string   `protobuf:"bytes,6,opt,name=os,proto3" json:"os,omitempty"`
	Operator             string   `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqDeviceAdd) Reset()         { *m = ReqDeviceAdd{} }
func (m *ReqDeviceAdd) String() string { return proto.CompactTextString(m) }
func (*ReqDeviceAdd) ProtoMessage()    {}
func (*ReqDeviceAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b6e6e2e28953de, []int{1}
}

func (m *ReqDeviceAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqDeviceAdd.Unmarshal(m, b)
}
func (m *ReqDeviceAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqDeviceAdd.Marshal(b, m, deterministic)
}
func (m *ReqDeviceAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqDeviceAdd.Merge(m, src)
}
func (m *ReqDeviceAdd) XXX_Size() int {
	return xxx_messageInfo_ReqDeviceAdd.Size(m)
}
func (m *ReqDeviceAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqDeviceAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqDeviceAdd proto.InternalMessageInfo

func (m *ReqDeviceAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqDeviceAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqDeviceAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqDeviceAdd) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *ReqDeviceAdd) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

func (m *ReqDeviceAdd) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *ReqDeviceAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReqDeviceBase struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqDeviceBase) Reset()         { *m = ReqDeviceBase{} }
func (m *ReqDeviceBase) String() string { return proto.CompactTextString(m) }
func (*ReqDeviceBase) ProtoMessage()    {}
func (*ReqDeviceBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b6e6e2e28953de, []int{2}
}

func (m *ReqDeviceBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqDeviceBase.Unmarshal(m, b)
}
func (m *ReqDeviceBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqDeviceBase.Marshal(b, m, deterministic)
}
func (m *ReqDeviceBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqDeviceBase.Merge(m, src)
}
func (m *ReqDeviceBase) XXX_Size() int {
	return xxx_messageInfo_ReqDeviceBase.Size(m)
}
func (m *ReqDeviceBase) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqDeviceBase.DiscardUnknown(m)
}

var xxx_messageInfo_ReqDeviceBase proto.InternalMessageInfo

func (m *ReqDeviceBase) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqDeviceBase) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqDeviceBase) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqDeviceBase) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyDeviceInfo struct {
	Info                 *DeviceInfo  `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyDeviceInfo) Reset()         { *m = ReplyDeviceInfo{} }
func (m *ReplyDeviceInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyDeviceInfo) ProtoMessage()    {}
func (*ReplyDeviceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b6e6e2e28953de, []int{3}
}

func (m *ReplyDeviceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyDeviceInfo.Unmarshal(m, b)
}
func (m *ReplyDeviceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyDeviceInfo.Marshal(b, m, deterministic)
}
func (m *ReplyDeviceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyDeviceInfo.Merge(m, src)
}
func (m *ReplyDeviceInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyDeviceInfo.Size(m)
}
func (m *ReplyDeviceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyDeviceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyDeviceInfo proto.InternalMessageInfo

func (m *ReplyDeviceInfo) GetInfo() *DeviceInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyDeviceInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyDeviceList struct {
	Total                uint32        `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Page                 uint32        `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Pages                uint32        `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Number               uint32        `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
	List                 []*DeviceInfo `protobuf:"bytes,5,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus  `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ReplyDeviceList) Reset()         { *m = ReplyDeviceList{} }
func (m *ReplyDeviceList) String() string { return proto.CompactTextString(m) }
func (*ReplyDeviceList) ProtoMessage()    {}
func (*ReplyDeviceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_68b6e6e2e28953de, []int{4}
}

func (m *ReplyDeviceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyDeviceList.Unmarshal(m, b)
}
func (m *ReplyDeviceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyDeviceList.Marshal(b, m, deterministic)
}
func (m *ReplyDeviceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyDeviceList.Merge(m, src)
}
func (m *ReplyDeviceList) XXX_Size() int {
	return xxx_messageInfo_ReplyDeviceList.Size(m)
}
func (m *ReplyDeviceList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyDeviceList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyDeviceList proto.InternalMessageInfo

func (m *ReplyDeviceList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplyDeviceList) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *ReplyDeviceList) GetPages() uint32 {
	if m != nil {
		return m.Pages
	}
	return 0
}

func (m *ReplyDeviceList) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ReplyDeviceList) GetList() []*DeviceInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyDeviceList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*DeviceInfo)(nil), "omo.msp.scene.DeviceInfo")
	proto.RegisterType((*ReqDeviceAdd)(nil), "omo.msp.scene.ReqDeviceAdd")
	proto.RegisterType((*ReqDeviceBase)(nil), "omo.msp.scene.ReqDeviceBase")
	proto.RegisterType((*ReplyDeviceInfo)(nil), "omo.msp.scene.ReplyDeviceInfo")
	proto.RegisterType((*ReplyDeviceList)(nil), "omo.msp.scene.ReplyDeviceList")
}

func init() {
	proto.RegisterFile("proto/organization/device.proto", fileDescriptor_68b6e6e2e28953de)
}

var fileDescriptor_68b6e6e2e28953de = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x4e, 0x1b, 0x3d,
	0x14, 0x65, 0x92, 0xc9, 0x00, 0x37, 0x0c, 0x7c, 0xb2, 0xd0, 0x27, 0x37, 0xa5, 0x6d, 0x94, 0x15,
	0x9b, 0x06, 0x29, 0x7d, 0x02, 0x68, 0x45, 0x54, 0x89, 0xaa, 0xaa, 0x51, 0x37, 0xdd, 0x0d, 0x33,
	0x17, 0x64, 0x35, 0x63, 0x27, 0xb6, 0x87, 0x8a, 0xbe, 0x4b, 0x37, 0x7d, 0x1e, 0x1e, 0xaa, 0xf2,
	0xf5, 0xe4, 0x07, 0x26, 0x04, 0x56, 0xf1, 0xf1, 0x3d, 0xf7, 0xf8, 0x9c, 0x6b, 0x67, 0xe0, 0xdd,
	0xd4, 0x68, 0xa7, 0x4f, 0xb4, 0xb9, 0xc9, 0x94, 0xfc, 0x9d, 0x39, 0xa9, 0xd5, 0x49, 0x81, 0xb7,
	0x32, 0xc7, 0x21, 0x55, 0x58, 0xaa, 0x4b, 0x3d, 0x2c, 0xed, 0x74, 0x68, 0x73, 0x54, 0xd8, 0x5b,
	0xc7, 0xcf, 0x75, 0x59, 0x6a, 0x15, 0xf8, 0x83, 0xfb, 0x16, 0xc0, 0x27, 0x12, 0xf8, 0xac, 0xae,
	0x35, 0xdb, 0x87, 0x96, 0x2c, 0x78, 0xd4, 0x8f, 0x8e, 0x63, 0xd1, 0x92, 0x05, 0xfb, 0x0f, 0xda,
	0x95, 0x2c, 0x78, 0xab, 0x1f, 0x1d, 0xef, 0x0a, 0xbf, 0x64, 0x0c, 0x62, 0x95, 0x95, 0xc8, 0xdb,
	0xb4, 0x45, 0x6b, 0xc6, 0x61, 0x3b, 0x37, 0x98, 0x39, 0x2c, 0x78, 0x4c, 0xad, 0x73, 0xe8, 0x2b,
	0xd5, 0xb4, 0xa0, 0x4a, 0x27, 0x54, 0x6a, 0xb8, 0xe8, 0xd1, 0x86, 0x27, 0x24, 0x35, 0x87, 0xac,
	0x07, 0x3b, 0x7a, 0x8a, 0x86, 0x4a, 0xdb, 0x54, 0x5a, 0x60, 0xf6, 0x3f, 0x24, 0x06, 0xcb, 0xcc,
	0xfc, 0xe4, 0x3b, 0x54, 0xa9, 0x11, 0x3b, 0x84, 0x8e, 0xfe, 0xa5, 0xd0, 0xf0, 0x5d, 0xda, 0x0e,
	0xc0, 0xa7, 0xb1, 0x8a, 0x03, 0x6d, 0xb5, 0xac, 0xf2, 0xac, 0x59, 0xa5, 0x1d, 0xf2, 0x6e, 0x60,
	0x11, 0xf0, 0x2c, 0x6d, 0xf9, 0x5e, 0x60, 0x69, 0xeb, 0x9d, 0x99, 0x4a, 0x29, 0xa9, 0x6e, 0x78,
	0xda, 0x8f, 0x8e, 0x53, 0x31, 0x87, 0xac, 0x0f, 0xdd, 0x1c, 0x8d, 0x93, 0xd7, 0x32, 0xcf, 0x1c,
	0xf2, 0x7d, 0x6a, 0x59, 0xdd, 0x1a, 0xfc, 0x8d, 0x60, 0x4f, 0xe0, 0x2c, 0x4c, 0xf4, 0xb4, 0x28,
	0x96, 0xc6, 0xa2, 0x55, 0x63, 0xf3, 0x21, 0xb6, 0x56, 0x86, 0xb8, 0x8c, 0xd6, 0x7e, 0x10, 0x2d,
	0x84, 0x88, 0x9b, 0x21, 0x3a, 0xcd, 0x10, 0xc9, 0x22, 0xc4, 0x86, 0x21, 0x0e, 0x24, 0xa4, 0x0b,
	0x8f, 0x67, 0x99, 0xc5, 0xf9, 0x2d, 0x47, 0xcd, 0x5b, 0x7e, 0x89, 0xc1, 0xd5, 0xa3, 0xe2, 0x47,
	0x47, 0x39, 0x38, 0x10, 0x38, 0x9d, 0xdc, 0xad, 0x3c, 0xb1, 0xf7, 0x10, 0x4b, 0x75, 0xad, 0xe9,
	0xb4, 0xee, 0xe8, 0xd5, 0xf0, 0xc1, 0x83, 0x1d, 0x2e, 0x89, 0x82, 0x68, 0x6c, 0x04, 0x89, 0x75,
	0x99, 0xab, 0x2c, 0x79, 0xe9, 0x8e, 0x7a, 0x8f, 0x1a, 0x48, 0xfe, 0x92, 0x18, 0xa2, 0x66, 0x0e,
	0xee, 0xa3, 0x07, 0xc7, 0x5e, 0x48, 0xeb, 0xfc, 0xd8, 0x9c, 0x76, 0xd9, 0x84, 0xce, 0x4d, 0x45,
	0x00, 0x3e, 0xe7, 0x34, 0xbb, 0x09, 0x39, 0x53, 0x41, 0x6b, 0xcf, 0xf4, 0xbf, 0x96, 0x62, 0xa6,
	0x22, 0x00, 0x9f, 0x5e, 0x55, 0xe5, 0x15, 0x86, 0x8c, 0xa9, 0xa8, 0x91, 0x8f, 0x33, 0x91, 0xd6,
	0xf1, 0x4e, 0xbf, 0xfd, 0x4c, 0x1c, 0x4f, 0x5b, 0x89, 0x93, 0xbc, 0x34, 0xce, 0xe8, 0x4f, 0x0c,
	0x69, 0x10, 0xba, 0x44, 0xe3, 0x7f, 0xd8, 0x18, 0x92, 0xd3, 0xa2, 0xf8, 0xaa, 0x90, 0xbd, 0x6e,
	0xf4, 0x2f, 0x1f, 0x5f, 0xef, 0xed, 0x3a, 0xf1, 0xa5, 0xa5, 0xc1, 0x16, 0x3b, 0x87, 0x64, 0x8c,
	0xce, 0x0b, 0x35, 0x8d, 0xcc, 0x2a, 0xb4, 0xce, 0xf3, 0x5e, 0xa0, 0xf3, 0x11, 0x76, 0x05, 0x96,
	0xfa, 0x16, 0x9f, 0x93, 0xe2, 0xeb, 0xa4, 0x6a, 0x91, 0x6f, 0x70, 0x30, 0x46, 0xe7, 0x6f, 0xeb,
	0xec, 0xee, 0x5c, 0x4e, 0x1c, 0x1a, 0x76, 0xb4, 0x5e, 0x2a, 0x54, 0x37, 0xf9, 0xf2, 0x2a, 0x83,
	0x2d, 0xf6, 0x05, 0xf6, 0xc6, 0xe8, 0xfc, 0x3c, 0xa5, 0x75, 0x32, 0x7f, 0x46, 0xef, 0xcd, 0x53,
	0x97, 0x41, 0xcd, 0x34, 0x2e, 0xf8, 0x4e, 0xdf, 0x2f, 0xfa, 0xdb, 0x1c, 0x3d, 0x35, 0x7b, 0x5f,
	0xdd, 0x98, 0xf4, 0x02, 0xf6, 0x6b, 0x9d, 0x79, 0xd0, 0x66, 0x94, 0x59, 0x60, 0xd4, 0xd6, 0x36,
	0xa8, 0x9d, 0x1d, 0xfe, 0x60, 0xcd, 0xcf, 0xfc, 0x55, 0x42, 0x7b, 0x1f, 0xfe, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x7c, 0x75, 0x3b, 0x2c, 0x33, 0x06, 0x00, 0x00,
}
