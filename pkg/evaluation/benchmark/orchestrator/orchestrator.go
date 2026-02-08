package orchestrator

import (
	"context"
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/evaluation"
	"github.com/example/a2ui-go-agent-platform/pkg/evaluation/runner"
)

type Target struct {
	ID   string
	Name string
	Kind string
}

type Budget struct {
	MaxCases  int
	MaxTokens int
	MaxCost   float64
}

type TargetRun struct {
	TargetID string
	Report   runner.RunReport
	Error    string
}

type AggregateReport struct {
	RunID       string
	StartedAt   time.Time
	CompletedAt time.Time
	Targets     []TargetRun
	Regressed   bool
	Summary     string
}

func Execute(ctx context.Context, runID string, targets []Target, cases map[string][]runner.CaseResult, baseline evaluation.EvalScore, budget Budget) AggregateReport {
	_ = ctx
	start := time.Now()
	out := AggregateReport{RunID: runID, StartedAt: start, Targets: make([]TargetRun, 0, len(targets))}
	for _, target := range targets {
		items := append([]runner.CaseResult(nil), cases[target.ID]...)
		if budget.MaxCases > 0 && len(items) > budget.MaxCases {
			items = items[:budget.MaxCases]
		}
		tokenSum := 0
		trimmed := make([]runner.CaseResult, 0, len(items))
		for _, c := range items {
			if budget.MaxTokens > 0 && tokenSum+c.TokenCount > budget.MaxTokens {
				break
			}
			tokenSum += c.TokenCount
			trimmed = append(trimmed, c)
		}
		report := runner.BuildReport(runID+"-"+target.ID, trimmed)
		tr := TargetRun{TargetID: target.ID, Report: report}
		if report.Total == 0 {
			tr.Error = "no cases executed"
		}
		if report.Total > 0 {
			score := report.Results[0].Score
			if evaluation.Regressed(score, baseline) {
				out.Regressed = true
			}
		}
		out.Targets = append(out.Targets, tr)
	}
	out.CompletedAt = time.Now()
	if out.Regressed {
		out.Summary = "regression detected"
	} else {
		out.Summary = "benchmark passed"
	}
	return out
}
