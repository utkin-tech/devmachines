package config

import (
	"fmt"
	"strconv"

	"github.com/docker/go-units"
)

type Config interface {
	VM() VM
	Storage() Storage
	User() User
}

type ConfigImpl struct {
	env *Environment
}

var _ Config = (*ConfigImpl)(nil)

func NewConfig(env *Environment) *ConfigImpl {
	return &ConfigImpl{
		env: env,
	}
}

func (c *ConfigImpl) Storage() Storage {
	sizeB := mustParse(c.env.Storage)
	size := strconv.FormatInt(sizeB, 10)

	storage := &storageImpl{
		size: size,
	}

	return storage
}

func (c *ConfigImpl) User() User {
	user := &userImpl{
		user:     c.env.User,
		password: c.env.Password,
		sshKeys:  c.env.SSHKeys,
	}

	return user
}

func (c *ConfigImpl) VM() VM {
	memoryB := mustParse(c.env.Memory)
	memoryM := memoryB / units.MiB

	vm := &vmImpl{
		cpu:    uint(c.env.CPU),
		memory: uint(memoryM),
	}

	return vm
}

func mustParse(str string) int64 {
	i, err := units.RAMInBytes(str)
	if err != nil {
		panic(fmt.Errorf("cannot parse '%v': %v", str, err))
	}
	return i
}
