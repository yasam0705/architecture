package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	App            string
	Environment    string
	LogLevel       string
	RpcPort        string
	ContextTimeout time.Duration
	Postgres       struct {
		Host     string
		Port     string
		Database string
		User     string
		Password string
		Sslmode  string
	}
}

func New() (*Config, error) {
	cfg := &Config{}

	cfg.App = getEnv("APP", "service")
	cfg.Environment = getEnv("ENVIRONMENT", "develop")
	cfg.LogLevel = getEnv("LOG_LEVEL", "debug")
	cfg.RpcPort = getEnv("RPC_PORT", ":80")
	contextTimeout, err := time.ParseDuration(getEnv("CONTEXT_TIMEOUT", "10s"))
	if err != nil {
		return nil, fmt.Errorf("error parse duration context timeout: %s", err.Error())
	}
	cfg.ContextTimeout = contextTimeout

	cfg.Postgres.Host = getEnv("POSTGRES_HOST", "localhost")
	cfg.Postgres.Port = getEnv("POSTGRES_PORT", "5432")
	cfg.Postgres.Database = getEnv("POSTGRES_DATABASE", "db")
	cfg.Postgres.User = getEnv("POSTGRES_USER", "user")
	cfg.Postgres.Password = getEnv("POSTGRES_PASSWORD", "password")
	cfg.Postgres.Sslmode = getEnv("POSTGRES_SSLMODE", "disable")

	return cfg, nil
}

func getEnv(key, value string) string {
	v, ok := os.LookupEnv(key)
	if ok {
		return v
	}
	return value
}
