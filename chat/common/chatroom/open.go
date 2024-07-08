package chatroom

import (
	"encoding/json"
	"fmt"
	"time"
)

func JoinChannelIds(uid string, channelIds ...string) error {
	if cli, ok := clients[uid]; ok {
		for _, channelId := range channelIds {
			if cli.hub.clients[channelId] == nil {
				cli.hub.clients[channelId] = map[*Client]bool{}
			}
			cli.hub.clients[channelId][cli] = true
		}
		cli.channelId = channelIds
		return nil
	} else {
		msgStore.ErrorLogServer(fmt.Errorf("用户`%s`未登录1", uid))
		return fmt.Errorf("用户`%s`未登录", uid)
	}
}

func UnJoinChannelIds(uid string, channelIds ...string) error {
	if cli, ok := clients[uid]; ok {
		delChannel := make(map[string]bool)
		for _, channelId := range channelIds {
			if cli.hub.clients[channelId] == nil {
				return fmt.Errorf("用户`%s`未监听：`%s` channel", uid, channelId)
			}

			if _, ok := cli.hub.clients[channelId][cli]; ok {
				delete(cli.hub.clients[channelId], cli)
				delChannel[channelId] = true
			}
			//如果频道的客户端都退出了，删除频道
			if len(cli.hub.clients[channelId]) == 0 {
				delete(cli.hub.clients, channelId)
			}
		}
		var hasChannelIds []string
		for _, v := range cli.channelId {
			if _, ok := delChannel[v]; !ok {
				hasChannelIds = append(hasChannelIds, v)
			}
		}
		cli.channelId = hasChannelIds
		return nil
	} else {
		msgStore.ErrorLogServer(fmt.Errorf("用户`%s`未登录2", uid))
		return fmt.Errorf("用户`%s`未登录", uid)
	}
}

func SendMessageToUid(uid, toUId string, msg string, tp uint8) {
	if cli, ok := clients[uid]; ok {
		message := Message{}
		if uid == toUId {
			json.Unmarshal([]byte(msg), &message)
			message = Message{
				Content:  msg,
				User:     cli.Id,
				Detail:   cli.Detail,
				ToUserId: toUId,
				Type:     tp,
				SendTime: time.Now().Format("2006-01-02 15:04:05"),
			}

		} else {
			message = Message{
				Content:  msg,
				User:     cli.Id,
				Detail:   cli.Detail,
				ToUserId: toUId,
				Type:     tp,
				SendTime: time.Now().Format("2006-01-02 15:04:05"),
			}
		}
		sendMessage(cli, message)
	}
}

func SendMessageToUids(uid string, msg string, tp uint8, toUIds ...string) {
	if cli, ok := clients[uid]; ok {
		message := Message{
			Content:  msg,
			User:     cli.Id,
			Detail:   cli.Detail,
			Type:     tp,
			SendTime: time.Now().Format("2006-01-02 15:04:05"),
		}
		for _, uid := range toUIds {
			message.ToUserId = uid
			sendMessage(cli, message)
		}
	}
}

func SendMessageToChannelIds(uid string, msg string, tp uint8, channelIds ...string) map[string]bool {
	resp := make(map[string]bool)
	if cli, ok := clients[uid]; ok {
		message := Message{
			Content:  msg,
			User:     uid,
			Detail:   cli.Detail,
			Type:     tp,
			SendTime: time.Now().Format("2006-01-02 15:04:05"),
		}

		for _, channelId := range channelIds {
			if !cli.hub.clients[channelId][cli] {
				resp[channelId] = false
				continue
			}
			message.ChannelId = channelId
			if sendMessage(cli, message) == nil {
				resp[channelId] = true
			} else {
				resp[channelId] = false
			}
		}
	}
	return resp
}

func sendMessage(c *Client, msg Message) error {
	c.hub.broadcast <- msg
	return nil
}

func SendAllMessage(uid string, message Message) {
	if cli, ok := clients[uid]; ok {

		message.User = cli.Id
		message.Detail = cli.Detail

		sendMessage(cli, message)
	}
}

func GetAllUser(uid, channelId string) (uidS []interface{}) {
	if cli, ok := clients[uid]; ok {
		if channel, ok2 := cli.hub.clients[channelId]; ok2 {
			for cliSon, _ := range channel {
				uidS = append(uidS, cliSon.Detail)
			}
		}
	}
	return
}

func Unregister(uid string) error {
	if cli, ok := clients[uid]; ok {
		cli.hub.unregister <- cli
		return nil
	} else {
		msgStore.ErrorLogServer(fmt.Errorf("用户`%s`已退出", uid))
		return fmt.Errorf("用户`%s`已退出", uid)
	}
	return nil
}

func UnJoin(uid string) {
	if cli, ok := clients[uid]; ok {
		for k, v := range cli.hub.clients {

			if _, ok1 := v[cli]; ok1 {
				delete(v, cli)
			}
			//如果频道的客户端都退出了，删除频道
			delete(cli.hub.clients, k)
		}
		delete(clients, uid)
	} else {
		msgStore.ErrorLogServer(fmt.Errorf("用户`%s`未登录3", uid))
	}
	return
}
