// This program demonstrates how channels coordinate communication
// between goroutines to control their execution order.
//
// Concurrency is not just random switching between goroutines.
// It can be synchronized and controlled using channels and
// synchronization tools from the sync package.


package main

import (
	"fmt"
	"sync"
)

func oddNumbers(wg *sync.WaitGroup, oddChan <-chan bool, evenChan chan<- bool) {  //The declaration of send-only & receive-only of channels is not compulsory, its purely for type safety and learning purpose.
	defer wg.Done()

	for i := 1; i <= 20; i += 2 {
		<-oddChan
		fmt.Println("Odd:", i)
		evenChan <- true

	}
}

func evenNumbers(wg *sync.WaitGroup, oddChan chan<- bool, evenChan <-chan bool) {
	defer wg.Done()

	for i := 2; i <= 20; i += 2 {
		<-evenChan
		fmt.Println("even:", i)
		if i < 20 {
			oddChan <- true
		}
	}
}

func main() {
	var wg sync.WaitGroup

	oddChan := make(chan bool)
	evenChan := make(chan bool)

	wg.Add(2)
	go oddNumbers(&wg, oddChan, evenChan)
	go evenNumbers(&wg, oddChan, evenChan)

	oddChan <- true
	wg.Wait()
}
