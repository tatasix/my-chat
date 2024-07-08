package handler

import (
	"context"
	"net/http"
	"time"

	"chat/service/chat/api/internal/logic"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AutoChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DoGenerateActiveChatContentRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		timeout := 20 * time.Minute
		ctxNew, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel()
		l := logic.NewAutoChatLogic(ctxNew, svcCtx)
		resp, err := l.AutoChat(&req)
		if err != nil {
			httpx.ErrorCtx(ctxNew, w, err)
		} else {
			httpx.OkJsonCtx(ctxNew, w, resp)
		}
	}
}
