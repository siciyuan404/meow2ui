package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/example/a2ui-go-agent-platform/internal/infra/bootstrap"
	"github.com/example/a2ui-go-agent-platform/internal/infra/config"
	"github.com/example/a2ui-go-agent-platform/internal/infra/db"
	"github.com/example/a2ui-go-agent-platform/pkg/agent"
	"github.com/example/a2ui-go-agent-platform/pkg/httpx"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config invalid: %v", err)
	}
	app, err := bootstrap.New(ctx)
	if err != nil {
		log.Fatalf("bootstrap failed: %v", err)
	}

	mux := http.NewServeMux()
	var buildVersion = os.Getenv("APP_VERSION")
	if buildVersion == "" {
		buildVersion = "dev"
	}

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"ok": true, "version": buildVersion})
	})

	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		dbOK := true
		if cfg.StoreDriver == "postgres" {
			pool, err := db.Connect(r.Context(), cfg.Postgres)
			if err != nil {
				dbOK = false
			} else {
				pool.Close()
			}
		}
		payload := httpx.ReadyPayload(true, dbOK, buildVersion)
		status := http.StatusOK
		if !dbOK {
			status = http.StatusServiceUnavailable
		}
		httpx.WriteJSON(w, status, payload)
	})

	mux.HandleFunc("/version", func(w http.ResponseWriter, _ *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"version": buildVersion})
	})

	mux.HandleFunc("/api/v1/ops/health", func(w http.ResponseWriter, _ *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"ok": true})
	})
	mux.HandleFunc("/api/v1/ops/errors", func(w http.ResponseWriter, _ *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"errors": []any{}})
	})
	mux.HandleFunc("/api/v1/ops/capacity", func(w http.ResponseWriter, _ *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"capacity": map[string]any{"qps": 0, "p95": 0}})
	})
	mux.HandleFunc("/api/v1/ops/alerts", func(w http.ResponseWriter, _ *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"alerts": []any{}})
	})

	mux.HandleFunc("/api/v1/cost/summary", func(w http.ResponseWriter, _ *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"total": 0, "currency": "USD"})
	})
	mux.HandleFunc("/api/v1/cost/usage", func(w http.ResponseWriter, _ *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"items": []any{}})
	})
	mux.HandleFunc("/api/v1/cost/budgets", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			httpx.WriteJSON(w, http.StatusOK, map[string]any{"budgets": []any{}})
			return
		}
		if r.Method == http.MethodPost {
			httpx.WriteJSON(w, http.StatusOK, map[string]any{"ok": true})
			return
		}
		httpx.WriteError(w, r, httpx.ErrMethodNotAllowed)
	})

	mux.HandleFunc("/api/v1/evaluation/benchmark-runs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			httpx.WriteJSON(w, http.StatusOK, map[string]any{"runId": "bench-run-1"})
			return
		}
		httpx.WriteError(w, r, httpx.ErrMethodNotAllowed)
	})
	mux.HandleFunc("/api/v1/evaluation/benchmark-runs/bench-run-1", func(w http.ResponseWriter, _ *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"runId": "bench-run-1", "status": "completed"})
	})
	mux.HandleFunc("/api/v1/evaluation/benchmark-runs/bench-run-1/results", func(w http.ResponseWriter, _ *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"results": []any{}})
	})

	mux.HandleFunc("/api/v1/orgs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			httpx.WriteJSON(w, http.StatusOK, map[string]any{"id": "org-1"})
			return
		}
		httpx.WriteError(w, r, httpx.ErrMethodNotAllowed)
	})
	mux.HandleFunc("/api/v1/projects", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			httpx.WriteJSON(w, http.StatusOK, map[string]any{"id": "project-1"})
			return
		}
		httpx.WriteError(w, r, httpx.ErrMethodNotAllowed)
	})
	mux.HandleFunc("/api/v1/projects/project-1/members", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			httpx.WriteJSON(w, http.StatusOK, map[string]any{"ok": true})
			return
		}
		httpx.WriteError(w, r, httpx.ErrMethodNotAllowed)
	})
	mux.HandleFunc("/api/v1/audit/exports", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			httpx.WriteJSON(w, http.StatusOK, map[string]any{"id": "aexp-1"})
			return
		}
		httpx.WriteError(w, r, httpx.ErrMethodNotAllowed)
	})
	mux.HandleFunc("/api/v1/audit/exports/aexp-1", func(w http.ResponseWriter, _ *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"id": "aexp-1", "status": "completed"})
	})

	mux.HandleFunc("/api/v1/enterprise/sso/config", func(w http.ResponseWriter, _ *http.Request) {
		httpx.WriteJSON(w, http.StatusNotImplemented, map[string]any{"code": "not_implemented"})
	})
	mux.HandleFunc("/api/v1/enterprise/scim/sync", func(w http.ResponseWriter, _ *http.Request) {
		httpx.WriteJSON(w, http.StatusNotImplemented, map[string]any{"code": "not_implemented"})
	})

	mux.HandleFunc("/workspace/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpx.WriteError(w, r, httpx.ErrMethodNotAllowed)
			return
		}
		var body struct {
			Name string `json:"name"`
			Root string `json:"root"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			httpx.WriteError(w, r, err)
			return
		}
		ws, err := app.Workspace.Create(r.Context(), body.Name, body.Root)
		if err != nil {
			httpx.WriteError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, ws)
	})

	mux.HandleFunc("/session/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpx.WriteError(w, r, httpx.ErrMethodNotAllowed)
			return
		}
		var body struct {
			WorkspaceID string `json:"workspaceId"`
			Title       string `json:"title"`
			ThemeID     string `json:"themeId"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			httpx.WriteError(w, r, err)
			return
		}
		session, version, err := app.Session.Create(r.Context(), body.WorkspaceID, body.Title, body.ThemeID)
		if err != nil {
			httpx.WriteError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"session": session, "version": version})
	})

	mux.HandleFunc("/agent/run", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpx.WriteError(w, r, httpx.ErrMethodNotAllowed)
			return
		}
		var body struct {
			SessionID string `json:"sessionId"`
			Prompt    string `json:"prompt"`
			OnlyArea  string `json:"onlyArea"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			httpx.WriteError(w, r, err)
			return
		}
		out, err := app.Agent.Run(r.Context(), agent.RunInput{SessionID: body.SessionID, Prompt: body.Prompt, OnlyArea: body.OnlyArea})
		if err != nil {
			httpx.WriteError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, out)
	})

	handler := httpx.TraceMiddleware(mux)
	addr := cfg.ServerAddr
	log.Printf("a2ui server listening on %s", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
