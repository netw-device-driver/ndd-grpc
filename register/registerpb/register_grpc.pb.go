// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package registerpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RegistrationClient is the client API for Registration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistrationClient interface {
	Create(ctx context.Context, in *RegistrationInfo, opts ...grpc.CallOption) (*DeviceType, error)
	Read(ctx context.Context, in *DeviceType, opts ...grpc.CallOption) (*RegistrationInfo, error)
	Update(ctx context.Context, in *RegistrationInfo, opts ...grpc.CallOption) (*DeviceType, error)
	Delete(ctx context.Context, in *DeviceType, opts ...grpc.CallOption) (*DeviceType, error)
}

type registrationClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistrationClient(cc grpc.ClientConnInterface) RegistrationClient {
	return &registrationClient{cc}
}

func (c *registrationClient) Create(ctx context.Context, in *RegistrationInfo, opts ...grpc.CallOption) (*DeviceType, error) {
	out := new(DeviceType)
	err := c.cc.Invoke(ctx, "/register.Registration/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) Read(ctx context.Context, in *DeviceType, opts ...grpc.CallOption) (*RegistrationInfo, error) {
	out := new(RegistrationInfo)
	err := c.cc.Invoke(ctx, "/register.Registration/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) Update(ctx context.Context, in *RegistrationInfo, opts ...grpc.CallOption) (*DeviceType, error) {
	out := new(DeviceType)
	err := c.cc.Invoke(ctx, "/register.Registration/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationClient) Delete(ctx context.Context, in *DeviceType, opts ...grpc.CallOption) (*DeviceType, error) {
	out := new(DeviceType)
	err := c.cc.Invoke(ctx, "/register.Registration/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistrationServer is the server API for Registration service.
// All implementations must embed UnimplementedRegistrationServer
// for forward compatibility
type RegistrationServer interface {
	Create(context.Context, *RegistrationInfo) (*DeviceType, error)
	Read(context.Context, *DeviceType) (*RegistrationInfo, error)
	Update(context.Context, *RegistrationInfo) (*DeviceType, error)
	Delete(context.Context, *DeviceType) (*DeviceType, error)
	mustEmbedUnimplementedRegistrationServer()
}

// UnimplementedRegistrationServer must be embedded to have forward compatible implementations.
type UnimplementedRegistrationServer struct {
}

func (UnimplementedRegistrationServer) Create(context.Context, *RegistrationInfo) (*DeviceType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRegistrationServer) Read(context.Context, *DeviceType) (*RegistrationInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedRegistrationServer) Update(context.Context, *RegistrationInfo) (*DeviceType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRegistrationServer) Delete(context.Context, *DeviceType) (*DeviceType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRegistrationServer) mustEmbedUnimplementedRegistrationServer() {}

// UnsafeRegistrationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistrationServer will
// result in compilation errors.
type UnsafeRegistrationServer interface {
	mustEmbedUnimplementedRegistrationServer()
}

func RegisterRegistrationServer(s grpc.ServiceRegistrar, srv RegistrationServer) {
	s.RegisterService(&Registration_ServiceDesc, srv)
}

func _Registration_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/register.Registration/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).Create(ctx, req.(*RegistrationInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/register.Registration/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).Read(ctx, req.(*DeviceType))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistrationInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/register.Registration/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).Update(ctx, req.(*RegistrationInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Registration_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceType)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/register.Registration/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationServer).Delete(ctx, req.(*DeviceType))
	}
	return interceptor(ctx, in, info, handler)
}

// Registration_ServiceDesc is the grpc.ServiceDesc for Registration service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Registration_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "register.Registration",
	HandlerType: (*RegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Registration_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Registration_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Registration_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Registration_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "register/registerpb/register.proto",
}
