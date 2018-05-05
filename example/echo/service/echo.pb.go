// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/api-gateway/example/echo/service/echo.proto

/*
Package echo is a generated protocol buffer package.

It is generated from these files:
	github.com/api-gateway/example/echo/service/echo.proto

It has these top-level messages:
	EchoRequest
	EchoResponse
*/
package echo

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import http "net/http"
import strings "strings"
import math "math"
import google_protobuf1 "github.com/gogo/protobuf/types"
import google_protobuf2 "github.com/gogo/protobuf/types"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type EchoRequest struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (m *EchoRequest) Reset()                    { *m = EchoRequest{} }
func (m *EchoRequest) String() string            { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()               {}
func (*EchoRequest) Descriptor() ([]byte, []int) { return fileDescriptorEcho, []int{0} }

func (m *EchoRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type EchoResponse struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (m *EchoResponse) Reset()                    { *m = EchoResponse{} }
func (m *EchoResponse) String() string            { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()               {}
func (*EchoResponse) Descriptor() ([]byte, []int) { return fileDescriptorEcho, []int{1} }

func (m *EchoResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "echo.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "echo.EchoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Echo service

type EchoClient interface {
	Ping(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*google_protobuf2.Timestamp, error)
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Ping(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*google_protobuf2.Timestamp, error) {
	out := new(google_protobuf2.Timestamp)
	err := grpc.Invoke(ctx, "/echo.Echo/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := grpc.Invoke(ctx, "/echo.Echo/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Echo service

type EchoServer interface {
	Ping(context.Context, *google_protobuf1.Empty) (*google_protobuf2.Timestamp, error)
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echo/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Ping(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Echo_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echo/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Echo_Ping_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _Echo_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/api-gateway/example/echo/service/echo.proto",
}

func init() {
	proto.RegisterFile("github.com/api-gateway/example/echo/service/echo.proto", fileDescriptorEcho)
}

var fileDescriptorEcho = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x59, 0x59, 0x44, 0xa3, 0x82, 0x46, 0x10, 0x59, 0x05, 0x35, 0x27, 0x11, 0x4c, 0x50,
	0xc1, 0x83, 0x47, 0xa1, 0x37, 0x0f, 0x52, 0x7c, 0x81, 0xec, 0x32, 0xee, 0x06, 0x9a, 0x4c, 0x6c,
	0x66, 0x6b, 0x7b, 0xf5, 0x15, 0xbc, 0xf8, 0x5e, 0xbe, 0x82, 0x0f, 0x22, 0x49, 0xb6, 0x50, 0xb4,
	0xb7, 0x99, 0xf9, 0xbf, 0xfc, 0xff, 0x4f, 0xd8, 0x7d, 0x6b, 0xa8, 0xeb, 0x6b, 0xd9, 0xa0, 0x55,
	0xda, 0x9b, 0xeb, 0x56, 0x13, 0xbc, 0xeb, 0x85, 0x82, 0xb9, 0xb6, 0x7e, 0x02, 0x0a, 0x9a, 0x0e,
	0x55, 0x80, 0xe9, 0xcc, 0x34, 0x79, 0x91, 0x7e, 0x8a, 0x84, 0xbc, 0x8c, 0x73, 0x75, 0xda, 0x22,
	0xb6, 0x13, 0x88, 0x2f, 0x95, 0x76, 0x0e, 0x49, 0x93, 0x41, 0x17, 0x32, 0x53, 0x9d, 0x0c, 0x6a,
	0xda, 0xea, 0xfe, 0x55, 0x81, 0xf5, 0xb4, 0x18, 0xc4, 0xb3, 0xbf, 0x22, 0x19, 0x0b, 0x81, 0xb4,
	0xf5, 0x19, 0x10, 0x17, 0x6c, 0x67, 0xd4, 0x74, 0x38, 0x86, 0xb7, 0x1e, 0x02, 0x71, 0xce, 0x4a,
	0x82, 0x39, 0x1d, 0x17, 0xe7, 0xc5, 0xe5, 0xf6, 0x38, 0xcd, 0x42, 0xb0, 0xdd, 0x8c, 0x04, 0x8f,
	0x2e, 0xc0, 0x3a, 0xe6, 0xf6, 0xab, 0x60, 0x65, 0x84, 0xf8, 0x13, 0x2b, 0x9f, 0x8d, 0x6b, 0xf9,
	0x91, 0xcc, 0xc9, 0x72, 0x99, 0x2c, 0x47, 0xb1, 0x56, 0x55, 0xfd, 0xbb, 0xbf, 0x2c, 0x1b, 0x89,
	0xfd, 0x8f, 0xef, 0x9f, 0xcf, 0x0d, 0xc6, 0xb7, 0xd4, 0xec, 0x46, 0xf9, 0xe8, 0xf2, 0x38, 0xb8,
	0x1e, 0xc8, 0xf4, 0x29, 0x2b, 0x4d, 0x2b, 0xbe, 0x7a, 0xca, 0xcd, 0xc4, 0x61, 0x32, 0xd8, 0x13,
	0xc9, 0x20, 0xca, 0x0f, 0xc5, 0x55, 0xbd, 0x99, 0x92, 0xee, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x41, 0x1f, 0xb1, 0x4f, 0x84, 0x01, 0x00, 0x00,
}

const PROTO_JSON ="[{\"Package\":\"echo\",\"Service\":\"Echo\",\"Method\":{\"name\":\"Ping\",\"input_type\":\".google.protobuf.Empty\",\"output_type\":\".google.protobuf.Timestamp\",\"options\":{}},\"InputTypeDescriptor\":null,\"Pattern\":{\"Verb\":\"GET\",\"Path\":\"/v1/ping\",\"Body\":\"\"},\"Options\":{}},{\"Package\":\"echo\",\"Service\":\"Echo\",\"Method\":{\"name\":\"Echo\",\"input_type\":\".echo.EchoRequest\",\"output_type\":\".echo.EchoResponse\",\"options\":{}},\"InputTypeDescriptor\":{\"name\":\"EchoRequest\",\"field\":[{\"name\":\"text\",\"number\":1,\"label\":1,\"type\":9,\"json_name\":\"text\"}]},\"Pattern\":{\"Verb\":\"POST\",\"Path\":\"/v1/echo\",\"Body\":\"*\"},\"Options\":{}}]"

func init() {
	if _, err := (&http.Client{}).Post("http://api-gateway:8080/rules", "", strings.NewReader(PROTO_JSON)); err != nil {
		panic(err)
	}
}