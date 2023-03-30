package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"
	"xuedengwang/dal/model"

	"xuedengwang/service/gradeclass/rpc/internal/svc"
	"xuedengwang/service/gradeclass/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddGradeClassLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddGradeClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddGradeClassLogic {
	return &AddGradeClassLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddGradeClassLogic) AddGradeClass(in *pb.AddGradeClassReq) (*pb.GradeClassInterface, error) {
	// todo: add your logic here and delete this line
	gradeClass := &model.GradeClass{
		Remarks:  in.Remarks,
		CreateBy: in.ByID,
		UpdateBy: in.ByID,
		Name:     in.Name,
		Code:     in.Code,
		Clazz:    in.Clazz,
		Grade:    in.Grade,
	}
	if err := l.svcCtx.Query.GradeClass.WithContext(l.ctx).Create(gradeClass); err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "add gradeClass err req: %+v", in)
	}
	return &pb.GradeClassInterface{}, nil
}
