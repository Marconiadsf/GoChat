package llm

type Provider interface {
	Send(messages []Message) (string, error)
}
