package db

import (
	"fmt"
	"sync"
	"time"
)

type HealthChecker struct {
	providers map[string]Provider
	mu        sync.RWMutex
}

func NewHealthChecker() *HealthChecker {
	return &HealthChecker{providers: make(map[string]Provider)}
}

func (hc *HealthChecker) Register(p Provider) {
	hc.mu.Lock()
	defer hc.mu.Unlock()
	hc.providers[p.Name()] = p
}

func (hc *HealthChecker) CheckAll() map[string]error {
	hc.mu.RLock()
	defer hc.mu.RUnlock()
	
	results := make(map[string]error)
	for name, p := range hc.providers {
		err := p.Health()
		if err != nil {
			results[name] = err
		}
	}
	return results
}

func (hc *HealthChecker) CheckOne(name string) error {
	hc.mu.RLock()
	p, ok := hc.providers[name]
	hc.mu.RUnlock()
	
	if !ok {
		return fmt.Errorf("provider  not found", name)
	}
	return p.Health()
}

type ProviderWithHealth struct {
	Provider
	healthChecker *HealthChecker
}

func (p *ProviderWithHealth) Health() error {
	if hi, ok := p.Provider.(interface{ Health() error }); ok {
		return hi.Health()
	}
	return nil
}
