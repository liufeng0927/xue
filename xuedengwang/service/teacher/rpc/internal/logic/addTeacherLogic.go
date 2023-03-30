package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"
	"xuedengwang/dal/model"

	"xuedengwang/service/teacher/rpc/internal/svc"
	"xuedengwang/service/teacher/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTeacherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTeacherLogic {
	return &AddTeacherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddTeacherLogic) AddTeacher(in *pb.AddTeacherReq) (*pb.TeacherInterface, error) {
	// todo: add your logic here and delete this line
	teacher := &model.Teacher{
		Remarks:  in.Remarks,
		Sex:      in.Sex,
		CreateBy: in.ByID,
		UpdateBy: in.ByID,
		Phone:    in.Phone,
		Qq:       in.QQ,
		Name:     in.Name,
		TeachNo:  in.Teachno,
		CourseID: in.CourseID,
	}
	if err := l.svcCtx.Query.Teacher.WithContext(l.ctx).Create(teacher); err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "add teacher err req: %+v", in)
	}
	return &pb.TeacherInterface{}, nil
}
