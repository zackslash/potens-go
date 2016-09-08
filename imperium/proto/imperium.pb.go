// Code generated by protoc-gen-go.
// source: proto/imperium.proto
// DO NOT EDIT!

/*
Package imperium is a generated protocol buffer package.

It is generated from these files:
	proto/imperium.proto

It has these top-level messages:
	CertificateRequest
	CertificateResponse
*/
package imperium

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type CertificateRequest struct {
	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
}

func (m *CertificateRequest) Reset()                    { *m = CertificateRequest{} }
func (m *CertificateRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateRequest) ProtoMessage()               {}
func (*CertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CertificateResponse struct {
	Certificate     string                     `protobuf:"bytes,1,opt,name=certificate" json:"certificate,omitempty"`
	PrivateKey      string                     `protobuf:"bytes,2,opt,name=private_key,json=privateKey" json:"private_key,omitempty"`
	RootCertificate string                     `protobuf:"bytes,3,opt,name=root_certificate,json=rootCertificate" json:"root_certificate,omitempty"`
	Hostname        string                     `protobuf:"bytes,4,opt,name=hostname" json:"hostname,omitempty"`
	ExpiryTime      *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=expiry_time,json=expiryTime" json:"expiry_time,omitempty"`
}

func (m *CertificateResponse) Reset()                    { *m = CertificateResponse{} }
func (m *CertificateResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateResponse) ProtoMessage()               {}
func (*CertificateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CertificateResponse) GetExpiryTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpiryTime
	}
	return nil
}

func init() {
	proto.RegisterType((*CertificateRequest)(nil), "imperium.CertificateRequest")
	proto.RegisterType((*CertificateResponse)(nil), "imperium.CertificateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Imperium service

type ImperiumClient interface {
	Request(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
}

type imperiumClient struct {
	cc *grpc.ClientConn
}

func NewImperiumClient(cc *grpc.ClientConn) ImperiumClient {
	return &imperiumClient{cc}
}

func (c *imperiumClient) Request(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := grpc.Invoke(ctx, "/imperium.Imperium/Request", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Imperium service

type ImperiumServer interface {
	Request(context.Context, *CertificateRequest) (*CertificateResponse, error)
}

func RegisterImperiumServer(s *grpc.Server, srv ImperiumServer) {
	s.RegisterService(&_Imperium_serviceDesc, srv)
}

func _Imperium_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImperiumServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/imperium.Imperium/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImperiumServer).Request(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Imperium_serviceDesc = grpc.ServiceDesc{
	ServiceName: "imperium.Imperium",
	HandlerType: (*ImperiumServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _Imperium_Request_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("proto/imperium.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x89, 0xda, 0x1a, 0x27, 0x07, 0x65, 0x54, 0x08, 0x41, 0x69, 0xe9, 0x49, 0x11, 0x36,
	0x50, 0x8f, 0x1e, 0x05, 0xa1, 0x78, 0x0b, 0xde, 0x43, 0xda, 0x4e, 0xeb, 0xa2, 0xdb, 0x1d, 0x37,
	0x1b, 0x31, 0xff, 0xd5, 0x1f, 0x63, 0xb2, 0xf9, 0x30, 0x22, 0xbd, 0xed, 0xfb, 0xce, 0xb3, 0x3b,
	0xec, 0x03, 0x17, 0x6c, 0xb4, 0xd5, 0xb1, 0x54, 0x4c, 0x46, 0x16, 0x4a, 0xb8, 0x88, 0x7e, 0x97,
	0xa3, 0xc9, 0x56, 0xeb, 0xed, 0x3b, 0xc5, 0xae, 0x5f, 0x16, 0x9b, 0xd8, 0x4a, 0x45, 0xb9, 0xcd,
	0x14, 0x37, 0xe8, 0xec, 0x0e, 0xf0, 0x91, 0x8c, 0x95, 0x1b, 0xb9, 0xca, 0x2c, 0x25, 0xf4, 0x51,
	0x54, 0x63, 0xbc, 0x84, 0x71, 0xc6, 0x9c, 0xca, 0x75, 0xe8, 0x4d, 0xbd, 0x9b, 0x93, 0x64, 0x54,
	0xa5, 0xc5, 0x7a, 0xf6, 0xed, 0xc1, 0xf9, 0x1f, 0x3a, 0x67, 0xbd, 0xcb, 0x09, 0xa7, 0x10, 0xac,
	0x7e, 0xeb, 0xf6, 0xce, 0xb0, 0xc2, 0x09, 0x04, 0x6c, 0xe4, 0x67, 0x75, 0x4c, 0xdf, 0xa8, 0x0c,
	0x0f, 0x1c, 0x01, 0x6d, 0xf5, 0x4c, 0x25, 0xde, 0xc2, 0x99, 0xd1, 0xda, 0xa6, 0xc3, 0x77, 0x0e,
	0x1d, 0x75, 0x5a, 0xf7, 0x83, 0xad, 0x18, 0x81, 0xff, 0xaa, 0x73, 0xbb, 0xcb, 0x14, 0x85, 0x47,
	0x0e, 0xe9, 0x33, 0x3e, 0x40, 0x40, 0x5f, 0x2c, 0x4d, 0x99, 0xd6, 0x1f, 0x0d, 0x47, 0xd5, 0x38,
	0x98, 0x47, 0xa2, 0xb1, 0x20, 0x3a, 0x0b, 0xe2, 0xa5, 0xb3, 0x90, 0x40, 0x83, 0xd7, 0xc5, 0x3c,
	0x01, 0x7f, 0xd1, 0x8a, 0xc3, 0x27, 0x38, 0xee, 0x64, 0x5c, 0x89, 0x5e, 0xef, 0x7f, 0x55, 0xd1,
	0xf5, 0x9e, 0x69, 0xa3, 0x66, 0x39, 0x76, 0x3b, 0xef, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf9,
	0xca, 0x51, 0xb7, 0xa9, 0x01, 0x00, 0x00,
}
