package course

import (
	"context"
	"xuedengwang/core/ctxdata"
	"xuedengwang/service/course/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCourseLogic {
	return &UpdateCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCourseLogic) UpdateCourse(req *types.UpdateCourseReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.CourseRpc.UpdateCourse(l.ctx, &pb.UpdateCourseReq{
		ID:         req.ID,
		Courseno:   req.Courseno,
		Coursename: req.Coursename,
		Remarks:    req.Remarks,
		UpdateByID: ctxdata.GetUidFromCtx(l.ctx),
	})
	if err != nil {
		return nil, err
	}
	return
}
