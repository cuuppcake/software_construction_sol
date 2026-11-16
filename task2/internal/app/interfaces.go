package app

import "pos/internal/domain"

type StudentRepository interface {
	Get(id string) (domain.Student, bool)
	List() []domain.Student
	Save(student domain.Student)
	Delete(id string)
}

type RubricRepository interface {
	GetRubric() *domain.Rubric
	SaveRubric(rubric *domain.Rubric)
}

type EvaluationRepository interface {
	ListByStudent(studentID string) []domain.Evaluation
	ListByCriterion(criterionID string) []domain.Evaluation
	SaveEvaluation(evaluation domain.Evaluation)
	DeleteEvaluation(studentID, criterionID string) bool
	DeleteByStudent(studentID string)
}

type ResultExporter interface {
	Export(student domain.Student, result domain.Result, path string) error
}
