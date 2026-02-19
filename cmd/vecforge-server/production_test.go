package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCORS(t *testing.T) {
	handler := corsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}))
	
	req := httptest.NewRequest("OPTIONS", "/", nil)
	w := httptest.NewRecorder()
	
	handler.ServeHTTP(w, req)
	
	if w.Header().Get("Access-Control-Allow-Origin") == "" {
		t.Error("CORS headers not set")
	}
}

func TestRateLimit(t *testing.T) {
	rl := NewRateLimiter(2, 100)
	
	if !rl.Allow("test") {
		t.Error("first request should be allowed")
	}
	if !rl.Allow("test") {
		t.Error("second request should be allowed")
	}
	if rl.Allow("test") {
		t.Error("third request should be denied")
	}
}
