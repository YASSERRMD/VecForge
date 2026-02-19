package agent

import "fmt"

type PromptBuilder struct {
	systemTemplate string
	userTemplate  string
}

func NewPromptBuilder() *PromptBuilder {
	return &PromptBuilder{
		systemTemplate: "You are a helpful assistant.",
		userTemplate:   "Context: %s\n\nQuestion: %s\n\nAnswer:",
	}
}

func (pb *PromptBuilder) Build(context, question string) string {
	return fmt.Sprintf(pb.userTemplate, context, question)
}

func (pb *PromptBuilder) WithSystemTemplate(t string) *PromptBuilder {
	pb.systemTemplate = t
	return pb
}

func (pb *PromptBuilder) WithUserTemplate(t string) *PromptBuilder {
	pb.userTemplate = t
	return pb
}

func (pb *PromptBuilder) SystemPrompt() string {
	return pb.systemTemplate
}
