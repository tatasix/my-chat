package chatroom

import (
	"chat/common/chatroom"
	"chat/service/chat/api/internal/service/im"
	"chat/service/chat/api/internal/svc"
	"context"
	"net/http"
)

func ChatRoomSocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.FormValue("user")
		name := r.FormValue("name")
		chatroom.NewServer(w, r, user, cliDetail{User: user, UserName: name}, &im.Store{Svc: svcCtx, Ctx: context.Background()})
	}
}

type cliDetail struct {
	User     string `json:"user"`
	UserName string `json:"user_name"`
}
