package app

import (
	"fmt"
	"pos/internal/domain"
)

type Service struct {
	students    StudentRepository
	rubrics     RubricRepository
	evaluations EvaluationRepository
	exporter    ResultExporter
	calculator  *domain.Calculator
	factory     domain.EvaluationFactory
}

func NewService(students StudentRepository, rubrics RubricRepository, evaluations EvaluationRepository, exporter ResultExporter, calculator *domain.Calculator) *Service {
	return &Service{
		students:    students,
		rubrics:     rubrics,
		evaluations: evaluations,
		exporter:    exporter,
		calculator:  calculator,
	}
}

func (s *Service) AddStudent(id, name string) error {
	student, err := domain.NewStudent(id, name)
	if err != nil {
		return err
	}
	if _, exists := s.students.Get(student.GetID()); exists {
		return fmt.Errorf("студент %q уже существует", student.GetID())
	}
	s.students.Save(student)
	return nil
}

func (s *Service) RenameStudent(id, name string) error {
	student, exists := s.students.Get(id)
	if !exists {
		return fmt.Errorf("студент %q не найден", id)
	}
	if err := student.Rename(name); err != nil {
		return err
	}
	s.students.Save(student)
	return nil
}

func (s *Service) RemoveStudent(id string) error {
	if _, exists := s.students.Get(id); !exists {
		return fmt.Errorf("студент %q не найден", id)
	}
	s.evaluations.DeleteByStudent(id)
	s.students.Delete(id)
	return nil
}

func (s *Service) Students() []domain.Student {
	return s.students.List()
}

func (s *Service) AddCriterion(id, name string, maxScore int, weight float64) error {
	criterion, err := domain.NewCriterion(id, name, maxScore, weight)
	if err != nil {
		return err
	}
	rubric := s.rubrics.GetRubric()
	if err := rubric.Add(criterion); err != nil {
		return err
	}
	s.rubrics.SaveRubric(rubric)
	return nil
}

func (s *Service) UpdateCriterion(id, name string, maxScore int, weight float64) error {
	criterion, err := domain.NewCriterion(id, name, maxScore, weight)
	if err != nil {
		return err
	}
	for _, evaluation := range s.evaluations.ListByCriterion(criterion.GetID()) {
		if evaluation.GetScore().GetPoints() > maxScore {
			return fmt.Errorf("по критерию %q есть оценка выше нового максимума", id)
		}
	}
	rubric := s.rubrics.GetRubric()
	if err := rubric.Replace(criterion); err != nil {
		return err
	}
	s.rubrics.SaveRubric(rubric)
	return nil
}

func (s *Service) RemoveCriterion(id string) error {
	if len(s.evaluations.ListByCriterion(id)) != 0 {
		return fmt.Errorf("по критерию %q ещё есть оценки", id)
	}
	rubric := s.rubrics.GetRubric()
	if err := rubric.Remove(id); err != nil {
		return err
	}
	s.rubrics.SaveRubric(rubric)
	return nil
}

func (s *Service) Criteria() []domain.Criterion {
	return s.rubrics.GetRubric().List()
}

func (s *Service) Grade(studentID, criterionID string, points int) error {
	student, exists := s.students.Get(studentID)
	if !exists {
		return fmt.Errorf("студент %q не найден", studentID)
	}
	evaluation, err := s.factory.Create(student, s.rubrics.GetRubric(), criterionID, points)
	if err != nil {
		return err
	}
	s.evaluations.SaveEvaluation(evaluation)
	return nil
}

func (s *Service) CancelGrade(studentID, criterionID string) error {
	if !s.evaluations.DeleteEvaluation(studentID, criterionID) {
		return fmt.Errorf("оценка студента %q по критерию %q не найдена", studentID, criterionID)
	}
	return nil
}

func (s *Service) Evaluations(studentID string) ([]domain.Evaluation, error) {
	if _, exists := s.students.Get(studentID); !exists {
		return nil, fmt.Errorf("студент %q не найден", studentID)
	}
	return s.evaluations.ListByStudent(studentID), nil
}

func (s *Service) Result(studentID string) (domain.Result, error) {
	student, exists := s.students.Get(studentID)
	if !exists {
		return domain.Result{}, fmt.Errorf("студент %q не найден", studentID)
	}
	return s.calculator.Calculate(student, s.rubrics.GetRubric(), s.evaluations.ListByStudent(studentID))
}

func (s *Service) ExportResult(studentID, path string) error {
	student, exists := s.students.Get(studentID)
	if !exists {
		return fmt.Errorf("студент %q не найден", studentID)
	}
	result, err := s.Result(studentID)
	if err != nil {
		return err
	}
	return s.exporter.Export(student, result, path)
}
