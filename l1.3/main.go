package main

import (
	"fmt"
	"time"
)

func main() {
	var N, counter int

	fmt.Print("Enter the number of workers: ")
	fmt.Scan(&N)

	ch := make(chan string, 100)
	for i := 0; i < N; i++ {
		go worker(i, ch)
	}

	for {
		time.Sleep(500 * time.Millisecond)
		ch <- fmt.Sprintf("number: %v", counter)
		counter++
	}
}

func worker(id int, message <-chan string) {
	for m := range message {
		fmt.Printf("Worker %d: got message %s\n", id, m)
	}
}
