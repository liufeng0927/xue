package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/user/rpc/internal/svc"
	"xuedengwang/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIDLogic {
	return &GetUserByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIDLogic) GetUserByID(in *pb.GetUserByIDReq) (*pb.GetUserByIDResp, error) {
	// todo: add your logic here and delete this line
	userDB := l.svcCtx.Query.User
	user, err := userDB.WithContext(l.ctx).Preload(userDB.Role).Where(userDB.ID.Eq(in.ID)).First()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find user by id : %+v", in.ID)
	}
	resp := &pb.GetUserByIDResp{
		ID:         user.ID,
		CreateTime: user.CreateTime.String(),
		UpdateTime: user.UpdateTime.String(),
		SysRoleID:  user.RoleID,
	}
	_ = copier.Copy(resp, user)
	return resp, nil
}
