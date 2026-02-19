package db

import (
	"context"
	"sync"
	"time"
)

type ProviderWrapper struct {
	Provider
	cache      Cache
	metrics    *Metrics
	circuitBreaker *CircuitBreaker
}

func WrapWithCache(p Provider, cache Cache) *ProviderWrapper {
	return &ProviderWrapper{
		Provider: p,
		cache:    cache,
		metrics:  NewMetrics(),
	}
}

func WrapWithCircuitBreaker(p Provider, cb *CircuitBreaker) *ProviderWrapper {
	return &ProviderWrapper{
		Provider:       p,
		circuitBreaker: cb,
		metrics:       NewMetrics(),
	}
}

func (w *ProviderWrapper) Search(query []float32, topK int) ([]Hit, error) {
	start := time.Now()
	
	if w.circuitBreaker != nil {
		err := w.circuitBreaker.Execute(func() error {
			return w.doSearch(query, topK)
		})
		w.metrics.AddLatency(time.Since(start).Microseconds())
		return nil, err
	}
	
	return w.doSearch(query, topK)
}

func (w *ProviderWrapper) doSearch(query []float32, topK int) ([]Hit, error) {
	w.metrics.IncSearches()
	hits, err := w.Provider.Search(query, topK)
	if err != nil {
		w.metrics.IncErrors()
	}
	w.metrics.AddLatency(time.Since(start).Microseconds())
	return hits, err
}

type PoolWithMetrics struct {
	pool   *ConnectionPool
	metrics map[string]*Metrics
	mu     sync.Mutex
}

func NewPoolWithMetrics(pool *ConnectionPool) *PoolWithMetrics {
	return &PoolWithMetrics{pool: pool, metrics: make(map[string]*Metrics)}
}

func (p *PoolWithMetrics) GetMetrics(name string) *Metrics {
	p.mu.Lock()
	defer p.mu.Unlock()
	if m, ok := p.metrics[name]; ok {
		return m
	}
	m := NewMetrics()
	p.metrics[name] = m
	return m
}
