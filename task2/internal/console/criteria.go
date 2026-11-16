package console

import (
	"fmt"

	"github.com/google/uuid"
)

func (c *Console) saveCriterion(update bool) error {
	var id uuid.UUID
	var err error
	if update {
		id, err = c.id("ID критерия")
		if err != nil {
			return err
		}
	}
	name, err := c.prompt("Название")
	if err != nil {
		return err
	}
	maxScore, err := c.integer("Максимальный балл")
	if err != nil {
		return err
	}
	weight, err := c.integer("Вес (например, 50)")
	if err != nil {
		return err
	}
	if update {
		return c.service.UpdateCriterion(id, name, maxScore, weight)
	}
	id, err = c.service.AddCriterion(name, maxScore, weight)
	if err != nil {
		return err
	}
	return c.println("ID критерия:", id)
}

func (c *Console) showCriteria() error {
	criteria := c.service.Criteria()
	if len(criteria) == 0 {
		return c.println("Критериев пока нет.")
	}
	for _, criterion := range criteria {
		if _, err := fmt.Fprintf(c.out, "%s — %s, максимум %d, вес %d\n", criterion.ID(), criterion.Name(), criterion.MaxScore(), criterion.Weight()); err != nil {
			return err
		}
	}
	return nil
}

func (c *Console) removeCriterion() error {
	id, err := c.id("ID критерия")
	if err != nil {
		return err
	}
	return c.service.RemoveCriterion(id)
}
