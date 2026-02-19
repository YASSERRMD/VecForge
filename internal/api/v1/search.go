package v1

import (
	"encoding/json"
	"net/http"
)

type SearchRequest struct {
	Query     string   `json:"q"`
	Providers []string `json:"providers"`
	Limit     int      `json:"limit"`
}

type Hit struct {
	ID       string      `json:"id"`
	Score    float64     `json:"score"`
	Payload  interface{} `json:"payload,omitempty"`
	Provider string      `json:"provider"`
}

type SearchResponse struct {
	Hits      []Hit  `json:"hits"`
	Query     string `json:"query"`
	LatencyUs int64  `json:"latency_us"`
	Provider  string `json:"provider"`
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	response := SearchResponse{
		Hits: []Hit{
			{ID: "demo_1", Score: 0.95, Provider: "qdrant"},
			{ID: "demo_2", Score: 0.87, Provider: "weaviate"},
		},
		Query:     req.Query,
		LatencyUs: 42,
		Provider:  "fused",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
