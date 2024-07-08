package tool

import (
	"chat/common/response"
	"net/http"

	"chat/service/chat/api/internal/logic/tool"
	"chat/service/chat/api/internal/svc"
)

func QuestionnaireQuestionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := tool.NewQuestionnaireQuestionLogic(r.Context(), svcCtx)
		resp, err := l.QuestionnaireQuestion()
		response.Response(r, w, resp, err)
	}
}
