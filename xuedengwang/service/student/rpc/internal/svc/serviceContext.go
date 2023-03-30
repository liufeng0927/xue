package svc

import (
	"github.com/go-redis/redis/v8"
	"xuedengwang/core/ormx"
	"xuedengwang/core/redisx"
	"xuedengwang/dal/query"
	"xuedengwang/service/student/rpc/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Client
	Query       *query.Query
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisClient := redisx.InitDefault(c.Rdb)
	return &ServiceContext{
		Config:      c,
		RedisClient: redisClient,
		Query:       query.Use(ormx.NewMySQL(&c.DB)),
	}
}