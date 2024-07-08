package activity

import (
	"chat/common/response"
	"net/http"

	"chat/service/chat/api/internal/logic/backend/activity"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetStatisticsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetStatisticsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := activity.NewGetStatisticsLogic(r.Context(), svcCtx)
		resp, err := l.GetStatistics(&req)
		response.Response(r, w, resp, err)

	}
}
