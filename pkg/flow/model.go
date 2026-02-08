package flow

import "github.com/example/a2ui-go-agent-platform/pkg/provider"

type FailureMode string

const (
	FailureModeFailFast        FailureMode = "fail_fast"
	FailureModeContinueOnError FailureMode = "continue_on_error"
)

type NodeType string

const (
	NodeTypePlan     NodeType = "plan"
	NodeTypeEmit     NodeType = "emit"
	NodeTypeValidate NodeType = "validate"
	NodeTypeRepair   NodeType = "repair"
	NodeTypeApply    NodeType = "apply"
	NodeTypeCustom   NodeType = "custom"
)

type Node struct {
	ID        string   `json:"id"`
	Type      NodeType `json:"type"`
	DependsOn []string `json:"depends_on,omitempty"`
	TimeoutMS int      `json:"timeout_ms,omitempty"`
}

type Edge struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Condition string `json:"condition,omitempty"`
}

type Policy struct {
	Parallelism int         `json:"parallelism"`
	FailureMode FailureMode `json:"failure_mode"`
}

type Definition struct {
	Name   string `json:"name"`
	Policy Policy `json:"policy"`
	Nodes  []Node `json:"nodes"`
	Edges  []Edge `json:"edges"`
}

type NodeResult struct {
	NodeID    string                    `json:"node_id"`
	Step      string                    `json:"step"`
	Status    string                    `json:"status"`
	LatencyMS int64                     `json:"latency_ms"`
	Error     string                    `json:"error,omitempty"`
	Tokens    int                       `json:"tokens,omitempty"`
	Output    map[string]any            `json:"output,omitempty"`
	Generated provider.GenerateResponse `json:"-"`
}

type RunResult struct {
	Success bool         `json:"success"`
	Nodes   []NodeResult `json:"nodes"`
}

func DefaultDefinition() Definition {
	return Definition{
		Name: "default",
		Policy: Policy{
			Parallelism: 1,
			FailureMode: FailureModeFailFast,
		},
		Nodes: []Node{
			{ID: "plan", Type: NodeTypePlan},
			{ID: "emit", Type: NodeTypeEmit, DependsOn: []string{"plan"}},
			{ID: "validate", Type: NodeTypeValidate, DependsOn: []string{"emit"}},
			{ID: "apply", Type: NodeTypeApply, DependsOn: []string{"validate"}},
		},
		Edges: []Edge{
			{From: "plan", To: "emit"},
			{From: "emit", To: "validate"},
			{From: "validate", To: "apply", Condition: "always"},
		},
	}
}
