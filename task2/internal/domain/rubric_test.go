package domain_test

import (
	"pos/internal/domain"
	"testing"
)

func testRubric(t *testing.T) (*domain.Rubric, domain.Criterion, domain.Criterion) {
	t.Helper()
	first, err := domain.NewCriterion("Работа", 10, 70)
	if err != nil {
		t.Fatal(err)
	}
	second, err := domain.NewCriterion("Экзамен", 20, 30)
	if err != nil {
		t.Fatal(err)
	}
	rubric := domain.NewRubric()
	if err := rubric.Add(first); err != nil {
		t.Fatal(err)
	}
	if err := rubric.Add(second); err != nil {
		t.Fatal(err)
	}
	return rubric, first, second
}

func TestRubricProtectsCriteria(t *testing.T) {
	rubric, first, _ := testRubric(t)
	if err := rubric.Add(domain.Criterion{}); err == nil {
		t.Fatal("пустой критерий был добавлен")
	}
	if err := rubric.Add(first); err == nil {
		t.Fatal("повторный критерий был добавлен")
	}
	clone := rubric.Clone()
	if err := clone.Remove(first.ID()); err != nil {
		t.Fatal(err)
	}
	if len(rubric.List()) != 2 {
		t.Fatal("изменение копии поменяло исходный рубрикатор")
	}
}
