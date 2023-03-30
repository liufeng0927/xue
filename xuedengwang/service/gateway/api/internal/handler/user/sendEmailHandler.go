package user

import (
	"net/http"
	"xuedengwang/core/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"xuedengwang/service/gateway/api/internal/logic/user"
	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"
)

func SendEmailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendEmailReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewSendEmailLogic(r.Context(), svcCtx)
		resp, err := l.SendEmail(&req)
		result.HttpResult(r, w, resp, err)
	}
}