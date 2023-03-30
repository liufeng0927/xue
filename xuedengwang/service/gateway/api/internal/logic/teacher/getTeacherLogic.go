package teacher

import (
	"context"
	"github.com/jinzhu/copier"
	"xuedengwang/service/teacher/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTeacherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTeacherLogic {
	return &GetTeacherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTeacherLogic) GetTeacher(req *types.GetTeacherReq) (resp *types.GetTeacherResp, err error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.TeacherRpc.GetTeacher(l.ctx, &pb.GetTeacherReq{
		SearchValue: req.SearchValue,
		PageIndex:   req.PageIndex,
		PageSize:    req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	var teachers []*types.Teacher
	_ = copier.Copy(&teachers, rpcResp.Content)

	return &types.GetTeacherResp{
		Content:       teachers,
		TotalElements: rpcResp.TotalElements,
	}, nil
}
