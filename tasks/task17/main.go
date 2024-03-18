package main

import (
	"fmt"
	"sort"
)

// бинарный поиск работает только в отсортированном массиве
// мы проверяем средний элемент в текущем подсписке на то, больше он целевого или меньше
// если меньше, то берем правую часть
// если больше, то левую
// если равен, то мы нашли целевой элемент
func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2 // поиск середины

		if arr[mid] == target { // равенство
			return mid
		}

		if arr[mid] < target { // если меньше
			left = mid + 1
		} else { // если больше
			right = mid - 1
		}
	}

	return -1
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !sort.SliceIsSorted(numbers, func(p, q int) bool { // проверка на то, что массив отсортирован
		return numbers[p] < numbers[q]
	}) {
		fmt.Printf("Массив должен быть отсортирован\n")
		return
	}
	target := 1
	result := binarySearch(numbers, target)

	if result != -1 {
		fmt.Printf("Найдено на позиции %d\n", result)
	} else {
		fmt.Printf("Не найдено\n")
	}
}
