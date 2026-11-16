package console

import "fmt"

func (c *Console) addStudent() error {
	id, err := c.prompt("ID студента")
	if err != nil {
		return err
	}
	name, err := c.prompt("Имя")
	if err != nil {
		return err
	}
	return c.service.AddStudent(id, name)
}

func (c *Console) showStudents() {
	students := c.service.Students()
	if len(students) == 0 {
		fmt.Fprintln(c.out, "Студентов пока нет.")
	}
	for _, student := range students {
		fmt.Fprintf(c.out, "%s — %s\n", student.GetID(), student.GetName())
	}
}

func (c *Console) renameStudent() error {
	id, err := c.prompt("ID студента")
	if err != nil {
		return err
	}
	name, err := c.prompt("Новое имя")
	if err != nil {
		return err
	}
	return c.service.RenameStudent(id, name)
}

func (c *Console) removeStudent() error {
	id, err := c.prompt("ID студента")
	if err != nil {
		return err
	}
	return c.service.RemoveStudent(id)
}
