package park

type Helmet struct {
	name   string
	number int
}

func NewHelmet(name string, number int) *Helmet {
	return &Helmet{name: name, number: number}
}

func (v *Helmet) GetName() string {
	return v.name
}

func (v *Helmet) GetNumber() int {
	return v.number
}
