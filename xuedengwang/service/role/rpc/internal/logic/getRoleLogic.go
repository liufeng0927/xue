package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/role/rpc/internal/svc"
	"xuedengwang/service/role/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleLogic {
	return &GetRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleLogic) GetRole(in *pb.GetRoleReq) (*pb.GetRoleResp, error) {
	// todo: add your logic here and delete this line
	roleDB := l.svcCtx.Query.Role
	where := roleDB.WithContext(l.ctx)

	if len(in.SearchValue) > 0 {
		searchValue := fmt.Sprintf("%%%s%%", in.SearchValue)
		where.Where(roleDB.Name.Like(searchValue)).Or(roleDB.Code.Like(searchValue)).Or(roleDB.Remarks.Like(searchValue))
	}
	offset := (in.PageIndex - 1) * in.PageSize
	roles, count, err := where.FindByPage(int(offset), int(in.PageSize))
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find role req: %+v", in)

	}
	var content []*pb.RoleDao
	for _, role := range roles {
		role := &pb.RoleDao{
			CreateTime: role.CreateTime.String(),
			CreateBy:   role.CreateBy,
			UpdateTime: role.CreateTime.String(),
			UpdateBy:   role.UpdateBy,
			Remarks:    role.Remarks,
			ID:         role.ID,
			Name:       role.Name,
			Code:       role.Code,
		}

		content = append(content, role)
	}

	return &pb.GetRoleResp{
		Content:       content,
		TotalElements: count,
	}, nil

}
