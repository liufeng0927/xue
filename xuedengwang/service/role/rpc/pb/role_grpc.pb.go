// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: role.proto

package pb

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

// RoleClient is the client API for Role service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleClient interface {
	RoleAll(ctx context.Context, in *RoleInterface, opts ...grpc.CallOption) (*RoleAllResp, error)
	GetRole(ctx context.Context, in *GetRoleReq, opts ...grpc.CallOption) (*GetRoleResp, error)
	GetRoleByID(ctx context.Context, in *GetRoleByIDReq, opts ...grpc.CallOption) (*GetRoleByIDResp, error)
	UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*RoleInterface, error)
	AddRole(ctx context.Context, in *AddRoleReq, opts ...grpc.CallOption) (*RoleInterface, error)
	DeleteRoleByID(ctx context.Context, in *DeleteRoleByIDReq, opts ...grpc.CallOption) (*RoleInterface, error)
}

type roleClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleClient(cc grpc.ClientConnInterface) RoleClient {
	return &roleClient{cc}
}

func (c *roleClient) RoleAll(ctx context.Context, in *RoleInterface, opts ...grpc.CallOption) (*RoleAllResp, error) {
	out := new(RoleAllResp)
	err := c.cc.Invoke(ctx, "/pb.role/roleAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) GetRole(ctx context.Context, in *GetRoleReq, opts ...grpc.CallOption) (*GetRoleResp, error) {
	out := new(GetRoleResp)
	err := c.cc.Invoke(ctx, "/pb.role/getRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) GetRoleByID(ctx context.Context, in *GetRoleByIDReq, opts ...grpc.CallOption) (*GetRoleByIDResp, error) {
	out := new(GetRoleByIDResp)
	err := c.cc.Invoke(ctx, "/pb.role/getRoleByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) UpdateRole(ctx context.Context, in *UpdateRoleReq, opts ...grpc.CallOption) (*RoleInterface, error) {
	out := new(RoleInterface)
	err := c.cc.Invoke(ctx, "/pb.role/updateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) AddRole(ctx context.Context, in *AddRoleReq, opts ...grpc.CallOption) (*RoleInterface, error) {
	out := new(RoleInterface)
	err := c.cc.Invoke(ctx, "/pb.role/addRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) DeleteRoleByID(ctx context.Context, in *DeleteRoleByIDReq, opts ...grpc.CallOption) (*RoleInterface, error) {
	out := new(RoleInterface)
	err := c.cc.Invoke(ctx, "/pb.role/deleteRoleByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServer is the server API for Role service.
// All implementations must embed UnimplementedRoleServer
// for forward compatibility
type RoleServer interface {
	RoleAll(context.Context, *RoleInterface) (*RoleAllResp, error)
	GetRole(context.Context, *GetRoleReq) (*GetRoleResp, error)
	GetRoleByID(context.Context, *GetRoleByIDReq) (*GetRoleByIDResp, error)
	UpdateRole(context.Context, *UpdateRoleReq) (*RoleInterface, error)
	AddRole(context.Context, *AddRoleReq) (*RoleInterface, error)
	DeleteRoleByID(context.Context, *DeleteRoleByIDReq) (*RoleInterface, error)
	mustEmbedUnimplementedRoleServer()
}

// UnimplementedRoleServer must be embedded to have forward compatible implementations.
type UnimplementedRoleServer struct {
}

func (UnimplementedRoleServer) RoleAll(context.Context, *RoleInterface) (*RoleAllResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleAll not implemented")
}
func (UnimplementedRoleServer) GetRole(context.Context, *GetRoleReq) (*GetRoleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}
func (UnimplementedRoleServer) GetRoleByID(context.Context, *GetRoleByIDReq) (*GetRoleByIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoleByID not implemented")
}
func (UnimplementedRoleServer) UpdateRole(context.Context, *UpdateRoleReq) (*RoleInterface, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}
func (UnimplementedRoleServer) AddRole(context.Context, *AddRoleReq) (*RoleInterface, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRole not implemented")
}
func (UnimplementedRoleServer) DeleteRoleByID(context.Context, *DeleteRoleByIDReq) (*RoleInterface, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoleByID not implemented")
}
func (UnimplementedRoleServer) mustEmbedUnimplementedRoleServer() {}

// UnsafeRoleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleServer will
// result in compilation errors.
type UnsafeRoleServer interface {
	mustEmbedUnimplementedRoleServer()
}

func RegisterRoleServer(s grpc.ServiceRegistrar, srv RoleServer) {
	s.RegisterService(&Role_ServiceDesc, srv)
}

func _Role_RoleAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleInterface)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).RoleAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.role/roleAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).RoleAll(ctx, req.(*RoleInterface))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.role/getRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).GetRole(ctx, req.(*GetRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_GetRoleByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).GetRoleByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.role/getRoleByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).GetRoleByID(ctx, req.(*GetRoleByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.role/updateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).UpdateRole(ctx, req.(*UpdateRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_AddRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).AddRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.role/addRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).AddRole(ctx, req.(*AddRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_DeleteRoleByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).DeleteRoleByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.role/deleteRoleByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).DeleteRoleByID(ctx, req.(*DeleteRoleByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Role_ServiceDesc is the grpc.ServiceDesc for Role service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Role_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.role",
	HandlerType: (*RoleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "roleAll",
			Handler:    _Role_RoleAll_Handler,
		},
		{
			MethodName: "getRole",
			Handler:    _Role_GetRole_Handler,
		},
		{
			MethodName: "getRoleByID",
			Handler:    _Role_GetRoleByID_Handler,
		},
		{
			MethodName: "updateRole",
			Handler:    _Role_UpdateRole_Handler,
		},
		{
			MethodName: "addRole",
			Handler:    _Role_AddRole_Handler,
		},
		{
			MethodName: "deleteRoleByID",
			Handler:    _Role_DeleteRoleByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "role.proto",
}