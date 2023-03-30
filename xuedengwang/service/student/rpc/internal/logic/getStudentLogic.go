package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"
	"xuedengwang/service/student/rpc/internal/svc"
	"xuedengwang/service/student/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentLogic {
	return &GetStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStudentLogic) GetStudent(in *pb.GetStudentReq) (*pb.GetStudentResp, error) {
	// todo: add your logic here and delete this line
	studentDB := l.svcCtx.Query.Student
	gradeClassDB := l.svcCtx.Query.GradeClass.As("GradeClass")
	where := studentDB.WithContext(l.ctx).Preload(studentDB.GradeClass).Joins(studentDB.GradeClass)

	if len(in.SearchValue) > 0 {
		searchValue := fmt.Sprintf("%%%s%%", in.SearchValue)
		where = where.Where(studentDB.Stuno.Like(searchValue)).Or(studentDB.Name.Like(searchValue)).Or(studentDB.Qq.Like(searchValue)).Or(studentDB.Phone.Like(searchValue)).Or(gradeClassDB.Name.Like(searchValue))
	}
	offset := (in.PageIndex - 1) * in.PageSize
	students, count, err := where.FindByPage(int(offset), int(in.PageSize))
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find student req: %+v", in)

	}
	var content []*pb.StudentDao
	for _, student := range students {
		resp := &pb.StudentDao{
			ID:             student.ID,
			CreateTime:     student.CreateTime.String(),
			UpdateTime:     student.UpdateTime.String(),
			GradeClassName: student.GradeClass.Name,
			CreateBy:       student.CreateBy,
			UpdateBy:       student.UpdateBy,
			Remarks:        student.Remarks,
			Stuno:          student.Stuno,
			Sex:            student.Sex,
			Phone:          student.Phone,
			QQ:             student.Qq,
			Name:           student.Name,
		}
		content = append(content, resp)
	}

	return &pb.GetStudentResp{
		Content:       content,
		TotalElements: count,
	}, nil
}
