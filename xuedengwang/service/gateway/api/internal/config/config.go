package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"xuedengwang/core/ormx"
)

type Config struct {
	rest.RestConf
	JwtAuth           JwtAuth
	Captcha           Captcha
	UserRpcConf       zrpc.RpcClientConf
	RoleRpcConf       zrpc.RpcClientConf
	CourseRpcConf     zrpc.RpcClientConf
	StudentRpcConf    zrpc.RpcClientConf
	TeacherRpcConf    zrpc.RpcClientConf
	GradeClassRpcConf zrpc.RpcClientConf
	ScoresRpcConf     zrpc.RpcClientConf
	DB                ormx.Config
}

type JwtAuth struct {
	AccessSecret string
	AccessExpire int64
}

type Captcha struct {
	KeyLength int // 验证码长度
	ImgWidth  int // 验证码宽度
	ImgHeight int // 验证码高度
}
