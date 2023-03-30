package scores

import (
	"context"
	"xuedengwang/core/ctxdata"
	"xuedengwang/service/score/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddScoresLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddScoresLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddScoresLogic {
	return &AddScoresLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddScoresLogic) AddScores(req *types.AddScoresReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.ScoresRpc.AddScores(l.ctx, &pb.AddScoresReq{
		ByID:         ctxdata.GetUidFromCtx(l.ctx),
		CourseID:     req.CourseID,
		GradeClassID: req.GradeClassID,
	})
	if err != nil {
		return nil, err
	}
	return
}
