package result

import (
	"net/http"
	"time"
)

type ResponseSuccessBean struct {
	Status    uint32      `json:"status"`
	Message   string      `json:"message"`
	Result    interface{} `json:"result"`
	Timestamp int64       `json:"timestamp"`
}
type NullJson struct{}

func Success(data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{http.StatusOK, "OK", data, time.Now().UnixMilli()}
}

type ResponseErrorBean struct {
	Status  uint32 `json:"status"`
	Message string `json:"message"`
}

func Error(errCode uint32, errMsg string) *ResponseErrorBean {
	return &ResponseErrorBean{errCode, errMsg}
}
