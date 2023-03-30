package user

import (
	"context"
	"xuedengwang/core/ctxdata"
	"xuedengwang/service/user/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmailLogic {
	return &UpdateEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateEmailLogic) UpdateEmail(req *types.UpdateEmailReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.UserRpc.VerifyCode(l.ctx, &pb.VerifyCodeReq{
		Code:   req.Code,
		UserID: ctxdata.GetUidFromCtx(l.ctx),
	})

	if err != nil {
		return nil, err
	}

	_, err = l.svcCtx.UserRpc.UpdateEmail(l.ctx, &pb.UpdateEmailReq{
		Email:  req.Email,
		UserID: ctxdata.GetUidFromCtx(l.ctx),
	})
	if err != nil {
		return nil, err
	}
	return
}
