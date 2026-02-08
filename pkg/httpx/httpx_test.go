package httpx

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/example/a2ui-go-agent-platform/pkg/store"
)

func TestWriteJSON(t *testing.T) {
	w := httptest.NewRecorder()
	WriteJSON(w, http.StatusOK, map[string]string{"hello": "world"})
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if ct := w.Header().Get("Content-Type"); ct != "application/json" {
		t.Fatalf("expected application/json, got %s", ct)
	}
}

func TestWriteError_MethodNotAllowed(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/test", nil)
	WriteError(w, r, ErrMethodNotAllowed)
	if w.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", w.Code)
	}
}

func TestWriteError_NotFound(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/test", nil)
	WriteError(w, r, store.ErrNotFound)
	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}

func TestWriteError_Conflict(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/test", nil)
	WriteError(w, r, store.ErrConflict)
	if w.Code != http.StatusConflict {
		t.Fatalf("expected 409, got %d", w.Code)
	}
}

func TestWriteError_Generic(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/test", nil)
	WriteError(w, r, errors.New("something broke"))
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestTraceMiddleware(t *testing.T) {
	handler := TraceMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := TraceIDFromContext(r.Context())
		if traceID == "" {
			t.Fatal("expected trace ID in context")
		}
		WriteJSON(w, http.StatusOK, map[string]string{"traceId": traceID})
	}))

	// Without trace header
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/test", nil)
	handler.ServeHTTP(w, r)
	if w.Header().Get("X-Trace-Id") == "" {
		t.Fatal("expected X-Trace-Id header")
	}

	// With trace header
	w2 := httptest.NewRecorder()
	r2 := httptest.NewRequest("GET", "/test", nil)
	r2.Header.Set("X-Trace-Id", "custom-trace-123")
	handler.ServeHTTP(w2, r2)
	if w2.Header().Get("X-Trace-Id") != "custom-trace-123" {
		t.Fatalf("expected custom-trace-123, got %s", w2.Header().Get("X-Trace-Id"))
	}
}

func TestTraceIDFromContext_Empty(t *testing.T) {
	r := httptest.NewRequest("GET", "/test", nil)
	traceID := TraceIDFromContext(r.Context())
	if traceID == "" {
		t.Fatal("expected generated trace ID")
	}
}

func TestReadyPayload(t *testing.T) {
	p := ReadyPayload(true, true, "v1.0.0")
	if p["ok"] != true {
		t.Fatal("expected ok=true")
	}
	if p["version"] != "v1.0.0" {
		t.Fatalf("expected v1.0.0, got %v", p["version"])
	}

	p2 := ReadyPayload(true, false, "v1.0.0")
	if p2["ok"] != false {
		t.Fatal("expected ok=false when db is down")
	}
}
