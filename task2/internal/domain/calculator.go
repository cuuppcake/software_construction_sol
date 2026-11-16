package domain

import "fmt"

type Result struct {
	percent   float64
	completed int
	total     int
}

func (r Result) GetPercent() float64 {
	return r.percent
}

func (r Result) GetCompleted() int {
	return r.completed
}

func (r Result) GetTotal() int {
	return r.total
}

type Calculator struct {
	policy ResultPolicy
}

func NewCalculator(policy ResultPolicy) *Calculator {
	return &Calculator{policy: policy}
}

func (c *Calculator) Calculate(student Student, rubric *Rubric, evaluations []Evaluation) (Result, error) {
	if rubric == nil || c == nil || c.policy == nil {
		return Result{}, fmt.Errorf("нужны рубрикатор и правило расчёта итога")
	}
	if student.GetID() == "" {
		return Result{}, fmt.Errorf("нужен студент для расчёта итога")
	}
	criteria := rubric.List()
	if len(criteria) == 0 {
		return Result{}, fmt.Errorf("рубрикатор пуст")
	}
	scores := make(map[string]Score)
	for _, evaluation := range evaluations {
		if evaluation.GetStudentID() != student.GetID() {
			return Result{}, fmt.Errorf("оценка относится к другому студенту")
		}
		criterion, exists := rubric.Criterion(evaluation.GetCriterionID())
		if !exists {
			return Result{}, fmt.Errorf("критерия %q нет в рубрикаторе", evaluation.GetCriterionID())
		}
		if _, exists := scores[evaluation.GetCriterionID()]; exists {
			return Result{}, fmt.Errorf("по критерию %q найдены две оценки", evaluation.GetCriterionID())
		}
		score := evaluation.GetScore()
		if _, err := NewScore(score.GetPoints(), criterion.GetMaxScore()); err != nil {
			return Result{}, err
		}
		scores[evaluation.GetCriterionID()] = score
	}
	return Result{percent: c.policy.Percent(criteria, scores), completed: len(scores), total: len(criteria)}, nil
}

type WeightedPolicy struct{}

func (p WeightedPolicy) Percent(criteria []Criterion, scores map[string]Score) float64 {
	var earned, totalWeight float64
	for _, criterion := range criteria {
		weight := criterion.GetWeight()
		totalWeight += weight
		earned += float64(scores[criterion.GetID()].GetPoints()) / float64(criterion.GetMaxScore()) * weight
	}
	return earned / totalWeight * 100
}

type EqualPolicy struct{}

func (p EqualPolicy) Percent(criteria []Criterion, scores map[string]Score) float64 {
	var earned float64
	for _, criterion := range criteria {
		earned += float64(scores[criterion.GetID()].GetPoints()) / float64(criterion.GetMaxScore())
	}
	return earned / float64(len(criteria)) * 100
}
