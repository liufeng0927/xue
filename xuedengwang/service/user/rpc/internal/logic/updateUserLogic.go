package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/auth"
	"xuedengwang/core/errx"
	"xuedengwang/service/user/rpc/internal/svc"
	"xuedengwang/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *pb.UpdateUserReq) (*pb.UserInterface, error) {
	// todo: add your logic here and delete this line
	userDB := l.svcCtx.Query.User
	pwd, _ := auth.HashAndSalt(in.Password)

	user, err := userDB.WithContext(l.ctx).Where(userDB.ID.Eq(in.UserID)).Take()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find user by id req: %+v", in)
	}
	{
		user.ID = in.UserID
		user.Remarks = in.Remarks
		user.Realname = in.Realname
		user.Sex = in.Sex
		user.Status = in.Status
		user.Email = in.Email
		user.Username = in.Username
		user.Password = pwd
		user.RoleID = in.RoleID
		user.UpdateBy = in.UpdateBy
		user.UpdateTime = nil
	}
	if _, err := userDB.WithContext(l.ctx).Updates(user); err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "update user err req: %+v", in)
	}

	return &pb.UserInterface{}, nil
}
