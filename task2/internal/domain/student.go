package domain

import (
	"errors"
	"strings"
)

type Student struct {
	id   string
	name string
}

func NewStudent(id, name string) (Student, error) {
	if strings.TrimSpace(id) == "" || strings.TrimSpace(name) == "" {
		return Student{}, errors.New("ID и имя студента не должны быть пустыми")
	}
	return Student{id: strings.TrimSpace(id), name: strings.TrimSpace(name)}, nil
}

func (s Student) GetID() string {
	return s.id
}

func (s Student) GetName() string {
	return s.name
}

func (s *Student) Rename(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("имя студента не должно быть пустым")
	}
	s.name = strings.TrimSpace(name)
	return nil
}
