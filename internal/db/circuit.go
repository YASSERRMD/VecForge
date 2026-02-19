package db

import (
	"errors"
	"sync"
	"time"
)

type CircuitState int

const (
	StateClosed CircuitState = iota
	StateOpen
	StateHalfOpen
)

type CircuitBreaker struct {
	mu           sync.Mutex
	state        CircuitState
	failures     int
	successes    int
	maxFailures  int
	timeout      time.Duration
	lastFailure  time.Time
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:       StateClosed,
		maxFailures: maxFailures,
		timeout:     timeout,
	}
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
	if !cb.allowRequest() {
		return errors.New("circuit open")
	}
	
	err := fn()
	cb.recordResult(err)
	return err
}

func (cb *CircuitBreaker) allowRequest() bool {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	
	switch cb.state {
	case StateClosed:
		return true
	case StateOpen:
		if time.Since(cb.lastFailure) > cb.timeout {
			cb.state = StateHalfOpen
			return true
		}
		return false
	case StateHalfOpen:
		return true
	}
	return false
}

func (cb *CircuitBreaker) recordResult(err error) {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	
	if err == nil {
		cb.successes++
		if cb.state == StateHalfOpen && cb.successes >= 3 {
			cb.state = StateClosed
			cb.successes = 0
			cb.failures = 0
		}
	} else {
		cb.failures++
		cb.lastFailure = time.Now()
		if cb.failures >= cb.maxFailures {
			cb.state = StateOpen
		}
	}
}
package db

import "testing"

func TestCircuitBreaker(t *testing.T) {
	cb := NewCircuitBreaker(3, 1)
	
	err := cb.Execute(func() error { return nil })
	if err != nil {
		t.Errorf("expected nil error, got 