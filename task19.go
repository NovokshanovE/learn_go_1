package main

import (
	"fmt"
)

func reverseString(s string) []byte {

	res := make([]byte, 0)

	for i := len(s) - 1; i >= 0; i-- {
		// r, _ :=
		// st := rune(s[i])
		// fmt.Println(utf8.AppendRune(res, st))

		// res = append(res, s[i])
	}
	return res

}
func task19() {
	fmt.Printf("%s\n", "главрыба")
	fmt.Printf("%s\n", reverseString("главрыба"))
}
