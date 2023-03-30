// Code generated by goctl. DO NOT EDIT!
// Source: course.proto

package server

import (
	"context"

	"xuedengwang/service/course/rpc/internal/logic"
	"xuedengwang/service/course/rpc/internal/svc"
	"xuedengwang/service/course/rpc/pb"
)

type CourseServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedCourseServer
}

func NewCourseServer(svcCtx *svc.ServiceContext) *CourseServer {
	return &CourseServer{
		svcCtx: svcCtx,
	}
}

func (s *CourseServer) CourseAll(ctx context.Context, in *pb.CourseInterface) (*pb.CourseAllResp, error) {
	l := logic.NewCourseAllLogic(ctx, s.svcCtx)
	return l.CourseAll(in)
}

func (s *CourseServer) GetCourse(ctx context.Context, in *pb.GetCourseReq) (*pb.GetCourseResp, error) {
	l := logic.NewGetCourseLogic(ctx, s.svcCtx)
	return l.GetCourse(in)
}

func (s *CourseServer) GetCourseByID(ctx context.Context, in *pb.GetCourseByIDReq) (*pb.GetCourseByIDResp, error) {
	l := logic.NewGetCourseByIDLogic(ctx, s.svcCtx)
	return l.GetCourseByID(in)
}

func (s *CourseServer) UpdateCourse(ctx context.Context, in *pb.UpdateCourseReq) (*pb.CourseInterface, error) {
	l := logic.NewUpdateCourseLogic(ctx, s.svcCtx)
	return l.UpdateCourse(in)
}

func (s *CourseServer) AddCourse(ctx context.Context, in *pb.AddCourseReq) (*pb.CourseInterface, error) {
	l := logic.NewAddCourseLogic(ctx, s.svcCtx)
	return l.AddCourse(in)
}

func (s *CourseServer) DeleteRCourseByID(ctx context.Context, in *pb.DeleteCourseByIDReq) (*pb.CourseInterface, error) {
	l := logic.NewDeleteRCourseByIDLogic(ctx, s.svcCtx)
	return l.DeleteRCourseByID(in)
}
