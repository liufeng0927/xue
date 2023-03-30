package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"
	globalkey "xuedengwang/core/global"

	"xuedengwang/service/score/rpc/internal/svc"
	"xuedengwang/service/score/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScoreCensusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetScoreCensusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScoreCensusLogic {
	return &GetScoreCensusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetScoreCensusLogic) GetScoreCensus(in *pb.GetScoreCensusReq) (*pb.GetScoreCensusResp, error) {
	// todo: add your logic here and delete this line
	scoreDB := l.svcCtx.Query.StudentScore
	var scores []int64
	where := scoreDB.WithContext(l.ctx).Select(scoreDB.Score)

	if in.CourseId != 0 {
		where = where.Where(scoreDB.CourseID.Eq(in.CourseId))
	}

	if in.GradeClassId != 0 {
		where = where.Where(scoreDB.GradeClassID.Eq(in.GradeClassId))
	}

	if err := where.Scan(&scores); err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find studentScore req: %+v", in)
	}

	var scoreScale [5]int64

	for _, score := range scores {
		if score >= 90 {
			scoreScale[0] = scoreScale[0] + 1
		} else if score >= 80 {
			scoreScale[1] = scoreScale[1] + 1
		} else if score >= 70 {
			scoreScale[2] = scoreScale[2] + 1
		} else if score > 60 {
			scoreScale[3] = scoreScale[3] + 1
		} else {
			scoreScale[4] = scoreScale[4] + 1
		}
	}

	var resp []*pb.Census
	if scoreScale[0] > 0 {
		resp = append(resp, &pb.Census{
			Name:  globalkey.ScoreScale_Excellence,
			Value: scoreScale[0],
		})
	}
	if scoreScale[1] > 0 {
		resp = append(resp, &pb.Census{
			Name:  globalkey.ScoreScale_GoodMarks,
			Value: scoreScale[1],
		})
	}
	if scoreScale[2] > 0 {
		resp = append(resp, &pb.Census{
			Name:  globalkey.ScoreScale_Pass,
			Value: scoreScale[2],
		})
	}
	if scoreScale[3] > 0 {
		resp = append(resp, &pb.Census{
			Name:  globalkey.ScoreScale_medium,
			Value: scoreScale[3],
		})
	}
	if scoreScale[4] > 0 {
		resp = append(resp, &pb.Census{
			Name:  globalkey.ScoreScale_Fail,
			Value: scoreScale[4],
		})
	}

	return &pb.GetScoreCensusResp{
		Census: resp,
	}, nil
}
