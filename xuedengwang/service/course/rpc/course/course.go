// Code generated by goctl. DO NOT EDIT!
// Source: course.proto

package course

import (
	"context"

	"xuedengwang/service/course/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCourseReq        = pb.AddCourseReq
	CourseAll           = pb.CourseAll
	CourseAllResp       = pb.CourseAllResp
	CourseDao           = pb.CourseDao
	CourseInterface     = pb.CourseInterface
	DeleteCourseByIDReq = pb.DeleteCourseByIDReq
	GetCourseByIDReq    = pb.GetCourseByIDReq
	GetCourseByIDResp   = pb.GetCourseByIDResp
	GetCourseReq        = pb.GetCourseReq
	GetCourseResp       = pb.GetCourseResp
	UpdateCourseReq     = pb.UpdateCourseReq

	Course interface {
		CourseAll(ctx context.Context, in *CourseInterface, opts ...grpc.CallOption) (*CourseAllResp, error)
		GetCourse(ctx context.Context, in *GetCourseReq, opts ...grpc.CallOption) (*GetCourseResp, error)
		GetCourseByID(ctx context.Context, in *GetCourseByIDReq, opts ...grpc.CallOption) (*GetCourseByIDResp, error)
		UpdateCourse(ctx context.Context, in *UpdateCourseReq, opts ...grpc.CallOption) (*CourseInterface, error)
		AddCourse(ctx context.Context, in *AddCourseReq, opts ...grpc.CallOption) (*CourseInterface, error)
		DeleteRCourseByID(ctx context.Context, in *DeleteCourseByIDReq, opts ...grpc.CallOption) (*CourseInterface, error)
	}

	defaultCourse struct {
		cli zrpc.Client
	}
)

func NewCourse(cli zrpc.Client) Course {
	return &defaultCourse{
		cli: cli,
	}
}

func (m *defaultCourse) CourseAll(ctx context.Context, in *CourseInterface, opts ...grpc.CallOption) (*CourseAllResp, error) {
	client := pb.NewCourseClient(m.cli.Conn())
	return client.CourseAll(ctx, in, opts...)
}

func (m *defaultCourse) GetCourse(ctx context.Context, in *GetCourseReq, opts ...grpc.CallOption) (*GetCourseResp, error) {
	client := pb.NewCourseClient(m.cli.Conn())
	return client.GetCourse(ctx, in, opts...)
}

func (m *defaultCourse) GetCourseByID(ctx context.Context, in *GetCourseByIDReq, opts ...grpc.CallOption) (*GetCourseByIDResp, error) {
	client := pb.NewCourseClient(m.cli.Conn())
	return client.GetCourseByID(ctx, in, opts...)
}

func (m *defaultCourse) UpdateCourse(ctx context.Context, in *UpdateCourseReq, opts ...grpc.CallOption) (*CourseInterface, error) {
	client := pb.NewCourseClient(m.cli.Conn())
	return client.UpdateCourse(ctx, in, opts...)
}

func (m *defaultCourse) AddCourse(ctx context.Context, in *AddCourseReq, opts ...grpc.CallOption) (*CourseInterface, error) {
	client := pb.NewCourseClient(m.cli.Conn())
	return client.AddCourse(ctx, in, opts...)
}

func (m *defaultCourse) DeleteRCourseByID(ctx context.Context, in *DeleteCourseByIDReq, opts ...grpc.CallOption) (*CourseInterface, error) {
	client := pb.NewCourseClient(m.cli.Conn())
	return client.DeleteRCourseByID(ctx, in, opts...)
}