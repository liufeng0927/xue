package student

import (
	"context"
	"xuedengwang/service/student/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStudentByIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteStudentByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStudentByIDLogic {
	return &DeleteStudentByIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteStudentByIDLogic) DeleteStudentByID(req *types.DeleteStudentByIDReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.StudentRpc.DeleteStudentByID(l.ctx, &pb.DeleteStudentByIDReq{StudentID: req.ID})
	if err != nil {
		return nil, err
	}
	return
}
