package email

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strings"
	globalkey "xuedengwang/core/global"

	"time"
)

type CodeCache struct {
	cacheClient *redis.Client
	ctx         context.Context
}

func NewCacheSmsCode(ctx context.Context, cacheClient *redis.Client) *CodeCache {
	return &CodeCache{
		cacheClient: cacheClient,
		ctx:         ctx,
	}
}

func (c *CodeCache) Set(userID int64, code string, expire time.Duration) error {
	key := fmt.Sprintf(globalkey.CacheEmailCodeKey, userID)
	return c.cacheClient.Set(c.ctx, key, code, expire).Err()
}

// Get 获取code
func (c *CodeCache) Get(userID int64) string {
	key := fmt.Sprintf(globalkey.CacheEmailCodeKey, userID)
	val, err := c.cacheClient.Get(c.ctx, key).Result()
	if err != nil {
		return ""
	}

	return val
}

// Verify 验证
func (c *CodeCache) Verify(userID int64, input string) bool {
	if len(input) == 0 {
		return false
	}
	v := c.Get(userID)

	clear := strings.EqualFold(v, input)
	//key := fmt.Sprintf(globalkey.CacheEmailCodeKey, userID)
	//if clear {
	//	//clear为true，删除这个验证码
	//	err := c.cacheClient.Del(c.ctx, key).Err()
	//	if err != nil {
	//		return true
	//	}
	//}
	//return false
	return clear
}
