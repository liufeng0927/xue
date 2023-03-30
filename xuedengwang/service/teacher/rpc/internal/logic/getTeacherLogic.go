package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"
	"xuedengwang/service/teacher/rpc/internal/svc"
	"xuedengwang/service/teacher/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTeacherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTeacherLogic {
	return &GetTeacherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTeacherLogic) GetTeacher(in *pb.GetTeacherReq) (*pb.GetTeacherResp, error) {
	// todo: add your logic here and delete this line
	teacherDB := l.svcCtx.Query.Teacher
	courseDB := l.svcCtx.Query.Course.As("Course")
	where := teacherDB.WithContext(l.ctx).Preload(teacherDB.Course).Joins(teacherDB.Course)

	if len(in.SearchValue) > 0 {
		searchValue := fmt.Sprintf("%%%s%%", in.SearchValue)
		where = where.Where(teacherDB.TeachNo.Like(searchValue)).Or(teacherDB.Name.Like(searchValue)).Or(teacherDB.Qq.Like(searchValue)).Or(teacherDB.Phone.Like(searchValue)).Or(courseDB.CourseName.Like(searchValue))
	}
	offset := (in.PageIndex - 1) * in.PageSize
	teachers, count, err := where.FindByPage(int(offset), int(in.PageSize))
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find teacher req: %+v", in)

	}
	var content []*pb.TeacherDao
	for _, teacher := range teachers {
		resp := &pb.TeacherDao{
			ID:         teacher.ID,
			CreateTime: teacher.CreateTime.String(),
			UpdateTime: teacher.UpdateTime.String(),
			CourseName: teacher.Course.CourseName,
			CreateBy:   teacher.CreateBy,
			UpdateBy:   teacher.UpdateBy,
			Remarks:    teacher.Remarks,
			Teachno:    teacher.TeachNo,
			Sex:        teacher.Sex,
			Phone:      teacher.Phone,
			QQ:         teacher.Qq,
			Name:       teacher.Name,
		}
		content = append(content, resp)
	}

	return &pb.GetTeacherResp{
		Content:       content,
		TotalElements: count,
	}, nil

}
