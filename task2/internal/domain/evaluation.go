package domain

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type Evaluation struct {
	studentID   uuid.UUID
	criterionID uuid.UUID
	score       Score
}

func (e Evaluation) StudentID() uuid.UUID {
	return e.studentID
}

func (e Evaluation) CriterionID() uuid.UUID {
	return e.criterionID
}

func (e Evaluation) Score() Score {
	return e.score
}

type EvaluationFactory struct{}

func (f EvaluationFactory) Create(student Student, rubric *Rubric, criterionID uuid.UUID, points int) (Evaluation, error) {
	if student.ID() == uuid.Nil || rubric == nil {
		return Evaluation{}, errors.New("нужны студент и рубрикатор")
	}
	criterion, exists := rubric.Criterion(criterionID)
	if !exists {
		return Evaluation{}, fmt.Errorf("критерия %q нет в рубрикаторе", criterionID)
	}
	score, err := NewScore(points, criterion.MaxScore())
	if err != nil {
		return Evaluation{}, err
	}
	return Evaluation{
		studentID:   student.ID(),
		criterionID: criterionID,
		score:       score,
	}, nil
}
