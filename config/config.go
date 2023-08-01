package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type (
	// Config -.
	Config struct {
		App      `yaml:"app"`
		HTTP     `yaml:"http"`
		Logger   `yaml:"logger"`
		Postgres `yaml:"postgres"`
		Redis    `yaml:"redis"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port         string        `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		ReadTimeout  time.Duration `yaml:"read-timeout" env:"HTTP-READ-TIMEOUT"`
		WriteTimeout time.Duration `yaml:"write-timeout" env:"HTTP-WRITE-TIMEOUT"`
	}

	// Logger -.
	Logger struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// Postgres -.
	Postgres struct {
		Username    string        `yaml:"username" env:"PSQL_USERNAME" env-required:"true"`
		Password    string        `yaml:"password" env:"PSQL_PASSWORD" env-required:"true"`
		Host        string        `yaml:"host" env:"PSQL_HOST" env-required:"true"`
		Port        string        `yaml:"port" env:"PSQL_PORT" env-required:"true"`
		Database    string        `yaml:"database" env:"PSQL_DATABASE" env-required:"true"`
		MaxAttempts int           `yaml:"max-attempts" env:"PSQL_MAX_ATTEMPTS" env-required:"true"`
		MaxDelay    time.Duration `yaml:"max-delay" env:"PSQL_MAX_DELAY" env-required:"true"`
	}

	// Redis -.
	Redis struct {
		Host     string `yaml:"host" env:"REDIS_HOST" env-required:"true"`
		Port     string `yaml:"port" env:"REDIS_PORT" env-required:"true"`
		DB       int    `yaml:"db" env:"REDIS_DB"`
		Password string `yaml:"password" env:"REDIS_PASSWORD"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
