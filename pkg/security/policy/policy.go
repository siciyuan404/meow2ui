package policy

import "strings"

type Action struct {
	Kind string
	Path string
}

type Decision struct {
	Allowed bool
	Risk    string
	Reason  string
	RuleID  string
}

type Engine struct{}

func NewEngine() *Engine { return &Engine{} }

func (e *Engine) Decide(a Action) Decision {
	k := strings.ToLower(strings.TrimSpace(a.Kind))
	switch k {
	case "read":
		return Decision{Allowed: true, Risk: "low", Reason: "read allowed", RuleID: "POL-READ-ALLOW"}
	case "write":
		return Decision{Allowed: true, Risk: "medium", Reason: "write allowed", RuleID: "POL-WRITE-ALLOW"}
	case "exec", "network":
		return Decision{Allowed: false, Risk: "high", Reason: "high risk action blocked", RuleID: "POL-HIGH-BLOCK"}
	default:
		return Decision{Allowed: false, Risk: "medium", Reason: "unknown action", RuleID: "POL-UNKNOWN-BLOCK"}
	}
}
