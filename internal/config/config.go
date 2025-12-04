package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ApiPort         string        `envconfig:"API_PORT"`
	DiagPort        string        `envconfig:"DIAG_PORT"`
	DBUrl           string        `envconfig:"DB_URL"`
	HTTPTimeout     time.Duration `envconfig:"HTTP_TIMEOUT"`
	IdleTimeout     time.Duration `envconfig:"IDLE_TIMEOUT"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT"`
}

func Load() (*Config, error) {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, fmt.Errorf("can't process the config: %w", err)
	}

	return &cfg, nil
}
