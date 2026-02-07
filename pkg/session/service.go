package session

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/domain"
	"github.com/example/a2ui-go-agent-platform/pkg/store"
	"github.com/example/a2ui-go-agent-platform/pkg/util"
)

type Service struct {
	sessions store.SessionRepository
	versions store.VersionRepository
}

func NewService(sessions store.SessionRepository, versions store.VersionRepository) *Service {
	return &Service{sessions: sessions, versions: versions}
}

func (s *Service) Create(ctx context.Context, workspaceID, title, themeID string) (domain.Session, domain.SchemaVersion, error) {
	if workspaceID == "" {
		return domain.Session{}, domain.SchemaVersion{}, fmt.Errorf("workspace id is required")
	}
	if themeID == "" {
		themeID = "default"
	}
	now := time.Now()
	session := domain.Session{
		ID:            util.NewID("ssn"),
		WorkspaceID:   workspaceID,
		Title:         title,
		ActiveThemeID: themeID,
		Status:        domain.SessionActive,
		CreatedAt:     now,
		UpdatedAt:     now,
	}
	if err := s.sessions.CreateSession(ctx, session); err != nil {
		return domain.Session{}, domain.SchemaVersion{}, err
	}

	initial := domain.UISchema{
		Version: "1.0.0",
		Root: domain.UIComponent{
			ID:   "root",
			Type: "Container",
			Props: map[string]any{
				"className": "min-h-screen",
			},
		},
	}
	buf, err := json.Marshal(initial)
	if err != nil {
		return domain.Session{}, domain.SchemaVersion{}, err
	}
	ver := domain.SchemaVersion{
		ID:              util.NewID("ver"),
		SessionID:       session.ID,
		ParentVersionID: "",
		SchemaPath:      filepath.Join("ui-prototype", session.ID+".json"),
		SchemaHash:      "",
		SchemaJSON:      string(buf),
		Summary:         "initial schema",
		ThemeSnapshotID: themeID,
		CreatedAt:       now,
	}
	if err := s.versions.CreateVersion(ctx, ver); err != nil {
		return domain.Session{}, domain.SchemaVersion{}, err
	}
	return session, ver, nil
}

func (s *Service) GetLatestVersion(ctx context.Context, sessionID string) (domain.SchemaVersion, error) {
	return s.versions.GetLatestBySession(ctx, sessionID)
}

func (s *Service) ListVersions(ctx context.Context, sessionID string) ([]domain.SchemaVersion, error) {
	return s.versions.ListBySession(ctx, sessionID)
}

func (s *Service) SetActiveTheme(ctx context.Context, sessionID, themeID string) error {
	ssn, err := s.sessions.GetSession(ctx, sessionID)
	if err != nil {
		return err
	}
	ssn.ActiveThemeID = themeID
	return s.sessions.UpdateSession(ctx, ssn)
}

type ContextBundle struct {
	SystemRules  map[string]any
	SessionFacts map[string]any
	TaskInput    map[string]any
}

func (s *Service) BuildContext(ctx context.Context, sessionID string, task map[string]any) (ContextBundle, error) {
	ssn, err := s.sessions.GetSession(ctx, sessionID)
	if err != nil {
		return ContextBundle{}, err
	}
	latest, err := s.versions.GetLatestBySession(ctx, sessionID)
	if err != nil {
		return ContextBundle{}, err
	}
	return ContextBundle{
		SystemRules: map[string]any{
			"output_mode":           "patch_preferred",
			"strict_json_schema":    true,
			"component_whitelist":   true,
			"local_edit_constraint": true,
		},
		SessionFacts: map[string]any{
			"session_id":      ssn.ID,
			"active_theme_id": ssn.ActiveThemeID,
			"metadata":        ssn.Metadata,
			"current_version": latest.ID,
			"schema":          latest.SchemaJSON,
		},
		TaskInput: task,
	}, nil
}
