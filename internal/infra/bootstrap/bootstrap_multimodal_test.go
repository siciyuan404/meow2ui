package bootstrap_test

import (
	"context"
	"testing"

	"github.com/example/a2ui-go-agent-platform/internal/infra/bootstrap"
	"github.com/example/a2ui-go-agent-platform/pkg/agent"
	"github.com/example/a2ui-go-agent-platform/pkg/domain"
)

func TestAgentRunRejectsBlockedMediaRef(t *testing.T) {
	ctx := context.Background()
	app, err := bootstrap.New(ctx)
	if err != nil {
		t.Fatalf("bootstrap: %v", err)
	}
	ws, err := app.Workspace.Create(ctx, "demo", "/tmp/demo")
	if err != nil {
		t.Fatalf("workspace: %v", err)
	}
	ssn, _, err := app.Session.Create(ctx, ws.ID, "s1", "default")
	if err != nil {
		t.Fatalf("session: %v", err)
	}

	_, err = app.Agent.Run(ctx, agent.RunInput{
		SessionID: ssn.ID,
		Prompt:    "build page",
		Media: []domain.MultimodalInput{{
			Type: domain.MediaTypeImage,
			Ref:  "http://127.0.0.1/private.png",
		}},
	})
	if err == nil {
		t.Fatalf("expected blocked media error")
	}
}
