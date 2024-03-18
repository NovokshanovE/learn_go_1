package main

var justString string

// -------------------------------------------------
// Befor
/*
	Данный участок кода может приводить к утечкам памяти,
	так как  строка v остается в памяти до тех пор, пока
	justString не будет очищена.
*/
func someFunc_Befor() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func task15_Befor() {
	someFunc()
}

// -------------------------------------------------
// After
/*
	В этом коде мы создаем новую строку из первых 100
	символов v, что позволяет освободить память,
	занимаемую остальной частью v, после выполнения
	функции someFunc()..
*/
var justString string

func someFunc_After() {
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100]))
}

func task15_After() {
	someFunc()
}
