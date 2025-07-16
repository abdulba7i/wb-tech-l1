package main

import "fmt"

func main() {
	arr := []int{3, 24, 234, 534, 61, 3, 35, 6, 8}
	sorted := quickSort(arr)
	fmt.Println(sorted)
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := len(arr) / 2 // в качестве опорного элемента возьмем середину

	var less, greater []int // создадим два слайса, в которых будем хранить все элементы меньше опорного, и больше

	for i, v := range arr {
		if i == pivot {
			continue // опорный элемент в учёт не берём
		}

		if v <= arr[pivot] {

			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	// result := append(quickSort(less), pivotValue)
	result := append(quickSort(less), arr[pivot])

	result = append(result, quickSort(greater)...)
	return result
}
