package main

import "fmt"

func main() {
	variableType(50)
	variableType("hello")
	variableType(true)
	variableType(make(chan any))
}

func variableType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("type int")
	case string:
		fmt.Println("type string")
	case bool:
		fmt.Println("type bool")
	case chan any:
		fmt.Println("type chan")
	}
}
