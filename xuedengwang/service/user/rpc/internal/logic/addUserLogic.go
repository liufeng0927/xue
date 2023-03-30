package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/auth"
	"xuedengwang/core/errx"
	"xuedengwang/dal/model"

	"xuedengwang/service/user/rpc/internal/svc"
	"xuedengwang/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserLogic) AddUser(in *pb.AddUserReq) (*pb.UserInterface, error) {
	// todo: add your logic here and delete this line
	pwd, _ := auth.HashAndSalt(in.Password)

	user := &model.User{
		Remarks:  in.Remarks,
		Realname: in.Realname,
		Sex:      in.Sex,
		Status:   in.Status,
		Email:    in.Email,
		Username: in.Username,
		Password: pwd,
		RoleID:   in.RoleID,
		CreateBy: in.ByID,
		UpdateBy: in.ByID,
	}
	if err := l.svcCtx.Query.User.WithContext(l.ctx).Create(user); err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "add user err req: %+v", in)
	}
	return &pb.UserInterface{}, nil
}
