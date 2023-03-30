package student

import (
	"context"
	"xuedengwang/core/ctxdata"
	"xuedengwang/service/student/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStudentLogic {
	return &UpdateStudentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStudentLogic) UpdateStudent(req *types.UpdateStudentReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.StudentRpc.UpdateStudent(l.ctx, &pb.UpdateStudentReq{
		Stuno:        req.Stuno,
		Remarks:      req.Remarks,
		ID:           req.ID,
		Sex:          req.Sex,
		QQ:           req.QQ,
		Phone:        req.Phone,
		Name:         req.Name,
		GradeClassID: req.GradeClassId,
		UpdateByID:   ctxdata.GetUidFromCtx(l.ctx),
	})
	if err != nil {
		return nil, err
	}
	return
}
