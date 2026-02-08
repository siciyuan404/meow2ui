package orchestrator

import (
	"context"
	"testing"

	"github.com/example/a2ui-go-agent-platform/pkg/evaluation"
	"github.com/example/a2ui-go-agent-platform/pkg/evaluation/runner"
)

func TestExecuteWithBudgetAndRegression(t *testing.T) {
	base := evaluation.EvalScore{SchemaValid: true, ComponentValidRate: 95, PropValidRate: 95, Success: true}
	cases := map[string][]runner.CaseResult{
		"t1": {
			{CaseID: "c1", Score: evaluation.EvalScore{SchemaValid: true, ComponentValidRate: 80, PropValidRate: 80, Success: false}, LatencyMS: 10, TokenCount: 50},
			{CaseID: "c2", Score: evaluation.EvalScore{SchemaValid: true, ComponentValidRate: 90, PropValidRate: 90, Success: true}, LatencyMS: 10, TokenCount: 60},
		},
	}
	report := Execute(context.Background(), "run-1", []Target{{ID: "t1", Name: "target1"}}, cases, base, Budget{MaxCases: 1, MaxTokens: 100})
	if len(report.Targets) != 1 {
		t.Fatalf("expected one target")
	}
	if report.Targets[0].Report.Total != 1 {
		t.Fatalf("expected max case cap applied")
	}
	if !report.Regressed {
		t.Fatalf("expected regression")
	}
}
