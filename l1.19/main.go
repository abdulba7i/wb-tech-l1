package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var text string
	fmt.Scan(&text)

	fmt.Println(ReversedWord(text))
}

func ReversedWord(text string) string {
	rn := []rune(text)
	newrn := ""
	for i := utf8.RuneCountInString(string(rn)) - 1; i > -1; i-- {
		newrn += string(rn[i])
	}

	return newrn
}
