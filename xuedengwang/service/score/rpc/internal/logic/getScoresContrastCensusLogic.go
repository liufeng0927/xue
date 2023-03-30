package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"xuedengwang/core/errx"

	"xuedengwang/service/score/rpc/internal/svc"
	"xuedengwang/service/score/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScoresContrastCensusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrFindCourseNameError = errx.NewErrMsg("查询课程名错误")
var ErrCountScoreError = errx.NewErrMsg("统计分数错误")

func NewGetScoresContrastCensusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScoresContrastCensusLogic {
	return &GetScoresContrastCensusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetScoresContrastCensusLogic) GetScoresContrastCensus(in *pb.GetScoresContrastCensusReq) (*pb.GetScoresContrastCensusResp, error) {
	// todo: add your logic here and delete this line
	courseDB := l.svcCtx.Query.Course
	studentScoreDB := l.svcCtx.Query.StudentScore

	var studentScores []struct {
		CourseID int64
		Avg      float64
		Max      float64
		Min      float64
		Count    int64
	}

	if err := studentScoreDB.WithContext(l.ctx).
		Preload(studentScoreDB.Course).
		Select(studentScoreDB.CourseID.As("course_id"),
			studentScoreDB.Score.Avg().As("avg"),
			studentScoreDB.Score.Max().As("max"),
			studentScoreDB.Score.Min().As("min"),
			studentScoreDB.ID.Count().As("count")).
		Where(studentScoreDB.CourseID.Eq(in.CourseId)).
		Group(studentScoreDB.CourseID).
		Scan(&studentScores); err != nil {
		return nil, errors.Wrapf(ErrCountScoreError, "count studentScore err  , err:%v", err)

	}
	categoryList := make([]string, 0)

	avgDate := make([]float32, 0)
	maxDate := make([]float32, 0)
	minDate := make([]float32, 0)
	countDate := make([]float32, 0)

	for _, score := range studentScores {
		var courseName string
		if err := courseDB.WithContext(l.ctx).Select(courseDB.CourseName).Where(courseDB.ID.Eq(score.CourseID)).Scan(&courseName); err != nil {
			return nil, errors.Wrapf(ErrFindCourseNameError, "find  courseName  err  , err:%v", err)
		}
		categoryList = append(categoryList, courseName)
		scoreFormat := fmt.Sprintf("%.2f", score.Avg)
		avgDate = append(avgDate, cast.ToFloat32(scoreFormat))
		maxDate = append(maxDate, float32(score.Max))
		minDate = append(minDate, float32(score.Min))
		countDate = append(countDate, float32(score.Count))
	}
	census := []*pb.DateType{
		{
			Data: avgDate,
			Type: "bar",
			Name: "平均成绩",
		},
		{
			Data: maxDate,
			Type: "bar",
			Name: "最高成绩",
		},
		{
			Data: minDate,
			Type: "bar",
			Name: "最低成绩",
		},
		{
			Data: countDate,
			Type: "bar",
			Name: "总人数",
		},
	}
	return &pb.GetScoresContrastCensusResp{ScoresContrastCensus: &pb.ScoresContrastCensus{
		CategoryList:         categoryList,
		BarEchartsSeriesList: census,
	},
	}, nil
}
