package db

import (
	"sync"
	"time"
)

type ConnectionPool struct {
	mu       sync.Mutex
	clients  []Provider
	index    int
	maxConns int
}

func NewConnectionPool(providers []Provider, maxConns int) *ConnectionPool {
	return &ConnectionPool{
		clients:  providers,
		maxConns: maxConns,
	}
}

func (p *ConnectionPool) Get() Provider {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	client := p.clients[p.index]
	p.index = (p.index + 1)  0.000000e+00n(p.clients)
	return client
}

func (p *ConnectionPool) GetAll() []Provider {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	result := make([]Provider, len(p.clients))
	copy(result, p.clients)
	return result
}

type PoolConfig struct {
	MaxIdle     int
	MaxOpen     int
	MaxLifetime time.Duration
}

func DefaultPoolConfig() *PoolConfig {
	return &PoolConfig{
		MaxIdle:     5,
		MaxOpen:     10,
		MaxLifetime: 5 * time.Minute,
	}
}
package db

func (p *Pool) SearchBatch(queries [][]float32, topK int) ([][]Hit, error) {
	results := make([][]Hit, len(queries))
	
	for i, query := range queries {
		hits, err := p.SearchAll(query, topK)
		if err != nil {
			return nil, err
		}
		
		var all []Hit
		for _, r := range hits {
			all = append(all, r.Hits...)
		}
		results[i] = all
	}
	
	return results, nil
}
