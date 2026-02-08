package sqlstore

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/domain"
	"github.com/example/a2ui-go-agent-platform/pkg/store"
	"github.com/example/a2ui-go-agent-platform/pkg/util"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Store {
	return &Store{pool: pool}
}

func (s *Store) Workspace() store.WorkspaceRepository   { return s }
func (s *Store) Session() store.SessionRepository       { return s }
func (s *Store) Version() store.VersionRepository       { return s }
func (s *Store) Provider() store.ProviderRepository     { return s }
func (s *Store) Theme() store.ThemeRepository           { return s }
func (s *Store) Playground() store.PlaygroundRepository { return s }
func (s *Store) Event() store.EventRepository           { return s }
func (s *Store) Flow() store.FlowRepository             { return s }

func (s *Store) CreateWorkspace(ctx context.Context, w domain.Workspace) error {
	_, err := s.pool.Exec(ctx, `INSERT INTO workspaces (id,name,root_path,created_at,updated_at) VALUES ($1,$2,$3,$4,$5)`, w.ID, w.Name, w.RootPath, w.CreatedAt, w.UpdatedAt)
	return mapErr(err)
}

func (s *Store) GetWorkspace(ctx context.Context, id string) (domain.Workspace, error) {
	var w domain.Workspace
	err := s.pool.QueryRow(ctx, `SELECT id,name,root_path,created_at,updated_at FROM workspaces WHERE id=$1`, id).Scan(&w.ID, &w.Name, &w.RootPath, &w.CreatedAt, &w.UpdatedAt)
	if err != nil {
		return domain.Workspace{}, mapErr(err)
	}
	return w, nil
}

func (s *Store) ListWorkspaces(ctx context.Context) ([]domain.Workspace, error) {
	rows, err := s.pool.Query(ctx, `SELECT id,name,root_path,created_at,updated_at FROM workspaces ORDER BY created_at ASC`)
	if err != nil {
		return nil, mapErr(err)
	}
	defer rows.Close()
	out := []domain.Workspace{}
	for rows.Next() {
		var w domain.Workspace
		if err := rows.Scan(&w.ID, &w.Name, &w.RootPath, &w.CreatedAt, &w.UpdatedAt); err != nil {
			return nil, mapErr(err)
		}
		out = append(out, w)
	}
	return out, nil
}

func (s *Store) CreateSession(ctx context.Context, ssn domain.Session) error {
	metadata, err := json.Marshal(ssn.Metadata)
	if err != nil {
		return err
	}
	_, err = s.pool.Exec(ctx, `INSERT INTO sessions (id,workspace_id,title,active_theme_id,status,metadata,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`, ssn.ID, ssn.WorkspaceID, ssn.Title, ssn.ActiveThemeID, string(ssn.Status), string(metadata), ssn.CreatedAt, ssn.UpdatedAt)
	return mapErr(err)
}

func (s *Store) GetSession(ctx context.Context, id string) (domain.Session, error) {
	var ssn domain.Session
	var status string
	var metadata string
	err := s.pool.QueryRow(ctx, `SELECT id,workspace_id,title,active_theme_id,status,metadata,created_at,updated_at FROM sessions WHERE id=$1`, id).Scan(&ssn.ID, &ssn.WorkspaceID, &ssn.Title, &ssn.ActiveThemeID, &status, &metadata, &ssn.CreatedAt, &ssn.UpdatedAt)
	if err != nil {
		return domain.Session{}, mapErr(err)
	}
	ssn.Status = domain.SessionStatus(status)
	_ = json.Unmarshal([]byte(metadata), &ssn.Metadata)
	return ssn, nil
}

func (s *Store) ListByWorkspace(ctx context.Context, workspaceID string) ([]domain.Session, error) {
	rows, err := s.pool.Query(ctx, `SELECT id,workspace_id,title,active_theme_id,status,metadata,created_at,updated_at FROM sessions WHERE workspace_id=$1 ORDER BY created_at ASC`, workspaceID)
	if err != nil {
		return nil, mapErr(err)
	}
	defer rows.Close()
	out := []domain.Session{}
	for rows.Next() {
		var ssn domain.Session
		var status string
		var metadata string
		if err := rows.Scan(&ssn.ID, &ssn.WorkspaceID, &ssn.Title, &ssn.ActiveThemeID, &status, &metadata, &ssn.CreatedAt, &ssn.UpdatedAt); err != nil {
			return nil, mapErr(err)
		}
		ssn.Status = domain.SessionStatus(status)
		_ = json.Unmarshal([]byte(metadata), &ssn.Metadata)
		out = append(out, ssn)
	}
	return out, nil
}

func (s *Store) UpdateSession(ctx context.Context, ssn domain.Session) error {
	metadata, err := json.Marshal(ssn.Metadata)
	if err != nil {
		return err
	}
	result, err := s.pool.Exec(ctx, `UPDATE sessions SET title=$2,active_theme_id=$3,status=$4,metadata=$5,updated_at=$6 WHERE id=$1`, ssn.ID, ssn.Title, ssn.ActiveThemeID, string(ssn.Status), string(metadata), time.Now())
	if err != nil {
		return mapErr(err)
	}
	if result.RowsAffected() == 0 {
		return store.ErrNotFound
	}
	return nil
}

func (s *Store) CreateVersion(ctx context.Context, v domain.SchemaVersion) error {
	_, err := s.pool.Exec(ctx, `INSERT INTO schema_versions (id,session_id,parent_version_id,schema_path,schema_hash,schema_json,summary,theme_snapshot_id,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`, v.ID, v.SessionID, nullableString(v.ParentVersionID), v.SchemaPath, v.SchemaHash, v.SchemaJSON, v.Summary, v.ThemeSnapshotID, v.CreatedAt)
	return mapErr(err)
}

func (s *Store) GetVersion(ctx context.Context, id string) (domain.SchemaVersion, error) {
	var v domain.SchemaVersion
	var parent *string
	err := s.pool.QueryRow(ctx, `SELECT id,session_id,parent_version_id,schema_path,schema_hash,schema_json,summary,theme_snapshot_id,created_at FROM schema_versions WHERE id=$1`, id).Scan(&v.ID, &v.SessionID, &parent, &v.SchemaPath, &v.SchemaHash, &v.SchemaJSON, &v.Summary, &v.ThemeSnapshotID, &v.CreatedAt)
	if err != nil {
		return domain.SchemaVersion{}, mapErr(err)
	}
	if parent != nil {
		v.ParentVersionID = *parent
	}
	return v, nil
}

func (s *Store) GetLatestBySession(ctx context.Context, sessionID string) (domain.SchemaVersion, error) {
	var v domain.SchemaVersion
	var parent *string
	err := s.pool.QueryRow(ctx, `SELECT id,session_id,parent_version_id,schema_path,schema_hash,schema_json,summary,theme_snapshot_id,created_at FROM schema_versions WHERE session_id=$1 ORDER BY created_at DESC LIMIT 1`, sessionID).Scan(&v.ID, &v.SessionID, &parent, &v.SchemaPath, &v.SchemaHash, &v.SchemaJSON, &v.Summary, &v.ThemeSnapshotID, &v.CreatedAt)
	if err != nil {
		return domain.SchemaVersion{}, mapErr(err)
	}
	if parent != nil {
		v.ParentVersionID = *parent
	}
	return v, nil
}

func (s *Store) ListBySession(ctx context.Context, sessionID string) ([]domain.SchemaVersion, error) {
	rows, err := s.pool.Query(ctx, `SELECT id,session_id,parent_version_id,schema_path,schema_hash,schema_json,summary,theme_snapshot_id,created_at FROM schema_versions WHERE session_id=$1 ORDER BY created_at ASC`, sessionID)
	if err != nil {
		return nil, mapErr(err)
	}
	defer rows.Close()
	out := []domain.SchemaVersion{}
	for rows.Next() {
		var v domain.SchemaVersion
		var parent *string
		if err := rows.Scan(&v.ID, &v.SessionID, &parent, &v.SchemaPath, &v.SchemaHash, &v.SchemaJSON, &v.Summary, &v.ThemeSnapshotID, &v.CreatedAt); err != nil {
			return nil, mapErr(err)
		}
		if parent != nil {
			v.ParentVersionID = *parent
		}
		out = append(out, v)
	}
	return out, nil
}

func (s *Store) CreateVersionAsset(ctx context.Context, asset domain.SchemaVersionAsset) error {
	_, err := s.pool.Exec(ctx, `INSERT INTO schema_version_assets (id,version_id,asset_type,asset_ref,metadata_json,created_at) VALUES ($1,$2,$3,$4,$5,$6)`, asset.ID, asset.VersionID, asset.AssetType, asset.AssetRef, asset.MetadataJSON, asset.CreatedAt)
	return mapErr(err)
}

func (s *Store) ListVersionAssets(ctx context.Context, versionID string) ([]domain.SchemaVersionAsset, error) {
	rows, err := s.pool.Query(ctx, `SELECT id,version_id,asset_type,asset_ref,metadata_json,created_at FROM schema_version_assets WHERE version_id=$1 ORDER BY created_at ASC`, versionID)
	if err != nil {
		return nil, mapErr(err)
	}
	defer rows.Close()
	out := []domain.SchemaVersionAsset{}
	for rows.Next() {
		var item domain.SchemaVersionAsset
		if err := rows.Scan(&item.ID, &item.VersionID, &item.AssetType, &item.AssetRef, &item.MetadataJSON, &item.CreatedAt); err != nil {
			return nil, mapErr(err)
		}
		out = append(out, item)
	}
	return out, nil
}

func (s *Store) CreateProvider(ctx context.Context, p domain.Provider) error {
	_, err := s.pool.Exec(ctx, `INSERT INTO providers (id,name,type,base_url,auth_ref,timeout_ms,enabled) VALUES ($1,$2,$3,$4,$5,$6,$7)`, p.ID, p.Name, p.Type, p.BaseURL, p.AuthRef, p.TimeoutMS, p.Enabled)
	return mapErr(err)
}

func (s *Store) CreateModel(ctx context.Context, model domain.Model) error {
	cap, err := json.Marshal(model.Capabilities)
	if err != nil {
		return err
	}
	metadata, err := json.Marshal(model.Metadata)
	if err != nil {
		return err
	}
	_, err = s.pool.Exec(ctx, `INSERT INTO models (id,provider_id,name,capabilities,metadata,context_limit,enabled) VALUES ($1,$2,$3,$4,$5,$6,$7)`, model.ID, model.ProviderID, model.Name, string(cap), string(metadata), model.ContextLimit, model.Enabled)
	return mapErr(err)
}

func (s *Store) ListProviders(ctx context.Context) ([]domain.Provider, error) {
	rows, err := s.pool.Query(ctx, `SELECT id,name,type,base_url,auth_ref,timeout_ms,enabled FROM providers ORDER BY name ASC`)
	if err != nil {
		return nil, mapErr(err)
	}
	defer rows.Close()
	out := []domain.Provider{}
	for rows.Next() {
		var p domain.Provider
		if err := rows.Scan(&p.ID, &p.Name, &p.Type, &p.BaseURL, &p.AuthRef, &p.TimeoutMS, &p.Enabled); err != nil {
			return nil, mapErr(err)
		}
		out = append(out, p)
	}
	return out, nil
}

func (s *Store) ListModels(ctx context.Context) ([]domain.Model, error) {
	rows, err := s.pool.Query(ctx, `SELECT id,provider_id,name,capabilities,metadata,context_limit,enabled FROM models ORDER BY name ASC`)
	if err != nil {
		return nil, mapErr(err)
	}
	defer rows.Close()
	return scanModels(rows)
}

func (s *Store) ListEnabledModels(ctx context.Context) ([]domain.Model, error) {
	rows, err := s.pool.Query(ctx, `SELECT id,provider_id,name,capabilities,metadata,context_limit,enabled FROM models WHERE enabled=true ORDER BY name ASC`)
	if err != nil {
		return nil, mapErr(err)
	}
	defer rows.Close()
	return scanModels(rows)
}

func (s *Store) CreateTheme(ctx context.Context, t domain.Theme) error {
	tokens, err := json.Marshal(t.TokenSet)
	if err != nil {
		return err
	}
	_, err = s.pool.Exec(ctx, `INSERT INTO themes (id,name,token_set,is_builtin,created_at) VALUES ($1,$2,$3,$4,$5)`, t.ID, t.Name, string(tokens), t.IsBuiltin, t.CreatedAt)
	return mapErr(err)
}

func (s *Store) GetTheme(ctx context.Context, id string) (domain.Theme, error) {
	var t domain.Theme
	var tokens string
	err := s.pool.QueryRow(ctx, `SELECT id,name,token_set,is_builtin,created_at FROM themes WHERE id=$1`, id).Scan(&t.ID, &t.Name, &tokens, &t.IsBuiltin, &t.CreatedAt)
	if err != nil {
		return domain.Theme{}, mapErr(err)
	}
	_ = json.Unmarshal([]byte(tokens), &t.TokenSet)
	return t, nil
}

func (s *Store) ListThemes(ctx context.Context) ([]domain.Theme, error) {
	rows, err := s.pool.Query(ctx, `SELECT id,name,token_set,is_builtin,created_at FROM themes ORDER BY created_at DESC`)
	if err != nil {
		return nil, mapErr(err)
	}
	defer rows.Close()
	out := []domain.Theme{}
	for rows.Next() {
		var t domain.Theme
		var tokens string
		if err := rows.Scan(&t.ID, &t.Name, &tokens, &t.IsBuiltin, &t.CreatedAt); err != nil {
			return nil, mapErr(err)
		}
		_ = json.Unmarshal([]byte(tokens), &t.TokenSet)
		out = append(out, t)
	}
	return out, nil
}

func (s *Store) CreateCategory(ctx context.Context, c domain.PlaygroundCategory) error {
	_, err := s.pool.Exec(ctx, `INSERT INTO playground_categories (id,name) VALUES ($1,$2)`, c.ID, c.Name)
	return mapErr(err)
}

func (s *Store) CreateTag(ctx context.Context, t domain.PlaygroundTag) error {
	_, err := s.pool.Exec(ctx, `INSERT INTO playground_tags (id,name) VALUES ($1,$2) ON CONFLICT (id) DO NOTHING`, t.ID, t.Name)
	return mapErr(err)
}

func (s *Store) CreateItem(ctx context.Context, item domain.PlaygroundItem) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return mapErr(err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	if item.CategoryID == "" {
		item.CategoryID = "default"
		_, _ = tx.Exec(ctx, `INSERT INTO playground_categories (id,name) VALUES ($1,$2) ON CONFLICT (id) DO NOTHING`, "default", "default")
	}

	_, err = tx.Exec(ctx, `INSERT INTO playground_items (id,title,category_id,source_session_id,source_version_id,theme_id,schema_snapshot,preview_ref,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`, item.ID, item.Title, item.CategoryID, safeEmpty(item.SourceSessionID), safeEmpty(item.SourceVersionID), safeEmpty(item.ThemeID), item.SchemaSnapshot, safeEmpty(item.PreviewRef), item.CreatedAt)
	if err != nil {
		return mapErr(err)
	}

	for _, tagName := range item.Tags {
		tagName = strings.TrimSpace(tagName)
		if tagName == "" {
			continue
		}
		var tagID string
		err := tx.QueryRow(ctx, `SELECT id FROM playground_tags WHERE name=$1 LIMIT 1`, tagName).Scan(&tagID)
		if errors.Is(err, pgx.ErrNoRows) {
			tagID = util.NewID("pgtag")
			if _, err := tx.Exec(ctx, `INSERT INTO playground_tags (id,name) VALUES ($1,$2)`, tagID, tagName); err != nil {
				return mapErr(err)
			}
		} else if err != nil {
			return mapErr(err)
		}
		if _, err := tx.Exec(ctx, `INSERT INTO playground_item_tags (item_id,tag_id) VALUES ($1,$2) ON CONFLICT (item_id,tag_id) DO NOTHING`, item.ID, tagID); err != nil {
			return mapErr(err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return mapErr(err)
	}
	return nil
}

func (s *Store) SearchItems(ctx context.Context, categoryID string, tags []string, query string) ([]domain.PlaygroundItem, error) {
	rows, err := s.pool.Query(ctx, `SELECT id,title,category_id,source_session_id,source_version_id,theme_id,schema_snapshot,preview_ref,created_at FROM playground_items ORDER BY created_at DESC`)
	if err != nil {
		return nil, mapErr(err)
	}
	defer rows.Close()
	out := []domain.PlaygroundItem{}
	for rows.Next() {
		var it domain.PlaygroundItem
		if err := rows.Scan(&it.ID, &it.Title, &it.CategoryID, &it.SourceSessionID, &it.SourceVersionID, &it.ThemeID, &it.SchemaSnapshot, &it.PreviewRef, &it.CreatedAt); err != nil {
			return nil, mapErr(err)
		}
		itemTags, err := s.listTagNamesByItem(ctx, it.ID)
		if err != nil {
			return nil, err
		}
		it.Tags = itemTags
		if !matchesFilter(it, categoryID, tags, query) {
			continue
		}
		out = append(out, it)
	}
	return out, nil
}

func (s *Store) listTagNamesByItem(ctx context.Context, itemID string) ([]string, error) {
	rows, err := s.pool.Query(ctx, `SELECT t.name FROM playground_tags t JOIN playground_item_tags it ON t.id=it.tag_id WHERE it.item_id=$1`, itemID)
	if err != nil {
		return nil, mapErr(err)
	}
	defer rows.Close()
	out := []string{}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, mapErr(err)
		}
		out = append(out, name)
	}
	return out, nil
}

func (s *Store) CreateRun(ctx context.Context, run domain.AgentRun) error {
	_, err := s.pool.Exec(ctx, `INSERT INTO agent_runs (id,session_id,request_text,status,started_at,ended_at) VALUES ($1,$2,$3,$4,$5,$6)`, run.ID, run.SessionID, run.RequestText, string(run.Status), run.StartedAt, run.EndedAt)
	return mapErr(err)
}

func (s *Store) UpdateRun(ctx context.Context, run domain.AgentRun) error {
	result, err := s.pool.Exec(ctx, `UPDATE agent_runs SET status=$2,ended_at=$3 WHERE id=$1`, run.ID, string(run.Status), run.EndedAt)
	if err != nil {
		return mapErr(err)
	}
	if result.RowsAffected() == 0 {
		return store.ErrNotFound
	}
	return nil
}

func (s *Store) GetRun(ctx context.Context, id string) (domain.AgentRun, error) {
	var run domain.AgentRun
	var status string
	err := s.pool.QueryRow(ctx, `SELECT id,session_id,request_text,status,started_at,ended_at FROM agent_runs WHERE id=$1`, id).Scan(&run.ID, &run.SessionID, &run.RequestText, &status, &run.StartedAt, &run.EndedAt)
	if err != nil {
		return domain.AgentRun{}, mapErr(err)
	}
	run.Status = domain.AgentRunStatus(status)
	return run, nil
}

func (s *Store) ListRuns(ctx context.Context) ([]domain.AgentRun, error) {
	rows, err := s.pool.Query(ctx, `SELECT id,session_id,request_text,status,started_at,ended_at FROM agent_runs ORDER BY started_at DESC`)
	if err != nil {
		return nil, mapErr(err)
	}
	defer rows.Close()
	out := []domain.AgentRun{}
	for rows.Next() {
		var run domain.AgentRun
		var status string
		if err := rows.Scan(&run.ID, &run.SessionID, &run.RequestText, &status, &run.StartedAt, &run.EndedAt); err != nil {
			return nil, mapErr(err)
		}
		run.Status = domain.AgentRunStatus(status)
		out = append(out, run)
	}
	return out, nil
}

func (s *Store) CreateEvent(ctx context.Context, e domain.AgentEvent) error {
	payload, err := json.Marshal(e.Payload)
	if err != nil {
		return err
	}
	_, err = s.pool.Exec(ctx, `INSERT INTO agent_events (id,run_id,step,payload,latency_ms,token_in,token_out,created_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`, e.ID, e.RunID, e.Step, string(payload), e.LatencyMS, e.TokenIn, e.TokenOut, e.CreatedAt)
	return mapErr(err)
}

func (s *Store) ListEventsByRun(ctx context.Context, runID string) ([]domain.AgentEvent, error) {
	rows, err := s.pool.Query(ctx, `SELECT id,run_id,step,payload,latency_ms,token_in,token_out,created_at FROM agent_events WHERE run_id=$1 ORDER BY created_at ASC`, runID)
	if err != nil {
		return nil, mapErr(err)
	}
	defer rows.Close()
	out := []domain.AgentEvent{}
	for rows.Next() {
		var e domain.AgentEvent
		var payload string
		if err := rows.Scan(&e.ID, &e.RunID, &e.Step, &payload, &e.LatencyMS, &e.TokenIn, &e.TokenOut, &e.CreatedAt); err != nil {
			return nil, mapErr(err)
		}
		_ = json.Unmarshal([]byte(payload), &e.Payload)
		out = append(out, e)
	}
	return out, nil
}

func (s *Store) CreateFlowTemplate(ctx context.Context, flow domain.FlowTemplate) error {
	_, err := s.pool.Exec(ctx, `INSERT INTO flow_templates (id,name,status,created_at) VALUES ($1,$2,$3,$4)`, flow.ID, flow.Name, string(flow.Status), flow.CreatedAt)
	return mapErr(err)
}

func (s *Store) ListFlowTemplates(ctx context.Context) ([]domain.FlowTemplate, error) {
	rows, err := s.pool.Query(ctx, `SELECT id,name,status,created_at FROM flow_templates ORDER BY created_at DESC`)
	if err != nil {
		return nil, mapErr(err)
	}
	defer rows.Close()
	out := []domain.FlowTemplate{}
	for rows.Next() {
		var item domain.FlowTemplate
		var status string
		if err := rows.Scan(&item.ID, &item.Name, &status, &item.CreatedAt); err != nil {
			return nil, mapErr(err)
		}
		item.Status = domain.FlowTemplateStatus(status)
		out = append(out, item)
	}
	return out, nil
}

func (s *Store) CreateFlowTemplateVersion(ctx context.Context, version domain.FlowTemplateVersion) error {
	_, err := s.pool.Exec(ctx, `INSERT INTO flow_template_versions (id,template_id,version,definition_json,created_at) VALUES ($1,$2,$3,$4,$5)`, version.ID, version.TemplateID, version.Version, version.DefinitionJSON, version.CreatedAt)
	return mapErr(err)
}

func (s *Store) GetFlowTemplateVersion(ctx context.Context, templateID, version string) (domain.FlowTemplateVersion, error) {
	var out domain.FlowTemplateVersion
	err := s.pool.QueryRow(ctx, `SELECT id,template_id,version,definition_json,created_at FROM flow_template_versions WHERE template_id=$1 AND version=$2`, templateID, version).Scan(&out.ID, &out.TemplateID, &out.Version, &out.DefinitionJSON, &out.CreatedAt)
	if err != nil {
		return domain.FlowTemplateVersion{}, mapErr(err)
	}
	return out, nil
}

func (s *Store) BindSessionFlow(ctx context.Context, binding domain.SessionFlowBinding) error {
	_, err := s.pool.Exec(ctx, `INSERT INTO session_flow_bindings (session_id,template_id,template_version,bound_at) VALUES ($1,$2,$3,$4) ON CONFLICT (session_id) DO UPDATE SET template_id=$2,template_version=$3,bound_at=$4`, binding.SessionID, binding.TemplateID, binding.TemplateVersion, binding.BoundAt)
	return mapErr(err)
}

func (s *Store) GetSessionFlowBinding(ctx context.Context, sessionID string) (domain.SessionFlowBinding, error) {
	var out domain.SessionFlowBinding
	err := s.pool.QueryRow(ctx, `SELECT session_id,template_id,template_version,bound_at FROM session_flow_bindings WHERE session_id=$1`, sessionID).Scan(&out.SessionID, &out.TemplateID, &out.TemplateVersion, &out.BoundAt)
	if err != nil {
		return domain.SessionFlowBinding{}, mapErr(err)
	}
	return out, nil
}

func mapErr(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return store.ErrNotFound
	}
	msg := strings.ToLower(err.Error())
	if strings.Contains(msg, "duplicate key") || strings.Contains(msg, "unique") {
		return fmt.Errorf("%w: %v", store.ErrConflict, err)
	}
	return err
}

func scanModels(rows pgx.Rows) ([]domain.Model, error) {
	out := []domain.Model{}
	for rows.Next() {
		var m domain.Model
		var capsRaw string
		var metadata string
		if err := rows.Scan(&m.ID, &m.ProviderID, &m.Name, &capsRaw, &metadata, &m.ContextLimit, &m.Enabled); err != nil {
			return nil, err
		}
		_ = json.Unmarshal([]byte(capsRaw), &m.Capabilities)
		_ = json.Unmarshal([]byte(metadata), &m.Metadata)
		out = append(out, m)
	}
	return out, nil
}

func nullableString(v string) any {
	if strings.TrimSpace(v) == "" {
		return nil
	}
	return v
}

func safeEmpty(v string) string {
	if strings.TrimSpace(v) == "" {
		return "-"
	}
	return v
}

func matchesFilter(item domain.PlaygroundItem, categoryID string, tags []string, query string) bool {
	if categoryID != "" && item.CategoryID != categoryID {
		return false
	}
	if query != "" && !strings.Contains(strings.ToLower(item.Title), strings.ToLower(strings.TrimSpace(query))) {
		return false
	}
	if len(tags) == 0 {
		return true
	}
	set := map[string]struct{}{}
	for _, t := range item.Tags {
		set[t] = struct{}{}
	}
	for _, t := range tags {
		if _, ok := set[t]; !ok {
			return false
		}
	}
	return true
}

var _ store.Repositories = (*Store)(nil)
var _ store.WorkspaceRepository = (*Store)(nil)
var _ store.SessionRepository = (*Store)(nil)
var _ store.VersionRepository = (*Store)(nil)
var _ store.ProviderRepository = (*Store)(nil)
var _ store.ThemeRepository = (*Store)(nil)
var _ store.PlaygroundRepository = (*Store)(nil)
var _ store.EventRepository = (*Store)(nil)
var _ store.FlowRepository = (*Store)(nil)
