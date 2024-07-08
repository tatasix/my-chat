package frontend

import (
	"chat/common/response"
	"chat/common/xerr"
	"chat/service/chat/api/internal/logic/frontend"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		isImg := frontend.CheckUploadImages(r)

		if !isImg {
			response.Response(r, w, nil, xerr.NewErrCodeMsg(xerr.FORBIDDEN, "非法文件"))
			return
		}

		l := frontend.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(&req, r)
		response.Response(r, w, resp, err)
	}
}
