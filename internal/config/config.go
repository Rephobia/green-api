package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env     string `env:"APP_ENV"  env-default:"local"`
	AppName string `env:"APP_NAME" env-default:"green-api"`
	Port    string `env:"PORT"     env-default:"8080"`
}

func New() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to read env file (%w)", err)
	}

	return &cfg, nil
}
