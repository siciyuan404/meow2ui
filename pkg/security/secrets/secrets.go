package secrets

import (
	"context"
	"fmt"
	"os"
	"strings"
)

type Provider interface {
	Get(context.Context, string) (string, error)
}

type EnvProvider struct{}

func (p EnvProvider) Get(_ context.Context, ref string) (string, error) {
	if strings.TrimSpace(ref) == "" {
		return "", fmt.Errorf("secret ref is empty")
	}
	v := os.Getenv(ref)
	if strings.TrimSpace(v) == "" {
		return "", fmt.Errorf("secret not found for ref: %s", ref)
	}
	return v, nil
}
