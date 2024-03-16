package main

import "fmt"

func reverseWords(s string) string {
	str := []rune(s)
	num_space := len(str)
	space := []int{}
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == ' ' {
			num_space = i
			space = append(space, -1)
		} else {
			space = append(space, num_space)
		}

	}

	center_start := 0
	for i := len(str) - 1; i >= 0; i-- {
		j := (len(str) - 1) - i
		if space[i] == -1 {
			center_start = j + 1

		} else {
			center := (center_start + space[i] - 1) / 2
			if j > center {
				continue
			}
			if (space[i]-center_start)%2 == 1 {

				str[j], str[center+center-j] = str[center+center-j], str[j]
			} else {

				str[j], str[center+center-j+1] = str[center+center-j+1], str[j]
			}

		}

	}
	return string(str)
}

func main() {
	fmt.Println(reverseWords("abs dfret rtere"))

}
