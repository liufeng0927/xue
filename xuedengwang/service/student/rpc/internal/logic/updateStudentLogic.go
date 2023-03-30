package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/student/rpc/internal/svc"
	"xuedengwang/service/student/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStudentLogic {
	return &UpdateStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStudentLogic) UpdateStudent(in *pb.UpdateStudentReq) (*pb.StudentInterface, error) {
	// todo: add your logic here and delete this line
	studentDB := l.svcCtx.Query.Student

	student, err := studentDB.WithContext(l.ctx).Where(studentDB.ID.Eq(in.ID)).Take()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find student by id req: %+v", in)
	}
	{
		student.ID = in.ID
		student.Name = in.Name
		student.Remarks = in.Remarks
		student.UpdateBy = in.UpdateByID
		student.UpdateTime = nil
		student.Sex = in.Sex
		student.Phone = in.Phone
		student.Qq = in.QQ
		student.Stuno = in.Stuno
		student.GradeClassID = in.GradeClassID
	}
	if _, err := studentDB.WithContext(l.ctx).Updates(student); err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "update teacher err req: %+v", in)
	}
	return &pb.StudentInterface{}, nil
}
