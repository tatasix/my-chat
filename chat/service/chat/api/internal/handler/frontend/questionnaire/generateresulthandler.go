package questionnaire

import (
	"chat/common/response"
	"chat/service/chat/api/internal/logic/frontend/questionnaire"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"net/http"

	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func GenerateResultHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()
		l := questionnaire.NewGenerateResultLogic(context.Background(), svcCtx)

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
				var req types.GenerateResultRequest
				_ = json.Unmarshal(message, &req)
				if req.Heartbeat {
					// 回复一个Pong作为心跳响应
					err = c.WriteMessage(websocket.PongMessage, []byte(""))
					if err != nil {
						log.Println("write pong message failed:", err)
					}
					break
				}

				if err != nil {
					var res types.GenerateResultResponse
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
				baseInfoCh := make(chan any, 1)

				go func() {
					defer func() {
						close(channel)
						close(baseInfoCh)
						// 处理 goroutine 泄漏
						recover()
					}()

					res, errChat := l.GenerateResult(&req, channel, baseInfoCh)
					if errChat != nil {
						logx.Error("GenerateResult error :", errChat)
						res.ErrorMessage = response.GetErrorMessage(errChat)
					}
					if res == nil {
						res = &types.GenerateResultResponse{}
					}
					res.IsEnd = true
					res.MessageID = req.MessageID

					logx.Infof("GenerateResultHandler message:%+v", res)
					returnMessage, _ := json.Marshal(res)
					//发送消息
					errWrite := c.WriteMessage(mt, returnMessage)
					if errWrite != nil {
						logx.Error("WriteMessage error :", errWrite)
					}

				}()
				baseInfo := <-baseInfoCh
				var rs []rune
				// 加快初次响应的时间 后续可改为阶梯式（用户体验好）
				length := 1
				for {
					s, ok := <-channel
					if !ok {
						// 数据接受完成
						if len(rs) > 0 {
							SendMessage(string(rs), req.MessageID, c, mt, baseInfo)
							rs = []rune{}
						}
						break
					}
					rs = append(rs, []rune(s)...)

					if len(rs) > length {
						SendMessage(string(rs), req.MessageID, c, mt, baseInfo)
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

func SendMessage(s, messageId string, c *websocket.Conn, mt int, baseInfo any) {

	baseReturn := types.GenerateResultResponse{}
	if baseInfo != nil {
		switch d := baseInfo.(type) {
		case types.GenerateResultResponse:
			baseReturn = d
		default:
			fmt.Println("Unsupported type:", d)
		}
	}
	baseReturn.MessageID = messageId
	if s != "" {
		baseReturn.Message = s
	}
	//发送消息
	returnMessage, _ := json.Marshal(baseReturn)
	errWrite := c.WriteMessage(mt, returnMessage)
	if errWrite != nil {
		logx.Error("WriteMessage error :", errWrite)
		_ = c.Close()

	}

	//
	//response := types.CustomerChatResponse{
	//	Message:   s,
	//	MessageID: messageId,
	//}
	////发送消息
	//returnMessage, _ := json.Marshal(response)
	//errWrite := c.WriteMessage(mt, returnMessage)
	//if errWrite != nil {
	//	logx.Error("WriteMessage error :", errWrite)
	//}
}
