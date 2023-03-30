package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"
	"xuedengwang/service/teacher/rpc/internal/svc"
	"xuedengwang/service/teacher/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTeacherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTeacherLogic {
	return &UpdateTeacherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTeacherLogic) UpdateTeacher(in *pb.UpdateTeacherReq) (*pb.TeacherInterface, error) {
	// todo: add your logic here and delete this line
	teacherDB := l.svcCtx.Query.Teacher

	teacher, err := teacherDB.WithContext(l.ctx).Where(teacherDB.ID.Eq(in.ID)).Take()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find teacher by id req: %+v", in)
	}
	{
		teacher.ID = in.ID
		teacher.Name = in.Name
		teacher.Remarks = in.Remarks
		teacher.UpdateBy = in.UpdateByID
		teacher.UpdateTime = nil
		teacher.Sex = in.Sex
		teacher.Phone = in.Phone
		teacher.Qq = in.QQ
		teacher.TeachNo = in.Teachno
		teacher.CourseID = in.CourseID
	}
	if _, err := teacherDB.WithContext(l.ctx).Updates(teacher); err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "update teacher err req: %+v", in)
	}
	return &pb.TeacherInterface{}, nil
}
