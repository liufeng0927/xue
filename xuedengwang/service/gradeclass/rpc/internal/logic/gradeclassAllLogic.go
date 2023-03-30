package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/gradeclass/rpc/internal/svc"
	"xuedengwang/service/gradeclass/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GradeclassAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGradeclassAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GradeclassAllLogic {
	return &GradeclassAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GradeclassAllLogic) GradeclassAll(in *pb.GradeClassInterface) (*pb.GradeclassAllResp, error) {
	// todo: add your logic here and delete this line
	gradeClassDB := l.svcCtx.Query.GradeClass

	gradeClasss, err := gradeClassDB.WithContext(l.ctx).Find()
	if err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "gradeClass db find , error:%v", err)
	}

	var gradeClassAll []*pb.GradeclassAll

	for _, gradeClass := range gradeClasss {
		gradeClass := &pb.GradeclassAll{
			ID:   gradeClass.ID,
			Name: gradeClass.Name,
		}
		gradeClassAll = append(gradeClassAll, gradeClass)

	}
	return &pb.GradeclassAllResp{GradeclassAll: gradeClassAll}, nil
}
