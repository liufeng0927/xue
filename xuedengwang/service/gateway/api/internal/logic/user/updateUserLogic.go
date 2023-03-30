package user

import (
	"context"
	"xuedengwang/core/ctxdata"
	"xuedengwang/service/user/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.UserRpc.UpdateUser(l.ctx, &pb.UpdateUserReq{
		UserID:   req.ID,
		Realname: req.Realname,
		Remarks:  req.Remarks,
		Username: req.Username,
		Password: req.Username,
		Sex:      req.Sex,
		Status:   req.Status,
		Email:    req.Email,
		RoleID:   req.SysRoleID,
		UpdateBy: ctxdata.GetUidFromCtx(l.ctx),
	})
	if err != nil {
		return nil, err
	}
	return
}
