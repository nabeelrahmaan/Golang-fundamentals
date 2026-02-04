package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(jobs chan<- int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Println("Producing Job:", i)
		jobs <- i 
		time.Sleep(200 * time.Millisecond)
	}
	close(jobs) 
}
func consumer(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs { 
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	jobs := make(chan int, 5) 

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go consumer(i, jobs, &wg)
	}
	go producer(jobs, 10)

	wg.Wait()
	fmt.Println("All jobs processed")
}
