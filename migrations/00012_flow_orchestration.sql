-- +goose Up
CREATE TABLE IF NOT EXISTS flow_templates (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  status TEXT NOT NULL DEFAULT 'draft',
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS flow_template_versions (
  id TEXT PRIMARY KEY,
  template_id TEXT NOT NULL,
  version TEXT NOT NULL,
  definition_json TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  FOREIGN KEY (template_id) REFERENCES flow_templates(id)
);

CREATE TABLE IF NOT EXISTS session_flow_bindings (
  session_id TEXT PRIMARY KEY,
  template_id TEXT NOT NULL,
  template_version TEXT NOT NULL,
  bound_at TIMESTAMP NOT NULL,
  FOREIGN KEY (session_id) REFERENCES sessions(id),
  FOREIGN KEY (template_id) REFERENCES flow_templates(id)
);

CREATE INDEX IF NOT EXISTS idx_flow_template_versions_template ON flow_template_versions(template_id);

-- +goose Down
DROP TABLE IF EXISTS session_flow_bindings;
DROP TABLE IF EXISTS flow_template_versions;
DROP TABLE IF EXISTS flow_templates;
