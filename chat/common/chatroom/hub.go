package chatroom

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	_ "net/http/pprof"
	"time"
)

type Hub struct {
	// 所有通道客户端
	clients map[string]map[*Client]bool
	// 发送消息
	broadcast chan Message
	// 注册客户端
	register chan *Client
	// 注销客户端
	unregister chan *Client
}

var globHub *Hub

func newServer() *Hub {
	return &Hub{
		clients:    make(map[string]map[*Client]bool),
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register: // 登录
			if old, ok := clients[client.Id]; ok {

				// 退出所有聊天室
				for _, sk := range old.channelId {
					if v, ok1 := h.clients[sk]; ok1 {
						delete(v, old)
						h.clients[sk][client] = true
					}
				}
				delete(clients, old.Id)
			}
			clients[client.Id] = client
			msgStore.InfoLogServer(fmt.Sprintf(" all clients: %+v ", clients))
			msgStore.LoginServer(client.Id)
		case client := <-h.unregister: // 注销 / 退出
			// 退出所有聊天室
			for _, sk := range client.channelId {
				if v, ok := h.clients[sk]; ok {
					delete(v, client)

					//如果频道的客户端都退出了，删除频道
					if len(h.clients[sk]) == 0 {
						delete(h.clients, sk)
					}
				}
			}
			// 退出聊天服务
			delete(clients, client.Id)
			msgStore.LogoutServer(client.Id)
		case message := <-h.broadcast: // 接受消息

			var sent []string

			if message.ToUserId != "" {
				if cli, ok := clients[message.ToUserId]; ok {
					fmt.Println(cli.Id, "send", message.Type)
					//send, _ := json.Marshal(message)
					//cli.send <- send // to one user
					h.SendSocketMessage(cli, &message)
					sent = append(sent, cli.Id)
				} else {
					if msgStore != nil {
						msgStore.DelaySendMessage(message.ChannelId, message, sent)
					}
				}
			} else {
				if channel, ok := h.clients[message.ChannelId]; ok {
					for cli, _ := range channel {
						//if cli.Id == message.User && !message.IsKf {
						//	continue
						//}
						message.ToUserId = cli.Id
						message.Detail = cli.Detail
						fmt.Println(cli.Id, "send1", message.Type)
						//send1, _ := json.Marshal(message)
						//cli.send <- send1 // to all users
						h.SendSocketMessage(cli, &message)
						sent = append(sent, cli.Id)
					}
				}
				if msgStore != nil {
					msgStore.DelaySendMessage(message.ChannelId, message, sent)
				}
			}
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func (h *Hub) SendSocketMessage(cli *Client, message *Message) {
	send, _ := json.Marshal(message)
	w, err := cli.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		return
	}
	_, _ = w.Write(send)
	n := len(cli.send)
	for i := 0; i < n; i++ {
		_, _ = w.Write(<-cli.send)
	}
	if err := w.Close(); err != nil {
		return
	}
}
