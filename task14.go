package main

import (
	"fmt"
	"reflect"
)

func main() {
	variables := []interface{}{123, "Hello, World!", true, make(chan int)}

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
