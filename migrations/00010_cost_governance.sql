-- +goose Up
CREATE TABLE IF NOT EXISTS model_pricing (
  id TEXT PRIMARY KEY,
  provider_id TEXT NOT NULL,
  model_id TEXT NOT NULL,
  currency TEXT NOT NULL,
  input_per_1k NUMERIC(18,8) NOT NULL,
  output_per_1k NUMERIC(18,8) NOT NULL,
  effective_from TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS cost_usage (
  id TEXT PRIMARY KEY,
  run_id TEXT,
  session_id TEXT,
  workspace_id TEXT,
  user_id TEXT,
  provider_id TEXT NOT NULL,
  model_id TEXT NOT NULL,
  token_in INTEGER NOT NULL,
  token_out INTEGER NOT NULL,
  estimated_cost NUMERIC(18,8) NOT NULL,
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS budgets (
  id TEXT PRIMARY KEY,
  scope_type TEXT NOT NULL,
  scope_id TEXT NOT NULL,
  period TEXT NOT NULL,
  amount NUMERIC(18,8) NOT NULL,
  currency TEXT NOT NULL,
  thresholds TEXT NOT NULL DEFAULT '{}',
  action_policy TEXT NOT NULL DEFAULT '{}',
  enabled BOOLEAN NOT NULL DEFAULT true
);

CREATE TABLE IF NOT EXISTS budget_events (
  id TEXT PRIMARY KEY,
  budget_id TEXT NOT NULL,
  scope_id TEXT NOT NULL,
  event_type TEXT NOT NULL,
  current_spent NUMERIC(18,8) NOT NULL,
  threshold NUMERIC(18,8) NOT NULL,
  action_taken TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  FOREIGN KEY (budget_id) REFERENCES budgets(id)
);

CREATE INDEX IF NOT EXISTS idx_cost_usage_created_at ON cost_usage(created_at);

-- +goose Down
DROP TABLE IF EXISTS budget_events;
DROP TABLE IF EXISTS budgets;
DROP TABLE IF EXISTS cost_usage;
DROP TABLE IF EXISTS model_pricing;
