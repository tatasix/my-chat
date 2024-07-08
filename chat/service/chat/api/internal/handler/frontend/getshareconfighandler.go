package frontend

import (
	"chat/common/response"
	"net/http"

	"chat/service/chat/api/internal/logic/frontend"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetShareConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetShareConfigRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := frontend.NewGetShareConfigLogic(r.Context(), svcCtx)
		resp, err := l.GetShareConfig(&req)
		response.Response(r, w, resp, err)
	}
}
