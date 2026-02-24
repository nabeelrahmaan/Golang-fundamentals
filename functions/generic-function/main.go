package main

import "fmt"
func Print[T any](val T){
	fmt.Println(val)
}

func main(){
Print(15)
Print(12.02)
Print(true)
Print("Nabeel")
}