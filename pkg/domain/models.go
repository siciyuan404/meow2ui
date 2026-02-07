package domain

import "time"

type Workspace struct {
	ID        string
	Name      string
	RootPath  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SessionStatus string

const (
	SessionActive SessionStatus = "active"
	SessionClosed SessionStatus = "closed"
)

type Session struct {
	ID            string
	WorkspaceID   string
	Title         string
	ActiveThemeID string
	Metadata      map[string]any
	Status        SessionStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UIComponent struct {
	ID       string                 `json:"id"`
	Type     string                 `json:"type"`
	Props    map[string]any         `json:"props,omitempty"`
	Children []UIComponent          `json:"children,omitempty"`
	Style    map[string]any         `json:"style,omitempty"`
	Binding  string                 `json:"binding,omitempty"`
	Events   []map[string]any       `json:"events,omitempty"`
	Loop     map[string]string      `json:"loop,omitempty"`
	Meta     map[string]interface{} `json:"meta,omitempty"`
}

type UISchema struct {
	Version string         `json:"version"`
	Root    UIComponent    `json:"root"`
	Data    map[string]any `json:"data,omitempty"`
}

type SchemaVersion struct {
	ID              string
	SessionID       string
	ParentVersionID string
	SchemaPath      string
	SchemaHash      string
	SchemaJSON      string
	Summary         string
	ThemeSnapshotID string
	CreatedAt       time.Time
}

type Provider struct {
	ID        string
	Name      string
	Type      string
	BaseURL   string
	AuthRef   string
	TimeoutMS int
	Enabled   bool
}

type Model struct {
	ID           string
	ProviderID   string
	Name         string
	Capabilities []string
	Metadata     map[string]any
	ContextLimit int
	Enabled      bool
}

type Theme struct {
	ID        string
	Name      string
	TokenSet  map[string]any
	IsBuiltin bool
	CreatedAt time.Time
}

type PlaygroundCategory struct {
	ID   string
	Name string
}

type PlaygroundTag struct {
	ID   string
	Name string
}

type PlaygroundItem struct {
	ID              string
	Title           string
	CategoryID      string
	SourceSessionID string
	SourceVersionID string
	ThemeID         string
	SchemaSnapshot  string
	PreviewRef      string
	Tags            []string
	CreatedAt       time.Time
}

type AgentRunStatus string

const (
	AgentRunRunning   AgentRunStatus = "running"
	AgentRunCompleted AgentRunStatus = "completed"
	AgentRunFailed    AgentRunStatus = "failed"
)

type AgentRun struct {
	ID          string
	SessionID   string
	RequestText string
	Status      AgentRunStatus
	StartedAt   time.Time
	EndedAt     *time.Time
}

type AgentEvent struct {
	ID        string
	RunID     string
	Step      string
	Payload   map[string]any
	LatencyMS int
	TokenIn   int
	TokenOut  int
	CreatedAt time.Time
}

type PatchOperation struct {
	Op     string         `json:"op"`
	Target string         `json:"target"`
	Value  map[string]any `json:"value,omitempty"`
	Reason string         `json:"reason,omitempty"`
}

type PatchDocument struct {
	Mode                string           `json:"mode"`
	TargetSchemaVersion string           `json:"targetSchemaVersion,omitempty"`
	Operations          []PatchOperation `json:"operations"`
}
