// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package configpb

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

// ConfigurationClient is the client API for Configuration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfigurationClient interface {
	Create(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	Get(ctx context.Context, in *ResourceKey, opts ...grpc.CallOption) (*Status, error)
	Update(ctx context.Context, in *Notification, opts ...grpc.CallOption) (*Reply, error)
	Delete(ctx context.Context, in *ResourceKey, opts ...grpc.CallOption) (*Reply, error)
	GetConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigReply, error)
	GetResourceName(ctx context.Context, in *ResourceRequest, opts ...grpc.CallOption) (*ResourceReply, error)
}

type configurationClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigurationClient(cc grpc.ClientConnInterface) ConfigurationClient {
	return &configurationClient{cc}
}

func (c *configurationClient) Create(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/config.Configuration/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) Get(ctx context.Context, in *ResourceKey, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/config.Configuration/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) Update(ctx context.Context, in *Notification, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/config.Configuration/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) Delete(ctx context.Context, in *ResourceKey, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/config.Configuration/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) GetConfig(ctx context.Context, in *ConfigRequest, opts ...grpc.CallOption) (*ConfigReply, error) {
	out := new(ConfigReply)
	err := c.cc.Invoke(ctx, "/config.Configuration/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configurationClient) GetResourceName(ctx context.Context, in *ResourceRequest, opts ...grpc.CallOption) (*ResourceReply, error) {
	out := new(ResourceReply)
	err := c.cc.Invoke(ctx, "/config.Configuration/GetResourceName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigurationServer is the server API for Configuration service.
// All implementations must embed UnimplementedConfigurationServer
// for forward compatibility
type ConfigurationServer interface {
	Create(context.Context, *Request) (*Reply, error)
	Get(context.Context, *ResourceKey) (*Status, error)
	Update(context.Context, *Notification) (*Reply, error)
	Delete(context.Context, *ResourceKey) (*Reply, error)
	GetConfig(context.Context, *ConfigRequest) (*ConfigReply, error)
	GetResourceName(context.Context, *ResourceRequest) (*ResourceReply, error)
	mustEmbedUnimplementedConfigurationServer()
}

// UnimplementedConfigurationServer must be embedded to have forward compatible implementations.
type UnimplementedConfigurationServer struct {
}

func (UnimplementedConfigurationServer) Create(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedConfigurationServer) Get(context.Context, *ResourceKey) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedConfigurationServer) Update(context.Context, *Notification) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedConfigurationServer) Delete(context.Context, *ResourceKey) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedConfigurationServer) GetConfig(context.Context, *ConfigRequest) (*ConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedConfigurationServer) GetResourceName(context.Context, *ResourceRequest) (*ResourceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResourceName not implemented")
}
func (UnimplementedConfigurationServer) mustEmbedUnimplementedConfigurationServer() {}

// UnsafeConfigurationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigurationServer will
// result in compilation errors.
type UnsafeConfigurationServer interface {
	mustEmbedUnimplementedConfigurationServer()
}

func RegisterConfigurationServer(s grpc.ServiceRegistrar, srv ConfigurationServer) {
	s.RegisterService(&Configuration_ServiceDesc, srv)
}

func _Configuration_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.Configuration/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).Create(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.Configuration/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).Get(ctx, req.(*ResourceKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Notification)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.Configuration/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).Update(ctx, req.(*Notification))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.Configuration/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).Delete(ctx, req.(*ResourceKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.Configuration/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).GetConfig(ctx, req.(*ConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configuration_GetResourceName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigurationServer).GetResourceName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config.Configuration/GetResourceName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigurationServer).GetResourceName(ctx, req.(*ResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Configuration_ServiceDesc is the grpc.ServiceDesc for Configuration service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Configuration_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "config.Configuration",
	HandlerType: (*ConfigurationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Configuration_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Configuration_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Configuration_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Configuration_Delete_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _Configuration_GetConfig_Handler,
		},
		{
			MethodName: "GetResourceName",
			Handler:    _Configuration_GetResourceName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config/configpb/config.proto",
}
