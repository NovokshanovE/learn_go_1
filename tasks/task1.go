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
		// Human: Human{
		// 	A:  "xexe",
		// 	B:  "ewle;",
		// 	Id: 78,
		// },
	}
}

func task1() {
	A := NewAction()
	A.smthPrint()
	fmt.Println(A.B)
	fmt.Println(A)
}
