package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/a2ui"
	"github.com/example/a2ui-go-agent-platform/pkg/domain"
	"github.com/example/a2ui-go-agent-platform/pkg/events"
	"github.com/example/a2ui-go-agent-platform/pkg/flow"
	flowruntime "github.com/example/a2ui-go-agent-platform/pkg/flow/runtime"
	"github.com/example/a2ui-go-agent-platform/pkg/guardrail"
	"github.com/example/a2ui-go-agent-platform/pkg/playground/retrieval"
	"github.com/example/a2ui-go-agent-platform/pkg/provider"
	"github.com/example/a2ui-go-agent-platform/pkg/security/policy"
	"github.com/example/a2ui-go-agent-platform/pkg/session"
	"github.com/example/a2ui-go-agent-platform/pkg/store"
	"github.com/example/a2ui-go-agent-platform/pkg/telemetry"
	"github.com/example/a2ui-go-agent-platform/pkg/util"
)

type Service struct {
	provider  *provider.Service
	a2ui      *a2ui.Service
	session   *session.Service
	events    *events.Service
	guardrail *guardrail.Service
	telemetry *telemetry.Service
	versions  store.VersionRepository
	retriever retrieval.Retriever
	flows     *flow.Service
	runtime   *flowruntime.Service
	policy    *policy.Engine
}

func NewService(
	providerSvc *provider.Service,
	a2uiSvc *a2ui.Service,
	sessionSvc *session.Service,
	eventsSvc *events.Service,
	guardrailSvc *guardrail.Service,
	telemetrySvc *telemetry.Service,
	versions store.VersionRepository,
	retriever retrieval.Retriever,
	flowSvc *flow.Service,
	flowRuntime *flowruntime.Service,
) *Service {
	return &Service{
		provider:  providerSvc,
		a2ui:      a2uiSvc,
		session:   sessionSvc,
		events:    eventsSvc,
		guardrail: guardrailSvc,
		telemetry: telemetrySvc,
		versions:  versions,
		retriever: retriever,
		flows:     flowSvc,
		runtime:   flowRuntime,
		policy:    policy.NewEngine(),
	}
}

type RunInput struct {
	SessionID string
	Prompt    string
	OnlyArea  string
	Media     []domain.MultimodalInput
}

type RunOutput struct {
	RunID      string
	VersionID  string
	SchemaJSON string
	Repaired   bool
}

func (s *Service) Run(ctx context.Context, in RunInput) (RunOutput, error) {
	started := time.Now()
	guard := s.guardrail.CheckPromptInjection(in.Prompt)
	if !guard.Allowed {
		s.telemetry.RecordRun(false, time.Since(started).Milliseconds())
		return RunOutput{}, fmt.Errorf("guardrail blocked request: %s", guard.Reason)
	}

	run, err := s.events.StartRun(ctx, in.SessionID, in.Prompt)
	if err != nil {
		s.telemetry.RecordRun(false, time.Since(started).Milliseconds())
		return RunOutput{}, err
	}

	complete := false
	defer func() {
		_ = s.events.CompleteRun(ctx, run.ID, complete)
		s.telemetry.RecordRun(complete, time.Since(started).Milliseconds())
	}()

	for _, media := range in.Media {
		d := s.policy.ValidateMediaRef(media.Ref, []string{"githubusercontent.com", "example-cdn.com"})
		if !d.Allowed {
			_ = s.events.Emit(ctx, run.ID, "media_blocked", map[string]any{"ref": media.Ref, "reason": d.Reason, "rule": d.RuleID}, 0, 0, 0)
			return RunOutput{}, fmt.Errorf("media blocked: %s", d.Reason)
		}
	}

	bundle, err := s.session.BuildContext(ctx, in.SessionID, map[string]any{
		"prompt":    in.Prompt,
		"only_area": in.OnlyArea,
		"media":     in.Media,
	})
	if err != nil {
		_ = s.events.Emit(ctx, run.ID, "init_failed", map[string]any{"error": err.Error()}, 0, 0, 0)
		return RunOutput{}, err
	}
	_ = s.events.Emit(ctx, run.ID, "init", map[string]any{"session": in.SessionID}, 1, 0, 0)

	if s.shouldUseRetrieval() && s.retriever != nil {
		start := time.Now()
		hits, rErr := s.retriever.Retrieve(ctx, retrieval.RetrievalQuery{
			Text:     in.Prompt,
			ThemeID:  fmt.Sprint(bundle.SessionFacts["active_theme_id"]),
			Limit:    retrievalLimit(),
			MinScore: retrievalMinScore(),
		})
		if rErr != nil {
			_ = s.events.Emit(ctx, run.ID, "retrieval_skipped", map[string]any{"reason": rErr.Error()}, int(time.Since(start).Milliseconds()), 0, 0)
		} else if len(hits) == 0 {
			_ = s.events.Emit(ctx, run.ID, "retrieval_skipped", map[string]any{"reason": "no_hits"}, int(time.Since(start).Milliseconds()), 0, 0)
		} else {
			examples := make([]map[string]any, 0, len(hits))
			itemIDs := make([]string, 0, len(hits))
			for _, h := range hits {
				examples = append(examples, map[string]any{"title": h.Title, "tags": h.Tags, "summary": h.Summary})
				itemIDs = append(itemIDs, h.ItemID)
			}
			bundle.TaskInput["examples"] = examples
			_ = s.events.Emit(ctx, run.ID, "retrieval_used", map[string]any{"retrieval_used": true, "retrieval_hits": len(hits), "retrieval_item_ids": itemIDs, "retrieval_latency_ms": int(time.Since(start).Milliseconds())}, int(time.Since(start).Milliseconds()), 0, 0)
		}
	} else {
		_ = s.events.Emit(ctx, run.ID, "retrieval_skipped", map[string]any{"reason": "disabled"}, 0, 0, 0)
	}

	def, _, err := s.flows.ResolveDefinition(ctx, in.SessionID)
	if err != nil {
		_ = s.events.Emit(ctx, run.ID, "flow_resolve_failed", map[string]any{"error": err.Error()}, 0, 0, 0)
		return RunOutput{}, err
	}

	planTask := provider.TaskPlan
	if len(in.Media) > 0 {
		switch in.Media[0].Type {
		case domain.MediaTypeImage:
			planTask = provider.TaskPlanImage
		case domain.MediaTypeAudio:
			planTask = provider.TaskPlanAudio
		}
	}

	flowResult, err := s.runtime.Execute(ctx, def, provider.GenerateRequest{
		SystemPrompt: "generate execution plan",
		UserPrompt:   in.Prompt,
		Context:      map[string]any{"expect": "plan", "bundle": bundle, "plan_task": string(planTask)},
	})
	if err != nil {
		if len(in.Media) > 0 && planTask != provider.TaskPlan {
			flowResult, err = s.runtime.Execute(ctx, def, provider.GenerateRequest{
				SystemPrompt: "generate execution plan",
				UserPrompt:   in.Prompt,
				Context:      map[string]any{"expect": "plan", "bundle": bundle, "plan_task": string(provider.TaskPlan), "degraded_from": string(planTask)},
			})
		}
	}
	if err != nil {
		_ = s.events.Emit(ctx, run.ID, "flow_failed", map[string]any{"error": err.Error()}, 0, 0, 0)
		return RunOutput{}, err
	}
	if err := flowruntime.RequireSuccess(flowResult); err != nil {
		_ = s.events.Emit(ctx, run.ID, "flow_failed", map[string]any{"error": err.Error()}, 0, 0, 0)
		return RunOutput{}, err
	}
	for _, node := range flowResult.Nodes {
		payload := map[string]any{"status": node.Status, "output": node.Output}
		if node.Error != "" {
			payload["error"] = node.Error
		}
		_ = s.events.Emit(ctx, run.ID, "flow."+node.Step, payload, int(node.LatencyMS), 0, node.Tokens)
	}

	latest, err := s.session.GetLatestVersion(ctx, in.SessionID)
	if err != nil {
		return RunOutput{}, err
	}
	baseSchema, err := s.a2ui.ParseSchema(latest.SchemaJSON)
	if err != nil {
		return RunOutput{}, err
	}

	patch := domain.PatchDocument{
		Mode:                "patch",
		TargetSchemaVersion: latest.ID,
		Operations: []domain.PatchOperation{
			{
				Op:     "update_root_props",
				Target: "root",
				Value: map[string]any{
					"title": in.Prompt,
				},
				Reason: "sync prompt to root title",
			},
		},
	}
	next, err := s.a2ui.ApplyPatch(baseSchema, patch)
	if err != nil {
		return RunOutput{}, err
	}
	validated := s.a2ui.ValidateSchema(next)
	repaired := false
	if !validated.Valid {
		errorJSON, _ := json.Marshal(validated.Errors)
		repairPrompt := fmt.Sprintf("repair schema errors: %s", string(errorJSON))
		repairResp, rErr := s.provider.Generate(ctx, provider.TaskRepair, provider.GenerateRequest{
			SystemPrompt: "repair ui schema",
			UserPrompt:   repairPrompt,
			Context:      map[string]any{"expect": "repair"},
		})
		if rErr != nil {
			return RunOutput{}, fmt.Errorf("repair failed: %w", rErr)
		}
		_ = s.events.Emit(ctx, run.ID, "repair", map[string]any{"text": repairResp.Text}, 8, 50, repairResp.Tokens)
		next.Root.Type = "Container"
		repaired = true
		validated = s.a2ui.ValidateSchema(next)
		if !validated.Valid {
			errorJSON, _ := json.Marshal(validated.Errors)
			_ = s.events.Emit(ctx, run.ID, "repair_failed", map[string]any{"validation_errors": string(errorJSON)}, 0, 0, 0)
			return RunOutput{}, fmt.Errorf("schema validation failed after repair: %s", string(errorJSON))
		}
	}

	outJSON, err := json.Marshal(next)
	if err != nil {
		return RunOutput{}, err
	}
	newVersion := domain.SchemaVersion{
		ID:              util.NewID("ver"),
		SessionID:       in.SessionID,
		ParentVersionID: latest.ID,
		SchemaPath:      latest.SchemaPath,
		SchemaHash:      "",
		SchemaJSON:      string(outJSON),
		Summary:         "agent applied patch",
		ThemeSnapshotID: latest.ThemeSnapshotID,
		CreatedAt:       time.Now(),
	}
	if err := s.versions.CreateVersion(ctx, newVersion); err != nil {
		return RunOutput{}, err
	}
	for _, media := range in.Media {
		metadataJSON, _ := json.Marshal(media.Metadata)
		_ = s.versions.CreateVersionAsset(ctx, domain.SchemaVersionAsset{
			ID:           util.NewID("asset"),
			VersionID:    newVersion.ID,
			AssetType:    string(media.Type),
			AssetRef:     media.Ref,
			MetadataJSON: string(metadataJSON),
			CreatedAt:    time.Now(),
		})
	}
	_ = s.events.Emit(ctx, run.ID, "apply", map[string]any{"version_id": newVersion.ID}, 5, 0, 0)

	complete = true
	return RunOutput{
		RunID:      run.ID,
		VersionID:  newVersion.ID,
		SchemaJSON: string(outJSON),
		Repaired:   repaired,
	}, nil
}

func (s *Service) shouldUseRetrieval() bool {
	v := os.Getenv("PLAYGROUND_RETRIEVAL_ENABLED")
	if v == "" {
		return true
	}
	return v == "true" || v == "1"
}

func retrievalLimit() int {
	v := os.Getenv("PLAYGROUND_RETRIEVAL_LIMIT")
	if v == "" {
		return 3
	}
	n, err := strconv.Atoi(v)
	if err != nil || n <= 0 {
		return 3
	}
	return n
}

func retrievalMinScore() float64 {
	v := os.Getenv("PLAYGROUND_RETRIEVAL_MIN_SCORE")
	if v == "" {
		return 1
	}
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 1
	}
	return f
}
