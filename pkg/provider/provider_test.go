package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/example/a2ui-go-agent-platform/pkg/domain"
	"github.com/example/a2ui-go-agent-platform/pkg/store"
)

// --- Mock ProviderRepository ---

type mockProviderRepo struct {
	providers []domain.Provider
	models    []domain.Model
	err       error
}

func (m *mockProviderRepo) CreateProvider(_ context.Context, p domain.Provider) error { return nil }
func (m *mockProviderRepo) CreateModel(_ context.Context, mod domain.Model) error     { return nil }
func (m *mockProviderRepo) ListProviders(_ context.Context) ([]domain.Provider, error) {
	return m.providers, m.err
}
func (m *mockProviderRepo) ListModels(_ context.Context) ([]domain.Model, error) {
	return m.models, m.err
}
func (m *mockProviderRepo) ListEnabledModels(_ context.Context) ([]domain.Model, error) {
	return m.models, m.err
}

var _ store.ProviderRepository = (*mockProviderRepo)(nil)

// --- Failing Adapter ---

type failingAdapter struct {
	retryable bool
	callCount int
}

func (f *failingAdapter) Name() string { return "failing" }
func (f *failingAdapter) Generate(_ context.Context, _ domain.Provider, _ domain.Model, _ GenerateRequest) (GenerateResponse, error) {
	f.callCount++
	return GenerateResponse{}, &ProviderError{
		Code:      "test_error",
		Retryable: f.retryable,
		Message:   "test failure",
	}
}

// ==================== Registry Tests ====================

func TestRegistry_RegisterAndGet(t *testing.T) {
	reg := NewRegistry()
	mock := MockAdapter{}
	reg.Register("OpenAI", mock)

	adapter, ok := reg.Get("openai")
	if !ok {
		t.Fatal("expected adapter to be found")
	}
	if adapter.Name() != "mock" {
		t.Fatalf("expected mock, got %s", adapter.Name())
	}
}

func TestRegistry_GetNotFound(t *testing.T) {
	reg := NewRegistry()
	_, ok := reg.Get("nonexistent")
	if ok {
		t.Fatal("expected adapter not found")
	}
}

func TestRegistry_CaseInsensitive(t *testing.T) {
	reg := NewRegistry()
	reg.Register("DeepSeek", MockAdapter{})

	for _, key := range []string{"deepseek", "DEEPSEEK", "DeepSeek"} {
		if _, ok := reg.Get(key); !ok {
			t.Fatalf("expected to find adapter with key %q", key)
		}
	}
}

// ==================== Router Tests ====================

func TestRouter_Route_Success(t *testing.T) {
	repo := &mockProviderRepo{
		providers: []domain.Provider{
			{ID: "p1", Name: "Test", Type: "mock", Enabled: true},
		},
		models: []domain.Model{
			{ID: "m1", ProviderID: "p1", Capabilities: []string{"text"}, Enabled: true},
		},
	}
	router := NewRouter(repo)
	p, m, err := router.Route(context.Background(), TaskPlan)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.ID != "p1" || m.ID != "m1" {
		t.Fatalf("unexpected result: provider=%s model=%s", p.ID, m.ID)
	}
}

func TestRouter_Route_NoModels(t *testing.T) {
	repo := &mockProviderRepo{
		providers: []domain.Provider{},
		models:    []domain.Model{},
	}
	router := NewRouter(repo)
	_, _, err := router.Route(context.Background(), TaskPlan)
	if err == nil {
		t.Fatal("expected error for no models")
	}
}

func TestRouter_Route_DisabledProvider(t *testing.T) {
	repo := &mockProviderRepo{
		providers: []domain.Provider{
			{ID: "p1", Name: "Test", Type: "mock", Enabled: false},
		},
		models: []domain.Model{
			{ID: "m1", ProviderID: "p1", Capabilities: []string{"text"}, Enabled: true},
		},
	}
	router := NewRouter(repo)
	_, _, err := router.Route(context.Background(), TaskPlan)
	if err == nil {
		t.Fatal("expected error for disabled provider")
	}
}

func TestRouter_Candidates_FilterByCapability(t *testing.T) {
	repo := &mockProviderRepo{
		providers: []domain.Provider{
			{ID: "p1", Type: "mock", Enabled: true},
		},
		models: []domain.Model{
			{ID: "m-text", ProviderID: "p1", Capabilities: []string{"text"}, Enabled: true},
			{ID: "m-image", ProviderID: "p1", Capabilities: []string{"image"}, Enabled: true},
		},
	}
	router := NewRouter(repo)

	candidates, err := router.Candidates(context.Background(), TaskPlanImage)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(candidates) != 1 || candidates[0].Model.ID != "m-image" {
		t.Fatalf("expected only image model, got %d candidates", len(candidates))
	}
}

func TestRouter_Candidates_SortByPriority(t *testing.T) {
	repo := &mockProviderRepo{
		providers: []domain.Provider{
			{ID: "p1", Type: "mock", Enabled: true},
		},
		models: []domain.Model{
			{ID: "m-low", ProviderID: "p1", Capabilities: []string{"text"}, Metadata: map[string]any{"priority": 50}, Enabled: true},
			{ID: "m-high", ProviderID: "p1", Capabilities: []string{"text"}, Metadata: map[string]any{"priority": 10}, Enabled: true},
		},
	}
	router := NewRouter(repo)
	candidates, _ := router.Candidates(context.Background(), TaskPlan)
	if len(candidates) != 2 {
		t.Fatalf("expected 2 candidates, got %d", len(candidates))
	}
	if candidates[0].Model.ID != "m-high" {
		t.Fatalf("expected m-high first, got %s", candidates[0].Model.ID)
	}
}

func TestRouter_Candidates_FilterByRole(t *testing.T) {
	repo := &mockProviderRepo{
		providers: []domain.Provider{
			{ID: "p1", Type: "mock", Enabled: true},
		},
		models: []domain.Model{
			{ID: "m-plan", ProviderID: "p1", Capabilities: []string{"text"}, Metadata: map[string]any{"role": "plan"}, Enabled: true},
			{ID: "m-emit", ProviderID: "p1", Capabilities: []string{"text"}, Metadata: map[string]any{"role": "emit"}, Enabled: true},
		},
	}
	router := NewRouter(repo)
	candidates, _ := router.Candidates(context.Background(), TaskPlan)
	if len(candidates) != 1 || candidates[0].Model.ID != "m-plan" {
		t.Fatalf("expected only plan model, got %d candidates", len(candidates))
	}
}

func TestRouter_Candidates_RepoError(t *testing.T) {
	repo := &mockProviderRepo{err: fmt.Errorf("db down")}
	router := NewRouter(repo)
	_, err := router.Candidates(context.Background(), TaskPlan)
	if err == nil {
		t.Fatal("expected error from repo")
	}
}

// ==================== Service Tests ====================

func TestService_Generate_Success(t *testing.T) {
	repo := &mockProviderRepo{
		providers: []domain.Provider{
			{ID: "p1", Type: "mock", Enabled: true},
		},
		models: []domain.Model{
			{ID: "m1", ProviderID: "p1", Capabilities: []string{"text"}, Enabled: true},
		},
	}
	reg := NewRegistry()
	reg.Register("mock", MockAdapter{})
	svc := NewService(NewRouter(repo), reg)

	resp, err := svc.Generate(context.Background(), TaskPlan, GenerateRequest{
		UserPrompt: "test",
		Context:    map[string]any{"expect": "plan"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Tokens != 120 {
		t.Fatalf("expected 120 tokens, got %d", resp.Tokens)
	}
}

func TestService_Generate_NoModels(t *testing.T) {
	repo := &mockProviderRepo{}
	reg := NewRegistry()
	svc := NewService(NewRouter(repo), reg)

	_, err := svc.Generate(context.Background(), TaskPlan, GenerateRequest{})
	if err == nil {
		t.Fatal("expected error for no models")
	}
}

func TestService_Generate_AdapterNotFound(t *testing.T) {
	repo := &mockProviderRepo{
		providers: []domain.Provider{
			{ID: "p1", Type: "unknown_type", Enabled: true},
		},
		models: []domain.Model{
			{ID: "m1", ProviderID: "p1", Capabilities: []string{"text"}, Enabled: true},
		},
	}
	reg := NewRegistry()
	svc := NewService(NewRouter(repo), reg)

	_, err := svc.Generate(context.Background(), TaskPlan, GenerateRequest{})
	if err == nil {
		t.Fatal("expected error for missing adapter")
	}
}

func TestService_Generate_RetryOnRetryableError(t *testing.T) {
	repo := &mockProviderRepo{
		providers: []domain.Provider{
			{ID: "p1", Type: "failing", Enabled: true},
		},
		models: []domain.Model{
			{ID: "m1", ProviderID: "p1", Capabilities: []string{"text"}, Enabled: true},
		},
	}
	fa := &failingAdapter{retryable: true}
	reg := NewRegistry()
	reg.Register("failing", fa)
	svc := NewService(NewRouter(repo), reg)

	_, err := svc.Generate(context.Background(), TaskPlan, GenerateRequest{})
	if err == nil {
		t.Fatal("expected error after retries")
	}
	if fa.callCount != 3 {
		t.Fatalf("expected 3 attempts, got %d", fa.callCount)
	}
}

func TestService_Generate_NoRetryOnNonRetryable(t *testing.T) {
	repo := &mockProviderRepo{
		providers: []domain.Provider{
			{ID: "p1", Type: "failing", Enabled: true},
		},
		models: []domain.Model{
			{ID: "m1", ProviderID: "p1", Capabilities: []string{"text"}, Enabled: true},
		},
	}
	fa := &failingAdapter{retryable: false}
	reg := NewRegistry()
	reg.Register("failing", fa)
	svc := NewService(NewRouter(repo), reg)

	_, err := svc.Generate(context.Background(), TaskPlan, GenerateRequest{})
	if err == nil {
		t.Fatal("expected error")
	}
	if fa.callCount != 1 {
		t.Fatalf("expected 1 attempt (no retry), got %d", fa.callCount)
	}
}

func TestService_Generate_FallbackToNextCandidate(t *testing.T) {
	repo := &mockProviderRepo{
		providers: []domain.Provider{
			{ID: "p1", Type: "failing", Enabled: true},
			{ID: "p2", Type: "mock", Enabled: true},
		},
		models: []domain.Model{
			{ID: "m1", ProviderID: "p1", Capabilities: []string{"text"}, Metadata: map[string]any{"priority": 1}, Enabled: true},
			{ID: "m2", ProviderID: "p2", Capabilities: []string{"text"}, Metadata: map[string]any{"priority": 2}, Enabled: true},
		},
	}
	reg := NewRegistry()
	reg.Register("failing", &failingAdapter{retryable: false})
	reg.Register("mock", MockAdapter{})
	svc := NewService(NewRouter(repo), reg)

	resp, err := svc.Generate(context.Background(), TaskPlan, GenerateRequest{UserPrompt: "hello"})
	if err != nil {
		t.Fatalf("expected fallback to succeed, got: %v", err)
	}
	if resp.Text != "hello" {
		t.Fatalf("expected fallback response, got: %s", resp.Text)
	}
}

// ==================== Helper Function Tests ====================

func TestCapabilityForTask(t *testing.T) {
	tests := []struct {
		task     TaskType
		expected string
	}{
		{TaskPlan, "text"},
		{TaskEmit, "text"},
		{TaskRepair, "text"},
		{TaskPlanImage, "image"},
		{TaskPlanAudio, "audio"},
		{TaskType("unknown"), "text"},
	}
	for _, tt := range tests {
		got := capabilityForTask(tt.task)
		if got != tt.expected {
			t.Errorf("capabilityForTask(%s) = %s, want %s", tt.task, got, tt.expected)
		}
	}
}

func TestHasCapability(t *testing.T) {
	if !hasCapability([]string{"text", "image"}, "TEXT") {
		t.Error("expected case-insensitive match")
	}
	if hasCapability([]string{"text"}, "image") {
		t.Error("expected no match")
	}
	if hasCapability(nil, "text") {
		t.Error("expected no match on nil")
	}
}

func TestModelMatchesRole(t *testing.T) {
	tests := []struct {
		name     string
		model    domain.Model
		role     string
		expected bool
	}{
		{"nil metadata", domain.Model{}, "plan", true},
		{"no role key", domain.Model{Metadata: map[string]any{"x": 1}}, "plan", true},
		{"matching role", domain.Model{Metadata: map[string]any{"role": "plan"}}, "plan", true},
		{"non-matching role", domain.Model{Metadata: map[string]any{"role": "emit"}}, "plan", false},
		{"case insensitive", domain.Model{Metadata: map[string]any{"role": "PLAN"}}, "plan", true},
		{"non-string role", domain.Model{Metadata: map[string]any{"role": 123}}, "plan", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := modelMatchesRole(tt.model, tt.role)
			if got != tt.expected {
				t.Errorf("modelMatchesRole() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestModelPriority(t *testing.T) {
	tests := []struct {
		name     string
		model    domain.Model
		expected int
	}{
		{"nil metadata", domain.Model{}, 100},
		{"no priority", domain.Model{Metadata: map[string]any{}}, 100},
		{"int priority", domain.Model{Metadata: map[string]any{"priority": 5}}, 5},
		{"float64 priority", domain.Model{Metadata: map[string]any{"priority": 3.0}}, 3},
		{"string priority", domain.Model{Metadata: map[string]any{"priority": "7"}}, 7},
		{"invalid string", domain.Model{Metadata: map[string]any{"priority": "abc"}}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := modelPriority(tt.model)
			if got != tt.expected {
				t.Errorf("modelPriority() = %d, want %d", got, tt.expected)
			}
		})
	}
}

// ==================== ProviderError Tests ====================

func TestProviderError_Error(t *testing.T) {
	e := &ProviderError{Code: "rate_limit", Message: "too many requests"}
	if e.Error() != "too many requests" {
		t.Fatalf("expected 'too many requests', got %s", e.Error())
	}
}

// ==================== MockAdapter Tests ====================

func TestMockAdapter_Name(t *testing.T) {
	m := MockAdapter{}
	if m.Name() != "mock" {
		t.Fatalf("expected 'mock', got %s", m.Name())
	}
}

func TestMockAdapter_Generate_Plan(t *testing.T) {
	m := MockAdapter{}
	resp, err := m.Generate(context.Background(), domain.Provider{}, domain.Model{}, GenerateRequest{
		Context: map[string]any{"expect": "plan"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Tokens != 120 {
		t.Fatalf("expected 120 tokens, got %d", resp.Tokens)
	}
}

func TestMockAdapter_Generate_Repair(t *testing.T) {
	m := MockAdapter{}
	resp, err := m.Generate(context.Background(), domain.Provider{}, domain.Model{}, GenerateRequest{
		UserPrompt: "fix this",
		Context:    map[string]any{"expect": "repair"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Text != "fix this" {
		t.Fatalf("expected echo, got %s", resp.Text)
	}
}

func TestMockAdapter_Generate_Default(t *testing.T) {
	m := MockAdapter{}
	resp, err := m.Generate(context.Background(), domain.Provider{}, domain.Model{}, GenerateRequest{
		UserPrompt: "hello",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Text != "hello" || resp.Tokens != 100 {
		t.Fatalf("unexpected response: text=%s tokens=%d", resp.Text, resp.Tokens)
	}
}
