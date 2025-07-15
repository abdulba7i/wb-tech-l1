package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	timer := time.After(5 * time.Second) // создаем канал, который через 5 секунд вернет значение

	go worker(ch, timer)

	for val := range ch {
		fmt.Printf("Got message: %d\n", val)
	}

	fmt.Println("Exit")
}

func worker(ch chan int, timer <-chan time.Time) {
	counter := 0

	// имитируем бесконечную работу, которая будет работать до тех пор, пока не пройдет заданное время
	for {
		select {
		case <-timer: // если таймер сработал, то закрываем канал и завершаем работу
			close(ch)
			return
		case ch <- counter: // иначе отправляем данные в канал
			counter++
			time.Sleep(300 * time.Millisecond)
		}
	}
}
