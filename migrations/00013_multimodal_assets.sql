-- +goose Up
CREATE TABLE IF NOT EXISTS schema_version_assets (
  id TEXT PRIMARY KEY,
  version_id TEXT NOT NULL,
  asset_type TEXT NOT NULL,
  asset_ref TEXT NOT NULL,
  metadata_json TEXT NOT NULL DEFAULT '{}',
  created_at TIMESTAMP NOT NULL,
  FOREIGN KEY (version_id) REFERENCES schema_versions(id)
);

CREATE INDEX IF NOT EXISTS idx_schema_version_assets_version ON schema_version_assets(version_id);

-- +goose Down
DROP TABLE IF EXISTS schema_version_assets;
