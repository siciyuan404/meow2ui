package store

import (
	"testing"
)

func TestErrNotFound(t *testing.T) {
	if ErrNotFound == nil {
		t.Fatal("expected non-nil ErrNotFound")
	}
	if ErrNotFound.Error() != "not found" {
		t.Fatalf("expected 'not found', got %q", ErrNotFound.Error())
	}
}

func TestErrConflict(t *testing.T) {
	if ErrConflict == nil {
		t.Fatal("expected non-nil ErrConflict")
	}
	if ErrConflict.Error() != "conflict" {
		t.Fatalf("expected 'conflict', got %q", ErrConflict.Error())
	}
}

func TestErrUnavailable(t *testing.T) {
	if ErrUnavailable == nil {
		t.Fatal("expected non-nil ErrUnavailable")
	}
	if ErrUnavailable.Error() != "unavailable" {
		t.Fatalf("expected 'unavailable', got %q", ErrUnavailable.Error())
	}
}

func TestErrorsAreDistinct(t *testing.T) {
	if ErrNotFound == ErrConflict {
		t.Fatal("ErrNotFound and ErrConflict should be distinct")
	}
	if ErrNotFound == ErrUnavailable {
		t.Fatal("ErrNotFound and ErrUnavailable should be distinct")
	}
	if ErrConflict == ErrUnavailable {
		t.Fatal("ErrConflict and ErrUnavailable should be distinct")
	}
}
