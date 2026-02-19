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

type Config struct {
	ServerPort string `json:"server_port"`
	RedisURL   string `json:"redis_url"`
	QdrantURL  string `json:"qdrant_url"`
}

func loadConfig() *Config {
	return &Config{
		ServerPort: getEnv("PORT", "8080"),
		RedisURL:   getEnv("REDIS_URL", "redis://localhost:6379"),
		QdrantURL:  getEnv("QDRANT_URL", "http://localhost:6333"),
	}
}

func getEnv(k, v string) string {
	if x := os.Getenv(k); x != "" {
		return x
	}
	return v
}

func main() {
	cfg := loadConfig()
	log.Printf("Config: port=%s", cfg.ServerPort)

	r := http.NewServeMux()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"status":"ok"}`))
	})
	r.HandleFunc("/v1/search", handleSearch)

	srv := &http.Server{Addr: ":" + cfg.ServerPort, Handler: logging(r), ReadTimeout: 15 * time.Second, WriteTimeout: 15 * time.Second}
	go srv.ListenAndServe()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	srv.Shutdown(context.Background())
	log.Println("Server stopped")
}

func logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		h.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(t))
	})
}

func handleSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"hits": []interface{}{}, "query": "", "latency_us": 42, "provider": "fused"})
}
