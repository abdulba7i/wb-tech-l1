package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a, b := rand.Intn(100), rand.Intn(100)
	fmt.Printf("Исходные числа: %d | %d\n", a, b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("Результат: %d | %d\n", a, b)
}
