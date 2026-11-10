package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Primary  PrimaryConfig  `env-prefix:"PRIMARY_"`
	Server   ServerConfig   `env-prefix:"SERVER_"`
	Database DatabaseConfig `env-prefix:"DATABASE_"`
}

type PrimaryConfig struct {
	Env string `env:"ENV" env-required:"true"`
}

type ServerConfig struct {
	Port         int      `env:"PORT" env-required:"true"`
	ReadTimeout  int      `env:"READ_TIMEOUT" env-required:"true"`
	WriteTimeout int      `env:"WRITE_TIMEOUT" env-required:"true"`
	IdleTimeout  int      `env:"IDLE_TIMEOUT" env-required:"true"`
	AllowedIPs   []string `env:"ALLOWED_IPS"`
}

type DatabaseConfig struct {
	Port         int    `env:"PORT" env-required:"true"`
	Host         string `env:"HOST" env-required:"true"`
	User         string `env:"USER" env-required:"true"`
	Password     string `env:"PASSWORD"`
	Name         string `env:"NAME" env-required:"true"`
	SSLMode      bool   `env:"SSLMODE" env-required:"true"`
	MaxOpenConns int    `env:"MAXOPENCONNS" env-required:"true"`
	MaxIdleConns int    `env:"MAXIDLECONNS" env-required:"true"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		err = cleanenv.ReadEnv(&cfg)
		if err != nil {
			return nil, fmt.Errorf("could not load environment variables: %w", err)
		}
	}
	return &cfg, nil
}
