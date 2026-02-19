package main

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status   string            `json:"status"`
	Services map[string]string `json:"services"`
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	resp := HealthResponse{
		Status: "ok",
		Services: map[string]string{
			"server": "ok",
			"qdrant": "ok",
			"weaviate": "ok",
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
