package config

import (
	"log/slog"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/knadh/koanf/providers/env/v2"
	"github.com/knadh/koanf/v2"
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
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	k := koanf.New(".")

	err := k.Load(env.Provider(".", env.Opt{
		Prefix: "GO_TODO_",
		TransformFunc: func(k, v string) (string, any) {
			k = strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(k, "GO_TODO_")), "__", ".")

			switch k {
			case "server.cors_allowed_origins":
				if strings.Contains(v, ",") {
					rawSlice := strings.Split(v, ",")
					var cleanSlice []string

					for _, item := range rawSlice {
						if trimmed := strings.TrimSpace(item); trimmed != "" {
							cleanSlice = append(cleanSlice, trimmed)
						}
					}
					return k, cleanSlice
				}
			}
			return k, v
		},
	}), nil)
	if err != nil {
		logger.Error("could not load initial env variables", "error:", err)
	}

	mainConfig := &Config{}

	err = k.Unmarshal("", mainConfig)
	if err != nil {
		logger.Error("could not unmarshal main config", "error:", err)
	}

	validate := validator.New()
	err = validate.Struct(mainConfig)
	if err != nil {
		logger.Error("config validation failed", "error:", err)
	}

	return mainConfig, nil
}
