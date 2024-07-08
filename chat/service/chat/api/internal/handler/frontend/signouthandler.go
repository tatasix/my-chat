package frontend

import (
	"chat/common/response"
	"net/http"

	"chat/service/chat/api/internal/logic/frontend"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SignoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	//账号注销
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SignoutRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := frontend.NewSignoutLogic(r.Context(), svcCtx)
		resp, err := l.Signout(&req)
		response.Response(r, w, resp, err)
	}
}
