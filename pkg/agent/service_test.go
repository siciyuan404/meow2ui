package agent

import (
	"testing"

	"github.com/example/a2ui-go-agent-platform/pkg/domain"
)

func TestNewService(t *testing.T) {
	svc := NewService(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	if svc == nil {
		t.Fatal("expected non-nil service")
	}
	if svc.provider != nil {
		t.Fatal("expected nil provider")
	}
}

func TestRunInput_Fields(t *testing.T) {
	in := RunInput{
		SessionID: "ssn-1",
		Prompt:    "create a dashboard",
		OnlyArea:  "header",
	}
	if in.SessionID != "ssn-1" {
		t.Fatalf("expected ssn-1, got %s", in.SessionID)
	}
	if in.Prompt != "create a dashboard" {
		t.Fatalf("unexpected prompt: %s", in.Prompt)
	}
	if in.OnlyArea != "header" {
		t.Fatalf("unexpected onlyArea: %s", in.OnlyArea)
	}
	if len(in.Media) != 0 {
		t.Fatalf("expected empty media, got %d", len(in.Media))
	}
}

func TestRunOutput_Fields(t *testing.T) {
	out := RunOutput{
		RunID:      "run-1",
		VersionID:  "ver-1",
		SchemaJSON: `{"version":"1.0.0"}`,
		Repaired:   true,
	}
	if out.RunID != "run-1" {
		t.Fatalf("expected run-1, got %s", out.RunID)
	}
	if !out.Repaired {
		t.Fatal("expected repaired=true")
	}
}

func TestShouldUseRetrieval_Default(t *testing.T) {
	svc := NewService(nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	// Default (no env var) should return true
	if !svc.shouldUseRetrieval() {
		t.Fatal("expected retrieval enabled by default")
	}
}

func TestRetrievalLimit_Default(t *testing.T) {
	limit := retrievalLimit()
	if limit != 3 {
		t.Fatalf("expected default limit 3, got %d", limit)
	}
}

func TestRetrievalMinScore_Default(t *testing.T) {
	score := retrievalMinScore()
	if score != 1 {
		t.Fatalf("expected default min score 1, got %f", score)
	}
}

func TestRunInput_WithMedia(t *testing.T) {
	tests := []struct {
		name      string
		input     RunInput
		mediaLen  int
		sessionID string
	}{
		{
			name:      "no media",
			input:     RunInput{SessionID: "s1", Prompt: "hello"},
			mediaLen:  0,
			sessionID: "s1",
		},
		{
			name: "with image media",
			input: RunInput{
				SessionID: "s2",
				Prompt:    "analyze this",
				Media:     []domain.MultimodalInput{},
			},
			mediaLen:  0,
			sessionID: "s2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.input.SessionID != tt.sessionID {
				t.Fatalf("expected %s, got %s", tt.sessionID, tt.input.SessionID)
			}
			if len(tt.input.Media) != tt.mediaLen {
				t.Fatalf("expected %d media, got %d", tt.mediaLen, len(tt.input.Media))
			}
		})
	}
}
