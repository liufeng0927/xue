package sys

import (
	"net/http"
	"xuedengwang/core/result"
	"xuedengwang/service/gateway/api/internal/logic/sys"
	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := sys.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		result.HttpResult(r, w, resp, err)
	}
}
