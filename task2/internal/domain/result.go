package domain

type Result struct {
	percent   float64
	completed int
	total     int
}

func NewResult(percent float64, completed, total int) Result {
	return Result{
		percent:   percent,
		completed: completed,
		total:     total,
	}
}

func (r Result) Percent() float64 {
	return r.percent
}

func (r Result) Completed() int {
	return r.completed
}

func (r Result) Total() int {
	return r.total
}
