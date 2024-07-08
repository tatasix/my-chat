package rights

import (
	"chat/common/response"
	"net/http"

	"chat/service/chat/api/internal/logic/frontend/rights"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CheckRightsTempHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckRightsTempRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := rights.NewCheckRightsTempLogic(r.Context(), svcCtx)
		resp, err := l.CheckRightsTemp(&req)
		response.Response(r, w, resp, err)
	}
}
