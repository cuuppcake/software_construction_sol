package app

import (
	"fmt"
	"pos/internal/calculation"
	"pos/internal/domain"

	"github.com/google/uuid"
)

type Service struct {
	students    StudentRepository
	rubrics     RubricRepository
	evaluations EvaluationRepository
	exporter    ResultExporter
	calculator  *calculation.Calculator
	factory     domain.EvaluationFactory
}

func NewService(students StudentRepository, rubrics RubricRepository, evaluations EvaluationRepository, exporter ResultExporter, calculator *calculation.Calculator) *Service {
	return &Service{
		students:    students,
		rubrics:     rubrics,
		evaluations: evaluations,
		exporter:    exporter,
		calculator:  calculator,
	}
}

func (s *Service) student(id uuid.UUID) (domain.Student, error) {
	student, exists := s.students.Get(id)
	if !exists {
		return domain.Student{}, fmt.Errorf("студент %q не найден", id)
	}
	return student, nil
}

func (s *Service) AddStudent(name string) (uuid.UUID, error) {
	student, err := domain.NewStudent(name)
	if err != nil {
		return uuid.Nil, err
	}
	s.students.Save(student)
	return student.ID(), nil
}

func (s *Service) RenameStudent(id uuid.UUID, name string) error {
	student, err := s.student(id)
	if err != nil {
		return err
	}
	if err := student.Rename(name); err != nil {
		return err
	}
	s.students.Save(student)
	return nil
}

func (s *Service) RemoveStudent(id uuid.UUID) error {
	if _, err := s.student(id); err != nil {
		return err
	}
	s.evaluations.DeleteByStudent(id)
	s.students.Delete(id)
	return nil
}

func (s *Service) Students() []domain.Student {
	return s.students.List()
}

func (s *Service) AddCriterion(name string, maxScore, weight int) (uuid.UUID, error) {
	criterion, err := domain.NewCriterion(name, maxScore, weight)
	if err != nil {
		return uuid.Nil, err
	}
	rubric := s.rubrics.Rubric()
	if err := rubric.Add(criterion); err != nil {
		return uuid.Nil, err
	}
	s.rubrics.SaveRubric(rubric)
	return criterion.ID(), nil
}

func (s *Service) UpdateCriterion(id uuid.UUID, name string, maxScore, weight int) error {
	rubric := s.rubrics.Rubric()
	criterion, exists := rubric.Criterion(id)
	if !exists {
		return fmt.Errorf("критерий %q не найден", id)
	}
	for _, evaluation := range s.evaluations.ListByCriterion(id) {
		if evaluation.Score().Points() > maxScore {
			return fmt.Errorf("по критерию %q есть оценка выше нового максимума", id)
		}
	}
	if err := criterion.Change(name, maxScore, weight); err != nil {
		return err
	}
	if err := rubric.Replace(criterion); err != nil {
		return err
	}
	s.rubrics.SaveRubric(rubric)
	return nil
}

func (s *Service) RemoveCriterion(id uuid.UUID) error {
	if len(s.evaluations.ListByCriterion(id)) != 0 {
		return fmt.Errorf("по критерию %q ещё есть оценки", id)
	}
	rubric := s.rubrics.Rubric()
	if err := rubric.Remove(id); err != nil {
		return err
	}
	s.rubrics.SaveRubric(rubric)
	return nil
}

func (s *Service) Criteria() []domain.Criterion {
	return s.rubrics.Rubric().List()
}

func (s *Service) Grade(studentID, criterionID uuid.UUID, points int) error {
	student, err := s.student(studentID)
	if err != nil {
		return err
	}
	evaluation, err := s.factory.Create(student, s.rubrics.Rubric(), criterionID, points)
	if err != nil {
		return err
	}
	s.evaluations.SaveEvaluation(evaluation)
	return nil
}

func (s *Service) CancelGrade(studentID, criterionID uuid.UUID) error {
	if !s.evaluations.DeleteEvaluation(studentID, criterionID) {
		return fmt.Errorf("оценка студента %q по критерию %q не найдена", studentID, criterionID)
	}
	return nil
}

func (s *Service) Evaluations(studentID uuid.UUID) ([]domain.Evaluation, error) {
	if _, err := s.student(studentID); err != nil {
		return nil, err
	}
	return s.evaluations.ListByStudent(studentID), nil
}

func (s *Service) resultFor(student domain.Student) (domain.Result, error) {
	return s.calculator.Calculate(student, s.rubrics.Rubric(), s.evaluations.ListByStudent(student.ID()))
}

func (s *Service) Result(studentID uuid.UUID) (domain.Result, error) {
	student, err := s.student(studentID)
	if err != nil {
		return domain.Result{}, err
	}
	return s.resultFor(student)
}

func (s *Service) ExportResult(studentID uuid.UUID, path string) error {
	student, err := s.student(studentID)
	if err != nil {
		return err
	}
	result, err := s.resultFor(student)
	if err != nil {
		return err
	}
	return s.exporter.Export(student, result, path)
}
