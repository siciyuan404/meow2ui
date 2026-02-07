package injection

import "strings"

type Result struct {
	Detected      bool
	Severity      string
	Patterns      []string
	SanitizedText string
}

func DetectPrompt(input string) Result {
	l := strings.ToLower(input)
	patterns := []string{}
	if strings.Contains(l, "ignore previous instructions") {
		patterns = append(patterns, "ignore_previous_instructions")
	}
	if strings.Contains(l, "reveal system prompt") {
		patterns = append(patterns, "reveal_system_prompt")
	}
	if strings.Contains(l, "delete all files") {
		patterns = append(patterns, "delete_all_files")
	}
	if len(patterns) == 0 {
		return Result{Detected: false, Severity: "none", SanitizedText: input}
	}
	clean := strings.ReplaceAll(input, "ignore previous instructions", "[filtered]")
	clean = strings.ReplaceAll(clean, "reveal system prompt", "[filtered]")
	return Result{Detected: true, Severity: "high", Patterns: patterns, SanitizedText: clean}
}
