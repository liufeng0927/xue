package role

import (
	"net/http"
	"xuedengwang/core/result"

	"xuedengwang/service/gateway/api/internal/logic/role"
	"xuedengwang/service/gateway/api/internal/svc"
)

func RoleAllHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := role.NewRoleAllLogic(r.Context(), svcCtx)
		resp, err := l.RoleAll()
		result.HttpResult(r, w, resp, err)
	}
}
