package util

import (
	"strings"
)

func TheTarotMap() map[uint32]string {
	// 定义一个map，key是数字，value是塔罗牌
	return map[uint32]string{
		0:  "世界",
		1:  "倒吊人",
		2:  "力量",
		3:  "命运之轮",
		4:  "圣杯七",
		5:  "圣杯三",
		6:  "圣杯九",
		7:  "圣杯二",
		8:  "圣杯五",
		9:  "圣杯侍从",
		10: "圣杯八",
		11: "圣杯六",
		12: "圣杯十",
		13: "圣杯四",
		14: "圣杯国王",
		15: "圣杯皇后",
		16: "圣杯首牌",
		17: "圣杯骑士",
		18: "太阳",
		19: "女皇",
		20: "女祭司",
		21: "宝剑七",
		22: "宝剑三",
		23: "宝剑九",
		24: "宝剑二",
		25: "宝剑五",
		26: "宝剑侍从",
		27: "宝剑八",
		28: "宝剑六",
		29: "宝剑十",
		30: "宝剑四",
		31: "宝剑国王",
		32: "宝剑皇后",
		33: "宝剑首牌",
		34: "宝剑骑士",
		35: "审判",
		36: "恋人",
		37: "愚人",
		38: "战车",
		39: "教皇",
		40: "星币七",
		41: "星币三",
		42: "星币九",
		43: "星币二",
		44: "星币五",
		45: "星币侍从",
		46: "星币八",
		47: "星币六",
		48: "星币十",
		49: "星币四",
		50: "星币国王",
		51: "星币皇后",
		52: "星币首牌",
		53: "星币骑士",
		54: "星星",
		55: "月亮",
		56: "权杖七",
		57: "权杖三",
		58: "权杖九",
		59: "权杖二",
		60: "权杖五",
		61: "权杖侍从",
		62: "权杖八",
		63: "权杖六",
		64: "权杖十",
		65: "权杖四",
		66: "权杖国王",
		67: "权杖手牌",
		68: "权杖皇后",
		69: "权杖骑士",
		70: "正义",
		71: "死神",
		72: "皇帝",
		73: "节制",
		74: "隐士",
		75: "高塔",
		76: "魔术师",
		77: "魔鬼",
	}

}

func TheTarotImage() map[uint32]string {
	// 定义一个map，key是数字，value是塔罗牌
	return map[uint32]string{
		0:  "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105917492891648世界.png",
		1:  "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105919367745536倒吊人.png",
		2:  "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105920009474048力量.png",
		3:  "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105920756060160命运之轮.png",
		4:  "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105921582338048圣杯七.png",
		5:  "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105922408615936圣杯三.png",
		6:  "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105923083898880圣杯九.png",
		7:  "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105923624964096圣杯二.png",
		8:  "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105924367355904圣杯五.png",
		9:  "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105925453680640圣杯侍从.png",
		10: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105926808440832圣杯八.png",
		11: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105927441780736圣杯六.png",
		12: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105929551515648圣杯十.png",
		13: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105930155495424圣杯四.png",
		14: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105931799662592圣杯国王.png",
		15: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105933972312064圣杯皇后.png",
		16: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105936132378624圣杯首牌.png",
		17: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105938342776832圣杯骑士.png",
		18: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105939009671168太阳.png",
		19: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105942750990336女皇.png",
		20: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105943426273280女祭司.png",
		21: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105947310198784宝剑七.png",
		22: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105949621260288宝剑三.png",
		23: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105953165447168宝剑九.png",
		24: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105955140964352宝剑二.png",
		25: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105956088877056宝剑五.png",
		26: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105957359751168宝剑侍从.png",
		27: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105958232166400宝剑八.png",
		28: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105960635502592宝剑六.png",
		29: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105962548105216宝剑十.png",
		30: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105964796252160宝剑四.png",
		31: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105965282791424宝剑国王.png",
		32: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105966134235136宝剑皇后.png",
		33: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105966968901632宝剑首牌.png",
		34: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105968336244736宝剑骑士.png",
		35: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105968990556160审判.png",
		36: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105969732947968恋人.png",
		37: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105970555031552愚人.png",
		38: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105971087708160战车.png",
		39: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105971549081600教皇.png",
		40: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105972740263936星币七.png",
		41: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105973558153216星币三.png",
		42: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105974023720960星币九.png",
		43: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105974929690624星币二.png",
		44: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105977144283136星币五.png",
		45: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105977844731904星币侍从.png",
		46: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105979388235776星币八.png",
		47: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105980592001024星币六.png",
		48: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105982139699200星币十.png",
		49: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105982735290368星币四.png",
		50: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105983754506240星币国王.png",
		51: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105984295571456星币皇后.png",
		52: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105984828248064星币首牌.png",
		53: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105985927155712星币骑士.png",
		54: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105986505969664星星.png",
		55: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105986967343104月亮.png",
		56: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105988137553920权杖七.png",
		57: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105988598927360权杖三.png",
		58: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105989219684352权杖九.png",
		59: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105989769138176权杖二.png",
		60: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105990364729344权杖五.png",
		61: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105990972903424权杖侍从.png",
		62: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105991430082560权杖八.png",
		63: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105992524795904权杖六.png",
		64: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105993040695296权杖十.png",
		65: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105993556594688权杖四.png",
		66: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105994227683328权杖国王.png",
		67: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105995251093504权杖手牌.png",
		68: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105995699884032权杖皇后.png",
		69: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105997750898688权杖骑士.png",
		70: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692105999940325376正义.png",
		71: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692106000573665280死神.png",
		72: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692106002263969792皇帝.png",
		73: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692106004344344576节制.png",
		74: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692106005095124992隐士.png",
		75: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692106006181449728高塔.png",
		76: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692106006928035840魔术师.png",
		77: "https://miaostatic.pansi.com/image/common/tarot/20230817/1692106008345710592魔鬼.png",
	}

}

func GetTheTarotString() string {
	emoji := TheTarotMap()
	var values []string
	for _, v := range emoji {
		values = append(values, v)
	}
	return strings.Join(values, ",")
}

func GetKeyFromTarotMap(text string) uint32 {
	emojiMap := TheTarotMap()
	for k, v := range emojiMap {
		if v == text {
			return k
		}
	}
	return 0
}

func GetTarotName(key uint32) string {
	emojiMap := TheTarotMap()
	if k, ok := emojiMap[key]; ok {
		return k
	}
	return ""
}

func GetTarotImage(key uint32) string {
	emojiMap := TheTarotImage()
	if k, ok := emojiMap[key]; ok {
		return k
	}
	return ""
}
