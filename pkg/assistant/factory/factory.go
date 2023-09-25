package factory

type IAssistant interface {
	GetChatResponse(sourceCode string, Language string) (string, error)
}
