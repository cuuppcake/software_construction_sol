package domain

import "fmt"

type Rubric struct {
	criteria map[string]Criterion
	order    []string
}

func NewRubric() *Rubric {
	return &Rubric{criteria: make(map[string]Criterion)}
}

func (r *Rubric) Add(c Criterion) error {
	criterion, err := NewCriterion(c.GetID(), c.GetName(), c.GetMaxScore(), c.GetWeight())
	if err != nil {
		return err
	}
	if _, exists := r.criteria[criterion.GetID()]; exists {
		return fmt.Errorf("критерий %q уже существует", criterion.GetID())
	}
	r.criteria[criterion.GetID()] = criterion
	r.order = append(r.order, criterion.GetID())
	return nil
}

func (r *Rubric) Replace(c Criterion) error {
	criterion, err := NewCriterion(c.GetID(), c.GetName(), c.GetMaxScore(), c.GetWeight())
	if err != nil {
		return err
	}
	if _, exists := r.criteria[criterion.GetID()]; !exists {
		return fmt.Errorf("критерий %q не найден", criterion.GetID())
	}
	r.criteria[criterion.GetID()] = criterion
	return nil
}

func (r *Rubric) Remove(id string) error {
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

func (r *Rubric) Criterion(id string) (Criterion, bool) {
	c, ok := r.criteria[id]
	return c, ok
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
	for id, c := range r.criteria {
		clone.criteria[id] = c
	}
	return clone
}
