// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: student.proto

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

// StudentClient is the client API for Student service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentClient interface {
	GetStudent(ctx context.Context, in *GetStudentReq, opts ...grpc.CallOption) (*GetStudentResp, error)
	GetStudentByID(ctx context.Context, in *GetStudentByIDReq, opts ...grpc.CallOption) (*GetStudentByIDResp, error)
	UpdateStudent(ctx context.Context, in *UpdateStudentReq, opts ...grpc.CallOption) (*StudentInterface, error)
	AddStudent(ctx context.Context, in *AddStudentReq, opts ...grpc.CallOption) (*StudentInterface, error)
	DeleteStudentByID(ctx context.Context, in *DeleteStudentByIDReq, opts ...grpc.CallOption) (*StudentInterface, error)
}

type studentClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentClient(cc grpc.ClientConnInterface) StudentClient {
	return &studentClient{cc}
}

func (c *studentClient) GetStudent(ctx context.Context, in *GetStudentReq, opts ...grpc.CallOption) (*GetStudentResp, error) {
	out := new(GetStudentResp)
	err := c.cc.Invoke(ctx, "/pb.student/getStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentClient) GetStudentByID(ctx context.Context, in *GetStudentByIDReq, opts ...grpc.CallOption) (*GetStudentByIDResp, error) {
	out := new(GetStudentByIDResp)
	err := c.cc.Invoke(ctx, "/pb.student/getStudentByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentClient) UpdateStudent(ctx context.Context, in *UpdateStudentReq, opts ...grpc.CallOption) (*StudentInterface, error) {
	out := new(StudentInterface)
	err := c.cc.Invoke(ctx, "/pb.student/updateStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentClient) AddStudent(ctx context.Context, in *AddStudentReq, opts ...grpc.CallOption) (*StudentInterface, error) {
	out := new(StudentInterface)
	err := c.cc.Invoke(ctx, "/pb.student/addStudent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentClient) DeleteStudentByID(ctx context.Context, in *DeleteStudentByIDReq, opts ...grpc.CallOption) (*StudentInterface, error) {
	out := new(StudentInterface)
	err := c.cc.Invoke(ctx, "/pb.student/deleteStudentByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServer is the server API for Student service.
// All implementations must embed UnimplementedStudentServer
// for forward compatibility
type StudentServer interface {
	GetStudent(context.Context, *GetStudentReq) (*GetStudentResp, error)
	GetStudentByID(context.Context, *GetStudentByIDReq) (*GetStudentByIDResp, error)
	UpdateStudent(context.Context, *UpdateStudentReq) (*StudentInterface, error)
	AddStudent(context.Context, *AddStudentReq) (*StudentInterface, error)
	DeleteStudentByID(context.Context, *DeleteStudentByIDReq) (*StudentInterface, error)
	mustEmbedUnimplementedStudentServer()
}

// UnimplementedStudentServer must be embedded to have forward compatible implementations.
type UnimplementedStudentServer struct {
}

func (UnimplementedStudentServer) GetStudent(context.Context, *GetStudentReq) (*GetStudentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudent not implemented")
}
func (UnimplementedStudentServer) GetStudentByID(context.Context, *GetStudentByIDReq) (*GetStudentByIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentByID not implemented")
}
func (UnimplementedStudentServer) UpdateStudent(context.Context, *UpdateStudentReq) (*StudentInterface, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStudent not implemented")
}
func (UnimplementedStudentServer) AddStudent(context.Context, *AddStudentReq) (*StudentInterface, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStudent not implemented")
}
func (UnimplementedStudentServer) DeleteStudentByID(context.Context, *DeleteStudentByIDReq) (*StudentInterface, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStudentByID not implemented")
}
func (UnimplementedStudentServer) mustEmbedUnimplementedStudentServer() {}

// UnsafeStudentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentServer will
// result in compilation errors.
type UnsafeStudentServer interface {
	mustEmbedUnimplementedStudentServer()
}

func RegisterStudentServer(s grpc.ServiceRegistrar, srv StudentServer) {
	s.RegisterService(&Student_ServiceDesc, srv)
}

func _Student_GetStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServer).GetStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.student/getStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServer).GetStudent(ctx, req.(*GetStudentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Student_GetStudentByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServer).GetStudentByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.student/getStudentByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServer).GetStudentByID(ctx, req.(*GetStudentByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Student_UpdateStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStudentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServer).UpdateStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.student/updateStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServer).UpdateStudent(ctx, req.(*UpdateStudentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Student_AddStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStudentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServer).AddStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.student/addStudent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServer).AddStudent(ctx, req.(*AddStudentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Student_DeleteStudentByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStudentByIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServer).DeleteStudentByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.student/deleteStudentByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServer).DeleteStudentByID(ctx, req.(*DeleteStudentByIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Student_ServiceDesc is the grpc.ServiceDesc for Student service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Student_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.student",
	HandlerType: (*StudentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getStudent",
			Handler:    _Student_GetStudent_Handler,
		},
		{
			MethodName: "getStudentByID",
			Handler:    _Student_GetStudentByID_Handler,
		},
		{
			MethodName: "updateStudent",
			Handler:    _Student_UpdateStudent_Handler,
		},
		{
			MethodName: "addStudent",
			Handler:    _Student_AddStudent_Handler,
		},
		{
			MethodName: "deleteStudentByID",
			Handler:    _Student_DeleteStudentByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "student.proto",
}