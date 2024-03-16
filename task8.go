package main

import "fmt"

func setBitByPos(number int64, pos int64, set int64) int64 {
	var one int64 = set << (pos)   // установка единицы в нужную позицию
	var mask int64 = ^(1 << (pos)) // маска для обнуления бита на позиции раной pos
	number &= mask                 // применяем маску на нашем числе для обнуления бита
	return number | (one)          // возвращаем число с установленным на позиции pos 1 или 0 (set)
}

// 0100
func main() {

	fmt.Println(setBitByPos(13, 2, 1)) //1101 -> 1101 - 13
	fmt.Println(setBitByPos(13, 2, 0)) //1101 -> 1001 - 9
	fmt.Println(setBitByPos(13, 1, 1)) //1101 -> 1111 - 15
	fmt.Println(setBitByPos(13, 3, 0)) //1101 -> 0101 - 5

}
