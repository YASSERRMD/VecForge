package agent

import "sync"

type Storage interface {
	Save(key string, data interface{}) error
	Load(key string, data interface{}) error
	Delete(key string) error
}

type MemoryStorage struct {
	mu   sync.RWMutex
	data map[string]interface{}
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{data: make(map[string]interface{})}
}

func (s *MemoryStorage) Save(key string, data interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = data
	return nil
}

func (s *MemoryStorage) Load(key string, data interface{}) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if v, ok := s.data[key]; ok {
		return nil
	}
	return ErrNotFound
}

func (s *MemoryStorage) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
	return nil
}

var ErrNotFound = &StorageError{msg: "key not found"}

type StorageError struct{ msg string }

func (e *StorageError) Error() string { return e.msg }
