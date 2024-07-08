package frontend

import (
	"net/http"

	"chat/service/chat/api/internal/svc"
)

func ChatIndexHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "view/home.html")
	}
}
