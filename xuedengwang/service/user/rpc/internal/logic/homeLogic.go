package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"xuedengwang/core/errx"
	"xuedengwang/service/user/rpc/internal/svc"
	"xuedengwang/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeLogic {
	return &HomeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var ErrCountTeacherError = errx.NewErrMsg("统计教师人数错误")
var ErrCountStudentError = errx.NewErrMsg("统计学生人数错误")
var ErrCountClassError = errx.NewErrMsg("统计班级人数错误")
var ErrCountCourseError = errx.NewErrMsg("统计课程门数错误")
var ErrCountScoreError = errx.NewErrMsg("统计分数错误")
var ErrFindCourseNameError = errx.NewErrMsg("查询课程名错误")

// 统计课程门数
func (l *HomeLogic) Home(in *pb.UserInterface) (*pb.HomeResp, error) {
	// todo: add your logic here and delete this line

	teacherDB := l.svcCtx.Query.Teacher
	classDB := l.svcCtx.Query.GradeClass
	studentDB := l.svcCtx.Query.Student
	courseDB := l.svcCtx.Query.Course
	studentScoreDB := l.svcCtx.Query.StudentScore

	// 统计教师个数
	teacherNums, err := teacherDB.WithContext(l.ctx).Count()
	if err != nil {
		return nil, errors.Wrapf(ErrCountTeacherError, "count teacherNums err  , err:%v", err)
	}
	// 统计班级数量
	classNums, err := classDB.WithContext(l.ctx).Count()
	if err != nil {
		return nil, errors.Wrapf(ErrCountClassError, "count classNums err  , err:%v", err)
	}
	// 统计学生人数
	studentNums, err := studentDB.WithContext(l.ctx).Count()
	if err != nil {
		return nil, errors.Wrapf(ErrCountStudentError, "count studentNums err  , err:%v", err)
	}
	// 统计课程门数
	courseNums, err := courseDB.WithContext(l.ctx).Count()
	if err != nil {
		return nil, errors.Wrapf(ErrCountCourseError, "count CourseNums err  , err:%v", err)
	}

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
		Group(studentScoreDB.CourseID).Scan(&studentScores); err != nil {
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
	barEchartsSeriesList := []*pb.DateTypeName{
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

	return &pb.HomeResp{
		TeacherNums: teacherNums,
		ClassNums:   classNums,
		StudentNums: studentNums,
		CourseNums:  courseNums,
		Scores: &pb.Scores{
			CategoryList:         categoryList,
			BarEchartsSeriesList: barEchartsSeriesList,
		},
	}, nil
}
