package calci

import "errors"

func Add(num1 int, num2 int) int {
	return num1 + num2
}

func Substract(num1 int, num2 int) int {
	return num1 - num2
}

func Multiply(num1, num2 int) int {
	return num1 * num2
}

func Divide(num1, num2 int) (int, error) {

	if num2 == 0 {
		return 0, errors.New("Divide by Zero")
	}

	return num1 / num2, nil
}
