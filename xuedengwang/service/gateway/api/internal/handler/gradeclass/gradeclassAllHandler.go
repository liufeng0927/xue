package gradeclass

import (
	"net/http"
	"xuedengwang/core/result"

	"xuedengwang/service/gateway/api/internal/logic/gradeclass"
	"xuedengwang/service/gateway/api/internal/svc"
)

func GradeclassAllHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := gradeclass.NewGradeclassAllLogic(r.Context(), svcCtx)
		resp, err := l.GradeclassAll()
		result.HttpResult(r, w, resp, err)
	}
}
