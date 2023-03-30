package scores

import (
	"context"
	"github.com/jinzhu/copier"
	"xuedengwang/service/score/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScoreCensusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetScoreCensusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScoreCensusLogic {
	return &GetScoreCensusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetScoreCensusLogic) GetScoreCensus(req *types.GetScoreCensusReq) (*types.GetScoreCensusResp, error) {
	// todo: add your logic here and delete this line
	rpcResp, err := l.svcCtx.ScoresRpc.GetScoreCensus(l.ctx, &pb.GetScoreCensusReq{
		GradeClassId: req.GradeClassId,
		CourseId:     req.CourseId,
	})
	if err != nil {
		return nil, err
	}
	var census types.GetScoreCensusResp
	_ = copier.Copy(&census, rpcResp)

	return &census, nil
}
