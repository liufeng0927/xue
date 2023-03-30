package user

import (
	"context"
	"xuedengwang/core/ctxdata"
	"xuedengwang/service/user/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInfoLogic {
	return &UpdateInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateInfoLogic) UpdateInfo(req *types.UpdateInfoReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.UserRpc.UpdateInfo(l.ctx, &pb.UpdateInfoReq{
		Realname: req.Realname,
		Sex:      req.Sex,
		UserIcon: req.UserIcon,
		UserID:   ctxdata.GetUidFromCtx(l.ctx),
	})
	if err != nil {
		return nil, err
	}

	return
}
