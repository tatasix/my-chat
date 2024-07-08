package chatroom

import (
	"chat/common/response"
	"net/http"

	"chat/service/chat/api/internal/logic/chatroom"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatRoomJoinHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatRoomJoinRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := chatroom.NewChatRoomJoinLogic(r.Context(), svcCtx)
		resp, err := l.ChatRoomJoin(&req)

		response.Response(r, w, resp, err)
	}
}
