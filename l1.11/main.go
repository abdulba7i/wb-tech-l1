package main

import (
	"fmt"
	"math/rand"
)

func randomArr() []int {
	arr := make([]int, rand.Intn(15))
	for i := range arr {
		arr[i] = rand.Intn(50)
	}
	return arr
}

func main() {
	arr1, arr2 := randomArr(), randomArr()
	fmt.Printf("Исходные массивы: %v # %v\n", arr1, arr2)

	fmt.Printf("Пересечение двух множеств: %v\n", searchIntersections(arr1, arr2))
}

func searchIntersections(arr1, arr2 []int) []int {
	m, repeatMap := make(map[int]struct{}), make(map[int]struct{})
	result := []int{}

	for i := 0; i < len(arr1); i++ {
		m[arr1[i]] = struct{}{}
	}

	for j := 0; j < len(arr2); j++ {
		if _, ok := repeatMap[arr2[j]]; ok {
			continue
		}
		if _, ok := m[arr2[j]]; ok {
			result = append(result, arr2[j])
			repeatMap[arr2[j]] = struct{}{}
		}
	}

	return result
}
