package network

type Network interface {
	InterfaceName() string
	Addresses() []string
	Gateway() string
}

type NetworkImpl struct {
	name    string
	address []string
	gateway string
}

var _ Network = (*NetworkImpl)(nil)

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
		name:    interfaceName,
		address: addresses,
		gateway: gateway,
	}, nil
}

func (n *NetworkImpl) InterfaceName() string {
	return n.name
}

func (n *NetworkImpl) Addresses() []string {
	return n.address
}

func (n *NetworkImpl) Gateway() string {
	return n.gateway
}
