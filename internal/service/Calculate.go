package service

import (
	"errors"
)

func Calculate(num1 float64, num2 float64, op string) (float64, error) {
	var answer float64
	switch op {
	case "+":
		answer = num1 + num2
	case "-":
		answer = num1 - num2

	case "*":
		answer = num1 * num2
	case "/":
		if num2 == 0.0 {
			return 0, errors.New("Cannot Divide By Zero")
		} else {
			answer = num1 / num2

		}

	default:
		return 0, errors.New("Invalid operator")

	}
	return answer, nil
}
