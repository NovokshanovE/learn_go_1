package main

import "fmt"

func main() {
	var a, b float64 = 5 << 21, 7 << 22 // значения больше 2^20

	fmt.Println("Сложение:", a+b)
	fmt.Println("Вычитание:", a-b)
	fmt.Println("Умножение:", a*b)
	fmt.Println("Деление:", a/b)
}
