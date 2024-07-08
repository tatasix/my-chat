package rights

import (
	"net/http"

	"chat/service/chat/api/internal/logic/frontend/rights"
	"chat/service/chat/api/internal/svc"
)

func NotifyApplePayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := rights.NewNotifyApplePayLogic(r.Context(), svcCtx)
		l.NotifyApplePay(r)
	}
}
