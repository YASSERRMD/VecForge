package db

import "time"

type QdrantClient struct {
	url    string
	timeout time.Duration
}

func NewQdrantClient(url string) *QdrantClient {
	return &QdrantClient{url: url, timeout: 5 * time.Second}
}

func (c *QdrantClient) Name() string   { return "qdrant" }
func (c *QdrantClient) URL() string    { return c.url }

func (c *QdrantClient) Search(query []float32, topK int) ([]Hit, error) {
	return []Hit{{ID: "q1", Score: 0.95}}, nil
}
