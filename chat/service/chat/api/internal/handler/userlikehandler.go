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

func UserLikeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLikeReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamError(r, w, err)
			return
		}

		l := logic.NewUserLikeLogic(r.Context(), svcCtx)
		resp, err := l.UserLike(&req)
		if err != nil {
			response.Response(r, w, nil, xerr.NewErrCodeMsg(xerr.RequestParamError, err.Error()))
			return
		}
		response.Response(r, w, resp, err)
	}
}
