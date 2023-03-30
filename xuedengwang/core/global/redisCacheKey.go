package globalkey

import "strings"

/**
redis key except "model cache key"  in here,
but "model cache key" in model
*/

// CacheUserTokenKey /** 用户登陆的token
var CacheUserTokenKey = "xue:user_token:%d"

var CacheEmailCodeKey = "xue:email:%d"

func BuildKey(prefix string) []string {
	if len(prefix) == 0 {
		prefix = "dev:"
	}
	prefix = strings.TrimSpace(prefix)

	var list []string
	//CacheUserTokenKey = prefix + CacheUserTokenKey
	//CacheEmailCodeKey = prefix + CacheEmailCodeKey
	return append(list, CacheUserTokenKey, CacheEmailCodeKey)
}
