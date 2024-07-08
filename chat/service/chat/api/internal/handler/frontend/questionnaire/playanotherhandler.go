package questionnaire

import (
	"chat/common/response"
	"net/http"

	"chat/service/chat/api/internal/logic/frontend/questionnaire"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PlayAnotherHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PlayAnotherRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := questionnaire.NewPlayAnotherLogic(r.Context(), svcCtx)
		resp, err := l.PlayAnother(&req)
		response.Response(r, w, resp, err)
	}
}