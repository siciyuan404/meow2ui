package session

import (
	"context"
	"fmt"
	"testing"

	"github.com/example/a2ui-go-agent-platform/pkg/domain"
	"github.com/example/a2ui-go-agent-platform/pkg/store"
)

// --- Mock Repositories ---

type mockSessionRepo struct {
	sessions map[string]domain.Session
	err      error
}

func newMockSessionRepo() *mockSessionRepo {
	return &mockSessionRepo{sessions: map[string]domain.Session{}}
}

func (m *mockSessionRepo) CreateSession(_ context.Context, s domain.Session) error {
	if m.err != nil {
		return m.err
	}
	m.sessions[s.ID] = s
	return nil
}
func (m *mockSessionRepo) GetSession(_ context.Context, id string) (domain.Session, error) {
	if m.err != nil {
		return domain.Session{}, m.err
	}
	s, ok := m.sessions[id]
	if !ok {
		return domain.Session{}, store.ErrNotFound
	}
	return s, nil
}
func (m *mockSessionRepo) ListByWorkspace(_ context.Context, _ string) ([]domain.Session, error) {
	return nil, nil
}
func (m *mockSessionRepo) UpdateSession(_ context.Context, s domain.Session) error {
	if m.err != nil {
		return m.err
	}
	m.sessions[s.ID] = s
	return nil
}

type mockVersionRepo struct {
	versions []domain.SchemaVersion
	err      error
}

func newMockVersionRepo() *mockVersionRepo {
	return &mockVersionRepo{versions: []domain.SchemaVersion{}}
}

func (m *mockVersionRepo) CreateVersion(_ context.Context, v domain.SchemaVersion) error {
	if m.err != nil {
		return m.err
	}
	m.versions = append(m.versions, v)
	return nil
}
func (m *mockVersionRepo) GetVersion(_ context.Context, id string) (domain.SchemaVersion, error) {
	for _, v := range m.versions {
		if v.ID == id {
			return v, nil
		}
	}
	return domain.SchemaVersion{}, store.ErrNotFound
}
func (m *mockVersionRepo) GetLatestBySession(_ context.Context, sessionID string) (domain.SchemaVersion, error) {
	if m.err != nil {
		return domain.SchemaVersion{}, m.err
	}
	var latest domain.SchemaVersion
	found := false
	for _, v := range m.versions {
		if v.SessionID == sessionID {
			latest = v
			found = true
		}
	}
	if !found {
		return domain.SchemaVersion{}, store.ErrNotFound
	}
	return latest, nil
}
func (m *mockVersionRepo) ListBySession(_ context.Context, sessionID string) ([]domain.SchemaVersion, error) {
	var result []domain.SchemaVersion
	for _, v := range m.versions {
		if v.SessionID == sessionID {
			result = append(result, v)
		}
	}
	return result, nil
}
func (m *mockVersionRepo) CreateVersionAsset(_ context.Context, _ domain.SchemaVersionAsset) error {
	return nil
}
func (m *mockVersionRepo) ListVersionAssets(_ context.Context, _ string) ([]domain.SchemaVersionAsset, error) {
	return nil, nil
}

// ==================== Create Tests ====================

func TestService_Create_Success(t *testing.T) {
	svc := NewService(newMockSessionRepo(), newMockVersionRepo())
	session, version, err := svc.Create(context.Background(), "ws-1", "Test Session", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if session.WorkspaceID != "ws-1" {
		t.Fatalf("expected workspace ws-1, got %s", session.WorkspaceID)
	}
	if session.ActiveThemeID != "default" {
		t.Fatalf("expected default theme, got %s", session.ActiveThemeID)
	}
	if session.Status != domain.SessionActive {
		t.Fatalf("expected active status, got %s", session.Status)
	}
	if version.SessionID != session.ID {
		t.Fatalf("version session mismatch")
	}
	if version.Summary != "initial schema" {
		t.Fatalf("expected initial schema summary, got %s", version.Summary)
	}
}

func TestService_Create_CustomTheme(t *testing.T) {
	svc := NewService(newMockSessionRepo(), newMockVersionRepo())
	session, _, err := svc.Create(context.Background(), "ws-1", "Test", "dark")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if session.ActiveThemeID != "dark" {
		t.Fatalf("expected dark theme, got %s", session.ActiveThemeID)
	}
}

func TestService_Create_EmptyWorkspaceID(t *testing.T) {
	svc := NewService(newMockSessionRepo(), newMockVersionRepo())
	_, _, err := svc.Create(context.Background(), "", "Test", "")
	if err == nil {
		t.Fatal("expected error for empty workspace ID")
	}
}

func TestService_Create_SessionRepoError(t *testing.T) {
	sr := newMockSessionRepo()
	sr.err = fmt.Errorf("db error")
	svc := NewService(sr, newMockVersionRepo())
	_, _, err := svc.Create(context.Background(), "ws-1", "Test", "")
	if err == nil {
		t.Fatal("expected error from session repo")
	}
}

func TestService_Create_VersionRepoError(t *testing.T) {
	vr := newMockVersionRepo()
	vr.err = fmt.Errorf("db error")
	svc := NewService(newMockSessionRepo(), vr)
	_, _, err := svc.Create(context.Background(), "ws-1", "Test", "")
	if err == nil {
		t.Fatal("expected error from version repo")
	}
}

// ==================== SetActiveTheme Tests ====================

func TestService_SetActiveTheme_Success(t *testing.T) {
	sr := newMockSessionRepo()
	svc := NewService(sr, newMockVersionRepo())
	session, _, _ := svc.Create(context.Background(), "ws-1", "Test", "light")

	err := svc.SetActiveTheme(context.Background(), session.ID, "dark")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	updated := sr.sessions[session.ID]
	if updated.ActiveThemeID != "dark" {
		t.Fatalf("expected dark, got %s", updated.ActiveThemeID)
	}
}

func TestService_SetActiveTheme_NotFound(t *testing.T) {
	svc := NewService(newMockSessionRepo(), newMockVersionRepo())
	err := svc.SetActiveTheme(context.Background(), "nonexistent", "dark")
	if err == nil {
		t.Fatal("expected error for nonexistent session")
	}
}

// ==================== BuildContext Tests ====================

func TestService_BuildContext_Success(t *testing.T) {
	sr := newMockSessionRepo()
	vr := newMockVersionRepo()
	svc := NewService(sr, vr)
	session, _, _ := svc.Create(context.Background(), "ws-1", "Test", "default")

	bundle, err := svc.BuildContext(context.Background(), session.ID, map[string]any{"prompt": "hello"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if bundle.SystemRules["strict_json_schema"] != true {
		t.Fatal("expected strict_json_schema in system rules")
	}
	if bundle.SessionFacts["session_id"] != session.ID {
		t.Fatal("expected session_id in session facts")
	}
	if bundle.TaskInput["prompt"] != "hello" {
		t.Fatal("expected prompt in task input")
	}
}

func TestService_BuildContext_WithMedia(t *testing.T) {
	sr := newMockSessionRepo()
	vr := newMockVersionRepo()
	svc := NewService(sr, vr)
	session, _, _ := svc.Create(context.Background(), "ws-1", "Test", "default")

	media := []domain.MultimodalInput{
		{Type: domain.MediaTypeImage, Ref: "https://example.com/img.png"},
	}
	task := map[string]any{"prompt": "describe", "media": media}
	bundle, err := svc.BuildContext(context.Background(), session.ID, task)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	mediaResult, ok := bundle.TaskInput["media"].([]map[string]any)
	if !ok {
		t.Fatal("expected media to be normalized to []map[string]any")
	}
	if len(mediaResult) != 1 {
		t.Fatalf("expected 1 media item, got %d", len(mediaResult))
	}
}

func TestService_BuildContext_SessionNotFound(t *testing.T) {
	svc := NewService(newMockSessionRepo(), newMockVersionRepo())
	_, err := svc.BuildContext(context.Background(), "nonexistent", map[string]any{})
	if err == nil {
		t.Fatal("expected error for nonexistent session")
	}
}

// ==================== GetLatestVersion / ListVersions ====================

func TestService_GetLatestVersion(t *testing.T) {
	sr := newMockSessionRepo()
	vr := newMockVersionRepo()
	svc := NewService(sr, vr)
	session, ver, _ := svc.Create(context.Background(), "ws-1", "Test", "")

	latest, err := svc.GetLatestVersion(context.Background(), session.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if latest.ID != ver.ID {
		t.Fatalf("expected version %s, got %s", ver.ID, latest.ID)
	}
}

func TestService_ListVersions(t *testing.T) {
	sr := newMockSessionRepo()
	vr := newMockVersionRepo()
	svc := NewService(sr, vr)
	session, _, _ := svc.Create(context.Background(), "ws-1", "Test", "")

	versions, err := svc.ListVersions(context.Background(), session.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(versions) != 1 {
		t.Fatalf("expected 1 version, got %d", len(versions))
	}
}
