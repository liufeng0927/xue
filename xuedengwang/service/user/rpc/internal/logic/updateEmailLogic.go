package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/user/rpc/internal/svc"
	"xuedengwang/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrUpdateEmailError = errx.NewErrMsg("修改用户邮箱错误！")

func NewUpdateEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmailLogic {
	return &UpdateEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateEmailLogic) UpdateEmail(in *pb.UpdateEmailReq) (*pb.UserInterface, error) {
	// todo: add your logic here and delete this line
	userDB := l.svcCtx.Query.User
	_, err := userDB.WithContext(l.ctx).Where(userDB.ID.Eq(in.UserID)).Update(userDB.Email, in.Email)
	if err != nil {
		return nil, errors.Wrapf(ErrUpdateEmailError, "update email by userid userID:%v , email:%s, err:%v", in.UserID, in.Email, err)
	}

	return &pb.UserInterface{}, nil
}
