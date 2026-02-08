package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/example/a2ui-go-agent-platform/internal/infra/bootstrap"
	"github.com/example/a2ui-go-agent-platform/internal/infra/config"
	"github.com/example/a2ui-go-agent-platform/internal/infra/db"
	"github.com/example/a2ui-go-agent-platform/pkg/agent"
	"github.com/example/a2ui-go-agent-platform/pkg/debugger"
	"github.com/example/a2ui-go-agent-platform/pkg/domain"
	"github.com/example/a2ui-go-agent-platform/pkg/flow"
	"github.com/example/a2ui-go-agent-platform/pkg/httpx"
	"github.com/example/a2ui-go-agent-platform/pkg/marketplace"
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
	debugSvc := debugger.NewService(app.Events)
	marketSvc := marketplace.NewService()
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

	mux.HandleFunc("/api/v1/debug/runs", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			httpx.WriteError(w, r, httpx.ErrMethodNotAllowed)
			return
		}
		filter := debugger.RunFilter{
			SessionID: r.URL.Query().Get("session_id"),
			Status:    r.URL.Query().Get("status"),
		}
		if from := r.URL.Query().Get("from"); from != "" {
			if t, err := time.Parse(time.RFC3339, from); err == nil {
				filter.From = t
			}
		}
		if to := r.URL.Query().Get("to"); to != "" {
			if t, err := time.Parse(time.RFC3339, to); err == nil {
				filter.To = t
			}
		}
		items, err := debugSvc.ListRuns(r.Context(), filter)
		if err != nil {
			httpx.WriteError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"runs": items})
	})
	mux.HandleFunc("/api/v1/debug/runs/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			httpx.WriteError(w, r, httpx.ErrMethodNotAllowed)
			return
		}
		path := r.URL.Path
		prefix := "/api/v1/debug/runs/"
		if len(path) <= len(prefix) {
			httpx.WriteError(w, r, httpx.ErrMethodNotAllowed)
			return
		}
		runID := path[len(prefix):]
		if len(runID) > 5 && runID[len(runID)-5:] == "/cost" {
			runID = runID[:len(runID)-5]
			cost, err := debugSvc.GetRunCost(r.Context(), runID)
			if err != nil {
				httpx.WriteError(w, r, err)
				return
			}
			httpx.WriteJSON(w, http.StatusOK, cost)
			return
		}
		detail, err := debugSvc.GetRunDetail(r.Context(), runID)
		if err != nil {
			httpx.WriteError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, detail)
	})

	mux.HandleFunc("/api/v1/flows", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			items, err := app.Flow.ListTemplates(r.Context())
			if err != nil {
				httpx.WriteError(w, r, err)
				return
			}
			httpx.WriteJSON(w, http.StatusOK, map[string]any{"items": items})
			return
		}
		if r.Method == http.MethodPost {
			var body struct {
				Name       string          `json:"name"`
				Version    string          `json:"version"`
				Definition flow.Definition `json:"definition"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				httpx.WriteError(w, r, err)
				return
			}
			tpl, ver, err := app.Flow.CreateTemplate(r.Context(), body.Name, body.Version, body.Definition)
			if err != nil {
				httpx.WriteError(w, r, err)
				return
			}
			httpx.WriteJSON(w, http.StatusOK, map[string]any{"template": tpl, "version": ver})
			return
		}
		httpx.WriteError(w, r, httpx.ErrMethodNotAllowed)
	})
	mux.HandleFunc("/api/v1/flows/bind-session", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpx.WriteError(w, r, httpx.ErrMethodNotAllowed)
			return
		}
		var body struct {
			SessionID  string `json:"sessionId"`
			TemplateID string `json:"templateId"`
			Version    string `json:"version"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			httpx.WriteError(w, r, err)
			return
		}
		if err := app.Flow.BindSession(r.Context(), body.SessionID, body.TemplateID, body.Version); err != nil {
			httpx.WriteError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"ok": true})
	})

	mux.HandleFunc("/api/v1/marketplace/templates", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			items := marketSvc.Search(r.Context(), marketplace.SearchInput{
				Category: r.URL.Query().Get("category"),
				Tag:      r.URL.Query().Get("tag"),
				Query:    r.URL.Query().Get("q"),
			})
			httpx.WriteJSON(w, http.StatusOK, map[string]any{"items": items})
			return
		}
		if r.Method == http.MethodPost {
			var body struct {
				Name      string   `json:"name"`
				Category  string   `json:"category"`
				Tags      []string `json:"tags"`
				Schema    string   `json:"schema"`
				Theme     string   `json:"theme"`
				Owner     string   `json:"owner"`
				SessionID string   `json:"sessionId"`
				VersionID string   `json:"versionId"`
			}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				httpx.WriteError(w, r, err)
				return
			}
			tpl, err := marketSvc.SaveFromSession(r.Context(), marketplace.SaveFromSessionInput{
				Name:      body.Name,
				Category:  body.Category,
				Tags:      body.Tags,
				Schema:    body.Schema,
				Theme:     body.Theme,
				Owner:     body.Owner,
				SessionID: body.SessionID,
				VersionID: body.VersionID,
			})
			if err != nil {
				httpx.WriteError(w, r, err)
				return
			}
			httpx.WriteJSON(w, http.StatusOK, tpl)
			return
		}
		httpx.WriteError(w, r, httpx.ErrMethodNotAllowed)
	})

	mux.HandleFunc("/api/v1/marketplace/review", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpx.WriteError(w, r, httpx.ErrMethodNotAllowed)
			return
		}
		var body struct {
			TemplateID string `json:"templateId"`
			Decision   string `json:"decision"`
			Note       string `json:"note"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			httpx.WriteError(w, r, err)
			return
		}
		if body.Decision == "submit" {
			if err := marketSvc.SubmitReview(r.Context(), body.TemplateID); err != nil {
				httpx.WriteError(w, r, err)
				return
			}
		} else {
			if err := marketSvc.Review(r.Context(), body.TemplateID, body.Decision, body.Note); err != nil {
				httpx.WriteError(w, r, err)
				return
			}
		}
		httpx.WriteJSON(w, http.StatusOK, map[string]any{"ok": true})
	})

	mux.HandleFunc("/api/v1/marketplace/ratings", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpx.WriteError(w, r, httpx.ErrMethodNotAllowed)
			return
		}
		var body struct {
			TemplateID string `json:"templateId"`
			UserID     string `json:"userId"`
			Score      int    `json:"score"`
			Comment    string `json:"comment"`
			Flag       bool   `json:"flag"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			httpx.WriteError(w, r, err)
			return
		}
		if body.Flag {
			if err := marketSvc.FlagComment(r.Context(), body.TemplateID, body.UserID); err != nil {
				httpx.WriteError(w, r, err)
				return
			}
			httpx.WriteJSON(w, http.StatusOK, map[string]any{"ok": true})
			return
		}
		item, err := marketSvc.AddRating(r.Context(), body.TemplateID, body.UserID, body.Score, body.Comment)
		if err != nil {
			httpx.WriteError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, item)
	})

	mux.HandleFunc("/api/v1/marketplace/apply", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpx.WriteError(w, r, httpx.ErrMethodNotAllowed)
			return
		}
		var body struct {
			TemplateID string `json:"templateId"`
			SessionID  string `json:"sessionId"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			httpx.WriteError(w, r, err)
			return
		}
		out, err := marketSvc.ApplyTemplate(r.Context(), body.TemplateID, body.SessionID)
		if err != nil {
			httpx.WriteError(w, r, err)
			return
		}
		httpx.WriteJSON(w, http.StatusOK, out)
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
			Media     []struct {
				Type     string         `json:"type"`
				Ref      string         `json:"ref"`
				Metadata map[string]any `json:"metadata"`
			} `json:"media"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			httpx.WriteError(w, r, err)
			return
		}
		media := make([]domain.MultimodalInput, 0, len(body.Media))
		for _, m := range body.Media {
			media = append(media, domain.MultimodalInput{Type: domain.MediaType(m.Type), Ref: m.Ref, Metadata: m.Metadata})
		}
		out, err := app.Agent.Run(r.Context(), agent.RunInput{SessionID: body.SessionID, Prompt: body.Prompt, OnlyArea: body.OnlyArea, Media: media})
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
