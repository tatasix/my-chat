package handler

import (
	"chat/common/response"
	"chat/common/xerr"
	"net/http"

	"chat/service/chat/api/internal/logic"
	"chat/service/chat/api/internal/svc"
)

func ChinaCalendarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewChinaCalendarLogic(r.Context(), svcCtx)
		resp, err := l.ChinaCalendar()

		if err != nil {
			response.Response(r, w, nil, xerr.NewErrCodeMsg(xerr.RequestParamError, err.Error()))
			return
		}

		response.Response(r, w, resp, err)
	}
}
