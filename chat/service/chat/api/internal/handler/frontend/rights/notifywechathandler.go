package rights

import (
	"net/http"

	"chat/service/chat/api/internal/logic/frontend/rights"
	"chat/service/chat/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func NotifyWechatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := rights.NewNotifyWechatLogic(r.Context(), svcCtx)
		res, err := l.NotifyWechat(r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJson(w, res)
		}
	}
}
