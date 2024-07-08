package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[SUCCESS] = "成功"
	message[ServerFail] = "调用失败请稍后重试"
	message[RequestParamError] = "参数错误"
	message[UNAUTHORIZED] = "无效的token"
	message[FORBIDDEN] = "权限不足"
	message[PasswordIncorrect] = "密码错误"
	message[RouteNotFound] = "请求资源不存在"
	message[RouteNotMatch] = "请求方式错误"
	message[DBError] = "数据库繁忙,请稍后再试"

	//common
	message[RecordNotFound] = "record not found"
	message[RepeatRequest] = "please do not repeat the request"
	message[ParamMiss] = "param miss"
	message[FileMiss] = "no such file"
	message[SystemError] = "系统错误"
	message[SystemBusyError] = "系统繁忙，请稍后再试～"
	message[RequestError] = "请求参数错误，请稍后再试～"
	message[InvalidToken] = "无效的token"
	message[ParamError] = "param error"
	message[SensitiveError] = "您的消息中含有敏感词信息，请重新输入"
	//config
	message[ConfigExist] = "当前配置在系统中已存在"
	message[ConfigOpenAiKeyEmpty] = "未配置key"
	message[ConfigEmpty] = "缺少配置"

	//login
	message[LoginMobileError] = "mobile error"
	message[LoginVerifyCodeError] = "verify code error"
	message[LoginAccountNotExist] = "account is not exist"
	message[LoginAccountOrPasswordError] = "account or password error"
	message[LoginAccountExist] = "account is exist"
	message[LoginCaptchaError] = "captcha error"

	//chat
	message[ChatCustomerNotExist] = "customer is not exist"
	message[ChatNodeNotExist] = "node not set"
	message[ChatPromptEmpty] = "prompt is empty"
	message[ChatTimesEmpty] = "The times has been used up"
	message[ChatRelationRecordEmpty] = "relation record not found"

	//chat room
	message[ChatRoomChannelNotFound] = "chat room is not exist"
	message[ChatRoomFull] = "chat room full"
	message[ChatRoomChannelClose] = "chat room close"
	message[ChatRoomOut] = "chat room out"
	message[ChatRoomAlreadyInRoom] = "You are already in the room"

	message[RightsAmountError] = "待付款金额错误"
	message[RightsNotHaveTimesError] = "你没有可使用的次数"
	message[RightsNotVip] = "没有可用次数，请充值"

	message[QuestionnaireHaveDone] = "测评已经完成"
	message[QuestionnaireStatusError] = "测评状态错误"
	message[QuestionnaireRelationError] = "relation id error"
	message[QuestionnaireResponseError] = "请先回答完所有的问题"

}

func MapErrMsg(errCode uint32) string {
	if msg, ok := message[errCode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errCode uint32) bool {
	if _, ok := message[errCode]; ok {
		return true
	} else {
		return false
	}
}
