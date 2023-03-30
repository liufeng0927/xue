package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/course/rpc/internal/svc"
	"xuedengwang/service/course/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCourseByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCourseByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCourseByIDLogic {
	return &GetCourseByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCourseByIDLogic) GetCourseByID(in *pb.GetCourseByIDReq) (*pb.GetCourseByIDResp, error) {
	// todo: add your logic here and delete this line
	courseDB := l.svcCtx.Query.Course
	course, err := courseDB.WithContext(l.ctx).Where(courseDB.ID.Eq(in.ID)).First()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find course by id : %+v", in.ID)
	}
	resp := &pb.GetCourseByIDResp{
		CreateTime: course.CreateTime.String(),
		CreateBy:   course.CreateBy,
		UpdateTime: course.CreateTime.String(),
		UpdateBy:   course.UpdateBy,
		Remarks:    course.Remarks,
		ID:         course.ID,
		Coursename: course.CourseName,
		Courseno:   course.CourseNo,
	}
	return resp, nil
}
