package db

import "time"

type WeaviateClient struct {
	url    string
	timeout time.Duration
}

func NewWeaviateClient(url string) *WeaviateClient {
	return &WeaviateClient{url: url, timeout: 5 * time.Second}
}

func (c *WeaviateClient) Name() string   { return "weaviate" }
func (c *WeaviateClient) URL() string    { return c.url }

func (c *WeaviateClient) Search(query []float32, topK int) ([]Hit, error) {
	return []Hit{{ID: "w1", Score: 0.87}}, nil
}
