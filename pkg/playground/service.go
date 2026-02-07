package playground

import (
	"context"
	"fmt"
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/domain"
	"github.com/example/a2ui-go-agent-platform/pkg/store"
	"github.com/example/a2ui-go-agent-platform/pkg/util"
)

type Service struct {
	repo store.PlaygroundRepository
}

func NewService(repo store.PlaygroundRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateCategory(ctx context.Context, name string) (domain.PlaygroundCategory, error) {
	if name == "" {
		return domain.PlaygroundCategory{}, fmt.Errorf("category name is required")
	}
	c := domain.PlaygroundCategory{ID: util.NewID("pgcat"), Name: name}
	if err := s.repo.CreateCategory(ctx, c); err != nil {
		return domain.PlaygroundCategory{}, err
	}
	return c, nil
}

func (s *Service) CreateTag(ctx context.Context, name string) (domain.PlaygroundTag, error) {
	if name == "" {
		return domain.PlaygroundTag{}, fmt.Errorf("tag name is required")
	}
	t := domain.PlaygroundTag{ID: util.NewID("pgtag"), Name: name}
	if err := s.repo.CreateTag(ctx, t); err != nil {
		return domain.PlaygroundTag{}, err
	}
	return t, nil
}

type SaveInput struct {
	Title           string
	CategoryID      string
	SourceSessionID string
	SourceVersionID string
	ThemeID         string
	SchemaSnapshot  string
	PreviewRef      string
	Tags            []string
}

func (s *Service) SaveItem(ctx context.Context, in SaveInput) (domain.PlaygroundItem, error) {
	if in.Title == "" {
		return domain.PlaygroundItem{}, fmt.Errorf("title is required")
	}
	if in.SchemaSnapshot == "" {
		return domain.PlaygroundItem{}, fmt.Errorf("schema snapshot is required")
	}
	item := domain.PlaygroundItem{
		ID:              util.NewID("pg"),
		Title:           in.Title,
		CategoryID:      in.CategoryID,
		SourceSessionID: in.SourceSessionID,
		SourceVersionID: in.SourceVersionID,
		ThemeID:         in.ThemeID,
		SchemaSnapshot:  in.SchemaSnapshot,
		PreviewRef:      in.PreviewRef,
		Tags:            append([]string(nil), in.Tags...),
		CreatedAt:       time.Now(),
	}
	if err := s.repo.CreateItem(ctx, item); err != nil {
		return domain.PlaygroundItem{}, err
	}
	return item, nil
}

func (s *Service) Search(ctx context.Context, categoryID string, tags []string, query string) ([]domain.PlaygroundItem, error) {
	return s.repo.SearchItems(ctx, categoryID, tags, query)
}
