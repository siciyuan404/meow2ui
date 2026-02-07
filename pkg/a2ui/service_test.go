package a2ui

import (
	"testing"

	"github.com/example/a2ui-go-agent-platform/pkg/domain"
)

func TestValidateSchema(t *testing.T) {
	svc := NewService()
	ok := svc.ValidateSchema(domain.UISchema{
		Version: "1.0.0",
		Root: domain.UIComponent{
			ID:   "root",
			Type: "Container",
		},
	})
	if !ok.Valid {
		t.Fatalf("expected valid schema, got errors=%v", ok.Errors)
	}

	bad := svc.ValidateSchema(domain.UISchema{})
	if bad.Valid {
		t.Fatalf("expected invalid schema")
	}
	if len(bad.Errors) == 0 || bad.Errors[0].Code == "" {
		t.Fatalf("expected structured validation errors")
	}
}

func TestApplyPatch(t *testing.T) {
	svc := NewService()
	base := domain.UISchema{Version: "1.0.0", Root: domain.UIComponent{ID: "root", Type: "Container"}}
	patch := domain.PatchDocument{
		Mode: "patch",
		Operations: []domain.PatchOperation{
			{Op: "update_root_props", Target: "root", Value: map[string]any{"title": "hello"}},
		},
	}
	next, err := svc.ApplyPatch(base, patch)
	if err != nil {
		t.Fatalf("apply patch failed: %v", err)
	}
	if next.Root.Props["title"] != "hello" {
		t.Fatalf("expected title updated")
	}
}

func TestValidateSchema_ComponentAndPropRules(t *testing.T) {
	svc := NewService()
	res := svc.ValidateSchema(domain.UISchema{
		Version: "1.0.0",
		Root: domain.UIComponent{
			ID:   "root",
			Type: "Unknown",
		},
	})
	if res.Valid {
		t.Fatalf("expected invalid due to unknown component")
	}

	res2 := svc.ValidateSchema(domain.UISchema{
		Version: "1.0.0",
		Root: domain.UIComponent{
			ID:   "root",
			Type: "Container",
			Props: map[string]any{
				"title": 100,
			},
		},
	})
	if res2.Valid {
		t.Fatalf("expected invalid due to prop type mismatch")
	}
}
