package config

type User interface {
	User() string
	Password() string
	SSHKeys() []string
}

type userImpl struct {
	user     string
	password string
	sshKeys  []string
}

var _ User = (*userImpl)(nil)

func (u *userImpl) Password() string {
	return u.password
}

func (u *userImpl) SSHKeys() []string {
	return u.sshKeys
}

func (u *userImpl) User() string {
	return u.user
}
