package im

import (
	"chat/common/chatroom"
	"chat/common/response"
	"chat/common/xerr"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type cliDetail struct {
	User     string `json:"user"`
	UserName string `json:"user_name"`
	KfId     string `json:"kf_id"`
}

type Server struct {
	Ctx *svc.ServiceContext
}

//func (i Server) Start(req *types.ChatRoomSocketRequest, w http.ResponseWriter, r *http.Request, store *Store) {
//chatroom.NewServer(w, r, req.User, cliDetail{User: req.User, UserName: req.UserName, KfId: req.OpenKfID}, store)
//}

func (i Server) Start() {
	Run(i.Ctx)
}
func (i Server) Stop() {
	fmt.Println("im service was stop...")
}

func Run(ctx *svc.ServiceContext) {
	mux := http.NewServeMux()
	mux.HandleFunc("/chatroom/socket1", func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatRoomSocketRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		if req.User == "" || req.UserName == "" {
			response.Response(r, w, nil, xerr.NewErrCode(xerr.SystemError))
			return
		}

		chatroom.NewServer(w, r, req.User, cliDetail{User: req.User, UserName: req.UserName}, &Store{Svc: ctx, Ctx: r.Context()})
	})
	fmt.Printf("Starting Ws Server at 0.0.0.0:8881...\n")
	err := http.ListenAndServe(":8881", mux)
	if err != nil {
		fmt.Println("ws err：", err)
		return
	}
}
