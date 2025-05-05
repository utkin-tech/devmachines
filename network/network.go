package network

type Network[T Addr] interface {
	InterfaceName() string
	Addresses() []T
	Gateway() string
}

type NetworkImpl struct {
	name      string
	addresses []*AddrImpl
	gateway   string
}

var _ Network[*AddrImpl] = (*NetworkImpl)(nil)

func NewNetwork(interfaceName string) (*NetworkImpl, error) {
	addresses, err := GetAddressesByInterface(interfaceName)
	if err != nil {
		return nil, err
	}

	gateway, err := GetDefaultGateway()
	if err != nil {
		return nil, err
	}

	return &NetworkImpl{
		name:      interfaceName,
		addresses: addresses,
		gateway:   gateway,
	}, nil
}

func (n *NetworkImpl) InterfaceName() string {
	return n.name
}

func (n *NetworkImpl) Addresses() []*AddrImpl {
	return n.addresses
}

func (n *NetworkImpl) Gateway() string {
	return n.gateway
}
