package main

import "fmt"

func solve_t3(a int, res chan int) { // функция для запуска горутины
	res <- a * a
}

func main() {
	array := [5]int{2, 4, 6, 8, 10}
	// from := make(chan bool)
	to := make(chan int)
	res := 0
	for _, i := range array {

		go solve_t3(i, to) // вызов горутины

	}

	for i := 0; i < 5; i++ {
		res += <-to // чтение данных из канала
	}
	fmt.Println(res) // вывод результата
	// <-from
}
