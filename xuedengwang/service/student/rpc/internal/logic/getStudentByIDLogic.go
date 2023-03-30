package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/student/rpc/internal/svc"
	"xuedengwang/service/student/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStudentByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentByIDLogic {
	return &GetStudentByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStudentByIDLogic) GetStudentByID(in *pb.GetStudentByIDReq) (*pb.GetStudentByIDResp, error) {
	// todo: add your logic here and delete this line
	studentDB := l.svcCtx.Query.Student
	student, err := studentDB.WithContext(l.ctx).Preload(studentDB.GradeClass).Where(studentDB.ID.Eq(in.ID)).First()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find student by id : %+v", in.ID)
	}
	resp := &pb.GetStudentByIDResp{
		ID:             student.ID,
		CreateTime:     student.CreateTime.String(),
		UpdateTime:     student.UpdateTime.String(),
		GradeClassID:   student.GradeClassID,
		GradeClassName: student.GradeClass.Name,
		CreateBy:       student.CreateBy,
		UpdateBy:       student.UpdateBy,
		Remarks:        student.Remarks,
		Stuno:          student.Stuno,
		Sex:            student.Sex,
		Phone:          student.Phone,
		QQ:             student.Qq,
		Name:           student.Name,
	}
	return resp, nil
}
