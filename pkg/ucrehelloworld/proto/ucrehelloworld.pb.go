// Code generated by protoc-gen-go.
// source: ucrehelloworld.proto
// DO NOT EDIT!

/*
Package ucrehelloworld is a generated protocol buffer package.

It is generated from these files:
	ucrehelloworld.proto

It has these top-level messages:
	SayHelloRequest
	SayHelloResponse
*/
package ucrehelloworld

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "gitlab.ucloudadmin.com/ucre/rpc/pkg/common/proto"

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
const _ = proto.ProtoPackageIsVersion1

type SayHelloRequest struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SayHelloRequest) Reset()                    { *m = SayHelloRequest{} }
func (m *SayHelloRequest) String() string            { return proto.CompactTextString(m) }
func (*SayHelloRequest) ProtoMessage()               {}
func (*SayHelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SayHelloRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type SayHelloResponse struct {
	Ret              *common.Ret `protobuf:"bytes,1,opt,name=ret" json:"ret,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SayHelloResponse) Reset()                    { *m = SayHelloResponse{} }
func (m *SayHelloResponse) String() string            { return proto.CompactTextString(m) }
func (*SayHelloResponse) ProtoMessage()               {}
func (*SayHelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SayHelloResponse) GetRet() *common.Ret {
	if m != nil {
		return m.Ret
	}
	return nil
}

func init() {
	proto.RegisterType((*SayHelloRequest)(nil), "ucrehelloworld.SayHelloRequest")
	proto.RegisterType((*SayHelloResponse)(nil), "ucrehelloworld.SayHelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for UcreHelloWorld service

type UcreHelloWorldClient interface {
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
}

type ucreHelloWorldClient struct {
	cc *grpc.ClientConn
}

func NewUcreHelloWorldClient(cc *grpc.ClientConn) UcreHelloWorldClient {
	return &ucreHelloWorldClient{cc}
}

func (c *ucreHelloWorldClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	out := new(SayHelloResponse)
	err := grpc.Invoke(ctx, "/ucrehelloworld.UcreHelloWorld/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UcreHelloWorld service

type UcreHelloWorldServer interface {
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
}

func RegisterUcreHelloWorldServer(s *grpc.Server, srv UcreHelloWorldServer) {
	s.RegisterService(&_UcreHelloWorld_serviceDesc, srv)
}

func _UcreHelloWorld_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UcreHelloWorldServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ucrehelloworld.UcreHelloWorld/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UcreHelloWorldServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UcreHelloWorld_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ucrehelloworld.UcreHelloWorld",
	HandlerType: (*UcreHelloWorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _UcreHelloWorld_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x4d, 0x2e, 0x4a,
	0xcd, 0x48, 0xcd, 0xc9, 0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x43, 0x15, 0x95, 0xe2, 0x49, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0xc8, 0x2a, 0xa9,
	0x72, 0xf1, 0x07, 0x27, 0x56, 0x7a, 0x80, 0xa4, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84,
	0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c,
	0x25, 0x43, 0x2e, 0x01, 0x84, 0xb2, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x59, 0x2e, 0xe6,
	0xa2, 0xd4, 0x12, 0xb0, 0x32, 0x6e, 0x23, 0x6e, 0x3d, 0xa8, 0xb1, 0x41, 0xa9, 0x25, 0x41, 0x20,
	0x71, 0xa3, 0x44, 0x2e, 0xbe, 0x50, 0xa0, 0xcd, 0x60, 0x3d, 0xe1, 0x20, 0x9b, 0x85, 0xfc, 0xb9,
	0x38, 0x60, 0x86, 0x08, 0xc9, 0xeb, 0xa1, 0x39, 0x16, 0xcd, 0x15, 0x52, 0x0a, 0xb8, 0x15, 0x40,
	0xec, 0x57, 0x62, 0x00, 0x04, 0x00, 0x00, 0xff, 0xff, 0xee, 0x28, 0x14, 0x65, 0xf1, 0x00, 0x00,
	0x00,
}