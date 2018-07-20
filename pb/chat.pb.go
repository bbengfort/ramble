// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	chat.proto

It has these top-level messages:
	ChatMessage
	PingRequest
	PingReply
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ChatMessage struct {
	Sequence  int64  `protobuf:"varint,1,opt,name=sequence" json:"sequence,omitempty"`
	Sender    string `protobuf:"bytes,2,opt,name=sender" json:"sender,omitempty"`
	Timestamp string `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Message   string `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
}

func (m *ChatMessage) Reset()                    { *m = ChatMessage{} }
func (m *ChatMessage) String() string            { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()               {}
func (*ChatMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ChatMessage) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *ChatMessage) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *ChatMessage) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *ChatMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PingRequest struct {
	Sequence int64  `protobuf:"varint,1,opt,name=sequence" json:"sequence,omitempty"`
	Sender   string `protobuf:"bytes,2,opt,name=sender" json:"sender,omitempty"`
	Ttl      int64  `protobuf:"varint,3,opt,name=ttl" json:"ttl,omitempty"`
	Payload  []byte `protobuf:"bytes,15,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingRequest) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *PingRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *PingRequest) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *PingRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type PingReply struct {
	Sequence int64        `protobuf:"varint,1,opt,name=sequence" json:"sequence,omitempty"`
	Receiver string       `protobuf:"bytes,2,opt,name=receiver" json:"receiver,omitempty"`
	Echo     *PingRequest `protobuf:"bytes,3,opt,name=echo" json:"echo,omitempty"`
}

func (m *PingReply) Reset()                    { *m = PingReply{} }
func (m *PingReply) String() string            { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()               {}
func (*PingReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PingReply) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *PingReply) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *PingReply) GetEcho() *PingRequest {
	if m != nil {
		return m.Echo
	}
	return nil
}

func init() {
	proto.RegisterType((*ChatMessage)(nil), "pb.ChatMessage")
	proto.RegisterType((*PingRequest)(nil), "pb.PingRequest")
	proto.RegisterType((*PingReply)(nil), "pb.PingReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Ramble service

type RambleClient interface {
	Chat(ctx context.Context, opts ...grpc.CallOption) (Ramble_ChatClient, error)
	Ping(ctx context.Context, opts ...grpc.CallOption) (Ramble_PingClient, error)
}

type rambleClient struct {
	cc *grpc.ClientConn
}

func NewRambleClient(cc *grpc.ClientConn) RambleClient {
	return &rambleClient{cc}
}

func (c *rambleClient) Chat(ctx context.Context, opts ...grpc.CallOption) (Ramble_ChatClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Ramble_serviceDesc.Streams[0], c.cc, "/pb.Ramble/Chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &rambleChatClient{stream}
	return x, nil
}

type Ramble_ChatClient interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type rambleChatClient struct {
	grpc.ClientStream
}

func (x *rambleChatClient) Send(m *ChatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *rambleChatClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rambleClient) Ping(ctx context.Context, opts ...grpc.CallOption) (Ramble_PingClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Ramble_serviceDesc.Streams[1], c.cc, "/pb.Ramble/Ping", opts...)
	if err != nil {
		return nil, err
	}
	x := &ramblePingClient{stream}
	return x, nil
}

type Ramble_PingClient interface {
	Send(*PingRequest) error
	CloseAndRecv() (*PingReply, error)
	grpc.ClientStream
}

type ramblePingClient struct {
	grpc.ClientStream
}

func (x *ramblePingClient) Send(m *PingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ramblePingClient) CloseAndRecv() (*PingReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PingReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Ramble service

type RambleServer interface {
	Chat(Ramble_ChatServer) error
	Ping(Ramble_PingServer) error
}

func RegisterRambleServer(s *grpc.Server, srv RambleServer) {
	s.RegisterService(&_Ramble_serviceDesc, srv)
}

func _Ramble_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RambleServer).Chat(&rambleChatServer{stream})
}

type Ramble_ChatServer interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ServerStream
}

type rambleChatServer struct {
	grpc.ServerStream
}

func (x *rambleChatServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *rambleChatServer) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Ramble_Ping_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RambleServer).Ping(&ramblePingServer{stream})
}

type Ramble_PingServer interface {
	SendAndClose(*PingReply) error
	Recv() (*PingRequest, error)
	grpc.ServerStream
}

type ramblePingServer struct {
	grpc.ServerStream
}

func (x *ramblePingServer) SendAndClose(m *PingReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ramblePingServer) Recv() (*PingRequest, error) {
	m := new(PingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Ramble_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Ramble",
	HandlerType: (*RambleServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chat",
			Handler:       _Ramble_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Ping",
			Handler:       _Ramble_Ping_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x41, 0x6b, 0xbc, 0x30,
	0x10, 0xc5, 0xff, 0x59, 0xc5, 0xff, 0x3a, 0xb6, 0x6c, 0x99, 0x43, 0x09, 0xd2, 0x83, 0xd8, 0x8b,
	0xf4, 0x20, 0x65, 0xfb, 0x11, 0x7a, 0x2e, 0x94, 0x7c, 0x83, 0xa8, 0xc3, 0x2a, 0x24, 0x9a, 0x9a,
	0xb4, 0xe0, 0xb7, 0x2f, 0xc9, 0x76, 0xdd, 0x65, 0x0f, 0x3d, 0xf4, 0x96, 0xf7, 0xde, 0xf0, 0x7e,
	0x64, 0x06, 0xa0, 0xed, 0xa5, 0xab, 0xcd, 0x3c, 0xb9, 0x09, 0x37, 0xa6, 0x29, 0x17, 0xc8, 0x5e,
	0x7b, 0xe9, 0xde, 0xc8, 0x5a, 0x79, 0x20, 0xcc, 0x61, 0x6b, 0xe9, 0xe3, 0x93, 0xc6, 0x96, 0x38,
	0x2b, 0x58, 0x15, 0x89, 0x55, 0xe3, 0x3d, 0x24, 0x96, 0xc6, 0x8e, 0x66, 0xbe, 0x29, 0x58, 0x95,
	0x8a, 0x1f, 0x85, 0x0f, 0x90, 0xba, 0x41, 0x93, 0x75, 0x52, 0x1b, 0x1e, 0x85, 0xe8, 0x6c, 0x20,
	0x87, 0xff, 0xfa, 0x58, 0xce, 0xe3, 0x90, 0x9d, 0x64, 0xa9, 0x21, 0x7b, 0x1f, 0xc6, 0x83, 0xf0,
	0xfd, 0xd6, 0xfd, 0x09, 0x7d, 0x07, 0x91, 0x73, 0x2a, 0x40, 0x23, 0xe1, 0x9f, 0x1e, 0x67, 0xe4,
	0xa2, 0x26, 0xd9, 0xf1, 0x5d, 0xc1, 0xaa, 0x1b, 0x71, 0x92, 0x65, 0x0f, 0xe9, 0x11, 0x67, 0xd4,
	0xf2, 0x2b, 0x2c, 0x87, 0xed, 0x4c, 0x2d, 0x0d, 0x5f, 0x2b, 0x6e, 0xd5, 0xf8, 0x08, 0x31, 0xb5,
	0xfd, 0x14, 0x88, 0xd9, 0x7e, 0x57, 0x9b, 0xa6, 0xbe, 0xf8, 0x83, 0x08, 0xe1, 0xbe, 0x83, 0x44,
	0x48, 0xdd, 0x28, 0xc2, 0x1a, 0x62, 0xbf, 0x5d, 0x0c, 0x83, 0x17, 0x7b, 0xce, 0xaf, 0x8d, 0xf2,
	0x5f, 0xc5, 0x9e, 0x19, 0x3e, 0x41, 0xec, 0xeb, 0xf0, 0xba, 0x38, 0xbf, 0x3d, 0x1b, 0x46, 0x2d,
	0x7e, 0xba, 0x49, 0xc2, 0x11, 0x5f, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xf3, 0x67, 0x11,
	0xd2, 0x01, 0x00, 0x00,
}