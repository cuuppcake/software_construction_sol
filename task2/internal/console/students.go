package console

func (c *Console) addStudent() error {
	name, err := c.prompt("Имя студента")
	if err != nil {
		return err
	}
	id, err := c.service.AddStudent(name)
	if err != nil {
		return err
	}
	return c.println("ID студента:", id)
}

func (c *Console) showStudents() error {
	students := c.service.Students()
	if len(students) == 0 {
		return c.println("Студентов пока нет.")
	}
	for _, student := range students {
		if err := c.printf("%s — %s\n", student.ID(), student.Name()); err != nil {
			return err
		}
	}
	return nil
}

func (c *Console) renameStudent() error {
	id, err := c.id("ID студента")
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
	id, err := c.id("ID студента")
	if err != nil {
		return err
	}
	return c.service.RemoveStudent(id)
}
