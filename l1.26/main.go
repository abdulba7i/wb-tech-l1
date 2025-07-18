package main

import (
	"fmt"
	"strings"
)

func main() {
	sExmp1 := "abcd"
	sExmp2 := "abCdefAaf"
	sExmp3 := "aabcd"

	fmt.Println(UnicSymbol(sExmp1))
	fmt.Println(UnicSymbol(sExmp2))
	fmt.Println(UnicSymbol(sExmp3))
}

func UnicSymbol(s string) bool {
	s = strings.ToLower(s) // приведем все символы в нижний регистр, так как регистр буквы не имеет значения
	m := map[string]int{}  // значение у нас будет типа int, он и будет сыграть у нас в роли счётчика

	for _, v := range s {
		m[string(v)]++

		// если в одном из ключе, значения встречается больше одного раза, возвращаем false
		if m[string(v)] > 1 {
			return false
		}
	}

	return true
}
