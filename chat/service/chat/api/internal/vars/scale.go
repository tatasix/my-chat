package vars

type Scale struct {
	OpenKfId string
	Info     string
	Name     string
}

func GetTheBaseInfo3() []Scale {
	return []Scale{}
}

func GetTheBaseInfo() []Scale {
	result := []Scale{
		{
			Name:     "抑郁体验问卷",
			OpenKfId: "aEKcuGDK2cV0xUmm5uLk5X1FFIYNDYlb",
			Info: `[
    {
        "QuestionNumber": 1,
        "Title": "我尽可能商地为自己设定目标",
        "type": 3,
        "score_letter": 4,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 2,
        "Title": "没有周围人的支持，我将会感到孤立无援",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 7,
        "Title": "我常发觉自己不能按自己的标准或理想行事",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 3,
        "Title": "我容易满足于目前的计划和目标，从不去追求更高的目标",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 4,
        "Title": "我有时觉得自己很高大，有时却又觉得自己很渺小",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 5,
        "Title": "当我与别人形成了密切的关系，我从来没有唯恐失去的感觉",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 6,
        "Title": "我迫切需要只有别人才能提供的东西",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 8,
        "Title": "我感到我总能充分发挥自己的潜能",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 9,
        "Title": "与人缺少长久的关系并不让我优虑",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 10,
        "Title": "如果不能达到自己的期望，我会觉得没有价值",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 11,
        "Title": "许多时候我觉得孤立无援",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 12,
        "Title": "我很少担心自己的言行会遭到非议",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 13,
        "Title": "在我目前的状况与我的希望之间有相当大的距离",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 14,
        "Title": "我在激烈的竞争中感到快乐",
        "type": 3,
        "score_letter": 4,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 15,
        "Title": "我觉得有许多责任必须承担",
        "type": 3,
        "score_letter": 4,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 16,
        "Title": "我有时感到内心“空虚”",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 17,
        "Title": "我不易满足于现状",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 18,
        "Title": "我不在乎是否达到了别人的要求",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 19,
        "Title": "当感到寂奥时我会变得恐慌",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 20,
        "Title": "假如失去一个很亲密的朋友，我会觉得好象是失去了自身的某个重要的部分",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 21,
        "Title": "不管我犯过多少错误人们都不会将我拒之门外",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 22,
        "Title": "我难于中断使我不愉快的关系",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 23,
        "Title": "我常常担心会有失去亲密朋友的危险",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 24,
        "Title": "别人对我要求太高",
        "type": 3,
        "score_letter": 4,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 25,
        "Title": "跟别人一道时，我容易低估或“贱卖”自己",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 26,
        "Title": "我不大在乎别人怎样报答",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 27,
        "Title": "两个人的关系不管多么亲密，仍会有摩擦和冲突",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 28,
        "Title": "我对被别人拒绝的暗示非常敏感",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 29,
        "Title": "我的成功对家庭很重要",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 30,
        "Title": "我常常觉得自己令人失望",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 31,
        "Title": "当别人惹我发火时，我会让他(她)知道我的感受",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 32,
        "Title": "我持之以恒、不遗余力地取悦或帮助周围的人",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 33,
        "Title": "我精力(能力、力量)充沛",
        "type": 3,
        "score_letter": 4,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 34,
        "Title": "我发觉很难对朋友的请求说“不”",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 35,
        "Title": "在一种密切的关系中我绝不会真正感到安全",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 36,
        "Title": "我对自己的看法常常改变，有时感到自己完美无缺，有时看到自己的不足又觉得自己一无是处",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 37,
        "Title": "我常因处境改变而恐惧",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 38,
        "Title": "即便最亲近的人即将离去，我也一样能自己生活下去",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 39,
        "Title": "人们必须坚持不懈地追求他人的爱：也就是说，爱必须争取",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 40,
        "Title": "我对他人对自己言行的感受特别敏感",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 41,
        "Title": "我常因自己的言行而内疚",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 42,
        "Title": "我是一个独立性强的人",
        "type": 3,
        "score_letter": 4,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 43,
        "Title": "我常感到有罪",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 44,
        "Title": "我想我是一个很复杂的人，一个具有“多种侧面”的人",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 45,
        "Title": "我十分担心会冒犯或伤害我所亲近的人",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 46,
        "Title": "发怒会使我惊慌失措",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 47,
        "Title": "重要的不是你的身份而是你所取得的成就",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 48,
        "Title": "不管成功还是失败，我都感觉良好",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 49,
        "Title": "我很容易把自己的感受和问题放到一边，全身心地关心别人的感受与问题",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 50,
        "Title": "假如一个我所关心的人冲我发火，我将担心他会离我而去",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 51,
        "Title": "当要担负重要责任时，我会感到不自在",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 52,
        "Title": "与朋友吵架后，我必须尽快承认错误",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 53,
        "Title": "我不愿承认自身的弱点",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 54,
        "Title": "重要的是我喜欢自己的工作，而不是我的工作是否得到称赞",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 55,
        "Title": "与人争吵后，我会感到非常孤独",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 56,
        "Title": "在与别人的交往中，我很注意别人能给我什么",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 57,
        "Title": "我很少想到我的家庭",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 58,
        "Title": "我对亲友的感受时常改变：有时感到怒发冲冠，有时却又柔情似水，情意绵绵",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 59,
        "Title": "我的言行对周围的人影响很大",
        "type": 3,
        "score_letter": 4,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 60,
        "Title": "我有时感到自己“特别气”",
        "type": 3,
        "score_letter": 4,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 61,
        "Title": "我成长在一个极端封闭的家庭中",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 62,
        "Title": "我对自己和自己的成就十分满意",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 63,
        "Title": "我希望能从亲友那儿得到许多东西",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 64,
        "Title": "我倾向于对自己过分严厉",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 65,
        "Title": "独自呆着一点也不令我心烦",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 66,
        "Title": "我经常用准则或目标来对照自己",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "比较反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "稍微反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "既不反对也不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "稍微同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "比较同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "完全同意",
                "score": 7,
                "id": 7
            }
        ]
    }
]`,
		},
		{
			Name:     "防御方式问卷",
			OpenKfId: "YTBNxzLNkCrpxZyCQkUmTCOa4Dk9QCXl",
			Info: `[
    {
        "QuestionNumber": 1,
        "Title": "我因帮助他人而获得满足，如果不这样做，我就会变得情绪抑郁",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 4,
        "Title": "人们总是不公平地对待我",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 2,
        "Title": "人们常说我是个脾气暴躁的人",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 7,
        "Title": "我不知道为什么会遇到相同的受挫情境",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 3,
        "Title": "在我没有时间处理某个棘手的事情时，我可以把它搁置一边",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 9,
        "Title": "我受到挫折时，表现就象个孩子",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 5,
        "Title": "我通过做一些积极的或创见性的事情来摆脱自己的焦虑不安，如绘画、做木工活等",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 8,
        "Title": "我能够相当轻松地嘲笑我自己",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 6,
        "Title": "偶尔，我把一些今天该做的事情推迟到明天再做",
        "type": 3,
        "score_letter": 4,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 10,
        "Title": "在维护我的利益方面，我羞于与人计较",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 11,
        "Title": "我比我认识人中的大多数都强",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 12,
        "Title": "人们往往虐待我",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 13,
        "Title": "如果某人骗了我或偷了我的钱，我宁愿他得到帮助，而不是受惩罚",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 14,
        "Title": "偶尔，我想一些坏得不能说出口的事情",
        "type": 3,
        "score_letter": 4,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 15,
        "Title": "偶尔，我因一些下流的笑话而大笑",
        "type": 3,
        "score_letter": 4,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 16,
        "Title": "人们说我象一只驼鸟，把自己的头埋入沙中，换句话说，我往往有意忽视一些不愉快的事情",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 17,
        "Title": "我常常不能竭尽全力地与人竞争",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 18,
        "Title": "我常感到我比和我在一起的人强",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 19,
        "Title": "某人正在想剥夺我所得到的一切",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 20,
        "Title": "我有时发怒",
        "type": 3,
        "score_letter": 4,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 21,
        "Title": "我时常在某种内在力量的驱使下，不由自主地做出些行为",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 22,
        "Title": "我宁愿饿死而不愿被迫吃饭",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 23,
        "Title": "我常常故意忽视一些危险，似乎我是个超人",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 24,
        "Title": "我以有贬低别人威望的能力而自豪",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 25,
        "Title": "人们告诉我：我总有被害的感觉",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 26,
        "Title": "有时感觉不好时，我就发脾气",
        "type": 3,
        "score_letter": 4,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 27,
        "Title": "当某些事情使我烦脑时，我常常不由自主地做出些行为",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 28,
        "Title": "当遇事不顺心时，我就会生病",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 29,
        "Title": "我是一个很有自制力的人",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 30,
        "Title": "我简直就像一个不得志的艺术家一样",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 31,
        "Title": "我不总是说真话",
        "type": 3,
        "score_letter": 4,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 32,
        "Title": "当我感到自尊心受伤害时，我就会回避",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 33,
        "Title": "我常常不由自主地迫使自己干些过头的事情，以至于其他人不得不限制我",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 34,
        "Title": "我的朋友们把我看做乡下佬",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 35,
        "Title": "在我愤怒的时候，我常常回避",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 36,
        "Title": "我往往对那些确实对我友好的人，比我应该怀疑的人保持更高的警惕性",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 37,
        "Title": "我已学得特殊的才能，足以使我毫无问题地渡过一生",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 38,
        "Title": "有时，在选举的时候，我往往选那些我几乎不了解的人",
        "type": 3,
        "score_letter": 4,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 39,
        "Title": "我常常不能按时赴约",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 40,
        "Title": "我幻想的多，可在现实生活中做的少",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 41,
        "Title": "我羞于与人打交道",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 42,
        "Title": "我什么都不怕",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 43,
        "Title": "有时我认为我是个天使，有时我认为我是个恶魔",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 44,
        "Title": "在比赛时，我宁要赢而不愿输",
        "type": 3,
        "score_letter": 4,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 45,
        "Title": "在我愤怒的时候，我变得很愿挖苦人",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 46,
        "Title": "在我自尊心受伤害时，我就公开还击",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 47,
        "Title": "我认为当我受伤害时，我就应该翻脸",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 48,
        "Title": "我每天读报时，不是每个版面都读",
        "type": 3,
        "score_letter": 4,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 49,
        "Title": "我沮丧时，就会避开",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 50,
        "Title": "我对性问题感到害羞",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 51,
        "Title": "我总是感到我所认识的某个人象个保护神",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 52,
        "Title": "我的处世哲学是：“非理勿信，非理勿做，非理勿视”",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 53,
        "Title": "我认为：人有好坏之分",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 54,
        "Title": "如果我的上司惹我生气，我可能会在工作中找麻烦或磨洋工，以报复他",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 55,
        "Title": "每个人都和我对着干",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 56,
        "Title": "我往往对那些我讨厌的人而表示友好",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 57,
        "Title": "如果我乘坐的飞机一个发动机失灵，我就会非常紧张",
        "type": 3,
        "score_letter": 4,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 58,
        "Title": "我认识这样一个人，他什么都能做而且做得合理正直",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 59,
        "Title": "如果我感情的发泄会防碍我正从事的事业，那么我就能控制住它",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 60,
        "Title": "一些人正在密谋要害我",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 61,
        "Title": "我通常可以看到恶境当中的好的一面",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 62,
        "Title": "在我不得不去做一些我不愿做的事情时，就头痛",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 63,
        "Title": "我常常发现我对那些理应仇视的人，表示很友好",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 64,
        "Title": "我认为：“人人都有善意”是不存在的，如果你不好，那么你一切都不好",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 65,
        "Title": "我决不会对那些我讨厌的人表示愤怒",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 66,
        "Title": "我确信生活对我是不公正的",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 67,
        "Title": "在严重的打击下，我就会垮下来",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 68,
        "Title": "在我意识到不得不面临一场困境的时候，如考试，招工会谈。我就试图想像它会如何，并计划出一些方法去应付它",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 69,
        "Title": "医生们决不会真的弄清我患的是什么病",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 70,
        "Title": "当某个和我很亲近的人死去时，我并不悲伤",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 71,
        "Title": "在我为了利益和人争斗之后，我往往因为我的粗鲁而向人道歉",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 72,
        "Title": "发生与我有关的大部分事情并不是我的责任",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 73,
        "Title": "当我感觉情绪压抑或焦虑不安时，吃点东西，可以使我感觉好些",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 74,
        "Title": "勤奋工作使我感觉好些",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 75,
        "Title": "医生不能真的帮我解决问题",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 76,
        "Title": "我常听人们说我不暴露自己的感情",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 77,
        "Title": "我认为，人们在看电影，戏剧或书籍对所领悟的意义，比这些作品所要表达的意义要多",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 78,
        "Title": "我感觉到我有一些不由自主要去做的习惯或仪式行为，并给我带来很多麻烦",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 79,
        "Title": "当我紧张时，就喝酒或吃药",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 80,
        "Title": "当我心情不愉快时，就想和别人呆在一起",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 81,
        "Title": "如果我能够预感到我会沮丧的话，我就能更好地应付它",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 82,
        "Title": "无论我怎样发牢骚，从未得到过满意的结果",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 83,
        "Title": "我常常发现当环境要引起我强烈的情绪反应时，我就会麻木不仁",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 84,
        "Title": "忘我地工作，可使我摆脱情绪上的忧郁和焦虑",
        "type": 3,
        "score_letter": 2,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 85,
        "Title": "紧张的时候，我就吸烟",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 86,
        "Title": "如果我陷入某种危机时，我就会寻找另一个和我具有同样命运的人",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 87,
        "Title": "如我做错了事情，不能受责备",
        "type": 3,
        "score_letter": 1,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    },
    {
        "QuestionNumber": 88,
        "Title": "如果我有攻击他人的想法，我就感觉有种做点事情的需要，以转移这种想法",
        "type": 3,
        "score_letter": 3,
        "answers": [
            {
                "content": "完全反对",
                "score": 1,
                "id": 1
            },
            {
                "content": "很反对",
                "score": 2,
                "id": 2
            },
            {
                "content": "比较反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "稍微反对",
                "score": 4,
                "id": 4
            },
            {
                "content": "既不反对也不同意",
                "score": 5,
                "id": 5
            },
            {
                "content": "稍微同意",
                "score": 6,
                "id": 6
            },
            {
                "content": "比较同意",
                "score": 7,
                "id": 7
            },
            {
                "content": "很同意",
                "score": 8,
                "id": 8
            },
            {
                "content": "完全同意",
                "score": 9,
                "id": 9
            }
        ]
    }
]`,
		},
	}
	return result
}

func GetTheBaseInfo1() []Scale {
	result := []Scale{
		{
			Name:     "蒙哥马利抑郁评估",
			OpenKfId: "UB5OWvU4XufQVjOy3rhj3TgAEm17ujdd",
			Info: `[
    {
        "QuestionNumber": 1,
        "Title": "结合您最近一周的情绪做出选择，可观察到的抑郁",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 0,
                "id": 1
            },
            {
                "content": "介于上下两种选项之间",
                "score": 1,
                "id": 2
            },
            {
                "content": "看起来是悲伤的，但能使之高兴一些",
                "score": 2,
                "id": 3
            },
            {
                "content": "介于上下两种选项之间",
                "score": 3,
                "id": 4
            },
            {
                "content": "突出的悲伤忧郁，但其情绪仍可受外界环境影响",
                "score": 4,
                "id": 5
            },
            {
                "content": "介于上下两种选项之间",
                "score": 5,
                "id": 6
            },
            {
                "content": "整天抑郁，极度严重",
                "score": 6,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 2,
        "Title": "结合您最近一周的情绪做出选择，抑郁主诉",
        "type": 2,
        "answers": [
            {
                "content": "在日常心境中偶有抑郁",
                "score": 0,
                "id": 1
            },
            {
                "content": "介于上下两种选项之间",
                "score": 1,
                "id": 2
            },
            {
                "content": "有抑郁或情绪低沉，但可使之愉快些",
                "score": 2,
                "id": 3
            },
            {
                "content": "介于上下两种选项之间",
                "score": 3,
                "id": 4
            },
            {
                "content": "沉湎于抑郁沮丧心境,但环境仍可对心境有些影响",
                "score": 4,
                "id": 5
            },
            {
                "content": "介于上下两种选项之间",
                "score": 5,
                "id": 6
            },
            {
                "content": "持久不断的深度抑郁沮丧",
                "score": 6,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 3,
        "Title": "结合您最近一周的情绪做出选择，内心紧张",
        "type": 2,
        "answers": [
            {
                "content": "平静，偶有瞬间的紧张",
                "score": 0,
                "id": 1
            },
            {
                "content": "介于上下两种选项之间",
                "score": 1,
                "id": 2
            },
            {
                "content": "偶有紧张不安及难以言明的不舒服感",
                "score": 2,
                "id": 3
            },
            {
                "content": "介于上下两种选项之间",
                "score": 3,
                "id": 4
            },
            {
                "content": "持久的内心紧张，或间歇呈现的恐惧状态，要花费 相当努力方能克制",
                "score": 4,
                "id": 5
            },
            {
                "content": "介于上下两种选项之间",
                "score": 5,
                "id": 6
            },
            {
                "content": "持续的恐惧和苦恼，极度惊恐",
                "score": 6,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 4,
        "Title": "结合您最近一周的情绪做出选择，睡眠减少",
        "type": 2,
        "answers": [
            {
                "content": "睡眠如常",
                "score": 0,
                "id": 1
            },
            {
                "content": "介于上下两种选项之间",
                "score": 1,
                "id": 2
            },
            {
                "content": "轻度入睡困难，或睡眠较浅，或时睡时醒",
                "score": 2,
                "id": 3
            },
            {
                "content": "介于上下两种选项之间",
                "score": 3,
                "id": 4
            },
            {
                "content": "睡眠减少或睡眠中断2小时以上",
                "score": 4,
                "id": 5
            },
            {
                "content": "介于上下两种选项之间",
                "score": 5,
                "id": 6
            },
            {
                "content": "每天睡眠总时间不超过2－3小时",
                "score": 6,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 5,
        "Title": "结合您最近一周的情绪做出选择，食欲减退",
        "type": 2,
        "answers": [
            {
                "content": "食欲正常或增进",
                "score": 0,
                "id": 1
            },
            {
                "content": "介于上下两种选项之间",
                "score": 1,
                "id": 2
            },
            {
                "content": "轻度食欲减退",
                "score": 2,
                "id": 3
            },
            {
                "content": "介于上下两种选项之间",
                "score": 3,
                "id": 4
            },
            {
                "content": "没有食欲，食而无味",
                "score": 4,
                "id": 5
            },
            {
                "content": "介于上下两种选项之间",
                "score": 5,
                "id": 6
            },
            {
                "content": "不愿进食，需他人帮助",
                "score": 6,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 6,
        "Title": "结合您最近一周的情绪做出选择，思想集中困难",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 0,
                "id": 1
            },
            {
                "content": "介于上下两种选项之间",
                "score": 1,
                "id": 2
            },
            {
                "content": "偶有思想集中困难",
                "score": 2,
                "id": 3
            },
            {
                "content": "介于上下两种选项之间",
                "score": 3,
                "id": 4
            },
            {
                "content": "思想难以集中，以致干扰阅读或交谈",
                "score": 4,
                "id": 5
            },
            {
                "content": "介于上下两种选项之间",
                "score": 5,
                "id": 6
            },
            {
                "content": "完全不能集中思想，无法阅读",
                "score": 6,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 7,
        "Title": "结合您最近一周的情绪做出选择，懒散",
        "type": 2,
        "answers": [
            {
                "content": "活动发动并无困难，动作不慢",
                "score": 0,
                "id": 1
            },
            {
                "content": "介于上下两种选项之间",
                "score": 1,
                "id": 2
            },
            {
                "content": "有始动困难",
                "score": 2,
                "id": 3
            },
            {
                "content": "介于上下两种选项之间",
                "score": 3,
                "id": 4
            },
            {
                "content": "即使简单的日常活动也难以发动，需花很大努力",
                "score": 4,
                "id": 5
            },
            {
                "content": "介于上下两种选项之间",
                "score": 5,
                "id": 6
            },
            {
                "content": "完全呈懒散状态，无人帮助什么也干不了",
                "score": 6,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 8,
        "Title": "结合您最近一周的情绪做出选择，感受力衰退",
        "type": 2,
        "answers": [
            {
                "content": "对周围的人和物的兴趣正常",
                "score": 0,
                "id": 1
            },
            {
                "content": "介于上下两种选项之间",
                "score": 1,
                "id": 2
            },
            {
                "content": "对日常趣事的享受减退",
                "score": 2,
                "id": 3
            },
            {
                "content": "介于上下两种选项之间",
                "score": 3,
                "id": 4
            },
            {
                "content": "对周围不感兴趣，对朋友和熟人缺乏感情",
                "score": 4,
                "id": 5
            },
            {
                "content": "介于上下两种选项之间",
                "score": 5,
                "id": 6
            },
            {
                "content": "呈情感麻木状态，不能体验愤怒、悲痛和愉快，对亲友全无感情",
                "score": 6,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 9,
        "Title": "结合您最近一周的情绪做出选择，悲观思想",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 0,
                "id": 1
            },
            {
                "content": "介于上下两种选项之间",
                "score": 1,
                "id": 2
            },
            {
                "content": "时有时无的失败，自责和自卑感",
                "score": 2,
                "id": 3
            },
            {
                "content": "介于上下两种选项之间",
                "score": 3,
                "id": 4
            },
            {
                "content": "持久的自责或肯定的但尚近情理的自罪，对前途悲观",
                "score": 4,
                "id": 5
            },
            {
                "content": "介于上下两种选项之间",
                "score": 5,
                "id": 6
            },
            {
                "content": "自我毁灭、自我悔恨或感罪恶深重的妄想，荒谬绝伦、难以动摇的自我谴责",
                "score": 6,
                "id": 7
            }
        ]
    },
    {
        "QuestionNumber": 10,
        "Title": "结合您最近一周的情绪做出选择，自杀观念",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 0,
                "id": 1
            },
            {
                "content": "介于上下两种选项之间",
                "score": 1,
                "id": 2
            },
            {
                "content": "对生活厌倦，偶有瞬间即逝的自杀念头",
                "score": 2,
                "id": 3
            },
            {
                "content": "介于上下两种选项之间",
                "score": 3,
                "id": 4
            },
            {
                "content": "感到不如死了的好，常有自杀念头，认为自杀是一 种可能的自我解决的方法，但尚无切实的自杀计划",
                "score": 4,
                "id": 5
            },
            {
                "content": "介于上下两种选项之间",
                "score": 5,
                "id": 6
            },
            {
                "content": "已拟适合时机的自杀计划，并积极准备",
                "score": 6,
                "id": 7
            }
        ]
    }
]`,
		},
		{
			Name:     "孤独感测试",
			OpenKfId: "FnyZNS9uIkw81dqrBRqeiwWyQyjcr1Sc",
			Info: `[
    {
        "QuestionNumber": 1,
        "Title": "你觉得和周围的人相处融洽，有“物以类聚”之感",
        "type": 2,
        "answers": [
            {
                "content": "经常",
                "score": 1,
                "id": 1
            },
            {
                "content": "有时",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少",
                "score": 3,
                "id": 3
            },
            {
                "content": "从不",
                "score": 4,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 2,
        "Title": "你觉得缺个伴儿",
        "type": 2,
        "answers": [
            {
                "content": "经常",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时",
                "score": 3,
                "id": 2
            },
            {
                "content": "很少",
                "score": 2,
                "id": 3
            },
            {
                "content": "从不",
                "score": 1,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 3,
        "Title": "你觉得没人可以求助、分享或依靠",
        "type": 2,
        "answers": [
            {
                "content": "经常",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时",
                "score": 3,
                "id": 2
            },
            {
                "content": "很少",
                "score": 2,
                "id": 3
            },
            {
                "content": "从不",
                "score": 1,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 4,
        "Title": "你觉得孤单",
        "type": 2,
        "answers": [
            {
                "content": "经常",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时",
                "score": 3,
                "id": 2
            },
            {
                "content": "很少",
                "score": 2,
                "id": 3
            },
            {
                "content": "从不",
                "score": 1,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 5,
        "Title": "你觉得是朋友群中的一员",
        "type": 2,
        "answers": [
            {
                "content": "经常",
                "score": 1,
                "id": 1
            },
            {
                "content": "有时",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少",
                "score": 3,
                "id": 3
            },
            {
                "content": "从不",
                "score": 4,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 9,
        "Title": "你觉得自己外向而友好",
        "type": 2,
        "answers": [
            {
                "content": "经常",
                "score": 1,
                "id": 1
            },
            {
                "content": "有时",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少",
                "score": 3,
                "id": 3
            },
            {
                "content": "从不",
                "score": 4,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 8,
        "Title": "你觉得你不能和周遭的人分享自己的兴趣和想法",
        "type": 2,
        "answers": [
            {
                "content": "经常",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时",
                "score": 3,
                "id": 2
            },
            {
                "content": "很少",
                "score": 2,
                "id": 3
            },
            {
                "content": "从不",
                "score": 1,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 7,
        "Title": "你觉得和任何人都不再亲近了",
        "type": 2,
        "answers": [
            {
                "content": "经常",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时",
                "score": 3,
                "id": 2
            },
            {
                "content": "很少",
                "score": 2,
                "id": 3
            },
            {
                "content": "从不",
                "score": 1,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 6,
        "Title": "你觉得和身边的人有很多共同点",
        "type": 2,
        "answers": [
            {
                "content": "经常",
                "score": 1,
                "id": 1
            },
            {
                "content": "有时",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少",
                "score": 3,
                "id": 3
            },
            {
                "content": "从不",
                "score": 4,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 10,
        "Title": "你觉得和别人很亲近",
        "type": 2,
        "answers": [
            {
                "content": "经常",
                "score": 1,
                "id": 1
            },
            {
                "content": "有时",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少",
                "score": 3,
                "id": 3
            },
            {
                "content": "从不",
                "score": 4,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 11,
        "Title": "你觉得自己遭人冷落",
        "type": 2,
        "answers": [
            {
                "content": "经常",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时",
                "score": 3,
                "id": 2
            },
            {
                "content": "很少",
                "score": 2,
                "id": 3
            },
            {
                "content": "从不",
                "score": 1,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 12,
        "Title": "你觉得自己和别人的交往没有意义",
        "type": 2,
        "answers": [
            {
                "content": "经常",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时",
                "score": 3,
                "id": 2
            },
            {
                "content": "很少",
                "score": 2,
                "id": 3
            },
            {
                "content": "从不",
                "score": 1,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 13,
        "Title": "你觉得没人真的了解你",
        "type": 2,
        "answers": [
            {
                "content": "经常",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时",
                "score": 3,
                "id": 2
            },
            {
                "content": "很少",
                "score": 2,
                "id": 3
            },
            {
                "content": "从不",
                "score": 1,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 14,
        "Title": "你觉得自己与他人隔绝了",
        "type": 2,
        "answers": [
            {
                "content": "经常",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时",
                "score": 3,
                "id": 2
            },
            {
                "content": "很少",
                "score": 2,
                "id": 3
            },
            {
                "content": "从不",
                "score": 1,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 15,
        "Title": "你觉得如果你想，就一定能找个伴儿",
        "type": 2,
        "answers": [
            {
                "content": "经常",
                "score": 1,
                "id": 1
            },
            {
                "content": "有时",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少",
                "score": 3,
                "id": 3
            },
            {
                "content": "从不",
                "score": 4,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 16,
        "Title": "你觉得还是有人真正理解你",
        "type": 2,
        "answers": [
            {
                "content": "经常",
                "score": 1,
                "id": 1
            },
            {
                "content": "有时",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少",
                "score": 3,
                "id": 3
            },
            {
                "content": "从不",
                "score": 4,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 17,
        "Title": "你觉得害羞",
        "type": 2,
        "answers": [
            {
                "content": "经常",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时",
                "score": 3,
                "id": 2
            },
            {
                "content": "很少",
                "score": 2,
                "id": 3
            },
            {
                "content": "从不",
                "score": 1,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 18,
        "Title": "你觉得你身边虽然有人，但他们却没真正和你在一起",
        "type": 2,
        "answers": [
            {
                "content": "经常",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时",
                "score": 3,
                "id": 2
            },
            {
                "content": "很少",
                "score": 2,
                "id": 3
            },
            {
                "content": "从不",
                "score": 1,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 19,
        "Title": "你觉得还是有人可以说说话",
        "type": 2,
        "answers": [
            {
                "content": "经常",
                "score": 1,
                "id": 1
            },
            {
                "content": "有时",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少",
                "score": 3,
                "id": 3
            },
            {
                "content": "从不",
                "score": 4,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 20,
        "Title": "你觉得还是有人可以求助、分享或依靠",
        "type": 2,
        "answers": [
            {
                "content": "经常",
                "score": 1,
                "id": 1
            },
            {
                "content": "有时",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少",
                "score": 3,
                "id": 3
            },
            {
                "content": "从不",
                "score": 4,
                "id": 4
            }
        ]
    }
]`,
		},
		{
			Name:     "焦虑自测",
			OpenKfId: "8HPjI81RCupXimX33iNWgMRGHKzlaaNB",
			Info: `[
    {
        "QuestionNumber": 1,
        "Title": "我感到比往常更加神经过敏的焦虑",
        "type": 2,
        "answers": [
            {
                "content": "绝大部分时间有",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时有",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少有",
                "score": 1,
                "id": 3
            },
            {
                "content": "大部分时间有",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 2,
        "Title": "我无缘无故感到担心",
        "type": 2,
        "answers": [
            {
                "content": "绝大部分时间有",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时有",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少有",
                "score": 1,
                "id": 3
            },
            {
                "content": "大部分时间有",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 3,
        "Title": "我容易心烦意乱或感到恐慌",
        "type": 2,
        "answers": [
            {
                "content": "绝大部分时间有",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时有",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少有",
                "score": 1,
                "id": 3
            },
            {
                "content": "大部分时间有",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 4,
        "Title": "我感到我的身体好像被分成几块，支离破碎",
        "type": 2,
        "answers": [
            {
                "content": "绝大部分时间有",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时有",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少有",
                "score": 1,
                "id": 3
            },
            {
                "content": "大部分时间有",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 5,
        "Title": "我感到事事都很顺利，不会有倒霉的事情发生",
        "type": 2,
        "answers": [
            {
                "content": "绝大部分时间有",
                "score": 1,
                "id": 1
            },
            {
                "content": "有时有",
                "score": 3,
                "id": 2
            },
            {
                "content": "很少有",
                "score": 4,
                "id": 3
            },
            {
                "content": "大部分时间有",
                "score": 2,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 6,
        "Title": "我的四肢拌动和震颤",
        "type": 2,
        "answers": [
            {
                "content": "绝大部分时间有",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时有",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少有",
                "score": 1,
                "id": 3
            },
            {
                "content": "大部分时间有",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 7,
        "Title": "我因头痛、颈痛和背痛而烦恼",
        "type": 2,
        "answers": [
            {
                "content": "绝大部分时间有",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时有",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少有",
                "score": 1,
                "id": 3
            },
            {
                "content": "大部分时间有",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 8,
        "Title": "我感到无力而且容易疲劳",
        "type": 2,
        "answers": [
            {
                "content": "绝大部分时间有",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时有",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少有",
                "score": 1,
                "id": 3
            },
            {
                "content": "大部分时间有",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 9,
        "Title": "我感到平静，能安静坐下来",
        "type": 2,
        "answers": [
            {
                "content": "绝大部分时间有",
                "score": 1,
                "id": 1
            },
            {
                "content": "有时有",
                "score": 3,
                "id": 2
            },
            {
                "content": "很少有",
                "score": 4,
                "id": 3
            },
            {
                "content": "大部分时间有",
                "score": 2,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 10,
        "Title": "我感到我的心跳较快",
        "type": 2,
        "answers": [
            {
                "content": "绝大部分时间有",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时有",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少有",
                "score": 1,
                "id": 3
            },
            {
                "content": "大部分时间有",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 11,
        "Title": "我因阵阵的眩晕而不舒服",
        "type": 2,
        "answers": [
            {
                "content": "绝大部分时间有",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时有",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少有",
                "score": 1,
                "id": 3
            },
            {
                "content": "大部分时间有",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 12,
        "Title": "我有阵阵要晕倒的感觉",
        "type": 2,
        "answers": [
            {
                "content": "绝大部分时间有",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时有",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少有",
                "score": 1,
                "id": 3
            },
            {
                "content": "大部分时间有",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 13,
        "Title": "我呼吸时进气和出气都不费力",
        "type": 2,
        "answers": [
            {
                "content": "绝大部分时间有",
                "score": 1,
                "id": 1
            },
            {
                "content": "有时有",
                "score": 3,
                "id": 2
            },
            {
                "content": "很少有",
                "score": 4,
                "id": 3
            },
            {
                "content": "大部分时间有",
                "score": 2,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 14,
        "Title": "我的手指和脚趾感到麻木和刺激",
        "type": 2,
        "answers": [
            {
                "content": "绝大部分时间有",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时有",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少有",
                "score": 1,
                "id": 3
            },
            {
                "content": "大部分时间有",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 15,
        "Title": "我因胃痛和消化不良而苦恼",
        "type": 2,
        "answers": [
            {
                "content": "绝大部分时间有",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时有",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少有",
                "score": 1,
                "id": 3
            },
            {
                "content": "大部分时间有",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 16,
        "Title": "我必须频繁排尿",
        "type": 2,
        "answers": [
            {
                "content": "绝大部分时间有",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时有",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少有",
                "score": 1,
                "id": 3
            },
            {
                "content": "大部分时间有",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 17,
        "Title": "我的手总是温暖而干燥",
        "type": 2,
        "answers": [
            {
                "content": "绝大部分时间有",
                "score": 1,
                "id": 1
            },
            {
                "content": "有时有",
                "score": 3,
                "id": 2
            },
            {
                "content": "很少有",
                "score": 4,
                "id": 3
            },
            {
                "content": "大部分时间有",
                "score": 2,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 18,
        "Title": "我觉得脸发烧发红",
        "type": 2,
        "answers": [
            {
                "content": "绝大部分时间有",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时有",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少有",
                "score": 1,
                "id": 3
            },
            {
                "content": "大部分时间有",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 19,
        "Title": "我容易入睡，晚上休息很好",
        "type": 2,
        "answers": [
            {
                "content": "绝大部分时间有",
                "score": 1,
                "id": 1
            },
            {
                "content": "有时有",
                "score": 3,
                "id": 2
            },
            {
                "content": "很少有",
                "score": 4,
                "id": 3
            },
            {
                "content": "大部分时间有",
                "score": 2,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 20,
        "Title": "我做恶梦",
        "type": 2,
        "answers": [
            {
                "content": "绝大部分时间有",
                "score": 4,
                "id": 1
            },
            {
                "content": "有时有",
                "score": 2,
                "id": 2
            },
            {
                "content": "很少有",
                "score": 1,
                "id": 3
            },
            {
                "content": "大部分时间有",
                "score": 3,
                "id": 4
            }
        ]
    }
]`,
		},
		{
			Name:     "马斯洛安全感测试",
			OpenKfId: "Tk3nQYFmSdx8PDFuflP2ITuCrIYwz5U9",
			Info: `[
    {
        "QuestionNumber": 1,
        "Title": "通常，我更愿与人呆在一起，而不是个人独处",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 2,
        "Title": "在社交方面我感到轻松",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 3,
        "Title": "我缺乏自信",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 4,
        "Title": "感到自己已经得到了足够的赞扬",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 5,
        "Title": "我经常感到对世事的不满",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 6,
        "Title": "我感到人们像尊重他人一样地尊重我",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 7,
        "Title": "一次窘迫的经历会使我在很长时间内感到不安和焦虑",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 8,
        "Title": "我对自己感到满意",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 9,
        "Title": "一般说来，我不是一个自私的人",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 10,
        "Title": "我倾向于通过逃避来避免一些不愉快的事情",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 11,
        "Title": "当我与别人在一起时，我也常常会有一种孤独的感觉",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 12,
        "Title": "我感到生活对我来说是公平的",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 13,
        "Title": "当朋友批评我时，我是可以接受的",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 14,
        "Title": "我很容易气馁",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 15,
        "Title": "我通常对绝大多数人都是友好的",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 16,
        "Title": "我经常感到活着没有意思",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 17,
        "Title": "一般说来，我是一个乐观主义者",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 18,
        "Title": "我认为我是一个相当敏感的人",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 19,
        "Title": "一般说来，我是一个快活的人",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 20,
        "Title": "通常，我对自己抱有信心",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 21,
        "Title": "我常常自己感到不自然",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 0,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 22,
        "Title": "我对自己不是很满意",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 23,
        "Title": "我经常情绪低落",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 24,
        "Title": "在我与每个人第一次见面时，我常常感到对方可能不会喜欢我",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 25,
        "Title": "我对自己有足够的信心",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 26,
        "Title": "通常，我认为大多数人都是可以信任的",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 27,
        "Title": "我认为，在这个世界上我是一个有用的人",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 28,
        "Title": "一般说来，我与他人相处很融洽",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 29,
        "Title": "我经常为自己的未来发愁",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 30,
        "Title": "我感到自己是坚强有力的",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 31,
        "Title": "我很健谈",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 32,
        "Title": "我有一种自己是别人的负担的感觉",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 33,
        "Title": "我在表达自己感情方面存在困难",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 34,
        "Title": "我时常为他人的幸运而感到欣喜",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 35,
        "Title": "我经常感到似乎遗忘了什么事情",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 36,
        "Title": "我是一个比较多疑的人",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 37,
        "Title": "一般说来，我认为世界是一个适于生存的好地方",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 38,
        "Title": "我很容易不安",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 39,
        "Title": "我经常反省自己",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 0,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 40,
        "Title": "我是在按照自己的意愿生活，而不是按照其他什么人的意愿在生活",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 41,
        "Title": "当事情没办好时，我为自己感到悲哀和伤心",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 0,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 42,
        "Title": "我感到自己在工作和职业上是一个成功者",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 43,
        "Title": "我通常愿意让别人了解我究竟是怎样一个人",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 44,
        "Title": "我感到自己没有能很好地适应生活",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 45,
        "Title": "我经常抱着“车到山前必有路”的信念而坚持将事情做下去",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 46,
        "Title": "我感到生活是一个沉重的负担",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 47,
        "Title": "我被自卑所困扰",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 48,
        "Title": "一般说来，我感到还好",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 49,
        "Title": "我与异性相处得很好",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 50,
        "Title": "在街上，我曾因感到人们在看我而烦恼",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 51,
        "Title": "我很容易受伤害",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 0,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 52,
        "Title": "在这个世界上，我感到温暖",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 53,
        "Title": "我为自己的智力而忧虑",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 54,
        "Title": "通常，我使别人感到轻松",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 55,
        "Title": "对于未来，我隐隐有一种恐惧感",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 56,
        "Title": "我的行为很自然",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 57,
        "Title": "一般说来，我是幸运的",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 58,
        "Title": "我有一个幸福的童年",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 59,
        "Title": "我有许多真正的朋友",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 60,
        "Title": "在多数时间中我都感到不安",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 61,
        "Title": "我不喜欢竞争",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 62,
        "Title": "我的家庭环境很幸福",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 63,
        "Title": "我时常担心会遇到飞来的横祸",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 64,
        "Title": "在与人相处时，我常常会感到很烦燥",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 65,
        "Title": "一般说来，我很容易满足",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 66,
        "Title": "我的情绪时常会一下子从非常高兴变得非常悲哀",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 67,
        "Title": "一般说来，我受到人们的尊重和尊敬",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 68,
        "Title": "我可以很好地与别人配合工作",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 69,
        "Title": "我感到自己不能控制自己的情感",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 70,
        "Title": "我有时感到人们在嘲笑我",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 71,
        "Title": "一般说来，我是一个比较陌生的人",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 72,
        "Title": "总的说来，我感到世界对我是公正的",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 73,
        "Title": "我曾经因怀疑一些事情并非真实而苦恼",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 74,
        "Title": "我经常受到羞辱",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 75,
        "Title": "我经常感到自己被人们视为异乎寻常",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            },
            {
                "content": "不清楚",
                "score": 1,
                "id": 3
            }
        ]
    }
]`,
		},
		{
			Name:     "Beck抑郁自测",
			OpenKfId: "TJfQvCKQXnplYXPrPrqzTS3O5lbYi8Cd",
			Info: `[
    {
        "QuestionNumber": 1,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "我不感到忧愁",
                "score": 0,
                "id": 1
            },
            {
                "content": "我感到忧愁",
                "score": 1,
                "id": 2
            },
            {
                "content": "我整天都感到忧愁，且不能改变这种情绪",
                "score": 2,
                "id": 3
            },
            {
                "content": "我非常忧伤或不愉快，以致我不能忍受",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 2,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "对于将来我不感到悲观",
                "score": 0,
                "id": 1
            },
            {
                "content": "我对将来感到悲观",
                "score": 1,
                "id": 2
            },
            {
                "content": "我感到没有什么可指望的",
                "score": 2,
                "id": 3
            },
            {
                "content": "我感到将来无望，事事都不能变好",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 3,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "我不象一个失败者",
                "score": 0,
                "id": 1
            },
            {
                "content": "我觉得我比一般人失败的次数多些",
                "score": 1,
                "id": 2
            },
            {
                "content": "当我回首过去我看到的是许多失败",
                "score": 2,
                "id": 3
            },
            {
                "content": "我感到我是一个彻底失败了的人",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 4,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "我对事物象往常一样满意",
                "score": 0,
                "id": 1
            },
            {
                "content": "我对事物不象往常一样满意",
                "score": 1,
                "id": 2
            },
            {
                "content": "我不再对任何事物感到真正的满意",
                "score": 2,
                "id": 3
            },
            {
                "content": "我对每件事都不满意或讨厌",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 5,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "我没有特别感到内疚",
                "score": 0,
                "id": 1
            },
            {
                "content": "在相当一部分时间内我感到内疚",
                "score": 1,
                "id": 2
            },
            {
                "content": "在部分时间里我感到内疚",
                "score": 2,
                "id": 3
            },
            {
                "content": "我时刻感到内疚",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 6,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "我没有感到正在受惩罚",
                "score": 0,
                "id": 1
            },
            {
                "content": "我感到我可能受惩罚",
                "score": 1,
                "id": 2
            },
            {
                "content": "我预感会受惩罚",
                "score": 2,
                "id": 3
            },
            {
                "content": "我感到我正在受惩罚",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 7,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "我感到我并不使人失望",
                "score": 0,
                "id": 1
            },
            {
                "content": "我对自己失望",
                "score": 1,
                "id": 2
            },
            {
                "content": "我讨厌自己",
                "score": 2,
                "id": 3
            },
            {
                "content": "我痛恨自己",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 8,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "我感觉我并不比别人差",
                "score": 0,
                "id": 1
            },
            {
                "content": "我对自己的缺点和错误常自我反省",
                "score": 1,
                "id": 2
            },
            {
                "content": "我经常责备自己的过失",
                "score": 2,
                "id": 3
            },
            {
                "content": "每次发生糟糕的事我都责备自己",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 9,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "我没有任何自杀的想法",
                "score": 0,
                "id": 1
            },
            {
                "content": "我有自杀的的念头但不会真去自杀",
                "score": 1,
                "id": 2
            },
            {
                "content": "我很想自杀",
                "score": 2,
                "id": 3
            },
            {
                "content": "如果我有机会我就会自杀",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 10,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "我并不比以往爱哭",
                "score": 0,
                "id": 1
            },
            {
                "content": "我现在比以前爱哭",
                "score": 1,
                "id": 2
            },
            {
                "content": "现在我经常哭",
                "score": 2,
                "id": 3
            },
            {
                "content": "我以往能哭，但现在即使我想哭也哭不出来",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 11,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "我并不比以往容易激惹",
                "score": 0,
                "id": 1
            },
            {
                "content": "我现在经常容易发火",
                "score": 1,
                "id": 2
            },
            {
                "content": "以往能激惹我的那些事情现在则完全不能激惹我了",
                "score": 2,
                "id": 3
            },
            {
                "content": "我比以往容易激惹或容易生气",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 12,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "我对他人的兴趣没有减少",
                "score": 0,
                "id": 1
            },
            {
                "content": "我对他人的兴趣比以往减少了",
                "score": 1,
                "id": 2
            },
            {
                "content": "我对他人丧失了大部分兴趣",
                "score": 2,
                "id": 3
            },
            {
                "content": "我对他人现在毫无兴趣",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 13,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "我与以往一样能作决定",
                "score": 0,
                "id": 1
            },
            {
                "content": "我现在作决定没有以前果断",
                "score": 1,
                "id": 2
            },
            {
                "content": "我现在作决定比以前困难得多",
                "score": 2,
                "id": 3
            },
            {
                "content": "我现在完全不能作决定",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 14,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "我觉得自己看上去和以前差不多",
                "score": 0,
                "id": 1
            },
            {
                "content": "我担心我看上去老了或没有以前好看了",
                "score": 1,
                "id": 2
            },
            {
                "content": "我觉得我的外貌变得不好看了，而且是永久性的改变",
                "score": 2,
                "id": 3
            },
            {
                "content": "我认为我看上去很丑了",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 15,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "我能象以往一样工作",
                "score": 0,
                "id": 1
            },
            {
                "content": "我要经一番特别努力才能开始做事",
                "score": 1,
                "id": 2
            },
            {
                "content": "我做任何事都必须作很大的努力，强迫自己去做",
                "score": 2,
                "id": 3
            },
            {
                "content": "我完全不能工作",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 16,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "我睡眠象以往一样好",
                "score": 0,
                "id": 1
            },
            {
                "content": "我睡眠没有以往那样好",
                "score": 1,
                "id": 2
            },
            {
                "content": "我比往常早醒1~2小时，再入睡有困难",
                "score": 2,
                "id": 3
            },
            {
                "content": "我比往常早醒几个小时",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 17,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "我现在并不比以往感到容易疲劳",
                "score": 0,
                "id": 1
            },
            {
                "content": "我现在比以往容易疲劳",
                "score": 1,
                "id": 2
            },
            {
                "content": "我做任何事都容易疲劳",
                "score": 2,
                "id": 3
            },
            {
                "content": "我太疲劳了以致我不能做任何事情",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 18,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "我的食欲与以前一样好",
                "score": 0,
                "id": 1
            },
            {
                "content": "我现在食欲没有往常那样好",
                "score": 1,
                "id": 2
            },
            {
                "content": "我的食欲现在差多了",
                "score": 2,
                "id": 3
            },
            {
                "content": "我完全没有食欲了",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 19,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "我最近没有明显的体重减轻",
                "score": 0,
                "id": 1
            },
            {
                "content": "我体重下降超过5斤",
                "score": 1,
                "id": 2
            },
            {
                "content": "我体重下降超过10斤",
                "score": 2,
                "id": 3
            },
            {
                "content": "我体重下降超过15斤，我在控制饮食来减轻体重",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 20,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "与以往比我并不过分担心身体健康",
                "score": 0,
                "id": 1
            },
            {
                "content": "我担心我身体的毛病如疼痛、反胃及便秘",
                "score": 1,
                "id": 2
            },
            {
                "content": "我很着急身体的毛病而妨碍我思考其他问题",
                "score": 2,
                "id": 3
            },
            {
                "content": "我非常着急身体疾病，以致不能思考任何其它事情",
                "score": 3,
                "id": 4
            }
        ]
    },
    {
        "QuestionNumber": 21,
        "Title": "最近一周，最适合您的情况是",
        "type": 2,
        "answers": [
            {
                "content": "我的性欲最近没有什么变化",
                "score": 0,
                "id": 1
            },
            {
                "content": "我的性欲比以往差些",
                "score": 1,
                "id": 2
            },
            {
                "content": "现在我的性欲比以往减退了许多",
                "score": 2,
                "id": 3
            },
            {
                "content": "我完全丧失了性欲",
                "score": 3,
                "id": 4
            }
        ]
    }
]`,
		},
		{
			Name:     "自动思维测试",
			OpenKfId: "YhNDr78RFtvlMSZ0InBWIWjSfVfL9KTt",
			Info: `[
    {
        "QuestionNumber": 1,
        "Title": "我觉得活在世上困难重重",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 2,
        "Title": "我不好",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 3,
        "Title": "为什么我总不能成功",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 4,
        "Title": "没有人理解我",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 5,
        "Title": "我让人失望",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 6,
        "Title": "我觉得过不下去了",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 7,
        "Title": "真希望我能好一点",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 8,
        "Title": "我很虚弱",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 9,
        "Title": "我的生活不按我的愿望发展",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 10,
        "Title": "我对自己很不满意",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 11,
        "Title": "我觉得一切都不好了",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 12,
        "Title": "我无法坚持下去",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 13,
        "Title": "我无法重新开始",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 14,
        "Title": "我究竟犯了什么毛病",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 15,
        "Title": "真希望我是在另外一个地方",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 16,
        "Title": "我无法同时对付这些事情",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 17,
        "Title": "我恨我自己",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 18,
        "Title": "我毫无价值",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 19,
        "Title": "真希望我一下子就消失了",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 20,
        "Title": "我这是怎么了",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 21,
        "Title": "我是个失败者",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 22,
        "Title": "我的生活一团糟",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 23,
        "Title": "我一事无成",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 24,
        "Title": "我不可能干好",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 25,
        "Title": "我觉得孤立无援",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 26,
        "Title": "有些东西必须改变",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 27,
        "Title": "我肯定有问题",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 28,
        "Title": "我的将来毫无希望",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 29,
        "Title": "这根本毫无价值",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 30,
        "Title": "我干什么事都有头无尾",
        "type": 2,
        "answers": [
            {
                "content": "无",
                "score": 1,
                "id": 1
            },
            {
                "content": "偶尔出现",
                "score": 2,
                "id": 2
            },
            {
                "content": "有时出现",
                "score": 3,
                "id": 3
            },
            {
                "content": "经常出现",
                "score": 4,
                "id": 4
            },
            {
                "content": "持续存在",
                "score": 5,
                "id": 5
            }
        ]
    }
]`,
		},
		{
			Name:     "自信心测验",
			OpenKfId: "J4TRSQ4j5pjO0FBW7BJ93x0nhKBzMqg5",
			Info: `[
    {
        "QuestionNumber": 1,
        "Title": "一旦你下了决心，即使 没有人赞同，你仍会坚持做到底吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 2,
        "Title": "参加晚宴时，即使很想上洗手间，你也会忍着直到宴会结束吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 3,
        "Title": "如果想买性感内衣，你会尽量邮购，而不亲自到店里去吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 4,
        "Title": "你认为自己是个较完美的人吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 5,
        "Title": "如果店员的服务态度不好，你会告诉他们的经理吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 6,
        "Title": "你不常欣赏自己的照片吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 7,
        "Title": "别人批评你，你会觉得难过吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 8,
        "Title": "你很少对人说出你真正的意见吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 9,
        "Title": "对别人的赞美，你持怀疑的态度吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 10,
        "Title": "你总是觉得自己比别人差吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 11,
        "Title": "你对自己的外表满意吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 12,
        "Title": "你认为自己的能力比别人差吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 13,
        "Title": "在聚会上，只有你一个人穿得不正式，你会感到不自在吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 14,
        "Title": "你是个受欢迎的人吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 15,
        "Title": "你认为自己很有魅力吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 16,
        "Title": "你有幽默感吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 17,
        "Title": "目前的工作是你的专长吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 18,
        "Title": "你懂得搭配衣服吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 19,
        "Title": "危急时，你很冷静吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 20,
        "Title": "你与别人合作无间吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 21,
        "Title": "你认为自己只是个寻常人吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 22,
        "Title": "你经常希望自己长得像某某人吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 23,
        "Title": "你经常羡慕别人的成就吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 24,
        "Title": "你为了不使他人难过，而放弃自己喜欢做的事吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 25,
        "Title": "你会为了讨好别人而打扮吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 26,
        "Title": "你勉强自己做许多不愿意做的事吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 27,
        "Title": "你任由他人来支配你的生活吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 28,
        "Title": "你认为你的优点比缺点多吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 29,
        "Title": "你经常跟人说抱歉吗即使在不是你错的情况下",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 30,
        "Title": "如果在非故意的情况下伤了别人的心，你会难过吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 31,
        "Title": "你希望自己具备更多的才能和天赋吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 32,
        "Title": "你经常听取别人的意见吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 33,
        "Title": "在聚会上，你经常等别人先跟你打招呼吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 34,
        "Title": "你每天照镜子超过三次吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 35,
        "Title": "你的个性很强吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 36,
        "Title": "你是个优秀的领导者吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 37,
        "Title": "你的记性很好吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 38,
        "Title": "你对异性有吸引力吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 39,
        "Title": "你懂得理财吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 1,
                "id": 1
            },
            {
                "content": "否",
                "score": 0,
                "id": 2
            }
        ]
    },
    {
        "QuestionNumber": 40,
        "Title": "买衣服前 , 你通常先听取别人的意见吗",
        "type": 1,
        "answers": [
            {
                "content": "是",
                "score": 0,
                "id": 1
            },
            {
                "content": "否",
                "score": 1,
                "id": 2
            }
        ]
    }
]`,
		},
		{
			Name:     "社会适应能力测试",
			OpenKfId: "hlvSHiCc9tavzp50cbO8wfTkZohkVWCg",
			Info: `[
    {
        "QuestionNumber": 1,
        "Title": "对自己的某次失败，我绝对不会提及，就怕被别人抓住自己的弱点",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": -2,
                "id": 1
            },
            {
                "content": "无法肯定",
                "score": 0,
                "id": 2
            },
            {
                "content": "不是",
                "score": 2,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 2,
        "Title": "每到一个新的地方，我很容易同别人接近",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 2,
                "id": 3
            },
            {
                "content": "无法肯定",
                "score": 0,
                "id": 2
            },
            {
                "content": "不是",
                "score": -2,
                "id": 1
            }
        ]
    },
    {
        "QuestionNumber": 3,
        "Title": "在陌生人面前，我常无话可说，以至感到尴尬",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": -2,
                "id": 1
            },
            {
                "content": "无法肯定",
                "score": 0,
                "id": 2
            },
            {
                "content": "不是",
                "score": 2,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 4,
        "Title": "当我选衣服时，我会想跟随潮流，希望适合自己",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 2,
                "id": 3
            },
            {
                "content": "无法肯定",
                "score": 0,
                "id": 2
            },
            {
                "content": "不是",
                "score": -2,
                "id": 1
            }
        ]
    },
    {
        "QuestionNumber": 5,
        "Title": "每到一个新地方，我第一天总是睡不好，就是在家里，只要换一张床，有时也会失眠",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": -2,
                "id": 1
            },
            {
                "content": "无法肯定",
                "score": 0,
                "id": 2
            },
            {
                "content": "不是",
                "score": 2,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 6,
        "Title": "我积极参加社会实践，以积累工作经验",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 2,
                "id": 3
            },
            {
                "content": "无法肯定",
                "score": 0,
                "id": 2
            },
            {
                "content": "不是",
                "score": -2,
                "id": 1
            }
        ]
    },
    {
        "QuestionNumber": 7,
        "Title": "越是人多的地方，我越感到紧张.",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": -2,
                "id": 1
            },
            {
                "content": "无法肯定",
                "score": 0,
                "id": 2
            },
            {
                "content": "不是",
                "score": 2,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 8,
        "Title": "在正式比赛或考试时，我的成绩多半不会比平时练习差",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 2,
                "id": 3
            },
            {
                "content": "无法肯定",
                "score": 0,
                "id": 2
            },
            {
                "content": "不是",
                "score": -2,
                "id": 1
            }
        ]
    },
    {
        "QuestionNumber": 9,
        "Title": "必须在大庭广众面前说话，我会因怯场，变得不知所措，说话语无伦次",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": -2,
                "id": 1
            },
            {
                "content": "无法肯定",
                "score": 0,
                "id": 2
            },
            {
                "content": "不是",
                "score": 2,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 10,
        "Title": "当我骑自行车到一个比较远的地方去参加社交活动，中途找不到路标时，我会耐心等待过路车或等人走过时问个清楚",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 2,
                "id": 3
            },
            {
                "content": "无法肯定",
                "score": 0,
                "id": 2
            },
            {
                "content": "不是",
                "score": -2,
                "id": 1
            }
        ]
    },
    {
        "QuestionNumber": 11,
        "Title": "我从来不愿为了在别人面前留下好印象而特地去做，就算是有机会也不做",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": -2,
                "id": 1
            },
            {
                "content": "无法肯定",
                "score": 0,
                "id": 2
            },
            {
                "content": "不是",
                "score": 2,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 12,
        "Title": "和同学、家人相处，我很少固执己见，乐于采纳别人的看法",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 2,
                "id": 3
            },
            {
                "content": "无法肯定",
                "score": 0,
                "id": 2
            },
            {
                "content": "不是",
                "score": -2,
                "id": 1
            }
        ]
    },
    {
        "QuestionNumber": 13,
        "Title": "同别人争论时，我常常感到语塞，事后才想起该怎样反驳对方，可惜已经太迟了",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": -2,
                "id": 1
            },
            {
                "content": "无法肯定",
                "score": 0,
                "id": 2
            },
            {
                "content": "不是",
                "score": 2,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 14,
        "Title": "我对生活条件要求不高，即使生活条件很艰苦，我也能过得很愉快",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 2,
                "id": 3
            },
            {
                "content": "无法肯定",
                "score": 0,
                "id": 2
            },
            {
                "content": "不是",
                "score": -2,
                "id": 1
            }
        ]
    },
    {
        "QuestionNumber": 15,
        "Title": "在受到别人批评时，我会想找到机会去批评他(她)",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": -2,
                "id": 1
            },
            {
                "content": "无法肯定",
                "score": 0,
                "id": 2
            },
            {
                "content": "不是",
                "score": 2,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 16,
        "Title": "在决定胜负成败的关键时刻，我虽然很紧张，但总能很快地使自己镇定下来",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 2,
                "id": 3
            },
            {
                "content": "无法肯定",
                "score": 0,
                "id": 2
            },
            {
                "content": "不是",
                "score": -2,
                "id": 1
            }
        ]
    },
    {
        "QuestionNumber": 17,
        "Title": "我不喜欢的东西，不管怎么学也学不会",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": -2,
                "id": 1
            },
            {
                "content": "无法肯定",
                "score": 0,
                "id": 2
            },
            {
                "content": "不是",
                "score": 2,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 18,
        "Title": "在嘈杂混乱的环境里，我仍然能集中精力学习，并且效率较高",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 2,
                "id": 3
            },
            {
                "content": "无法肯定",
                "score": 0,
                "id": 2
            },
            {
                "content": "不是",
                "score": -2,
                "id": 1
            }
        ]
    },
    {
        "QuestionNumber": 19,
        "Title": "我不喜欢陌生人来家里做客，每逢这种情况，我就有意回避",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": -2,
                "id": 1
            },
            {
                "content": "无法肯定",
                "score": 0,
                "id": 2
            },
            {
                "content": "不是",
                "score": 2,
                "id": 3
            }
        ]
    },
    {
        "QuestionNumber": 20,
        "Title": "我很喜欢参加社交活动，我感到这是交朋友的好机会",
        "type": 2,
        "answers": [
            {
                "content": "是",
                "score": 2,
                "id": 3
            },
            {
                "content": "无法肯定",
                "score": 0,
                "id": 2
            },
            {
                "content": "不是",
                "score": -2,
                "id": 1
            }
        ]
    }
]`,
		},
		{
			Name:     "自评抑郁量表",
			OpenKfId: "cOmxua5h8DRdIFRiqYqjiTcNwsrQBsIp",
			Info: `[
    {
        "QuestionNumber": 1,
        "Title": "我觉得闷闷不乐，情绪低沉",
        "type": 2,
        "answers": [
            {
                "content": "从无或偶尔",
                "id": 1,
                "score": 1
            },
            {
                "content": "有时",
                "id": 2,
                "score": 2
            },
            {
                "content": "经常",
                "id": 3,
                "score": 3
            },
            {
                "content": "总是如此",
                "id": 4,
                "score": 4
            }
        ]
    },
    {
        "QuestionNumber": 2,
        "Title": "我觉得一天之中早晨最好",
        "type": 2,
        "answers": [
            {
                "content": "从无或偶尔",
                "id": 1,
                "score": 4
            },
            {
                "content": "有时",
                "id": 2,
                "score": 3
            },
            {
                "content": "经常",
                "id": 3,
                "score": 2
            },
            {
                "content": "总是如此",
                "id": 4,
                "score": 1
            }
        ]
    },
    {
        "QuestionNumber": 3,
        "Title": "我一阵阵哭出来或觉得想哭",
        "type": 2,
        "answers": [
            {
                "content": "从无或偶尔",
                "id": 1,
                "score": 1
            },
            {
                "content": "有时",
                "id": 2,
                "score": 2
            },
            {
                "content": "经常",
                "id": 3,
                "score": 3
            },
            {
                "content": "总是如此",
                "id": 4,
                "score": 4
            }
        ]
    },
    {
        "QuestionNumber": 4,
        "Title": "我晚上睡眠不好",
        "type": 2,
        "answers": [
            {
                "content": "从无或偶尔",
                "id": 1,
                "score": 1
            },
            {
                "content": "有时",
                "id": 2,
                "score": 2
            },
            {
                "content": "经常",
                "id": 3,
                "score": 3
            },
            {
                "content": "总是如此",
                "id": 4,
                "score": 4
            }
        ]
    },
    {
        "QuestionNumber": 5,
        "Title": "我吃得跟平常一样多",
        "type": 2,
        "answers": [
            {
                "content": "从无或偶尔",
                "id": 1,
                "score": 4
            },
            {
                "content": "有时",
                "id": 2,
                "score": 3
            },
            {
                "content": "经常",
                "id": 3,
                "score": 2
            },
            {
                "content": "总是如此",
                "id": 4,
                "score": 1
            }
        ]
    },
    {
        "QuestionNumber": 6,
        "Title": "我与异性密切接触时和以往一样感到愉快",
        "type": 2,
        "answers": [
            {
                "content": "从无或偶尔",
                "id": 1,
                "score": 4
            },
            {
                "content": "有时",
                "id": 2,
                "score": 3
            },
            {
                "content": "经常",
                "id": 3,
                "score": 2
            },
            {
                "content": "总是如此",
                "id": 4,
                "score": 1
            }
        ]
    },
    {
        "QuestionNumber": 7,
        "Title": "我发觉我的体重在下降",
        "type": 2,
        "answers": [
            {
                "content": "从无或偶尔",
                "id": 1,
                "score": 1
            },
            {
                "content": "有时",
                "id": 2,
                "score": 2
            },
            {
                "content": "经常",
                "id": 3,
                "score": 3
            },
            {
                "content": "总是如此",
                "id": 4,
                "score": 4
            }
        ]
    },
    {
        "QuestionNumber": 8,
        "Title": "我有便秘的苦恼",
        "type": 2,
        "answers": [
            {
                "content": "从无或偶尔",
                "id": 1,
                "score": 1
            },
            {
                "content": "有时",
                "id": 2,
                "score": 2
            },
            {
                "content": "经常",
                "id": 3,
                "score": 3
            },
            {
                "content": "总是如此",
                "id": 4,
                "score": 4
            }
        ]
    },
    {
        "QuestionNumber": 9,
        "Title": "我心跳比平常快",
        "type": 2,
        "answers": [
            {
                "content": "从无或偶尔",
                "id": 1,
                "score": 1
            },
            {
                "content": "有时",
                "id": 2,
                "score": 2
            },
            {
                "content": "经常",
                "id": 3,
                "score": 3
            },
            {
                "content": "总是如此",
                "id": 4,
                "score": 4
            }
        ]
    },
    {
        "QuestionNumber": 10,
        "Title": "我无缘无故地感到疲乏",
        "type": 2,
        "answers": [
            {
                "content": "从无或偶尔",
                "id": 1,
                "score": 1
            },
            {
                "content": "有时",
                "id": 2,
                "score": 2
            },
            {
                "content": "经常",
                "id": 3,
                "score": 3
            },
            {
                "content": "总是如此",
                "id": 4,
                "score": 4
            }
        ]
    },
    {
        "QuestionNumber": 11,
        "Title": "我的头脑跟平常一样清楚",
        "type": 2,
        "answers": [
            {
                "content": "从无或偶尔",
                "id": 1,
                "score": 4
            },
            {
                "content": "有时",
                "id": 2,
                "score": 3
            },
            {
                "content": "经常",
                "id": 3,
                "score": 2
            },
            {
                "content": "总是如此",
                "id": 4,
                "score": 1
            }
        ]
    },
    {
        "QuestionNumber": 12,
        "Title": "我觉得经常做的事情并没有困难",
        "type": 2,
        "answers": [
            {
                "content": "从无或偶尔",
                "id": 1,
                "score": 4
            },
            {
                "content": "有时",
                "id": 2,
                "score": 3
            },
            {
                "content": "经常",
                "id": 3,
                "score": 2
            },
            {
                "content": "总是如此",
                "id": 4,
                "score": 1
            }
        ]
    },
    {
        "QuestionNumber": 13,
        "Title": "我觉得不安而平静不下来",
        "type": 2,
        "answers": [
            {
                "content": "从无或偶尔",
                "id": 1,
                "score": 1
            },
            {
                "content": "有时",
                "id": 2,
                "score": 2
            },
            {
                "content": "经常",
                "id": 3,
                "score": 3
            },
            {
                "content": "总是如此",
                "id": 4,
                "score": 4
            }
        ]
    },
    {
        "QuestionNumber": 14,
        "Title": "我对将来抱有希望",
        "type": 2,
        "answers": [
            {
                "content": "从无或偶尔",
                "id": 1,
                "score": 4
            },
            {
                "content": "有时",
                "id": 2,
                "score": 3
            },
            {
                "content": "经常",
                "id": 3,
                "score": 2
            },
            {
                "content": "总是如此",
                "id": 4,
                "score": 1
            }
        ]
    },
    {
        "QuestionNumber": 15,
        "Title": "我比平常容易生气激动",
        "type": 2,
        "answers": [
            {
                "content": "从无或偶尔",
                "id": 1,
                "score": 4
            },
            {
                "content": "有时",
                "id": 2,
                "score": 3
            },
            {
                "content": "经常",
                "id": 3,
                "score": 2
            },
            {
                "content": "总是如此",
                "id": 4,
                "score": 1
            }
        ]
    },
    {
        "QuestionNumber": 16,
        "Title": "我觉得作出决定是容易的",
        "type": 2,
        "answers": [
            {
                "content": "从无或偶尔",
                "id": 1,
                "score": 4
            },
            {
                "content": "有时",
                "id": 2,
                "score": 3
            },
            {
                "content": "经常",
                "id": 3,
                "score": 2
            },
            {
                "content": "总是如此",
                "id": 4,
                "score": 1
            }
        ]
    },
    {
        "QuestionNumber": 17,
        "Title": "我觉得自己是个有用的人，有人需要我",
        "type": 2,
        "answers": [
            {
                "content": "从无或偶尔",
                "id": 1,
                "score": 4
            },
            {
                "content": "有时",
                "id": 2,
                "score": 3
            },
            {
                "content": "经常",
                "id": 3,
                "score": 2
            },
            {
                "content": "总是如此",
                "id": 4,
                "score": 1
            }
        ]
    },
    {
        "QuestionNumber": 18,
        "Title": "我的生活过得很有意思",
        "type": 2,
        "answers": [
            {
                "content": "从无或偶尔",
                "id": 1,
                "score": 4
            },
            {
                "content": "有时",
                "id": 2,
                "score": 3
            },
            {
                "content": "经常",
                "id": 3,
                "score": 2
            },
            {
                "content": "总是如此",
                "id": 4,
                "score": 1
            }
        ]
    },
    {
        "QuestionNumber": 19,
        "Title": "我认为如果我死了，别人会生活得好些",
        "type": 2,
        "answers": [
            {
                "content": "从无或偶尔",
                "id": 1,
                "score": 1
            },
            {
                "content": "有时",
                "id": 2,
                "score": 2
            },
            {
                "content": "经常",
                "id": 3,
                "score": 3
            },
            {
                "content": "总是如此",
                "id": 4,
                "score": 4
            }
        ]
    },
    {
        "QuestionNumber": 20,
        "Title": "平常感兴趣的事我仍然照样感兴趣",
        "type": 2,
        "answers": [
            {
                "content": "从无或偶尔",
                "id": 1,
                "score": 4
            },
            {
                "content": "有时",
                "id": 2,
                "score": 3
            },
            {
                "content": "经常",
                "id": 3,
                "score": 2
            },
            {
                "content": "总是如此",
                "id": 4,
                "score": 1
            }
        ]
    }
]`,
		},
		{
			Name:     "述情障碍测试",
			OpenKfId: "ERc6gf0tDSTs6yDOBy1Tpx1pJWqMRo8D",
			Info: `[
    {
        "QuestionNumber": 1,
        "Title": "当我哭泣时，我知道是什么原因",
        "type": 2,
        "answers": [
            {
                "content": "完全同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全不同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 2,
        "Title": "空想纯粹是浪费时间",
        "type": 2,
        "answers": [
            {
                "content": "完全不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 3,
        "Title": "我希望自己不那么害羞",
        "type": 2,
        "answers": [
            {
                "content": "完全不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 4,
        "Title": "我常搞不清自己是什么样的感受",
        "type": 2,
        "answers": [
            {
                "content": "完全不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 5,
        "Title": "我常幻想着将来",
        "type": 2,
        "answers": [
            {
                "content": "完全同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全不同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 6,
        "Title": "我似乎交朋友和别人一样容易",
        "type": 2,
        "answers": [
            {
                "content": "完全不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 7,
        "Title": "知道问题的答案比知道其原因更重要",
        "type": 2,
        "answers": [
            {
                "content": "完全不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 8,
        "Title": "我难以用恰当的词描述自己的情感",
        "type": 2,
        "answers": [
            {
                "content": "完全不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 9,
        "Title": "我喜欢别人知道我对事物的态度",
        "type": 2,
        "answers": [
            {
                "content": "完全同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全不同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 10,
        "Title": "有些身体感觉连医生也不理解",
        "type": 2,
        "answers": [
            {
                "content": "完全不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 11,
        "Title": "只做工作是不够的我需知道为何做和如何做好",
        "type": 2,
        "answers": [
            {
                "content": "完全同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全不同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 12,
        "Title": "我很容易地描述自己的感受",
        "type": 2,
        "answers": [
            {
                "content": "完全不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 13,
        "Title": "我更喜欢分析间题而不仅仅描述它",
        "type": 2,
        "answers": [
            {
                "content": "完全同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全不同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 14,
        "Title": "当我心烦意乱时，我不知是伤心、害怕、愤怒",
        "type": 2,
        "answers": [
            {
                "content": "完全不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 15,
        "Title": "我常好幻想",
        "type": 2,
        "answers": [
            {
                "content": "完全同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全不同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 16,
        "Title": "当我无事可做时，常好空想",
        "type": 2,
        "answers": [
            {
                "content": "完全同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全不同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 17,
        "Title": "我常为体内的感觉所困惑",
        "type": 2,
        "answers": [
            {
                "content": "完全不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 18,
        "Title": "我极少做白日梦",
        "type": 2,
        "answers": [
            {
                "content": "完全不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 19,
        "Title": "我更关心事情的发生，而不注惫为何发生",
        "type": 2,
        "answers": [
            {
                "content": "完全不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 20,
        "Title": "我有些难以识别的感受",
        "type": 2,
        "answers": [
            {
                "content": "完全不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 21,
        "Title": "情感的沟通是很重要的",
        "type": 2,
        "answers": [
            {
                "content": "完全同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全不同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 22,
        "Title": "我觉得难以描述对别人的情感",
        "type": 2,
        "answers": [
            {
                "content": "完全不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 23,
        "Title": "别人告诉我，要更多地表达自己的感受",
        "type": 2,
        "answers": [
            {
                "content": "完全不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 24,
        "Title": "一个人应寻求更深刻的理解",
        "type": 2,
        "answers": [
            {
                "content": "完全同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本不同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全不同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 25,
        "Title": "我不知道我的内心发生了什么",
        "type": 2,
        "answers": [
            {
                "content": "完全不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 26,
        "Title": "我常不知道自己为什么气愤",
        "type": 2,
        "answers": [
            {
                "content": "完全不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "基本不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "不同意也不反对",
                "score": 3,
                "id": 3
            },
            {
                "content": "基本同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "完全同意",
                "score": 5,
                "id": 5
            }
        ]
    }
]`,
		},
		{
			Name:     "考试焦虑自测",
			OpenKfId: "FKtYIPRJqw2P4Md34KJbSVexztrEEJDK",
			Info: `[
    {
        "QuestionNumber": 1,
        "Title": "当一次重大考试就要来临时，我总是在想别人比我聪明得多",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 2,
        "Title": "如果我将要做一次智能测试，在做之前我会非常焦虑",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 3,
        "Title": "如果我知道将会有一次智能测试，在此之前我感到很自信、很轻松",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 0
            },
            {
                "id": 2,
                "content": "否",
                "score": 1
            }
        ]
    },
    {
        "QuestionNumber": 4,
        "Title": "参加重大考试时，我会出很多汗",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 5,
        "Title": "考试期间，我发现自己总是在想一些和考试内容无关的事",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 6,
        "Title": "当一次突然袭击式的考试来到时，我感到很怕",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 7,
        "Title": "考试期间我经常想到会失败",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 8,
        "Title": "重大考试后我经常感到紧张，以至胃不舒服",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 9,
        "Title": "我对智能考试和期末考试之类的事总感到发怵",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 10,
        "Title": "在一次考试中取得好成绩似乎并不能增加我在第二次考试中的信心",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 11,
        "Title": "在重大考试期间我有时感到心跳很快",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 12,
        "Title": "考试结束后我总是觉得可以比实际上做得更好",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 13,
        "Title": "考试完毕后我总是感到很抑郁",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 14,
        "Title": "每次期末考试之前，我总有一种紧张不安的感觉",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 15,
        "Title": "考试时，我的情绪反应不会干扰我考试",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 0
            },
            {
                "id": 2,
                "content": "否",
                "score": 1
            }
        ]
    },
    {
        "QuestionNumber": 16,
        "Title": "考试期间我经常很紧张，以至本来知道的东西也忘了",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 17,
        "Title": "复习重要的考试对我来说似乎是一个很大的挑战",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 18,
        "Title": "对某一门考试，我越努力复习越感到困惑",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 19,
        "Title": "某门考试一结束，我试图停止有关担忧，但做不到",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 20,
        "Title": "考试期间我有时会想我是否能完成大学学业",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 21,
        "Title": "我宁愿写一篇论文，而不是参加一次考试，作为某门课程的成绩",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 22,
        "Title": "我真希望考试不要那么烦人",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 23,
        "Title": "我相信如果我单独参加考试而且没有时间限制的话，我会考得更好",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 24,
        "Title": "想着我在考试中能得多少分，影响了我的复习和考试",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 25,
        "Title": "如果考试能废除的话，我想我能学得更好",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 26,
        "Title": "我对考试抱这样的态度：虽然我现在不懂，但我并不担心",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 0
            },
            {
                "id": 2,
                "content": "否",
                "score": 1
            }
        ]
    },
    {
        "QuestionNumber": 27,
        "Title": "我真不明白为什么有些人对考试那么紧张",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 0
            },
            {
                "id": 2,
                "content": "否",
                "score": 1
            }
        ]
    },
    {
        "QuestionNumber": 28,
        "Title": "我很差劲的想法会干扰我在考试中的表现",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 29,
        "Title": "我复习期末考试并不比复习平时考试更卖力",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 0
            },
            {
                "id": 2,
                "content": "否",
                "score": 1
            }
        ]
    },
    {
        "QuestionNumber": 30,
        "Title": "尽管我对某门考试复习很好，但我仍然感到焦虑",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 31,
        "Title": "在重大考试前，我吃不香",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 32,
        "Title": "在重大考试前我发现我的手臂会颤抖",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 33,
        "Title": "在考试前我很少有“临时抱佛脚”的需要",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 0
            },
            {
                "id": 2,
                "content": "否",
                "score": 1
            }
        ]
    },
    {
        "QuestionNumber": 34,
        "Title": "校方应认识到有些学生对考试较为焦虑，而这会影响他们的考试成绩",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 35,
        "Title": "我认为考试期间似乎不应该搞得那么紧张",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 36,
        "Title": "一接触到发下的试卷，我就觉得很不自在",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    },
    {
        "QuestionNumber": 37,
        "Title": "我讨厌老师喜欢搞“突然袭击”式考试的课程",
        "type": 2,
        "answers": [
            {
                "id": 1,
                "content": "是",
                "score": 1
            },
            {
                "id": 2,
                "content": "否",
                "score": 0
            }
        ]
    }
]`,
		},
		{
			Name:     "48题",
			OpenKfId: "HHByMTLYej29rIKxm51S8vvhi63dZNXt",
			Info: `[
    {
        "QuestionNumber": 1,
        "Title": "我在大型社交活动中感到舒适",
        "type": 3,
        "score_letter": "E",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 2,
        "Title": "我更喜欢与一两个亲密朋友交往，而不是大团体",
        "type": 3,
        "score_letter": "I",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 3,
        "Title": "与他人交往能为我充电",
        "type": 3,
        "score_letter": "E",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 4,
        "Title": "长时间与他人互动会让我感到疲劳",
        "type": 3,
        "score_letter": "I",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 5,
        "Title": "我喜欢独自时间",
        "type": 3,
        "score_letter": "I",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 6,
        "Title": "我喜欢团队活动",
        "type": 3,
        "score_letter": "E",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 7,
        "Title": "在社交聚会后，我通常感到充满活力",
        "type": 3,
        "score_letter": "E",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 8,
        "Title": "我发现与陌生人交往很容易",
        "type": 3,
        "score_letter": "E",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 9,
        "Title": "在一个安静的环境中独自工作会使我感到更舒适",
        "type": 3,
        "score_letter": "I",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 10,
        "Title": "我更喜欢独自完成任务，而不是在团队中工作",
        "type": 3,
        "score_letter": "I",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 11,
        "Title": "在参加大型活动后，我需要一些独自的时间来放松",
        "type": 3,
        "score_letter": "I",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 12,
        "Title": "与人群互动可以为我带来活力和动力",
        "type": 3,
        "score_letter": "E",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 13,
        "Title": "我更依赖经验和现实",
        "type": 3,
        "score_letter": "S",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 14,
        "Title": "我经常梦想并探索可能性",
        "type": 3,
        "score_letter": "N",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 15,
        "Title": "我关注细节",
        "type": 3,
        "score_letter": "S",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 16,
        "Title": "我喜欢思考和想象未来",
        "type": 3,
        "score_letter": "N",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 17,
        "Title": "对我来说，过去的经验非常重要",
        "type": 3,
        "score_letter": "S",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 18,
        "Title": "我经常考虑新的可能性",
        "type": 3,
        "score_letter": "N",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 19,
        "Title": "当我考虑问题时，我更关注“为什么”而不是“如何”",
        "type": 3,
        "score_letter": "N",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 20,
        "Title": "我倾向于根据我目前的经验来做决策，而不是可能的未来发展",
        "type": 3,
        "score_letter": "S",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 21,
        "Title": "抽象理论和概念使我兴奋",
        "type": 3,
        "score_letter": "N",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 22,
        "Title": "我在做决策时更依赖具体事实和明确的数据",
        "type": 3,
        "score_letter": "S",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 23,
        "Title": "我常常沉浸在对未来的幻想和设想中",
        "type": 3,
        "score_letter": "N",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 24,
        "Title": "对我来说，观察和体验当前的现实是非常重要的",
        "type": 3,
        "score_letter": "S",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 25,
        "Title": "我做决策时首先考虑事实和逻辑",
        "type": 3,
        "score_letter": "T",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 26,
        "Title": "我更多地考虑人们的情感和需求",
        "type": 3,
        "score_letter": "F",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 27,
        "Title": "在争议中，我试图保持公正",
        "type": 3,
        "score_letter": "T",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 28,
        "Title": "我的决策通常基于我的价值观",
        "type": 3,
        "score_letter": "F",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 29,
        "Title": "我试图客观地看待情境",
        "type": 3,
        "score_letter": "T",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 30,
        "Title": "我深受他人情感的影响",
        "type": 3,
        "score_letter": "F",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 31,
        "Title": "当面对决策时，我通常会依赖逻辑和分析",
        "type": 3,
        "score_letter": "T",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 32,
        "Title": "人们经常说我很有同情心",
        "type": 3,
        "score_letter": "F",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 33,
        "Title": "在决策时，我更重视客观事实而不是人们的感受",
        "type": 3,
        "score_letter": "T",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 34,
        "Title": "我更喜欢与他人建立深厚的情感联系",
        "type": 3,
        "score_letter": "F",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 35,
        "Title": "在评估某个情况时，我的情感经常会介入",
        "type": 3,
        "score_letter": "F",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 36,
        "Title": "对于大多数情况，我认为逻辑胜过感觉",
        "type": 3,
        "score_letter": "T",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 37,
        "Title": "我喜欢有明确的计划和日程",
        "type": 3,
        "score_letter": "J",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 38,
        "Title": "我喜欢随意和灵活的生活",
        "type": 3,
        "score_letter": "P",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 39,
        "Title": "我总是提前准备",
        "type": 3,
        "score_letter": "J",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 40,
        "Title": "我喜欢探索新的选择，而不是立即做决定",
        "type": 3,
        "score_letter": "P",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 41,
        "Title": "我遵循我的计划",
        "type": 3,
        "score_letter": "J",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 42,
        "Title": "我通常是即兴的",
        "type": 3,
        "score_letter": "P",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 43,
        "Title": "我喜欢按计划行事，而不是随心所欲",
        "type": 3,
        "score_letter": "J",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 44,
        "Title": "我经常会延迟决策，以便收集更多的信息",
        "type": 3,
        "score_letter": "P",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 45,
        "Title": "我认为有规划的生活比随遇而安的生活更好",
        "type": 3,
        "score_letter": "J",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 46,
        "Title": "我更喜欢在一个灵活和开放的环境中工作",
        "type": 3,
        "score_letter": "P",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 47,
        "Title": "我很少更改我的计划",
        "type": 3,
        "score_letter": "J",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    },
    {
        "QuestionNumber": 48,
        "Title": "我喜欢探索各种可能性，而不是坚持一种固定的方式",
        "type": 3,
        "score_letter": "P",
        "answers": [
            {
                "content": "非常不同意",
                "score": 1,
                "id": 1
            },
            {
                "content": "一般不同意",
                "score": 2,
                "id": 2
            },
            {
                "content": "中立",
                "score": 3,
                "id": 3
            },
            {
                "content": "一般同意",
                "score": 4,
                "id": 4
            },
            {
                "content": "非常同意",
                "score": 5,
                "id": 5
            }
        ]
    }
]`,
		},
	}

	//UB5OWvU4XufQVjOy3rhj3TgAEm17ujdd,蒙哥马利抑郁评估
	//FnyZNS9uIkw81dqrBRqeiwWyQyjcr1Sc,孤独感测试
	//8HPjI81RCupXimX33iNWgMRGHKzlaaNB,焦虑自测
	//Tk3nQYFmSdx8PDFuflP2ITuCrIYwz5U9,马斯洛安全感测试
	//TJfQvCKQXnplYXPrPrqzTS3O5lbYi8Cd,Beck抑郁自测
	//l5sjTQZKBSYGkWrupnq0rYLxdRb7n2AO,焦虑自评
	//YhNDr78RFtvlMSZ0InBWIWjSfVfL9KTt,自动思维测试
	//J4TRSQ4j5pjO0FBW7BJ93x0nhKBzMqg5,自信心测验
	//hlvSHiCc9tavzp50cbO8wfTkZohkVWCg,社会适应能力测试
	//cOmxua5h8DRdIFRiqYqjiTcNwsrQBsIp,自评抑郁量表
	//ERc6gf0tDSTs6yDOBy1Tpx1pJWqMRo8D,述情障碍测试
	//FKtYIPRJqw2P4Md34KJbSVexztrEEJDK,考试焦虑自测
	//HHByMTLYej29rIKxm51S8vvhi63dZNXt,48题

	return result
}
