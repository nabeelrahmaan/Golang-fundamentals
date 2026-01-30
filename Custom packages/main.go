package main

import (
	"Custom-packages/arithmatic"
	"fmt"
)

func main() {
	a := 20
	b := 5

	fmt.Println("Add:", arithmatic.Add(a, b))
	fmt.Println("Subtract:", arithmatic.Subtract(a, b))
	fmt.Println("Multiply:", arithmatic.Multiply(a, b))
	result, err := arithmatic.Devide(a, b)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Devide:", result)
}
