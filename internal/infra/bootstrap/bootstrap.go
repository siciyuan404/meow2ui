package bootstrap

import (
	"context"
	"fmt"
	"os"

	"github.com/example/a2ui-go-agent-platform/internal/infra/db"
	"github.com/example/a2ui-go-agent-platform/internal/infra/memorystore"
	"github.com/example/a2ui-go-agent-platform/internal/infra/sqlstore"
	"github.com/example/a2ui-go-agent-platform/pkg/a2ui"
	"github.com/example/a2ui-go-agent-platform/pkg/agent"
	"github.com/example/a2ui-go-agent-platform/pkg/domain"
	"github.com/example/a2ui-go-agent-platform/pkg/events"
	"github.com/example/a2ui-go-agent-platform/pkg/flow"
	flowruntime "github.com/example/a2ui-go-agent-platform/pkg/flow/runtime"
	"github.com/example/a2ui-go-agent-platform/pkg/guardrail"
	"github.com/example/a2ui-go-agent-platform/pkg/playground"
	"github.com/example/a2ui-go-agent-platform/pkg/playground/retrieval"
	"github.com/example/a2ui-go-agent-platform/pkg/provider"
	"github.com/example/a2ui-go-agent-platform/pkg/session"
	"github.com/example/a2ui-go-agent-platform/pkg/store"
	"github.com/example/a2ui-go-agent-platform/pkg/telemetry"
	"github.com/example/a2ui-go-agent-platform/pkg/theme"
	"github.com/example/a2ui-go-agent-platform/pkg/util"
	"github.com/example/a2ui-go-agent-platform/pkg/workspace"
)

type App struct {
	Store      store.Repositories
	Workspace  *workspace.Service
	Session    *session.Service
	Provider   *provider.Service
	A2UI       *a2ui.Service
	Guardrail  *guardrail.Service
	Theme      *theme.Service
	Playground *playground.Service
	Events     *events.Service
	Telemetry  *telemetry.Service
	Flow       *flow.Service
	Agent      *agent.Service
}

func New(ctx context.Context) (*App, error) {
	st, err := initStore(ctx)
	if err != nil {
		return nil, err
	}
	workspaceSvc := workspace.NewService(st.Workspace())
	sessionSvc := session.NewService(st.Session(), st.Version())
	registry := provider.NewRegistry()
	registry.Register("mock", provider.MockAdapter{})
	router := provider.NewRouter(st.Provider())
	providerSvc := provider.NewService(router, registry)
	a2uiSvc := a2ui.NewService()
	guardSvc := guardrail.NewService()
	themeSvc := theme.NewService(st.Theme(), st.Session())
	playSvc := playground.NewService(st.Playground())
	eventSvc := events.NewService(st.Event())
	telemetrySvc := telemetry.NewService()
	flowSvc := flow.NewService(st.Flow())
	flowRuntimeSvc := flowruntime.NewService(providerSvc)
	retriever := retrieval.NewKeywordRetriever(st.Playground())
	agentSvc := agent.NewService(providerSvc, a2uiSvc, sessionSvc, eventSvc, guardSvc, telemetrySvc, st.Version(), retriever, flowSvc, flowRuntimeSvc)

	providers, err := st.Provider().ListProviders(ctx)
	if err != nil {
		return nil, err
	}
	if len(providers) == 0 {
		seedProvider := domain.Provider{ID: util.NewID("prov"), Name: "mock-provider", Type: "mock", BaseURL: "local://mock", AuthRef: "", TimeoutMS: 30000, Enabled: true}
		if err := st.Provider().CreateProvider(ctx, seedProvider); err != nil {
			return nil, err
		}
		if err := st.Provider().CreateModel(ctx, domain.Model{ID: util.NewID("model"), ProviderID: seedProvider.ID, Name: "mock-text", Capabilities: []string{"text"}, Metadata: map[string]any{"role": "plan", "priority": 1}, ContextLimit: 8192, Enabled: true}); err != nil {
			return nil, err
		}
		if err := st.Provider().CreateModel(ctx, domain.Model{ID: util.NewID("model"), ProviderID: seedProvider.ID, Name: "mock-image-plan", Capabilities: []string{"image", "text"}, Metadata: map[string]any{"role": "plan_image", "priority": 1}, ContextLimit: 8192, Enabled: true}); err != nil {
			return nil, err
		}
		if err := st.Provider().CreateModel(ctx, domain.Model{ID: util.NewID("model"), ProviderID: seedProvider.ID, Name: "mock-audio-plan", Capabilities: []string{"audio", "text"}, Metadata: map[string]any{"role": "plan_audio", "priority": 1}, ContextLimit: 8192, Enabled: true}); err != nil {
			return nil, err
		}
		if err := st.Provider().CreateModel(ctx, domain.Model{ID: util.NewID("model"), ProviderID: seedProvider.ID, Name: "mock-text-emit", Capabilities: []string{"text"}, Metadata: map[string]any{"role": "emit", "priority": 1}, ContextLimit: 8192, Enabled: true}); err != nil {
			return nil, err
		}
		if err := st.Provider().CreateModel(ctx, domain.Model{ID: util.NewID("model"), ProviderID: seedProvider.ID, Name: "mock-text-repair", Capabilities: []string{"text"}, Metadata: map[string]any{"role": "repair", "priority": 1}, ContextLimit: 8192, Enabled: true}); err != nil {
			return nil, err
		}
	}

	return &App{
		Store:      st,
		Workspace:  workspaceSvc,
		Session:    sessionSvc,
		Provider:   providerSvc,
		A2UI:       a2uiSvc,
		Guardrail:  guardSvc,
		Theme:      themeSvc,
		Playground: playSvc,
		Events:     eventSvc,
		Telemetry:  telemetrySvc,
		Flow:       flowSvc,
		Agent:      agentSvc,
	}, nil
}

func initStore(ctx context.Context) (store.Repositories, error) {
	driver := os.Getenv("STORE_DRIVER")
	if driver == "" || driver == "memory" {
		return memorystore.New(), nil
	}
	if driver != "postgres" {
		return nil, fmt.Errorf("unsupported STORE_DRIVER: %s", driver)
	}
	cfg, err := db.LoadConfigFromEnv()
	if err != nil {
		return nil, err
	}
	pool, err := db.Connect(ctx, cfg)
	if err != nil {
		return nil, err
	}
	if os.Getenv("AUTO_MIGRATE") == "true" {
		if err := db.MigrateUp(ctx, cfg, "migrations"); err != nil {
			pool.Close()
			return nil, err
		}
	}
	return sqlstore.New(pool), nil
}
