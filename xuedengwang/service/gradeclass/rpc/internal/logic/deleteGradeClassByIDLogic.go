package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/gradeclass/rpc/internal/svc"
	"xuedengwang/service/gradeclass/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGradeClassByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteGradeClassByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGradeClassByIDLogic {
	return &DeleteGradeClassByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteGradeClassByIDLogic) DeleteGradeClassByID(in *pb.DeleteGradeClassByIDReq) (*pb.GradeClassInterface, error) {
	// todo: add your logic here and delete this line
	gradeClassDB := l.svcCtx.Query.GradeClass
	_, err := gradeClassDB.WithContext(l.ctx).Where(gradeClassDB.ID.Eq(in.GradeClassID)).Delete()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "delete gradeClass err req: %+v", in)
	}
	return &pb.GradeClassInterface{}, nil
}
