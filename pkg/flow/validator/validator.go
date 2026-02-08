package validator

import (
	"fmt"
	"strings"
)

type Definition struct {
	Name   string
	Policy Policy
	Nodes  []Node
	Edges  []Edge
}

type Policy struct {
	Parallelism int
}

type Node struct {
	ID        string
	DependsOn []string
}

type Edge struct {
	From      string
	To        string
	Condition string
}

func Validate(def Definition) error {
	if strings.TrimSpace(def.Name) == "" {
		return fmt.Errorf("flow name is required")
	}
	if len(def.Nodes) == 0 {
		return fmt.Errorf("flow nodes are required")
	}
	if def.Policy.Parallelism <= 0 {
		return fmt.Errorf("parallelism must be > 0")
	}
	nodes := map[string]Node{}
	for _, n := range def.Nodes {
		if strings.TrimSpace(n.ID) == "" {
			return fmt.Errorf("node id is required")
		}
		if _, ok := nodes[n.ID]; ok {
			return fmt.Errorf("duplicate node id: %s", n.ID)
		}
		nodes[n.ID] = n
	}
	for _, n := range def.Nodes {
		for _, dep := range n.DependsOn {
			if _, ok := nodes[dep]; !ok {
				return fmt.Errorf("unknown dependency %s for node %s", dep, n.ID)
			}
		}
	}
	for _, e := range def.Edges {
		if _, ok := nodes[e.From]; !ok {
			return fmt.Errorf("unknown edge from node: %s", e.From)
		}
		if _, ok := nodes[e.To]; !ok {
			return fmt.Errorf("unknown edge to node: %s", e.To)
		}
		if strings.TrimSpace(e.Condition) != "" {
			if err := validateCondition(e.Condition); err != nil {
				return fmt.Errorf("invalid condition on edge %s->%s: %w", e.From, e.To, err)
			}
		}
	}
	if hasCycle(def.Nodes) {
		return fmt.Errorf("flow has cycle")
	}
	return nil
}

func validateCondition(condition string) error {
	allowed := []string{"validate.ok", "!validate.ok", "always", "success", "failed"}
	for _, token := range allowed {
		if condition == token {
			return nil
		}
	}
	return fmt.Errorf("unsupported condition")
}

func hasCycle(nodes []Node) bool {
	graph := make(map[string][]string)
	for _, n := range nodes {
		for _, dep := range n.DependsOn {
			graph[dep] = append(graph[dep], n.ID)
		}
		if _, ok := graph[n.ID]; !ok {
			graph[n.ID] = nil
		}
	}
	seen := map[string]bool{}
	inStack := map[string]bool{}
	var visit func(string) bool
	visit = func(id string) bool {
		if inStack[id] {
			return true
		}
		if seen[id] {
			return false
		}
		seen[id] = true
		inStack[id] = true
		for _, next := range graph[id] {
			if visit(next) {
				return true
			}
		}
		inStack[id] = false
		return false
	}
	for _, n := range nodes {
		if visit(n.ID) {
			return true
		}
	}
	return false
}
