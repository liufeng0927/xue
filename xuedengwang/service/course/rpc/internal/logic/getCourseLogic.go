package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/course/rpc/internal/svc"
	"xuedengwang/service/course/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCourseLogic {
	return &GetCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCourseLogic) GetCourse(in *pb.GetCourseReq) (*pb.GetCourseResp, error) {
	// todo: add your logic here and delete this line
	courseDB := l.svcCtx.Query.Course
	where := courseDB.WithContext(l.ctx)

	if len(in.SearchValue) > 0 {
		searchValue := fmt.Sprintf("%%%s%%", in.SearchValue)
		where = where.Where(courseDB.CourseName.Like(searchValue)).Or(courseDB.CourseNo.Like(searchValue))
	}
	offset := (in.PageIndex - 1) * in.PageSize
	courses, count, err := where.FindByPage(int(offset), int(in.PageSize))
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find course req: %+v", in)

	}
	var content []*pb.CourseDao
	for _, course := range courses {
		course := &pb.CourseDao{
			CreateTime: course.CreateTime.String(),
			CreateBy:   course.CreateBy,
			UpdateTime: course.CreateTime.String(),
			UpdateBy:   course.UpdateBy,
			Remarks:    course.Remarks,
			ID:         course.ID,
			Coursename: course.CourseName,
			Courseno:   course.CourseNo,
		}

		content = append(content, course)
	}

	return &pb.GetCourseResp{
		Content:       content,
		TotalElements: count,
	}, nil

}
