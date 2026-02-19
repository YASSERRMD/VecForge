package agent

import "testing"

func TestQueryRewriter(t *testing.T) {
	r := NewQueryRewriter()
	
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello World", "hello world"},
		{"AI ML", "artificial intelligence machine learning"},
	}
	
	for _, tt := range tests {
		result := r.Rewrite(tt.input)
		if result != tt.expected {
			t.Errorf("expected %q, got %q", tt.expected, result)
		}
	}
}
