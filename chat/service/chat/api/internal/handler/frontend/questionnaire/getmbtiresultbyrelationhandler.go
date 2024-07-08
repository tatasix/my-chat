package questionnaire

import (
	"chat/common/response"
	"net/http"

	"chat/service/chat/api/internal/logic/frontend/questionnaire"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetMbtiResultByRelationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMbtiResultByRelationRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := questionnaire.NewGetMbtiResultByRelationLogic(r.Context(), svcCtx)
		resp, err := l.GetMbtiResultByRelation(&req)
		response.Response(r, w, resp, err)
	}
}
