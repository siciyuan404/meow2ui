-- +goose Up
CREATE TABLE IF NOT EXISTS benchmark_targets (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  kind TEXT NOT NULL,
  status TEXT NOT NULL DEFAULT 'active',
  metadata_json TEXT NOT NULL DEFAULT '{}',
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS benchmark_env_snapshots (
  id TEXT PRIMARY KEY,
  run_id TEXT NOT NULL,
  git_sha TEXT NOT NULL,
  dataset_version TEXT NOT NULL,
  runtime_info TEXT NOT NULL DEFAULT '{}',
  created_at TIMESTAMP NOT NULL,
  FOREIGN KEY (run_id) REFERENCES benchmark_runs(id)
);

CREATE INDEX IF NOT EXISTS idx_benchmark_env_run ON benchmark_env_snapshots(run_id);

-- +goose Down
DROP TABLE IF EXISTS benchmark_env_snapshots;
DROP TABLE IF EXISTS benchmark_targets;
