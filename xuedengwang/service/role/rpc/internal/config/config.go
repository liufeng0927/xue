package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"xuedengwang/core/ormx"
	"xuedengwang/core/redisx"
)

type Config struct {
	zrpc.RpcServerConf
	DB  ormx.Config
	Rdb redisx.Config
}
