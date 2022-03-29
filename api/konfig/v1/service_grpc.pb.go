// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: konfig/v1/service.proto

package v1

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

// ConfiguratorServiceClient is the client API for ConfiguratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfiguratorServiceClient interface {
	// GetConfig for a specific application, version, environment and command.
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error)
}

type configuratorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfiguratorServiceClient(cc grpc.ClientConnInterface) ConfiguratorServiceClient {
	return &configuratorServiceClient{cc}
}

func (c *configuratorServiceClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	out := new(GetConfigResponse)
	err := c.cc.Invoke(ctx, "/konfig.v1.ConfiguratorService/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfiguratorServiceServer is the server API for ConfiguratorService service.
// All implementations must embed UnimplementedConfiguratorServiceServer
// for forward compatibility
type ConfiguratorServiceServer interface {
	// GetConfig for a specific application, version, environment and command.
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)
	mustEmbedUnimplementedConfiguratorServiceServer()
}

// UnimplementedConfiguratorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConfiguratorServiceServer struct {
}

func (UnimplementedConfiguratorServiceServer) GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedConfiguratorServiceServer) mustEmbedUnimplementedConfiguratorServiceServer() {}

// UnsafeConfiguratorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfiguratorServiceServer will
// result in compilation errors.
type UnsafeConfiguratorServiceServer interface {
	mustEmbedUnimplementedConfiguratorServiceServer()
}

func RegisterConfiguratorServiceServer(s grpc.ServiceRegistrar, srv ConfiguratorServiceServer) {
	s.RegisterService(&ConfiguratorService_ServiceDesc, srv)
}

func _ConfiguratorService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfiguratorServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/konfig.v1.ConfiguratorService/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfiguratorServiceServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfiguratorService_ServiceDesc is the grpc.ServiceDesc for ConfiguratorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfiguratorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "konfig.v1.ConfiguratorService",
	HandlerType: (*ConfiguratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfig",
			Handler:    _ConfiguratorService_GetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "konfig/v1/service.proto",
}
