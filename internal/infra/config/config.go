package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/example/a2ui-go-agent-platform/internal/infra/db"
)

type AppConfig struct {
	StoreDriver string
	ServerAddr  string
	AutoMigrate bool
	Postgres    db.Config
}

func Load() (AppConfig, error) {
	storeDriver := strings.TrimSpace(os.Getenv("STORE_DRIVER"))
	if storeDriver == "" {
		storeDriver = "memory"
	}
	addr := strings.TrimSpace(os.Getenv("A2UI_ADDR"))
	if addr == "" {
		addr = ":8080"
	}
	autoMigrate := strings.EqualFold(strings.TrimSpace(os.Getenv("AUTO_MIGRATE")), "true")

	cfg := AppConfig{StoreDriver: storeDriver, ServerAddr: addr, AutoMigrate: autoMigrate}
	if storeDriver == "postgres" {
		pg, err := db.LoadConfigFromEnv()
		if err != nil {
			return AppConfig{}, err
		}
		cfg.Postgres = pg
	}
	if err := cfg.Validate(); err != nil {
		return AppConfig{}, err
	}
	return cfg, nil
}

func (c AppConfig) Validate() error {
	if c.StoreDriver != "memory" && c.StoreDriver != "postgres" {
		return fmt.Errorf("invalid STORE_DRIVER: %s", c.StoreDriver)
	}
	if strings.TrimSpace(c.ServerAddr) == "" {
		return fmt.Errorf("A2UI_ADDR is required")
	}
	if c.StoreDriver == "postgres" {
		if strings.TrimSpace(c.Postgres.Host) == "" || strings.TrimSpace(c.Postgres.User) == "" || strings.TrimSpace(c.Postgres.Database) == "" {
			return fmt.Errorf("invalid postgres config")
		}
	}
	return nil
}
