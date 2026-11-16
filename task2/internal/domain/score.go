package domain

import "fmt"

type Score struct {
	points int
}

func NewScore(points, maxScore int) (Score, error) {
	if maxScore <= 0 || points < 0 || points > maxScore {
		return Score{}, fmt.Errorf("балл %d должен быть в диапазоне от 0 до %d", points, maxScore)
	}
	return Score{points: points}, nil
}

func (s Score) Points() int {
	return s.points
}
