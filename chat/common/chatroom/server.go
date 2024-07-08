package chatroom

import (
	"fmt"
	"net/http"
)

func NewServer(w http.ResponseWriter, r *http.Request, clientId string, detail interface{}, store messageInterface) {
	if globHub == nil {
		globHub = newServer()
		go globHub.run()
	}
	if clients == nil {
		clients = make(map[string]*Client)
	}
	if msgStore == nil {
		msgStore = store
	}
	runWs(globHub, w, r, clientId, detail)
}

func runWs(hub *Hub, w http.ResponseWriter, r *http.Request, clientId string, detail interface{}) {
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("升级get请求错误", err)
		return
	}
	client := &Client{
		hub:        hub,
		conn:       conn,
		send:       make(chan []byte),
		Id:         clientId,
		Detail:     detail,
		userActive: make(chan struct{}),
		isClose:    false,
	}
	//连接时休眠1秒  防止刷新页面 先连接后退出
	//time.Sleep(time.Duration(1) * time.Second)
	client.hub.register <- client
	//var userActive = make(chan struct{})
	go client.ReadMsg()
	//go client.WriteMsg()
	/*go func() {
		d := 20 * time.Second
		timer := time.NewTimer(d)
		defer timer.Stop()
		for {
			select {
			case <-timer.C:
				fmt.Println("超时退出")
				client.Close()
				UnJoin(clientId)
				close(client.userActive)
				return // 返回并结束goroutine
			case <-client.userActive:
				fmt.Println("用户活跃")
				if !timer.Stop() {
					// 如果还没响应，说明到时间了，重置定时器，否则重新开始计时
					<-timer.C
				}
				timer.Reset(d)
			}
			time.Sleep(time.Millisecond * 1000)
		}
	}()*/
}
