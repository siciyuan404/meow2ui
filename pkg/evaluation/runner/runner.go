package runner

import "github.com/example/a2ui-go-agent-platform/pkg/evaluation"

type CaseResult struct {
	CaseID     string
	Score      evaluation.EvalScore
	LatencyMS  int
	TokenCount int
}

type RunReport struct {
	RunID        string
	Total        int
	Passed       int
	AverageDelay int
	Results      []CaseResult
}

func BuildReport(runID string, results []CaseResult) RunReport {
	passed := 0
	totalDelay := 0
	for _, r := range results {
		if r.Score.Success {
			passed++
		}
		totalDelay += r.LatencyMS
	}
	avg := 0
	if len(results) > 0 {
		avg = totalDelay / len(results)
	}
	return RunReport{RunID: runID, Total: len(results), Passed: passed, AverageDelay: avg, Results: results}
}
