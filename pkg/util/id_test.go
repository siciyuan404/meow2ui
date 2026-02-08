package util

import (
	"strings"
	"testing"
)

func TestNewID(t *testing.T) {
	id := NewID("test")
	if !strings.HasPrefix(id, "test-") {
		t.Fatalf("expected prefix 'test-', got %s", id)
	}
}

func TestNewID_Unique(t *testing.T) {
	ids := map[string]bool{}
	for i := 0; i < 100; i++ {
		id := NewID("u")
		if ids[id] {
			t.Fatalf("duplicate ID: %s", id)
		}
		ids[id] = true
	}
}
