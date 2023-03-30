package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/gradeclass/rpc/internal/svc"
	"xuedengwang/service/gradeclass/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGradeClassLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGradeClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGradeClassLogic {
	return &GetGradeClassLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGradeClassLogic) GetGradeClass(in *pb.GetGradeClassReq) (*pb.GetGradeClassResp, error) {
	// todo: add your logic here and delete this line
	gradeClassDB := l.svcCtx.Query.GradeClass
	studentDB := l.svcCtx.Query.Student
	where := gradeClassDB.WithContext(l.ctx)

	if len(in.SearchValue) > 0 {
		searchValue := fmt.Sprintf("%%%s%%", in.SearchValue)
		where = where.Where(gradeClassDB.Name.Like(searchValue))
	}
	offset := (in.PageIndex - 1) * in.PageSize
	gradeClasss, count, err := where.FindByPage(int(offset), int(in.PageSize))
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find gradeClass req: %+v", in)

	}
	var content []*pb.GradeClassDao
	for _, gradeClass := range gradeClasss {
		students, err := studentDB.WithContext(l.ctx).Where(studentDB.GradeClassID.Eq(gradeClass.ID)).Count()
		if err != nil {
			return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find gradeClass students by id : %+v", gradeClass.ID)
		}

		resp := &pb.GradeClassDao{
			ID:         gradeClass.ID,
			CreateTime: gradeClass.CreateTime.String(),
			UpdateTime: gradeClass.UpdateTime.String(),
			CreateBy:   gradeClass.CreateBy,
			UpdateBy:   gradeClass.UpdateBy,
			Remarks:    gradeClass.Remarks,
			Code:       gradeClass.Code,
			Grade:      gradeClass.Grade,
			Clazz:      gradeClass.Clazz,
			Name:       gradeClass.Name,
			Students:   students,
		}
		content = append(content, resp)
	}

	return &pb.GetGradeClassResp{
		Content:       content,
		TotalElements: count,
	}, nil
}
