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
		fmt.Printf("Worker = %d, data = %d\n", id, n)
	}
}

func main() {
	workers := 5 // количество воркеров

	data := make(chan int)

	// Отслеживание Ctrl+C
	c := make(chan os.Signal, 1)   // создание канала для отслеживания
	signal.Notify(c, os.Interrupt) // отслеживание именно Ctrl+C

	var wg sync.WaitGroup // создание WaitGroup для синхронизации
	wg.Add(workers)

	for i := 1; i <= workers; i++ {
		go worker(i, data, &wg) // Запуск воркеров
	}

	go func() {

		for i := 0; ; i++ {
			select { // блокировка до тех пор, пока не будет получен сигнал Ctrl+C
			case <-c: // если Ctrl+C, то закрываем канал data и завершениее отправляющей горутины
				close(data)
				return
			default:
				data <- i // передача данных в канал
			}
		}
	}()

	wg.Wait()
	fmt.Println("Программа успешно завершила работу")
}
