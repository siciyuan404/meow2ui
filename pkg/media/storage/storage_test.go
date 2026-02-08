package storage

import (
	"context"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestLocalBackendRoundTrip(t *testing.T) {
	root := t.TempDir()
	b := LocalBackend{Root: root}
	ctx := context.Background()

	ref := filepath.ToSlash("images/demo.txt")
	if err := b.Put(ctx, Item{Ref: ref, Data: []byte("hello")}); err != nil {
		t.Fatalf("put: %v", err)
	}
	item, err := b.Get(ctx, ref)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if string(item.Data) != "hello" {
		t.Fatalf("unexpected data: %s", string(item.Data))
	}
	url, err := b.SignURL(ctx, ref, time.Minute)
	if err != nil || url == "" {
		t.Fatalf("sign url failed: %v", err)
	}
	if err := b.Delete(ctx, ref); err != nil {
		t.Fatalf("delete: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, ref)); err == nil {
		t.Fatalf("file should be deleted")
	}
}
