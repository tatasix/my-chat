package rights

import (
	"chat/common/response"
	"chat/service/chat/api/internal/types"
	"net/http"

	"chat/service/chat/api/internal/logic/frontend/rights"
	"chat/service/chat/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ApplePayConfirmHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ApplePayConfirmReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamError(r, w, err)
			return
		}

		l := rights.NewApplePayConfirmLogic(r.Context(), svcCtx)
		err := l.ApplePayConfirm(req)

		response.Response(r, w, nil, err)
	}
}
