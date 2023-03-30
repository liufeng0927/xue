package teacher

import (
	"context"
	"github.com/jinzhu/copier"
	"xuedengwang/service/teacher/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTeacherByIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTeacherByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTeacherByIDLogic {
	return &GetTeacherByIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTeacherByIDLogic) GetTeacherByID(req *types.GetTeacherByIDReq) (*types.GetTeacherByIDResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.TeacherRpc.GetTeacherByID(l.ctx, &pb.GetTeacherByIDReq{ID: req.ID})
	if err != nil {
		return nil, err
	}
	var resp types.GetTeacherByIDResp
	_ = copier.Copy(&resp, rpcResp)
	resp.CourseID = rpcResp.CourseID
	return &resp, nil
}
