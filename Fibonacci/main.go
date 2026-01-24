package main

import "fmt"

func main() {
	var limit int
	a, b := 0, 1
	fmt.Print("Enter the limit:")
	fmt.Scan(&limit)
	for n := 0; n < limit; n++ {
		fmt.Println(a)
		a, b = b, a+b
	}
}
