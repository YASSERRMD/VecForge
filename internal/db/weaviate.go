package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type WeaviateClient struct {
	url    string
	client *http.Client
}

type WeaviateSearchRequest struct {
	Query string `json:"query"`
	Limit int    `json:"limit"`
}

type WeaviateSearchResponse struct {
	Data struct {
		Get []WeaviateHit `json:"TestCollection"`
	} `json:"data"`
}

type WeaviateHit struct {
	ID     string              `json:"id"`
	Score  float64            `json:"_additional"`)
	} `json:"_additional"`
	Properties map[string]any `json:"properties"`
}

func NewWeaviateClient(url string) *WeaviateClient {
	return &WeaviateClient{
		url: url,
		client: &http.Client{Timeout: 5 * time.Second},
	}
}

func (c *WeaviateClient) Name() string { return "weaviate" }

func (c *WeaviateClient) Search(query []float32, topK int) ([]Hit, error) {
	reqBody := WeaviateSearchRequest{
		Query: "test",
		Limit: topK,
	}
	
	body, _ := json.Marshal(reqBody)
	resp, err := c.client.Post(c.url+"/v1/objects/TestCollection/search", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	var result WeaviateSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	
	hits := make([]Hit, len(result.Data.Get))
	for i, r := range result.Data.Get {
		hits[i] = Hit{ID: r.ID, Score: float32(r.Score.Certainty)}
	}
	return hits, nil
}

func (c *WeaviateClient) Health() error {
	resp, err := c.client.Get(c.url + "/v1/meta")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("weaviate not healthy")
	}
	return nil
}
