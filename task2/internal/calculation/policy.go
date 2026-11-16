package calculation

import (
	"pos/internal/domain"

	"github.com/google/uuid"
)

type ResultPolicy interface {
	Percent(criteria []domain.Criterion, scores map[uuid.UUID]domain.Score) float64
}

type WeightedPolicy struct{}

func (p WeightedPolicy) Percent(criteria []domain.Criterion, scores map[uuid.UUID]domain.Score) float64 {
	var earned, totalWeight float64
	for _, criterion := range criteria {
		weight := float64(criterion.Weight())
		totalWeight += weight
		earned += float64(scores[criterion.ID()].Points()) / float64(criterion.MaxScore()) * weight
	}
	return earned / totalWeight * 100
}

type EqualPolicy struct{}

func (p EqualPolicy) Percent(criteria []domain.Criterion, scores map[uuid.UUID]domain.Score) float64 {
	var earned float64
	for _, criterion := range criteria {
		earned += float64(scores[criterion.ID()].Points()) / float64(criterion.MaxScore())
	}
	return earned / float64(len(criteria)) * 100
}
