package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

type SillyNephewError struct {
}

func (s *SillyNephewError) Error(i int) string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", i)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0, errors.New("division by zero")
	}

	fodderAmount, fodderError := weightFodder.FodderAmount()

	if fodderAmount < 0 {
		if (fodderError == ErrScaleMalfunction) || (fodderError == nil) {
			return 0, errors.New("negative fodder")
		}
	}

	if fodderError != nil {
		if fodderError == ErrScaleMalfunction {
			if fodderAmount >= 0 {
				return fodderAmount * 2 / float64(cows), nil
			}
		} else {
			return 0, fodderError
		}
	}

	if cows < 0 {
		var snr = &SillyNephewError{}
		return 0.0, errors.New(snr.Error(cows))
	}

	return fodderAmount / float64(cows), nil
}
