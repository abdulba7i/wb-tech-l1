package main

import "fmt"

type Human struct {
	ID      int
	Name    string
	Age     int
	Message string
}

func (h Human) InfoHuman() {
	fmt.Printf("ID: %d, Name: %s, Age: %d\n", h.ID, h.Name, h.Age)
}

func (h Human) Say() {
	fmt.Printf("A man named %s say: %s\n", h.Name, h.Message)
}

type Action struct {
	Human
}

func main() {
	// Первый вариант объявления
	a := Action{
		Human{
			ID:      1,
			Name:    "AbdulBari",
			Age:     21,
			Message: "Hi",
		},
	}

	// Второй вариант объявления
	// var a Action
	// a.ID, a.Name, a.Age, a.Message = 1, "AbdulBari", 21, "Hi"

	a.InfoHuman()
	a.Say()
}
