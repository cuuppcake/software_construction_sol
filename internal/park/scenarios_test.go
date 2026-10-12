package park

import "testing"

func TestEmptyPark(t *testing.T) {
	p := NewPark(RegularService{})
	if p.TransportCount() != 0 || p.ItemCount() != 0 || p.TotalEnergy() != 0 {
		t.Fatal("empty park must have zero counts and energy")
	}
	if len(p.Inventory()) != 0 || len(p.ForBeginners()) != 0 {
		t.Fatal("empty park must return empty lists")
	}
}

func TestMixedInventory(t *testing.T) {
	p := NewPark(RegularService{})
	vehicles := []Transport{
		NewScooter("Scooter", 1, 0.5, 5),
		NewBicycle("Bicycle", 2, 6),
		NewElectricBicycle("Electric", 3, 0.8, 10),
	}
	for _, vehicle := range vehicles {
		if err := p.AddTransport(vehicle); err != nil {
			t.Fatal(err)
		}
	}
	items := []InventoryItem{NewHelmet("Helmet", 4), NewChargingStation("Station", 5, 1.2)}
	for _, item := range items {
		if err := p.AddItem(item); err != nil {
			t.Fatal(err)
		}
	}
	if p.TransportCount() != 3 || p.ItemCount() != 2 {
		t.Fatal("counts must distinguish transport from items")
	}
	energy := p.TotalEnergy()
	if energy < 2.4999 || energy > 2.5001 {
		t.Fatalf("expected total energy 2.5, got %v", energy)
	}
	beginners := p.ForBeginners()
	if len(beginners) != 2 {
		t.Fatal("expected two beginner vehicles")
	}
	if beginners[0].GetName() != "Bicycle" || beginners[1].GetName() != "Electric" {
		t.Fatal("wrong vehicles selected for beginners")
	}
	inventory := p.Inventory()
	if len(inventory) != 5 {
		t.Fatal("expected five inventory items")
	}
	names := []string{"Scooter", "Bicycle", "Electric", "Helmet", "Station"}
	for i, item := range inventory {
		if item.GetNumber() != i+1 || item.GetName() != names[i] {
			t.Fatal("wrong inventory name or number")
		}
	}
	copyOfInventory := p.Inventory()
	copyOfInventory[0] = NewHelmet("Replacement", 99)
	if p.Inventory()[0].GetName() != "Scooter" {
		t.Fatal("changing the returned slice must not change park inventory")
	}
}

func TestFailedAdditionsPreserveInventory(t *testing.T) {
	p := NewPark(RegularService{})
	if err := p.AddTransport(NewScooter("Existing", 1, 0.5, 8)); err != nil {
		t.Fatal(err)
	}
	for _, vehicle := range []Transport{
		NewBicycle("Duplicate", 1, 7),
		NewBicycle("Invalid ease", 2, 0),
		NewElectricBicycle("Broken", 3, -1, 7),
	} {
		if err := p.AddTransport(vehicle); err == nil {
			t.Fatalf("accepted invalid transport %q", vehicle.GetName())
		}
		inventory := p.Inventory()
		if len(inventory) != 1 {
			t.Fatal("failed transport addition changed inventory")
		}
		if inventory[0].GetName() != "Existing" || inventory[0].GetNumber() != 1 {
			t.Fatal("failed addition replaced existing transport")
		}
	}
	if err := p.AddItem(NewHelmet("Duplicate", 1)); err == nil {
		t.Fatal("item must not reuse a transport number")
	}
	if len(p.Inventory()) != 1 || p.ItemCount() != 0 || p.TotalEnergy() != 0.5 {
		t.Fatal("rejected additions changed park state")
	}
	if err := p.AddTransport(NewElectricBicycle("Repaired", 3, 0, 7)); err != nil {
		t.Fatalf("rejected transport must not reserve its number: %v", err)
	}
}

func TestRegularServiceIdentityAndEnergy(t *testing.T) {
	service := RegularService{}
	invalid := []Transport{
		NewBicycle("", 1, 7),
		NewBicycle("Bike", 0, 7),
		NewBicycle("Bike", -1, 7),
		NewElectricBicycle("Electric", 1, -0.1, 7),
	}
	for _, vehicle := range invalid {
		if service.Check(vehicle) {
			t.Fatal("invalid transport passed inspection")
		}
	}
	if !service.Check(NewElectricBicycle("Electric", 1, 0, 7)) {
		t.Fatal("zero energy should be allowed")
	}
	if !service.Check(NewScooter("Scooter", 1, 0.5, 7)) {
		t.Fatal("valid scooter should pass inspection")
	}
}
