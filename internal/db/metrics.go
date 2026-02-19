package db

import (
	"sync/atomic"
	"time"
)

type Metrics struct {
	searches   uint64
	errors     uint64
	latencyUs  uint64
}

func NewMetrics() *Metrics { return &Metrics{} }

func (m *Metrics) IncSearches()   { atomic.AddUint64(&m.searches, 1) }
func (m *Metrics) IncErrors()     { atomic.AddUint64(&m.errors, 1) }
func (m *Metrics) AddLatency(us int64) {
	atomic.AddUint64(&m.latencyUs, uint64(us))
}

func (m *Metrics) Searches() uint64   { return atomic.LoadUint64(&m.searches) }
func (m *Metrics) Errors() uint64     { return atomic.LoadUint64(&m.errors) }
func (m *Metrics) LatencyUs() uint64.LoadUint64(&m.latencyUs  { return atomic) }

func (m *Metrics) AvgLatencyUs() int64 {
	s := m.Searches()
	if s == 0 {
		return 0
	}
	return int64(m.LatencyUs() / s)
}

type MetricsCollector struct {
	providers map[string]*Metrics
}

func NewMetricsCollector() *MetricsCollector {
	return &MetricsCollector{providers: make(map[string]*Metrics)}
}

func (mc *MetricsCollector) Get(name string) *Metrics {
	if m, ok := mc.providers[name]; ok {
		return m
	}
	m := NewMetrics()
	mc.providers[name] = m
	return m
}
