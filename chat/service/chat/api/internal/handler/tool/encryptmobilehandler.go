package tool

import (
	"chat/common/response"
	"net/http"

	"chat/service/chat/api/internal/logic/tool"
	"chat/service/chat/api/internal/svc"
)

func EncryptMobileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := tool.NewEncryptMobileLogic(r.Context(), svcCtx)
		resp, err := l.EncryptMobile()
		response.Response(r, w, resp, err)
	}
}
