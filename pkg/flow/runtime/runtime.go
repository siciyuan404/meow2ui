package runtime

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/flow"
	"github.com/example/a2ui-go-agent-platform/pkg/provider"
)

type Executor interface {
	Execute(context.Context, flow.Definition, provider.GenerateRequest) (flow.RunResult, error)
}

type Service struct {
	provider *provider.Service
}

func NewService(providerSvc *provider.Service) *Service {
	return &Service{provider: providerSvc}
}

func (s *Service) Execute(ctx context.Context, def flow.Definition, req provider.GenerateRequest) (flow.RunResult, error) {
	nodes := map[string]flow.Node{}
	pending := map[string]int{}
	next := map[string][]flow.Edge{}
	for _, n := range def.Nodes {
		nodes[n.ID] = n
		pending[n.ID] = len(n.DependsOn)
	}
	for _, e := range def.Edges {
		next[e.From] = append(next[e.From], e)
	}

	queue := make([]string, 0, len(def.Nodes))
	for _, n := range def.Nodes {
		if pending[n.ID] == 0 {
			queue = append(queue, n.ID)
		}
	}

	var mu sync.Mutex
	results := make([]flow.NodeResult, 0, len(def.Nodes))
	executed := map[string]flow.NodeResult{}

	workerN := def.Policy.Parallelism
	if workerN <= 0 {
		workerN = 1
	}

	for len(queue) > 0 {
		batch := append([]string(nil), queue...)
		queue = queue[:0]
		var wg sync.WaitGroup
		errCh := make(chan error, len(batch))
		sem := make(chan struct{}, workerN)
		for _, nodeID := range batch {
			sem <- struct{}{}
			wg.Add(1)
			go func(id string) {
				defer wg.Done()
				defer func() { <-sem }()
				res, err := s.executeNode(ctx, nodes[id], req)
				mu.Lock()
				results = append(results, res)
				executed[id] = res
				mu.Unlock()
				if err != nil {
					errCh <- err
				}
			}(nodeID)
		}
		wg.Wait()
		close(errCh)

		if len(errCh) > 0 && def.Policy.FailureMode == flow.FailureModeFailFast {
			return flow.RunResult{Success: false, Nodes: results}, <-errCh
		}

		for _, id := range batch {
			for _, edge := range next[id] {
				if !allowEdge(edge.Condition, executed[id]) {
					continue
				}
				pending[edge.To]--
				if pending[edge.To] == 0 {
					queue = append(queue, edge.To)
				}
			}
		}
	}

	return flow.RunResult{Success: true, Nodes: results}, nil
}

func (s *Service) executeNode(ctx context.Context, node flow.Node, req provider.GenerateRequest) (flow.NodeResult, error) {
	start := time.Now()
	result := flow.NodeResult{NodeID: node.ID, Step: string(node.Type), Status: "completed"}

	taskType := provider.TaskEmit
	switch node.Type {
	case flow.NodeTypePlan:
		taskType = resolvePlanTask(req)
	case flow.NodeTypeEmit:
		taskType = provider.TaskEmit
	case flow.NodeTypeRepair:
		taskType = provider.TaskRepair
	case flow.NodeTypeValidate, flow.NodeTypeApply, flow.NodeTypeCustom:
		result.LatencyMS = time.Since(start).Milliseconds()
		result.Output = map[string]any{"ok": true}
		return result, nil
	}

	resp, err := s.provider.Generate(ctx, taskType, req)
	result.LatencyMS = time.Since(start).Milliseconds()
	if err != nil {
		result.Status = "failed"
		result.Error = err.Error()
		return result, err
	}
	result.Tokens = resp.Tokens
	result.Generated = resp
	result.Output = map[string]any{"text": resp.Text}
	return result, nil
}

func resolvePlanTask(req provider.GenerateRequest) provider.TaskType {
	v, ok := req.Context["plan_task"]
	if !ok {
		return provider.TaskPlan
	}
	s, ok := v.(string)
	if !ok {
		return provider.TaskPlan
	}
	switch provider.TaskType(s) {
	case provider.TaskPlanImage:
		return provider.TaskPlanImage
	case provider.TaskPlanAudio:
		return provider.TaskPlanAudio
	default:
		return provider.TaskPlan
	}
}

func allowEdge(condition string, from flow.NodeResult) bool {
	switch condition {
	case "", "always":
		return true
	case "validate.ok":
		ok, _ := from.Output["ok"].(bool)
		return ok
	case "!validate.ok":
		ok, _ := from.Output["ok"].(bool)
		return !ok
	case "success":
		return from.Status == "completed"
	case "failed":
		return from.Status == "failed"
	default:
		return false
	}
}

func RequireSuccess(result flow.RunResult) error {
	if result.Success {
		return nil
	}
	for _, item := range result.Nodes {
		if item.Status == "failed" {
			return fmt.Errorf("node %s failed: %s", item.NodeID, item.Error)
		}
	}
	return fmt.Errorf("flow failed")
}
