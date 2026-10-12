package park

import (
	"strings"
	"testing"
)

type FakeService struct {
	allow bool
}

func (s FakeService) Check(t Transport) bool {
	return s.allow
}

func TestInspection(t *testing.T) {
	p := NewContainer(FakeService{allow: false}).Park
	if err := p.AddTransport(NewScooter("Ninebot", 1, 0.5, 8)); err == nil {
		t.Fatal("rejected transport should return an error")
	}
	if p.TransportCount() != 0 {
		t.Fatal("rejected transport entered the park")
	}
	p = NewContainer(FakeService{allow: true}).Park
	if err := p.AddTransport(NewScooter("Ninebot", 1, 0.5, 8)); err != nil {
		t.Fatal(err)
	}
}

func TestBeginners(t *testing.T) {
	p := NewPark(FakeService{allow: true})
	p.AddTransport(NewBicycle("Easy", 1, 6))
	p.AddTransport(NewBicycle("Hard", 2, 5))
	result := p.ForBeginners()
	if len(result) != 1 || result[0].GetName() != "Easy" {
		t.Fatal("only vehicles with ease >= 6 should be returned")
	}
}

func TestEnergy(t *testing.T) {
	p := NewPark(FakeService{allow: true})
	p.AddTransport(NewScooter("A", 1, 0.5, 7))
	p.AddTransport(NewBicycle("B", 2, 7))
	p.AddItem(NewChargingStation("C", 3, 1.2))
	if energy := p.TotalEnergy(); energy < 1.6999 || energy > 1.7001 {
		t.Fatalf("expected 1.7, got %v", energy)
	}
}

func TestInventoryAndDuplicate(t *testing.T) {
	p := NewPark(FakeService{allow: true})
	if err := p.AddItem(NewHelmet("Helmet", 1)); err != nil {
		t.Fatal(err)
	}
	if err := p.AddTransport(NewBicycle("Bike", 1, 7)); err == nil || !strings.Contains(err.Error(), "1") {
		t.Fatal("duplicate error should mention the number")
	}
	if err := p.AddTransport(NewBicycle("Bike", 2, 7)); err != nil {
		t.Fatal(err)
	}
	if len(p.Inventory()) != 2 {
		t.Fatal("inventory should include transport and items")
	}
}

func TestRegularService(t *testing.T) {
	service := RegularService{}
	cases := []struct {
		ease     int
		expected bool
	}{{0, false}, {1, true}, {10, true}, {11, false}}
	for _, c := range cases {
		if got := service.Check(NewBicycle("Bike", 1, c.ease)); got != c.expected {
			t.Errorf("ease %d: expected %v", c.ease, c.expected)
		}
	}
	if service.Check(nil) {
		t.Fatal("nil transport should fail inspection")
	}
}

func TestMissingService(t *testing.T) {
	p := NewPark(nil)
	if p.AddTransport(NewBicycle("Bike", 1, 7)) == nil {
		t.Fatal("missing service must return an error")
	}
}

func TestInvalidItems(t *testing.T) {
	p := NewPark(FakeService{allow: true})
	for _, item := range []InventoryItem{nil, NewHelmet("", 1), NewHelmet("Helmet", 0)} {
		if p.AddItem(item) == nil {
			t.Fatal("invalid item was accepted")
		}
	}
	if err := p.AddItem(NewHelmet("Helmet", 3)); err != nil {
		t.Fatal(err)
	}
	if p.AddItem(NewHelmet("Second", 3)) == nil {
		t.Fatal("duplicate item number was accepted")
	}
}

func TestTransportValidation(t *testing.T) {
	p := NewPark(RegularService{})
	for _, vehicle := range []Transport{nil, NewBicycle("", 1, 7), NewBicycle("Bike", 0, 7), NewBicycle("Bike", 2, 11)} {
		if p.AddTransport(vehicle) == nil {
			t.Fatal("invalid transport was accepted")
		}
	}
}

func TestTypedNil(t *testing.T) {
	p := NewPark(RegularService{})
	var scooter *Scooter
	if p.AddTransport(scooter) == nil {
		t.Fatal("nil scooter pointer was accepted")
	}
	var helmet *Helmet
	if p.AddItem(helmet) == nil {
		t.Fatal("nil helmet pointer was accepted")
	}
}

func TestConstructors(t *testing.T) {
	if NewScooter("S", 1, 0.5, 7).GetEnergy() != 0.5 || NewScooter("S", 1, 0.5, 7).GetEase() != 7 {
		t.Fatal("scooter fields incorrect")
	}
	if NewElectricBicycle("E", 2, 0.8, 6).GetEnergy() != 0.8 {
		t.Fatal("electric bicycle energy incorrect")
	}
	if NewHelmet("H", 3).GetNumber() != 3 || NewChargingStation("C", 4, 1.2).GetName() != "C" {
		t.Fatal("item fields incorrect")
	}
}

func TestInspectionRejectsNegativeEnergy(t *testing.T) {
	p := NewPark(RegularService{})
	if p.AddTransport(NewScooter("Broken", 10, -0.4, 7)) == nil {
		t.Fatal("inspection must reject negative energy")
	}
}
