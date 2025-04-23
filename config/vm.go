package config

type VM interface {
	CPU() uint
	Memory() uint
}

type vmImpl struct {
	cpu    uint
	memory uint
}

var _ VM = (*vmImpl)(nil)

func (v *vmImpl) CPU() uint {
	return v.cpu
}

func (v *vmImpl) Memory() uint {
	return v.memory
}
