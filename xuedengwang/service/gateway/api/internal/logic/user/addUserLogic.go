package user

import (
	"context"
	"xuedengwang/core/ctxdata"
	"xuedengwang/service/user/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.AddUserReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.UserRpc.AddUser(l.ctx, &pb.AddUserReq{
		Realname: req.Realname,
		Remarks:  req.Remarks,
		Username: req.Username,
		Password: req.Username,
		Sex:      req.Sex,
		Status:   req.Status,
		Email:    req.Email,
		RoleID:   req.SysRoleID,
		ByID:     ctxdata.GetUidFromCtx(l.ctx),
	})

	if err != nil {
		return nil, err
	}
	return
}
