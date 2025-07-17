package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Atomics struct {
	ops atomic.Uint64
	wg  sync.WaitGroup
}

func main() {
	var a Atomics

	for i := 0; i < 100; i++ {
		a.wg.Add(1)

		go func() {
			for i := 0; i < 1000; i++ {
				a.ops.Add(1)
			}
			a.wg.Done()
		}()
	}
	a.wg.Wait()

	fmt.Println(a.ops.Load())
}
