package course

import (
	"context"
	"xuedengwang/core/ctxdata"
	"xuedengwang/service/course/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCourseLogic {
	return &AddCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCourseLogic) AddCourse(req *types.AddCourseReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.CourseRpc.AddCourse(l.ctx, &pb.AddCourseReq{
		Coursename: req.Coursename,
		Courseno:   req.Courseno,
		Remarks:    req.Remarks,
		ByID:       ctxdata.GetUidFromCtx(l.ctx),
	})
	if err != nil {
		return nil, err
	}
	return
}
