package course

import (
	"net/http"
	"xuedengwang/core/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"xuedengwang/service/gateway/api/internal/logic/course"
	"xuedengwang/service/gateway/api/internal/svc"
	"xuedengwang/service/gateway/api/internal/types"
)

func AddCourseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCourseReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := course.NewAddCourseLogic(r.Context(), svcCtx)
		resp, err := l.AddCourse(&req)
		result.HttpResult(r, w, resp, err)
	}
}
