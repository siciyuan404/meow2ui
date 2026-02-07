package metrics

import "time"

type DailyMetrics struct {
	Date            string  `json:"date"`
	DAU             int     `json:"dau"`
	WorkspaceCount  int     `json:"workspaceCount"`
	SessionCount    int     `json:"sessionCount"`
	AgentSuccess    float64 `json:"agentSuccessRate"`
	AgentAvgLatency float64 `json:"agentAvgLatency"`
}

type FunnelStep struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Retention struct {
	D1 float64 `json:"d1"`
	D7 float64 `json:"d7"`
}

func DefaultDaily(now time.Time) []DailyMetrics {
	return []DailyMetrics{{
		Date:            now.Format("2006-01-02"),
		DAU:             0,
		WorkspaceCount:  0,
		SessionCount:    0,
		AgentSuccess:    0,
		AgentAvgLatency: 0,
	}}
}
