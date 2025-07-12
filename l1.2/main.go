package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	wg.Add(5)

	for _, v := range arr {
		go squareNum(v, &wg)
	}

	wg.Wait()
}

func squareNum(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%d: %d\n", n, n*n)
}
