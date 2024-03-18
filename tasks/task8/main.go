package main

import (
	"fmt"
	"strconv"
)

func setBitByPos(number int64, pos int64, set int64) int64 {
	if number > 0 {
		var one int64 = set << (pos)   // установка единицы в нужную позицию
		var mask int64 = ^(1 << (pos)) // маска для обнуления бита на позиции раной pos
		number &= mask                 // применяем маску на нашем числе для обнуления бита
		return number | (one)          // возвращаем число с установленным на позиции pos 1 или 0 (set)
	} else {
		// set = 1 - set // если число отрицательное
		//(это сделано для того, чтобы мы получали по модулю то же число что и при работе с положительными числами)

		var one int64 = set << (pos)   // установка единицы в нужную позицию
		var mask int64 = ^(1 << (pos)) // маска для обнуления бита на позиции раной pos
		number &= mask                 // применяем маску на нашем числе для обнуления бита
		return number | (one)          // возвращаем число с установленным на позиции pos 1 или 0 (set)
	}

}

// 0100
func main() {

	fmt.Println(setBitByPos(13, 2, 1))  //1101 -> 1101 - 13
	fmt.Println(setBitByPos(13, 2, 0))  //1101 -> 1001 - 9
	fmt.Println(setBitByPos(13, 1, 1))  //1101 -> 1111 - 15
	fmt.Println(setBitByPos(13, 3, 0))  //1101 -> 0101 - 5
	fmt.Println(setBitByPos(13, 3, 0))  //1101 -> 0101 - 5
	fmt.Println(setBitByPos(-13, 3, 0)) //-1101 -> -0101 - -5
	fmt.Println(setBitByPos(-13, 1, 1)) //-1101 -> -1111 - -15
	fmt.Println(strconv.FormatInt(-15, 2))

}
