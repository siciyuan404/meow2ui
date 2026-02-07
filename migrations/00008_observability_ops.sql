-- +goose Up
CREATE TABLE IF NOT EXISTS ops_alerts (
  id TEXT PRIMARY KEY,
  rule_id TEXT NOT NULL,
  status TEXT NOT NULL,
  message TEXT NOT NULL,
  started_at TIMESTAMP NOT NULL,
  resolved_at TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS ops_alerts;
