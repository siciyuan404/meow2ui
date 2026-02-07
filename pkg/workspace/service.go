package workspace

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/domain"
	"github.com/example/a2ui-go-agent-platform/pkg/store"
	"github.com/example/a2ui-go-agent-platform/pkg/util"
)

type Service struct {
	repo store.WorkspaceRepository
}

func NewService(repo store.WorkspaceRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, name, rootPath string) (domain.Workspace, error) {
	if name == "" {
		return domain.Workspace{}, fmt.Errorf("workspace name is required")
	}
	if rootPath == "" {
		return domain.Workspace{}, fmt.Errorf("workspace root path is required")
	}
	now := time.Now()
	w := domain.Workspace{
		ID:        util.NewID("ws"),
		Name:      name,
		RootPath:  filepath.Clean(rootPath),
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := s.repo.CreateWorkspace(ctx, w); err != nil {
		return domain.Workspace{}, err
	}
	return w, nil
}

func (s *Service) List(ctx context.Context) ([]domain.Workspace, error) {
	return s.repo.ListWorkspaces(ctx)
}

func (s *Service) Get(ctx context.Context, id string) (domain.Workspace, error) {
	return s.repo.GetWorkspace(ctx, id)
}
