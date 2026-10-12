package park

import (
	"errors"
	"fmt"
	"reflect"
)

type Park struct {
	vehicles []Transport
	items    []InventoryItem
	service  ServiceCenter
}

func NewPark(service ServiceCenter) *Park {
	return &Park{service: service}
}

func isNil(value any) bool {
	if value == nil {
		return true
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return v.IsNil()
	}
	return false
}

func (p *Park) numberExists(number int) bool {
	for _, v := range p.vehicles {
		if v.GetNumber() == number {
			return true
		}
	}
	for _, item := range p.items {
		if item.GetNumber() == number {
			return true
		}
	}
	return false
}

func (p *Park) checkTransport(t Transport) error {
	if isNil(t) {
		return errors.New("transport is nil")
	}
	if t.GetName() == "" || t.GetNumber() <= 0 || t.GetEase() < 1 || t.GetEase() > 10 {
		return errors.New("invalid transport")
	}
	if p.numberExists(t.GetNumber()) {
		return fmt.Errorf("inventory number %d already exists", t.GetNumber())
	}
	if isNil(p.service) {
		return errors.New("service center is not configured")
	}
	if !p.service.Check(t) {
		return errors.New("inspection failed")
	}
	return nil
}

func (p *Park) AddTransport(t Transport) error {
	if err := p.checkTransport(t); err != nil {
		return err
	}
	p.vehicles = append(p.vehicles, t)
	return nil
}

func (p *Park) checkItem(item InventoryItem) error {
	if isNil(item) {
		return errors.New("item is nil")
	}
	if item.GetName() == "" || item.GetNumber() <= 0 {
		return errors.New("invalid item")
	}
	if p.numberExists(item.GetNumber()) {
		return fmt.Errorf("inventory number %d already exists", item.GetNumber())
	}
	return nil
}

func (p *Park) AddItem(item InventoryItem) error {
	if err := p.checkItem(item); err != nil {
		return err
	}
	p.items = append(p.items, item)
	return nil
}

func (p *Park) TransportCount() int {
	return len(p.vehicles)
}

func (p *Park) ItemCount() int {
	return len(p.items)
}

func (p *Park) TotalEnergy() float64 {
	var total float64
	for _, v := range p.vehicles {
		if e, ok := v.(EnergyConsumer); ok {
			total += e.GetEnergy()
		}
	}
	for _, item := range p.items {
		if e, ok := item.(EnergyConsumer); ok {
			total += e.GetEnergy()
		}
	}
	return total
}

func (p *Park) ForBeginners() []Transport {
	var result []Transport
	for _, v := range p.vehicles {
		if v.GetEase() >= 6 {
			result = append(result, v)
		}
	}
	return result
}

func (p *Park) Inventory() []InventoryItem {
	var result []InventoryItem
	for _, v := range p.vehicles {
		result = append(result, v)
	}
	result = append(result, p.items...)
	return result
}
