package memorystore

import (
	"context"
	"testing"
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/domain"
)

func TestNew(t *testing.T) {
	ms := New()
	if ms == nil {
		t.Fatal("expected non-nil MemoryStore")
	}
	if ms.Workspace() == nil || ms.Session() == nil || ms.Version() == nil ||
		ms.Provider() == nil || ms.Theme() == nil || ms.Playground() == nil ||
		ms.Event() == nil || ms.Flow() == nil {
		t.Fatal("expected all repository accessors to return non-nil")
	}
}

// ==================== Workspace ====================

func TestWorkspace_CRUD(t *testing.T) {
	ms := New()
	ctx := context.Background()

	err := ms.CreateWorkspace(ctx, domain.Workspace{ID: "ws-1", Name: "Test"})
	if err != nil {
		t.Fatalf("CreateWorkspace: %v", err)
	}

	ws, err := ms.GetWorkspace(ctx, "ws-1")
	if err != nil {
		t.Fatalf("GetWorkspace: %v", err)
	}
	if ws.Name != "Test" {
		t.Fatalf("expected Test, got %s", ws.Name)
	}
	if ws.CreatedAt.IsZero() {
		t.Fatal("expected CreatedAt to be set")
	}

	_, err = ms.GetWorkspace(ctx, "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent workspace")
	}

	list, err := ms.ListWorkspaces(ctx)
	if err != nil {
		t.Fatalf("ListWorkspaces: %v", err)
	}
	if len(list) != 1 {
		t.Fatalf("expected 1 workspace, got %d", len(list))
	}
}

// ==================== Session ====================

func TestSession_CRUD(t *testing.T) {
	ms := New()
	ctx := context.Background()

	err := ms.CreateSession(ctx, domain.Session{ID: "s-1", WorkspaceID: "ws-1", Title: "Session 1"})
	if err != nil {
		t.Fatalf("CreateSession: %v", err)
	}

	s, err := ms.GetSession(ctx, "s-1")
	if err != nil {
		t.Fatalf("GetSession: %v", err)
	}
	if s.Title != "Session 1" {
		t.Fatalf("expected Session 1, got %s", s.Title)
	}

	_, err = ms.GetSession(ctx, "nonexistent")
	if err == nil {
		t.Fatal("expected error")
	}

	s.Title = "Updated"
	err = ms.UpdateSession(ctx, s)
	if err != nil {
		t.Fatalf("UpdateSession: %v", err)
	}
	updated, _ := ms.GetSession(ctx, "s-1")
	if updated.Title != "Updated" {
		t.Fatalf("expected Updated, got %s", updated.Title)
	}

	err = ms.UpdateSession(ctx, domain.Session{ID: "nonexistent"})
	if err == nil {
		t.Fatal("expected error for nonexistent session update")
	}

	list, err := ms.ListByWorkspace(ctx, "ws-1")
	if err != nil {
		t.Fatalf("ListByWorkspace: %v", err)
	}
	if len(list) != 1 {
		t.Fatalf("expected 1, got %d", len(list))
	}

	empty, _ := ms.ListByWorkspace(ctx, "ws-other")
	if len(empty) != 0 {
		t.Fatalf("expected 0, got %d", len(empty))
	}
}

// ==================== Version ====================

func TestVersion_CRUD(t *testing.T) {
	ms := New()
	ctx := context.Background()

	v1 := domain.SchemaVersion{ID: "v-1", SessionID: "s-1", SchemaJSON: "{}", CreatedAt: time.Now().Add(-time.Hour)}
	v2 := domain.SchemaVersion{ID: "v-2", SessionID: "s-1", SchemaJSON: "{\"updated\":true}", CreatedAt: time.Now()}

	ms.CreateVersion(ctx, v1)
	ms.CreateVersion(ctx, v2)

	got, err := ms.GetVersion(ctx, "v-1")
	if err != nil {
		t.Fatalf("GetVersion: %v", err)
	}
	if got.ID != "v-1" {
		t.Fatalf("expected v-1, got %s", got.ID)
	}

	_, err = ms.GetVersion(ctx, "nonexistent")
	if err == nil {
		t.Fatal("expected error")
	}

	latest, err := ms.GetLatestBySession(ctx, "s-1")
	if err != nil {
		t.Fatalf("GetLatestBySession: %v", err)
	}
	if latest.ID != "v-2" {
		t.Fatalf("expected v-2 as latest, got %s", latest.ID)
	}

	_, err = ms.GetLatestBySession(ctx, "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent session")
	}

	list, _ := ms.ListBySession(ctx, "s-1")
	if len(list) != 2 {
		t.Fatalf("expected 2, got %d", len(list))
	}
}

// ==================== Version Assets ====================

func TestVersionAssets(t *testing.T) {
	ms := New()
	ctx := context.Background()

	ms.CreateVersionAsset(ctx, domain.SchemaVersionAsset{ID: "a-1", VersionID: "v-1", AssetType: "image", AssetRef: "img.png"})
	ms.CreateVersionAsset(ctx, domain.SchemaVersionAsset{ID: "a-2", VersionID: "v-1", AssetType: "audio", AssetRef: "audio.mp3"})

	assets, err := ms.ListVersionAssets(ctx, "v-1")
	if err != nil {
		t.Fatalf("ListVersionAssets: %v", err)
	}
	if len(assets) != 2 {
		t.Fatalf("expected 2, got %d", len(assets))
	}

	empty, _ := ms.ListVersionAssets(ctx, "v-other")
	if len(empty) != 0 {
		t.Fatalf("expected 0, got %d", len(empty))
	}
}

// ==================== Provider & Model ====================

func TestProvider_CRUD(t *testing.T) {
	ms := New()
	ctx := context.Background()

	ms.CreateProvider(ctx, domain.Provider{ID: "p-1", Name: "OpenAI", Enabled: true})
	ms.CreateModel(ctx, domain.Model{ID: "m-1", ProviderID: "p-1", Enabled: true})
	ms.CreateModel(ctx, domain.Model{ID: "m-2", ProviderID: "p-1", Enabled: false})

	providers, _ := ms.ListProviders(ctx)
	if len(providers) != 1 {
		t.Fatalf("expected 1 provider, got %d", len(providers))
	}

	models, _ := ms.ListModels(ctx)
	if len(models) != 2 {
		t.Fatalf("expected 2 models, got %d", len(models))
	}

	enabled, _ := ms.ListEnabledModels(ctx)
	if len(enabled) != 1 {
		t.Fatalf("expected 1 enabled model, got %d", len(enabled))
	}
}

// ==================== Theme ====================

func TestTheme_CRUD(t *testing.T) {
	ms := New()
	ctx := context.Background()

	ms.CreateTheme(ctx, domain.Theme{ID: "t-1", Name: "Dark"})

	theme, err := ms.GetTheme(ctx, "t-1")
	if err != nil {
		t.Fatalf("GetTheme: %v", err)
	}
	if theme.Name != "Dark" {
		t.Fatalf("expected Dark, got %s", theme.Name)
	}

	_, err = ms.GetTheme(ctx, "nonexistent")
	if err == nil {
		t.Fatal("expected error")
	}

	list, _ := ms.ListThemes(ctx)
	if len(list) != 1 {
		t.Fatalf("expected 1, got %d", len(list))
	}
}

// ==================== Playground ====================

func TestPlayground_CRUD(t *testing.T) {
	ms := New()
	ctx := context.Background()

	ms.CreateCategory(ctx, domain.PlaygroundCategory{ID: "cat-1", Name: "Dashboard"})
	ms.CreateTag(ctx, domain.PlaygroundTag{ID: "tag-1", Name: "modern"})
	ms.CreateItem(ctx, domain.PlaygroundItem{
		ID: "item-1", Title: "Modern Dashboard", CategoryID: "cat-1", Tags: []string{"modern"},
	})
	ms.CreateItem(ctx, domain.PlaygroundItem{
		ID: "item-2", Title: "Classic Form", CategoryID: "cat-2", Tags: []string{"classic"},
	})

	// Search by category
	items, _ := ms.SearchItems(ctx, "cat-1", nil, "")
	if len(items) != 1 || items[0].ID != "item-1" {
		t.Fatalf("expected 1 item in cat-1, got %d", len(items))
	}

	// Search by query
	items, _ = ms.SearchItems(ctx, "", nil, "modern")
	if len(items) != 1 {
		t.Fatalf("expected 1 item matching 'modern', got %d", len(items))
	}

	// Search by tag
	items, _ = ms.SearchItems(ctx, "", []string{"classic"}, "")
	if len(items) != 1 || items[0].ID != "item-2" {
		t.Fatalf("expected 1 item with tag 'classic', got %d", len(items))
	}

	// Search all
	items, _ = ms.SearchItems(ctx, "", nil, "")
	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}

	// Search no match
	items, _ = ms.SearchItems(ctx, "", []string{"nonexistent"}, "")
	if len(items) != 0 {
		t.Fatalf("expected 0, got %d", len(items))
	}
}

// ==================== Agent Runs & Events ====================

func TestAgentRun_CRUD(t *testing.T) {
	ms := New()
	ctx := context.Background()

	ms.CreateRun(ctx, domain.AgentRun{ID: "run-1", SessionID: "s-1", Status: domain.AgentRunRunning})

	run, err := ms.GetRun(ctx, "run-1")
	if err != nil {
		t.Fatalf("GetRun: %v", err)
	}
	if run.Status != domain.AgentRunRunning {
		t.Fatalf("expected running, got %s", run.Status)
	}

	_, err = ms.GetRun(ctx, "nonexistent")
	if err == nil {
		t.Fatal("expected error")
	}

	run.Status = domain.AgentRunCompleted
	ms.UpdateRun(ctx, run)
	updated, _ := ms.GetRun(ctx, "run-1")
	if updated.Status != domain.AgentRunCompleted {
		t.Fatalf("expected completed, got %s", updated.Status)
	}

	err = ms.UpdateRun(ctx, domain.AgentRun{ID: "nonexistent"})
	if err == nil {
		t.Fatal("expected error for nonexistent run update")
	}

	runs, _ := ms.ListRuns(ctx)
	if len(runs) != 1 {
		t.Fatalf("expected 1, got %d", len(runs))
	}
}

func TestAgentEvent_CRUD(t *testing.T) {
	ms := New()
	ctx := context.Background()

	ms.CreateEvent(ctx, domain.AgentEvent{ID: "e-1", RunID: "run-1", Step: "plan"})
	ms.CreateEvent(ctx, domain.AgentEvent{ID: "e-2", RunID: "run-1", Step: "emit"})

	events, _ := ms.ListEventsByRun(ctx, "run-1")
	if len(events) != 2 {
		t.Fatalf("expected 2, got %d", len(events))
	}

	empty, _ := ms.ListEventsByRun(ctx, "run-other")
	if len(empty) != 0 {
		t.Fatalf("expected 0, got %d", len(empty))
	}
}

// ==================== Flow ====================

func TestFlow_CRUD(t *testing.T) {
	ms := New()
	ctx := context.Background()

	ms.CreateFlowTemplate(ctx, domain.FlowTemplate{ID: "ft-1", Name: "Default Flow"})

	templates, _ := ms.ListFlowTemplates(ctx)
	if len(templates) != 1 {
		t.Fatalf("expected 1, got %d", len(templates))
	}

	ms.CreateFlowTemplateVersion(ctx, domain.FlowTemplateVersion{
		ID: "fv-1", TemplateID: "ft-1", Version: "1.0", DefinitionJSON: "{}",
	})

	ver, err := ms.GetFlowTemplateVersion(ctx, "ft-1", "1.0")
	if err != nil {
		t.Fatalf("GetFlowTemplateVersion: %v", err)
	}
	if ver.Version != "1.0" {
		t.Fatalf("expected 1.0, got %s", ver.Version)
	}

	_, err = ms.GetFlowTemplateVersion(ctx, "ft-1", "2.0")
	if err == nil {
		t.Fatal("expected error for nonexistent version")
	}

	ms.BindSessionFlow(ctx, domain.SessionFlowBinding{
		SessionID: "s-1", TemplateID: "ft-1", TemplateVersion: "1.0",
	})

	binding, err := ms.GetSessionFlowBinding(ctx, "s-1")
	if err != nil {
		t.Fatalf("GetSessionFlowBinding: %v", err)
	}
	if binding.TemplateID != "ft-1" {
		t.Fatalf("expected ft-1, got %s", binding.TemplateID)
	}

	_, err = ms.GetSessionFlowBinding(ctx, "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent binding")
	}
}
