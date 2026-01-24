package main

import "fmt"

func main(){
	var num1 int
	var num2 int
	var operation int
	var result int

	fmt.Println("Enter the first number:")
	fmt.Scan(&num1)
	fmt.Println("Enter the second number:")
	fmt.Scan(&num2)
	fmt.Println("1.Addition\n2.substraction\n3.multiplication\n4.division\nSelect operation:")
	fmt.Scan(&operation)
	switch operation{
	case 1:
		result=num1+num2
	case 2:
		result=num1-num2
	case 3:
		result=num1*num2
	case 4:
		result=num1/num2
	default:
		fmt.Println("Invalid operation")
	}
fmt.Println("Result:",result)
}