package main

import (
	"errors"
	"fmt"
)

func devide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Devision by zero not allowed")
	}
	return x / y, nil
}

func main() {
	var a int
	var b int
	fmt.Println("Enter the Divident:")
	fmt.Scan(&a)
	fmt.Println("Enter the Devisor:")
	fmt.Scan(&b)

	result, err := devide(a, b)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}
