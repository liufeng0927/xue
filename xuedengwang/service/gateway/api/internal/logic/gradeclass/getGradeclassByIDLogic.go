package gradeclass

import (
	"context"
	"github.com/jinzhu/copier"
	"xuedengwang/service/gradeclass/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGradeclassByIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGradeclassByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGradeclassByIDLogic {
	return &GetGradeclassByIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGradeclassByIDLogic) GetGradeclassByID(req *types.GetGradeClassByIDReq) (*types.GetGradeClassByIDResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.GradeClassRpc.GetGradeClassByID(l.ctx, &pb.GetGradeClassByIDReq{ID: req.ID})
	if err != nil {
		return nil, err
	}
	var resp types.GetGradeClassByIDResp
	_ = copier.Copy(&resp, rpcResp)
	return &resp, nil
}
