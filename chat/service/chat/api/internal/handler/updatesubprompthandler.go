package handler

import (
	"chat/common/response"
	"chat/common/xerr"
	"net/http"

	"chat/service/chat/api/internal/logic"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateSubPromptHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateSubPromptRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateSubPromptLogic(r.Context(), svcCtx)
		err := l.UpdateSubPrompt(&req)
		if err != nil {
			response.Response(r, w, nil, xerr.NewErrCodeMsg(xerr.RecordNotFound, err.Error()))
			return
		}

		response.Response(r, w, nil, err)
	}
}
