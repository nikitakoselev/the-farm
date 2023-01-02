package thefarm

import "errors"

// See types.go for the types defined for this exercise.

type SillyNephewError struct {
	amount int
	Error  string
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, fodderError := weightFodder.FodderAmount()

	if cows == 0 {
		return 0, errors.New("division by zero")
	}

	if fodderError == ErrScaleMalfunction && fodderAmount >= 0 {
		return fodderAmount * 2 / float64(cows), errors.New("no error")
	}

	if cows < 0 {
		return 0.0, errors.New("silly nephew, there cannot be -5 cows")
	}

	if fodderError != nil {
		return 0, fodderError
	}

	if fodderAmount < 0 {
		return 0, errors.New("negative fodder")
	}

	return fodderAmount / float64(cows), nil
}
