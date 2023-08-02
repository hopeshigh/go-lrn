package fizzbuzz

import (
	"errors"
	"fmt"
)

func Fizzbuzz(input int) (string, error) {
	if input <= 0 {
		return "", errors.New("input must be a positive integer")
	}

	switch {
	case input%15 == 0:
		return "FizzBuzz", nil
	case input%5 == 0:
		return "Buzz", nil
	case input%3 == 0:
		return "Fizz", nil
	default:
		return fmt.Sprintf("%d", input), nil
	}
}
