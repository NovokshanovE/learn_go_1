package main

import (
	"fmt"
)

func main() {
	data := []int{1, 2, 3, 4, 5} // Исходные данные

	ch1 := make(chan int) // Первый канал
	ch2 := make(chan int) // Второй канал

	go func() { // Запуск горутины для записи данных в первый канал
		for _, x := range data {
			ch1 <- x
		}
		close(ch1) // Закрытие канала
	}()

	go func() { // Запуск горутины для чтения данных из первого канала
		for x := range ch1 {
			ch2 <- x * 2
		}
		close(ch2) // Закрытие канала
	}()

	for res := range ch2 { // Чтение результатов из второго канала и вывод их в stdout
		fmt.Println(res)
	}
}
