package gradeclass

import (
	"net/http"
	"xuedengwang/core/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"xuedengwang/service/gateway/api/internal/logic/gradeclass"
	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"
)

func UpdateGradeclassHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateGradeClassReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := gradeclass.NewUpdateGradeclassLogic(r.Context(), svcCtx)
		resp, err := l.UpdateGradeclass(&req)
		result.HttpResult(r, w, resp, err)
	}
}
