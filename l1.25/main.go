package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()

	value := <-ch

	Sleep(1 * time.Second)

	fmt.Println(value)
}

func Sleep(t time.Duration) {
	<-time.After(t)
}
