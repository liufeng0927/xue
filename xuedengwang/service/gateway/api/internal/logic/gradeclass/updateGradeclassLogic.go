package gradeclass

import (
	"context"
	"xuedengwang/core/ctxdata"
	"xuedengwang/service/gradeclass/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGradeclassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateGradeclassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGradeclassLogic {
	return &UpdateGradeclassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGradeclassLogic) UpdateGradeclass(req *types.UpdateGradeClassReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.GradeClassRpc.UpdateGradeClass(l.ctx, &pb.UpdateGradeClassReq{
		ID:         req.ID,
		Code:       req.Code,
		Name:       req.Name,
		Remarks:    req.Remarks,
		Clazz:      req.Clazz,
		Grade:      req.Grade,
		UpdateByID: ctxdata.GetUidFromCtx(l.ctx),
	})
	if err != nil {
		return nil, err
	}
	return
}
