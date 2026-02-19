package agent

import "strings"

type ContextBuilder struct {
	maxTokens int
}

func NewContextBuilder() *ContextBuilder {
	return &ContextBuilder{maxTokens: 2000}
}

func (cb *ContextBuilder) Build(hits []Hit) string {
	var sb strings.Builder
	sb.WriteString("Relevant information:\n")
	
	for i, h := range hits {
		sb.WriteString(h.Content)
		sb.WriteString("\n---\n")
		if i >= 5 {
			break
		}
	}
	
	return sb.String()
}

type Hit struct {
	Content string
	Score   float32
	Source  string
}

func (cb *ContextBuilder) WithMaxTokens(tokens int) *ContextBuilder {
	cb.maxTokens = tokens
	return cb
}
