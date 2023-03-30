package course

import (
	"context"
	"xuedengwang/service/course/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCourseByIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteCourseByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCourseByIDLogic {
	return &DeleteCourseByIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCourseByIDLogic) DeleteCourseByID(req *types.DeleteCourseByIDReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.CourseRpc.DeleteRCourseByID(l.ctx, &pb.DeleteCourseByIDReq{CourseID: req.ID})
	if err != nil {
		return nil, err
	}
	return
}
