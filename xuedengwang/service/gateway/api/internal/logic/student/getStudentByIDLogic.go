package student

import (
	"context"
	"github.com/jinzhu/copier"
	"xuedengwang/service/student/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentByIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStudentByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentByIDLogic {
	return &GetStudentByIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStudentByIDLogic) GetStudentByID(req *types.GetStudentByIDReq) (*types.GetStudentByIDResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.StudentRpc.GetStudentByID(l.ctx, &pb.GetStudentByIDReq{ID: req.ID})
	if err != nil {
		return nil, err
	}
	var resp types.GetStudentByIDResp
	_ = copier.Copy(&resp, rpcResp)
	resp.GradeClassID = rpcResp.GradeClassID
	return &resp, nil
}
