package main

import (
	"fmt"
)

func reverseString(s string) string {
	runes := []rune(s)                                    // для того чтобы перевернуть строку переведем ее в числовой формат rune
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 { // пройдемся по строке до центра 2-мя указателями и поменяем i-й c j-м
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func main() {
	fmt.Printf("%s\n", "главрыба")
	fmt.Printf("%s\n", reverseString("главрыба"))
}
