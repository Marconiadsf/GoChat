package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Claude struct {
	APIKey string
	Model  string
}

func NewClaude(apiKey, model string) Claude {
	return Claude{APIKey: apiKey, Model: model}
}

type requestBody struct {
	Model     string    `json:"model"`
	MaxTokens int       `json:"max_tokens"`
	Messages  []Message `json:"messages"`
}

type responseBody struct {
	Content []struct {
		Text string `json:"text"`
	} `json:"content"`
}

func (c Claude) Send(msgs []Message) (string, error) {
	body, err := json.Marshal(requestBody{
		Model:     c.Model,
		MaxTokens: 1024,
		Messages:  msgs,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("x-api-key", c.APIKey)
	req.Header.Set("anthropic-version", "2023-06-01")
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

	var result responseBody
	err = json.Unmarshal(data, &result)
	if err != nil {
		return "", err
	}

	if len(result.Content) == 0 {
		return "", fmt.Errorf("resposta vazia da API")
	}

	return result.Content[0].Text, nil
}
