package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, id int, msg <-chan string) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%d: exit programm\n", id)
			return
		case m, ok := <-msg:
			if !ok {
				fmt.Printf("%d: chan close\n", id)
				return
			}
			fmt.Printf("Worker %d: got message %s\n", id, m)
		}
	}
}

func main() {
	var N, counter int // количество воркеров и счётчик
	var wg sync.WaitGroup
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	fmt.Print("Enter the number of workers: ")
	fmt.Scan(&N)

	msg := make(chan string, 100)
	// wg.Add(100) — если знаем заранее о количестве задач, то можем указать сразу

	for i := 0; i < N; i++ {
		wg.Add(1) // если же не известно заранее, то будем вместес с циклом прибавлять по одной задачки
		go worker(ctx, &wg, i, msg)
	}

	// будем параллельно бегать по второму циклу, чтобы основной код мог спокойно рабоать,
	// и мы могли бы в любой момент завершить программу Ctrl+C
	go func() {
		for {
			select {
			case <-ctx.Done(): // если пришло завершение, то прерываем программу
				close(msg) // и сразу же закрываем канал
				return
			case msg <- fmt.Sprintf("number: %v", counter): // иначе отправляем поток сообщений
				time.Sleep(300 * time.Millisecond)
				counter++
			}
		}
	}()

	// Ожидаем завершения всех воркеров
	wg.Wait()
	fmt.Println("Programm exit.")
}
