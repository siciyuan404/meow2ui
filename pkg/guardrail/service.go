package guardrail

import (
	"fmt"
	"path/filepath"
	"strings"
)

type RiskLevel string

const (
	RiskLow    RiskLevel = "low"
	RiskMedium RiskLevel = "medium"
	RiskHigh   RiskLevel = "high"
)

type ToolAction struct {
	Kind string
	Path string
	Args map[string]any
}

type Policy struct {
	AllowExec      bool
	AllowNetwork   bool
	AllowedRootDir string
}

type Result struct {
	Allowed bool
	Risk    RiskLevel
	Reason  string
}

type Service struct{}

func NewService() *Service { return &Service{} }

func (s *Service) CheckPromptInjection(input string) Result {
	l := strings.ToLower(input)
	patterns := []string{
		"ignore previous instructions",
		"reveal system prompt",
		"run shell command",
		"delete all files",
	}
	for _, p := range patterns {
		if strings.Contains(l, p) {
			return Result{Allowed: false, Risk: RiskHigh, Reason: "potential prompt injection detected"}
		}
	}
	return Result{Allowed: true, Risk: RiskLow, Reason: "ok"}
}

func (s *Service) CheckToolAction(action ToolAction, policy Policy) Result {
	kind := strings.ToLower(action.Kind)
	switch kind {
	case "read":
		if policy.AllowedRootDir == "" {
			return Result{Allowed: true, Risk: RiskLow, Reason: "read allowed"}
		}
		clean := filepath.Clean(action.Path)
		if !strings.HasPrefix(clean, filepath.Clean(policy.AllowedRootDir)) {
			return Result{Allowed: false, Risk: RiskHigh, Reason: "path outside allowed root"}
		}
		return Result{Allowed: true, Risk: RiskLow, Reason: "read allowed"}
	case "write":
		if policy.AllowedRootDir == "" {
			return Result{Allowed: false, Risk: RiskMedium, Reason: "write blocked without allowed root"}
		}
		clean := filepath.Clean(action.Path)
		if !strings.HasPrefix(clean, filepath.Clean(policy.AllowedRootDir)) {
			return Result{Allowed: false, Risk: RiskHigh, Reason: "write path outside allowed root"}
		}
		return Result{Allowed: true, Risk: RiskMedium, Reason: "write allowed in allowed root"}
	case "exec":
		if !policy.AllowExec {
			return Result{Allowed: false, Risk: RiskHigh, Reason: "exec disabled by policy"}
		}
		return Result{Allowed: true, Risk: RiskHigh, Reason: "exec allowed by policy"}
	case "network":
		if !policy.AllowNetwork {
			return Result{Allowed: false, Risk: RiskHigh, Reason: "network disabled by policy"}
		}
		return Result{Allowed: true, Risk: RiskMedium, Reason: "network allowed by policy"}
	default:
		return Result{Allowed: false, Risk: RiskMedium, Reason: fmt.Sprintf("unknown action kind: %s", action.Kind)}
	}
}
