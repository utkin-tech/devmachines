package main

const (
	Username = "user"
	Password = "$5$KXasYhNX$QkJkEVjIhA/.W1qjTPhJzgJXeZpvu8RSsGx1HqaxX23"
	DiskSize = "10G"
)

var AuthorizedKeys = []string{
	"ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIMBEMPGKBgwtQaQ3un8I7j3wIzrknlCUUoetLWJCwfzn erlnby@DESKTOP-46HDCTP",
	"ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIJx/50sox9LvHx1rrkmOYjn4hitH3jvkm4JomiMdwLUz utkin_danil@mail.ru",
	"ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAICeyDWnYtatyEiHKjT/jpLUA7FLMoe97ItrkaLA2cYsX qemu_machine",
}

type User interface {
	User() string
	Password() string
	SSHKeys() []string
	DiskSize() string
}

type UserImpl struct{}

var _ User = (*UserImpl)(nil)

func NewUser() *UserImpl {
	return &UserImpl{}
}

func (u *UserImpl) DiskSize() string {
	return DiskSize
}

func (u *UserImpl) Password() string {
	return Password
}

func (u *UserImpl) SSHKeys() []string {
	return AuthorizedKeys
}

func (u *UserImpl) User() string {
	return Username
}
