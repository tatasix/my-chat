package rights

import (
	"chat/common/response"
	"net/http"

	"chat/service/chat/api/internal/logic/frontend/rights"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetRightsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRightsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := rights.NewGetRightsLogic(r.Context(), svcCtx)
		resp, err := l.GetRights(&req)
		response.Response(r, w, resp, err)
	}
}
