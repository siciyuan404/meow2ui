-- +goose Up
CREATE TABLE IF NOT EXISTS orgs (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  status TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS projects (
  id TEXT PRIMARY KEY,
  org_id TEXT NOT NULL,
  name TEXT NOT NULL,
  status TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  FOREIGN KEY (org_id) REFERENCES orgs(id)
);

CREATE TABLE IF NOT EXISTS project_members (
  id TEXT PRIMARY KEY,
  project_id TEXT NOT NULL,
  user_id TEXT NOT NULL,
  role TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  FOREIGN KEY (project_id) REFERENCES projects(id)
);

ALTER TABLE workspaces ADD COLUMN IF NOT EXISTS project_id TEXT;
ALTER TABLE sessions ADD COLUMN IF NOT EXISTS project_id TEXT;

CREATE TABLE IF NOT EXISTS audit_export_jobs (
  id TEXT PRIMARY KEY,
  org_id TEXT NOT NULL,
  requested_by TEXT NOT NULL,
  format TEXT NOT NULL,
  range_start TIMESTAMP NOT NULL,
  range_end TIMESTAMP NOT NULL,
  status TEXT NOT NULL,
  artifact_uri TEXT,
  created_at TIMESTAMP NOT NULL,
  finished_at TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS audit_export_jobs;
ALTER TABLE sessions DROP COLUMN IF EXISTS project_id;
ALTER TABLE workspaces DROP COLUMN IF EXISTS project_id;
DROP TABLE IF EXISTS project_members;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS orgs;
