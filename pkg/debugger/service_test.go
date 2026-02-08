package debugger

import (
	"context"
	"testing"
	"time"

	"github.com/example/a2ui-go-agent-platform/internal/infra/memorystore"
	"github.com/example/a2ui-go-agent-platform/pkg/domain"
	"github.com/example/a2ui-go-agent-platform/pkg/events"
)

func TestServiceListRunsAndDetail(t *testing.T) {
	ctx := context.Background()
	store := memorystore.New()
	eventSvc := events.NewService(store.Event())
	dbg := NewService(eventSvc)

	run := domain.AgentRun{ID: "run-1", SessionID: "ssn-1", RequestText: "hello", Status: domain.AgentRunCompleted, StartedAt: time.Now().Add(-2 * time.Second)}
	if err := store.Event().CreateRun(ctx, run); err != nil {
		t.Fatalf("create run: %v", err)
	}
	if err := store.Event().CreateEvent(ctx, domain.AgentEvent{ID: "evt-1", RunID: run.ID, Step: "plan", Payload: map[string]any{"authorization": "Bearer abcdefghijklmnopqrstuvwxyz123"}, LatencyMS: 5, TokenIn: 100, TokenOut: 30, CreatedAt: time.Now()}); err != nil {
		t.Fatalf("create event: %v", err)
	}

	items, err := dbg.ListRuns(ctx, RunFilter{SessionID: "ssn-1"})
	if err != nil {
		t.Fatalf("list runs: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("expected 1 run, got %d", len(items))
	}

	detail, err := dbg.GetRunDetail(ctx, run.ID)
	if err != nil {
		t.Fatalf("detail: %v", err)
	}
	if len(detail.Steps) != 1 {
		t.Fatalf("expected 1 step, got %d", len(detail.Steps))
	}
	if detail.Steps[0].Payload["authorization"] != "***REDACTED***" {
		t.Fatalf("payload should be redacted")
	}
	if detail.Cost.TotalTokens <= 0 {
		t.Fatalf("expected positive token cost")
	}
}
