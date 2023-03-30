package user

import (
	"context"
	"github.com/jinzhu/copier"
	"xuedengwang/service/user/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserReq) (*types.GetUserResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &pb.GetUserReq{
		SearchValue: req.SearchValue,
		Status:      req.Status,
		PageIndex:   req.PageIndex,
		PageSize:    req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	var users []*types.User
	_ = copier.Copy(&users, rpcResp.Content)

	return &types.GetUserResp{
		Content:       users,
		TotalElements: rpcResp.TotalElements,
	}, nil
}
