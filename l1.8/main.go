package main

import (
	"fmt"
)

func UpdateBit(n int64, i int, bit int) int64 {
	if bit == 1 {
		n |= 1 << i
	} else {
		n &^= 1 << i
	}
	return n
}

func main() {
	var num int64 = 12
	var i int = 1
	var bit = 1
	fmt.Printf("Исходное число %d: %b\n", num, num)

	newNum := UpdateBit(num, i, bit)

	fmt.Printf("Результат %d: %b\n", newNum, newNum)

}
