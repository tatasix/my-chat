package xerr

// SUCCESS 成功返回
const SUCCESS uint32 = 200
const ErrorCommon uint32 = 1
const LogicSuccess uint32 = 0

/** 全局错误码 **/

const ServerFail uint32 = 500
const RequestParamError uint32 = 400
const UNAUTHORIZED uint32 = 401
const FORBIDDEN uint32 = 403
const RouteNotFound uint32 = 404
const RouteNotMatch uint32 = 405
const PasswordIncorrect uint32 = 406

const DBError uint32 = 600

const (
	ParamMiss       uint32 = 100001
	RecordNotFound  uint32 = 100002
	RepeatRequest   uint32 = 100003
	FileMiss        uint32 = 100004
	SystemError     uint32 = 100005
	SystemBusyError uint32 = 100006
	RequestError    uint32 = 100007
	InvalidToken    uint32 = 100008
	ParamError      uint32 = 100009
	SensitiveError  uint32 = 100011
)

const (
	ConfigExist          uint32 = 300001
	ConfigOpenAiKeyEmpty uint32 = 300002
	ConfigEmpty          uint32 = 300003
)

const (
	LoginMobileError            uint32 = 200001
	LoginVerifyCodeError        uint32 = 200002
	LoginAccountNotExist        uint32 = 200003
	LoginAccountOrPasswordError uint32 = 200004
	LoginAccountExist           uint32 = 200005
	LoginCaptchaError           uint32 = 200006
)

const (
	ChatCustomerNotExist    uint32 = 400001
	ChatNodeNotExist        uint32 = 400002
	ChatPromptEmpty         uint32 = 400003
	ChatTimesEmpty          uint32 = 400004
	ChatRelationRecordEmpty uint32 = 400006
)

const (
	ChatRoomChannelNotFound uint32 = 500001
	ChatRoomFull            uint32 = 500002
	ChatRoomChannelClose    uint32 = 500003
	ChatRoomOut             uint32 = 500004
	ChatRoomAlreadyInRoom   uint32 = 500006
)

const (
	RightsAmountError       uint32 = 600001
	RightsNotHaveTimesError uint32 = 600002
	RightsNotVip            uint32 = 600003
)

const (
	QuestionnaireHaveDone      uint32 = 700001
	QuestionnaireStatusError   uint32 = 700002
	QuestionnaireRelationError uint32 = 700003
	QuestionnaireResponseError uint32 = 700004
)
