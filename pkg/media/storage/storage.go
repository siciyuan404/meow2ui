package storage

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Item struct {
	Ref      string
	Data     []byte
	Metadata map[string]any
}

type Backend interface {
	Put(context.Context, Item) error
	Get(context.Context, string) (Item, error)
	Delete(context.Context, string) error
	SignURL(context.Context, string, time.Duration) (string, error)
}

type LocalBackend struct {
	Root string
}

func (b LocalBackend) Put(_ context.Context, item Item) error {
	if item.Ref == "" {
		return fmt.Errorf("ref is required")
	}
	path := filepath.Join(b.Root, filepath.Clean(item.Ref))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, item.Data, 0o644)
}

func (b LocalBackend) Get(_ context.Context, ref string) (Item, error) {
	path := filepath.Join(b.Root, filepath.Clean(ref))
	data, err := os.ReadFile(path)
	if err != nil {
		return Item{}, err
	}
	return Item{Ref: ref, Data: data}, nil
}

func (b LocalBackend) Delete(_ context.Context, ref string) error {
	path := filepath.Join(b.Root, filepath.Clean(ref))
	return os.Remove(path)
}

func (b LocalBackend) SignURL(_ context.Context, ref string, _ time.Duration) (string, error) {
	if ref == "" {
		return "", fmt.Errorf("ref is required")
	}
	return "local://" + ref, nil
}
