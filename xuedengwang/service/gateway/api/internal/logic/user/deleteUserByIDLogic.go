package user

import (
	"context"
	"xuedengwang/service/user/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserByIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserByIDLogic {
	return &DeleteUserByIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserByIDLogic) DeleteUserByID(req *types.DeleteUserByIDReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.UserRpc.DeleteUserByID(l.ctx, &pb.DeleteUserByIDReq{UserID: req.ID})
	if err != nil {
		return nil, err
	}
	return
}
