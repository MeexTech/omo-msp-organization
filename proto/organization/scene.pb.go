// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/organization/scene.proto

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

type SceneInfo struct {
	Uid                  string         `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Id                   uint64         `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Created              int64          `protobuf:"varint,3,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64          `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	Name                 string         `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Type                 int32          `protobuf:"varint,6,opt,name=type,proto3" json:"type,omitempty"`
	Status               int32          `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	Cover                string         `protobuf:"bytes,8,opt,name=cover,proto3" json:"cover,omitempty"`
	Master               string         `protobuf:"bytes,9,opt,name=master,proto3" json:"master,omitempty"`
	Remark               string         `protobuf:"bytes,10,opt,name=remark,proto3" json:"remark,omitempty"`
	Location             string         `protobuf:"bytes,11,opt,name=location,proto3" json:"location,omitempty"`
	Entity               string         `protobuf:"bytes,12,opt,name=entity,proto3" json:"entity,omitempty"`
	Creator              string         `protobuf:"bytes,13,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string         `protobuf:"bytes,14,opt,name=operator,proto3" json:"operator,omitempty"`
	Short                string         `protobuf:"bytes,15,opt,name=short,proto3" json:"short,omitempty"`
	Address              *AddressInfo   `protobuf:"bytes,16,opt,name=address,proto3" json:"address,omitempty"`
	Supporter            string         `protobuf:"bytes,17,opt,name=supporter,proto3" json:"supporter,omitempty"`
	Bucket               string         `protobuf:"bytes,19,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Parents              []string       `protobuf:"bytes,20,rep,name=parents,proto3" json:"parents,omitempty"`
	Members              []string       `protobuf:"bytes,21,rep,name=members,proto3" json:"members,omitempty"`
	Domains              []*ProductInfo `protobuf:"bytes,22,rep,name=domains,proto3" json:"domains,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SceneInfo) Reset()         { *m = SceneInfo{} }
func (m *SceneInfo) String() string { return proto.CompactTextString(m) }
func (*SceneInfo) ProtoMessage()    {}
func (*SceneInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_059c9a53c993103b, []int{0}
}

func (m *SceneInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SceneInfo.Unmarshal(m, b)
}
func (m *SceneInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SceneInfo.Marshal(b, m, deterministic)
}
func (m *SceneInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SceneInfo.Merge(m, src)
}
func (m *SceneInfo) XXX_Size() int {
	return xxx_messageInfo_SceneInfo.Size(m)
}
func (m *SceneInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SceneInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SceneInfo proto.InternalMessageInfo

func (m *SceneInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *SceneInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SceneInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *SceneInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *SceneInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SceneInfo) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *SceneInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *SceneInfo) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *SceneInfo) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *SceneInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *SceneInfo) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *SceneInfo) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *SceneInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *SceneInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *SceneInfo) GetShort() string {
	if m != nil {
		return m.Short
	}
	return ""
}

func (m *SceneInfo) GetAddress() *AddressInfo {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *SceneInfo) GetSupporter() string {
	if m != nil {
		return m.Supporter
	}
	return ""
}

func (m *SceneInfo) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *SceneInfo) GetParents() []string {
	if m != nil {
		return m.Parents
	}
	return nil
}

func (m *SceneInfo) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *SceneInfo) GetDomains() []*ProductInfo {
	if m != nil {
		return m.Domains
	}
	return nil
}

type ReqSceneAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 int32    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Cover                string   `protobuf:"bytes,3,opt,name=cover,proto3" json:"cover,omitempty"`
	Master               string   `protobuf:"bytes,4,opt,name=master,proto3" json:"master,omitempty"`
	Remark               string   `protobuf:"bytes,5,opt,name=remark,proto3" json:"remark,omitempty"`
	Location             string   `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	Entity               string   `protobuf:"bytes,7,opt,name=entity,proto3" json:"entity,omitempty"`
	Operator             string   `protobuf:"bytes,8,opt,name=operator,proto3" json:"operator,omitempty"`
	Short                string   `protobuf:"bytes,9,opt,name=short,proto3" json:"short,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSceneAdd) Reset()         { *m = ReqSceneAdd{} }
func (m *ReqSceneAdd) String() string { return proto.CompactTextString(m) }
func (*ReqSceneAdd) ProtoMessage()    {}
func (*ReqSceneAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_059c9a53c993103b, []int{1}
}

func (m *ReqSceneAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSceneAdd.Unmarshal(m, b)
}
func (m *ReqSceneAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSceneAdd.Marshal(b, m, deterministic)
}
func (m *ReqSceneAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSceneAdd.Merge(m, src)
}
func (m *ReqSceneAdd) XXX_Size() int {
	return xxx_messageInfo_ReqSceneAdd.Size(m)
}
func (m *ReqSceneAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSceneAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSceneAdd proto.InternalMessageInfo

func (m *ReqSceneAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqSceneAdd) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ReqSceneAdd) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqSceneAdd) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *ReqSceneAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqSceneAdd) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *ReqSceneAdd) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *ReqSceneAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqSceneAdd) GetShort() string {
	if m != nil {
		return m.Short
	}
	return ""
}

type ReplySceneOne struct {
	Info                 *SceneInfo   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplySceneOne) Reset()         { *m = ReplySceneOne{} }
func (m *ReplySceneOne) String() string { return proto.CompactTextString(m) }
func (*ReplySceneOne) ProtoMessage()    {}
func (*ReplySceneOne) Descriptor() ([]byte, []int) {
	return fileDescriptor_059c9a53c993103b, []int{2}
}

func (m *ReplySceneOne) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplySceneOne.Unmarshal(m, b)
}
func (m *ReplySceneOne) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplySceneOne.Marshal(b, m, deterministic)
}
func (m *ReplySceneOne) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplySceneOne.Merge(m, src)
}
func (m *ReplySceneOne) XXX_Size() int {
	return xxx_messageInfo_ReplySceneOne.Size(m)
}
func (m *ReplySceneOne) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplySceneOne.DiscardUnknown(m)
}

var xxx_messageInfo_ReplySceneOne proto.InternalMessageInfo

func (m *ReplySceneOne) GetInfo() *SceneInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplySceneOne) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplySceneList struct {
	Total                uint32       `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	PageMax              uint32       `protobuf:"varint,2,opt,name=pageMax,proto3" json:"pageMax,omitempty"`
	PageNow              uint32       `protobuf:"varint,3,opt,name=pageNow,proto3" json:"pageNow,omitempty"`
	List                 []*SceneInfo `protobuf:"bytes,4,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplySceneList) Reset()         { *m = ReplySceneList{} }
func (m *ReplySceneList) String() string { return proto.CompactTextString(m) }
func (*ReplySceneList) ProtoMessage()    {}
func (*ReplySceneList) Descriptor() ([]byte, []int) {
	return fileDescriptor_059c9a53c993103b, []int{3}
}

func (m *ReplySceneList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplySceneList.Unmarshal(m, b)
}
func (m *ReplySceneList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplySceneList.Marshal(b, m, deterministic)
}
func (m *ReplySceneList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplySceneList.Merge(m, src)
}
func (m *ReplySceneList) XXX_Size() int {
	return xxx_messageInfo_ReplySceneList.Size(m)
}
func (m *ReplySceneList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplySceneList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplySceneList proto.InternalMessageInfo

func (m *ReplySceneList) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ReplySceneList) GetPageMax() uint32 {
	if m != nil {
		return m.PageMax
	}
	return 0
}

func (m *ReplySceneList) GetPageNow() uint32 {
	if m != nil {
		return m.PageNow
	}
	return 0
}

func (m *ReplySceneList) GetList() []*SceneInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplySceneList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReqSceneUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cover                string   `protobuf:"bytes,3,opt,name=cover,proto3" json:"cover,omitempty"`
	Remark               string   `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark,omitempty"`
	Master               string   `protobuf:"bytes,6,opt,name=master,proto3" json:"master,omitempty"`
	Operator             string   `protobuf:"bytes,7,opt,name=operator,proto3" json:"operator,omitempty"`
	Type                 uint32   `protobuf:"varint,8,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSceneUpdate) Reset()         { *m = ReqSceneUpdate{} }
func (m *ReqSceneUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqSceneUpdate) ProtoMessage()    {}
func (*ReqSceneUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_059c9a53c993103b, []int{4}
}

func (m *ReqSceneUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSceneUpdate.Unmarshal(m, b)
}
func (m *ReqSceneUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSceneUpdate.Marshal(b, m, deterministic)
}
func (m *ReqSceneUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSceneUpdate.Merge(m, src)
}
func (m *ReqSceneUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqSceneUpdate.Size(m)
}
func (m *ReqSceneUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSceneUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSceneUpdate proto.InternalMessageInfo

func (m *ReqSceneUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqSceneUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqSceneUpdate) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *ReqSceneUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqSceneUpdate) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *ReqSceneUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqSceneUpdate) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type ReplyMasterUsed struct {
	Master               string       `protobuf:"bytes,1,opt,name=master,proto3" json:"master,omitempty"`
	Used                 bool         `protobuf:"varint,2,opt,name=used,proto3" json:"used,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyMasterUsed) Reset()         { *m = ReplyMasterUsed{} }
func (m *ReplyMasterUsed) String() string { return proto.CompactTextString(m) }
func (*ReplyMasterUsed) ProtoMessage()    {}
func (*ReplyMasterUsed) Descriptor() ([]byte, []int) {
	return fileDescriptor_059c9a53c993103b, []int{5}
}

func (m *ReplyMasterUsed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyMasterUsed.Unmarshal(m, b)
}
func (m *ReplyMasterUsed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyMasterUsed.Marshal(b, m, deterministic)
}
func (m *ReplyMasterUsed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyMasterUsed.Merge(m, src)
}
func (m *ReplyMasterUsed) XXX_Size() int {
	return xxx_messageInfo_ReplyMasterUsed.Size(m)
}
func (m *ReplyMasterUsed) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyMasterUsed.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyMasterUsed proto.InternalMessageInfo

func (m *ReplyMasterUsed) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *ReplyMasterUsed) GetUsed() bool {
	if m != nil {
		return m.Used
	}
	return false
}

func (m *ReplyMasterUsed) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReqSceneStatus struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSceneStatus) Reset()         { *m = ReqSceneStatus{} }
func (m *ReqSceneStatus) String() string { return proto.CompactTextString(m) }
func (*ReqSceneStatus) ProtoMessage()    {}
func (*ReqSceneStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_059c9a53c993103b, []int{6}
}

func (m *ReqSceneStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSceneStatus.Unmarshal(m, b)
}
func (m *ReqSceneStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSceneStatus.Marshal(b, m, deterministic)
}
func (m *ReqSceneStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSceneStatus.Merge(m, src)
}
func (m *ReqSceneStatus) XXX_Size() int {
	return xxx_messageInfo_ReqSceneStatus.Size(m)
}
func (m *ReqSceneStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSceneStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSceneStatus proto.InternalMessageInfo

func (m *ReqSceneStatus) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqSceneStatus) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReqSceneStatus) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReqSceneDomains struct {
	Uid                  string         `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Operator             string         `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	List                 []*ProductInfo `protobuf:"bytes,3,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ReqSceneDomains) Reset()         { *m = ReqSceneDomains{} }
func (m *ReqSceneDomains) String() string { return proto.CompactTextString(m) }
func (*ReqSceneDomains) ProtoMessage()    {}
func (*ReqSceneDomains) Descriptor() ([]byte, []int) {
	return fileDescriptor_059c9a53c993103b, []int{7}
}

func (m *ReqSceneDomains) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSceneDomains.Unmarshal(m, b)
}
func (m *ReqSceneDomains) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSceneDomains.Marshal(b, m, deterministic)
}
func (m *ReqSceneDomains) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSceneDomains.Merge(m, src)
}
func (m *ReqSceneDomains) XXX_Size() int {
	return xxx_messageInfo_ReqSceneDomains.Size(m)
}
func (m *ReqSceneDomains) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSceneDomains.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSceneDomains proto.InternalMessageInfo

func (m *ReqSceneDomains) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqSceneDomains) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqSceneDomains) GetList() []*ProductInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*SceneInfo)(nil), "omo.msp.scene.SceneInfo")
	proto.RegisterType((*ReqSceneAdd)(nil), "omo.msp.scene.ReqSceneAdd")
	proto.RegisterType((*ReplySceneOne)(nil), "omo.msp.scene.ReplySceneOne")
	proto.RegisterType((*ReplySceneList)(nil), "omo.msp.scene.ReplySceneList")
	proto.RegisterType((*ReqSceneUpdate)(nil), "omo.msp.scene.ReqSceneUpdate")
	proto.RegisterType((*ReplyMasterUsed)(nil), "omo.msp.scene.ReplyMasterUsed")
	proto.RegisterType((*ReqSceneStatus)(nil), "omo.msp.scene.ReqSceneStatus")
	proto.RegisterType((*ReqSceneDomains)(nil), "omo.msp.scene.ReqSceneDomains")
}

func init() {
	proto.RegisterFile("proto/organization/scene.proto", fileDescriptor_059c9a53c993103b)
}

var fileDescriptor_059c9a53c993103b = []byte{
	// 949 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xae, 0x2c, 0xf9, 0x8f, 0xfe, 0x49, 0xc7, 0x65, 0x01, 0x61, 0xa4, 0x99, 0xa1, 0x2b, 0x5f,
	0x0c, 0x2e, 0x90, 0xed, 0x05, 0x92, 0x15, 0x71, 0x83, 0x25, 0x59, 0xa0, 0xa2, 0xbb, 0xd8, 0x9d,
	0x22, 0x9d, 0x66, 0x42, 0x2d, 0x51, 0x21, 0xa9, 0x74, 0xd9, 0x73, 0xec, 0x31, 0xf6, 0x08, 0xdb,
	0xc3, 0xec, 0x4d, 0x06, 0x1e, 0x52, 0xb6, 0x9c, 0x4a, 0x76, 0x97, 0xdc, 0xe9, 0xf0, 0x1c, 0x7e,
	0xfa, 0xce, 0xf9, 0x3e, 0x51, 0x24, 0x47, 0xb9, 0xe0, 0x8a, 0xbf, 0xe6, 0xe2, 0x36, 0xcc, 0x92,
	0x3f, 0x42, 0x95, 0xf0, 0xec, 0xb5, 0x8c, 0x20, 0x83, 0x39, 0x26, 0xe8, 0x88, 0xa7, 0x7c, 0x9e,
	0xca, 0x7c, 0x8e, 0x8b, 0x93, 0x6f, 0x6b, 0xca, 0x23, 0x9e, 0xa6, 0x3c, 0x33, 0xf5, 0xfe, 0x3f,
	0x1e, 0xe9, 0xbf, 0xd3, 0xa5, 0xe7, 0xd9, 0x07, 0x4e, 0x5f, 0x12, 0xb7, 0x48, 0x62, 0xe6, 0x4c,
	0x9d, 0x59, 0x3f, 0xd0, 0x8f, 0x74, 0x4c, 0x5a, 0x49, 0xcc, 0x5a, 0x53, 0x67, 0xe6, 0x05, 0xad,
	0x24, 0xa6, 0x8c, 0x74, 0x23, 0x01, 0xa1, 0x82, 0x98, 0xb9, 0x53, 0x67, 0xe6, 0x06, 0x65, 0xa8,
	0x33, 0x45, 0x1e, 0x63, 0xc6, 0x33, 0x19, 0x1b, 0x52, 0x4a, 0xbc, 0x2c, 0x4c, 0x81, 0xb5, 0x11,
	0x16, 0x9f, 0xf5, 0x9a, 0x7a, 0xc8, 0x81, 0x75, 0xa6, 0xce, 0xac, 0x1d, 0xe0, 0x33, 0x3d, 0x20,
	0x1d, 0xa9, 0x42, 0x55, 0x48, 0xd6, 0xc5, 0x55, 0x1b, 0xd1, 0x7d, 0xd2, 0x8e, 0xf8, 0x3d, 0x08,
	0xd6, 0x43, 0x00, 0x13, 0xe8, 0xea, 0x34, 0x94, 0x0a, 0x04, 0xeb, 0xe3, 0xb2, 0x8d, 0xf4, 0xba,
	0x80, 0x34, 0x14, 0x1f, 0x19, 0x31, 0xeb, 0x26, 0xa2, 0x13, 0xd2, 0x5b, 0xf2, 0x08, 0x47, 0xc0,
	0x06, 0x98, 0x59, 0xc5, 0x7a, 0x0f, 0x64, 0x2a, 0x51, 0x0f, 0x6c, 0x68, 0xf6, 0x98, 0x68, 0xd5,
	0x2d, 0x17, 0x6c, 0x84, 0x89, 0x32, 0xd4, 0x68, 0x3c, 0x07, 0x81, 0xa9, 0xb1, 0x41, 0x2b, 0x63,
	0xcd, 0x57, 0xfe, 0xc6, 0x85, 0x62, 0x7b, 0x86, 0x2f, 0x06, 0xf4, 0x07, 0xd2, 0x0d, 0xe3, 0x58,
	0x80, 0x94, 0xec, 0xe5, 0xd4, 0x99, 0x0d, 0x8e, 0x27, 0xf3, 0x0d, 0xad, 0xe6, 0x27, 0x26, 0xab,
	0x85, 0x08, 0xca, 0x52, 0x7a, 0x48, 0xfa, 0xb2, 0xc8, 0x73, 0x2e, 0x74, 0xa3, 0x5f, 0x21, 0xde,
	0x7a, 0x41, 0xf3, 0xbe, 0x29, 0xa2, 0x8f, 0xa0, 0xd8, 0xd7, 0x86, 0xb7, 0x89, 0x34, 0xef, 0x3c,
	0x14, 0x90, 0x29, 0xc9, 0xf6, 0xa7, 0xae, 0xe6, 0x6d, 0x43, 0x9d, 0x49, 0x21, 0xbd, 0x01, 0x21,
	0xd9, 0x37, 0x26, 0x63, 0x43, 0xcd, 0x2f, 0xe6, 0x69, 0x98, 0x64, 0x92, 0x1d, 0x4c, 0xdd, 0x1a,
	0x7e, 0xd7, 0x82, 0xc7, 0x45, 0xa4, 0x0c, 0x3f, 0x5b, 0xea, 0xff, 0xeb, 0x90, 0x41, 0x00, 0x77,
	0x68, 0xa1, 0x93, 0x78, 0xad, 0xb5, 0x53, 0xa3, 0x75, 0xab, 0xa2, 0xf5, 0x4a, 0x53, 0xb7, 0x5e,
	0x53, 0xaf, 0x41, 0xd3, 0x76, 0xa3, 0xa6, 0x9d, 0x46, 0x4d, 0xbb, 0x1b, 0x9a, 0x56, 0x95, 0xeb,
	0x35, 0x29, 0xd7, 0xaf, 0x28, 0xe7, 0xdf, 0x91, 0x51, 0x00, 0xf9, 0xf2, 0x01, 0x9b, 0xfc, 0x39,
	0x03, 0xfa, 0x1d, 0xf1, 0x92, 0xec, 0x03, 0xc7, 0x26, 0x07, 0xc7, 0xec, 0xd1, 0x9c, 0x56, 0x9f,
	0x53, 0x80, 0x55, 0xf4, 0x78, 0x65, 0xeb, 0x56, 0xad, 0xee, 0x06, 0x1b, 0x2b, 0x4a, 0xcb, 0xfb,
	0x7f, 0x3b, 0x64, 0xbc, 0x7e, 0xe7, 0x45, 0x22, 0x95, 0xe6, 0xa6, 0xb8, 0x0a, 0x97, 0xf8, 0xd6,
	0x51, 0x60, 0x02, 0xa3, 0xf4, 0x2d, 0x5c, 0x86, 0xbf, 0x23, 0xfa, 0x28, 0x28, 0xc3, 0x32, 0x73,
	0xc5, 0x3f, 0xe1, 0x8c, 0x6d, 0xe6, 0x8a, 0x7f, 0xd2, 0xf4, 0x97, 0x89, 0x54, 0xcc, 0x43, 0x99,
	0xb7, 0xd0, 0xd7, 0x55, 0x15, 0xfa, 0xed, 0x2f, 0xa6, 0xff, 0x17, 0xd2, 0x37, 0xae, 0x78, 0x8f,
	0xa7, 0x40, 0xcd, 0xd1, 0x52, 0x5a, 0xa5, 0x55, 0xb1, 0x4a, 0xa3, 0x2d, 0xac, 0xfc, 0xde, 0x86,
	0xfc, 0x6b, 0xbb, 0x74, 0x36, 0xec, 0x52, 0x95, 0xb8, 0xfb, 0x48, 0xe2, 0xd2, 0x8c, 0x3d, 0x9c,
	0x09, 0x3e, 0xfb, 0x77, 0x64, 0x0f, 0xbb, 0xb8, 0xc4, 0xed, 0xef, 0x25, 0xc4, 0x15, 0x68, 0x67,
	0x03, 0x9a, 0x12, 0xaf, 0x90, 0x60, 0x4e, 0xc4, 0x5e, 0x80, 0xcf, 0x95, 0x09, 0xb9, 0x5f, 0x3c,
	0xa1, 0x5f, 0xd6, 0x03, 0x32, 0x99, 0x9a, 0x01, 0x1d, 0x6c, 0x18, 0x67, 0x7d, 0x1e, 0x56, 0xdb,
	0x73, 0x37, 0xdb, 0xf3, 0xb9, 0x6e, 0xc5, 0xe0, 0xbe, 0x31, 0x9f, 0x68, 0x0d, 0x70, 0x15, 0xa0,
	0xf5, 0x68, 0x3e, 0x73, 0x6b, 0x0e, 0x77, 0xe7, 0x19, 0x80, 0x75, 0xc7, 0x7f, 0x0e, 0xc8, 0xd0,
	0xb4, 0x01, 0xe2, 0x3e, 0x89, 0x80, 0xbe, 0x21, 0x9d, 0x93, 0x38, 0xd6, 0x9f, 0xc9, 0xe7, 0x73,
	0x58, 0x9d, 0x13, 0x93, 0xc3, 0xda, 0x19, 0xd9, 0x0f, 0xcc, 0x7f, 0xa1, 0x51, 0x16, 0xa0, 0x1a,
	0x50, 0x0a, 0x90, 0x48, 0x61, 0x27, 0xca, 0x05, 0x19, 0x1b, 0x94, 0x53, 0xab, 0xed, 0xb3, 0xd0,
	0x7e, 0x24, 0xfd, 0x00, 0x52, 0x7e, 0x0f, 0xbb, 0x68, 0xb1, 0x3a, 0x20, 0x9d, 0x41, 0x4a, 0xc3,
	0x73, 0x59, 0x31, 0xda, 0x36, 0x9c, 0xa3, 0x3a, 0x9c, 0xf5, 0x5e, 0xff, 0x05, 0x3d, 0x23, 0xdd,
	0x05, 0x28, 0x3c, 0x1f, 0x1a, 0x80, 0xae, 0xc3, 0x5b, 0x98, 0xbc, 0x6a, 0xec, 0x4c, 0x6f, 0x45,
	0x56, 0x83, 0x05, 0xa8, 0xd3, 0x87, 0xb3, 0x64, 0xa9, 0xa7, 0x74, 0x58, 0x8f, 0x65, 0xb2, 0xbb,
	0xd1, 0x2e, 0xc9, 0x70, 0x01, 0x4a, 0xfb, 0x3a, 0x91, 0x2a, 0x89, 0x9e, 0x04, 0x57, 0x6e, 0xf6,
	0x5f, 0xd0, 0x05, 0x21, 0xe6, 0x10, 0x39, 0x0d, 0x25, 0xd0, 0x57, 0x0d, 0xae, 0x32, 0x25, 0x5b,
	0x67, 0x7f, 0x4e, 0x86, 0xa6, 0xca, 0x7e, 0x72, 0x4d, 0x50, 0x26, 0xbd, 0x15, 0xea, 0x8a, 0x8c,
	0x0c, 0x94, 0xfd, 0x6b, 0xd7, 0x61, 0xe9, 0x1e, 0x6d, 0x7a, 0xa7, 0xb7, 0xde, 0x92, 0xb1, 0xc1,
	0xbb, 0x28, 0xff, 0x5f, 0x0d, 0x7a, 0x9e, 0x2d, 0xc3, 0xdb, 0xad, 0xcc, 0xce, 0xc8, 0xf0, 0x24,
	0xcf, 0x21, 0x8b, 0x2f, 0xf1, 0xc7, 0xfe, 0xff, 0x8d, 0x6a, 0x45, 0x7c, 0x4b, 0xc6, 0xef, 0x8a,
	0x1b, 0x25, 0xc2, 0x48, 0x3d, 0x13, 0x69, 0x51, 0xce, 0xea, 0xda, 0x5e, 0x42, 0x1a, 0x80, 0x74,
	0xf9, 0x56, 0xa0, 0x73, 0xb2, 0x67, 0xf5, 0x5b, 0xdd, 0x80, 0x9e, 0x3a, 0xa5, 0x9f, 0x4a, 0x4e,
	0xe5, 0x29, 0x79, 0xd4, 0xe0, 0x05, 0x9b, 0xdf, 0xe1, 0xab, 0x92, 0x97, 0xbe, 0x2f, 0x5c, 0xe9,
	0x1f, 0xd9, 0x33, 0xd4, 0xb3, 0x5e, 0x37, 0x37, 0xb9, 0x27, 0xe2, 0x9c, 0xee, 0xff, 0x4a, 0x3f,
	0xbf, 0xfa, 0xdf, 0x74, 0x70, 0xed, 0xfb, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x4b, 0xf1,
	0xf9, 0x46, 0x0c, 0x00, 0x00,
}
