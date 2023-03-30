package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"
	"xuedengwang/dal/model"

	"xuedengwang/service/course/rpc/internal/svc"
	"xuedengwang/service/course/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCourseLogic {
	return &AddCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCourseLogic) AddCourse(in *pb.AddCourseReq) (*pb.CourseInterface, error) {
	// todo: add your logic here and delete this line
	course := &model.Course{
		Remarks:    in.Remarks,
		CourseName: in.Coursename,
		CourseNo:   in.Courseno,
		CreateBy:   in.ByID,
		UpdateBy:   in.ByID,
	}
	if err := l.svcCtx.Query.Course.WithContext(l.ctx).Create(course); err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "add course err req: %+v", in)
	}
	return &pb.CourseInterface{}, nil
}
