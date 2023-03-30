package course

import (
	"context"
	"github.com/jinzhu/copier"
	"xuedengwang/service/course/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCourseLogic {
	return &GetCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCourseLogic) GetCourse(req *types.GetCourseReq) (*types.GetCourseResp, error) {
	// todo: add your logic here and delete this line

	rpcResp, err := l.svcCtx.CourseRpc.GetCourse(l.ctx, &pb.GetCourseReq{
		PageIndex:   req.PageIndex,
		PageSize:    req.PageSize,
		SearchValue: req.SearchValue,
	})
	if err != nil {
		return nil, err
	}
	var courses []*types.Course
	_ = copier.Copy(&courses, rpcResp.Content)

	return &types.GetCourseResp{
		Content:       courses,
		TotalElements: rpcResp.TotalElements,
	}, nil
}
