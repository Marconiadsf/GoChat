package chat

import "github.com/Marconiadsf/GoChat/llm"

type Chat struct {
	provider llm.Provider
	history  []llm.Message
}

func NewChat(provider llm.Provider) Chat {
	return Chat{provider: provider}
}

func (c *Chat) Send(userMsg string) (string, error) {
	c.history = append(c.history, llm.Message{Role: "user", Content: userMsg})

	response, err := c.provider.Send(c.history)
	if err != nil {
		c.history = c.history[:len(c.history)-1]
		return "", err
	}

	c.history = append(c.history, llm.Message{Role: "assistant", Content: response})
	return response, nil
}
