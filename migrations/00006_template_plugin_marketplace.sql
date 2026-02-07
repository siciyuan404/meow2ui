-- +goose Up
CREATE TABLE IF NOT EXISTS templates (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  version TEXT NOT NULL,
  category TEXT NOT NULL,
  tags TEXT NOT NULL DEFAULT '[]',
  schema_snapshot TEXT NOT NULL,
  owner_user_id TEXT NOT NULL DEFAULT 'system',
  status TEXT NOT NULL DEFAULT 'draft',
  reviewed_by TEXT,
  reviewed_at TIMESTAMP,
  review_note TEXT,
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS template_versions (
  id TEXT PRIMARY KEY,
  template_id TEXT NOT NULL,
  version TEXT NOT NULL,
  schema_snapshot TEXT NOT NULL,
  changelog TEXT NOT NULL DEFAULT '',
  created_at TIMESTAMP NOT NULL,
  FOREIGN KEY (template_id) REFERENCES templates(id)
);

CREATE TABLE IF NOT EXISTS plugins (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  version TEXT NOT NULL,
  entrypoint TEXT NOT NULL,
  capabilities TEXT NOT NULL DEFAULT '[]',
  permissions TEXT NOT NULL DEFAULT '[]',
  owner_user_id TEXT NOT NULL DEFAULT 'system',
  status TEXT NOT NULL DEFAULT 'draft',
  reviewed_by TEXT,
  reviewed_at TIMESTAMP,
  review_note TEXT,
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS project_plugins (
  workspace_id TEXT NOT NULL,
  plugin_id TEXT NOT NULL,
  enabled BOOLEAN NOT NULL DEFAULT false,
  config TEXT NOT NULL DEFAULT '{}',
  PRIMARY KEY (workspace_id, plugin_id),
  FOREIGN KEY (workspace_id) REFERENCES workspaces(id),
  FOREIGN KEY (plugin_id) REFERENCES plugins(id)
);

CREATE INDEX IF NOT EXISTS idx_templates_status ON templates(status);
CREATE INDEX IF NOT EXISTS idx_templates_category ON templates(category);
CREATE INDEX IF NOT EXISTS idx_template_versions_template ON template_versions(template_id);
CREATE INDEX IF NOT EXISTS idx_plugins_status ON plugins(status);

-- +goose Down
DROP TABLE IF EXISTS project_plugins;
DROP TABLE IF EXISTS template_versions;
DROP TABLE IF EXISTS plugins;
DROP TABLE IF EXISTS templates;
