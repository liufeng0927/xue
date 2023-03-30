package httpx

import (
	"context"
	"net/http"
	"time"
)

const (
	// DefaultRetryTimes 如果请求失败，最多重试3次
	DefaultRetryTimes = 3
	// DefaultRetryDelay 在重试前，延迟等待10毫秒
	DefaultRetryDelay = time.Millisecond * 10

	// StatusReadRespErr read resp body err, should re-call doReq again.
	StatusReadRespErr = -204

	// StatusDoReqErr do req err, should re-call doReq again.
	StatusDoReqErr = -500
)

// RetryVerify Verify parse the body and verify that it is correct
type RetryVerify func(body []byte) (shouldRetry bool)

func shouldRetry(ctx context.Context, httpCode int) bool {
	select {
	case <-ctx.Done():
		return false
	default:
	}

	switch httpCode {
	case
		StatusReadRespErr,
		StatusDoReqErr,

		http.StatusRequestTimeout,
		http.StatusLocked,
		http.StatusTooEarly,
		http.StatusTooManyRequests,

		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:

		return true

	default:
		return false
	}
}
