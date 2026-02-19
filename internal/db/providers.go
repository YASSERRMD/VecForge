package db

type Provider interface {
	Search(query []float32, topK int) ([]Hit, error)
	Name() string
}

type Hit struct {
	ID    string
	Score float32
}

type SearchResult struct {
	Hits     []Hit
	Provider string
	Latency  int64
}
