package memorystore

import (
	"context"
	"errors"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/domain"
	"github.com/example/a2ui-go-agent-platform/pkg/store"
)

var ErrNotFound = errors.New("not found")

type MemoryStore struct {
	mu sync.RWMutex

	workspaces map[string]domain.Workspace
	sessions   map[string]domain.Session
	versions   map[string]domain.SchemaVersion

	providers map[string]domain.Provider
	models    map[string]domain.Model

	themes map[string]domain.Theme

	categories map[string]domain.PlaygroundCategory
	tags       map[string]domain.PlaygroundTag
	items      map[string]domain.PlaygroundItem

	runs   map[string]domain.AgentRun
	events map[string][]domain.AgentEvent
}

func New() *MemoryStore {
	return &MemoryStore{
		workspaces: map[string]domain.Workspace{},
		sessions:   map[string]domain.Session{},
		versions:   map[string]domain.SchemaVersion{},
		providers:  map[string]domain.Provider{},
		models:     map[string]domain.Model{},
		themes:     map[string]domain.Theme{},
		categories: map[string]domain.PlaygroundCategory{},
		tags:       map[string]domain.PlaygroundTag{},
		items:      map[string]domain.PlaygroundItem{},
		runs:       map[string]domain.AgentRun{},
		events:     map[string][]domain.AgentEvent{},
	}
}

func (m *MemoryStore) Workspace() store.WorkspaceRepository   { return m }
func (m *MemoryStore) Session() store.SessionRepository       { return m }
func (m *MemoryStore) Version() store.VersionRepository       { return m }
func (m *MemoryStore) Provider() store.ProviderRepository     { return m }
func (m *MemoryStore) Theme() store.ThemeRepository           { return m }
func (m *MemoryStore) Playground() store.PlaygroundRepository { return m }
func (m *MemoryStore) Event() store.EventRepository           { return m }

func (m *MemoryStore) CreateWorkspace(_ context.Context, w domain.Workspace) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if w.CreatedAt.IsZero() {
		now := time.Now()
		w.CreatedAt = now
		w.UpdatedAt = now
	}
	m.workspaces[w.ID] = w
	return nil
}

func (m *MemoryStore) GetWorkspace(_ context.Context, id string) (domain.Workspace, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	w, ok := m.workspaces[id]
	if !ok {
		return domain.Workspace{}, ErrNotFound
	}
	return w, nil
}

func (m *MemoryStore) ListWorkspaces(_ context.Context) ([]domain.Workspace, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]domain.Workspace, 0, len(m.workspaces))
	for _, v := range m.workspaces {
		out = append(out, v)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].CreatedAt.Before(out[j].CreatedAt) })
	return out, nil
}

func (m *MemoryStore) CreateSession(_ context.Context, s domain.Session) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if s.CreatedAt.IsZero() {
		now := time.Now()
		s.CreatedAt = now
		s.UpdatedAt = now
	}
	m.sessions[s.ID] = s
	return nil
}

func (m *MemoryStore) GetSession(_ context.Context, id string) (domain.Session, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	s, ok := m.sessions[id]
	if !ok {
		return domain.Session{}, ErrNotFound
	}
	return s, nil
}

func (m *MemoryStore) ListByWorkspace(_ context.Context, workspaceID string) ([]domain.Session, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := []domain.Session{}
	for _, s := range m.sessions {
		if s.WorkspaceID == workspaceID {
			out = append(out, s)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].CreatedAt.Before(out[j].CreatedAt) })
	return out, nil
}

func (m *MemoryStore) UpdateSession(_ context.Context, s domain.Session) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.sessions[s.ID]; !ok {
		return ErrNotFound
	}
	s.UpdatedAt = time.Now()
	m.sessions[s.ID] = s
	return nil
}

func (m *MemoryStore) updateSessionMetadata(_ context.Context, sessionID string, metadata map[string]any) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	ssn, ok := m.sessions[sessionID]
	if !ok {
		return ErrNotFound
	}
	ssn.Metadata = metadata
	ssn.UpdatedAt = time.Now()
	m.sessions[sessionID] = ssn
	return nil
}

func (m *MemoryStore) CreateVersion(_ context.Context, v domain.SchemaVersion) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if v.CreatedAt.IsZero() {
		v.CreatedAt = time.Now()
	}
	m.versions[v.ID] = v
	return nil
}

func (m *MemoryStore) GetVersion(_ context.Context, id string) (domain.SchemaVersion, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.versions[id]
	if !ok {
		return domain.SchemaVersion{}, ErrNotFound
	}
	return v, nil
}

func (m *MemoryStore) GetLatestBySession(_ context.Context, sessionID string) (domain.SchemaVersion, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var latest *domain.SchemaVersion
	for _, v := range m.versions {
		if v.SessionID != sessionID {
			continue
		}
		if latest == nil || v.CreatedAt.After(latest.CreatedAt) {
			t := v
			latest = &t
		}
	}
	if latest == nil {
		return domain.SchemaVersion{}, ErrNotFound
	}
	return *latest, nil
}

func (m *MemoryStore) ListBySession(_ context.Context, sessionID string) ([]domain.SchemaVersion, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := []domain.SchemaVersion{}
	for _, v := range m.versions {
		if v.SessionID == sessionID {
			out = append(out, v)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].CreatedAt.Before(out[j].CreatedAt) })
	return out, nil
}

func (m *MemoryStore) CreateProvider(_ context.Context, p domain.Provider) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.providers[p.ID] = p
	return nil
}

func (m *MemoryStore) CreateModel(_ context.Context, model domain.Model) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.models[model.ID] = model
	return nil
}

func (m *MemoryStore) ListProviders(_ context.Context) ([]domain.Provider, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]domain.Provider, 0, len(m.providers))
	for _, p := range m.providers {
		out = append(out, p)
	}
	return out, nil
}

func (m *MemoryStore) ListModels(_ context.Context) ([]domain.Model, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := make([]domain.Model, 0, len(m.models))
	for _, model := range m.models {
		out = append(out, model)
	}
	return out, nil
}

func (m *MemoryStore) ListEnabledModels(_ context.Context) ([]domain.Model, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := []domain.Model{}
	for _, model := range m.models {
		if model.Enabled {
			out = append(out, model)
		}
	}
	return out, nil
}

func (m *MemoryStore) CreateTheme(_ context.Context, t domain.Theme) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if t.CreatedAt.IsZero() {
		t.CreatedAt = time.Now()
	}
	m.themes[t.ID] = t
	return nil
}

func (m *MemoryStore) GetTheme(_ context.Context, id string) (domain.Theme, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	t, ok := m.themes[id]
	if !ok {
		return domain.Theme{}, ErrNotFound
	}
	return t, nil
}

func (m *MemoryStore) ListThemes(_ context.Context) ([]domain.Theme, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	out := []domain.Theme{}
	for _, t := range m.themes {
		out = append(out, t)
	}
	return out, nil
}

func (m *MemoryStore) CreateCategory(_ context.Context, c domain.PlaygroundCategory) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.categories[c.ID] = c
	return nil
}

func (m *MemoryStore) CreateTag(_ context.Context, t domain.PlaygroundTag) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.tags[t.ID] = t
	return nil
}

func (m *MemoryStore) CreateItem(_ context.Context, item domain.PlaygroundItem) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if item.CreatedAt.IsZero() {
		item.CreatedAt = time.Now()
	}
	m.items[item.ID] = item
	return nil
}

func (m *MemoryStore) SearchItems(_ context.Context, categoryID string, tags []string, query string) ([]domain.PlaygroundItem, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	tagSet := map[string]struct{}{}
	for _, t := range tags {
		tagSet[t] = struct{}{}
	}
	q := strings.ToLower(strings.TrimSpace(query))
	out := []domain.PlaygroundItem{}
	for _, it := range m.items {
		if categoryID != "" && it.CategoryID != categoryID {
			continue
		}
		if q != "" && !strings.Contains(strings.ToLower(it.Title), q) {
			continue
		}
		ok := true
		for t := range tagSet {
			found := false
			for _, has := range it.Tags {
				if has == t {
					found = true
					break
				}
			}
			if !found {
				ok = false
				break
			}
		}
		if ok {
			out = append(out, it)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].CreatedAt.After(out[j].CreatedAt) })
	return out, nil
}

func (m *MemoryStore) CreateRun(_ context.Context, run domain.AgentRun) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if run.StartedAt.IsZero() {
		run.StartedAt = time.Now()
	}
	m.runs[run.ID] = run
	return nil
}

func (m *MemoryStore) UpdateRun(_ context.Context, run domain.AgentRun) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.runs[run.ID]; !ok {
		return ErrNotFound
	}
	m.runs[run.ID] = run
	return nil
}

func (m *MemoryStore) GetRun(_ context.Context, id string) (domain.AgentRun, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	r, ok := m.runs[id]
	if !ok {
		return domain.AgentRun{}, ErrNotFound
	}
	return r, nil
}

func (m *MemoryStore) CreateEvent(_ context.Context, e domain.AgentEvent) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if e.CreatedAt.IsZero() {
		e.CreatedAt = time.Now()
	}
	m.events[e.RunID] = append(m.events[e.RunID], e)
	return nil
}

func (m *MemoryStore) ListEventsByRun(_ context.Context, runID string) ([]domain.AgentEvent, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	items := append([]domain.AgentEvent(nil), m.events[runID]...)
	sort.Slice(items, func(i, j int) bool { return items[i].CreatedAt.Before(items[j].CreatedAt) })
	return items, nil
}

var _ store.Repositories = (*MemoryStore)(nil)
var _ store.WorkspaceRepository = (*MemoryStore)(nil)
var _ store.SessionRepository = (*MemoryStore)(nil)
var _ store.VersionRepository = (*MemoryStore)(nil)
var _ store.ProviderRepository = (*MemoryStore)(nil)
var _ store.ThemeRepository = (*MemoryStore)(nil)
var _ store.PlaygroundRepository = (*MemoryStore)(nil)
var _ store.EventRepository = (*MemoryStore)(nil)
