package console

import "fmt"

func (c *Console) saveCriterion(update bool) error {
	id, err := c.prompt("ID критерия")
	if err != nil {
		return err
	}
	name, err := c.prompt("Название")
	if err != nil {
		return err
	}
	maxScore, err := c.integer("Максимальный балл")
	if err != nil {
		return err
	}
	weight, err := c.decimal("Вес (например, 0.5)")
	if err != nil {
		return err
	}
	if update {
		return c.service.UpdateCriterion(id, name, maxScore, weight)
	}
	return c.service.AddCriterion(id, name, maxScore, weight)
}

func (c *Console) showCriteria() {
	criteria := c.service.Criteria()
	if len(criteria) == 0 {
		fmt.Fprintln(c.out, "Критериев пока нет.")
	}
	for _, criterion := range criteria {
		fmt.Fprintf(c.out, "%s — %s, максимум %d, вес %.2f\n", criterion.GetID(), criterion.GetName(), criterion.GetMaxScore(), criterion.GetWeight())
	}
}

func (c *Console) removeCriterion() error {
	id, err := c.prompt("ID критерия")
	if err != nil {
		return err
	}
	return c.service.RemoveCriterion(id)
}
