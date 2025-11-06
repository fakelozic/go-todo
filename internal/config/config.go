package config

import (
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog"
)

type Config struct {
	Server   ServerConfig   `koanf:"server" validate:"required"`
	Database DatabaseConfig `koanf:"database" validate:"required"`
}

type ServerConfig struct {
	Port               string   `koanf:"port" validate:"required"`
	CORSAllowedOrigins []string `koanf:"cors_allowed_origins" validate:"required"`
}

type DatabaseConfig struct {
	DSN string `koanf:"dsn" validate:"required"`
}

func LoadConfig() (*Config, error) {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	k := koanf.New(".")

	err := k.Load(env.Provider("GO_TODO_", ".", func(s string) string {
		return strings.ToLower(strings.TrimPrefix(s, "GO_TODO_"))
	}), nil)
	if err != nil {
		logger.Fatal().Err(err).Msg("could not load initial env variables")
	}

	mainConfig := &Config{}

	err = k.Unmarshal("", mainConfig)
	if err != nil {
		logger.Fatal().Err(err).Msg("could not unmarshal main config")
	}

	validate := validator.New()

	err = validate.Struct(mainConfig)
	if err != nil {
		logger.Fatal().Err(err).Msg("config validation failed")
	}

	return mainConfig, nil
}
