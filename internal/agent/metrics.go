package agent

import (
	"sync/atomic"
	"time"
)

type AgentMetrics struct {
	queries    uint64
	rewrites  uint64
	errors    uint64
	latencyNs uint64
}

func NewAgentMetrics() *AgentMetrics { return &AgentMetrics{} }

func (m *AgentMetrics) IncQueries()   { atomic.AddUint64(&m.queries, 1) }
func (m *AgentMetrics) IncRewrites()  { atomic.AddUint64(&m.rewrites, 1) }
func (m *AgentMetrics) IncErrors()    { atomic.AddUint64(&m.errors, 1) }
func (m *AgentMetrics) AddLatency(d time.Duration) {
	atomic.AddUint64(&m.latencyNs, uint64(d.Nanoseconds()))
}

func (m *AgentMetrics) Queries() uint64  { return atomic.LoadUint64(&m.queries) }
func (m *AgentMetrics) Rewrites() uint64 { return atomic.LoadUint64(&m.rewrites) }
func (m *AgentMetrics) Errors() uint64    { return atomic.LoadUint64(&m.errors) }

func (m *AgentMetrics) AvgLatencyMs() float64 {
	q := m.Queries()
	if q == 0 {
		return 0
	}
	return float64(atomic.LoadUint64(&m.latencyNs)) / float64(q) / 1e6
}
