package storage

import (
	"pos/internal/domain"

	"github.com/google/uuid"
)

func (m *MemoryStore) ListByStudent(studentID uuid.UUID) []domain.Evaluation {
	var result []domain.Evaluation
	for _, key := range m.gradeOrder {
		evaluation := m.evaluations[key]
		if evaluation.StudentID() == studentID {
			result = append(result, evaluation)
		}
	}
	return result
}

func (m *MemoryStore) ListByCriterion(criterionID uuid.UUID) []domain.Evaluation {
	var result []domain.Evaluation
	for _, key := range m.gradeOrder {
		evaluation := m.evaluations[key]
		if evaluation.CriterionID() == criterionID {
			result = append(result, evaluation)
		}
	}
	return result
}

func (m *MemoryStore) SaveEvaluation(evaluation domain.Evaluation) {
	key := gradeKey{
		studentID:   evaluation.StudentID(),
		criterionID: evaluation.CriterionID(),
	}
	if _, exists := m.evaluations[key]; !exists {
		m.gradeOrder = append(m.gradeOrder, key)
	}
	m.evaluations[key] = evaluation
}

func (m *MemoryStore) DeleteEvaluation(studentID, criterionID uuid.UUID) bool {
	key := gradeKey{
		studentID:   studentID,
		criterionID: criterionID,
	}
	if _, exists := m.evaluations[key]; !exists {
		return false
	}
	delete(m.evaluations, key)
	for i, current := range m.gradeOrder {
		if current == key {
			m.gradeOrder = append(m.gradeOrder[:i], m.gradeOrder[i+1:]...)
			break
		}
	}
	return true
}

func (m *MemoryStore) DeleteByStudent(studentID uuid.UUID) {
	for _, evaluation := range m.ListByStudent(studentID) {
		m.DeleteEvaluation(studentID, evaluation.CriterionID())
	}
}
