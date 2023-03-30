package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"xuedengwang/core/auth"
	"xuedengwang/core/ctxdata"
	"xuedengwang/core/errx"
	"xuedengwang/core/jwtx"
	"xuedengwang/service/user/rpc/internal/svc"
	"xuedengwang/service/user/rpc/pb"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrUserNotFindError = errx.NewErrMsg("用户不存在，请注册后重新登录")
var ErrUsernamePasswordError = errx.NewErrMsg("账号或密码不正确")
var ErrGenerateTokenError = errx.NewErrMsg("生成token失败")

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {

	userDB := l.svcCtx.Query.User

	user, err := userDB.WithContext(l.ctx).Preload(userDB.Role).Where(userDB.Username.Eq(in.Username)).First()
	if err != nil {
		return nil, errors.Wrapf(ErrUserNotFindError, "login username not find err username:%v , err:%v", in.Username, err)
	}
	if !auth.ComparePasswords(user.Password, in.Password) {
		return nil, errors.Wrap(ErrUsernamePasswordError, "loginByPassword loginKey or password error")
	}
	//2、生成token

	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := jwtx.GenerateToken(map[string]interface{}{ctxdata.CtxKeyJwtUserId: user.ID}, l.svcCtx.Config.JwtAuth.AccessSecret, accessExpire)
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "getJwtToken err userId:%d , err:%v", user.ID, err)
	}

	role := &pb.Role{
		CreateTime: user.Role.CreateTime.String(),
		CreateBy:   user.Role.CreateBy,
		UpdateTime: user.Role.CreateTime.String(),
		UpdateBy:   user.Role.UpdateBy,
		Remarks:    user.Role.Remarks,
		ID:         user.Role.ID,
		Name:       user.Role.Name,
		Code:       user.Role.Code,
	}

	return &pb.LoginResp{
		Role:       role,
		CreateTime: user.CreateTime.String(),
		Sex:        user.Sex,
		UserIcon:   *user.UserIcon,
		Username:   user.Username,
		Realname:   user.Realname,
		Email:      user.Email,
		Token:      accessToken,
	}, nil

}
