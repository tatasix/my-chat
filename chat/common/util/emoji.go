package util

import "strings"

func GetKeyFromEmojiMap(text string) int64 {
	emojiMap := getTheEmojiMap()
	for k, v := range emojiMap {
		if v == text {
			return k
		}
	}
	return 0
}

func getTheEmojiMap() map[int64]string {
	return map[int64]string{
		2:  "爱你",
		3:  "抱歉",
		4:  "不开心",
		5:  "吃惊",
		6:  "愤怒",
		7:  "鼓励",
		8:  "好奇",
		9:  "开心",
		10: "难过",
		11: "快乐",
		12: "困",
		14: "请求",
		15: "生气",
		16: "无聊",
		17: "拥抱",
		18: "着急",
	}

}

func GetTheEmojiString() string {
	emoji := getTheEmojiMap()
	var values []string
	for _, v := range emoji {
		values = append(values, v)
	}
	return strings.Join(values, ",")
}
