package domain

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type Student struct {
	id   uuid.UUID
	name string
}

func NewStudent(name string) (Student, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return Student{}, errors.New("имя студента не должно быть пустым")
	}
	return Student{
		id:   uuid.New(),
		name: name,
	}, nil
}

func (s Student) ID() uuid.UUID {
	return s.id
}

func (s Student) Name() string {
	return s.name
}

func (s *Student) Rename(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("имя студента не должно быть пустым")
	}
	s.name = name
	return nil
}
