package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/course/rpc/internal/svc"
	"xuedengwang/service/course/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CourseAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCourseAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CourseAllLogic {
	return &CourseAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CourseAllLogic) CourseAll(in *pb.CourseInterface) (*pb.CourseAllResp, error) {
	// todo: add your logic here and delete this line

	courseDB := l.svcCtx.Query.Course

	courses, err := courseDB.WithContext(l.ctx).Find()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "role db find , error:%v", err)
	}

	var courseAll []*pb.CourseAll

	for _, course := range courses {
		role := &pb.CourseAll{
			Name: course.CourseName,
			ID:   course.ID,
		}
		courseAll = append(courseAll, role)

	}

	return &pb.CourseAllResp{CourseAll: courseAll}, nil
}
