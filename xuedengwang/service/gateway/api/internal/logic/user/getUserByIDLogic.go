package user

import (
	"context"
	"github.com/jinzhu/copier"
	"xuedengwang/service/user/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIDLogic {
	return &GetUserByIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserByIDLogic) GetUserByID(req *types.GetUserByIDReq) (*types.GetUserByIDResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.UserRpc.GetUserByID(l.ctx, &pb.GetUserByIDReq{
		ID: req.ID,
	})
	if err != nil {
		return nil, err
	}
	var resp types.GetUserByIDResp
	_ = copier.Copy(&resp, rpcResp)
	resp.SysRoleID = rpcResp.SysRoleID
	return &resp, nil
}
