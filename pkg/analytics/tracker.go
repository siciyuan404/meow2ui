package analytics

import (
	"context"
	"encoding/json"
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Event struct {
	EventType   string
	UserID      string
	WorkspaceID string
	SessionID   string
	RunID       string
	Properties  map[string]any
	OccurredAt  time.Time
}

type Tracker struct {
	pool *pgxpool.Pool
}

func NewTracker(pool *pgxpool.Pool) *Tracker {
	return &Tracker{pool: pool}
}

func (t *Tracker) Track(ctx context.Context, e Event) error {
	if e.OccurredAt.IsZero() {
		e.OccurredAt = time.Now()
	}
	b, _ := json.Marshal(sanitizeProperties(e.Properties))
	_, err := t.pool.Exec(ctx, `INSERT INTO product_events (id,event_type,user_id,workspace_id,session_id,run_id,properties,occurred_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`, util.NewID("pevt"), e.EventType, nullable(e.UserID), nullable(e.WorkspaceID), nullable(e.SessionID), nullable(e.RunID), string(b), e.OccurredAt)
	return err
}

func sanitizeProperties(in map[string]any) map[string]any {
	if in == nil {
		return map[string]any{}
	}
	out := map[string]any{}
	for k, v := range in {
		if k == "apiKey" || k == "password" || k == "prompt" {
			continue
		}
		out[k] = v
	}
	return out
}

func nullable(v string) any {
	if v == "" {
		return nil
	}
	return v
}
