// Code generated by protoc-gen-go.
// source: proto/discovery.proto
// DO NOT EDIT!

/*
Package discovery is a generated protocol buffer package.

It is generated from these files:
	proto/discovery.proto

It has these top-level messages:
	RegisterRequest
	DeRegisterRequest
	HeartBeatRequest
	LocationRequest
	StatusRequest
	DiscoveryResponse
	ServiceLocation
*/
package discovery

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

type ServiceStatus int32

const (
	ServiceStatus_OFFLINE              ServiceStatus = 0
	ServiceStatus_ONLINE               ServiceStatus = 1
	ServiceStatus_DEGRADED_PERFORMANCE ServiceStatus = 2
	ServiceStatus_PARTIAL_OUTAGE       ServiceStatus = 3
	ServiceStatus_UNDER_MAINTENANCE    ServiceStatus = 5
)

var ServiceStatus_name = map[int32]string{
	0: "OFFLINE",
	1: "ONLINE",
	2: "DEGRADED_PERFORMANCE",
	3: "PARTIAL_OUTAGE",
	5: "UNDER_MAINTENANCE",
}
var ServiceStatus_value = map[string]int32{
	"OFFLINE":              0,
	"ONLINE":               1,
	"DEGRADED_PERFORMANCE": 2,
	"PARTIAL_OUTAGE":       3,
	"UNDER_MAINTENANCE":    5,
}

func (x ServiceStatus) String() string {
	return proto.EnumName(ServiceStatus_name, int32(x))
}
func (ServiceStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StatusTarget int32

const (
	StatusTarget_INSTANCE StatusTarget = 0
	StatusTarget_SERVICE  StatusTarget = 1
	StatusTarget_BOTH     StatusTarget = 2
)

var StatusTarget_name = map[int32]string{
	0: "INSTANCE",
	1: "SERVICE",
	2: "BOTH",
}
var StatusTarget_value = map[string]int32{
	"INSTANCE": 0,
	"SERVICE":  1,
	"BOTH":     2,
}

func (x StatusTarget) String() string {
	return proto.EnumName(StatusTarget_name, int32(x))
}
func (StatusTarget) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type RegisterRequest struct {
	AppId        string `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	InstanceUuid string `protobuf:"bytes,2,opt,name=instance_uuid,json=instanceUuid" json:"instance_uuid,omitempty"`
	ServiceHost  string `protobuf:"bytes,3,opt,name=service_host,json=serviceHost" json:"service_host,omitempty"`
	ServicePort  int32  `protobuf:"varint,4,opt,name=service_port,json=servicePort" json:"service_port,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DeRegisterRequest struct {
	AppId        string `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	InstanceUuid string `protobuf:"bytes,2,opt,name=instance_uuid,json=instanceUuid" json:"instance_uuid,omitempty"`
}

func (m *DeRegisterRequest) Reset()                    { *m = DeRegisterRequest{} }
func (m *DeRegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*DeRegisterRequest) ProtoMessage()               {}
func (*DeRegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type HeartBeatRequest struct {
	AppId        string `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	InstanceUuid string `protobuf:"bytes,2,opt,name=instance_uuid,json=instanceUuid" json:"instance_uuid,omitempty"`
}

func (m *HeartBeatRequest) Reset()                    { *m = HeartBeatRequest{} }
func (m *HeartBeatRequest) String() string            { return proto.CompactTextString(m) }
func (*HeartBeatRequest) ProtoMessage()               {}
func (*HeartBeatRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type LocationRequest struct {
	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
}

func (m *LocationRequest) Reset()                    { *m = LocationRequest{} }
func (m *LocationRequest) String() string            { return proto.CompactTextString(m) }
func (*LocationRequest) ProtoMessage()               {}
func (*LocationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type StatusRequest struct {
	AppId        string        `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	InstanceUuid string        `protobuf:"bytes,2,opt,name=instance_uuid,json=instanceUuid" json:"instance_uuid,omitempty"`
	Message      string        `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Status       ServiceStatus `protobuf:"varint,4,opt,name=status,enum=discovery.ServiceStatus" json:"status,omitempty"`
	Target       StatusTarget  `protobuf:"varint,5,opt,name=target,enum=discovery.StatusTarget" json:"target,omitempty"`
}

func (m *StatusRequest) Reset()                    { *m = StatusRequest{} }
func (m *StatusRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()               {}
func (*StatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type DiscoveryResponse struct {
	Recorded bool `protobuf:"varint,1,opt,name=recorded" json:"recorded,omitempty"`
}

func (m *DiscoveryResponse) Reset()                    { *m = DiscoveryResponse{} }
func (m *DiscoveryResponse) String() string            { return proto.CompactTextString(m) }
func (*DiscoveryResponse) ProtoMessage()               {}
func (*DiscoveryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ServiceLocation struct {
	AppId       string        `protobuf:"bytes,1,opt,name=app_id,json=appId" json:"app_id,omitempty"`
	ServiceHost string        `protobuf:"bytes,2,opt,name=service_host,json=serviceHost" json:"service_host,omitempty"`
	ServicePort int32         `protobuf:"varint,3,opt,name=service_port,json=servicePort" json:"service_port,omitempty"`
	Status      ServiceStatus `protobuf:"varint,4,opt,name=status,enum=discovery.ServiceStatus" json:"status,omitempty"`
}

func (m *ServiceLocation) Reset()                    { *m = ServiceLocation{} }
func (m *ServiceLocation) String() string            { return proto.CompactTextString(m) }
func (*ServiceLocation) ProtoMessage()               {}
func (*ServiceLocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*RegisterRequest)(nil), "discovery.RegisterRequest")
	proto.RegisterType((*DeRegisterRequest)(nil), "discovery.DeRegisterRequest")
	proto.RegisterType((*HeartBeatRequest)(nil), "discovery.HeartBeatRequest")
	proto.RegisterType((*LocationRequest)(nil), "discovery.LocationRequest")
	proto.RegisterType((*StatusRequest)(nil), "discovery.StatusRequest")
	proto.RegisterType((*DiscoveryResponse)(nil), "discovery.DiscoveryResponse")
	proto.RegisterType((*ServiceLocation)(nil), "discovery.ServiceLocation")
	proto.RegisterEnum("discovery.ServiceStatus", ServiceStatus_name, ServiceStatus_value)
	proto.RegisterEnum("discovery.StatusTarget", StatusTarget_name, StatusTarget_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Discovery service

type DiscoveryClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
	DeRegister(ctx context.Context, in *DeRegisterRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
	HeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
	GetLocation(ctx context.Context, in *LocationRequest, opts ...grpc.CallOption) (*ServiceLocation, error)
}

type discoveryClient struct {
	cc *grpc.ClientConn
}

func NewDiscoveryClient(cc *grpc.ClientConn) DiscoveryClient {
	return &discoveryClient{cc}
}

func (c *discoveryClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := grpc.Invoke(ctx, "/discovery.Discovery/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) DeRegister(ctx context.Context, in *DeRegisterRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := grpc.Invoke(ctx, "/discovery.Discovery/DeRegister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) HeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := grpc.Invoke(ctx, "/discovery.Discovery/HeartBeat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := grpc.Invoke(ctx, "/discovery.Discovery/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) GetLocation(ctx context.Context, in *LocationRequest, opts ...grpc.CallOption) (*ServiceLocation, error) {
	out := new(ServiceLocation)
	err := grpc.Invoke(ctx, "/discovery.Discovery/GetLocation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Discovery service

type DiscoveryServer interface {
	Register(context.Context, *RegisterRequest) (*DiscoveryResponse, error)
	DeRegister(context.Context, *DeRegisterRequest) (*DiscoveryResponse, error)
	HeartBeat(context.Context, *HeartBeatRequest) (*DiscoveryResponse, error)
	Status(context.Context, *StatusRequest) (*DiscoveryResponse, error)
	GetLocation(context.Context, *LocationRequest) (*ServiceLocation, error)
}

func RegisterDiscoveryServer(s *grpc.Server, srv DiscoveryServer) {
	s.RegisterService(&_Discovery_serviceDesc, srv)
}

func _Discovery_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_DeRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).DeRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/DeRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).DeRegister(ctx, req.(*DeRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).HeartBeat(ctx, req.(*HeartBeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_GetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).GetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/discovery.Discovery/GetLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).GetLocation(ctx, req.(*LocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Discovery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "discovery.Discovery",
	HandlerType: (*DiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Discovery_Register_Handler,
		},
		{
			MethodName: "DeRegister",
			Handler:    _Discovery_DeRegister_Handler,
		},
		{
			MethodName: "HeartBeat",
			Handler:    _Discovery_HeartBeat_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Discovery_Status_Handler,
		},
		{
			MethodName: "GetLocation",
			Handler:    _Discovery_GetLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("proto/discovery.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x94, 0x51, 0x6e, 0xda, 0x4c,
	0x10, 0xc7, 0x03, 0x04, 0xc7, 0x0c, 0x24, 0x71, 0x56, 0x5f, 0xf4, 0x59, 0x6e, 0x1f, 0x5a, 0xf7,
	0x25, 0xca, 0x43, 0xa8, 0x92, 0x0b, 0xd4, 0x89, 0x17, 0xb0, 0x44, 0x6c, 0xb4, 0x98, 0xbe, 0x5a,
	0x2e, 0x8c, 0xa8, 0x1f, 0x82, 0x5d, 0xef, 0x12, 0xa9, 0xe7, 0xe8, 0x05, 0x7a, 0x9c, 0x1e, 0xa6,
	0x87, 0xe8, 0x66, 0x63, 0x3b, 0x2e, 0xa8, 0x4d, 0x23, 0xf1, 0xc6, 0xcc, 0xfc, 0x3c, 0xb3, 0xf3,
	0x9f, 0x19, 0xe0, 0x34, 0xcb, 0x53, 0x91, 0xf6, 0x17, 0x09, 0x9f, 0xa7, 0xf7, 0x98, 0x7f, 0xbd,
	0x50, 0x36, 0xe9, 0x54, 0x0e, 0xfb, 0x5b, 0x03, 0x8e, 0x19, 0x2e, 0x13, 0x2e, 0x30, 0x67, 0xf8,
	0x65, 0x8d, 0x5c, 0x90, 0x53, 0xd0, 0xe2, 0x2c, 0x8b, 0x92, 0x85, 0xd9, 0x78, 0xd3, 0x38, 0xeb,
	0xb0, 0xb6, 0xb4, 0xbc, 0x05, 0x79, 0x07, 0x87, 0xc9, 0x8a, 0x8b, 0x78, 0x35, 0xc7, 0x68, 0xbd,
	0x96, 0xd1, 0xa6, 0x8a, 0xf6, 0x4a, 0xe7, 0x4c, 0xfa, 0xc8, 0x5b, 0xe8, 0x71, 0xcc, 0xef, 0x13,
	0xc9, 0x7c, 0x4e, 0xb9, 0x30, 0x5b, 0x8a, 0xe9, 0x16, 0xbe, 0x91, 0x74, 0xd5, 0x91, 0x2c, 0xcd,
	0x85, 0xb9, 0x2f, 0x91, 0x76, 0x85, 0x4c, 0xa4, 0xcb, 0x0e, 0xe0, 0xc4, 0xc5, 0x1d, 0x3e, 0xcb,
	0xf6, 0xc1, 0x18, 0x61, 0x9c, 0x8b, 0x6b, 0x8c, 0xc5, 0x2e, 0xf2, 0x9d, 0xc1, 0xf1, 0x38, 0x9d,
	0xc7, 0x22, 0x49, 0x57, 0x7f, 0x4f, 0x67, 0xff, 0x68, 0xc0, 0xe1, 0x54, 0xc4, 0x62, 0xcd, 0x77,
	0x21, 0xaf, 0x09, 0x07, 0x77, 0xc8, 0x79, 0xbc, 0xc4, 0x42, 0xd9, 0xd2, 0x24, 0xef, 0x41, 0xe3,
	0xaa, 0x8c, 0xd2, 0xf3, 0xe8, 0xd2, 0xbc, 0x78, 0x9a, 0xfa, 0xf4, 0x51, 0xda, 0xe2, 0x19, 0x05,
	0x47, 0xfa, 0xa0, 0x89, 0x38, 0x5f, 0xa2, 0x30, 0xdb, 0xea, 0x8b, 0xff, 0xeb, 0x5f, 0x28, 0x24,
	0x54, 0x61, 0x56, 0x60, 0x76, 0x5f, 0x4e, 0xa5, 0x24, 0x18, 0xf2, 0x2c, 0x5d, 0x71, 0x24, 0x16,
	0xe8, 0x39, 0xce, 0xd3, 0x7c, 0x81, 0x8f, 0xfd, 0xe8, 0xac, 0xb2, 0xed, 0xef, 0x72, 0xb9, 0x8a,
	0xda, 0xa5, 0x5a, 0x7f, 0xea, 0x7e, 0x73, 0x6f, 0x9a, 0xcf, 0xef, 0x4d, 0x6b, 0x6b, 0x6f, 0x5e,
	0x2e, 0xc2, 0xf9, 0x9d, 0x9c, 0x4e, 0x3d, 0x40, 0xba, 0x70, 0x10, 0x0c, 0x06, 0x63, 0xcf, 0xa7,
	0xc6, 0x1e, 0x01, 0xd0, 0x02, 0x5f, 0xfd, 0x6e, 0x48, 0xe9, 0xff, 0x73, 0xe9, 0x90, 0x39, 0x2e,
	0x75, 0xa3, 0x09, 0x65, 0x83, 0x80, 0xdd, 0x3a, 0xfe, 0x0d, 0x35, 0x9a, 0x84, 0xc0, 0xd1, 0xc4,
	0x61, 0xa1, 0xe7, 0x8c, 0xa3, 0x60, 0x16, 0x3a, 0x43, 0x6a, 0xb4, 0x64, 0x9b, 0x27, 0x33, 0xdf,
	0xa5, 0x2c, 0xba, 0x75, 0x3c, 0x3f, 0xa4, 0xbe, 0x42, 0xdb, 0xe7, 0x57, 0xd0, 0xab, 0x4b, 0x4b,
	0x7a, 0xa0, 0x7b, 0xfe, 0x34, 0x54, 0xd1, 0xbd, 0x87, 0xda, 0x53, 0xca, 0x3e, 0x7a, 0x37, 0x0f,
	0xf5, 0x74, 0xd8, 0xbf, 0x0e, 0xc2, 0x91, 0xd1, 0xbc, 0xfc, 0xd9, 0x84, 0x4e, 0x25, 0x3c, 0x71,
	0x41, 0x2f, 0x2f, 0x83, 0x58, 0xb5, 0xfe, 0x36, 0xce, 0xc5, 0x7a, 0x5d, 0x8b, 0x6d, 0x8f, 0x6d,
	0x04, 0xf0, 0x74, 0x61, 0xe4, 0x37, 0x16, 0x5f, 0x96, 0x69, 0x00, 0x9d, 0xea, 0xb4, 0xc8, 0xab,
	0x1a, 0xba, 0x79, 0x70, 0xcf, 0xe4, 0xf9, 0x00, 0x5a, 0x31, 0x02, 0x73, 0x6b, 0x11, 0xff, 0x2d,
	0x03, 0x85, 0xee, 0x10, 0x45, 0xb5, 0x69, 0x75, 0x71, 0x36, 0x8e, 0xd5, 0xb2, 0xb6, 0x17, 0xa3,
	0x44, 0x3e, 0x69, 0xea, 0x4f, 0xf2, 0xea, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x35, 0x5b, 0x00,
	0xd1, 0x3d, 0x05, 0x00, 0x00,
}
