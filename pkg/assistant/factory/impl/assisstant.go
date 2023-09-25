package impl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Ammyy9908/code-assisstant/pkg/common/dto/classes"
)

type Assistant struct {
	APIKey string
	Model  string
}

type Payload struct {
	Messages []classes.Message `json:"messages"`
	Model    string            `json:"model"`
}

func NewAssistant(apiKey, model string) *Assistant {

	return &Assistant{
		APIKey: apiKey,
		Model:  model,
	}
}

func (a *Assistant) GetChatResponse(sourceCode, Language string) (string, error) {

	//check if api key or model is not defined
	if a.APIKey == "" || a.Model == "" {
		return "", fmt.Errorf("api key or model is not defined")
	}
	messages := []classes.Message{}
	systemMessage := classes.Message{
		Role:    "system",
		Content: "You are an employee at google and you are asked to review the following code",
	}

	userMessage := classes.Message{
		Role:    "user",
		Content: "Can you tell me the time and space complexity of this code and Review the following  " + Language + " code  and tell me if there are any syntax issues, without suggesting corrections:?" + sourceCode,
	}
	messages = append(messages, systemMessage)
	messages = append(messages, userMessage)
	data := Payload{
		Messages: messages,
		Model:    a.Model,
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(payloadBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+a.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var responseData classes.AssistantResponse
	err = json.Unmarshal(bodyBytes, &responseData)
	if err != nil {
		return "", err
	}

	if len(responseData.Choices) > 0 {
		return responseData.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response found")
}
