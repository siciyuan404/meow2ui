package redaction

import (
	"fmt"
	"strings"
)

var sensitiveKeys = []string{"api_key", "token", "secret", "authorization", "password", "credential"}

func RedactMap(input map[string]any) map[string]any {
	if input == nil {
		return nil
	}
	out := make(map[string]any, len(input))
	for k, v := range input {
		if isSensitiveKey(k) {
			out[k] = "***REDACTED***"
			continue
		}
		switch tv := v.(type) {
		case map[string]any:
			out[k] = RedactMap(tv)
		case []any:
			items := make([]any, 0, len(tv))
			for _, item := range tv {
				if m, ok := item.(map[string]any); ok {
					items = append(items, RedactMap(m))
					continue
				}
				items = append(items, redactValue(item))
			}
			out[k] = items
		default:
			out[k] = redactValue(v)
		}
	}
	return out
}

func redactValue(v any) any {
	s := strings.TrimSpace(fmt.Sprint(v))
	l := strings.ToLower(s)
	if strings.HasPrefix(l, "bearer ") {
		return "Bearer ***REDACTED***"
	}
	if len(s) >= 24 && looksLikeSecret(s) {
		return "***REDACTED***"
	}
	return v
}

func isSensitiveKey(k string) bool {
	lower := strings.ToLower(strings.TrimSpace(k))
	for _, token := range sensitiveKeys {
		if strings.Contains(lower, token) {
			return true
		}
	}
	return false
}

func looksLikeSecret(v string) bool {
	hasDigit := false
	hasLetter := false
	for _, ch := range v {
		if ch >= '0' && ch <= '9' {
			hasDigit = true
		}
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			hasLetter = true
		}
	}
	return hasDigit && hasLetter
}
