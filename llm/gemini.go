package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Gemini struct {
	APIKey string
	Model  string
}

func NewGemini(apiKey, model string) Gemini {
	return Gemini{APIKey: apiKey, Model: model}
}

type geminiRequest struct {
	Contents []geminiContent `json:"contents"`
}

type geminiContent struct {
	Role  string       `json:"role"`
	Parts []geminiPart `json:"parts"`
}

func toGeminiContents(msgs []Message) []geminiContent {
	contents := make([]geminiContent, len(msgs))
	for i, m := range msgs {
		role := m.Role
		if role == "assistant" {
			role = "model"
		}
		contents[i] = geminiContent{
			Role:  role,
			Parts: []geminiPart{{Text: m.Content}},
		}
	}
	return contents
}

type geminiPart struct {
	Text string `json:"text"`
}

type geminiResponse struct {
	Error *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func (g Gemini) Send(msgs []Message) (string, error) {
	body, err := json.Marshal(geminiRequest{
		Contents: toGeminiContents(msgs),
	})
	if err != nil {
		return "", err
	}

	url := "https://generativelanguage.googleapis.com/v1beta/models/" + g.Model + ":generateContent"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("x-goog-api-key", g.APIKey)
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

	var result geminiResponse
	err = json.Unmarshal(data, &result)
	if err != nil {
		return "", err
	}

	if result.Error != nil {
		return "", fmt.Errorf("erro da API (%d): %s", result.Error.Code, result.Error.Message)
	}

	if len(result.Candidates) == 0 || len(result.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("resposta vazia da API")
	}

	return result.Candidates[0].Content.Parts[0].Text, nil
}
