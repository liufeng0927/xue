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

type UpdatePwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePwdLogic {
	return &UpdatePwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var ErrUserPwdError = errx.NewErrMsg("旧密码不正确")

func (l *UpdatePwdLogic) UpdatePwd(in *pb.UpdatePwdReq) (*pb.UserInterface, error) {
	// todo: add your logic here and delete this line

	userDB := l.svcCtx.Query.User

	user, err := userDB.WithContext(l.ctx).Where(userDB.ID.Eq(in.UserID)).First()
	if err != nil {
		return nil, errors.Wrapf(ErrUserNotFindError, "update pwd find user not find err userID:%v , err:%v", in.UserID, err)
	}
	if !auth.ComparePasswords(user.Password, in.UsedPass) {
		return nil, errors.Wrap(ErrUserPwdError, "update pwd error")
	}

	newPass, err := auth.HashAndSalt(in.NewPass)
	if err != nil {
		return nil, err
	}
	user.Password = newPass
	err = userDB.WithContext(l.ctx).Save(user)
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "req: %+v", in)
	}

	return &pb.UserInterface{}, nil
}
