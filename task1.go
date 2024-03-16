package main

import "fmt"

type Human struct {
	Id int
	A  string
	B  string
}

type Action struct {
	Id int
	Human
	C int
	B float64
}

func (h Human) smthPrint() {
	fmt.Println("Human method")
}

func NewAction() *Action {
	return &Action{
		C:  12,
		B:  132.829,
		Id: 17,
	}
}

func main() {
	A := NewAction()
	A.smthPrint()
}
