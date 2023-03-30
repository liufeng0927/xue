package scores

import (
	"context"
	"github.com/jinzhu/copier"
	"xuedengwang/service/score/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScoresContrastCensusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetScoresContrastCensusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScoresContrastCensusLogic {
	return &GetScoresContrastCensusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetScoresContrastCensusLogic) GetScoresContrastCensus(req *types.GetScoresContrastCensusReq) (resp *types.GetScoresContrastCensusResp, err error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.ScoresRpc.GetScoresContrastCensus(l.ctx, &pb.GetScoresContrastCensusReq{
		CourseId: req.CourseId,
	})
	if err != nil {
		return nil, err
	}
	var census types.GetScoresContrastCensusResp
	_ = copier.Copy(&census.Scores, rpcResp.ScoresContrastCensus)

	return &census, nil

}
