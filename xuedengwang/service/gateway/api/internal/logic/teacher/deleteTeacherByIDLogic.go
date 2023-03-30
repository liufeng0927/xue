package teacher

import (
	"context"
	"xuedengwang/service/teacher/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTeacherByIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteTeacherByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTeacherByIDLogic {
	return &DeleteTeacherByIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTeacherByIDLogic) DeleteTeacherByID(req *types.DeleteTeacherByIDReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.TeacherRpc.DeleteTeacherByID(l.ctx, &pb.DeleteTeacherByIDReq{TeacherID: req.ID})
	if err != nil {
		return nil, err
	}
	return
}
