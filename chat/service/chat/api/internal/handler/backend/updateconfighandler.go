package backend

import (
	"net/http"

	"chat/service/chat/api/internal/logic/backend"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateConfigRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := backend.NewUpdateConfigLogic(r.Context(), svcCtx)
		resp, err := l.UpdateConfig(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
