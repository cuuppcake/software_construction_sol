package domain_test

import (
	"pos/internal/domain"
	"testing"

	"github.com/google/uuid"
)

func TestNewCriterion(t *testing.T) {
	tests := []struct {
		scenario string
		name     string
		maxScore int
		weight   int
		valid    bool
	}{
		{
			scenario: "пустое название",
			name:     "",
			maxScore: 10,
			weight:   50,
		},
		{
			scenario: "нулевой максимум",
			name:     "Работа",
			maxScore: 0,
			weight:   50,
		},
		{
			scenario: "нулевой вес",
			name:     "Работа",
			maxScore: 10,
			weight:   0,
		},
		{
			scenario: "отрицательный вес",
			name:     "Работа",
			maxScore: 10,
			weight:   -1,
		},
		{
			scenario: "пробелы по краям убираются",
			name:     " Работа ",
			maxScore: 10,
			weight:   50,
			valid:    true,
		},
	}
	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			criterion, err := domain.NewCriterion(test.name, test.maxScore, test.weight)
			if (err == nil) != test.valid {
				t.Fatalf("неожиданная ошибка: %v", err)
			}
			if test.valid && (criterion.ID() == uuid.Nil || criterion.Name() != "Работа") {
				t.Fatal("критерий создан с неверными данными")
			}
		})
	}
}

func TestCriterionChangePreservesID(t *testing.T) {
	criterion, err := domain.NewCriterion("Работа", 10, 50)
	if err != nil {
		t.Fatal(err)
	}
	id := criterion.ID()
	if err := criterion.Change("Экзамен", 20, 30); err != nil {
		t.Fatal(err)
	}
	if criterion.ID() != id || criterion.Name() != "Экзамен" || criterion.MaxScore() != 20 {
		t.Fatal("изменение критерия поменяло ID или не сохранило поля")
	}
	if err := criterion.Change("", 5, 10); err == nil {
		t.Fatal("пустое название было принято")
	}
	if criterion.Name() != "Экзамен" {
		t.Fatal("отклонённое изменение испортило критерий")
	}
}
