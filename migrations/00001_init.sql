-- +goose Up
CREATE TABLE IF NOT EXISTS workspaces (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  root_path TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS sessions (
  id TEXT PRIMARY KEY,
  workspace_id TEXT NOT NULL,
  title TEXT NOT NULL,
  active_theme_id TEXT NOT NULL,
  status TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  FOREIGN KEY (workspace_id) REFERENCES workspaces(id)
);

CREATE TABLE IF NOT EXISTS schema_versions (
  id TEXT PRIMARY KEY,
  session_id TEXT NOT NULL,
  parent_version_id TEXT,
  schema_path TEXT NOT NULL,
  schema_hash TEXT NOT NULL,
  schema_json TEXT NOT NULL,
  summary TEXT NOT NULL,
  theme_snapshot_id TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  FOREIGN KEY (session_id) REFERENCES sessions(id)
);

CREATE TABLE IF NOT EXISTS providers (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  type TEXT NOT NULL,
  base_url TEXT NOT NULL,
  auth_ref TEXT NOT NULL,
  timeout_ms INTEGER NOT NULL,
  enabled BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS models (
  id TEXT PRIMARY KEY,
  provider_id TEXT NOT NULL,
  name TEXT NOT NULL,
  capabilities TEXT NOT NULL,
  context_limit INTEGER NOT NULL,
  enabled BOOLEAN NOT NULL,
  FOREIGN KEY (provider_id) REFERENCES providers(id)
);

CREATE TABLE IF NOT EXISTS themes (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  token_set TEXT NOT NULL,
  is_builtin BOOLEAN NOT NULL,
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS playground_categories (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS playground_tags (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS playground_items (
  id TEXT PRIMARY KEY,
  title TEXT NOT NULL,
  category_id TEXT NOT NULL,
  source_session_id TEXT NOT NULL,
  source_version_id TEXT NOT NULL,
  theme_id TEXT NOT NULL,
  schema_snapshot TEXT NOT NULL,
  preview_ref TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  FOREIGN KEY (category_id) REFERENCES playground_categories(id)
);

CREATE TABLE IF NOT EXISTS playground_item_tags (
  item_id TEXT NOT NULL,
  tag_id TEXT NOT NULL,
  PRIMARY KEY (item_id, tag_id),
  FOREIGN KEY (item_id) REFERENCES playground_items(id),
  FOREIGN KEY (tag_id) REFERENCES playground_tags(id)
);

CREATE TABLE IF NOT EXISTS agent_runs (
  id TEXT PRIMARY KEY,
  session_id TEXT NOT NULL,
  request_text TEXT NOT NULL,
  status TEXT NOT NULL,
  started_at TIMESTAMP NOT NULL,
  ended_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS agent_events (
  id TEXT PRIMARY KEY,
  run_id TEXT NOT NULL,
  step TEXT NOT NULL,
  payload TEXT NOT NULL,
  latency_ms INTEGER NOT NULL,
  token_in INTEGER NOT NULL,
  token_out INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL,
  FOREIGN KEY (run_id) REFERENCES agent_runs(id)
);

CREATE INDEX IF NOT EXISTS idx_sessions_workspace ON sessions(workspace_id);
CREATE INDEX IF NOT EXISTS idx_schema_versions_session ON schema_versions(session_id);
CREATE INDEX IF NOT EXISTS idx_models_provider ON models(provider_id);
CREATE INDEX IF NOT EXISTS idx_playground_items_category ON playground_items(category_id);
CREATE INDEX IF NOT EXISTS idx_agent_events_run ON agent_events(run_id);

-- +goose Down
DROP TABLE IF EXISTS agent_events;
DROP TABLE IF EXISTS agent_runs;
DROP TABLE IF EXISTS playground_item_tags;
DROP TABLE IF EXISTS playground_items;
DROP TABLE IF EXISTS playground_tags;
DROP TABLE IF EXISTS playground_categories;
DROP TABLE IF EXISTS themes;
DROP TABLE IF EXISTS models;
DROP TABLE IF EXISTS providers;
DROP TABLE IF EXISTS schema_versions;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS workspaces;
