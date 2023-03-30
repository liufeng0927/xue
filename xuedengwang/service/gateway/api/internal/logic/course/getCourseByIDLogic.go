package course

import (
	"context"
	"github.com/jinzhu/copier"
	"xuedengwang/service/course/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCourseByIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCourseByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCourseByIDLogic {
	return &GetCourseByIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCourseByIDLogic) GetCourseByID(req *types.GetCourseByIDReq) (*types.GetCourseByIDResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.CourseRpc.GetCourseByID(l.ctx, &pb.GetCourseByIDReq{ID: req.ID})
	if err != nil {
		return nil, err
	}
	var resp types.GetCourseByIDResp
	_ = copier.Copy(&resp, rpcResp)
	return &resp, nil
}
