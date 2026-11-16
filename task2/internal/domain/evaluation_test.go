package domain_test

import (
	"pos/internal/domain"
	"testing"

	"github.com/google/uuid"
)

func TestEvaluationFactory(t *testing.T) {
	student, err := domain.NewStudent("Анна")
	if err != nil {
		t.Fatal(err)
	}
	rubric, criterion, _ := testRubric(t)
	tests := []struct {
		name        string
		criterionID uuid.UUID
		points      int
		valid       bool
	}{
		{
			name:        "ноль",
			criterionID: criterion.ID(),
			points:      0,
			valid:       true,
		},
		{
			name:        "максимум",
			criterionID: criterion.ID(),
			points:      10,
			valid:       true,
		},
		{
			name:        "отрицательный",
			criterionID: criterion.ID(),
			points:      -1,
		},
		{
			name:        "выше максимума",
			criterionID: criterion.ID(),
			points:      11,
		},
		{
			name:        "нет критерия",
			criterionID: uuid.Nil,
			points:      5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			evaluation, err := (domain.EvaluationFactory{}).Create(student, rubric, test.criterionID, test.points)
			if (err == nil) != test.valid {
				t.Fatalf("неожиданная ошибка: %v", err)
			}
			if test.valid && (evaluation.StudentID() != student.ID() || evaluation.Score().Points() != test.points) {
				t.Fatal("оценка создана с неверными данными")
			}
		})
	}
}
