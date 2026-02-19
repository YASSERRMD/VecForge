package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type SearchRequest struct {
	Query     string   `json:"q"`
	Providers []string `json:"providers"`
	Limit     int      `json:"limit"`
}

type Hit struct {
	ID       string  `json:"id"`
	Score   float64 `json:"score"`
	Provider string `json:"provider"`
}

type SearchResponse struct {
	Hits      []Hit  `json:"hits"`
	Query     string `json:"query"`
	LatencyUs int64  `json:"latency_us"`
	Provider  string `json:"provider"`
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{\"status\":\"ok\"}"))
	})

	router.HandleFunc("/v1/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var req SearchRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}
		if req.Limit == 0 {
			req.Limit = 10
		}
		resp := SearchResponse{
			Hits: []Hit{
				{ID: "doc_1", Score: 0.95, Provider: "qdrant"},
				{ID: "doc_2", Score: 0.87, Provider: "weaviate"},
			},
			Query:     req.Query,
			LatencyUs: 42,
			Provider:  "fused",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      requestIDMiddleware(loggingMiddleware(router)),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	go func() {
		log.Println("VecForge server starting on port 8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := time.Now().Format("20060102150405")
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r)
	})
}
