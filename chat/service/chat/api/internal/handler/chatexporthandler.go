package handler

import (
	"net/http"

	"chat/service/chat/api/internal/logic"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatExportHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetChatListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewChatExportLogic(r.Context(), svcCtx)
		resp, err := l.ChatExport(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
