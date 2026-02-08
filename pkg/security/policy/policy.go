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

func (e *Engine) ValidateMediaRef(ref string, allowedHosts []string) Decision {
	r := strings.TrimSpace(strings.ToLower(ref))
	if r == "" {
		return Decision{Allowed: false, Risk: "low", Reason: "empty media ref", RuleID: "POL-MEDIA-EMPTY"}
	}
	blockedPrefixes := []string{"http://127.", "http://10.", "http://192.168.", "http://169.254.", "http://localhost", "https://localhost"}
	for _, p := range blockedPrefixes {
		if strings.HasPrefix(r, p) {
			return Decision{Allowed: false, Risk: "high", Reason: "blocked internal address", RuleID: "POL-MEDIA-SSRF-BLOCK"}
		}
	}
	if strings.HasPrefix(r, "http://") || strings.HasPrefix(r, "https://") {
		if len(allowedHosts) == 0 {
			return Decision{Allowed: false, Risk: "medium", Reason: "host not in allow list", RuleID: "POL-MEDIA-HOST-BLOCK"}
		}
		for _, host := range allowedHosts {
			h := strings.TrimSpace(strings.ToLower(host))
			if h != "" && strings.Contains(r, h) {
				return Decision{Allowed: true, Risk: "low", Reason: "media ref allowed", RuleID: "POL-MEDIA-ALLOW"}
			}
		}
		return Decision{Allowed: false, Risk: "medium", Reason: "host not in allow list", RuleID: "POL-MEDIA-HOST-BLOCK"}
	}
	if strings.HasPrefix(r, "s3://") || strings.HasPrefix(r, "oss://") || strings.HasPrefix(r, "local://") {
		return Decision{Allowed: true, Risk: "low", Reason: "media ref allowed", RuleID: "POL-MEDIA-ALLOW"}
	}
	return Decision{Allowed: false, Risk: "medium", Reason: "unsupported media ref", RuleID: "POL-MEDIA-UNSUPPORTED"}
}
