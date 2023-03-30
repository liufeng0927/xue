package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/score/rpc/internal/svc"
	"xuedengwang/service/score/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetScoresLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetScoresLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetScoresLogic {
	return &GetScoresLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetScoresLogic) GetScores(in *pb.GetScoresReq) (*pb.GetScoresResp, error) {
	// todo: add your logic here and delete this line
	studentScoreDB := l.svcCtx.Query.StudentScore
	studentDB := l.svcCtx.Query.Student.As("Student")

	where := studentScoreDB.WithContext(l.ctx).Preload(studentScoreDB.Student, studentScoreDB.Course, studentScoreDB.GradeClass).Joins(studentScoreDB.Student)

	if len(in.Stuno) > 0 {
		stuno := fmt.Sprintf("%%%s%%", in.Stuno)
		where = where.Where(studentDB.Stuno.Like(stuno))
	}

	if len(in.StudentName) > 0 {
		StudentName := fmt.Sprintf("%%%s%%", in.StudentName)
		where = where.Where(studentDB.Name.Like(StudentName))
	}
	if in.GradeClassId != 0 {
		where = where.Where(studentScoreDB.GradeClassID.Eq(in.GradeClassId))
	}
	if in.CourseId != 0 {
		where = where.Where(studentScoreDB.CourseID.Eq(in.CourseId))
	}

	offset := (in.PageIndex - 1) * in.PageSize
	studentScore, count, err := where.FindByPage(int(offset), int(in.PageSize))
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find studentScore req: %+v", in)

	}
	var content []*pb.ScoresDao
	for _, studentScore := range studentScore {

		resp := &pb.ScoresDao{
			ID:          studentScore.ID,
			CreateTime:  studentScore.CreateTime.String(),
			UpdateTime:  studentScore.UpdateTime.String(),
			Course:      studentScore.Course.CourseName,
			StudentName: studentScore.Student.Name,
			Stuno:       studentScore.Student.Stuno,
			GradeClass:  studentScore.GradeClass.Name,
			CreateBy:    studentScore.CreateBy,
			UpdateBy:    studentScore.UpdateBy,
			Remarks:     studentScore.Remarks,
			Score:       studentScore.Score,
			Type:        studentScore.Type,
		}
		content = append(content, resp)
	}

	return &pb.GetScoresResp{
		Content:       content,
		TotalElements: count,
	}, nil
}
