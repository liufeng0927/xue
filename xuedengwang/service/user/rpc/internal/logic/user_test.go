package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"xuedengwang/core/email"

	"xuedengwang/service/user/rpc/internal/config"
	"xuedengwang/service/user/rpc/internal/svc"

	"testing"
)

var svcCtx *svc.ServiceContext

func TestMain(t *testing.M) {
	var c config.Config
	conf.MustLoad("../../etc/user.yaml", &c)
	svcCtx = svc.NewServiceContext(c)
	t.Run()
}

func TestEnroll(t *testing.T) {
	ctx := context.Background()
	studentScoreDB := svcCtx.Query.StudentScore

	var studentScores []struct {
		Avg   float64
		Max   float64
		Min   float64
		Count int64
	}

	_ = studentScoreDB.WithContext(ctx).Select(
		studentScoreDB.Score.Avg().As("avg"),
		studentScoreDB.Score.Max().As("max"),
		studentScoreDB.Score.Min().As("min"),
		studentScoreDB.ID.Count().As("count")).
		Group(studentScoreDB.CourseID).Scan(&studentScores)

	fmt.Println(studentScores)
}
func TestEmail(t *testing.T) {
	client := email.NewClient("1945198227@qq.com", "", svcCtx.Config.Email)
	send, err := client.Send()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(send)
}
func TestAddScore(t *testing.T) {
	ctx := context.Background()
	studentDB := svcCtx.Query.Student
	var studentIDs []int64
	_ = studentDB.WithContext(ctx).Select(studentDB.ID).Where(studentDB.GradeClassID.Eq(1)).Scan(&studentIDs)
	fmt.Println(studentIDs)
}
