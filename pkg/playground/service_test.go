package playground

import (
	"context"
	"testing"

	"github.com/example/a2ui-go-agent-platform/internal/infra/memorystore"
)

func TestPlaygroundSaveAndSearch(t *testing.T) {
	ctx := context.Background()
	st := memorystore.New()
	svc := NewService(st)

	cat, err := svc.CreateCategory(ctx, "dashboard")
	if err != nil {
		t.Fatalf("create category failed: %v", err)
	}
	_, _ = svc.CreateTag(ctx, "finance")

	_, err = svc.SaveItem(ctx, SaveInput{
		Title:          "Finance Board",
		CategoryID:     cat.ID,
		SchemaSnapshot: `{"version":"1.0.0"}`,
		Tags:           []string{"finance"},
	})
	if err != nil {
		t.Fatalf("save item failed: %v", err)
	}

	items, err := svc.Search(ctx, cat.ID, []string{"finance"}, "board")
	if err != nil {
		t.Fatalf("search failed: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
}
