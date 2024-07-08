package frontend

import (
	"chat/common/response"

	"net/http"

	"chat/service/chat/api/internal/logic/frontend"
	"chat/service/chat/api/internal/svc"
)

func GenerateUniqueIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := frontend.NewGenerateUniqueIdLogic(r.Context(), svcCtx)
		resp, err := l.GenerateUniqueId()
		response.Response(r, w, resp, err)

	}
}
