package calculator

import "errors"

//define struct
type Analytic struct {
	TotalOperations int
}

//pointer receiver method to increment the total operations
func NewAnalytic() *Analytic {
	return &Analytic{
		TotalOperations: 0,
	}
}

func (a *Analytic) CalculateBonus(score int) (int, error) {
	if score <= 0 {
		return 0, errors.New("score must be greater than 0")
	}
	a.TotalOperations++
	bonus := score * 10
	return bonus, nil
}
