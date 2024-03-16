package main

import (
	"fmt"
)

func main() {
	data := []int{1, 2, 3, 4, 5} // Исходные данные

	ch1 := make(chan int) // Первый канал
	ch2 := make(chan int) // Второй канал

	// Запуск горутины для записи данных в первый канал
	go func() {
		for _, x := range data {
			ch1 <- x
		}
		close(ch1) // Закрытие канала после записи всех данных
	}()

	// Запуск горутины для чтения данных из первого канала, умножения их на 2 и записи во второй канал
	go func() {
		for x := range ch1 {
			ch2 <- x * 2
		}
		close(ch2) // Закрытие канала после записи всех данных
	}()

	// Чтение результатов из второго канала и вывод их в stdout
	for res := range ch2 {
		fmt.Println(res)
	}
}