package debugger

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/debugger/redaction"
	"github.com/example/a2ui-go-agent-platform/pkg/events"
	"github.com/example/a2ui-go-agent-platform/pkg/httpx"
)

type RunFilter struct {
	SessionID string
	Status    string
	From      time.Time
	To        time.Time
}

type RunSummary struct {
	RunID      string     `json:"runId"`
	SessionID  string     `json:"sessionId"`
	Status     string     `json:"status"`
	StartedAt  time.Time  `json:"startedAt"`
	EndedAt    *time.Time `json:"endedAt,omitempty"`
	DurationMS int64      `json:"durationMs"`
}

type StepView struct {
	Step      string         `json:"step"`
	Status    string         `json:"status"`
	LatencyMS int            `json:"latencyMs"`
	TokenIn   int            `json:"tokenIn"`
	TokenOut  int            `json:"tokenOut"`
	Payload   map[string]any `json:"payload,omitempty"`
}

type ToolCallView struct {
	Step      string         `json:"step"`
	Status    string         `json:"status"`
	Params    map[string]any `json:"params,omitempty"`
	Error     string         `json:"error,omitempty"`
	LatencyMS int            `json:"latencyMs"`
}

type ContextWindowView struct {
	SystemTokens  int `json:"systemTokens"`
	SessionTokens int `json:"sessionTokens"`
	TaskTokens    int `json:"taskTokens"`
	TotalTokens   int `json:"totalTokens"`
}

type CostView struct {
	TotalTokens int                `json:"totalTokens"`
	TotalCost   float64            `json:"totalCost"`
	ByModel     map[string]float64 `json:"byModel"`
	ByProvider  map[string]float64 `json:"byProvider"`
}

type RunDetail struct {
	Run     RunSummary        `json:"run"`
	Steps   []StepView        `json:"steps"`
	Tools   []ToolCallView    `json:"tools"`
	Context ContextWindowView `json:"context"`
	Cost    CostView          `json:"cost"`
	TraceID string            `json:"traceId"`
}

type Service struct {
	events *events.Service
}

func NewService(eventsSvc *events.Service) *Service {
	return &Service{events: eventsSvc}
}

func (s *Service) ListRuns(ctx context.Context, filter RunFilter) ([]RunSummary, error) {
	runs, err := s.events.ListRuns(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]RunSummary, 0, len(runs))
	for _, run := range runs {
		if !matchRun(run.SessionID, string(run.Status), run.StartedAt, filter) {
			continue
		}
		d := runDurationMS(run.StartedAt, run.EndedAt)
		out = append(out, RunSummary{RunID: run.ID, SessionID: run.SessionID, Status: string(run.Status), StartedAt: run.StartedAt, EndedAt: run.EndedAt, DurationMS: d})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].StartedAt.After(out[j].StartedAt) })
	return out, nil
}

func (s *Service) GetRunDetail(ctx context.Context, runID string) (RunDetail, error) {
	run, err := s.events.GetRun(ctx, runID)
	if err != nil {
		return RunDetail{}, err
	}
	eventsList, err := s.events.ListEventsByRun(ctx, runID)
	if err != nil {
		return RunDetail{}, err
	}
	steps := make([]StepView, 0, len(eventsList))
	tools := make([]ToolCallView, 0)
	totalIn := 0
	totalOut := 0
	for _, item := range eventsList {
		payload := redaction.RedactMap(item.Payload)
		stepStatus := statusFromStep(item.Step)
		steps = append(steps, StepView{
			Step:      item.Step,
			Status:    stepStatus,
			LatencyMS: item.LatencyMS,
			TokenIn:   item.TokenIn,
			TokenOut:  item.TokenOut,
			Payload:   payload,
		})
		totalIn += item.TokenIn
		totalOut += item.TokenOut
		if strings.Contains(strings.ToLower(item.Step), "tool") {
			tools = append(tools, ToolCallView{
				Step:      item.Step,
				Status:    stepStatus,
				LatencyMS: item.LatencyMS,
				Params:    payload,
				Error:     fmt.Sprint(payload["error"]),
			})
		}
	}
	cost := computeCost(totalIn + totalOut)
	ctxView := ContextWindowView{
		SystemTokens:  60,
		SessionTokens: max(0, totalIn/2),
		TaskTokens:    max(0, totalIn/2),
		TotalTokens:   totalIn,
	}

	traceID := httpx.TraceIDFromContext(ctx)
	return RunDetail{
		Run: RunSummary{
			RunID:      run.ID,
			SessionID:  run.SessionID,
			Status:     string(run.Status),
			StartedAt:  run.StartedAt,
			EndedAt:    run.EndedAt,
			DurationMS: runDurationMS(run.StartedAt, run.EndedAt),
		},
		Steps:   steps,
		Tools:   tools,
		Context: ctxView,
		Cost:    cost,
		TraceID: traceID,
	}, nil
}

func (s *Service) GetRunCost(ctx context.Context, runID string) (CostView, error) {
	eventsList, err := s.events.ListEventsByRun(ctx, runID)
	if err != nil {
		return CostView{}, err
	}
	total := 0
	for _, item := range eventsList {
		total += item.TokenIn + item.TokenOut
	}
	return computeCost(total), nil
}

func computeCost(tokens int) CostView {
	totalCost := float64(tokens) * 0.000002
	return CostView{
		TotalTokens: tokens,
		TotalCost:   totalCost,
		ByModel: map[string]float64{
			"mock-text": totalCost,
		},
		ByProvider: map[string]float64{
			"mock-provider": totalCost,
		},
	}
}

func matchRun(sessionID, status string, startedAt time.Time, filter RunFilter) bool {
	if filter.SessionID != "" && sessionID != filter.SessionID {
		return false
	}
	if filter.Status != "" && status != filter.Status {
		return false
	}
	if !filter.From.IsZero() && startedAt.Before(filter.From) {
		return false
	}
	if !filter.To.IsZero() && startedAt.After(filter.To) {
		return false
	}
	return true
}

func runDurationMS(start time.Time, end *time.Time) int64 {
	if end == nil {
		return time.Since(start).Milliseconds()
	}
	return end.Sub(start).Milliseconds()
}

func statusFromStep(step string) string {
	lower := strings.ToLower(step)
	if strings.Contains(lower, "failed") || strings.Contains(lower, "error") {
		return "failed"
	}
	return "completed"
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
