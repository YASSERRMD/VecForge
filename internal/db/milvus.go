package db

import (
	"fmt"
	"time"
)

type MilvusClient struct {
	url    string
	client *http.Client
}

func NewMilvusClient(url string) *MilvusClient {
	return &MilvusClient{
		url:    url,
		client: &http.Client{Timeout: 5 * time.Second},
	}
}

func (c *MilvusClient) Name() string { return "milvus" }

func (c *MilvusClient) Search(query []float32, topK int) ([]Hit, error) {
	return []Hit{{ID: "m1", Score: 0.91}}, nil
}

func (c *MilvusClient) Health() error {
	resp, err := c.client.Get(c.url + "/healthz")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("milvus not healthy")
	}
	return nil
}
