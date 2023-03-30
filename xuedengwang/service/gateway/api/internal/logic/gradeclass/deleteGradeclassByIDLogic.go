package gradeclass

import (
	"context"
	"xuedengwang/service/gradeclass/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGradeclassByIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteGradeclassByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGradeclassByIDLogic {
	return &DeleteGradeclassByIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteGradeclassByIDLogic) DeleteGradeclassByID(req *types.DeleteGradeClassByIDReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.GradeClassRpc.DeleteGradeClassByID(l.ctx, &pb.DeleteGradeClassByIDReq{GradeClassID: req.ID})
	if err != nil {
		return nil, err
	}
	return
}
