package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/user/rpc/internal/svc"
	"xuedengwang/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserByIDLogic {
	return &DeleteUserByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserByIDLogic) DeleteUserByID(in *pb.DeleteUserByIDReq) (*pb.UserInterface, error) {
	// todo: add your logic here and delete this line
	userDB := l.svcCtx.Query.User
	_, err := userDB.WithContext(l.ctx).Where(userDB.ID.Eq(in.UserID)).Delete()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "delete user err req: %+v", in)

	}
	return &pb.UserInterface{}, nil
}
