package main

import "fmt"

func main() {
	arr1 := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("Unique words from arr1:", UnicArr(arr1))

	arr2 := []string{"apple", "banana", "apple", "orange", "banana", "grape"}
	fmt.Println("Unique words from arr2:", UnicArr(arr2))

	arr3 := []string{"one", "two", "three", "two", "one", "four"}
	fmt.Println("Unique words from arr3:", UnicArr(arr3))
}

func UnicArr(arr []string) []string {
	result := []string{}
	mapUnic := map[string]struct{}{}

	for _, v := range arr {
		mapUnic[v] = struct{}{}
	}

	for key := range mapUnic {
		result = append(result, key)
	}

	return result
}
