package provider

import (
	"context"
	"encoding/json"

	"github.com/example/a2ui-go-agent-platform/pkg/domain"
)

type MockAdapter struct{}

func (m MockAdapter) Name() string { return "mock" }

func (m MockAdapter) Generate(_ context.Context, _ domain.Provider, _ domain.Model, req GenerateRequest) (GenerateResponse, error) {
	if req.Context != nil {
		if mode, ok := req.Context["expect"]; ok {
			s, _ := mode.(string)
			switch s {
			case "plan":
				b, _ := json.Marshal(map[string]any{
					"goal":    "update ui",
					"changes": []string{"modify title", "add card"},
				})
				return GenerateResponse{Text: string(b), Tokens: 120}, nil
			case "repair":
				return GenerateResponse{Text: req.UserPrompt, Tokens: 80}, nil
			}
		}
	}
	return GenerateResponse{Text: req.UserPrompt, Tokens: 100}, nil
}
