package domain

import (
	"errors"
	"math"
	"strings"
)

type Criterion struct {
	id       string
	name     string
	maxScore int
	weight   float64
}

func NewCriterion(id, name string, maxScore int, weight float64) (Criterion, error) {
	if strings.TrimSpace(id) == "" || strings.TrimSpace(name) == "" {
		return Criterion{}, errors.New("ID и название критерия не должны быть пустыми")
	}
	if maxScore <= 0 {
		return Criterion{}, errors.New("максимальный балл должен быть положительным")
	}
	if weight <= 0 || math.IsNaN(weight) || math.IsInf(weight, 0) {
		return Criterion{}, errors.New("вес должен быть положительным конечным числом")
	}
	return Criterion{id: strings.TrimSpace(id), name: strings.TrimSpace(name), maxScore: maxScore, weight: weight}, nil
}

func (c Criterion) GetID() string {
	return c.id
}

func (c Criterion) GetName() string {
	return c.name
}

func (c Criterion) GetMaxScore() int {
	return c.maxScore
}

func (c Criterion) GetWeight() float64 {
	return c.weight
}
