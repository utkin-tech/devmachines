package config

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/docker/go-connections/nat"
	"github.com/docker/go-units"
	"github.com/go-playground/validator/v10"
)

type NetworkType string

const (
	NetworkTypeNat    = "NAT"
	NetworkTypeBridge = "BRIDGE"
)

type Environment struct {
	CPU      int         `env:"CPU" validate:"required"`
	Memory   string      `env:"MEMORY" validate:"required,validBytes"`
	Storage  string      `env:"STORAGE" validate:"required,validBytes"`
	User     string      `env:"USER" validate:"required"`
	Password string      `env:"PASSWORD" validate:"required"`
	SSHKeys  []string    `env:"SSH_KEYS" validate:"required"`
	Network  NetworkType `env:"NETWORK" validate:"required,validNetwork" envDefault:"NAT"`
	VNC      string      `env:"VNC"`
	Ports    []string    `env:"PORTS" validate:"validPorts"`
}

func LoadEnvironment() (*Environment, error) {
	var cfg Environment

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("failed load environment variables: %w", err)
	}

	validate := validator.New()

	validate.RegisterValidation("validBytes", validateBytes)
	validate.RegisterValidation("validNetwork", validateNetwork)
	validate.RegisterValidation("validPorts", validatePorts)

	if err := validate.Struct(cfg); err != nil {
		return nil, fmt.Errorf("failed validate config: %w", err)
	}

	return &cfg, nil
}

func validateBytes(fl validator.FieldLevel) bool {
	_, err := units.RAMInBytes(fl.Field().String())
	return err == nil
}

func validateNetwork(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return value == NetworkTypeNat || value == NetworkTypeBridge
}

func validatePorts(fl validator.FieldLevel) bool {
	ports := fl.Field()

	for i := range ports.Len() {
		port := ports.Index(i).String()
		_, err := nat.ParsePortSpec(port)
		if err != nil {
			return false
		}
	}

	return true
}
