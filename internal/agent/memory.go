package agent

import (
	"sync"
	"time"
)

type Memory struct {
	mu       sync.RWMutex
	sessions map[string]*Session
}

type Session struct {
	ID        string
	Queries   []string
	Results   []string
	LastActive time.Time
}

func NewMemory() *Memory {
	return &Memory{sessions: make(map[string]*Session)}
}

func (m *Memory) GetOrCreate(id string) *Session {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	if s, ok := m.sessions[id]; ok {
		s.LastActive = time.Now()
		return s
	}
	
	s := &Session{ID: id, LastActive: time.Now()}
	m.sessions[id] = s
	return s
}

func (m *Memory) AddQuery(sid, query string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if s, ok := m.sessions[sid]; ok {
		s.Queries = append(s.Queries, query)
		s.LastActive = time.Now()
	}
}

func (m *Memory) AddResult(sid, result string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if s, ok := m.sessions[sid]; ok {
		s.Results = append(s.Results, result)
	}
}

func (m *Memory) CleanOld(maxAge time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	
	now := time.Now()
	for id, s := range m.sessions {
		if now.Sub(s.LastActive) > maxAge {
			delete(m.sessions, id)
		}
	}
}
