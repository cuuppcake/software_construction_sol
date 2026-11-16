package calculation

import (
	"errors"
	"fmt"
	"pos/internal/domain"

	"github.com/google/uuid"
)

type Calculator struct {
	policy ResultPolicy
}

func NewCalculator(policy ResultPolicy) *Calculator {
	return &Calculator{policy: policy}
}

func (c *Calculator) Calculate(student domain.Student, rubric *domain.Rubric, evaluations []domain.Evaluation) (domain.Result, error) {
	if c == nil || c.policy == nil || rubric == nil {
		return domain.Result{}, errors.New("нужны рубрикатор и правило расчёта итога")
	}
	if student.ID() == uuid.Nil {
		return domain.Result{}, errors.New("нужен студент для расчёта итога")
	}
	criteria := rubric.List()
	if len(criteria) == 0 {
		return domain.Result{}, errors.New("рубрикатор пуст")
	}
	scores := make(map[uuid.UUID]domain.Score)
	for _, evaluation := range evaluations {
		if evaluation.StudentID() != student.ID() {
			return domain.Result{}, errors.New("оценка относится к другому студенту")
		}
		criterion, exists := rubric.Criterion(evaluation.CriterionID())
		if !exists {
			return domain.Result{}, fmt.Errorf("критерия %q нет в рубрикаторе", evaluation.CriterionID())
		}
		if _, exists := scores[evaluation.CriterionID()]; exists {
			return domain.Result{}, fmt.Errorf("по критерию %q найдены две оценки", evaluation.CriterionID())
		}
		score := evaluation.Score()
		if _, err := domain.NewScore(score.Points(), criterion.MaxScore()); err != nil {
			return domain.Result{}, err
		}
		scores[evaluation.CriterionID()] = score
	}
	return domain.NewResult(c.policy.Percent(criteria, scores), len(scores), len(criteria)), nil
}
