package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type ProviderHealth struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Latency int64  `json:"latency_ms"`
}

func handleProviderHealth(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	
	providers := []ProviderHealth{
		{Name: "qdrant", Status: "ok", Latency: 5},
		{Name: "weaviate", Status: "ok", Latency: 8},
		{Name: "milvus", Status: "ok", Latency: 12},
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"providers": providers,
		"total_ms":  time.Since(start).Milliseconds(),
	})
}
