package alerting

import "time"

type Alert struct {
	RuleID     string
	Status     string
	Message    string
	StartedAt  time.Time
	ResolvedAt *time.Time
}

func Fire(ruleID, message string) Alert {
	return Alert{RuleID: ruleID, Status: "firing", Message: message, StartedAt: time.Now()}
}

func Resolve(a *Alert) {
	now := time.Now()
	a.Status = "resolved"
	a.ResolvedAt = &now
}
