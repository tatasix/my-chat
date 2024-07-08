package im

import (
	"chat/common/chatroom"
)

func JoinChannelIds(uid string, channelIds ...string) error {
	return chatroom.JoinChannelIds(uid, channelIds...)
}

func SendMessageToUid(uid, toUId string, msg string, tp uint8) {
	chatroom.SendMessageToUid(uid, toUId, msg, tp)
}

func SendMessageToUids(uid string, msg string, tp uint8, toUIds ...string) {
	chatroom.SendMessageToUids(uid, msg, tp, toUIds...)
}

func SendMessageToChannelIds(uid string, msg string, tp uint8, channelIds ...string) {
	chatroom.SendMessageToChannelIds(uid, msg, tp, channelIds...)
}

func SendAllMessage(uid string, msg chatroom.Message) {
	chatroom.SendAllMessage(uid, msg)
}

func GetAllUser(uid, channelId string) (uidS []interface{}) {
	return chatroom.GetAllUser(uid, channelId)
}

func UnJoinChannelIds(uid string, channelIds ...string) error {
	return chatroom.UnJoinChannelIds(uid, channelIds...)
}

func Unregister(uid string) error {
	return chatroom.Unregister(uid)
}
