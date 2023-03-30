package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"xuedengwang/core/errx"
	"xuedengwang/service/role/rpc/internal/svc"
	"xuedengwang/service/role/rpc/pb"
)

type GetRoleByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleByIDLogic {
	return &GetRoleByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleByIDLogic) GetRoleByID(in *pb.GetRoleByIDReq) (*pb.GetRoleByIDResp, error) {
	// todo: add your logic here and delete this line
	roleDB := l.svcCtx.Query.Role
	role, err := roleDB.WithContext(l.ctx).Where(roleDB.ID.Eq(in.ID)).First()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find role by id : %+v", in.ID)
	}
	resp := &pb.GetRoleByIDResp{
		CreateTime: role.CreateTime.String(),
		CreateBy:   role.CreateBy,
		UpdateTime: role.CreateTime.String(),
		UpdateBy:   role.UpdateBy,
		Remarks:    role.Remarks,
		ID:         role.ID,
		Name:       role.Name,
		Code:       role.Code,
	}
	return resp, nil
}
