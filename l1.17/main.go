package main

import "fmt"

func main() {
	arr := []int{1, 6, 10, 100, 134, 193, 200, 400}

	fmt.Println(binarySearch(arr, 193))
}

func binarySearch(arr []int, s int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		middle := (left + right) / 2
		middleElem := arr[middle]

		if middleElem == s {
			return middle
		} else if middleElem < s {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1
}
