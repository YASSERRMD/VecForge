package db

import (
	"time"
)

type ProviderWrapper struct {
	Provider
	cache        Cache
	metrics      *Metrics
	circuitBreak *CircuitBreaker
}

func WrapWithCache(p Provider, c Cache) *ProviderWrapper {
	return &ProviderWrapper{Provider: p, cache: c, metrics: NewMetrics()}
}

func WrapWithCircuit(p Provider, cb *CircuitBreaker) *ProviderWrapper {
	return &ProviderWrapper{Provider: p, circuitBreak: cb, metrics: NewMetrics()}
}

func (w *ProviderWrapper) Search(query []float32, topK int) ([]Hit, error) {
	start := time.Now()
	w.metrics.IncSearches()
	
	hits, err := w.Provider.Search(query, topK)
	if err != nil {
		w.metrics.IncErrors()
	}
	w.metrics.AddLatency(time.Since(start).Microseconds())
	return hits, err
}

func (w *ProviderWrapper) Metrics() *Metrics { return w.metrics }
