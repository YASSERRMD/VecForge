package db

import (
	"context"
	"errors"
	"time"
)

type TimeoutProvider struct {
	Provider
	timeout time.Duration
}

func WithTimeout(p Provider, timeout time.Duration) *TimeoutProvider {
	return &TimeoutProvider{Provider: p, timeout: timeout}
}

func (p *TimeoutProvider) Search(query []float32, topK int) ([]Hit, error) {
	type result struct {
		hits []Hit
		err  error
	}
	
	ch := make(chan result, 1)
	
	go func() {
		hits, err := p.Provider.Search(query, topK)
		ch <- result{hits: hits, err: err}
	}()
	
	select {
	case r := <-ch:
		return r.hits, r.err
	case <-time.After(p.timeout):
		return nil, errors.New("search timeout")
	}
}

func SearchWithContext(ctx context.Context, p Provider, query []float32, topK int) ([]Hit, error) {
	type result struct {
		hits []Hit
		err  error
	}
	
	ch := make(chan result, 1)
	
	go func() {
		hits, err := p.Search(query, topK)
		ch <- result{hits: hits, err: err}
	}()
	
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case r := <-ch:
		return r.hits, r.err
	}
}
