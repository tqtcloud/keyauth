// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: apps/policy/pb/policy.proto

package policy

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

// RPCClient is the client API for RPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RPCClient interface {
	ValidatePermission(ctx context.Context, in *ValidatePermissionRequest, opts ...grpc.CallOption) (*Policy, error)
	QueryPolicy(ctx context.Context, in *QueryPolicyRequest, opts ...grpc.CallOption) (*PolicySet, error)
}

type rPCClient struct {
	cc grpc.ClientConnInterface
}

func NewRPCClient(cc grpc.ClientConnInterface) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) ValidatePermission(ctx context.Context, in *ValidatePermissionRequest, opts ...grpc.CallOption) (*Policy, error) {
	out := new(Policy)
	err := c.cc.Invoke(ctx, "/keyauth.policy.RPC/ValidatePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) QueryPolicy(ctx context.Context, in *QueryPolicyRequest, opts ...grpc.CallOption) (*PolicySet, error) {
	out := new(PolicySet)
	err := c.cc.Invoke(ctx, "/keyauth.policy.RPC/QueryPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RPCServer is the server API for RPC service.
// All implementations must embed UnimplementedRPCServer
// for forward compatibility
type RPCServer interface {
	ValidatePermission(context.Context, *ValidatePermissionRequest) (*Policy, error)
	QueryPolicy(context.Context, *QueryPolicyRequest) (*PolicySet, error)
	mustEmbedUnimplementedRPCServer()
}

// UnimplementedRPCServer must be embedded to have forward compatible implementations.
type UnimplementedRPCServer struct {
}

func (UnimplementedRPCServer) ValidatePermission(context.Context, *ValidatePermissionRequest) (*Policy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatePermission not implemented")
}
func (UnimplementedRPCServer) QueryPolicy(context.Context, *QueryPolicyRequest) (*PolicySet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPolicy not implemented")
}
func (UnimplementedRPCServer) mustEmbedUnimplementedRPCServer() {}

// UnsafeRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RPCServer will
// result in compilation errors.
type UnsafeRPCServer interface {
	mustEmbedUnimplementedRPCServer()
}

func RegisterRPCServer(s grpc.ServiceRegistrar, srv RPCServer) {
	s.RegisterService(&RPC_ServiceDesc, srv)
}

func _RPC_ValidatePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidatePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).ValidatePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.policy.RPC/ValidatePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).ValidatePermission(ctx, req.(*ValidatePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RPC_QueryPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RPCServer).QueryPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keyauth.policy.RPC/QueryPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RPCServer).QueryPolicy(ctx, req.(*QueryPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RPC_ServiceDesc is the grpc.ServiceDesc for RPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "keyauth.policy.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidatePermission",
			Handler:    _RPC_ValidatePermission_Handler,
		},
		{
			MethodName: "QueryPolicy",
			Handler:    _RPC_QueryPolicy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/policy/pb/policy.proto",
}
