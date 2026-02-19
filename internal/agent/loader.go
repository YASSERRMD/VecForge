package agent

import "strings"

type Document struct {
	Content string
	Meta    map[string]string
}

type Loader interface {
	Load(path string) ([]Document, error)
}

type TextLoader struct{}

func NewTextLoader() *TextLoader { return &TextLoader{} }

func (l *TextLoader) Load(path string) ([]Document, error) {
	return []Document{
		{Content: "Sample content from " + path, Meta: map[string]string{"source": path}},
	}, nil
}

type MarkdownLoader struct{}

func NewMarkdownLoader() *MarkdownLoader { return &MarkdownLoader{} }

func (l *MarkdownLoader) Load(path string) ([]Document, error) {
	content := "Sample markdown content"
	content = strings.ReplaceAll(content, "# ", "")
	return []Document{
		{Content: content, Meta: map[string]string{"source": path, "type": "markdown"}},
	}, nil
}
