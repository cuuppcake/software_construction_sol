package calculation_test

import (
	"math"
	"pos/internal/calculation"
	"pos/internal/domain"
	"testing"
)

func rubricAndGrade(t *testing.T) (domain.Student, *domain.Rubric, domain.Evaluation) {
	t.Helper()
	student, err := domain.NewStudent("Анна")
	if err != nil {
		t.Fatal(err)
	}
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
	evaluation, err := (domain.EvaluationFactory{}).Create(student, rubric, first.ID(), 8)
	if err != nil {
		t.Fatal(err)
	}
	return student, rubric, evaluation
}

func TestResultPoliciesWithMissingGrade(t *testing.T) {
	student, rubric, evaluation := rubricAndGrade(t)
	tests := []struct {
		name     string
		policy   calculation.ResultPolicy
		expected float64
	}{
		{
			name:     "по весам",
			policy:   calculation.WeightedPolicy{},
			expected: 56,
		},
		{
			name:     "поровну",
			policy:   calculation.EqualPolicy{},
			expected: 40,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := calculation.NewCalculator(test.policy).Calculate(student, rubric, []domain.Evaluation{evaluation})
			if err != nil {
				t.Fatal(err)
			}
			if math.Abs(result.Percent()-test.expected) > 0.0001 || result.Completed() != 1 || result.Total() != 2 {
				t.Fatalf("ожидались %.2f%% и 1/2 критериев, получены %.2f%% и %d/%d", test.expected, result.Percent(), result.Completed(), result.Total())
			}
		})
	}
}

func TestCalculatorRejectsAnotherStudentsGrade(t *testing.T) {
	student, rubric, _ := rubricAndGrade(t)
	other, err := domain.NewStudent("Борис")
	if err != nil {
		t.Fatal(err)
	}
	criterion := rubric.List()[0]
	evaluation, err := (domain.EvaluationFactory{}).Create(other, rubric, criterion.ID(), 8)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := calculation.NewCalculator(calculation.WeightedPolicy{}).Calculate(student, rubric, []domain.Evaluation{evaluation}); err == nil {
		t.Fatal("чужая оценка попала в итог студента")
	}
}
