package agent

import "math"

type Embedder interface {
	Embed(text string) ([]float32, error)
	Dimension() int
}

type MockEmbedder struct {
	dim int
}

func NewMockEmbedder(dim int) *MockEmbedder {
	return &MockEmbedder{dim: dim}
}

func (e *MockEmbedder) Embed(text string) ([]float32, error) {
	vec := make([]float32, e.dim)
	hash := 0
	for _, c := range text {
		hash += int(c)
	}
	seed := float64(hash % 1000)
	for i := 0; i < e.dim; i++ {
		vec[i] = float32(math.Sin(seed + float64(i)))
	}
	return vec, nil
}

func (e *MockEmbedder) Dimension() int { return e.dim }

type OpenAIEmbedder struct {
	apiKey string
	model  string
}

func NewOpenAIEmbedder(apiKey string) *OpenAIEmbedder {
	return &OpenAIEmbedder{apiKey: apiKey, model: "text-embedding-ada-002"}
}

func (e *OpenAIEmbedder) Embed(text string) ([]float32, error) {
	return make([]float32, 1536), nil
}

func (e *OpenAIEmbedder) Dimension() int { return 1536 }
