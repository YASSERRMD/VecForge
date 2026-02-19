package main

import (
	"encoding/json"
	"net/http"
	"sync/atomic"
)

var metrics = &ServerMetrics{}

type ServerMetrics struct {
	queries   uint64
	errors    uint64
	latencyUs uint64
}

func (m *ServerMetrics) IncQueries()   { atomic.AddUint64(&m.queries, 1) }
func (m *ServerMetrics) IncErrors()    { atomic.AddUint64(&m.errors, 1) }
func (m *ServerMetrics) AddLatency(us int64) {
	atomic.AddUint64(&m.latencyUs, uint64(us))
}

func handleMetrics(w http.ResponseWriter, r *http.Request) {
	q := atomic.LoadUint64(&metrics.queries)
	e := atomic.LoadUint64(&metrics.errors)
	l := atomic.LoadUint64(&metrics.latencyUs)
	
	avgLatency := float64(0)
	if q > 0 {
		avgLatency = float64(l) / float64(q) / 1000
	}
	
	json.NewEncoder(w).Encode(map[string]interface{}{
		"queries":      q,
		"errors":       e,
		"avg_latency": avgLatency,
	})
}
