package role

import (
	"context"
	"xuedengwang/core/ctxdata"
	"xuedengwang/service/role/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleLogic) UpdateRole(req *types.UpdateRoleReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.RoleRpc.UpdateRole(l.ctx, &pb.UpdateRoleReq{
		RoleID:     req.ID,
		Code:       req.Code,
		Name:       req.Name,
		Remarks:    req.Remarks,
		UpdateByID: ctxdata.GetUidFromCtx(l.ctx),
	})
	if err != nil {
		return nil, err
	}
	return
}
