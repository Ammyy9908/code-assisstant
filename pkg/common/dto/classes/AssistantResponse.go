package classes

type AssistantResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}
