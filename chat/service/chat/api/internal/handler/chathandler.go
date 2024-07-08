package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"

	"chat/common/response"
	"chat/common/validator"
	"chat/service/chat/api/internal/logic"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
)

func ChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamError(r, w, err)
			return
		}
		// validate check
		timeout := 10 * time.Minute
		ctxNew, _ := context.WithTimeout(context.Background(), timeout)
		//defer cancel()
		if err := validator.Validate.StructCtx(ctxNew, req); err != nil {
			errMap := validator.Translate(err, &req)
			for _, errFormat := range errMap {
				response.ParamError(r, w, fmt.Errorf(errFormat))
				return
			}
			response.ParamError(r, w, err)
			return
		}

		l := logic.NewChatLogic(ctxNew, svcCtx)
		resp, err := l.Chat(&req)
		response.Response(r, w, resp, err)
	}
}
