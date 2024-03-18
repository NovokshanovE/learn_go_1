package main

import (
	"fmt"
	"reflect"
)

func main() {
	variables := []interface{}{123, "Hello, World!", true, make(chan int)}
	// для нахождения типа переменной можно воспользоваться пакетом reflect и методом TypeOf
	for _, variable := range variables {
		switch reflect.TypeOf(variable).Kind() {
		case reflect.Int:
			fmt.Println("Тип переменной: int")
		case reflect.String:
			fmt.Println("Тип переменной: string")
		case reflect.Bool:
			fmt.Println("Тип переменной: bool")
		case reflect.Chan:
			fmt.Println("Тип переменной: channel")
		default:
			fmt.Println("Неизвестный тип переменной")
		}
	}
}
