package park

type ServiceCenter interface {
	Check(t Transport) bool
}

type RegularService struct{}

func (s RegularService) Check(t Transport) bool {
	if t == nil {
		return false
	}
	if t.GetName() == "" || t.GetNumber() <= 0 {
		return false
	}
	if t.GetEase() < 1 || t.GetEase() > 10 {
		return false
	}
	if e, ok := t.(EnergyConsumer); ok && e.GetEnergy() < 0 {
		return false
	}
	return true
}
