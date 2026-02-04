package main

import (
	"fmt"
	"sync"
)

func printOdd(oddCh, evenCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 9; i += 2 {
		<-oddCh
		fmt.Println("Odd:", i)
		evenCh <- i
	}
}

func printEven(oddCh, evenCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 2; i <= 10; i += 2 {
		<-evenCh
		fmt.Println("Even:", i)

		if i < 10 {   
			oddCh <- i
		}
	}
}

func main() {
	oddCh := make(chan int)
	evenCh := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go printOdd(oddCh, evenCh, &wg)
	go printEven(oddCh, evenCh, &wg)

	oddCh <- 1 

	wg.Wait()
	fmt.Println("Done")
}
