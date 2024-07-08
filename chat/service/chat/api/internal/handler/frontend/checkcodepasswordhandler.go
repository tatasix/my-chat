package frontend

import (
	"chat/common/response"
	"net/http"

	"chat/service/chat/api/internal/logic/frontend"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CheckCodePasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckCodeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := frontend.NewCheckCodePasswordLogic(r.Context(), svcCtx)
		resp, err := l.CheckCodePassword(&req)
		response.Response(r, w, resp, err)

	}
}
