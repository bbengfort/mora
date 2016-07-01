// Code generated by protoc-gen-go.
// source: echo.proto
// DO NOT EDIT!

/*
Package echo is a generated protocol buffer package.

It is generated from these files:
	echo.proto

It has these top-level messages:
	Time
	Node
	EchoRequest
	EchoReply
*/
package echo

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
// const _ = proto.ProtoPackageIsVersion1

// Time preserves nanosecond latency measurements by using this custom time
// struct which should include either seconds or nanoseconds since the Unix
// epoch as unsigned int64. In Go, you can use time.Unix to parse this field.
type Time struct {
	Seconds     int64 `protobuf:"varint,1,opt,name=seconds" json:"seconds,omitempty"`
	Nanoseconds int64 `protobuf:"varint,2,opt,name=nanoseconds" json:"nanoseconds,omitempty"`
}

// Reset the message
func (m *Time) Reset() { *m = Time{} }

// String returns a string representation of the message
func (m *Time) String() string { return proto.CompactTextString(m) }

// ProtoMessage is a generated method
func (*Time) ProtoMessage() {}

// Descriptor is a generated method
func (*Time) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Node describes an echoing machine on the network, either the sender
// (the source) or the receiver (the target). This is distinct from mora.Node.
type Node struct {
	Id      int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	Dns     string `protobuf:"bytes,4,opt,name=dns" json:"dns,omitempty"`
}

// Reset the message
func (m *Node) Reset() { *m = Node{} }

// String returns a string representation of the message
func (m *Node) String() string { return proto.CompactTextString(m) }

// ProtoMessage is a generated method
func (*Node) ProtoMessage() {}

// Descriptor is a generated method
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// EchoRequest is used to measure latency and uptime as with a ping, but
// designed to be more application-layer specific for measuring system latency.
type EchoRequest struct {
	Source  *Node  `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Target  *Node  `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
	Sent    *Time  `protobuf:"bytes,3,opt,name=sent" json:"sent,omitempty"`
	Payload []byte `protobuf:"bytes,15,opt,name=payload,proto3" json:"payload,omitempty"`
}

// Reset the message
func (m *EchoRequest) Reset() { *m = EchoRequest{} }

// String returns a string representation of the message
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }

// ProtoMessage is a generated method
func (*EchoRequest) ProtoMessage() {}

// Descriptor is a generated method
func (*EchoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// GetSource returns the source node
func (m *EchoRequest) GetSource() *Node {
	if m != nil {
		return m.Source
	}
	return nil
}

// GetTarget returns the target node
func (m *EchoRequest) GetTarget() *Node {
	if m != nil {
		return m.Target
	}
	return nil
}

// GetSent returns the time struct
func (m *EchoRequest) GetSent() *Time {
	if m != nil {
		return m.Sent
	}
	return nil
}

// EchoReply is used to echo a message containing the actual receiver node as
// well as the received timestamp and a payload containing the original echo.
type EchoReply struct {
	Receiver *Node        `protobuf:"bytes,1,opt,name=receiver" json:"receiver,omitempty"`
	Received *Time        `protobuf:"bytes,2,opt,name=received" json:"received,omitempty"`
	Echo     *EchoRequest `protobuf:"bytes,3,opt,name=echo" json:"echo,omitempty"`
}

// Reset the message
func (m *EchoReply) Reset() { *m = EchoReply{} }

// String returns a string representation of the message
func (m *EchoReply) String() string { return proto.CompactTextString(m) }

// ProtoMessage is a generated method
func (*EchoReply) ProtoMessage() {}

// Descriptor is a generated method
func (*EchoReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// GetReceiver returns the actual receiving node
func (m *EchoReply) GetReceiver() *Node {
	if m != nil {
		return m.Receiver
	}
	return nil
}

// GetReceived returns the timestamp of reciept
func (m *EchoReply) GetReceived() *Time {
	if m != nil {
		return m.Received
	}
	return nil
}

// GetEcho returns the original echo payload
func (m *EchoReply) GetEcho() *EchoRequest {
	if m != nil {
		return m.Echo
	}
	return nil
}

func init() {
	proto.RegisterType((*Time)(nil), "echo.Time")
	proto.RegisterType((*Node)(nil), "echo.Node")
	proto.RegisterType((*EchoRequest)(nil), "echo.EchoRequest")
	proto.RegisterType((*EchoReply)(nil), "echo.EchoReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Echo service

// EchoClient is an interface for sending bounce messages.
type EchoClient interface {
	// Bounce allows nodes to respond to echo requests with echo replies.
	Bounce(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

// NewEchoClient creates an echo client with a connection
func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

// Bounce implements the client interface
func (c *echoClient) Bounce(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error) {
	out := new(EchoReply)
	err := grpc.Invoke(ctx, "/echo.Echo/Bounce", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Echo service

// EchoServer is an interface for echo responders
type EchoServer interface {
	// Bounce allows nodes to respond to echo requests with echo replies.
	Bounce(context.Context, *EchoRequest) (*EchoReply, error)
}

// RegisterEchoServer registers a service description on the context
func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Bounce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Bounce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echo/Bounce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Bounce(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bounce",
			Handler:    _Echo_Bounce_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0x4d, 0x4b, 0x03, 0x31,
	0x10, 0x75, 0xdb, 0x50, 0xed, 0x54, 0xac, 0x9d, 0xd3, 0xe2, 0x41, 0x4a, 0x40, 0xf1, 0xb4, 0x87,
	0x0a, 0xfe, 0x80, 0x05, 0xaf, 0x52, 0x82, 0x78, 0x4f, 0x37, 0x63, 0x2d, 0xb4, 0x49, 0x4d, 0x52,
	0xa1, 0x57, 0xff, 0x80, 0x7f, 0xd9, 0xcd, 0xb4, 0xbb, 0xac, 0x1f, 0xb7, 0x37, 0xef, 0xbd, 0xf0,
	0xde, 0x24, 0x01, 0xa0, 0xea, 0xcd, 0x15, 0x5b, 0xef, 0xa2, 0x43, 0x91, 0xb0, 0x2c, 0x41, 0x3c,
	0xaf, 0x36, 0x84, 0x39, 0x9c, 0x06, 0xaa, 0x9c, 0x35, 0x21, 0xcf, 0xa6, 0xd9, 0x5d, 0x5f, 0x35,
	0x23, 0x4e, 0x61, 0x64, 0xb5, 0x75, 0x8d, 0xda, 0x63, 0xb5, 0x4b, 0xc9, 0x17, 0x10, 0x4f, 0xce,
	0x10, 0x5e, 0x40, 0x6f, 0x65, 0x8e, 0xc7, 0x6b, 0x84, 0x08, 0xc2, 0xea, 0x0d, 0xf1, 0x91, 0xa1,
	0x62, 0x9c, 0x72, 0xb4, 0x31, 0x9e, 0x42, 0xc8, 0xfb, 0x4c, 0x37, 0x23, 0x5e, 0x42, 0xdf, 0xd8,
	0x90, 0x0b, 0x66, 0x13, 0x94, 0x5f, 0x19, 0x8c, 0x1e, 0xeb, 0x92, 0x8a, 0xde, 0x77, 0x14, 0x22,
	0x4a, 0x18, 0x04, 0xb7, 0xf3, 0x15, 0x71, 0xc6, 0x68, 0x06, 0x05, 0xaf, 0x93, 0xb2, 0xd5, 0x51,
	0x49, 0x9e, 0xa8, 0xfd, 0x92, 0x22, 0xa7, 0xfe, 0xf2, 0x1c, 0x14, 0xbc, 0x06, 0x11, 0xc8, 0x46,
	0x2e, 0xd0, 0x3a, 0xd2, 0x2d, 0x28, 0xe6, 0x53, 0xc7, 0xad, 0xde, 0xaf, 0x9d, 0x36, 0xf9, 0xb8,
	0xb6, 0x9c, 0xab, 0x66, 0x94, 0x9f, 0x19, 0x0c, 0x0f, 0x8d, 0xb6, 0xeb, 0x3d, 0xde, 0xc2, 0x99,
	0xa7, 0x8a, 0x56, 0x1f, 0xe4, 0xff, 0x69, 0xd4, 0x6a, 0x1d, 0x9f, 0xf9, 0xd9, 0x8a, 0x33, 0x5b,
	0x0d, 0x6f, 0x80, 0xdf, 0xe4, 0xd8, 0x6b, 0x72, 0xf0, 0x74, 0x2e, 0x40, 0xb1, 0x3c, 0x7b, 0x00,
	0x91, 0x48, 0x2c, 0x60, 0x50, 0xba, 0x9d, 0xad, 0x97, 0xfe, 0x6b, 0xbd, 0x1a, 0x77, 0xa9, 0xba,
	0xac, 0x3c, 0x29, 0xa7, 0x30, 0xa9, 0xdc, 0xa6, 0x58, 0x90, 0x5d, 0xbe, 0x3a, 0x1f, 0x59, 0x2b,
	0x79, 0x9d, 0x79, 0xfa, 0x10, 0xf3, 0x6c, 0x31, 0xe0, 0x9f, 0x71, 0xff, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0x55, 0xdf, 0xdd, 0x1b, 0x27, 0x02, 0x00, 0x00,
}
