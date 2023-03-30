package student

import (
	"context"
	"github.com/jinzhu/copier"
	"xuedengwang/service/student/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentLogic {
	return &GetStudentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStudentLogic) GetStudent(req *types.GetStudentReq) (resp *types.GetStudentResp, err error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.StudentRpc.GetStudent(l.ctx, &pb.GetStudentReq{
		SearchValue: req.SearchValue,
		PageIndex:   req.PageIndex,
		PageSize:    req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	var students []*types.Student
	_ = copier.Copy(&students, rpcResp.Content)

	return &types.GetStudentResp{
		Content:       students,
		TotalElements: rpcResp.TotalElements,
	}, nil

}
