package metrics

import "sync"

type Store struct {
	mu      sync.Mutex
	counter map[string]int64
}

func NewStore() *Store {
	return &Store{counter: map[string]int64{}}
}

func (s *Store) Inc(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.counter[name]++
}

func (s *Store) Get(name string) int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.counter[name]
}
