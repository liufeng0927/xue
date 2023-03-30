package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"
	"xuedengwang/dal/model"

	"xuedengwang/service/role/rpc/internal/svc"
	"xuedengwang/service/role/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRoleLogic {
	return &AddRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddRoleLogic) AddRole(in *pb.AddRoleReq) (*pb.RoleInterface, error) {
	// todo: add your logic here and delete this line
	role := &model.Role{
		Remarks:  in.Remarks,
		Code:     in.Code,
		Name:     in.Name,
		CreateBy: in.ByID,
		UpdateBy: in.ByID,
	}
	if err := l.svcCtx.Query.Role.WithContext(l.ctx).Create(role); err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "add role err req: %+v", in)
	}
	return &pb.RoleInterface{}, nil
}
