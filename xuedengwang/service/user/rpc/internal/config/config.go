package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"xuedengwang/core/email"
	"xuedengwang/core/ormx"
	"xuedengwang/core/redisx"
)

type Config struct {
	zrpc.RpcServerConf
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	DB    ormx.Config
	Rdb   redisx.Config
	Email email.Config
}
