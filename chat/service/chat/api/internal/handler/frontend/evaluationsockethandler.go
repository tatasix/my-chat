package frontend

import (
	"chat/common/response"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"net/http"

	"chat/service/chat/api/internal/logic/frontend"
	"chat/service/chat/api/internal/svc"
	"chat/service/chat/api/internal/types"
)

var socketUpgrader = websocket.Upgrader{
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

func EvaluationSocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := socketUpgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("socketUpgrader:", err)
			return
		}
		defer c.Close()
		l := frontend.NewEvaluationSocketLogic(r.Context(), svcCtx)

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
				var req types.EvaluationSocketRequest
				_ = json.Unmarshal(message, &req)
				if req.Heartbeat {
					// 回复一个Pong作为心跳响应
					err = c.WriteMessage(websocket.PongMessage, []byte(""))
					if err != nil {
						log.Println("write pong message failed:", err)
					}
					break
				}
				log.Printf("recv: %s", message)
				//处理消息
				channel := make(chan string, 100)
				baseInfoCh := make(chan any, 1)
				go func() {
					defer close(channel)
					defer close(baseInfoCh)
					res, errChat := l.EvaluationSocket(&req, channel, baseInfoCh)
					if res == nil {
						res = &types.EvaluationSocketResponse{}
					}
					if errChat != nil {
						logx.Error("EvaluationSocket error :", errChat)
						res.ErrorMessage = response.GetErrorMessage(errChat)
					}
					res.MessageId = req.MessageID
					logx.Infof("EvaluationSocketHandler message:%+v", res)
					returnMessage, _ := json.Marshal(res)
					//发送消息
					errWrite := c.WriteMessage(mt, returnMessage)
					if errWrite != nil {
						logx.Error("WriteMessage error :", errWrite)
						_ = c.Close()
					}
				}()
				baseInfo := <-baseInfoCh
				//close(baseInfoCh)
				var rs []rune
				// 加快初次响应的时间 后续可改为阶梯式（用户体验好）
				length := 1
				for {
					s, ok := <-channel
					if !ok {
						// 数据接受完成
						if len(rs) > 0 {
							SendSocketMessage(string(rs), req.MessageID, c, mt, baseInfo)
						}
						break
					}
					rs = append(rs, []rune(s)...)

					if len(rs) > length {
						SendSocketMessage(string(rs), req.MessageID, c, mt, baseInfo)
						rs = []rune{}
						if length < 4 {
							length++
						}
					}
				}

			default:
				log.Println("Received unknown message type")
				break
			}

		}
	}
}

func SendSocketMessage(s, messageId string, c *websocket.Conn, mt int, baseInfo any) {
	baseReturn := types.EvaluationSocketResponse{}
	if baseInfo != nil {
		switch d := baseInfo.(type) {
		case types.EvaluationSocketResponse:
			baseReturn = d
		default:
			fmt.Println("Unsupported type:", d)
		}
	}

	//发送消息
	returnMessage, _ := json.Marshal(types.EvaluationSocketResponse{
		Message:    s,
		MessageId:  messageId,
		Tarot:      baseReturn.Tarot,
		TarotImage: baseReturn.TarotImage,
		TarotName:  baseReturn.TarotName,
	})
	errWrite := c.WriteMessage(mt, returnMessage)
	if errWrite != nil {
		logx.Error("WriteMessage error :", errWrite)
		_ = c.Close()

	}
}
