package db

import (
	"errors"
	"testing"
)

func TestCircuitBreaker(t *testing.T) {
	cb := NewCircuitBreaker(3, 1)
	
	err := cb.Execute(func() error { return nil })
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	
	err = cb.Execute(func() error { return errors.New("fail") })
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	
	if cb.state != StateOpen {
		t.Errorf("expected state Open, got %v", cb.state)
	}
}
