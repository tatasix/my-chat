package frontend

import (
	"chat/common/response"
	"chat/common/validator"
	"fmt"
	"net/http"

	"chat/service/chat/api/internal/logic/frontend"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCustomerChatRecordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCustomerChatRecordRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		// validate check
		if err := validator.Validate.StructCtx(r.Context(), req); err != nil {
			errMap := validator.Translate(err, &req)
			for _, errFormat := range errMap {
				response.ParamError(r, w, fmt.Errorf(errFormat))
				return
			}
			response.ParamError(r, w, err)
			return
		}
		l := frontend.NewGetCustomerChatRecordLogic(r.Context(), svcCtx)
		resp, err := l.GetCustomerChatRecord(&req)

		response.Response(r, w, resp, err)

	}
}
