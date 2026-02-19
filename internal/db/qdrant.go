package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type QdrantClient struct {
	url    string
	client *http.Client
}

type QdrantSearchRequest struct {
	Query       []float32 `json:"query"`
	Limit       int       `json:"limit"`
	WithPayload bool      `json:"with_payload"`
}

type QdrantSearchResponse struct {
	Result []QdrantHit `json:"result"`
	Status string      `json:"status"`
}

type QdrantHit struct {
	ID      string          `json:"id"`
	Score   float64         `json:"score"`
	Payload map[string]any `json:"payload"`
}

func NewQdrantClient(url string) *QdrantClient {
	return &QdrantClient{
		url: url,
		client: &http.Client{Timeout: 5 * time.Second},
	}
}

func (c *QdrantClient) Name() string { return "qdrant" }

func (c *QdrantClient) Search(query []float32, topK int) ([]Hit, error) {
	reqBody := QdrantSearchRequest{
		Query:       query,
		Limit:       topK,
		WithPayload: true,
	}
	
	body, _ := json.Marshal(reqBody)
	resp, err := c.client.Post(c.url+"/collections/test/points/search", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	var result QdrantSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	
	hits := make([]Hit, len(result.Result))
	for i, r := range result.Result {
		hits[i] = Hit{ID: r.ID, Score: float32(r.Score)}
	}
	return hits, nil
}

func (c *QdrantClient) Health() error {
	resp, err := c.client.Get(c.url + "/readyz")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("qdrant not healthy")
	}
	return nil
}
