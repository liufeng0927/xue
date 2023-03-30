package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/errx"

	"xuedengwang/service/user/rpc/internal/svc"
	"xuedengwang/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrCodeError = errx.NewErrMsg("验证码错误")

func NewVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyCodeLogic {
	return &VerifyCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VerifyCodeLogic) VerifyCode(in *pb.VerifyCodeReq) (*pb.UserInterface, error) {
	// todo: add your logic here and delete this line
	if verify := l.svcCtx.EmailCodeCache.Verify(in.UserID, in.Code); !verify {
		return nil, errors.Wrapf(ErrCodeError, "code  find  not  err userID:%v ,code :%s", in.UserID, in.Code)
	}
	return &pb.UserInterface{}, nil
}
