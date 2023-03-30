package user

import (
	"context"
	"github.com/jinzhu/copier"
	"xuedengwang/service/user/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeLogic {
	return &HomeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeLogic) Home() (*types.HomeResp, error) {
	// todo: add your logic here and delete this line
	homeResp, err := l.svcCtx.UserRpc.Home(l.ctx, &pb.UserInterface{})
	if err != nil {
		return nil, err
	}
	var resp types.HomeResp
	_ = copier.Copy(&resp, homeResp)

	return &resp, nil
}
