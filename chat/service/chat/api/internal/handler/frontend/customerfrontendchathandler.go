package frontend

import (
	"chat/common/response"
	"chat/common/xerr"
	"chat/service/chat/api/internal/logic"
	"chat/service/chat/api/internal/logic/frontend"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		//origin := r.Header.Get("Origin")
		//if origin != "http://example.com" {
		//	return false
		//}
		return true
	},
}

func CustomerFrontendChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()
		l := frontend.NewCustomerFrontendChatLogic(context.Background(), svcCtx)

		for {
			//读取消息
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					logx.Error("[ERROR] Unexpected read message error: ", err)
				}
				break
			}
			// 根据消息类型进行处理
			switch mt {
			case websocket.TextMessage:
				log.Println("Received chat message:", string(message))

				// 处理聊天消息
				// 将消息进行反序列化，并处理
				var req types.CustomerChatRequest
				_ = json.Unmarshal(message, &req)
				if req.Heartbeat {
					// 回复一个Pong作为心跳响应
					err = c.WriteMessage(websocket.PongMessage, []byte(""))
					if err != nil {
						log.Println("write pong message failed:", err)
					}
					break
				}

				////检查请求次数
				//chatLimitLogic := logic.NewChatLimitRequestLogic(r.Context(), svcCtx)
				//userId, err := chatLimitLogic.CheckUserRequestedTimes(req.Token, req.User)

				if err != nil {
					var res types.CustomerChatResponse
					res.ErrorMessage = err.Error()
					returnMessage, _ := json.Marshal(res)
					//发送消息
					errWrite := c.WriteMessage(mt, returnMessage)
					if errWrite != nil {
						logx.Error("WriteMessage error :", errWrite)
					}
					break
				}

				log.Printf("recv: %s", message)
				//处理消息
				channel := make(chan string, 50)
				go func() {
					defer func() {
						close(channel)
						// 处理 goroutine 泄漏
						recover()
					}()

					res, errChat := l.CustomerChat(&req, channel)
					if errChat != nil {
						logx.Error("CustomerChatV2 error :", errChat)
						res.ErrorMessage = errChat.Error()
						//close(channel)
					}

					logx.Infof("CustomerFrontendChatHandler message:%+v", res)
					returnMessage, _ := json.Marshal(res)
					//发送消息
					errWrite := c.WriteMessage(mt, returnMessage)
					if errWrite != nil {
						logx.Error("WriteMessage error :", errWrite)
					}

					if res.IsEnd && res.Message != "系统繁忙，请稍后再试～" {
						//请求次数 +1
						//chatLimitLogic.AddUserTodayRequestedTimes(userId)
					}
				}()
				var rs []rune
				// 加快初次响应的时间 后续可改为阶梯式（用户体验好）
				length := 1
				for {
					s, ok := <-channel
					if !ok {
						// 数据接受完成
						if len(rs) > 0 {
							SendMessage(string(rs), req.MessageID, c, mt)
							rs = []rune{}
						}
						break
					}
					rs = append(rs, []rune(s)...)

					if len(rs) > length {
						SendMessage(string(rs), req.MessageID, c, mt)
						rs = []rune{}
						if length < 4 {
							length++
						}
					}
				}

			case websocket.PingMessage:
				log.Println("Received ping message")
				// 处理心跳消息

				// 回复一个Pong作为心跳响应
				err := c.WriteMessage(websocket.PongMessage, []byte(""))
				if err != nil {
					log.Println("write pong message failed:", err)
					break
				}

			default:
				log.Println("Received unknown message type")
				break
			}

		}
	}
}

func SendMessage(s, messageId string, c *websocket.Conn, mt int) {
	response := types.CustomerChatResponse{
		Message:   s,
		MessageID: messageId,
	}
	//发送消息
	returnMessage, _ := json.Marshal(response)
	errWrite := c.WriteMessage(mt, returnMessage)
	if errWrite != nil {
		logx.Error("WriteMessage error :", errWrite)
	}
}

func ResetRequestTimesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//检查请求次数
		var req types.CustomerChatRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		chatLimitLogic := logic.NewChatLimitRequestLogic(r.Context(), svcCtx)
		chatLimitLogic.ResetRequestTimesByUserToken(req.Token)
		response.Response(r, w, nil, xerr.NewErrCodeMsg(xerr.SUCCESS, "success"))
	}
}
