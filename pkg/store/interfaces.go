package store

import (
	"context"

	"github.com/example/a2ui-go-agent-platform/pkg/domain"
)

type WorkspaceRepository interface {
	CreateWorkspace(context.Context, domain.Workspace) error
	GetWorkspace(context.Context, string) (domain.Workspace, error)
	ListWorkspaces(context.Context) ([]domain.Workspace, error)
}

type SessionRepository interface {
	CreateSession(context.Context, domain.Session) error
	GetSession(context.Context, string) (domain.Session, error)
	ListByWorkspace(context.Context, string) ([]domain.Session, error)
	UpdateSession(context.Context, domain.Session) error
}

type VersionRepository interface {
	CreateVersion(context.Context, domain.SchemaVersion) error
	GetVersion(context.Context, string) (domain.SchemaVersion, error)
	GetLatestBySession(context.Context, string) (domain.SchemaVersion, error)
	ListBySession(context.Context, string) ([]domain.SchemaVersion, error)
	CreateVersionAsset(context.Context, domain.SchemaVersionAsset) error
	ListVersionAssets(context.Context, string) ([]domain.SchemaVersionAsset, error)
}

type ProviderRepository interface {
	CreateProvider(context.Context, domain.Provider) error
	CreateModel(context.Context, domain.Model) error
	ListProviders(context.Context) ([]domain.Provider, error)
	ListModels(context.Context) ([]domain.Model, error)
	ListEnabledModels(context.Context) ([]domain.Model, error)
}

type ThemeRepository interface {
	CreateTheme(context.Context, domain.Theme) error
	GetTheme(context.Context, string) (domain.Theme, error)
	ListThemes(context.Context) ([]domain.Theme, error)
}

type PlaygroundRepository interface {
	CreateCategory(context.Context, domain.PlaygroundCategory) error
	CreateTag(context.Context, domain.PlaygroundTag) error
	CreateItem(context.Context, domain.PlaygroundItem) error
	SearchItems(context.Context, string, []string, string) ([]domain.PlaygroundItem, error)
}

type EventRepository interface {
	CreateRun(context.Context, domain.AgentRun) error
	UpdateRun(context.Context, domain.AgentRun) error
	GetRun(context.Context, string) (domain.AgentRun, error)
	ListRuns(context.Context) ([]domain.AgentRun, error)
	CreateEvent(context.Context, domain.AgentEvent) error
	ListEventsByRun(context.Context, string) ([]domain.AgentEvent, error)
}

type FlowRepository interface {
	CreateFlowTemplate(context.Context, domain.FlowTemplate) error
	ListFlowTemplates(context.Context) ([]domain.FlowTemplate, error)
	CreateFlowTemplateVersion(context.Context, domain.FlowTemplateVersion) error
	GetFlowTemplateVersion(context.Context, string, string) (domain.FlowTemplateVersion, error)
	BindSessionFlow(context.Context, domain.SessionFlowBinding) error
	GetSessionFlowBinding(context.Context, string) (domain.SessionFlowBinding, error)
}

type Repositories interface {
	Workspace() WorkspaceRepository
	Session() SessionRepository
	Version() VersionRepository
	Provider() ProviderRepository
	Theme() ThemeRepository
	Playground() PlaygroundRepository
	Event() EventRepository
	Flow() FlowRepository
}
