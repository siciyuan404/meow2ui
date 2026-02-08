package workspace

import (
	"context"
	"testing"

	"github.com/example/a2ui-go-agent-platform/internal/infra/memorystore"
)

func TestService_Create_Success(t *testing.T) {
	ms := memorystore.New()
	svc := NewService(ms)

	ws, err := svc.Create(context.Background(), "My Workspace", "/tmp/ws")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ws.Name != "My Workspace" {
		t.Fatalf("expected My Workspace, got %s", ws.Name)
	}
}

func TestService_Create_EmptyName(t *testing.T) {
	ms := memorystore.New()
	svc := NewService(ms)
	_, err := svc.Create(context.Background(), "", "/tmp/ws")
	if err == nil {
		t.Fatal("expected error for empty name")
	}
}

func TestService_Create_EmptyRoot(t *testing.T) {
	ms := memorystore.New()
	svc := NewService(ms)
	_, err := svc.Create(context.Background(), "Test", "")
	if err == nil {
		t.Fatal("expected error for empty root")
	}
}

func TestService_List(t *testing.T) {
	ms := memorystore.New()
	svc := NewService(ms)
	svc.Create(context.Background(), "WS1", "/tmp/1")
	svc.Create(context.Background(), "WS2", "/tmp/2")

	list, err := svc.List(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(list) != 2 {
		t.Fatalf("expected 2, got %d", len(list))
	}
}

func TestService_Get(t *testing.T) {
	ms := memorystore.New()
	svc := NewService(ms)
	ws, _ := svc.Create(context.Background(), "Test", "/tmp/test")

	got, err := svc.Get(context.Background(), ws.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Name != "Test" {
		t.Fatalf("expected Test, got %s", got.Name)
	}

	_, err = svc.Get(context.Background(), "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent workspace")
	}
}
