package scores

import (
	"context"
	"github.com/jinzhu/copier"
	"xuedengwang/service/score/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScoresLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetScoresLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScoresLogic {
	return &GetScoresLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetScoresLogic) GetScores(req *types.GetScoresReq) (resp *types.GetScoresResp, err error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.ScoresRpc.GetScores(l.ctx, &pb.GetScoresReq{
		PageIndex:    req.PageIndex,
		PageSize:     req.PageSize,
		GradeClassId: req.GradeClassId,
		CourseId:     req.CourseId,
		Stuno:        req.Stuno,
		StudentName:  req.Name,
	})
	if err != nil {
		return nil, err
	}
	var scores []*types.Score
	_ = copier.Copy(&scores, rpcResp.Content)

	return &types.GetScoresResp{
		Content:       scores,
		TotalElements: rpcResp.TotalElements,
	}, nil
}
