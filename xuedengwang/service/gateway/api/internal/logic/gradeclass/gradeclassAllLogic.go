package gradeclass

import (
	"context"
	"xuedengwang/service/gradeclass/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GradeclassAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGradeclassAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GradeclassAllLogic {
	return &GradeclassAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GradeclassAllLogic) GradeclassAll() (*types.GradeClassAllResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.GradeClassRpc.GradeclassAll(l.ctx, &pb.GradeClassInterface{})
	if err != nil {
		return nil, err
	}
	var gradeClassAlls []*types.GradeClassAll
	for _, gradeClassAll := range rpcResp.GradeclassAll {
		gradeClassAlls = append(gradeClassAlls, &types.GradeClassAll{
			ID:   gradeClassAll.ID,
			Name: gradeClassAll.Name,
		})
	}
	return &types.GradeClassAllResp{
		GradeClassAll: gradeClassAlls,
	}, nil
}
