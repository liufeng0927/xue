// Code generated by goctl. DO NOT EDIT!
// Source: gradeclass.proto

package gradeclass

import (
	"context"

	"xuedengwang/service/gradeclass/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddGradeClassReq        = pb.AddGradeClassReq
	DeleteGradeClassByIDReq = pb.DeleteGradeClassByIDReq
	GetGradeClassByIDReq    = pb.GetGradeClassByIDReq
	GetGradeClassByIDResp   = pb.GetGradeClassByIDResp
	GetGradeClassReq        = pb.GetGradeClassReq
	GetGradeClassResp       = pb.GetGradeClassResp
	GradeClassDao           = pb.GradeClassDao
	GradeClassInterface     = pb.GradeClassInterface
	GradeclassAll           = pb.GradeclassAll
	GradeclassAllResp       = pb.GradeclassAllResp
	UpdateGradeClassReq     = pb.UpdateGradeClassReq

	Gradeclass interface {
		GradeclassAll(ctx context.Context, in *GradeClassInterface, opts ...grpc.CallOption) (*GradeclassAllResp, error)
		GetGradeClass(ctx context.Context, in *GetGradeClassReq, opts ...grpc.CallOption) (*GetGradeClassResp, error)
		GetGradeClassByID(ctx context.Context, in *GetGradeClassByIDReq, opts ...grpc.CallOption) (*GetGradeClassByIDResp, error)
		UpdateGradeClass(ctx context.Context, in *UpdateGradeClassReq, opts ...grpc.CallOption) (*GradeClassInterface, error)
		AddGradeClass(ctx context.Context, in *AddGradeClassReq, opts ...grpc.CallOption) (*GradeClassInterface, error)
		DeleteGradeClassByID(ctx context.Context, in *DeleteGradeClassByIDReq, opts ...grpc.CallOption) (*GradeClassInterface, error)
	}

	defaultGradeclass struct {
		cli zrpc.Client
	}
)

func NewGradeclass(cli zrpc.Client) Gradeclass {
	return &defaultGradeclass{
		cli: cli,
	}
}

func (m *defaultGradeclass) GradeclassAll(ctx context.Context, in *GradeClassInterface, opts ...grpc.CallOption) (*GradeclassAllResp, error) {
	client := pb.NewGradeclassClient(m.cli.Conn())
	return client.GradeclassAll(ctx, in, opts...)
}

func (m *defaultGradeclass) GetGradeClass(ctx context.Context, in *GetGradeClassReq, opts ...grpc.CallOption) (*GetGradeClassResp, error) {
	client := pb.NewGradeclassClient(m.cli.Conn())
	return client.GetGradeClass(ctx, in, opts...)
}

func (m *defaultGradeclass) GetGradeClassByID(ctx context.Context, in *GetGradeClassByIDReq, opts ...grpc.CallOption) (*GetGradeClassByIDResp, error) {
	client := pb.NewGradeclassClient(m.cli.Conn())
	return client.GetGradeClassByID(ctx, in, opts...)
}

func (m *defaultGradeclass) UpdateGradeClass(ctx context.Context, in *UpdateGradeClassReq, opts ...grpc.CallOption) (*GradeClassInterface, error) {
	client := pb.NewGradeclassClient(m.cli.Conn())
	return client.UpdateGradeClass(ctx, in, opts...)
}

func (m *defaultGradeclass) AddGradeClass(ctx context.Context, in *AddGradeClassReq, opts ...grpc.CallOption) (*GradeClassInterface, error) {
	client := pb.NewGradeclassClient(m.cli.Conn())
	return client.AddGradeClass(ctx, in, opts...)
}

func (m *defaultGradeclass) DeleteGradeClassByID(ctx context.Context, in *DeleteGradeClassByIDReq, opts ...grpc.CallOption) (*GradeClassInterface, error) {
	client := pb.NewGradeclassClient(m.cli.Conn())
	return client.DeleteGradeClassByID(ctx, in, opts...)
}
