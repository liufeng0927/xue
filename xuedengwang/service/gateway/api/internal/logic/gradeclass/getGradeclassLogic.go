package gradeclass

import (
	"context"
	"github.com/jinzhu/copier"
	"xuedengwang/service/gradeclass/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGradeclassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGradeclassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGradeclassLogic {
	return &GetGradeclassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGradeclassLogic) GetGradeclass(req *types.GetGradeClassReq) (*types.GetGradeClassResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.GradeClassRpc.GetGradeClass(l.ctx, &pb.GetGradeClassReq{
		PageIndex:   req.PageIndex,
		PageSize:    req.PageSize,
		SearchValue: req.SearchValue,
	})
	if err != nil {
		return nil, err
	}
	var gradeClasss []*types.GradeClass
	_ = copier.Copy(&gradeClasss, rpcResp.Content)

	return &types.GetGradeClassResp{
		Content:       gradeClasss,
		TotalElements: rpcResp.TotalElements,
	}, nil
}
