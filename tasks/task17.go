package main

import (
	"fmt"
	"sort"
)

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func task17() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if !sort.SliceIsSorted(numbers, func(p, q int) bool {
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
