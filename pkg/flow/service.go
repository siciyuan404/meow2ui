package flow

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/domain"
	"github.com/example/a2ui-go-agent-platform/pkg/flow/validator"
	"github.com/example/a2ui-go-agent-platform/pkg/store"
	"github.com/example/a2ui-go-agent-platform/pkg/util"
)

type Service struct {
	repo store.FlowRepository
}

func NewService(repo store.FlowRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTemplate(ctx context.Context, name, version string, definition Definition) (domain.FlowTemplate, domain.FlowTemplateVersion, error) {
	if err := validator.Validate(toValidatorDefinition(definition)); err != nil {
		return domain.FlowTemplate{}, domain.FlowTemplateVersion{}, err
	}
	tpl := domain.FlowTemplate{
		ID:        util.NewID("flow"),
		Name:      name,
		Status:    domain.FlowTemplateDraft,
		CreatedAt: time.Now(),
	}
	if err := s.repo.CreateFlowTemplate(ctx, tpl); err != nil {
		return domain.FlowTemplate{}, domain.FlowTemplateVersion{}, err
	}
	buf, err := json.Marshal(definition)
	if err != nil {
		return domain.FlowTemplate{}, domain.FlowTemplateVersion{}, err
	}
	v := domain.FlowTemplateVersion{
		ID:             util.NewID("flowver"),
		TemplateID:     tpl.ID,
		Version:        version,
		DefinitionJSON: string(buf),
		CreatedAt:      time.Now(),
	}
	if err := s.repo.CreateFlowTemplateVersion(ctx, v); err != nil {
		return domain.FlowTemplate{}, domain.FlowTemplateVersion{}, err
	}
	return tpl, v, nil
}

func (s *Service) ListTemplates(ctx context.Context) ([]domain.FlowTemplate, error) {
	return s.repo.ListFlowTemplates(ctx)
}

func (s *Service) BindSession(ctx context.Context, sessionID, templateID, version string) error {
	if _, err := s.repo.GetFlowTemplateVersion(ctx, templateID, version); err != nil {
		return err
	}
	return s.repo.BindSessionFlow(ctx, domain.SessionFlowBinding{
		SessionID:       sessionID,
		TemplateID:      templateID,
		TemplateVersion: version,
		BoundAt:         time.Now(),
	})
}

func (s *Service) ResolveDefinition(ctx context.Context, sessionID string) (Definition, bool, error) {
	b, err := s.repo.GetSessionFlowBinding(ctx, sessionID)
	if err != nil {
		if err == store.ErrNotFound {
			return DefaultDefinition(), false, nil
		}
		return Definition{}, false, err
	}
	v, err := s.repo.GetFlowTemplateVersion(ctx, b.TemplateID, b.TemplateVersion)
	if err != nil {
		return Definition{}, false, err
	}
	var def Definition
	if err := json.Unmarshal([]byte(v.DefinitionJSON), &def); err != nil {
		return Definition{}, false, fmt.Errorf("invalid flow definition: %w", err)
	}
	if err := validator.Validate(toValidatorDefinition(def)); err != nil {
		return Definition{}, false, err
	}
	return def, true, nil
}

func toValidatorDefinition(def Definition) validator.Definition {
	nodes := make([]validator.Node, 0, len(def.Nodes))
	for _, n := range def.Nodes {
		nodes = append(nodes, validator.Node{ID: n.ID, DependsOn: n.DependsOn})
	}
	edges := make([]validator.Edge, 0, len(def.Edges))
	for _, e := range def.Edges {
		edges = append(edges, validator.Edge{From: e.From, To: e.To, Condition: e.Condition})
	}
	return validator.Definition{
		Name:   def.Name,
		Policy: validator.Policy{Parallelism: def.Policy.Parallelism},
		Nodes:  nodes,
		Edges:  edges,
	}
}
