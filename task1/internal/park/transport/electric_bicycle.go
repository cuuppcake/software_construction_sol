package transport

type ElectricBicycle struct {
	name   string
	number int
	energy float64
	ease   int
}

func NewElectricBicycle(name string, number int, energy float64, ease int) *ElectricBicycle {
	return &ElectricBicycle{name: name, number: number, energy: energy, ease: ease}
}

func (v *ElectricBicycle) GetName() string {
	return v.name
}

func (v *ElectricBicycle) GetNumber() int {
	return v.number
}

func (v *ElectricBicycle) GetEnergy() float64 {
	return v.energy
}

func (v *ElectricBicycle) GetEase() int {
	return v.ease
}
