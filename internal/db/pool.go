package db

import "time"

type Pool struct {
	providers []Provider
	timeout   time.Duration
}

func NewPool(providers []Provider) *Pool {
	return &Pool{providers: providers, timeout: 5 * time.Second}
}

func (p *Pool) SearchAll(query []float32, topK int) ([]SearchResult, error) {
	ch := make(chan SearchResult, len(p.providers))
	
	for _, prov := range p.providers {
		go func(pr Provider) {
			hits, err := pr.Search(query, topK)
			if err == nil {
				ch <- SearchResult{Hits: hits, Provider: pr.Name(), Latency: 42}
			}
		}(prov)
	}
	
	results := make([]SearchResult, 0, len(p.providers))
	for range p.providers {
		results = append(results, <-ch)
	}
	return results, nil
}
