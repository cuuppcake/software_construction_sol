package console

import "fmt"

func (c *Console) grade() error {
	studentID, err := c.id("ID студента")
	if err != nil {
		return err
	}
	criterionID, err := c.id("ID критерия")
	if err != nil {
		return err
	}
	points, err := c.integer("Балл")
	if err != nil {
		return err
	}
	return c.service.Grade(studentID, criterionID, points)
}

func (c *Console) cancelGrade() error {
	studentID, err := c.id("ID студента")
	if err != nil {
		return err
	}
	criterionID, err := c.id("ID критерия")
	if err != nil {
		return err
	}
	return c.service.CancelGrade(studentID, criterionID)
}

func (c *Console) showEvaluations() error {
	studentID, err := c.id("ID студента")
	if err != nil {
		return err
	}
	evaluations, err := c.service.Evaluations(studentID)
	if err != nil {
		return err
	}
	if len(evaluations) == 0 {
		return c.println("У студента пока нет оценок.")
	}
	for _, evaluation := range evaluations {
		if err := c.printf("%s — %d\n", evaluation.CriterionID(), evaluation.Score().Points()); err != nil {
			return err
		}
	}
	return nil
}

func (c *Console) showResult() error {
	studentID, err := c.id("ID студента")
	if err != nil {
		return err
	}
	result, err := c.service.Result(studentID)
	if err != nil {
		return err
	}
	return c.printf("Итог: %.2f%%, оценено %d из %d критериев\n", result.Percent(), result.Completed(), result.Total())
}

func (c *Console) exportResult() error {
	studentID, err := c.id("ID студента")
	if err != nil {
		return err
	}
	path, err := c.prompt("Путь к CSV")
	if err != nil {
		return err
	}
	if path == "" {
		return fmt.Errorf("путь к CSV-файлу не должен быть пустым")
	}
	if err := c.service.ExportResult(studentID, path); err != nil {
		return err
	}
	return c.println("Экспортировано в", path)
}
