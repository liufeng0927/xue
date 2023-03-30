package logic

import (
	"context"
	"xuedengwang/dal/model"

	"xuedengwang/service/user/rpc/internal/svc"
	"xuedengwang/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateInfoLogic {
	return &UpdateInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateInfoLogic) UpdateInfo(in *pb.UpdateInfoReq) (*pb.UserInterface, error) {
	// todo: add your logic here and delete this line
	userDB := l.svcCtx.Query.User
	user := &model.User{
		Realname: in.Realname,
		Sex:      in.Sex,
		UserIcon: &in.UserIcon,
	}
	_, err := userDB.WithContext(l.ctx).Where(userDB.ID.Eq(in.UserID)).Updates(user)
	if err != nil {
		return nil, err
	}

	return &pb.UserInterface{}, nil
}
