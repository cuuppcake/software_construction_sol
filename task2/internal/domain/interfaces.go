package domain

type ResultPolicy interface {
	Percent(criteria []Criterion, scores map[string]Score) float64
}
