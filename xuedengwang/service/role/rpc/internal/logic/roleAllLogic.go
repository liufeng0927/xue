package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/role/rpc/internal/svc"
	"xuedengwang/service/role/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAllLogic {
	return &RoleAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleAllLogic) RoleAll(in *pb.RoleInterface) (*pb.RoleAllResp, error) {
	// todo: add your logic here and delete this line
	roleDB := l.svcCtx.Query.Role

	roles, err := roleDB.WithContext(l.ctx).Find()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "role db find , error:%v", err)
	}

	var roleAll []*pb.RoleAll

	for _, role := range roles {
		role := &pb.RoleAll{
			ID:   role.ID,
			Name: role.Name,
		}
		roleAll = append(roleAll, role)

	}

	return &pb.RoleAllResp{RoleAll: roleAll}, nil
}
