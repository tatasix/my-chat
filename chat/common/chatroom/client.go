package chatroom

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub        *Hub
	conn       *websocket.Conn
	send       chan []byte
	channelId  []string
	Id         string
	Detail     interface{}
	userActive chan struct{}
	isClose    bool
}

var clients map[string]*Client

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (c *Client) Close() {
	if c.isClose {
		return
	}
	_ = c.conn.Close()
	c.isClose = true
	fmt.Println("Close", c.isClose)
	c.send <- []byte{}
}

func (c *Client) WriteMsg() {
	defer func() {
		fmt.Println("WriteMsg 退出")
		c.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if c.isClose {
				return
			}
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{}) // 错误 关闭 channel
				msgStore.ErrorLogServer(fmt.Errorf("系统错误：未知"))
				return
			}
			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, _ = w.Write(message)
			n := len(c.send)
			for i := 0; i < n; i++ {
				_, _ = w.Write(<-c.send)
			}
			if err := w.Close(); err != nil {
				return
			}
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func (c *Client) ReadMsg() {
	defer func() {
		fmt.Println(c.Id, "ReadMsg 退出")
		c.hub.unregister <- c
		c.Close()
	}()
	fmt.Println(c.Id, "ReadMsg1", time.Now().Format("2006-01-02 15:04:05"))
	for {
		fmt.Println(c.Id, "ReadMsg2", time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Millisecond * 100)
		_, strByte, err := c.conn.ReadMessage()
		if err != nil {
			// 检查连接是否关闭
			if _, isCloseErr := err.(*websocket.CloseError); isCloseErr {
				// 连接已关闭，可以采取相应的处理
				return
			}
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				//c.hub.unregister <- c
				msgStore.ErrorLogServer(fmt.Errorf("系统错误：%v", err))
			}
			msgStore.ErrorLogServer(fmt.Errorf("hub:`%+v` hub.clients: `%+v ` err: `%+v `", c.hub, c.hub.clients, err))

			break
		}
		if clients[c.Id] != c {
			msgStore.ErrorLogServer(fmt.Errorf("用户`%s`未登录，不能发送消息 c.detail %s,clients %s", c.Id, c.Detail, clients[c.Id].Detail))
			continue
		}
		if c.isClose {
			return
		}
		// 用户活跃
		//c.userActive <- struct{}{}
		message := Message{}
		err = json.Unmarshal(strByte, &message)
		msgStore.InfoLogServer(fmt.Sprintf(" ReadMsg message:%v", message))

		if message.Heartbeat {
			// 回复一个Pong作为心跳响应
			pong := Message{
				Type:      SendMessagePong,
				Heartbeat: true,
			}
			pongJson, _ := json.Marshal(pong)
			err = c.conn.WriteMessage(websocket.TextMessage, pongJson)
			if err != nil {
				msgStore.ErrorLogServer(fmt.Errorf("write pong message failed:%v", err))
			}
			continue
		}
		message.SendTime = time.Now().Format("2006-01-02 15:04:05")
		message.Detail = c.Detail
		if message.Type == SendMessageCreate || message.Type == SendMessageJoin || message.Type == SendMessageGenerateOne {
			if string(strByte) != "" {
				if msgStore != nil {
					msgStore.SendMessage(message)
				}
			}
			continue
		}
		if message.ToUserId != "" {
			if _, ok := clients[message.ToUserId]; !ok {
				continue
			}
		} else if _, ok := c.hub.clients[message.ChannelId][c]; !ok {
			msgStore.ErrorLogServer(fmt.Errorf("用户`%s`未监听`%s`频道，不能发送消息", c.Id, message.ChannelId))
			msgStore.InfoLogServer(fmt.Sprintf("hub:`%+v` hub.clients: `%+v ` message: `%+v `", c.hub, c.hub.clients, message))
			continue
		}

		if string(strByte) != "" {
			c.hub.broadcast <- message // 转发读取到的channel消息
			if msgStore != nil {
				msgStore.SendMessage(message)
			}
		}
	}
}
