package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func worker(id int, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for n := range data {
		fmt.Printf("Воркер %d обработал данные %d\n", id, n)
	}
}

func task4() {
	workers := 5 // количество воркеров

	data := make(chan int)

	// Отслеживание Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var wg sync.WaitGroup
	wg.Add(workers)

	// Запуск воркеров
	for i := 1; i <= workers; i++ {
		go worker(i, data, &wg)
	}

	go func() {
		// Постоянная запись данных в канал
		for i := 0; ; i++ {
			select {
			case <-c:
				close(data)
				return
			default:
				data <- i
			}
		}
	}()

	wg.Wait()
	fmt.Println("Программа успешно завершила работу")
}
