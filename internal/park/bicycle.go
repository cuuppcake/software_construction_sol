package park

type Bicycle struct {
	name   string
	number int
	ease   int
}

func NewBicycle(name string, number int, ease int) *Bicycle {
	return &Bicycle{name: name, number: number, ease: ease}
}

func (v *Bicycle) GetName() string {
	return v.name
}

func (v *Bicycle) GetNumber() int {
	return v.number
}

func (v *Bicycle) GetEase() int {
	return v.ease
}
