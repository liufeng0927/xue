package gradeclass

import (
	"context"
	"xuedengwang/core/ctxdata"
	"xuedengwang/service/gradeclass/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGradeclassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddGradeclassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGradeclassLogic {
	return &AddGradeclassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddGradeclassLogic) AddGradeclass(req *types.AddGradeClassReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.GradeClassRpc.AddGradeClass(l.ctx, &pb.AddGradeClassReq{
		Code:    req.Code,
		Name:    req.Name,
		Remarks: req.Remarks,
		Clazz:   req.Clazz,
		Grade:   req.Grade,
		ByID:    ctxdata.GetUidFromCtx(l.ctx),
	})
	if err != nil {
		return nil, err
	}
	return
}
