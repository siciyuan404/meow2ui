-- +goose Up
CREATE TABLE IF NOT EXISTS users (
  id TEXT PRIMARY KEY,
  email TEXT NOT NULL UNIQUE,
  password_hash TEXT NOT NULL,
  status TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS auth_sessions (
  id TEXT PRIMARY KEY,
  user_id TEXT NOT NULL,
  token_hash TEXT NOT NULL,
  expires_at TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

ALTER TABLE workspaces ADD COLUMN IF NOT EXISTS owner_user_id TEXT NOT NULL DEFAULT 'system';
ALTER TABLE sessions ADD COLUMN IF NOT EXISTS owner_user_id TEXT NOT NULL DEFAULT 'system';
ALTER TABLE playground_items ADD COLUMN IF NOT EXISTS owner_user_id TEXT NOT NULL DEFAULT 'system';

-- +goose Down
ALTER TABLE playground_items DROP COLUMN IF EXISTS owner_user_id;
ALTER TABLE sessions DROP COLUMN IF EXISTS owner_user_id;
ALTER TABLE workspaces DROP COLUMN IF EXISTS owner_user_id;
DROP TABLE IF EXISTS auth_sessions;
DROP TABLE IF EXISTS users;
