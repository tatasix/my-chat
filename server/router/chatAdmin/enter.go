package chatAdmin

type RouterGroup struct {
	ApplicationConfigRouter
	CustomerConfigRouter
	ChatRouter
	FeedbackRouter
	ConfigOpenRouter
}
