package marketplace

import (
	"context"
	"testing"
)

func TestMarketplaceLifecycle(t *testing.T) {
	ctx := context.Background()
	svc := NewService()

	tpl, err := svc.SaveFromSession(ctx, SaveFromSessionInput{Name: "Dashboard", Category: "Dashboard", Tags: []string{"React"}, Schema: "{}", Theme: "default", Owner: "u1"})
	if err != nil {
		t.Fatalf("save: %v", err)
	}
	if err := svc.SubmitReview(ctx, tpl.ID); err != nil {
		t.Fatalf("submit: %v", err)
	}
	if err := svc.Review(ctx, tpl.ID, "publish", "ok"); err != nil {
		t.Fatalf("review: %v", err)
	}
	if _, err := svc.AddVersion(ctx, tpl.ID, "{\"v\":2}", "improve cards"); err != nil {
		t.Fatalf("version: %v", err)
	}
	if _, err := svc.AddRating(ctx, tpl.ID, "u2", 5, "great"); err != nil {
		t.Fatalf("rating: %v", err)
	}
	if err := svc.FlagComment(ctx, tpl.ID, "u2"); err != nil {
		t.Fatalf("flag: %v", err)
	}
	items := svc.Search(ctx, SearchInput{Category: "Dashboard", Tag: "React", Query: "dash"})
	if len(items) == 0 {
		t.Fatalf("expected search result")
	}
	if _, err := svc.ApplyTemplate(ctx, tpl.ID, "ssn-1"); err != nil {
		t.Fatalf("apply: %v", err)
	}
}
