package main

import "fmt"

func numb(a, b int) {
	fmt.Println("Sum:", a+b)
	fmt.Println("Difference:", a-b)
	fmt.Println("Product:", a*b)
	if b == 0 {
		fmt.Println("Quotient: Error! cant devide a number by zero")
	} else {
		fmt.Println("Quotient:", a/b)
	}
}

func main() {
	var x int
	var y int
	fmt.Println("enter two numbers")
	fmt.Scan(&x)
	fmt.Scan(&y)
	numb(x, y)
}
