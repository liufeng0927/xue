package user

import (
	"context"
	"xuedengwang/core/ctxdata"
	"xuedengwang/service/user/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePwdLogic {
	return &UpdatePwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatePwdLogic) UpdatePwd(req *types.UpdatePwdReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line

	_, err = l.svcCtx.UserRpc.UpdatePwd(l.ctx, &pb.UpdatePwdReq{
		UserID:   ctxdata.GetUidFromCtx(l.ctx),
		NewPass:  req.NewPass,
		UsedPass: req.UsedPass,
	})

	if err != nil {
		return nil, err
	}

	return
}
