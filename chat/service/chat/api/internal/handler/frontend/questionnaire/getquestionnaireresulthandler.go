package questionnaire

import (
	"chat/common/response"
	"net/http"

	"chat/service/chat/api/internal/logic/frontend/questionnaire"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetQuestionnaireResultHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetQuestionnaireRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := questionnaire.NewGetQuestionnaireResultLogic(r.Context(), svcCtx)
		resp, err := l.GetQuestionnaireResult(&req)
		response.Response(r, w, resp, err)
	}
}
