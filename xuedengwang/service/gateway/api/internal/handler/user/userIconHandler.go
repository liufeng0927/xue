package user

import (
	"net/http"
	"xuedengwang/core/result"

	"xuedengwang/service/gateway/api/internal/logic/user"
	"xuedengwang/service/gateway/api/internal/svc"
)

func UserIconHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUserIconLogic(r.Context(), svcCtx)

		file, fileHeader, err := r.FormFile("fileResource")
		if err != nil {
			result.ParamErrorResult(r, w, err)
		}
		resp, err := l.UserIcon(file, fileHeader)
		result.HttpResult(r, w, resp, err)
	}
}
