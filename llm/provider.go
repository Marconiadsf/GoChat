package llm

import "fmt"

type Provider interface {
	Send(messages []Message) (string, error)
}

func NewProvider(providerName, apiKey, model string) (Provider, error) {
	switch providerName {
	case "claude":
		return NewClaude(apiKey, model), nil
	case "gemini":
		return NewGemini(apiKey, model), nil
	case "ollama":
		return NewOllama(model, "http://localhost:11434"), nil
	default:
		return nil, fmt.Errorf("provider desconhecido: %s", providerName)
	}
}
