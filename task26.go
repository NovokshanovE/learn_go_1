package main

import "fmt"

func unique(s string) bool {
	// str := []rune(s)
	res := make(map[rune]int)
	for _, i := range s {

		if 65 <= i && i <= 90 {
			i += 32
		}
		// fmt.Println(i)
		if _, ok := res[i]; ok {
			return false
		} else {
			res[i] = 1
		}
	}
	return true
}

func task26() {
	fmt.Println(unique("aabcd"))
}
