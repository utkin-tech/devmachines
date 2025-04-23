package config

type Storage interface {
	Size() string
}

type storageImpl struct {
	size string
}

var _ Storage = (*storageImpl)(nil)

func (s *storageImpl) Size() string {
	return s.size
}
