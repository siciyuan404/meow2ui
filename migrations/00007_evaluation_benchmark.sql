-- +goose Up
CREATE TABLE IF NOT EXISTS benchmark_suites (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  version TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS benchmark_cases (
  id TEXT PRIMARY KEY,
  suite_id TEXT NOT NULL,
  prompt TEXT NOT NULL,
  expected_constraints TEXT NOT NULL DEFAULT '{}',
  tags TEXT NOT NULL DEFAULT '[]',
  FOREIGN KEY (suite_id) REFERENCES benchmark_suites(id)
);

CREATE TABLE IF NOT EXISTS benchmark_runs (
  id TEXT PRIMARY KEY,
  suite_id TEXT NOT NULL,
  model_id TEXT NOT NULL,
  prompt_profile TEXT NOT NULL,
  started_at TIMESTAMP NOT NULL,
  ended_at TIMESTAMP,
  FOREIGN KEY (suite_id) REFERENCES benchmark_suites(id)
);

CREATE TABLE IF NOT EXISTS benchmark_run_results (
  id TEXT PRIMARY KEY,
  run_id TEXT NOT NULL,
  case_id TEXT NOT NULL,
  score TEXT NOT NULL,
  pass BOOLEAN NOT NULL,
  latency_ms INTEGER NOT NULL,
  tokens INTEGER NOT NULL,
  FOREIGN KEY (run_id) REFERENCES benchmark_runs(id),
  FOREIGN KEY (case_id) REFERENCES benchmark_cases(id)
);

CREATE INDEX IF NOT EXISTS idx_benchmark_cases_suite ON benchmark_cases(suite_id);
CREATE INDEX IF NOT EXISTS idx_benchmark_results_run ON benchmark_run_results(run_id);

-- +goose Down
DROP TABLE IF EXISTS benchmark_run_results;
DROP TABLE IF EXISTS benchmark_runs;
DROP TABLE IF EXISTS benchmark_cases;
DROP TABLE IF EXISTS benchmark_suites;
