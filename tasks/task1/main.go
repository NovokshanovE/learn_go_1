package main

import "fmt"

// структура Human
type Human struct {
	Id int
	A  string
	B  string
}

// структура Action
type Action struct {
	Id    int
	Human // объект класса Human неявно задан в струкутре Action
	C     int
	B     float64
}

// Метод структуры Human
func (h Human) smthPrint() {
	fmt.Println("Human method")
}

// конструктор для инициализации объектов структуры Action
func NewAction() *Action {
	return &Action{
		C:  12,
		B:  132.829,
		Id: 17,
	}
}

func main() {
	A := NewAction()
	A.smthPrint() //мы можем использовать методы структуры Human без явного обращения к ней
}
