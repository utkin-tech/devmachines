package config

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/docker/go-units"
	"github.com/go-playground/validator/v10"
)

type Environment struct {
	CPU      int      `env:"CPU" validate:"required"`
	Memory   string   `env:"MEMORY" validate:"required,validMemory"`
	Storage  string   `env:"STORAGE" validate:"required,validStorage"`
	User     string   `env:"USER" validate:"required"`
	Password string   `env:"PASSWORD" validate:"required"`
	SSHKeys  []string `env:"SSH_KEYS" validate:"required"`
}

func LoadEnvironment() (*Environment, error) {
	var cfg Environment

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("failed load environment variables: %w", err)
	}

	validate := validator.New()

	validate.RegisterValidation("validMemory", validateMemory)
	validate.RegisterValidation("validStorage", validateStorage)

	if err := validate.Struct(cfg); err != nil {
		return nil, fmt.Errorf("failed validate config: %w", err)
	}

	return &cfg, nil
}

func validateMemory(fl validator.FieldLevel) bool {
	_, err := units.RAMInBytes(fl.Field().String())
	return err == nil
}

func validateStorage(fl validator.FieldLevel) bool {
	_, err := units.RAMInBytes(fl.Field().String())
	return err == nil
}
