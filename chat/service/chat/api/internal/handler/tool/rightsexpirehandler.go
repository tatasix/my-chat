package tool

import (
	"net/http"

	"chat/service/chat/api/internal/logic/tool"
	"chat/service/chat/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RightsExpireHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := tool.NewRightsExpireLogic(r.Context(), svcCtx)
		resp, err := l.RightsExpire()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
