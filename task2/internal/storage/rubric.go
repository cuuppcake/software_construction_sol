package storage

import "pos/internal/domain"

func (m *MemoryStore) Rubric() *domain.Rubric {
	return m.rubric.Clone()
}

func (m *MemoryStore) SaveRubric(rubric *domain.Rubric) {
	m.rubric = rubric.Clone()
}
