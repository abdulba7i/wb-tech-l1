package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 4 // индекса элемента, который желаем удалить

	fmt.Println(elemRemove(slice, i))
}

func elemRemove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1] // отрезаем последний элемент
}
