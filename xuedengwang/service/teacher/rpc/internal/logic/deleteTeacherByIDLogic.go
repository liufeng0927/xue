package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/teacher/rpc/internal/svc"
	"xuedengwang/service/teacher/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTeacherByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTeacherByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTeacherByIDLogic {
	return &DeleteTeacherByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTeacherByIDLogic) DeleteTeacherByID(in *pb.DeleteTeacherByIDReq) (*pb.TeacherInterface, error) {
	// todo: add your logic here and delete this line
	teacherDB := l.svcCtx.Query.Teacher
	_, err := teacherDB.WithContext(l.ctx).Where(teacherDB.ID.Eq(in.TeacherID)).Delete()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "delete teacher err req: %+v", in)
	}
	return &pb.TeacherInterface{}, nil
}
