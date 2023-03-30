package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/score/rpc/internal/svc"
	"xuedengwang/service/score/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateScoresLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrScoreError = errx.NewErrMsg("请输入0~100之间的分数")

func NewUpdateScoresLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateScoresLogic {
	return &UpdateScoresLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateScoresLogic) UpdateScores(in *pb.UpdateScoresReq) (*pb.ScoresInterface, error) {
	// todo: add your logic here and delete this line
	studentScoreDB := l.svcCtx.Query.StudentScore

	if in.Score > 100 || in.Score < 0 {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find teacher by id req: %+v", in)
	}
	studentScore, err := studentScoreDB.WithContext(l.ctx).Where(studentScoreDB.ID.Eq(in.ID)).Take()
	if err != nil {
		return nil, errors.Wrapf(ErrScoreError, "Please enter a score between 0~100 score: %+v", in.Score)
	}
	{
		studentScore.ID = in.ID
		studentScore.UpdateBy = in.UpdateByID
		studentScore.UpdateTime = nil
		studentScore.Score = in.Score
		studentScore.Type = "已批改"
	}
	if _, err := studentScoreDB.WithContext(l.ctx).Updates(studentScore); err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "update studentScore err req: %+v", in)
	}
	return &pb.ScoresInterface{}, nil
}
