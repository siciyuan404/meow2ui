package httpx

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/store"
	"github.com/example/a2ui-go-agent-platform/pkg/util"
)

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
	TraceID string `json:"traceId"`
}

var ErrMethodNotAllowed = errors.New("method not allowed")

type ctxKey string

const traceIDKey ctxKey = "trace_id"

func TraceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get("X-Trace-Id")
		if traceID == "" {
			traceID = util.NewID("trace")
		}
		w.Header().Set("X-Trace-Id", traceID)
		ctx := context.WithValue(r.Context(), traceIDKey, traceID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func TraceIDFromContext(ctx context.Context) string {
	v, _ := ctx.Value(traceIDKey).(string)
	if v == "" {
		return util.NewID("trace")
	}
	return v
}

func WriteJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, r *http.Request, err error) {
	status := http.StatusInternalServerError
	code := "INTERNAL_ERROR"
	message := "internal error"

	switch {
	case errors.Is(err, ErrMethodNotAllowed):
		status, code, message = http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "method not allowed"
	case errors.Is(err, store.ErrNotFound):
		status, code, message = http.StatusNotFound, "NOT_FOUND", "resource not found"
	case errors.Is(err, store.ErrConflict):
		status, code, message = http.StatusConflict, "CONFLICT", "resource conflict"
	default:
		status, code, message = http.StatusBadRequest, "BAD_REQUEST", err.Error()
	}

	WriteJSON(w, status, APIError{
		Code:    code,
		Message: message,
		Detail:  err.Error(),
		TraceID: TraceIDFromContext(r.Context()),
	})
}

func ReadyPayload(appOK bool, dbOK bool, version string) map[string]any {
	return map[string]any{
		"ok":      appOK && dbOK,
		"version": version,
		"time":    time.Now().UTC().Format(time.RFC3339),
		"checks": map[string]any{
			"app": map[string]any{"ok": appOK},
			"db":  map[string]any{"ok": dbOK},
		},
	}
}
