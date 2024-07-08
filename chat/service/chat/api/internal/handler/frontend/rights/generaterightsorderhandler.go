package rights

import (
	"chat/common/response"
	"net/http"

	"chat/service/chat/api/internal/logic/frontend/rights"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GenerateRightsOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GenerateRightsOrderRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := rights.NewGenerateRightsOrderLogic(r.Context(), svcCtx)
		resp, err := l.GenerateRightsOrder(&req)
		response.Response(r, w, resp, err)
	}
}
