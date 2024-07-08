package handler

import (
	"chat/common/response"
	"net/http"

	"chat/service/chat/api/internal/logic"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateCustomerConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CustomerConfig
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCreateCustomerConfigLogic(r.Context(), svcCtx)
		req.KfId = l.GenerateRandomString(32)

		resp, err := l.CreateCustomerConfig(&req)
		response.Response(r, w, resp, err)
	}
}
