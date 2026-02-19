package agent

import "encoding/json"

type LLM interface {
	Generate(prompt string) (string, error)
}

type MockLLM struct{}

func NewMockLLM() *MockLLM { return &MockLLM{} }

func (l *MockLLM) Generate(prompt string) (string, error) {
	return "This is a mock response to: " + prompt[:min(50, len(prompt))], nil
}

type OpenAILLM struct {
	apiKey string
	model  string
}

func NewOpenAILLM(apiKey string) *OpenAILLM {
	return &OpenAILLM{apiKey: apiKey, model: "gpt-4"}
}

func (l *OpenAILLM) Generate(prompt string) (string, error) {
	data, _ := json.Marshal(map[string]string{"model": l.model, "prompt": prompt})
	return string(data), nil
}
