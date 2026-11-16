package storage

import (
	"pos/internal/domain"

	"github.com/google/uuid"
)

type MemoryStore struct {
	students     map[uuid.UUID]domain.Student
	studentOrder []uuid.UUID
	rubric       *domain.Rubric
	evaluations  map[gradeKey]domain.Evaluation
	gradeOrder   []gradeKey
}

type gradeKey struct {
	studentID   uuid.UUID
	criterionID uuid.UUID
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		students:    make(map[uuid.UUID]domain.Student),
		rubric:      domain.NewRubric(),
		evaluations: make(map[gradeKey]domain.Evaluation),
	}
}
