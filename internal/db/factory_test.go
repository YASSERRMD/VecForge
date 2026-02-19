package db

import "testing"

func TestFactory(t *testing.T) {
	p, err := CreateProvider("qdrant", "http://localhost:6333")
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if p == nil {
		t.Error("expected non-nil provider")
	}
	if p.Name() != "qdrant" {
		t.Errorf("expected qdrant, got %s", p.Name())
	}
}

func TestFactoryUnknown(t *testing.T) {
	_, err := CreateProvider("unknown", "http://localhost:9999")
	if err == nil {
		t.Error("expected error for unknown provider")
	}
}
