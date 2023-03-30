package redisx

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-redis/redis/extra/redisotel/v8"
	"github.com/go-redis/redis/v8"
)

// Client redis 客户端
var Client *redis.Client

const (
	// ErrRedisNotFound not exist in redis
	ErrRedisNotFound = redis.Nil
	// DefaultRedisIndex default redis index
	DefaultRedisIndex = 0
)

// RedisManager define a redis manager
//nolint
type RedisManager struct {
	clients []*redis.Client
	*sync.RWMutex
}

// InitDefault init a default redis instance
func InitDefault(c Config) *redis.Client {
	clientManager := NewRedisManager()
	rdb, err := clientManager.GetClient(c, DefaultRedisIndex)
	if err != nil {
		panic(fmt.Sprintf("init redis err: %s", err.Error()))
	}
	Client = rdb

	return rdb
}

// NewRedisManager create a redis manager
func NewRedisManager() *RedisManager {
	return &RedisManager{
		clients: make([]*redis.Client, 10),
		RWMutex: &sync.RWMutex{},
	}
}

// GetClient get a redis instance
func (r *RedisManager) GetClient(conf Config, index int) (*redis.Client, error) {
	// get client from map
	r.RLock()
	if client := r.clients[index]; client != nil {
		r.RUnlock()
		return client, nil
	}
	r.RUnlock()

	c := conf[index]

	// create a redis client
	r.Lock()
	defer r.Unlock()
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Host,
		Password:     c.Password,
		DB:           c.DB,
		MinIdleConns: c.MinIdleConn,
		DialTimeout:  c.DialTimeout,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
		PoolSize:     c.PoolSize,
		PoolTimeout:  c.PoolTimeout,
	})

	// check redis if is ok
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	// hook tracing (using open telemetry)
	if c.EnableTrace {
		rdb.AddHook(redisotel.NewTracingHook())
	}
	r.clients[index] = rdb

	return rdb, nil
}
