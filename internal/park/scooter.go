package park

type Scooter struct {
	name   string
	number int
	energy float64
	ease   int
}

func NewScooter(name string, number int, energy float64, ease int) *Scooter {
	return &Scooter{name: name, number: number, energy: energy, ease: ease}
}

func (v *Scooter) GetName() string {
	return v.name
}

func (v *Scooter) GetNumber() int {
	return v.number
}

func (v *Scooter) GetEnergy() float64 {
	return v.energy
}

func (v *Scooter) GetEase() int {
	return v.ease
}
