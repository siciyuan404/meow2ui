package events

import (
	"context"
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/domain"
	"github.com/example/a2ui-go-agent-platform/pkg/store"
	"github.com/example/a2ui-go-agent-platform/pkg/util"
)

type Service struct {
	repo store.EventRepository
}

func NewService(repo store.EventRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) StartRun(ctx context.Context, sessionID, request string) (domain.AgentRun, error) {
	run := domain.AgentRun{
		ID:          util.NewID("run"),
		SessionID:   sessionID,
		RequestText: request,
		Status:      domain.AgentRunRunning,
		StartedAt:   time.Now(),
	}
	if err := s.repo.CreateRun(ctx, run); err != nil {
		return domain.AgentRun{}, err
	}
	return run, nil
}

func (s *Service) CompleteRun(ctx context.Context, runID string, success bool) error {
	run, err := s.repo.GetRun(ctx, runID)
	if err != nil {
		return err
	}
	now := time.Now()
	run.EndedAt = &now
	if success {
		run.Status = domain.AgentRunCompleted
	} else {
		run.Status = domain.AgentRunFailed
	}
	return s.repo.UpdateRun(ctx, run)
}

func (s *Service) Emit(ctx context.Context, runID, step string, payload map[string]any, latencyMS, tokenIn, tokenOut int) error {
	e := domain.AgentEvent{
		ID:        util.NewID("evt"),
		RunID:     runID,
		Step:      step,
		Payload:   payload,
		LatencyMS: latencyMS,
		TokenIn:   tokenIn,
		TokenOut:  tokenOut,
		CreatedAt: time.Now(),
	}
	return s.repo.CreateEvent(ctx, e)
}
