package park

type InventoryItem interface {
	GetName() string
	GetNumber() int
}

type Transport interface {
	InventoryItem
	GetEase() int
}

type EnergyConsumer interface {
	GetEnergy() float64
}

type ServiceCenter interface {
	Check(t Transport) bool
}
