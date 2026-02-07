-- +goose Up
CREATE TABLE IF NOT EXISTS backup_jobs (
  id TEXT PRIMARY KEY,
  type TEXT NOT NULL,
  status TEXT NOT NULL,
  started_at TIMESTAMP NOT NULL,
  ended_at TIMESTAMP,
  artifact_uri TEXT,
  size_bytes BIGINT,
  checksum TEXT,
  error TEXT
);

-- +goose Down
DROP TABLE IF EXISTS backup_jobs;
