package app

import (
	"testing"
	"vyshkat/internal/park"
)

func TestRunDemo(t *testing.T) {
	p := NewContainer(park.RegularService{}).Park
	RunDemo(p)

	if p.TransportCount() != 3 {
		t.Fatal("demo should add three vehicles")
	}
	if p.ItemCount() != 2 {
		t.Fatal("demo should add two items")
	}
	energy := p.TotalEnergy()
	if energy < 2.4999 || energy > 2.5001 {
		t.Fatal("demo should use 2.5 units of energy")
	}
	if len(p.ForBeginners()) != 3 {
		t.Fatal("all three demo vehicles should suit beginners")
	}
	inventory := p.Inventory()
	if len(inventory) != 5 {
		t.Fatal("demo should have five inventory items")
	}
	numbers := []int{101, 102, 103, 201, 202}
	for i, item := range inventory {
		if item.GetNumber() != numbers[i] {
			t.Fatal("unexpected inventory number after demo")
		}
		if item.GetName() == "Broken" || item.GetName() == "Duplicate" {
			t.Fatal("rejected transport entered inventory")
		}
	}
}
