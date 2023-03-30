package gradeclass

import (
	"net/http"
	"xuedengwang/core/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"xuedengwang/service/gateway/api/internal/logic/gradeclass"
	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"
)

func DeleteGradeclassByIDHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteGradeClassByIDReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := gradeclass.NewDeleteGradeclassByIDLogic(r.Context(), svcCtx)
		resp, err := l.DeleteGradeclassByID(&req)
		result.HttpResult(r, w, resp, err)
	}
}
