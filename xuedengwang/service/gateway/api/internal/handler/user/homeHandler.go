package user

import (
	"net/http"
	"xuedengwang/core/result"
	"xuedengwang/service/gateway/api/internal/logic/user"

	"xuedengwang/service/gateway/api/internal/svc"
)

func HomeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewHomeLogic(r.Context(), svcCtx)
		resp, err := l.Home()
		result.HttpResult(r, w, resp, err)
	}
}
