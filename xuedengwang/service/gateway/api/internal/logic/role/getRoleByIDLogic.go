package role

import (
	"context"
	"github.com/jinzhu/copier"
	"xuedengwang/service/role/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleByIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleByIDLogic {
	return &GetRoleByIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleByIDLogic) GetRoleByID(req *types.GetRoleByIDReq) (*types.GetRoleByIDResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.RoleRpc.GetRoleByID(l.ctx, &pb.GetRoleByIDReq{ID: req.ID})
	if err != nil {
		return nil, err
	}
	var resp types.GetRoleByIDResp
	_ = copier.Copy(&resp, rpcResp)
	return &resp, nil
}
