package student

import (
	"context"
	"xuedengwang/core/ctxdata"
	"xuedengwang/service/student/rpc/pb"

	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStudentLogic {
	return &AddStudentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddStudentLogic) AddStudent(req *types.AddStudentReq) (resp *types.Interface, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.StudentRpc.AddStudent(l.ctx, &pb.AddStudentReq{
		Stuno:        req.Stuno,
		Sex:          req.Sex,
		Phone:        req.Phone,
		QQ:           req.QQ,
		GradeClassID: req.GradeClassID,
		Name:         req.Name,
		Remarks:      req.Remarks,
		ByID:         ctxdata.GetUidFromCtx(l.ctx),
	})
	if err != nil {
		return nil, err
	}
	return
}
