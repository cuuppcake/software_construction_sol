package storage

import "pos/internal/domain"

type MemoryStore struct {
	students     map[string]domain.Student
	studentOrder []string
	rubric       *domain.Rubric
	evaluations  map[gradeKey]domain.Evaluation
	gradeOrder   []gradeKey
}

type gradeKey struct {
	studentID   string
	criterionID string
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		students:    make(map[string]domain.Student),
		rubric:      domain.NewRubric(),
		evaluations: make(map[gradeKey]domain.Evaluation),
	}
}

func (m *MemoryStore) Get(id string) (domain.Student, bool) {
	student, exists := m.students[id]
	return student, exists
}

func (m *MemoryStore) List() []domain.Student {
	result := make([]domain.Student, 0, len(m.studentOrder))
	for _, id := range m.studentOrder {
		result = append(result, m.students[id])
	}
	return result
}

func (m *MemoryStore) Save(student domain.Student) {
	if _, exists := m.students[student.GetID()]; !exists {
		m.studentOrder = append(m.studentOrder, student.GetID())
	}
	m.students[student.GetID()] = student
}

func (m *MemoryStore) Delete(id string) {
	delete(m.students, id)
	for i, current := range m.studentOrder {
		if current == id {
			m.studentOrder = append(m.studentOrder[:i], m.studentOrder[i+1:]...)
			break
		}
	}
}

func (m *MemoryStore) GetRubric() *domain.Rubric {
	return m.rubric.Clone()
}

func (m *MemoryStore) SaveRubric(rubric *domain.Rubric) {
	m.rubric = rubric.Clone()
}

func (m *MemoryStore) ListByStudent(studentID string) []domain.Evaluation {
	var result []domain.Evaluation
	for _, key := range m.gradeOrder {
		evaluation := m.evaluations[key]
		if evaluation.GetStudentID() == studentID {
			result = append(result, evaluation)
		}
	}
	return result
}

func (m *MemoryStore) ListByCriterion(criterionID string) []domain.Evaluation {
	var result []domain.Evaluation
	for _, key := range m.gradeOrder {
		evaluation := m.evaluations[key]
		if evaluation.GetCriterionID() == criterionID {
			result = append(result, evaluation)
		}
	}
	return result
}

func (m *MemoryStore) SaveEvaluation(evaluation domain.Evaluation) {
	key := gradeKey{studentID: evaluation.GetStudentID(), criterionID: evaluation.GetCriterionID()}
	if _, exists := m.evaluations[key]; !exists {
		m.gradeOrder = append(m.gradeOrder, key)
	}
	m.evaluations[key] = evaluation
}

func (m *MemoryStore) DeleteEvaluation(studentID, criterionID string) bool {
	key := gradeKey{studentID: studentID, criterionID: criterionID}
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

func (m *MemoryStore) DeleteByStudent(studentID string) {
	for _, evaluation := range m.ListByStudent(studentID) {
		m.DeleteEvaluation(studentID, evaluation.GetCriterionID())
	}
}
