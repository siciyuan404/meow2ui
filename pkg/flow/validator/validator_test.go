package validator

import "testing"

func TestValidateAcceptsSimpleDAG(t *testing.T) {
	def := Definition{
		Name:   "default",
		Policy: Policy{Parallelism: 1},
		Nodes: []Node{
			{ID: "a"},
			{ID: "b", DependsOn: []string{"a"}},
		},
		Edges: []Edge{{From: "a", To: "b", Condition: "always"}},
	}
	if err := Validate(def); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestValidateRejectsCycle(t *testing.T) {
	def := Definition{
		Name:   "cycle",
		Policy: Policy{Parallelism: 1},
		Nodes: []Node{
			{ID: "a", DependsOn: []string{"b"}},
			{ID: "b", DependsOn: []string{"a"}},
		},
	}
	if err := Validate(def); err == nil {
		t.Fatalf("expected cycle error")
	}
}
