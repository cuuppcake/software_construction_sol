package domain_test

import (
	"math"
	"pos/internal/domain"
	"testing"
)

func testRubric(t *testing.T) *domain.Rubric {
	t.Helper()
	rubric := domain.NewRubric()
	first, err := domain.NewCriterion("work", "Work", 10, 0.7)
	if err != nil {
		t.Fatal(err)
	}
	second, err := domain.NewCriterion("exam", "Exam", 20, 0.3)
	if err != nil {
		t.Fatal(err)
	}
	if err := rubric.Add(first); err != nil {
		t.Fatal(err)
	}
	if err := rubric.Add(second); err != nil {
		t.Fatal(err)
	}
	return rubric
}

func TestEvaluationFactoryRejectsInvalidGrades(t *testing.T) {
	student, err := domain.NewStudent("s1", "Anna")
	if err != nil {
		t.Fatal(err)
	}
	rubric := testRubric(t)
	factory := domain.EvaluationFactory{}
	for _, test := range []struct {
		criterion string
		points    int
	}{
		{"work", -1}, {"work", 11}, {"unknown", 5},
	} {
		if _, err := factory.Create(student, rubric, test.criterion, test.points); err == nil {
			t.Fatalf("приняты критерий %q и балл %d", test.criterion, test.points)
		}
	}
	evaluation, err := factory.Create(student, rubric, "work", 0)
	if err != nil || evaluation.GetScore().GetPoints() != 0 {
		t.Fatal("ноль должен быть допустимой оценкой")
	}
}

func TestCalculatorPoliciesAndMissingGrades(t *testing.T) {
	student, _ := domain.NewStudent("s1", "Anna")
	rubric := testRubric(t)
	evaluation, err := (domain.EvaluationFactory{}).Create(student, rubric, "work", 8)
	if err != nil {
		t.Fatal(err)
	}
	weighted, err := domain.NewCalculator(domain.WeightedPolicy{}).Calculate(student, rubric, []domain.Evaluation{evaluation})
	if err != nil {
		t.Fatal(err)
	}
	if weighted.GetPercent() < 55.9999 || weighted.GetPercent() > 56.0001 || weighted.GetCompleted() != 1 || weighted.GetTotal() != 2 {
		t.Fatalf("ожидались 56%% и 1/2 заполненных критериев, получены %.2f%% и %d/%d", weighted.GetPercent(), weighted.GetCompleted(), weighted.GetTotal())
	}
	equal, err := domain.NewCalculator(domain.EqualPolicy{}).Calculate(student, rubric, []domain.Evaluation{evaluation})
	if err != nil {
		t.Fatal(err)
	}
	if equal.GetPercent() < 39.9999 || equal.GetPercent() > 40.0001 {
		t.Fatalf("ожидались 40%%, получено %.2f%%", equal.GetPercent())
	}
}

func TestCalculatorRejectsAnotherStudentsEvaluation(t *testing.T) {
	student, err := domain.NewStudent("s1", "Анна")
	if err != nil {
		t.Fatal(err)
	}
	other, err := domain.NewStudent("s2", "Борис")
	if err != nil {
		t.Fatal(err)
	}
	rubric := testRubric(t)
	evaluation, err := (domain.EvaluationFactory{}).Create(other, rubric, "work", 8)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := domain.NewCalculator(domain.WeightedPolicy{}).Calculate(student, rubric, []domain.Evaluation{evaluation}); err == nil {
		t.Fatal("чужая оценка попала в итог студента")
	}
}

func TestRubricProtectsCriteria(t *testing.T) {
	rubric := testRubric(t)
	if err := rubric.Add(domain.Criterion{}); err == nil {
		t.Fatal("пустой критерий был добавлен")
	}
	if err := rubric.Add(rubric.List()[0]); err == nil {
		t.Fatal("повторный критерий был добавлен")
	}
	copy := rubric.Clone()
	if err := copy.Remove("work"); err != nil {
		t.Fatal(err)
	}
	if len(rubric.List()) != 2 {
		t.Fatal("изменение копии поменяло исходный рубрикатор")
	}
}

func TestCriterionRejectsInvalidWeight(t *testing.T) {
	for _, weight := range []float64{0, -1, math.NaN(), math.Inf(1)} {
		if _, err := domain.NewCriterion("work", "Работа", 10, weight); err == nil {
			t.Fatalf("критерий с весом %v был создан", weight)
		}
	}
}
