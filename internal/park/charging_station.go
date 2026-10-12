package park

type ChargingStation struct {
	name   string
	number int
	energy float64
}

func NewChargingStation(name string, number int, energy float64) *ChargingStation {
	return &ChargingStation{name: name, number: number, energy: energy}
}

func (v *ChargingStation) GetName() string {
	return v.name
}

func (v *ChargingStation) GetNumber() int {
	return v.number
}

func (v *ChargingStation) GetEnergy() float64 {
	return v.energy
}
