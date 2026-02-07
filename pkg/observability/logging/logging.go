package logging

import (
	"encoding/json"
	"log"
	"time"
)

type Entry struct {
	Timestamp string `json:"timestamp"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	TraceID   string `json:"trace_id,omitempty"`
	RunID     string `json:"run_id,omitempty"`
	SessionID string `json:"session_id,omitempty"`
	Code      string `json:"code,omitempty"`
	Component string `json:"component,omitempty"`
}

func Log(level, message, traceID, runID, sessionID, code, component string) {
	e := Entry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Level:     level,
		Message:   message,
		TraceID:   traceID,
		RunID:     runID,
		SessionID: sessionID,
		Code:      code,
		Component: component,
	}
	b, _ := json.Marshal(e)
	log.Print(string(b))
}
