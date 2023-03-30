package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/course/rpc/internal/svc"
	"xuedengwang/service/course/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCourseLogic {
	return &UpdateCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCourseLogic) UpdateCourse(in *pb.UpdateCourseReq) (*pb.CourseInterface, error) {
	// todo: add your logic here and delete this line
	courseDB := l.svcCtx.Query.Course

	course, err := courseDB.WithContext(l.ctx).Where(courseDB.ID.Eq(in.ID)).Take()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find course by id req: %+v", in)
	}
	{
		course.CourseName = in.Coursename
		course.Remarks = in.Remarks
		course.UpdateBy = in.UpdateByID
		course.CourseNo = in.Courseno
		course.UpdateTime = nil
	}
	if _, err := courseDB.WithContext(l.ctx).Updates(course); err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "update course err req: %+v", in)
	}
	return &pb.CourseInterface{}, nil
}
