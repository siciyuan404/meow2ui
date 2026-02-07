package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/example/a2ui-go-agent-platform/pkg/domain"
	"github.com/example/a2ui-go-agent-platform/pkg/security/secrets"
)

type OpenAICompatibleAdapter struct{}

func (a OpenAICompatibleAdapter) Name() string { return "openai_compatible" }

func (a OpenAICompatibleAdapter) Generate(ctx context.Context, providerCfg domain.Provider, model domain.Model, req GenerateRequest) (GenerateResponse, error) {
	secretProvider := secrets.EnvProvider{}
	key := ""
	if strings.TrimSpace(providerCfg.AuthRef) != "" {
		v, err := secretProvider.Get(ctx, providerCfg.AuthRef)
		if err == nil {
			key = v
		}
	}
	if strings.TrimSpace(key) == "" {
		return GenerateResponse{}, &ProviderError{Code: "PROVIDER_AUTH_ERROR", ProviderID: providerCfg.ID, ModelID: model.ID, Retryable: false, Message: "missing api key from auth_ref"}
	}

	timeout := time.Duration(providerCfg.TimeoutMS) * time.Millisecond
	if timeout <= 0 {
		timeout = 30 * time.Second
	}
	client := &http.Client{Timeout: timeout}

	body := map[string]any{
		"model": model.Name,
		"messages": []map[string]string{
			{"role": "system", "content": req.SystemPrompt},
			{"role": "user", "content": req.UserPrompt},
		},
		"temperature": 0.2,
	}
	b, _ := json.Marshal(body)

	url := strings.TrimRight(providerCfg.BaseURL, "/") + "/chat/completions"
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return GenerateResponse{}, err
	}
	httpReq.Header.Set("Authorization", "Bearer "+key)
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(httpReq)
	if err != nil {
		return GenerateResponse{}, &ProviderError{Code: "PROVIDER_TIMEOUT", ProviderID: providerCfg.ID, ModelID: model.ID, Retryable: true, Message: err.Error()}
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		return GenerateResponse{}, &ProviderError{Code: "PROVIDER_AUTH_ERROR", ProviderID: providerCfg.ID, ModelID: model.ID, Retryable: false, Message: fmt.Sprintf("upstream status: %d", resp.StatusCode)}
	}
	if resp.StatusCode == http.StatusTooManyRequests {
		return GenerateResponse{}, &ProviderError{Code: "PROVIDER_RATE_LIMIT", ProviderID: providerCfg.ID, ModelID: model.ID, Retryable: true, Message: "rate limited"}
	}
	if resp.StatusCode >= 500 {
		return GenerateResponse{}, &ProviderError{Code: "PROVIDER_UPSTREAM_ERROR", ProviderID: providerCfg.ID, ModelID: model.ID, Retryable: true, Message: fmt.Sprintf("upstream status: %d", resp.StatusCode)}
	}
	if resp.StatusCode >= 400 {
		return GenerateResponse{}, &ProviderError{Code: "PROVIDER_UPSTREAM_ERROR", ProviderID: providerCfg.ID, ModelID: model.ID, Retryable: false, Message: fmt.Sprintf("upstream status: %d", resp.StatusCode)}
	}

	var payload struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
		Usage struct {
			TotalTokens int `json:"total_tokens"`
		} `json:"usage"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return GenerateResponse{}, err
	}
	if len(payload.Choices) == 0 {
		return GenerateResponse{}, &ProviderError{Code: "PROVIDER_UPSTREAM_ERROR", ProviderID: providerCfg.ID, ModelID: model.ID, Retryable: false, Message: "empty choices"}
	}
	return GenerateResponse{Text: payload.Choices[0].Message.Content, Tokens: payload.Usage.TotalTokens}, nil
}
