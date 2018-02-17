// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package grpchantesting is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	Message
*/
package grpchantesting

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"

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

type Message struct {
	Payload              []byte            `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Count                int32             `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
	Code                 int32             `protobuf:"varint,3,opt,name=code" json:"code,omitempty"`
	DelayMillis          int32             `protobuf:"varint,4,opt,name=delay_millis,json=delayMillis" json:"delay_millis,omitempty"`
	Headers              map[string]string `protobuf:"bytes,5,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Trailers             map[string]string `protobuf:"bytes,6,rep,name=trailers" json:"trailers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `protobuf_unrecognized:"proto3" json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }
func (m *Message) Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Message) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Message) GetDelayMillis() int32 {
	if m != nil {
		return m.DelayMillis
	}
	return 0
}

func (m *Message) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Message) GetTrailers() map[string]string {
	if m != nil {
		return m.Trailers
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "grpchantesting.Message")
	proto.RegisterMapType((map[string]string)(nil), "grpchantesting.Message.HeadersEntry")
	proto.RegisterMapType((map[string]string)(nil), "grpchantesting.Message.TrailersEntry")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TestService service

type TestServiceClient interface {
	Unary(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (TestService_ClientStreamClient, error)
	ServerStream(ctx context.Context, in *Message, opts ...grpc.CallOption) (TestService_ServerStreamClient, error)
	BidiStream(ctx context.Context, opts ...grpc.CallOption) (TestService_BidiStreamClient, error)
	// UseExternalMessageTwice is here purely to test the protoc-gen-gochanstubs plug-in
	UseExternalMessageTwice(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type testServiceClient struct {
	cc *grpc.ClientConn
}

func NewTestServiceClient(cc *grpc.ClientConn) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) Unary(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := grpc.Invoke(ctx, "/grpchantesting.TestService/Unary", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (TestService_ClientStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[0], c.cc, "/grpchantesting.TestService/ClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceClientStreamClient{stream}
	return x, nil
}

type TestService_ClientStreamClient interface {
	Send(*Message) error
	CloseAndRecv() (*Message, error)
	grpc.ClientStream
}

type testServiceClientStreamClient struct {
	grpc.ClientStream
}

func (x *testServiceClientStreamClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceClientStreamClient) CloseAndRecv() (*Message, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) ServerStream(ctx context.Context, in *Message, opts ...grpc.CallOption) (TestService_ServerStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[1], c.cc, "/grpchantesting.TestService/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_ServerStreamClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type testServiceServerStreamClient struct {
	grpc.ClientStream
}

func (x *testServiceServerStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) BidiStream(ctx context.Context, opts ...grpc.CallOption) (TestService_BidiStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TestService_serviceDesc.Streams[2], c.cc, "/grpchantesting.TestService/BidiStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceBidiStreamClient{stream}
	return x, nil
}

type TestService_BidiStreamClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type testServiceBidiStreamClient struct {
	grpc.ClientStream
}

func (x *testServiceBidiStreamClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceBidiStreamClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) UseExternalMessageTwice(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/grpchantesting.TestService/UseExternalMessageTwice", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TestService service

type TestServiceServer interface {
	Unary(context.Context, *Message) (*Message, error)
	ClientStream(TestService_ClientStreamServer) error
	ServerStream(*Message, TestService_ServerStreamServer) error
	BidiStream(TestService_BidiStreamServer) error
	// UseExternalMessageTwice is here purely to test the protoc-gen-gochanstubs plug-in
	UseExternalMessageTwice(context.Context, *google_protobuf.Empty) (*google_protobuf.Empty, error)
}

func RegisterTestServiceServer(s *grpc.Server, srv TestServiceServer) {
	s.RegisterService(&_TestService_serviceDesc, srv)
}

func _TestService_Unary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).Unary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpchantesting.TestService/Unary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).Unary(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_ClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).ClientStream(&testServiceClientStreamServer{stream})
}

type TestService_ClientStreamServer interface {
	SendAndClose(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type testServiceClientStreamServer struct {
	grpc.ServerStream
}

func (x *testServiceClientStreamServer) SendAndClose(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceClientStreamServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Message)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).ServerStream(m, &testServiceServerStreamServer{stream})
}

type TestService_ServerStreamServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type testServiceServerStreamServer struct {
	grpc.ServerStream
}

func (x *testServiceServerStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_BidiStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).BidiStream(&testServiceBidiStreamServer{stream})
}

type TestService_BidiStreamServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type testServiceBidiStreamServer struct {
	grpc.ServerStream
}

func (x *testServiceBidiStreamServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceBidiStreamServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_UseExternalMessageTwice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).UseExternalMessageTwice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpchantesting.TestService/UseExternalMessageTwice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).UseExternalMessageTwice(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpchantesting.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unary",
			Handler:    _TestService_Unary_Handler,
		},
		{
			MethodName: "UseExternalMessageTwice",
			Handler:    _TestService_UseExternalMessageTwice_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientStream",
			Handler:       _TestService_ClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ServerStream",
			Handler:       _TestService_ServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BidiStream",
			Handler:       _TestService_BidiStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "test.proto",
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x4f, 0xea, 0x40,
	0x10, 0xc7, 0xd3, 0x42, 0xe1, 0x31, 0xf4, 0xbd, 0xbc, 0x6c, 0x8c, 0x34, 0xf5, 0x82, 0x46, 0x93,
	0x9e, 0x0a, 0xc1, 0x8b, 0x62, 0x62, 0x22, 0x86, 0x44, 0x0f, 0x5c, 0x0a, 0x9c, 0xcd, 0xd2, 0x8e,
	0x65, 0xe3, 0xd2, 0x36, 0xdb, 0x05, 0xed, 0xa7, 0xf5, 0x0b, 0xf8, 0x21, 0xcc, 0x6e, 0xc1, 0xc8,
	0xa1, 0x07, 0xb8, 0xed, 0xfc, 0x67, 0xe6, 0x97, 0xff, 0x7f, 0x33, 0x00, 0x12, 0x73, 0xe9, 0x67,
	0x22, 0x95, 0x29, 0xf9, 0x17, 0x8b, 0x2c, 0x5c, 0xd2, 0x44, 0x49, 0x2c, 0x89, 0xdd, 0xb3, 0x38,
	0x4d, 0x63, 0x8e, 0x3d, 0xdd, 0x5d, 0xac, 0x5f, 0x7b, 0xb8, 0xca, 0x64, 0x51, 0x0e, 0x5f, 0x7c,
	0x99, 0xd0, 0x9c, 0x60, 0x9e, 0xd3, 0x18, 0x89, 0x03, 0xcd, 0x8c, 0x16, 0x3c, 0xa5, 0x91, 0x63,
	0x74, 0x0d, 0xcf, 0x0e, 0x76, 0x25, 0x39, 0x01, 0x2b, 0x4c, 0xd7, 0x89, 0x74, 0xcc, 0xae, 0xe1,
	0x59, 0x41, 0x59, 0x10, 0x02, 0xf5, 0x30, 0x8d, 0xd0, 0xa9, 0x69, 0x51, 0xbf, 0xc9, 0x39, 0xd8,
	0x11, 0x72, 0x5a, 0xbc, 0xac, 0x18, 0xe7, 0x2c, 0x77, 0xea, 0xba, 0xd7, 0xd6, 0xda, 0x44, 0x4b,
	0xe4, 0x1e, 0x9a, 0x4b, 0xa4, 0x11, 0x8a, 0xdc, 0xb1, 0xba, 0x35, 0xaf, 0x3d, 0xb8, 0xf4, 0xf7,
	0x1d, 0xfb, 0x5b, 0x43, 0xfe, 0x53, 0x39, 0x36, 0x4e, 0xa4, 0x28, 0x82, 0xdd, 0x12, 0x79, 0x80,
	0x3f, 0x52, 0x50, 0xc6, 0x15, 0xa0, 0xa1, 0x01, 0x57, 0x55, 0x80, 0xd9, 0x76, 0xae, 0x24, 0xfc,
	0xac, 0xb9, 0x43, 0xb0, 0x7f, 0xb3, 0xc9, 0x7f, 0xa8, 0xbd, 0x61, 0xa1, 0x53, 0xb7, 0x02, 0xf5,
	0x54, 0x89, 0x37, 0x94, 0xaf, 0x51, 0x27, 0x6e, 0x05, 0x65, 0x31, 0x34, 0x6f, 0x0c, 0xf7, 0x0e,
	0xfe, 0xee, 0x61, 0x0f, 0x59, 0x1e, 0x7c, 0x9a, 0xd0, 0x9e, 0x61, 0x2e, 0xa7, 0x28, 0x36, 0x2c,
	0x44, 0x72, 0x0b, 0xd6, 0x3c, 0xa1, 0xa2, 0x20, 0x9d, 0x8a, 0x08, 0x6e, 0x55, 0x83, 0x8c, 0xc0,
	0x7e, 0xe4, 0x0c, 0x13, 0x39, 0x95, 0x02, 0xe9, 0xea, 0x70, 0x82, 0x67, 0x28, 0x86, 0x72, 0x82,
	0xe2, 0x58, 0x46, 0x5f, 0x31, 0x60, 0xc4, 0x22, 0x76, 0xbc, 0x8b, 0xbe, 0x41, 0x9e, 0xa1, 0x33,
	0xcf, 0x71, 0xfc, 0x21, 0x51, 0x24, 0x94, 0x6f, 0x3b, 0xb3, 0x77, 0xf5, 0x43, 0xa7, 0x7e, 0x79,
	0xbe, 0xfe, 0xee, 0x7c, 0xfd, 0xb1, 0x3a, 0x5f, 0xb7, 0x42, 0x5f, 0x34, 0x74, 0x7d, 0xfd, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0xb7, 0x9d, 0xe4, 0x24, 0x12, 0x03, 0x00, 0x00,
}
