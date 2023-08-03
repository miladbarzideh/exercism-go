package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(foodCalculator FodderCalculator, cows int) (float64, error) {
	amount, err := foodCalculator.FodderAmount(cows)
	if err != nil {
		return 0, err
	}
	fatteningFactor, err := foodCalculator.FatteningFactor()
	if err != nil {
		return 0, err
	}
	foodPerCow := (amount * fatteningFactor) / float64(cows)
	return foodPerCow, nil
}

func ValidateInputAndDivideFood(foodCalculator FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(foodCalculator, cows)
	}
	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	Message string
	Cows int
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.Cows, e.Message)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			Cows: cows,
			Message: "there are no negative cows",
		}
	}
	if cows == 0 {
		return &InvalidCowsError{
			Cows: cows,
			Message: "no cows don't need food",
		}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
