package runtime

import (
	"context"
	"testing"

	"github.com/example/a2ui-go-agent-platform/internal/infra/memorystore"
	"github.com/example/a2ui-go-agent-platform/pkg/flow"
	"github.com/example/a2ui-go-agent-platform/pkg/provider"
)

func TestExecuteDefaultFlow(t *testing.T) {
	ctx := context.Background()
	st := memorystore.New()
	router := provider.NewRouter(st.Provider())
	registry := provider.NewRegistry()
	registry.Register("mock", provider.MockAdapter{})
	providerSvc := provider.NewService(router, registry)

	rt := NewService(providerSvc)
	result, err := rt.Execute(ctx, flow.DefaultDefinition(), provider.GenerateRequest{SystemPrompt: "test", UserPrompt: "hello"})
	if err == nil {
		t.Fatalf("expected error when no model configured")
	}
	if result.Success {
		t.Fatalf("expected failed result")
	}
}
