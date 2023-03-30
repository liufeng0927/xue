package role

import (
	"context"
	"xuedengwang/service/role/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAllLogic {
	return &RoleAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleAllLogic) RoleAll() (*types.RoleAllResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.RoleRpc.RoleAll(l.ctx, &pb.RoleInterface{})
	if err != nil {
		return nil, err
	}
	var roleAll []*types.RoleAll
	for _, role := range rpcResp.RoleAll {
		roleAll = append(roleAll, &types.RoleAll{
			ID:   role.ID,
			Name: role.Name,
		})
	}
	return &types.RoleAllResp{
		RoleAll: roleAll,
	}, nil
}
