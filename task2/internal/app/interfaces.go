package app

import (
	"pos/internal/domain"

	"github.com/google/uuid"
)

type StudentRepository interface {
	Get(id uuid.UUID) (domain.Student, bool)
	List() []domain.Student
	Save(student domain.Student)
	Delete(id uuid.UUID)
}

type RubricRepository interface {
	Rubric() *domain.Rubric
	SaveRubric(rubric *domain.Rubric)
}

type EvaluationRepository interface {
	ListByStudent(studentID uuid.UUID) []domain.Evaluation
	ListByCriterion(criterionID uuid.UUID) []domain.Evaluation
	SaveEvaluation(evaluation domain.Evaluation)
	DeleteEvaluation(studentID, criterionID uuid.UUID) bool
	DeleteByStudent(studentID uuid.UUID)
}

type ResultExporter interface {
	Export(student domain.Student, result domain.Result, path string) error
}
