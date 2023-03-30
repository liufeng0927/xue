package scores

import (
	"context"
	"xuedengwang/service/score/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteScoresIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteScoresIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteScoresIDLogic {
	return &DeleteScoresIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteScoresIDLogic) DeleteScoresID(req *types.DeleteScoresByIDReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.ScoresRpc.DeleteScoresByID(l.ctx, &pb.DeleteScoresByIDReq{ScoresID: req.ID})
	if err != nil {
		return nil, err
	}

	return
}
