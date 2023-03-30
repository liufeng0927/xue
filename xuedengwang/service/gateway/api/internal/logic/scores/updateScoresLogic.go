package scores

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"xuedengwang/core/ctxdata"
	"xuedengwang/service/score/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateScoresLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateScoresLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateScoresLogic {
	return &UpdateScoresLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateScoresLogic) UpdateScores(req *types.UpdateScoresReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.ScoresRpc.UpdateScores(l.ctx, &pb.UpdateScoresReq{
		ID:         req.ID,
		UpdateByID: ctxdata.GetUidFromCtx(l.ctx),
		Score:      cast.ToFloat32(fmt.Sprintf("%.2f", cast.ToFloat32(req.Score))),
	})
	if err != nil {
		return nil, err
	}
	return
}
