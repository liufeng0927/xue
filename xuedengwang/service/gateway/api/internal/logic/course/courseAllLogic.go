package course

import (
	"context"
	"xuedengwang/service/course/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CourseAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCourseAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CourseAllLogic {
	return &CourseAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CourseAllLogic) CourseAll() (*types.CourseAllResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.CourseRpc.CourseAll(l.ctx, &pb.CourseInterface{})
	if err != nil {
		return nil, err
	}
	var courseAll []*types.CourseAll
	for _, course := range rpcResp.CourseAll {
		courseAll = append(courseAll, &types.CourseAll{
			ID:   course.ID,
			Name: course.Name,
		})
	}
	return &types.CourseAllResp{
		CourseAll: courseAll,
	}, nil
}
