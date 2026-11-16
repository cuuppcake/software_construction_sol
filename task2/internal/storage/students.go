package storage

import (
	"pos/internal/domain"

	"github.com/google/uuid"
)

func (m *MemoryStore) Get(id uuid.UUID) (domain.Student, bool) {
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
	if _, exists := m.students[student.ID()]; !exists {
		m.studentOrder = append(m.studentOrder, student.ID())
	}
	m.students[student.ID()] = student
}

func (m *MemoryStore) Delete(id uuid.UUID) {
	delete(m.students, id)
	for i, current := range m.studentOrder {
		if current == id {
			m.studentOrder = append(m.studentOrder[:i], m.studentOrder[i+1:]...)
			break
		}
	}
}
