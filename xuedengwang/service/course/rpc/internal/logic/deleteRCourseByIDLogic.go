package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/course/rpc/internal/svc"
	"xuedengwang/service/course/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRCourseByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRCourseByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRCourseByIDLogic {
	return &DeleteRCourseByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteRCourseByIDLogic) DeleteRCourseByID(in *pb.DeleteCourseByIDReq) (*pb.CourseInterface, error) {
	// todo: add your logic here and delete this line
	courseDB := l.svcCtx.Query.Course
	_, err := courseDB.WithContext(l.ctx).Where(courseDB.ID.Eq(in.CourseID)).Delete()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "delete course err req: %+v", in)
	}
	return &pb.CourseInterface{}, nil
}
