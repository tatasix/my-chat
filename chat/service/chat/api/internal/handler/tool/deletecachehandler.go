package tool

import (
	"chat/common/response"
	"net/http"

	"chat/service/chat/api/internal/logic/tool"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeleteCacheHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CacheRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tool.NewDeleteCacheLogic(r.Context(), svcCtx)
		resp, err := l.DeleteCache(&req)
		response.Response(r, w, resp, err)
	}
}
