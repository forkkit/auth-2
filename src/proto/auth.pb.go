// Code generated by protoc-gen-go.
// source: auth.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	auth.proto

It has these top-level messages:
	Auth
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto1.ProtoPackageIsVersion1

type Auth_CertificateType int32

const (
	Auth_UUID     Auth_CertificateType = 0
	Auth_PLAIN    Auth_CertificateType = 1
	Auth_TOKEN    Auth_CertificateType = 2
	Auth_FACEBOOK Auth_CertificateType = 3
)

var Auth_CertificateType_name = map[int32]string{
	0: "UUID",
	1: "PLAIN",
	2: "TOKEN",
	3: "FACEBOOK",
}
var Auth_CertificateType_value = map[string]int32{
	"UUID":     0,
	"PLAIN":    1,
	"TOKEN":    2,
	"FACEBOOK": 3,
}

func (x Auth_CertificateType) String() string {
	return proto1.EnumName(Auth_CertificateType_name, int32(x))
}
func (Auth_CertificateType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Auth struct {
}

func (m *Auth) Reset()                    { *m = Auth{} }
func (m *Auth) String() string            { return proto1.CompactTextString(m) }
func (*Auth) ProtoMessage()               {}
func (*Auth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Auth_Certificate struct {
	Type  Auth_CertificateType `protobuf:"varint,1,opt,name=Type,enum=proto.Auth_CertificateType" json:"Type,omitempty"`
	Proof []byte               `protobuf:"bytes,2,opt,name=Proof,proto3" json:"Proof,omitempty"`
}

func (m *Auth_Certificate) Reset()                    { *m = Auth_Certificate{} }
func (m *Auth_Certificate) String() string            { return proto1.CompactTextString(m) }
func (*Auth_Certificate) ProtoMessage()               {}
func (*Auth_Certificate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Auth_Result struct {
	OK     bool   `protobuf:"varint,1,opt,name=OK" json:"OK,omitempty"`
	UserId int32  `protobuf:"varint,2,opt,name=UserId" json:"UserId,omitempty"`
	Body   []byte `protobuf:"bytes,3,opt,name=Body,proto3" json:"Body,omitempty"`
}

func (m *Auth_Result) Reset()                    { *m = Auth_Result{} }
func (m *Auth_Result) String() string            { return proto1.CompactTextString(m) }
func (*Auth_Result) ProtoMessage()               {}
func (*Auth_Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

func init() {
	proto1.RegisterType((*Auth)(nil), "proto.Auth")
	proto1.RegisterType((*Auth_Certificate)(nil), "proto.Auth.Certificate")
	proto1.RegisterType((*Auth_Result)(nil), "proto.Auth.Result")
	proto1.RegisterEnum("proto.Auth_CertificateType", Auth_CertificateType_name, Auth_CertificateType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for AuthService service

type AuthServiceClient interface {
	Auth(ctx context.Context, in *Auth_Certificate, opts ...grpc.CallOption) (*Auth_Result, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Auth(ctx context.Context, in *Auth_Certificate, opts ...grpc.CallOption) (*Auth_Result, error) {
	out := new(Auth_Result)
	err := grpc.Invoke(ctx, "/proto.AuthService/Auth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthService service

type AuthServiceServer interface {
	Auth(context.Context, *Auth_Certificate) (*Auth_Result, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auth_Certificate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Auth(ctx, req.(*Auth_Certificate))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _AuthService_Auth_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x47, 0x18, 0xb9, 0x58, 0x1c,
	0x81, 0xa2, 0x52, 0xee, 0x5c, 0xdc, 0xce, 0xa9, 0x45, 0x25, 0x99, 0x69, 0x99, 0xc9, 0x89, 0x25,
	0xa9, 0x42, 0x9a, 0x5c, 0x2c, 0x21, 0x95, 0x05, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x7c, 0x46,
	0xd2, 0x10, 0x4d, 0x7a, 0x20, 0x95, 0x7a, 0x48, 0xca, 0x40, 0x4a, 0x84, 0x78, 0xb9, 0x58, 0x03,
	0x8a, 0xf2, 0xf3, 0xd3, 0x24, 0x98, 0x80, 0x6a, 0x79, 0xa4, 0x8c, 0xb8, 0xd8, 0x82, 0x52, 0x8b,
	0x4b, 0x73, 0x4a, 0x84, 0xb8, 0xb8, 0x98, 0xfc, 0xbd, 0xc1, 0x26, 0x70, 0x08, 0xf1, 0x71, 0xb1,
	0x85, 0x16, 0xa7, 0x16, 0x79, 0xa6, 0x80, 0x55, 0xb1, 0x0a, 0xf1, 0x70, 0xb1, 0x38, 0xe5, 0xa7,
	0x54, 0x4a, 0x30, 0x83, 0xf4, 0x28, 0xd9, 0x73, 0xf1, 0xa3, 0x9b, 0xca, 0xc1, 0xc5, 0x12, 0x1a,
	0xea, 0xe9, 0x22, 0xc0, 0x20, 0xc4, 0x09, 0x34, 0xdf, 0xc7, 0xd1, 0xd3, 0x4f, 0x80, 0x11, 0xc4,
	0x0c, 0xf1, 0xf7, 0x76, 0xf5, 0x13, 0x60, 0x02, 0x1a, 0xc0, 0xe1, 0xe6, 0xe8, 0xec, 0xea, 0xe4,
	0xef, 0xef, 0x2d, 0xc0, 0x6c, 0xe4, 0xc2, 0xc5, 0x0d, 0x72, 0x5b, 0x70, 0x6a, 0x51, 0x59, 0x66,
	0x72, 0xaa, 0x90, 0x29, 0xc4, 0x53, 0x42, 0xe2, 0x38, 0xdc, 0x2d, 0x25, 0x84, 0x2c, 0x01, 0x71,
	0xae, 0x12, 0x43, 0x12, 0x1b, 0x58, 0xd0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x22, 0x0b,
	0x19, 0x28, 0x01, 0x00, 0x00,
}
