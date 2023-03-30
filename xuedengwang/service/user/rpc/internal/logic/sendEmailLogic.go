package logic

import (
	"context"
	"github.com/pkg/errors"
	"xuedengwang/core/email"
	"xuedengwang/core/errx"
	"xuedengwang/core/randx"
	"xuedengwang/service/user/rpc/internal/svc"
	"xuedengwang/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendEmailLogic) SendEmail(in *pb.SendEmailReq) (*pb.UserInterface, error) {
	// todo: add your logic here and delete this line

	if len(in.Email) == 0 {
		userDB := l.svcCtx.Query.User
		if err := userDB.WithContext(l.ctx).Select(userDB.Email).Where(userDB.ID.Eq(in.UserID)).Scan(&in.Email); err != nil {
			return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "user db find email err:%v,req:%+v", err, in)
		}
	}

	randCode := randx.Krand(6, randx.KC_RAND_KIND_NUM)
	client := email.NewClient(in.Email, randCode, l.svcCtx.Config.Email)
	go client.Send()

	if err := l.svcCtx.EmailCodeCache.Set(in.UserID, randCode, l.svcCtx.Config.Email.ExpireTime); err != nil {
		return nil, errors.Wrapf(errx.NewErrCode(errx.DB_ERROR), "emailcode set userid:%d to cache error:%v", in.UserID, err)
	}

	return &pb.UserInterface{}, nil
}
