// Code generated by goctl. DO NOT EDIT!
// Source: gradeclass.proto

package server

import (
	"context"

	"xuedengwang/service/gradeclass/rpc/internal/logic"
	"xuedengwang/service/gradeclass/rpc/internal/svc"
	"xuedengwang/service/gradeclass/rpc/pb"
)

type GradeclassServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedGradeclassServer
}

func NewGradeclassServer(svcCtx *svc.ServiceContext) *GradeclassServer {
	return &GradeclassServer{
		svcCtx: svcCtx,
	}
}

func (s *GradeclassServer) GradeclassAll(ctx context.Context, in *pb.GradeClassInterface) (*pb.GradeclassAllResp, error) {
	l := logic.NewGradeclassAllLogic(ctx, s.svcCtx)
	return l.GradeclassAll(in)
}

func (s *GradeclassServer) GetGradeClass(ctx context.Context, in *pb.GetGradeClassReq) (*pb.GetGradeClassResp, error) {
	l := logic.NewGetGradeClassLogic(ctx, s.svcCtx)
	return l.GetGradeClass(in)
}

func (s *GradeclassServer) GetGradeClassByID(ctx context.Context, in *pb.GetGradeClassByIDReq) (*pb.GetGradeClassByIDResp, error) {
	l := logic.NewGetGradeClassByIDLogic(ctx, s.svcCtx)
	return l.GetGradeClassByID(in)
}

func (s *GradeclassServer) UpdateGradeClass(ctx context.Context, in *pb.UpdateGradeClassReq) (*pb.GradeClassInterface, error) {
	l := logic.NewUpdateGradeClassLogic(ctx, s.svcCtx)
	return l.UpdateGradeClass(in)
}

func (s *GradeclassServer) AddGradeClass(ctx context.Context, in *pb.AddGradeClassReq) (*pb.GradeClassInterface, error) {
	l := logic.NewAddGradeClassLogic(ctx, s.svcCtx)
	return l.AddGradeClass(in)
}

func (s *GradeclassServer) DeleteGradeClassByID(ctx context.Context, in *pb.DeleteGradeClassByIDReq) (*pb.GradeClassInterface, error) {
	l := logic.NewDeleteGradeClassByIDLogic(ctx, s.svcCtx)
	return l.DeleteGradeClassByID(in)
}
