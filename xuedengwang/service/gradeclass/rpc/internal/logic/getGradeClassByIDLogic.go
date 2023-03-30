package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/gradeclass/rpc/internal/svc"
	"xuedengwang/service/gradeclass/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGradeClassByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGradeClassByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGradeClassByIDLogic {
	return &GetGradeClassByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGradeClassByIDLogic) GetGradeClassByID(in *pb.GetGradeClassByIDReq) (*pb.GetGradeClassByIDResp, error) {
	// todo: add your logic here and delete this line
	gradeClassDB := l.svcCtx.Query.GradeClass
	gradeClass, err := gradeClassDB.WithContext(l.ctx).Where(gradeClassDB.ID.Eq(in.ID)).First()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find gradeClass by id : %+v", in.ID)
	}
	studentDB := l.svcCtx.Query.Student
	count, err := studentDB.WithContext(l.ctx).Where(studentDB.GradeClassID.Eq(in.ID)).Count()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find gradeClass students by id : %+v", in.ID)
	}
	resp := &pb.GetGradeClassByIDResp{
		ID:         gradeClass.ID,
		CreateTime: gradeClass.CreateTime.String(),
		UpdateTime: gradeClass.UpdateTime.String(),
		CreateBy:   gradeClass.CreateBy,
		UpdateBy:   gradeClass.UpdateBy,
		Remarks:    gradeClass.Remarks,
		Code:       gradeClass.Code,
		Grade:      gradeClass.Grade,
		Clazz:      gradeClass.Clazz,
		Name:       gradeClass.Name,
		Students:   count,
	}
	return resp, nil
}
