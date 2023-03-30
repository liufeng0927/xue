package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/role/rpc/internal/svc"
	"xuedengwang/service/role/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRoleByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRoleByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRoleByIDLogic {
	return &DeleteRoleByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRoleByIDLogic) DeleteRoleByID(in *pb.DeleteRoleByIDReq) (*pb.RoleInterface, error) {
	// todo: add your logic here and delete this line
	roleDB := l.svcCtx.Query.Role
	_, err := roleDB.WithContext(l.ctx).Where(roleDB.ID.Eq(in.RoleID)).Delete()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "delete role err req: %+v", in)

	}
	return &pb.RoleInterface{}, nil
}
