package teacher

import (
	"context"
	"xuedengwang/core/ctxdata"
	"xuedengwang/service/teacher/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTeacherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTeacherLogic {
	return &UpdateTeacherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTeacherLogic) UpdateTeacher(req *types.UpdateTeacherReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.TeacherRpc.UpdateTeacher(l.ctx, &pb.UpdateTeacherReq{
		Teachno:    req.Teachno,
		Remarks:    req.Remarks,
		ID:         req.ID,
		Sex:        req.Sex,
		QQ:         req.QQ,
		Phone:      req.Phone,
		Name:       req.Name,
		CourseID:   req.CourseID,
		UpdateByID: ctxdata.GetUidFromCtx(l.ctx),
	})
	if err != nil {
		return nil, err
	}
	return
}
