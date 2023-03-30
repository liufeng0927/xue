package role

import (
	"context"
	"xuedengwang/service/role/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleByIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRoleByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleByIDLogic {
	return &DeleteRoleByIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRoleByIDLogic) DeleteRoleByID(req *types.DeleteRoleByIDReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.RoleRpc.DeleteRoleByID(l.ctx, &pb.DeleteRoleByIDReq{RoleID: req.ID})
	if err != nil {
		return nil, err
	}
	return
}
