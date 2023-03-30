package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"
	"xuedengwang/service/gradeclass/rpc/internal/svc"
	"xuedengwang/service/gradeclass/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGradeClassLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGradeClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGradeClassLogic {
	return &UpdateGradeClassLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGradeClassLogic) UpdateGradeClass(in *pb.UpdateGradeClassReq) (*pb.GradeClassInterface, error) {
	// todo: add your logic here and delete this line
	gradeClassDB := l.svcCtx.Query.GradeClass
	gradeClass, err := gradeClassDB.WithContext(l.ctx).Where(gradeClassDB.ID.Eq(in.ID)).Take()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find teacher by id req: %+v", in)
	}
	{
		gradeClass.ID = in.ID
		gradeClass.Name = in.Name
		gradeClass.Remarks = in.Remarks
		gradeClass.UpdateBy = in.UpdateByID
		gradeClass.UpdateTime = nil
		gradeClass.Grade = in.Grade
		gradeClass.Code = in.Code
		gradeClass.Clazz = in.Clazz
	}
	if _, err := gradeClassDB.WithContext(l.ctx).Updates(gradeClass); err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "update teacher err req: %+v", in)
	}
	return &pb.GradeClassInterface{}, nil
}
