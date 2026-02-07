package theme

import (
	"context"
	"fmt"
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/domain"
	"github.com/example/a2ui-go-agent-platform/pkg/store"
	"github.com/example/a2ui-go-agent-platform/pkg/util"
)

type Service struct {
	themes   store.ThemeRepository
	sessions store.SessionRepository
}

func NewService(themes store.ThemeRepository, sessions store.SessionRepository) *Service {
	return &Service{themes: themes, sessions: sessions}
}

func (s *Service) Create(ctx context.Context, name string, tokenSet map[string]any, isBuiltin bool) (domain.Theme, error) {
	if name == "" {
		return domain.Theme{}, fmt.Errorf("theme name is required")
	}
	t := domain.Theme{
		ID:        util.NewID("theme"),
		Name:      name,
		TokenSet:  tokenSet,
		IsBuiltin: isBuiltin,
		CreatedAt: time.Now(),
	}
	if err := s.themes.CreateTheme(ctx, t); err != nil {
		return domain.Theme{}, err
	}
	return t, nil
}

func (s *Service) BindToSession(ctx context.Context, sessionID, themeID string) error {
	if _, err := s.themes.GetTheme(ctx, themeID); err != nil {
		return err
	}
	ssn, err := s.sessions.GetSession(ctx, sessionID)
	if err != nil {
		return err
	}
	ssn.ActiveThemeID = themeID
	return s.sessions.UpdateSession(ctx, ssn)
}

func (s *Service) List(ctx context.Context) ([]domain.Theme, error) {
	return s.themes.ListThemes(ctx)
}
