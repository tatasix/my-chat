package activity

import (
	"chat/common/response"
	"net/http"

	"chat/service/chat/api/internal/logic/frontend/activity"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChangeUserIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChangeUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := activity.NewChangeUserIdLogic(r.Context(), svcCtx)
		resp, err := l.ChangeUserId(&req)
		response.Response(r, w, resp, err)
	}
}
