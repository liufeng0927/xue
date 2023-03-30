package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/score/rpc/internal/svc"
	"xuedengwang/service/score/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteScoresByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteScoresByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteScoresByIDLogic {
	return &DeleteScoresByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteScoresByIDLogic) DeleteScoresByID(in *pb.DeleteScoresByIDReq) (*pb.ScoresInterface, error) {
	// todo: add your logic here and delete this line
	studentScoreDB := l.svcCtx.Query.StudentScore
	_, err := studentScoreDB.WithContext(l.ctx).Where(studentScoreDB.ID.Eq(in.ScoresID)).Delete()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "delete studentScore err req: %+v", in)
	}
	return &pb.ScoresInterface{}, nil
}
