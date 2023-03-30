package svc

import (
	"context"
	"github.com/go-redis/redis/v8"
	"xuedengwang/core/email"
	"xuedengwang/core/ormx"
	"xuedengwang/core/redisx"
	"xuedengwang/dal/query"
	"xuedengwang/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config         config.Config
	RedisClient    *redis.Client
	Query          *query.Query
	EmailCodeCache *email.CodeCache
}

func NewServiceContext(c config.Config) *ServiceContext {
	ctx := context.Background()
	redisClient := redisx.InitDefault(c.Rdb)
	return &ServiceContext{
		Config:         c,
		RedisClient:    redisClient,
		Query:          query.Use(ormx.NewMySQL(&c.DB)),
		EmailCodeCache: email.NewCacheSmsCode(ctx, redisClient),
	}
}
