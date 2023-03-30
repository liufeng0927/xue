package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"xuedengwang/core/errx"
	"xuedengwang/dal/model"
	"xuedengwang/dal/query"
	"xuedengwang/service/score/rpc/internal/svc"
	"xuedengwang/service/score/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddScoresLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddScoresLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddScoresLogic {
	return &AddScoresLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddScoresLogic) AddScores(in *pb.AddScoresReq) (*pb.ScoresInterface, error) {
	// todo: add your logic here and delete this line

	studentDB := l.svcCtx.Query.Student
	studentScoreDB := l.svcCtx.Query.StudentScore
	var studentIDs []int64
	_ = studentDB.WithContext(l.ctx).Select(studentDB.ID).Where(studentDB.GradeClassID.Eq(in.GradeClassID)).Scan(&studentIDs)
	err := l.svcCtx.Query.Transaction(func(tx *query.Query) error {
		for _, studentID := range studentIDs {

			_, err := studentScoreDB.WithContext(l.ctx).Where(studentScoreDB.StudentID.Eq(studentID), studentScoreDB.GradeClassID.Eq(in.GradeClassID), studentScoreDB.CourseID.Eq(in.CourseID)).Take()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				score := &model.StudentScore{
					CreateBy:     in.ByID,
					UpdateBy:     in.ByID,
					GradeClassID: in.GradeClassID,
					CourseID:     in.CourseID,
					StudentID:    studentID,
					Type:         "未批改",
					Remarks:      "初试成绩",
				}
				if err := studentScoreDB.WithContext(l.ctx).Create(score); err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "add StudentScore err req: %+v", in)
	}
	return &pb.ScoresInterface{}, nil
}
