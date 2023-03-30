package course

import (
	"net/http"
	"xuedengwang/core/result"

	"xuedengwang/service/gateway/api/internal/logic/course"
	"xuedengwang/service/gateway/api/internal/svc"
)

func CourseAllHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := course.NewCourseAllLogic(r.Context(), svcCtx)
		resp, err := l.CourseAll()
		result.HttpResult(r, w, resp, err)
	}
}
