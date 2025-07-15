package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// №1 - Выход по условию
	fmt.Println("№1 - Выход по условию")
	ch1 := make(chan int)

	go stopIF(ch1)
	for val := range ch1 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(val)
	}
	fmt.Println("Work Stop")

	time.Sleep(1 * time.Second)

	// №2 - Через канал уведомления
	fmt.Println("№2 - Через канал уведомления")
	ch2 := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		notificationsChanStop(ch2)
	}()

	time.Sleep(3 * time.Second)
	close(ch2)
	wg.Wait()

	// №3 - через Контекст
	fmt.Println("№3 - Через контекст")
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)

	go func() {
		defer wg.Done()
		contextStop(ctx)
	}()

	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()

	// №4 - Остановка работы runtime.Goexit()
	fmt.Println("№4 - Остановка работы runtime.Goexit()")

	ch4 := make(chan int)
	wg.Add(1)

	go func() {
		defer wg.Done()
		GoexitStop(ch4)
	}()

	go func() {
		time.Sleep(1 * time.Second)
		close(ch4)
	}()

	for val := range ch4 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(val)
	}

	wg.Wait()

	// №5 - Остановка за счёт time.After
	fmt.Println("№5 - Остановка за счёт time.After")

	ch5 := make(chan int)
	wg.Add(1)
	timer := time.After(2 * time.Second)

	go timeAfterStop(ch5, timer, &wg)

	for val := range ch5 {
		fmt.Printf("Got message: %d\n", val)
	}

	wg.Wait()
	fmt.Println("Exit.")
	fmt.Println("Все горутины завершены!")
}

// №1
func stopIF(ch chan int) {
	count := 0
	for {
		if count == 10 {
			close(ch)
			break
		}
		ch <- count
		count++
	}
}

// №2
func notificationsChanStop(stop <-chan struct{}) {
	for {
		select {
		case <-stop:
			fmt.Println("stop!")
			return
		default:
			fmt.Println("work...")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

// №3
func contextStop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop!")
			return
		default:
			fmt.Println("working...")
			time.Sleep(700 * time.Millisecond)
		}
	}
}

// №4
func GoexitStop(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	runtime.Goexit()
}

// №5
func timeAfterStop(ch chan int, timer <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	counter := 0

	for {
		select {
		case <-timer:
			close(ch)
			return
		case ch <- counter:
			counter++
			time.Sleep(300 * time.Millisecond)
		}
	}
}
