package db

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
}

func LoadConfigFromEnv() (Config, error) {
	port := 5432
	if raw := os.Getenv("PG_PORT"); raw != "" {
		p, err := strconv.Atoi(raw)
		if err != nil {
			return Config{}, fmt.Errorf("invalid PG_PORT: %w", err)
		}
		port = p
	}

	cfg := Config{
		Host:     envOrDefault("PG_HOST", "localhost"),
		Port:     port,
		User:     envOrDefault("PG_USER", "postgres"),
		Password: envOrDefault("PG_PASSWORD", "postgres"),
		Database: envOrDefault("PG_DATABASE", "a2ui_platform"),
		SSLMode:  envOrDefault("PG_SSLMODE", "disable"),
	}
	if err := validateDatabaseName(cfg.Database); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (c Config) AppDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", c.Host, c.Port, c.User, c.Password, c.Database, c.SSLMode)
}

func (c Config) AdminDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=postgres sslmode=%s", c.Host, c.Port, c.User, c.Password, c.SSLMode)
}

func envOrDefault(name, defaultValue string) string {
	v := os.Getenv(name)
	if v == "" {
		return defaultValue
	}
	return v
}

func validateDatabaseName(db string) error {
	if db == "" {
		return fmt.Errorf("database name is required")
	}
	ok, _ := regexp.MatchString(`^[a-zA-Z_][a-zA-Z0-9_]*$`, db)
	if !ok {
		return fmt.Errorf("invalid database name: %s", db)
	}
	return nil
}
