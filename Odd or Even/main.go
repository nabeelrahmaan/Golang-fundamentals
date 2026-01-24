package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a number:")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("is a odd number")
	}
}
