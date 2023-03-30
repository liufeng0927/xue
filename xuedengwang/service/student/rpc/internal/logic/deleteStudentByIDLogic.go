package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/student/rpc/internal/svc"
	"xuedengwang/service/student/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStudentByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteStudentByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStudentByIDLogic {
	return &DeleteStudentByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteStudentByIDLogic) DeleteStudentByID(in *pb.DeleteStudentByIDReq) (*pb.StudentInterface, error) {
	// todo: add your logic here and delete this line
	studentDB := l.svcCtx.Query.Student
	_, err := studentDB.WithContext(l.ctx).Where(studentDB.ID.Eq(in.StudentID)).Delete()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "delete student err req: %+v", in)
	}
	return &pb.StudentInterface{}, nil
}
