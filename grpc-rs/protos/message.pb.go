// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package proto

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

type Message_MessageType int32

const (
	Message_DEFAULT Message_MessageType = 0
	// Component message
	Message_COMPONENT_CONTRACT_REGISTER  Message_MessageType = 100
	Message_COMPONENT_CONSENSUS_REGISTER Message_MessageType = 101
	Message_COMPONENT_UNREGISTER         Message_MessageType = 102
	// Peer message
	Message_PEER_HANDSHAKE_REQUEST      Message_MessageType = 200
	Message_PEER_HANDSHAKE_RESPONSE     Message_MessageType = 201
	Message_PEER_HEARTBEAT_REQUEST      Message_MessageType = 202
	Message_PEER_HEARTBEAT_RESPONSE     Message_MessageType = 203
	Message_PEER_BROADCAST_TRANSACTION  Message_MessageType = 204
	Message_PEER_BROADCAST_BLOCK_NUMBER Message_MessageType = 205
	// pull blocks from other peer nodes
	Message_PEER_DELIVER_BLOCK Message_MessageType = 206
	// Consensus notification messages
	Message_CONSENSUS_TRANSACTION_ARRIVED Message_MessageType = 300
	Message_CONSENSUS_NOTIFY_BLOCK_COMMIT Message_MessageType = 301
)

var Message_MessageType_name = map[int32]string{
	0:   "DEFAULT",
	100: "COMPONENT_CONTRACT_REGISTER",
	101: "COMPONENT_CONSENSUS_REGISTER",
	102: "COMPONENT_UNREGISTER",
	200: "PEER_HANDSHAKE_REQUEST",
	201: "PEER_HANDSHAKE_RESPONSE",
	202: "PEER_HEARTBEAT_REQUEST",
	203: "PEER_HEARTBEAT_RESPONSE",
	204: "PEER_BROADCAST_TRANSACTION",
	205: "PEER_BROADCAST_BLOCK_NUMBER",
	206: "PEER_DELIVER_BLOCK",
	300: "CONSENSUS_TRANSACTION_ARRIVED",
	301: "CONSENSUS_NOTIFY_BLOCK_COMMIT",
}

var Message_MessageType_value = map[string]int32{
	"DEFAULT":                       0,
	"COMPONENT_CONTRACT_REGISTER":   100,
	"COMPONENT_CONSENSUS_REGISTER":  101,
	"COMPONENT_UNREGISTER":          102,
	"PEER_HANDSHAKE_REQUEST":        200,
	"PEER_HANDSHAKE_RESPONSE":       201,
	"PEER_HEARTBEAT_REQUEST":        202,
	"PEER_HEARTBEAT_RESPONSE":       203,
	"PEER_BROADCAST_TRANSACTION":    204,
	"PEER_BROADCAST_BLOCK_NUMBER":   205,
	"PEER_DELIVER_BLOCK":            206,
	"CONSENSUS_TRANSACTION_ARRIVED": 300,
	"CONSENSUS_NOTIFY_BLOCK_COMMIT": 301,
}

func (x Message_MessageType) String() string {
	return proto.EnumName(Message_MessageType_name, int32(x))
}

func (Message_MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}

type Message struct {
	MessageType Message_MessageType `protobuf:"varint,1,opt,name=message_type,json=messageType,proto3,enum=proto.Message_MessageType" json:"message_type,omitempty"`
	// The identifier used to correlate response messages to their related
	// request messages.  correlation_id should be set to a random string
	// for messages which are not responses to previously sent messages.  For
	// response messages, correlation_id should be set to the same string as
	// contained in the request message.
	CorrelationId string `protobuf:"bytes,2,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	// The content of the message, defined by message_type.  In many
	// cases, this data has been serialized with Protocol Buffers
	Content              []byte   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMessageType() Message_MessageType {
	if m != nil {
		return m.MessageType
	}
	return Message_DEFAULT
}

func (m *Message) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *Message) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type ComponentContractRegister struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Decorations          map[string][]byte `protobuf:"bytes,2,rep,name=decorations,proto3" json:"decorations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ComponentContractRegister) Reset()         { *m = ComponentContractRegister{} }
func (m *ComponentContractRegister) String() string { return proto.CompactTextString(m) }
func (*ComponentContractRegister) ProtoMessage()    {}
func (*ComponentContractRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *ComponentContractRegister) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentContractRegister.Unmarshal(m, b)
}
func (m *ComponentContractRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentContractRegister.Marshal(b, m, deterministic)
}
func (m *ComponentContractRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentContractRegister.Merge(m, src)
}
func (m *ComponentContractRegister) XXX_Size() int {
	return xxx_messageInfo_ComponentContractRegister.Size(m)
}
func (m *ComponentContractRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentContractRegister.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentContractRegister proto.InternalMessageInfo

func (m *ComponentContractRegister) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ComponentContractRegister) GetDecorations() map[string][]byte {
	if m != nil {
		return m.Decorations
	}
	return nil
}

type ComponentConsensusRegister struct {
	Alg                  string            `protobuf:"bytes,1,opt,name=alg,proto3" json:"alg,omitempty"`
	Decorations          map[string][]byte `protobuf:"bytes,2,rep,name=decorations,proto3" json:"decorations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ComponentConsensusRegister) Reset()         { *m = ComponentConsensusRegister{} }
func (m *ComponentConsensusRegister) String() string { return proto.CompactTextString(m) }
func (*ComponentConsensusRegister) ProtoMessage()    {}
func (*ComponentConsensusRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *ComponentConsensusRegister) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentConsensusRegister.Unmarshal(m, b)
}
func (m *ComponentConsensusRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentConsensusRegister.Marshal(b, m, deterministic)
}
func (m *ComponentConsensusRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentConsensusRegister.Merge(m, src)
}
func (m *ComponentConsensusRegister) XXX_Size() int {
	return xxx_messageInfo_ComponentConsensusRegister.Size(m)
}
func (m *ComponentConsensusRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentConsensusRegister.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentConsensusRegister proto.InternalMessageInfo

func (m *ComponentConsensusRegister) GetAlg() string {
	if m != nil {
		return m.Alg
	}
	return ""
}

func (m *ComponentConsensusRegister) GetDecorations() map[string][]byte {
	if m != nil {
		return m.Decorations
	}
	return nil
}

type ComponentUnregister struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComponentUnregister) Reset()         { *m = ComponentUnregister{} }
func (m *ComponentUnregister) String() string { return proto.CompactTextString(m) }
func (*ComponentUnregister) ProtoMessage()    {}
func (*ComponentUnregister) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *ComponentUnregister) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentUnregister.Unmarshal(m, b)
}
func (m *ComponentUnregister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentUnregister.Marshal(b, m, deterministic)
}
func (m *ComponentUnregister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentUnregister.Merge(m, src)
}
func (m *ComponentUnregister) XXX_Size() int {
	return xxx_messageInfo_ComponentUnregister.Size(m)
}
func (m *ComponentUnregister) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentUnregister.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentUnregister proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("proto.Message_MessageType", Message_MessageType_name, Message_MessageType_value)
	proto.RegisterType((*Message)(nil), "proto.Message")
	proto.RegisterType((*ComponentContractRegister)(nil), "proto.ComponentContractRegister")
	proto.RegisterMapType((map[string][]byte)(nil), "proto.ComponentContractRegister.DecorationsEntry")
	proto.RegisterType((*ComponentConsensusRegister)(nil), "proto.ComponentConsensusRegister")
	proto.RegisterMapType((map[string][]byte)(nil), "proto.ComponentConsensusRegister.DecorationsEntry")
	proto.RegisterType((*ComponentUnregister)(nil), "proto.ComponentUnregister")
}

func init() {
	proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd)
}

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 520 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0xfe, 0x27, 0x69, 0xff, 0xaa, 0x27, 0x6d, 0x35, 0x3a, 0x14, 0x6a, 0x92, 0xa2, 0x5a, 0x91,
	0x90, 0xb2, 0x8a, 0x44, 0xd8, 0x20, 0x24, 0x90, 0x26, 0xf6, 0x94, 0x5a, 0x4d, 0xec, 0x30, 0x1e,
	0x57, 0x62, 0x65, 0x99, 0x64, 0x88, 0x22, 0x12, 0x3b, 0xb2, 0x5d, 0xa4, 0x3c, 0x06, 0x0f, 0xc0,
	0x1b, 0xc0, 0x5b, 0xb0, 0xe0, 0xfe, 0x4c, 0x28, 0x8e, 0x13, 0x3b, 0xa5, 0x2c, 0x59, 0xcd, 0x1c,
	0x7f, 0x37, 0x7f, 0xf6, 0x81, 0xc3, 0x99, 0x4a, 0x92, 0x60, 0xac, 0xda, 0xf3, 0x38, 0x4a, 0x23,
	0xdc, 0xcd, 0x8e, 0xe6, 0x87, 0x1d, 0xd8, 0xeb, 0xaf, 0x00, 0x7c, 0x06, 0x07, 0x39, 0xc7, 0x4f,
	0x17, 0x73, 0xa5, 0x11, 0x9d, 0xb4, 0x8e, 0x3a, 0xf5, 0x95, 0xa0, 0x9d, 0xb3, 0xd6, 0xa7, 0x5c,
	0xcc, 0x95, 0xa8, 0xcd, 0x8a, 0x01, 0x1f, 0xc2, 0xd1, 0x30, 0x8a, 0x63, 0x35, 0x0d, 0xd2, 0x49,
	0x14, 0xfa, 0x93, 0x91, 0x56, 0xd1, 0x49, 0x6b, 0x5f, 0x1c, 0x96, 0x9e, 0x5a, 0x23, 0xd4, 0x60,
	0x6f, 0x18, 0x85, 0xa9, 0x0a, 0x53, 0xad, 0xaa, 0x93, 0xd6, 0x81, 0x58, 0x8f, 0xcd, 0xf7, 0x55,
	0xa8, 0x95, 0xdc, 0xb1, 0x06, 0x7b, 0x26, 0x3f, 0x67, 0x5e, 0x4f, 0xd2, 0xff, 0xf0, 0x0c, 0x1a,
	0x86, 0xd3, 0x1f, 0x38, 0x36, 0xb7, 0xa5, 0x6f, 0x38, 0xb6, 0x14, 0xcc, 0x90, 0xbe, 0xe0, 0x2f,
	0x2c, 0x57, 0x72, 0x41, 0x47, 0xa8, 0xc3, 0xe9, 0x16, 0xc1, 0xe5, 0xb6, 0xeb, 0xb9, 0x05, 0x43,
	0xa1, 0x06, 0xc7, 0x05, 0xc3, 0xb3, 0x37, 0xc8, 0x1b, 0x6c, 0xc0, 0xbd, 0x01, 0xe7, 0xc2, 0xbf,
	0x60, 0xb6, 0xe9, 0x5e, 0xb0, 0x4b, 0xee, 0x0b, 0xfe, 0xd2, 0xe3, 0xae, 0xa4, 0x5f, 0x08, 0x9e,
	0xc2, 0xc9, 0x1f, 0xa0, 0x3b, 0x58, 0x06, 0xd0, 0xaf, 0xa4, 0x90, 0x72, 0x26, 0x64, 0x97, 0x33,
	0xb9, 0x91, 0x7e, 0x2b, 0x49, 0x4b, 0x60, 0x2e, 0xfd, 0x4e, 0xf0, 0x0c, 0xea, 0x19, 0xda, 0x15,
	0x0e, 0x33, 0x0d, 0xe6, 0x4a, 0x5f, 0x0a, 0x66, 0xbb, 0xcc, 0x90, 0x96, 0x63, 0xd3, 0x1f, 0x04,
	0x75, 0x68, 0xdc, 0x20, 0x74, 0x7b, 0x8e, 0x71, 0xe9, 0xdb, 0x5e, 0xbf, 0xcb, 0x05, 0xfd, 0x49,
	0xf0, 0x04, 0x30, 0x63, 0x98, 0xbc, 0x67, 0x5d, 0x2d, 0x99, 0x4b, 0x9c, 0xfe, 0x22, 0xd8, 0x84,
	0x07, 0xc5, 0x37, 0x28, 0xd9, 0xfa, 0x4c, 0x08, 0xeb, 0x8a, 0x9b, 0xf4, 0x63, 0x65, 0x9b, 0x63,
	0x3b, 0xd2, 0x3a, 0x7f, 0x95, 0x07, 0x18, 0x4e, 0xbf, 0x6f, 0x49, 0xfa, 0xa9, 0xd2, 0xfc, 0x4c,
	0xe0, 0xbe, 0x11, 0xcd, 0xe6, 0x51, 0xa8, 0xc2, 0xd4, 0x88, 0xc2, 0x34, 0x0e, 0x86, 0xa9, 0x50,
	0xe3, 0x49, 0x92, 0xaa, 0x18, 0x11, 0x76, 0xc2, 0x60, 0xb6, 0xda, 0x94, 0x7d, 0x91, 0xdd, 0xd1,
	0x85, 0xda, 0x48, 0x0d, 0xa3, 0x38, 0xfb, 0xdf, 0x89, 0x56, 0xd1, 0xab, 0xad, 0x5a, 0xe7, 0x51,
	0xbe, 0x44, 0x7f, 0xb5, 0x6a, 0x9b, 0x85, 0x86, 0x87, 0x69, 0xbc, 0x10, 0x65, 0x97, 0xfa, 0x73,
	0xa0, 0x37, 0x09, 0x48, 0xa1, 0xfa, 0x56, 0x2d, 0xf2, 0xec, 0xe5, 0x15, 0x8f, 0x61, 0xf7, 0x5d,
	0x30, 0xbd, 0x56, 0xd9, 0xe2, 0x1d, 0x88, 0xd5, 0xf0, 0xb4, 0xf2, 0x84, 0x2c, 0x6b, 0xd4, 0xcb,
	0xd9, 0x89, 0x0a, 0x93, 0xeb, 0x64, 0xd3, 0x83, 0x42, 0x35, 0x98, 0x8e, 0xd7, 0x56, 0xc1, 0x74,
	0x8c, 0xf2, 0xb6, 0x16, 0x9d, 0x5b, 0x5a, 0x6c, 0x3b, 0xfd, 0xe3, 0x1a, 0x77, 0xe1, 0xce, 0x26,
	0xdb, 0x0b, 0xe3, 0x3c, 0xf4, 0xf5, 0xff, 0xd9, 0x6b, 0x3d, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff,
	0x5e, 0x17, 0xbc, 0x7e, 0xe3, 0x03, 0x00, 0x00,
}