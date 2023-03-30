package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"
	"xuedengwang/dal/model"

	"xuedengwang/service/student/rpc/internal/svc"
	"xuedengwang/service/student/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStudentLogic {
	return &AddStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddStudentLogic) AddStudent(in *pb.AddStudentReq) (*pb.StudentInterface, error) {
	// todo: add your logic here and delete this line
	student := &model.Student{
		Remarks:      in.Remarks,
		Sex:          in.Sex,
		CreateBy:     in.ByID,
		UpdateBy:     in.ByID,
		Phone:        in.Phone,
		Qq:           in.QQ,
		Name:         in.Name,
		Stuno:        in.Stuno,
		GradeClassID: in.GradeClassID,
	}
	if err := l.svcCtx.Query.Student.WithContext(l.ctx).Create(student); err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "add student err req: %+v", in)
	}
	return &pb.StudentInterface{}, nil
}
