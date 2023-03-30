package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/teacher/rpc/internal/svc"
	"xuedengwang/service/teacher/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTeacherByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTeacherByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTeacherByIDLogic {
	return &GetTeacherByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTeacherByIDLogic) GetTeacherByID(in *pb.GetTeacherByIDReq) (*pb.GetTeacherByIDResp, error) {
	// todo: add your logic here and delete this line
	teacherDB := l.svcCtx.Query.Teacher
	teacher, err := teacherDB.WithContext(l.ctx).Where(teacherDB.ID.Eq(in.ID)).First()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "find student by id : %+v", in.ID)
	}
	resp := &pb.GetTeacherByIDResp{
		ID:         teacher.ID,
		CreateTime: teacher.CreateTime.String(),
		UpdateTime: teacher.UpdateTime.String(),
		CourseID:   teacher.CourseID,
		CreateBy:   teacher.CreateBy,
		UpdateBy:   teacher.UpdateBy,
		Remarks:    teacher.Remarks,
		Teachno:    teacher.TeachNo,
		Sex:        teacher.Sex,
		Phone:      teacher.Phone,
		QQ:         teacher.Qq,
		Name:       teacher.Name,
	}
	return resp, nil
}
