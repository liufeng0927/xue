package ctxdata

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

// CtxKeyJwtUserId get uid from ctx
var CtxKeyJwtUserId = "jwtUserId"

func GetUidFromCtx(ctx context.Context) int64 {
	var uid int64
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(float64); ok {
		uid = int64(jsonUid)
	} else {
		logx.WithContext(ctx).Errorf("[GetUidFromCtx] user id not exist")
	}
	return uid
}
