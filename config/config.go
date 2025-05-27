package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type (
	// Config -.
	Config struct {
		App   App
		HTTP  HTTP
		Log   Log
		MYSQL MYSQL
		JWT   JWT
	}

	// App -.
	App struct {
		Name    string `env:"APP_NAME,required"`
		Env     string `env:"APP_ENV,required"`
		Version string `env:"APP_VERSION,required"`
	}

	// HTTP -.
	HTTP struct {
		Port           string `env:"HTTP_PORT,required"`
		UsePreforkMode bool   `env:"HTTP_USE_PREFORK_MODE" envDefault:"false"`
	}

	JWT struct {
		Secret string `env:"JWT_SECRET,required"`
	}

	// Log -.
	Log struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	// MYSQL -.
	MYSQL struct {
		PoolMax  int    `env:"MYSQL_POOL_MAX,required"`
		Username string `env:"MYSQL_USERNAME,required"`
		Password string `env:"MYSQL_PASSWORD"`
		Database string `env:"MYSQL_DATABASE,required"`
		Host     string `env:"MYSQL_HOST,required"`
		Port     string `env:"MYSQL_PORT,required"`
	}
)

// New returns app config.
func New() (*Config, error) {
	_ = godotenv.Load(".env")

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
