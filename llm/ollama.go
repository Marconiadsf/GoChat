package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Ollama struct {
	Model string
	Host  string
}

func NewOllama(model, host string) Ollama {
	return Ollama{Model: model, Host: host}
}

type ollamaRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type ollamaResponse struct {
	Error   string `json:"error"`
	Message struct {
		Content string `json:"content"`
	} `json:"message"`
}

func (o Ollama) Send(msgs []Message) (string, error) {
	body, err := json.Marshal(ollamaRequest{
		Model:    o.Model,
		Messages: msgs,
		Stream:   false,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", o.Host+"/api/chat", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result ollamaResponse
	err = json.Unmarshal(data, &result)
	if err != nil {
		return "", err
	}

	if result.Error != "" {
		return "", fmt.Errorf("erro do Ollama: %s", result.Error)
	}

	return result.Message.Content, nil
}
