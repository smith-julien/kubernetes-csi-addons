// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: networkfence.proto

package proto

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

// NetworkFenceClient is the client API for NetworkFence service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkFenceClient interface {
	// FenceClusterNetwork RPC call to fence the cluster network.
	FenceClusterNetwork(ctx context.Context, in *NetworkFenceRequest, opts ...grpc.CallOption) (*NetworkFenceResponse, error)
	// UnFenceClusterNetwork RPC call to un-fence the cluster network.
	UnFenceClusterNetwork(ctx context.Context, in *NetworkFenceRequest, opts ...grpc.CallOption) (*NetworkFenceResponse, error)
}

type networkFenceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkFenceClient(cc grpc.ClientConnInterface) NetworkFenceClient {
	return &networkFenceClient{cc}
}

func (c *networkFenceClient) FenceClusterNetwork(ctx context.Context, in *NetworkFenceRequest, opts ...grpc.CallOption) (*NetworkFenceResponse, error) {
	out := new(NetworkFenceResponse)
	err := c.cc.Invoke(ctx, "/proto.NetworkFence/FenceClusterNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkFenceClient) UnFenceClusterNetwork(ctx context.Context, in *NetworkFenceRequest, opts ...grpc.CallOption) (*NetworkFenceResponse, error) {
	out := new(NetworkFenceResponse)
	err := c.cc.Invoke(ctx, "/proto.NetworkFence/UnFenceClusterNetwork", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkFenceServer is the server API for NetworkFence service.
// All implementations must embed UnimplementedNetworkFenceServer
// for forward compatibility
type NetworkFenceServer interface {
	// FenceClusterNetwork RPC call to fence the cluster network.
	FenceClusterNetwork(context.Context, *NetworkFenceRequest) (*NetworkFenceResponse, error)
	// UnFenceClusterNetwork RPC call to un-fence the cluster network.
	UnFenceClusterNetwork(context.Context, *NetworkFenceRequest) (*NetworkFenceResponse, error)
	mustEmbedUnimplementedNetworkFenceServer()
}

// UnimplementedNetworkFenceServer must be embedded to have forward compatible implementations.
type UnimplementedNetworkFenceServer struct {
}

func (UnimplementedNetworkFenceServer) FenceClusterNetwork(context.Context, *NetworkFenceRequest) (*NetworkFenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FenceClusterNetwork not implemented")
}
func (UnimplementedNetworkFenceServer) UnFenceClusterNetwork(context.Context, *NetworkFenceRequest) (*NetworkFenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnFenceClusterNetwork not implemented")
}
func (UnimplementedNetworkFenceServer) mustEmbedUnimplementedNetworkFenceServer() {}

// UnsafeNetworkFenceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkFenceServer will
// result in compilation errors.
type UnsafeNetworkFenceServer interface {
	mustEmbedUnimplementedNetworkFenceServer()
}

func RegisterNetworkFenceServer(s grpc.ServiceRegistrar, srv NetworkFenceServer) {
	s.RegisterService(&NetworkFence_ServiceDesc, srv)
}

func _NetworkFence_FenceClusterNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkFenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkFenceServer).FenceClusterNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NetworkFence/FenceClusterNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkFenceServer).FenceClusterNetwork(ctx, req.(*NetworkFenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkFence_UnFenceClusterNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkFenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkFenceServer).UnFenceClusterNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NetworkFence/UnFenceClusterNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkFenceServer).UnFenceClusterNetwork(ctx, req.(*NetworkFenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkFence_ServiceDesc is the grpc.ServiceDesc for NetworkFence service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkFence_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.NetworkFence",
	HandlerType: (*NetworkFenceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FenceClusterNetwork",
			Handler:    _NetworkFence_FenceClusterNetwork_Handler,
		},
		{
			MethodName: "UnFenceClusterNetwork",
			Handler:    _NetworkFence_UnFenceClusterNetwork_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "networkfence.proto",
}
