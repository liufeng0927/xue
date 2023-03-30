package sys

import (
	"net/http"
	"xuedengwang/core/result"

	"xuedengwang/service/gateway/api/internal/logic/sys"
	"xuedengwang/service/gateway/api/internal/svc"
)

func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := sys.NewLogoutLogic(r.Context(), svcCtx)
		resp, err := l.Logout()
		result.HttpResult(r, w, resp, err)
	}
}
