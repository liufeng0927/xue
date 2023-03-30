package role

import (
	"context"
	"github.com/jinzhu/copier"
	"xuedengwang/service/role/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleLogic {
	return &GetRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleLogic) GetRole(req *types.GetRoleReq) (*types.GetRoleResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.RoleRpc.GetRole(l.ctx, &pb.GetRoleReq{
		PageIndex:   req.PageIndex,
		PageSize:    req.PageSize,
		SearchValue: req.SearchValue,
	})
	if err != nil {
		return nil, err
	}
	var roles []*types.Role
	_ = copier.Copy(&roles, rpcResp.Content)

	return &types.GetRoleResp{
		Content:       roles,
		TotalElements: rpcResp.TotalElements,
	}, nil

}
