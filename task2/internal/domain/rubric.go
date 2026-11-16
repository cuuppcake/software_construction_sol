package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type Rubric struct {
	criteria map[uuid.UUID]Criterion
	order    []uuid.UUID
}

func NewRubric() *Rubric {
	return &Rubric{criteria: make(map[uuid.UUID]Criterion)}
}

func (r *Rubric) Add(criterion Criterion) error {
	if err := criterion.valid(); err != nil {
		return err
	}
	if _, exists := r.criteria[criterion.ID()]; exists {
		return fmt.Errorf("критерий %q уже существует", criterion.ID())
	}
	r.criteria[criterion.ID()] = criterion
	r.order = append(r.order, criterion.ID())
	return nil
}

func (r *Rubric) Replace(criterion Criterion) error {
	if err := criterion.valid(); err != nil {
		return err
	}
	if _, exists := r.criteria[criterion.ID()]; !exists {
		return fmt.Errorf("критерий %q не найден", criterion.ID())
	}
	r.criteria[criterion.ID()] = criterion
	return nil
}

func (r *Rubric) Remove(id uuid.UUID) error {
	if _, exists := r.criteria[id]; !exists {
		return fmt.Errorf("критерий %q не найден", id)
	}
	delete(r.criteria, id)
	for i, current := range r.order {
		if current == id {
			r.order = append(r.order[:i], r.order[i+1:]...)
			break
		}
	}
	return nil
}

func (r *Rubric) Criterion(id uuid.UUID) (Criterion, bool) {
	criterion, exists := r.criteria[id]
	return criterion, exists
}

func (r *Rubric) List() []Criterion {
	result := make([]Criterion, 0, len(r.order))
	for _, id := range r.order {
		result = append(result, r.criteria[id])
	}
	return result
}

func (r *Rubric) Clone() *Rubric {
	clone := NewRubric()
	clone.order = append(clone.order, r.order...)
	for id, criterion := range r.criteria {
		clone.criteria[id] = criterion
	}
	return clone
}
