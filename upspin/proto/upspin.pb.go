// Code generated by protoc-gen-go.
// source: upspin.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	upspin.proto

It has these top-level messages:
	Endpoint
	Location
	StoreGetRequest
	StoreGetResponse
	StorePutRequest
	StorePutResponse
	StoreDeleteRequest
	StoreDeleteResponse
	ConfigureRequest
	ConfigureResponse
	EndpointRequest
	EndpointResponse
	ServerUserNameRequest
	ServerUserNameResponse
	AuthenticateRequest
	Signature
	AuthenticateResponse
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

// Endpoint mirrors upspin.Endpoint.
type Endpoint struct {
	Transport int32  `protobuf:"varint,1,opt,name=transport" json:"transport,omitempty"`
	NetAddr   string `protobuf:"bytes,2,opt,name=net_addr,json=netAddr" json:"net_addr,omitempty"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (m *Endpoint) String() string            { return proto1.CompactTextString(m) }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Location mirrors upspin.Location.
type Location struct {
	Endpoint  *Endpoint `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
	Reference string    `protobuf:"bytes,2,opt,name=reference" json:"reference,omitempty"`
}

func (m *Location) Reset()                    { *m = Location{} }
func (m *Location) String() string            { return proto1.CompactTextString(m) }
func (*Location) ProtoMessage()               {}
func (*Location) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Location) GetEndpoint() *Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

type StoreGetRequest struct {
	Reference string `protobuf:"bytes,1,opt,name=reference" json:"reference,omitempty"`
}

func (m *StoreGetRequest) Reset()                    { *m = StoreGetRequest{} }
func (m *StoreGetRequest) String() string            { return proto1.CompactTextString(m) }
func (*StoreGetRequest) ProtoMessage()               {}
func (*StoreGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type StoreGetResponse struct {
	Data      []byte      `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Locations []*Location `protobuf:"bytes,2,rep,name=locations" json:"locations,omitempty"`
}

func (m *StoreGetResponse) Reset()                    { *m = StoreGetResponse{} }
func (m *StoreGetResponse) String() string            { return proto1.CompactTextString(m) }
func (*StoreGetResponse) ProtoMessage()               {}
func (*StoreGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StoreGetResponse) GetLocations() []*Location {
	if m != nil {
		return m.Locations
	}
	return nil
}

type StorePutRequest struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *StorePutRequest) Reset()                    { *m = StorePutRequest{} }
func (m *StorePutRequest) String() string            { return proto1.CompactTextString(m) }
func (*StorePutRequest) ProtoMessage()               {}
func (*StorePutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type StorePutResponse struct {
	Reference string `protobuf:"bytes,1,opt,name=reference" json:"reference,omitempty"`
}

func (m *StorePutResponse) Reset()                    { *m = StorePutResponse{} }
func (m *StorePutResponse) String() string            { return proto1.CompactTextString(m) }
func (*StorePutResponse) ProtoMessage()               {}
func (*StorePutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type StoreDeleteRequest struct {
	Reference string `protobuf:"bytes,1,opt,name=reference" json:"reference,omitempty"`
}

func (m *StoreDeleteRequest) Reset()                    { *m = StoreDeleteRequest{} }
func (m *StoreDeleteRequest) String() string            { return proto1.CompactTextString(m) }
func (*StoreDeleteRequest) ProtoMessage()               {}
func (*StoreDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type StoreDeleteResponse struct {
}

func (m *StoreDeleteResponse) Reset()                    { *m = StoreDeleteResponse{} }
func (m *StoreDeleteResponse) String() string            { return proto1.CompactTextString(m) }
func (*StoreDeleteResponse) ProtoMessage()               {}
func (*StoreDeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type ConfigureRequest struct {
	Options []string `protobuf:"bytes,1,rep,name=options" json:"options,omitempty"`
}

func (m *ConfigureRequest) Reset()                    { *m = ConfigureRequest{} }
func (m *ConfigureRequest) String() string            { return proto1.CompactTextString(m) }
func (*ConfigureRequest) ProtoMessage()               {}
func (*ConfigureRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type ConfigureResponse struct {
}

func (m *ConfigureResponse) Reset()                    { *m = ConfigureResponse{} }
func (m *ConfigureResponse) String() string            { return proto1.CompactTextString(m) }
func (*ConfigureResponse) ProtoMessage()               {}
func (*ConfigureResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type EndpointRequest struct {
}

func (m *EndpointRequest) Reset()                    { *m = EndpointRequest{} }
func (m *EndpointRequest) String() string            { return proto1.CompactTextString(m) }
func (*EndpointRequest) ProtoMessage()               {}
func (*EndpointRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type EndpointResponse struct {
	Endpoint *Endpoint `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
}

func (m *EndpointResponse) Reset()                    { *m = EndpointResponse{} }
func (m *EndpointResponse) String() string            { return proto1.CompactTextString(m) }
func (*EndpointResponse) ProtoMessage()               {}
func (*EndpointResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *EndpointResponse) GetEndpoint() *Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

type ServerUserNameRequest struct {
}

func (m *ServerUserNameRequest) Reset()                    { *m = ServerUserNameRequest{} }
func (m *ServerUserNameRequest) String() string            { return proto1.CompactTextString(m) }
func (*ServerUserNameRequest) ProtoMessage()               {}
func (*ServerUserNameRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type ServerUserNameResponse struct {
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
}

func (m *ServerUserNameResponse) Reset()                    { *m = ServerUserNameResponse{} }
func (m *ServerUserNameResponse) String() string            { return proto1.CompactTextString(m) }
func (*ServerUserNameResponse) ProtoMessage()               {}
func (*ServerUserNameResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

// Authenticate is not part of the service definitions in upspin.
// It is used only in the remote implementations.
type AuthenticateRequest struct {
	UserName  string     `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Now       string     `protobuf:"bytes,2,opt,name=now" json:"now,omitempty"`
	Signature *Signature `protobuf:"bytes,3,opt,name=signature" json:"signature,omitempty"`
}

func (m *AuthenticateRequest) Reset()                    { *m = AuthenticateRequest{} }
func (m *AuthenticateRequest) String() string            { return proto1.CompactTextString(m) }
func (*AuthenticateRequest) ProtoMessage()               {}
func (*AuthenticateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *AuthenticateRequest) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Signature struct {
	R string `protobuf:"bytes,1,opt,name=r" json:"r,omitempty"`
	S string `protobuf:"bytes,2,opt,name=s" json:"s,omitempty"`
}

func (m *Signature) Reset()                    { *m = Signature{} }
func (m *Signature) String() string            { return proto1.CompactTextString(m) }
func (*Signature) ProtoMessage()               {}
func (*Signature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

type AuthenticateResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *AuthenticateResponse) Reset()                    { *m = AuthenticateResponse{} }
func (m *AuthenticateResponse) String() string            { return proto1.CompactTextString(m) }
func (*AuthenticateResponse) ProtoMessage()               {}
func (*AuthenticateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func init() {
	proto1.RegisterType((*Endpoint)(nil), "proto.Endpoint")
	proto1.RegisterType((*Location)(nil), "proto.Location")
	proto1.RegisterType((*StoreGetRequest)(nil), "proto.StoreGetRequest")
	proto1.RegisterType((*StoreGetResponse)(nil), "proto.StoreGetResponse")
	proto1.RegisterType((*StorePutRequest)(nil), "proto.StorePutRequest")
	proto1.RegisterType((*StorePutResponse)(nil), "proto.StorePutResponse")
	proto1.RegisterType((*StoreDeleteRequest)(nil), "proto.StoreDeleteRequest")
	proto1.RegisterType((*StoreDeleteResponse)(nil), "proto.StoreDeleteResponse")
	proto1.RegisterType((*ConfigureRequest)(nil), "proto.ConfigureRequest")
	proto1.RegisterType((*ConfigureResponse)(nil), "proto.ConfigureResponse")
	proto1.RegisterType((*EndpointRequest)(nil), "proto.EndpointRequest")
	proto1.RegisterType((*EndpointResponse)(nil), "proto.EndpointResponse")
	proto1.RegisterType((*ServerUserNameRequest)(nil), "proto.ServerUserNameRequest")
	proto1.RegisterType((*ServerUserNameResponse)(nil), "proto.ServerUserNameResponse")
	proto1.RegisterType((*AuthenticateRequest)(nil), "proto.AuthenticateRequest")
	proto1.RegisterType((*Signature)(nil), "proto.Signature")
	proto1.RegisterType((*AuthenticateResponse)(nil), "proto.AuthenticateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Store service

type StoreClient interface {
	Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error)
	Get(ctx context.Context, in *StoreGetRequest, opts ...grpc.CallOption) (*StoreGetResponse, error)
	Put(ctx context.Context, in *StorePutRequest, opts ...grpc.CallOption) (*StorePutResponse, error)
	Delete(ctx context.Context, in *StoreDeleteRequest, opts ...grpc.CallOption) (*StoreDeleteResponse, error)
	Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*ConfigureResponse, error)
	Endpoint(ctx context.Context, in *EndpointRequest, opts ...grpc.CallOption) (*EndpointResponse, error)
	ServerUserName(ctx context.Context, in *ServerUserNameRequest, opts ...grpc.CallOption) (*ServerUserNameResponse, error)
}

type storeClient struct {
	cc *grpc.ClientConn
}

func NewStoreClient(cc *grpc.ClientConn) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) Authenticate(ctx context.Context, in *AuthenticateRequest, opts ...grpc.CallOption) (*AuthenticateResponse, error) {
	out := new(AuthenticateResponse)
	err := grpc.Invoke(ctx, "/proto.Store/Authenticate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Get(ctx context.Context, in *StoreGetRequest, opts ...grpc.CallOption) (*StoreGetResponse, error) {
	out := new(StoreGetResponse)
	err := grpc.Invoke(ctx, "/proto.Store/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Put(ctx context.Context, in *StorePutRequest, opts ...grpc.CallOption) (*StorePutResponse, error) {
	out := new(StorePutResponse)
	err := grpc.Invoke(ctx, "/proto.Store/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Delete(ctx context.Context, in *StoreDeleteRequest, opts ...grpc.CallOption) (*StoreDeleteResponse, error) {
	out := new(StoreDeleteResponse)
	err := grpc.Invoke(ctx, "/proto.Store/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Configure(ctx context.Context, in *ConfigureRequest, opts ...grpc.CallOption) (*ConfigureResponse, error) {
	out := new(ConfigureResponse)
	err := grpc.Invoke(ctx, "/proto.Store/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Endpoint(ctx context.Context, in *EndpointRequest, opts ...grpc.CallOption) (*EndpointResponse, error) {
	out := new(EndpointResponse)
	err := grpc.Invoke(ctx, "/proto.Store/Endpoint", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) ServerUserName(ctx context.Context, in *ServerUserNameRequest, opts ...grpc.CallOption) (*ServerUserNameResponse, error) {
	out := new(ServerUserNameResponse)
	err := grpc.Invoke(ctx, "/proto.Store/ServerUserName", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Store service

type StoreServer interface {
	Authenticate(context.Context, *AuthenticateRequest) (*AuthenticateResponse, error)
	Get(context.Context, *StoreGetRequest) (*StoreGetResponse, error)
	Put(context.Context, *StorePutRequest) (*StorePutResponse, error)
	Delete(context.Context, *StoreDeleteRequest) (*StoreDeleteResponse, error)
	Configure(context.Context, *ConfigureRequest) (*ConfigureResponse, error)
	Endpoint(context.Context, *EndpointRequest) (*EndpointResponse, error)
	ServerUserName(context.Context, *ServerUserNameRequest) (*ServerUserNameResponse, error)
}

func RegisterStoreServer(s *grpc.Server, srv StoreServer) {
	s.RegisterService(&_Store_serviceDesc, srv)
}

func _Store_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Store/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Authenticate(ctx, req.(*AuthenticateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Store/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Get(ctx, req.(*StoreGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StorePutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Store/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Put(ctx, req.(*StorePutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Store/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Delete(ctx, req.(*StoreDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Store/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Configure(ctx, req.(*ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Endpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Endpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Store/Endpoint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Endpoint(ctx, req.(*EndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_ServerUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerUserNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).ServerUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Store/ServerUserName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).ServerUserName(ctx, req.(*ServerUserNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Store_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _Store_Authenticate_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Store_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Store_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Store_Delete_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Store_Configure_Handler,
		},
		{
			MethodName: "Endpoint",
			Handler:    _Store_Endpoint_Handler,
		},
		{
			MethodName: "ServerUserName",
			Handler:    _Store_ServerUserName_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x54, 0x5f, 0x6f, 0xd3, 0x4e,
	0x10, 0x8c, 0xeb, 0x5f, 0xda, 0x78, 0x7f, 0x11, 0x71, 0x37, 0x6d, 0xea, 0xba, 0x45, 0x8a, 0x4e,
	0x42, 0x54, 0xa2, 0x04, 0x14, 0x84, 0xc4, 0x0b, 0x82, 0xa8, 0xa0, 0x0a, 0x09, 0x41, 0xe5, 0x28,
	0xcf, 0x95, 0x89, 0xb7, 0xc5, 0xa2, 0xbd, 0x33, 0x77, 0x67, 0xf8, 0xb0, 0x7c, 0x19, 0x14, 0xfb,
	0xfc, 0xb7, 0xa6, 0x82, 0xa7, 0xd8, 0xbb, 0x3b, 0xb3, 0x33, 0xd9, 0x49, 0x60, 0x98, 0x26, 0x2a,
	0x89, 0xf9, 0x2c, 0x91, 0x42, 0x0b, 0xec, 0x67, 0x1f, 0xec, 0x0c, 0x06, 0xef, 0x79, 0x94, 0x88,
	0x98, 0x6b, 0x3c, 0x06, 0x47, 0xcb, 0x90, 0xab, 0x44, 0x48, 0xed, 0x59, 0x53, 0xeb, 0xa4, 0x1f,
	0x54, 0x05, 0x3c, 0x84, 0x01, 0x27, 0x7d, 0x19, 0x46, 0x91, 0xf4, 0xb6, 0xa6, 0xd6, 0x89, 0x13,
	0xec, 0x70, 0xd2, 0x8b, 0x28, 0x92, 0x6c, 0x05, 0x83, 0x8f, 0x62, 0x1d, 0xea, 0x58, 0x70, 0x7c,
	0x02, 0x03, 0x32, 0x84, 0x19, 0xc7, 0xff, 0xf3, 0x51, 0xbe, 0x71, 0x56, 0xec, 0x09, 0xca, 0x81,
	0xcd, 0x46, 0x49, 0x57, 0x24, 0x89, 0xaf, 0xc9, 0x90, 0x56, 0x05, 0xf6, 0x0c, 0x46, 0x4b, 0x2d,
	0x24, 0x9d, 0x93, 0x0e, 0xe8, 0x7b, 0x4a, 0xaa, 0x05, 0xb0, 0xda, 0x80, 0x15, 0xb8, 0x15, 0x40,
	0x25, 0x82, 0x2b, 0x42, 0x84, 0xff, 0xa2, 0x50, 0x87, 0xd9, 0xf0, 0x30, 0xc8, 0x9e, 0xf1, 0x29,
	0x38, 0x37, 0x46, 0xaf, 0xf2, 0xb6, 0xa6, 0x76, 0x4d, 0x64, 0xe1, 0x23, 0xa8, 0x26, 0xd8, 0x23,
	0xa3, 0xe3, 0x22, 0x2d, 0x75, 0x74, 0xb0, 0xb2, 0xe7, 0x66, 0x7b, 0x36, 0x66, 0xb6, 0xdf, 0xaf,
	0x77, 0x0e, 0x98, 0x21, 0xde, 0xd1, 0x0d, 0x69, 0xfa, 0x3b, 0x8f, 0xfb, 0x30, 0x6e, 0x60, 0xf2,
	0x45, 0xec, 0x14, 0xdc, 0x33, 0xc1, 0xaf, 0xe2, 0xeb, 0x54, 0x96, 0x44, 0x1e, 0xec, 0x88, 0x24,
	0x37, 0x69, 0x4d, 0xed, 0xcd, 0xc1, 0xcc, 0x2b, 0x1b, 0xc3, 0x6e, 0x6d, 0xda, 0x50, 0xec, 0xc2,
	0xa8, 0x3c, 0x51, 0xce, 0xc0, 0xde, 0x80, 0x5b, 0x95, 0x8c, 0xa5, 0x7f, 0x39, 0x30, 0x3b, 0x80,
	0xfd, 0x25, 0xc9, 0x1f, 0x24, 0x57, 0x8a, 0xe4, 0xa7, 0xf0, 0xb6, 0xd0, 0xc6, 0x5e, 0xc2, 0xa4,
	0xdd, 0x30, 0xfc, 0x47, 0xe0, 0xa4, 0x8a, 0xe4, 0x25, 0x0f, 0x6f, 0x0b, 0xfb, 0x83, 0xd4, 0x0c,
	0x31, 0x0d, 0xe3, 0x45, 0xaa, 0xbf, 0x12, 0xd7, 0xf1, 0x3a, 0xac, 0xbe, 0xb2, 0xfb, 0x30, 0xe8,
	0x82, 0xcd, 0xc5, 0x4f, 0x13, 0xaf, 0xcd, 0x23, 0xce, 0xc0, 0x51, 0xf1, 0x35, 0x0f, 0x75, 0x2a,
	0xc9, 0xb3, 0x33, 0x0f, 0xae, 0xf1, 0xb0, 0x2c, 0xea, 0x41, 0x35, 0xc2, 0x1e, 0x83, 0x53, 0xd6,
	0x71, 0x08, 0x96, 0x34, 0x3b, 0x2c, 0xb9, 0x79, 0x53, 0x86, 0xda, 0x52, 0xec, 0x14, 0xf6, 0x9a,
	0xf2, 0x8c, 0xa7, 0x3d, 0xe8, 0x6b, 0xf1, 0x8d, 0xb8, 0xc1, 0xe5, 0x2f, 0xf3, 0x5f, 0x36, 0xf4,
	0xb3, 0x5b, 0xe2, 0x07, 0x18, 0xd6, 0x71, 0xe8, 0x1b, 0x35, 0x1d, 0x5e, 0xfd, 0xa3, 0xce, 0x9e,
	0xb9, 0x61, 0x0f, 0x5f, 0x81, 0x7d, 0x4e, 0x1a, 0x27, 0x85, 0x9f, 0xe6, 0x0f, 0xc8, 0x3f, 0xb8,
	0x53, 0xaf, 0x23, 0x2f, 0xd2, 0x16, 0xb2, 0x8a, 0x7c, 0x13, 0x59, 0xcb, 0x38, 0xeb, 0xe1, 0x02,
	0xb6, 0xf3, 0x38, 0xe2, 0x61, 0x7d, 0xa8, 0x11, 0x6b, 0xdf, 0xef, 0x6a, 0x95, 0x14, 0x6f, 0xc1,
	0x29, 0x13, 0x89, 0xc5, 0xaa, 0x76, 0xa2, 0x7d, 0xef, 0x6e, 0xa3, 0x64, 0x78, 0x5d, 0xfb, 0x27,
	0x9b, 0xb4, 0x13, 0xd9, 0xf2, 0xd0, 0x0e, 0x35, 0xeb, 0xe1, 0x67, 0x78, 0xd0, 0x0c, 0x24, 0x1e,
	0x17, 0x82, 0xbb, 0x02, 0xec, 0x3f, 0xfc, 0x43, 0xb7, 0x20, 0xfc, 0xb2, 0x9d, 0xf5, 0x5f, 0xfc,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0x40, 0xa7, 0xb4, 0xb9, 0x77, 0x05, 0x00, 0x00,
}