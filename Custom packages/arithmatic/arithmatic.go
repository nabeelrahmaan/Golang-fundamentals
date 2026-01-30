package arithmatic

import "errors"

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Devide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cant devide by zero")
	}
	return a/b, nil;
}