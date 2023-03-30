package teacher

import (
	"context"
	"xuedengwang/core/ctxdata"
	"xuedengwang/service/teacher/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTeacherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTeacherLogic {
	return &AddTeacherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTeacherLogic) AddTeacher(req *types.AddTeacherReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.TeacherRpc.AddTeacher(l.ctx, &pb.AddTeacherReq{
		Teachno:  req.Teachno,
		Sex:      req.Sex,
		Phone:    req.Phone,
		QQ:       req.QQ,
		CourseID: req.CourseID,
		Name:     req.Name,
		Remarks:  req.Remarks,
		ByID:     ctxdata.GetUidFromCtx(l.ctx),
	})
	if err != nil {
		return nil, err
	}
	return
}
