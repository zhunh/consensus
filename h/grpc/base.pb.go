// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

package base

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

//pow消息格式
type PowReqMsg struct {
	PowReqMsg            string   `protobuf:"bytes,1,opt,name=PowReqMsg,proto3" json:"PowReqMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PowReqMsg) Reset()         { *m = PowReqMsg{} }
func (m *PowReqMsg) String() string { return proto.CompactTextString(m) }
func (*PowReqMsg) ProtoMessage()    {}
func (*PowReqMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{0}
}

func (m *PowReqMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PowReqMsg.Unmarshal(m, b)
}
func (m *PowReqMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PowReqMsg.Marshal(b, m, deterministic)
}
func (m *PowReqMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PowReqMsg.Merge(m, src)
}
func (m *PowReqMsg) XXX_Size() int {
	return xxx_messageInfo_PowReqMsg.Size(m)
}
func (m *PowReqMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_PowReqMsg.DiscardUnknown(m)
}

var xxx_messageInfo_PowReqMsg proto.InternalMessageInfo

func (m *PowReqMsg) GetPowReqMsg() string {
	if m != nil {
		return m.PowReqMsg
	}
	return ""
}

type PowRspMsg struct {
	PowRspMsg            string   `protobuf:"bytes,1,opt,name=PowRspMsg,proto3" json:"PowRspMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PowRspMsg) Reset()         { *m = PowRspMsg{} }
func (m *PowRspMsg) String() string { return proto.CompactTextString(m) }
func (*PowRspMsg) ProtoMessage()    {}
func (*PowRspMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{1}
}

func (m *PowRspMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PowRspMsg.Unmarshal(m, b)
}
func (m *PowRspMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PowRspMsg.Marshal(b, m, deterministic)
}
func (m *PowRspMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PowRspMsg.Merge(m, src)
}
func (m *PowRspMsg) XXX_Size() int {
	return xxx_messageInfo_PowRspMsg.Size(m)
}
func (m *PowRspMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_PowRspMsg.DiscardUnknown(m)
}

var xxx_messageInfo_PowRspMsg proto.InternalMessageInfo

func (m *PowRspMsg) GetPowRspMsg() string {
	if m != nil {
		return m.PowRspMsg
	}
	return ""
}

//pos消息格式
type PosReqMsg struct {
	PosReqMsg            string   `protobuf:"bytes,1,opt,name=PosReqMsg,proto3" json:"PosReqMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PosReqMsg) Reset()         { *m = PosReqMsg{} }
func (m *PosReqMsg) String() string { return proto.CompactTextString(m) }
func (*PosReqMsg) ProtoMessage()    {}
func (*PosReqMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{2}
}

func (m *PosReqMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PosReqMsg.Unmarshal(m, b)
}
func (m *PosReqMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PosReqMsg.Marshal(b, m, deterministic)
}
func (m *PosReqMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PosReqMsg.Merge(m, src)
}
func (m *PosReqMsg) XXX_Size() int {
	return xxx_messageInfo_PosReqMsg.Size(m)
}
func (m *PosReqMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_PosReqMsg.DiscardUnknown(m)
}

var xxx_messageInfo_PosReqMsg proto.InternalMessageInfo

func (m *PosReqMsg) GetPosReqMsg() string {
	if m != nil {
		return m.PosReqMsg
	}
	return ""
}

type PosRspMsg struct {
	PosRspMsg            string   `protobuf:"bytes,1,opt,name=PosRspMsg,proto3" json:"PosRspMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PosRspMsg) Reset()         { *m = PosRspMsg{} }
func (m *PosRspMsg) String() string { return proto.CompactTextString(m) }
func (*PosRspMsg) ProtoMessage()    {}
func (*PosRspMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{3}
}

func (m *PosRspMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PosRspMsg.Unmarshal(m, b)
}
func (m *PosRspMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PosRspMsg.Marshal(b, m, deterministic)
}
func (m *PosRspMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PosRspMsg.Merge(m, src)
}
func (m *PosRspMsg) XXX_Size() int {
	return xxx_messageInfo_PosRspMsg.Size(m)
}
func (m *PosRspMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_PosRspMsg.DiscardUnknown(m)
}

var xxx_messageInfo_PosRspMsg proto.InternalMessageInfo

func (m *PosRspMsg) GetPosRspMsg() string {
	if m != nil {
		return m.PosRspMsg
	}
	return ""
}

//难度值
type NonceReq struct {
	GetNonce             string   `protobuf:"bytes,1,opt,name=GetNonce,proto3" json:"GetNonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NonceReq) Reset()         { *m = NonceReq{} }
func (m *NonceReq) String() string { return proto.CompactTextString(m) }
func (*NonceReq) ProtoMessage()    {}
func (*NonceReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{4}
}

func (m *NonceReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonceReq.Unmarshal(m, b)
}
func (m *NonceReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonceReq.Marshal(b, m, deterministic)
}
func (m *NonceReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonceReq.Merge(m, src)
}
func (m *NonceReq) XXX_Size() int {
	return xxx_messageInfo_NonceReq.Size(m)
}
func (m *NonceReq) XXX_DiscardUnknown() {
	xxx_messageInfo_NonceReq.DiscardUnknown(m)
}

var xxx_messageInfo_NonceReq proto.InternalMessageInfo

func (m *NonceReq) GetGetNonce() string {
	if m != nil {
		return m.GetNonce
	}
	return ""
}

type NonceRsp struct {
	Difficulty           string   `protobuf:"bytes,1,opt,name=Difficulty,proto3" json:"Difficulty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NonceRsp) Reset()         { *m = NonceRsp{} }
func (m *NonceRsp) String() string { return proto.CompactTextString(m) }
func (*NonceRsp) ProtoMessage()    {}
func (*NonceRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{5}
}

func (m *NonceRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonceRsp.Unmarshal(m, b)
}
func (m *NonceRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonceRsp.Marshal(b, m, deterministic)
}
func (m *NonceRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonceRsp.Merge(m, src)
}
func (m *NonceRsp) XXX_Size() int {
	return xxx_messageInfo_NonceRsp.Size(m)
}
func (m *NonceRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_NonceRsp.DiscardUnknown(m)
}

var xxx_messageInfo_NonceRsp proto.InternalMessageInfo

func (m *NonceRsp) GetDifficulty() string {
	if m != nil {
		return m.Difficulty
	}
	return ""
}

//记账节点确认
type LedgerReq struct {
	Nonce                string   `protobuf:"bytes,1,opt,name=Nonce,proto3" json:"Nonce,omitempty"`
	CostTime             string   `protobuf:"bytes,2,opt,name=CostTime,proto3" json:"CostTime,omitempty"`
	Result               string   `protobuf:"bytes,3,opt,name=Result,proto3" json:"Result,omitempty"`
	Ip                   string   `protobuf:"bytes,4,opt,name=Ip,proto3" json:"Ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LedgerReq) Reset()         { *m = LedgerReq{} }
func (m *LedgerReq) String() string { return proto.CompactTextString(m) }
func (*LedgerReq) ProtoMessage()    {}
func (*LedgerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{6}
}

func (m *LedgerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LedgerReq.Unmarshal(m, b)
}
func (m *LedgerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LedgerReq.Marshal(b, m, deterministic)
}
func (m *LedgerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LedgerReq.Merge(m, src)
}
func (m *LedgerReq) XXX_Size() int {
	return xxx_messageInfo_LedgerReq.Size(m)
}
func (m *LedgerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LedgerReq.DiscardUnknown(m)
}

var xxx_messageInfo_LedgerReq proto.InternalMessageInfo

func (m *LedgerReq) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *LedgerReq) GetCostTime() string {
	if m != nil {
		return m.CostTime
	}
	return ""
}

func (m *LedgerReq) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *LedgerReq) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type LedgerRsp struct {
	LedgerRsp            string   `protobuf:"bytes,1,opt,name=LedgerRsp,proto3" json:"LedgerRsp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LedgerRsp) Reset()         { *m = LedgerRsp{} }
func (m *LedgerRsp) String() string { return proto.CompactTextString(m) }
func (*LedgerRsp) ProtoMessage()    {}
func (*LedgerRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{7}
}

func (m *LedgerRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LedgerRsp.Unmarshal(m, b)
}
func (m *LedgerRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LedgerRsp.Marshal(b, m, deterministic)
}
func (m *LedgerRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LedgerRsp.Merge(m, src)
}
func (m *LedgerRsp) XXX_Size() int {
	return xxx_messageInfo_LedgerRsp.Size(m)
}
func (m *LedgerRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_LedgerRsp.DiscardUnknown(m)
}

var xxx_messageInfo_LedgerRsp proto.InternalMessageInfo

func (m *LedgerRsp) GetLedgerRsp() string {
	if m != nil {
		return m.LedgerRsp
	}
	return ""
}

//注册节点
type RegisterReq struct {
	NodePort             string   `protobuf:"bytes,1,opt,name=nodePort,proto3" json:"nodePort,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{8}
}

func (m *RegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq.Unmarshal(m, b)
}
func (m *RegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq.Marshal(b, m, deterministic)
}
func (m *RegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq.Merge(m, src)
}
func (m *RegisterReq) XXX_Size() int {
	return xxx_messageInfo_RegisterReq.Size(m)
}
func (m *RegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq proto.InternalMessageInfo

func (m *RegisterReq) GetNodePort() string {
	if m != nil {
		return m.NodePort
	}
	return ""
}

func (m *RegisterReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RegisterRsp struct {
	RegisterRsp          string   `protobuf:"bytes,1,opt,name=RegisterRsp,proto3" json:"RegisterRsp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRsp) Reset()         { *m = RegisterRsp{} }
func (m *RegisterRsp) String() string { return proto.CompactTextString(m) }
func (*RegisterRsp) ProtoMessage()    {}
func (*RegisterRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{9}
}

func (m *RegisterRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRsp.Unmarshal(m, b)
}
func (m *RegisterRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRsp.Marshal(b, m, deterministic)
}
func (m *RegisterRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRsp.Merge(m, src)
}
func (m *RegisterRsp) XXX_Size() int {
	return xxx_messageInfo_RegisterRsp.Size(m)
}
func (m *RegisterRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRsp proto.InternalMessageInfo

func (m *RegisterRsp) GetRegisterRsp() string {
	if m != nil {
		return m.RegisterRsp
	}
	return ""
}

//挖矿请求
type StartMiningReq struct {
	DT                   string   `protobuf:"bytes,1,opt,name=DT,proto3" json:"DT,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartMiningReq) Reset()         { *m = StartMiningReq{} }
func (m *StartMiningReq) String() string { return proto.CompactTextString(m) }
func (*StartMiningReq) ProtoMessage()    {}
func (*StartMiningReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{10}
}

func (m *StartMiningReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartMiningReq.Unmarshal(m, b)
}
func (m *StartMiningReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartMiningReq.Marshal(b, m, deterministic)
}
func (m *StartMiningReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartMiningReq.Merge(m, src)
}
func (m *StartMiningReq) XXX_Size() int {
	return xxx_messageInfo_StartMiningReq.Size(m)
}
func (m *StartMiningReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StartMiningReq.DiscardUnknown(m)
}

var xxx_messageInfo_StartMiningReq proto.InternalMessageInfo

func (m *StartMiningReq) GetDT() string {
	if m != nil {
		return m.DT
	}
	return ""
}

type StartMiningRsp struct {
	StartMiningRsp       string   `protobuf:"bytes,1,opt,name=StartMiningRsp,proto3" json:"StartMiningRsp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartMiningRsp) Reset()         { *m = StartMiningRsp{} }
func (m *StartMiningRsp) String() string { return proto.CompactTextString(m) }
func (*StartMiningRsp) ProtoMessage()    {}
func (*StartMiningRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_db1b6b0986796150, []int{11}
}

func (m *StartMiningRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartMiningRsp.Unmarshal(m, b)
}
func (m *StartMiningRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartMiningRsp.Marshal(b, m, deterministic)
}
func (m *StartMiningRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartMiningRsp.Merge(m, src)
}
func (m *StartMiningRsp) XXX_Size() int {
	return xxx_messageInfo_StartMiningRsp.Size(m)
}
func (m *StartMiningRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_StartMiningRsp.DiscardUnknown(m)
}

var xxx_messageInfo_StartMiningRsp proto.InternalMessageInfo

func (m *StartMiningRsp) GetStartMiningRsp() string {
	if m != nil {
		return m.StartMiningRsp
	}
	return ""
}

func init() {
	proto.RegisterType((*PowReqMsg)(nil), "PowReqMsg")
	proto.RegisterType((*PowRspMsg)(nil), "PowRspMsg")
	proto.RegisterType((*PosReqMsg)(nil), "PosReqMsg")
	proto.RegisterType((*PosRspMsg)(nil), "PosRspMsg")
	proto.RegisterType((*NonceReq)(nil), "NonceReq")
	proto.RegisterType((*NonceRsp)(nil), "NonceRsp")
	proto.RegisterType((*LedgerReq)(nil), "LedgerReq")
	proto.RegisterType((*LedgerRsp)(nil), "LedgerRsp")
	proto.RegisterType((*RegisterReq)(nil), "RegisterReq")
	proto.RegisterType((*RegisterRsp)(nil), "RegisterRsp")
	proto.RegisterType((*StartMiningReq)(nil), "StartMiningReq")
	proto.RegisterType((*StartMiningRsp)(nil), "StartMiningRsp")
}

func init() { proto.RegisterFile("base.proto", fileDescriptor_db1b6b0986796150) }

var fileDescriptor_db1b6b0986796150 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0x5f, 0xcb, 0xd3, 0x30,
	0x14, 0xc6, 0x59, 0xd5, 0x97, 0xed, 0xbc, 0x63, 0x83, 0x20, 0x52, 0x8a, 0xc8, 0x08, 0x32, 0x74,
	0x17, 0x1d, 0xcc, 0x1b, 0xef, 0x04, 0x1d, 0xc8, 0xc0, 0x8d, 0xd2, 0xed, 0x0b, 0x74, 0xed, 0x59,
	0x09, 0xce, 0x26, 0xcb, 0xc9, 0x1c, 0x7e, 0x08, 0xbf, 0xb3, 0xb4, 0x4d, 0x6a, 0xda, 0xbb, 0x3c,
	0x4f, 0x7f, 0x3c, 0xe7, 0x0f, 0xa7, 0x00, 0xe7, 0x8c, 0x30, 0x56, 0x5a, 0x1a, 0xc9, 0x3f, 0xc2,
	0x24, 0x91, 0x8f, 0x14, 0x6f, 0x7b, 0x2a, 0xd9, 0x5b, 0x4f, 0x84, 0xa3, 0xc5, 0xe8, 0xc3, 0x24,
	0xfd, 0x6f, 0x38, 0x94, 0x94, 0x87, 0x36, 0xc2, 0x47, 0x1b, 0xa3, 0x45, 0xc9, 0x4f, 0xa5, 0x61,
	0x2a, 0xf9, 0xa9, 0xe4, 0xa7, 0xd2, 0x30, 0xd5, 0x1a, 0x7c, 0x09, 0xe3, 0x83, 0xac, 0x72, 0x4c,
	0xf1, 0xc6, 0x22, 0x18, 0x7f, 0x47, 0xd3, 0x48, 0x0b, 0x76, 0x9a, 0xaf, 0x1c, 0x47, 0x8a, 0xbd,
	0x03, 0xd8, 0x8a, 0xcb, 0x45, 0xe4, 0xf7, 0xab, 0xf9, 0x63, 0x49, 0xcf, 0xe1, 0x08, 0x93, 0x1f,
	0x58, 0x94, 0xa8, 0xeb, 0xd0, 0xd7, 0xf0, 0xca, 0x4f, 0x6c, 0x45, 0x5d, 0xea, 0x9b, 0x24, 0x73,
	0x12, 0xbf, 0x30, 0x0c, 0xda, 0x52, 0x4e, 0xb3, 0x37, 0xf0, 0x94, 0x22, 0xdd, 0xaf, 0x26, 0x7c,
	0xd1, 0x7c, 0xb1, 0x8a, 0xcd, 0x20, 0xd8, 0xa9, 0xf0, 0x65, 0xe3, 0x05, 0x3b, 0x55, 0x4f, 0x69,
	0xcb, 0x90, 0xaa, 0xa7, 0xec, 0x84, 0x9b, 0xb2, 0x33, 0xf8, 0x17, 0x78, 0x4e, 0xb1, 0x14, 0x64,
	0xda, 0x9e, 0x22, 0x18, 0x57, 0xb2, 0xc0, 0x44, 0x6a, 0xe3, 0x06, 0x75, 0xba, 0xee, 0xd7, 0xc8,
	0x9f, 0x58, 0xd9, 0xb6, 0x5a, 0xc1, 0xd7, 0x5e, 0x00, 0x29, 0xb6, 0xe8, 0x49, 0x9b, 0xe1, 0x5b,
	0x7c, 0x01, 0xb3, 0xa3, 0xc9, 0xb4, 0xd9, 0x8b, 0x4a, 0x54, 0x65, 0x5d, 0x74, 0x06, 0xc1, 0xf6,
	0x64, 0xd1, 0x60, 0x7b, 0xe2, 0x9f, 0xfb, 0x04, 0x29, 0xb6, 0x1c, 0x3a, 0x96, 0x1e, 0xb8, 0x9b,
	0xbf, 0x23, 0x80, 0x83, 0x2c, 0xf0, 0x88, 0xfa, 0x37, 0x6a, 0xb6, 0x84, 0x69, 0x22, 0x1f, 0x5f,
	0xb5, 0xcc, 0x8a, 0x3c, 0x23, 0xc3, 0x20, 0xee, 0xee, 0x2b, 0x6a, 0xdf, 0xed, 0x21, 0x34, 0x1c,
	0xf5, 0x39, 0xf2, 0x38, 0x77, 0x30, 0x6b, 0x78, 0xf6, 0x0a, 0xb2, 0x79, 0xdc, 0x1f, 0x24, 0xea,
	0x1b, 0xa4, 0x36, 0x39, 0xcc, 0xf7, 0x99, 0xa8, 0x5c, 0x4b, 0x22, 0x47, 0xb6, 0x82, 0xa9, 0xdb,
	0x46, 0x6d, 0xb3, 0x69, 0xec, 0xed, 0x3f, 0xf2, 0x14, 0x29, 0xf6, 0x1e, 0x20, 0xc5, 0x5c, 0xea,
	0x22, 0x41, 0xd4, 0x0c, 0xe2, 0xee, 0x76, 0xa2, 0xee, 0x4d, 0xea, 0xfc, 0xd4, 0xfc, 0x5b, 0x9f,
	0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x48, 0x03, 0xc3, 0x65, 0x69, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeServerClient is the client API for NodeServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeServerClient interface {
	PowBroadcast(ctx context.Context, in *PowReqMsg, opts ...grpc.CallOption) (*PowRspMsg, error)
	PosBroadcast(ctx context.Context, in *PosReqMsg, opts ...grpc.CallOption) (*PosRspMsg, error)
	StartMining(ctx context.Context, in *StartMiningReq, opts ...grpc.CallOption) (*StartMiningRsp, error)
}

type nodeServerClient struct {
	cc *grpc.ClientConn
}

func NewNodeServerClient(cc *grpc.ClientConn) NodeServerClient {
	return &nodeServerClient{cc}
}

func (c *nodeServerClient) PowBroadcast(ctx context.Context, in *PowReqMsg, opts ...grpc.CallOption) (*PowRspMsg, error) {
	out := new(PowRspMsg)
	err := c.cc.Invoke(ctx, "/NodeServer/PowBroadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServerClient) PosBroadcast(ctx context.Context, in *PosReqMsg, opts ...grpc.CallOption) (*PosRspMsg, error) {
	out := new(PosRspMsg)
	err := c.cc.Invoke(ctx, "/NodeServer/PosBroadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServerClient) StartMining(ctx context.Context, in *StartMiningReq, opts ...grpc.CallOption) (*StartMiningRsp, error) {
	out := new(StartMiningRsp)
	err := c.cc.Invoke(ctx, "/NodeServer/StartMining", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServerServer is the server API for NodeServer service.
type NodeServerServer interface {
	PowBroadcast(context.Context, *PowReqMsg) (*PowRspMsg, error)
	PosBroadcast(context.Context, *PosReqMsg) (*PosRspMsg, error)
	StartMining(context.Context, *StartMiningReq) (*StartMiningRsp, error)
}

// UnimplementedNodeServerServer can be embedded to have forward compatible implementations.
type UnimplementedNodeServerServer struct {
}

func (*UnimplementedNodeServerServer) PowBroadcast(ctx context.Context, req *PowReqMsg) (*PowRspMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PowBroadcast not implemented")
}
func (*UnimplementedNodeServerServer) PosBroadcast(ctx context.Context, req *PosReqMsg) (*PosRspMsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PosBroadcast not implemented")
}
func (*UnimplementedNodeServerServer) StartMining(ctx context.Context, req *StartMiningReq) (*StartMiningRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMining not implemented")
}

func RegisterNodeServerServer(s *grpc.Server, srv NodeServerServer) {
	s.RegisterService(&_NodeServer_serviceDesc, srv)
}

func _NodeServer_PowBroadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PowReqMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServerServer).PowBroadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NodeServer/PowBroadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServerServer).PowBroadcast(ctx, req.(*PowReqMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeServer_PosBroadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PosReqMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServerServer).PosBroadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NodeServer/PosBroadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServerServer).PosBroadcast(ctx, req.(*PosReqMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeServer_StartMining_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMiningReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServerServer).StartMining(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/NodeServer/StartMining",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServerServer).StartMining(ctx, req.(*StartMiningReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "NodeServer",
	HandlerType: (*NodeServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PowBroadcast",
			Handler:    _NodeServer_PowBroadcast_Handler,
		},
		{
			MethodName: "PosBroadcast",
			Handler:    _NodeServer_PosBroadcast_Handler,
		},
		{
			MethodName: "StartMining",
			Handler:    _NodeServer_StartMining_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base.proto",
}

// MainNodeServiceClient is the client API for MainNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MainNodeServiceClient interface {
	//注册节点
	RegisterNode(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRsp, error)
	//记账权确认
	RecordPeer(ctx context.Context, in *LedgerReq, opts ...grpc.CallOption) (*LedgerRsp, error)
}

type mainNodeServiceClient struct {
	cc *grpc.ClientConn
}

func NewMainNodeServiceClient(cc *grpc.ClientConn) MainNodeServiceClient {
	return &mainNodeServiceClient{cc}
}

func (c *mainNodeServiceClient) RegisterNode(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterRsp, error) {
	out := new(RegisterRsp)
	err := c.cc.Invoke(ctx, "/MainNodeService/RegisterNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainNodeServiceClient) RecordPeer(ctx context.Context, in *LedgerReq, opts ...grpc.CallOption) (*LedgerRsp, error) {
	out := new(LedgerRsp)
	err := c.cc.Invoke(ctx, "/MainNodeService/RecordPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MainNodeServiceServer is the server API for MainNodeService service.
type MainNodeServiceServer interface {
	//注册节点
	RegisterNode(context.Context, *RegisterReq) (*RegisterRsp, error)
	//记账权确认
	RecordPeer(context.Context, *LedgerReq) (*LedgerRsp, error)
}

// UnimplementedMainNodeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMainNodeServiceServer struct {
}

func (*UnimplementedMainNodeServiceServer) RegisterNode(ctx context.Context, req *RegisterReq) (*RegisterRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNode not implemented")
}
func (*UnimplementedMainNodeServiceServer) RecordPeer(ctx context.Context, req *LedgerReq) (*LedgerRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordPeer not implemented")
}

func RegisterMainNodeServiceServer(s *grpc.Server, srv MainNodeServiceServer) {
	s.RegisterService(&_MainNodeService_serviceDesc, srv)
}

func _MainNodeService_RegisterNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainNodeServiceServer).RegisterNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MainNodeService/RegisterNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainNodeServiceServer).RegisterNode(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MainNodeService_RecordPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LedgerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainNodeServiceServer).RecordPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MainNodeService/RecordPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainNodeServiceServer).RecordPeer(ctx, req.(*LedgerReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _MainNodeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MainNodeService",
	HandlerType: (*MainNodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNode",
			Handler:    _MainNodeService_RegisterNode_Handler,
		},
		{
			MethodName: "RecordPeer",
			Handler:    _MainNodeService_RecordPeer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base.proto",
}
