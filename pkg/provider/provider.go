package provider

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/domain"
	"github.com/example/a2ui-go-agent-platform/pkg/store"
)

type TaskType string

const (
	TaskPlan   TaskType = "plan"
	TaskEmit   TaskType = "emit"
	TaskRepair TaskType = "repair"
)

type GenerateRequest struct {
	SystemPrompt string
	UserPrompt   string
	Context      map[string]any
}

type GenerateResponse struct {
	Text   string
	Tokens int
}

type Adapter interface {
	Generate(context.Context, domain.Provider, domain.Model, GenerateRequest) (GenerateResponse, error)
	Name() string
}

type Registry struct {
	adapters map[string]Adapter
}

func NewRegistry() *Registry {
	return &Registry{adapters: map[string]Adapter{}}
}

func (r *Registry) Register(providerType string, adapter Adapter) {
	r.adapters[strings.ToLower(providerType)] = adapter
}

func (r *Registry) Get(providerType string) (Adapter, bool) {
	a, ok := r.adapters[strings.ToLower(providerType)]
	return a, ok
}

type Router struct {
	repo store.ProviderRepository
}

func NewRouter(repo store.ProviderRepository) *Router {
	return &Router{repo: repo}
}

func (r *Router) Route(ctx context.Context, task TaskType) (domain.Provider, domain.Model, error) {
	candidates, err := r.Candidates(ctx, task)
	if err != nil {
		return domain.Provider{}, domain.Model{}, err
	}
	if len(candidates) == 0 {
		return domain.Provider{}, domain.Model{}, fmt.Errorf("no enabled model for task %s", task)
	}
	return candidates[0].Provider, candidates[0].Model, nil
}

type Candidate struct {
	Provider domain.Provider
	Model    domain.Model
}

func (r *Router) Candidates(ctx context.Context, task TaskType) ([]Candidate, error) {
	providers, err := r.repo.ListProviders(ctx)
	if err != nil {
		return nil, err
	}
	models, err := r.repo.ListEnabledModels(ctx)
	if err != nil {
		return nil, err
	}
	required := capabilityForTask(task)
	role := strings.ToLower(string(task))
	candidates := []Candidate{}
	providerMap := map[string]domain.Provider{}
	for _, p := range providers {
		providerMap[p.ID] = p
	}
	for _, m := range models {
		if !hasCapability(m.Capabilities, required) {
			continue
		}
		if !modelMatchesRole(m, role) {
			continue
		}
		p, ok := providerMap[m.ProviderID]
		if !ok || !p.Enabled {
			continue
		}
		candidates = append(candidates, Candidate{Provider: p, Model: m})
	}
	sort.Slice(candidates, func(i, j int) bool {
		return modelPriority(candidates[i].Model) < modelPriority(candidates[j].Model)
	})
	return candidates, nil
}

func capabilityForTask(task TaskType) string {
	switch task {
	case TaskPlan, TaskEmit, TaskRepair:
		return "text"
	default:
		return "text"
	}
}

func hasCapability(have []string, required string) bool {
	for _, c := range have {
		if strings.EqualFold(c, required) {
			return true
		}
	}
	return false
}

type Service struct {
	router   *Router
	registry *Registry
}

func NewService(router *Router, registry *Registry) *Service {
	return &Service{router: router, registry: registry}
}

func (s *Service) Generate(ctx context.Context, task TaskType, req GenerateRequest) (GenerateResponse, error) {
	candidates, err := s.router.Candidates(ctx, task)
	if err != nil {
		return GenerateResponse{}, err
	}
	if len(candidates) == 0 {
		return GenerateResponse{}, fmt.Errorf("no enabled model for task %s", task)
	}

	var lastErr error
	for _, c := range candidates {
		adapter, ok := s.registry.Get(c.Provider.Type)
		if !ok {
			lastErr = fmt.Errorf("adapter not found for provider type: %s", c.Provider.Type)
			continue
		}
		for attempt := 0; attempt < 3; attempt++ {
			resp, err := adapter.Generate(ctx, c.Provider, c.Model, req)
			if err == nil {
				return resp, nil
			}
			lastErr = err
			pErr, ok := err.(*ProviderError)
			if !ok || !pErr.Retryable || attempt == 2 {
				break
			}
			time.Sleep(time.Duration(150*(attempt+1)) * time.Millisecond)
		}
	}
	if lastErr == nil {
		lastErr = errors.New("provider call failed")
	}
	return GenerateResponse{}, lastErr
}

func modelMatchesRole(model domain.Model, role string) bool {
	if model.Metadata == nil {
		return true
	}
	v, ok := model.Metadata["role"]
	if !ok {
		return true
	}
	r, ok := v.(string)
	if !ok {
		return true
	}
	return strings.EqualFold(strings.TrimSpace(r), role)
}

func modelPriority(model domain.Model) int {
	if model.Metadata == nil {
		return 100
	}
	v, ok := model.Metadata["priority"]
	if !ok {
		return 100
	}
	switch t := v.(type) {
	case int:
		return t
	case int32:
		return int(t)
	case int64:
		return int(t)
	case float64:
		return int(t)
	case string:
		n, err := strconv.Atoi(t)
		if err == nil {
			return n
		}
	}
	return 100
}
