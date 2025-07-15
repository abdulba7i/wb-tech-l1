package main

import (
	"fmt"
	"sync"
)

type MapWork struct {
	model map[int]string
	mu    sync.Mutex
}

func (m *MapWork) DataEntry(key int, value string) {
	m.mu.Lock()
	m.model[key] = value
	m.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	m := MapWork{
		model: make(map[int]string),
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			msg := fmt.Sprintf("msg #%d", i)
			m.DataEntry(i, msg)
		}(i)
	}

	wg.Wait()

	fmt.Println(m.model)
}
