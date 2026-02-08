package flow

import (
	"context"
	"testing"

	"github.com/example/a2ui-go-agent-platform/internal/infra/memorystore"
)

func TestFlowTemplateCreateAndResolve(t *testing.T) {
	ctx := context.Background()
	st := memorystore.New()
	svc := NewService(st.Flow())

	def := DefaultDefinition()
	tpl, ver, err := svc.CreateTemplate(ctx, "default", "v1", def)
	if err != nil {
		t.Fatalf("create template: %v", err)
	}
	if tpl.ID == "" || ver.ID == "" {
		t.Fatalf("expected ids")
	}

	if err := svc.BindSession(ctx, "ssn-1", tpl.ID, ver.Version); err != nil {
		t.Fatalf("bind: %v", err)
	}

	resolved, bound, err := svc.ResolveDefinition(ctx, "ssn-1")
	if err != nil {
		t.Fatalf("resolve: %v", err)
	}
	if !bound {
		t.Fatalf("expected bound definition")
	}
	if resolved.Name != def.Name {
		t.Fatalf("expected %s, got %s", def.Name, resolved.Name)
	}
}
