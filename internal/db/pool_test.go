package db

import (
	"testing"
	"time"
)

type mockProvider struct {
	name string
}

func (m *mockProvider) Name() string     { return m.name }
func (m *mockProvider) Search(q []float32, k int) ([]Hit, error) { return nil, nil }
func (m *mockProvider) Health() error     { return nil }

func TestLoadBalancer(t *testing.T) {
	providers := []Provider{&mockProvider{name: "a"}, &mockProvider{name: "b"}, &mockProvider{name: "c"}}
	lb := NewLoadBalancer(providers)
	
	for i := 0; i < 6; i++ {
		p := lb.Next()
		expected := []string{"a", "b", "c", "a", "b", "c"}[i]
		if p.Name() != expected {
			t.Errorf("expected %s, got %s", expected, p.Name())
		}
	}
}

func TestConnectionPool(t *testing.T) {
	providers := []Provider{&mockProvider{name: "x"}, &mockProvider{name: "y"}}
	pool := NewConnectionPool(providers, 5)
	
	p1 := pool.Get()
	p2 := pool.Get()
	
	if p1 == nil || p2 == nil {
		t.Error("expected non-nil providers")
	}
}
