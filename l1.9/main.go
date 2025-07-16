package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	producer, consumer := make(chan int), make(chan int)
	arr := []int{1, 2, 3, 4, 5}

	go func(producer <-chan int, consumer chan<- int) {
		for v := range producer {
			consumer <- v * 2
			wg.Done()
		}
		close(consumer)
	}(producer, consumer)

	go func() {
		for _, v := range arr {
			wg.Add(1)
			producer <- v
		}
		close(producer)
	}()

	for v := range consumer {
		fmt.Println(v)
	}

	wg.Wait()
}
