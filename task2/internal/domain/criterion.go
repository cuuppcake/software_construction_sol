package domain

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

type Criterion struct {
	id       uuid.UUID
	name     string
	maxScore int
	weight   int
}

func NewCriterion(name string, maxScore, weight int) (Criterion, error) {
	name = strings.TrimSpace(name)
	if err := checkCriterion(name, maxScore, weight); err != nil {
		return Criterion{}, err
	}
	return Criterion{
		id:       uuid.New(),
		name:     name,
		maxScore: maxScore,
		weight:   weight,
	}, nil
}

func checkCriterion(name string, maxScore, weight int) error {
	if name == "" {
		return errors.New("название критерия не должно быть пустым")
	}
	if maxScore <= 0 {
		return errors.New("максимальный балл должен быть положительным")
	}
	if weight <= 0 {
		return errors.New("вес должен быть положительным")
	}
	return nil
}

func (c Criterion) valid() error {
	if c.id == uuid.Nil {
		return errors.New("критерий не создан")
	}
	return checkCriterion(c.name, c.maxScore, c.weight)
}

func (c *Criterion) Change(name string, maxScore, weight int) error {
	name = strings.TrimSpace(name)
	if err := checkCriterion(name, maxScore, weight); err != nil {
		return err
	}
	c.name = name
	c.maxScore = maxScore
	c.weight = weight
	return nil
}

func (c Criterion) ID() uuid.UUID {
	return c.id
}

func (c Criterion) Name() string {
	return c.name
}

func (c Criterion) MaxScore() int {
	return c.maxScore
}

func (c Criterion) Weight() int {
	return c.weight
}
