package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"
	"xuedengwang/service/role/rpc/internal/svc"
	"xuedengwang/service/role/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRoleLogic) UpdateRole(in *pb.UpdateRoleReq) (*pb.RoleInterface, error) {
	// todo: add your logic here and delete this line

	roleDB := l.svcCtx.Query.Role

	role, err := roleDB.WithContext(l.ctx).Where(roleDB.ID.Eq(in.RoleID)).Take()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find role by id req: %+v", in)
	}
	{
		role.Name = in.Name
		role.Remarks = in.Remarks
		role.UpdateBy = in.UpdateByID
		role.Code = in.Code
		role.UpdateTime = nil
	}
	if _, err := roleDB.WithContext(l.ctx).Updates(role); err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "update role err req: %+v", in)
	}
	return &pb.RoleInterface{}, nil
}
