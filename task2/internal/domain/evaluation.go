package domain

import "fmt"

type Evaluation struct {
	studentID   string
	criterionID string
	score       Score
}

func (e Evaluation) GetStudentID() string {
	return e.studentID
}

func (e Evaluation) GetCriterionID() string {
	return e.criterionID
}

func (e Evaluation) GetScore() Score {
	return e.score
}

type EvaluationFactory struct{}

func (f EvaluationFactory) Create(student Student, rubric *Rubric, criterionID string, points int) (Evaluation, error) {
	if student.GetID() == "" || rubric == nil {
		return Evaluation{}, fmt.Errorf("нужны студент и рубрикатор")
	}
	criterion, exists := rubric.Criterion(criterionID)
	if !exists {
		return Evaluation{}, fmt.Errorf("критерия %q нет в рубрикаторе", criterionID)
	}
	score, err := NewScore(points, criterion.GetMaxScore())
	if err != nil {
		return Evaluation{}, err
	}
	return Evaluation{studentID: student.GetID(), criterionID: criterionID, score: score}, nil
}
