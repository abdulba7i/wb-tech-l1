package main

import (
	"fmt"
)

var justString string

func main() {
	result := someFunc()
	fmt.Println(result)
}

// Основная проблема - это утечка памяти, она обращается на срез от большой строки v, из-за чего вся огромная строка остаётся в памяти
// даже если используется только небольшая её часть. Чтобы это исправить нужно создать копию нужно части строки.
func someFunc() string {
	v := createHugeString(1 << 10)
	justStringBytes := make([]byte, 100)

	copy(justStringBytes, v)
	justString := string(justStringBytes) // создаём новую строку, не зависящая от v

	return justString
}

func createHugeString(sizeStr int) string {
	result := make([]byte, sizeStr)
	for i := 0; i < sizeStr; i++ {
		result[i] = '-'
	}
	return string(result)
}
