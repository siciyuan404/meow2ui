-- +goose Up
CREATE TABLE IF NOT EXISTS product_events (
  id TEXT PRIMARY KEY,
  event_type TEXT NOT NULL,
  user_id TEXT,
  workspace_id TEXT,
  session_id TEXT,
  run_id TEXT,
  properties TEXT NOT NULL,
  occurred_at TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_product_events_type_time ON product_events(event_type, occurred_at);
CREATE INDEX IF NOT EXISTS idx_product_events_user_time ON product_events(user_id, occurred_at);

-- +goose Down
DROP TABLE IF EXISTS product_events;
