package telemetry

import (
	"sync"
	"time"
)

type Snapshot struct {
	TotalRuns      int
	FailedRuns     int
	TotalLatencyMS int64
	LastUpdated    time.Time
}

type Service struct {
	mu      sync.Mutex
	total   int
	failed  int
	latency int64
	updated time.Time
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) RecordRun(success bool, latencyMS int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.total++
	if !success {
		s.failed++
	}
	s.latency += latencyMS
	s.updated = time.Now()
}

func (s *Service) Snapshot() Snapshot {
	s.mu.Lock()
	defer s.mu.Unlock()
	return Snapshot{
		TotalRuns:      s.total,
		FailedRuns:     s.failed,
		TotalLatencyMS: s.latency,
		LastUpdated:    s.updated,
	}
}
