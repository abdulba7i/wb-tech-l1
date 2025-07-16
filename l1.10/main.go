package main

import "fmt"

func main() {
	slice := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(sorted(slice))
}

func sorted(slice []float64) map[int][]float64 {
	set := map[int][]float64{}

	for _, v := range slice {
		intV := int(v)
		keySet := intV - (intV % 10)
		set[keySet] = append(set[keySet], v)
	}
	return set
}
