package park

type Transport interface {
	InventoryItem
	GetEase() int
}

type EnergyConsumer interface {
	GetEnergy() float64
}
