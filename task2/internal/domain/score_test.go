package domain_test

import (
	"pos/internal/domain"
	"testing"
)

func TestScoreRange(t *testing.T) {
	tests := []struct {
		name     string
		points   int
		maxScore int
		valid    bool
	}{
		{
			name:     "отрицательный балл",
			points:   -1,
			maxScore: 10,
		},
		{
			name:     "ноль",
			points:   0,
			maxScore: 10,
			valid:    true,
		},
		{
			name:     "максимальный балл",
			points:   10,
			maxScore: 10,
			valid:    true,
		},
		{
			name:     "балл выше максимума",
			points:   11,
			maxScore: 10,
		},
		{
			name:     "нулевой максимум",
			points:   0,
			maxScore: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			score, err := domain.NewScore(test.points, test.maxScore)
			if (err == nil) != test.valid {
				t.Fatalf("неожиданная ошибка: %v", err)
			}
			if test.valid && score.Points() != test.points {
				t.Fatal("создан неверный балл")
			}
		})
	}
}
