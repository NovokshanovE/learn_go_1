package main

import (
	"fmt"
)

// логика быстрой соритировки заключается в нахождении опорного элемента и сортировки элементов относительно его
// справа числа, которые больше, слева - меньше
func quickSort(arr []int, left, right int) {
	if left < right {
		var pivot = partition(arr, left, right) // новое положение опорного элемента, относительно которого мы разделим на список на 2 других для сортировки подсписков
		quickSort(arr, left, pivot)
		quickSort(arr, pivot+1, right)
	}
}

// ыункция для сортировки чсел в подсписке относительно опорного элемента
func partition(arr []int, left, right int) int {

	var pivot = arr[left] // выбор опорного элемента может быть другим
	var i = left
	var j = right

	for i < j {
		for arr[i] <= pivot && i < right {
			i++
		}
		for arr[j] > pivot && j > left {
			j--
		}

		if i < j {
			var temp = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	arr[left] = arr[j]
	arr[j] = pivot

	return j
}

func main() {
	var arr = []int{15, 3, 12, 6, -9, 9, 0}

	fmt.Print("Before Sorting: ")
	fmt.Println(arr)

	quickSort(arr, 0, len(arr)-1)
	fmt.Print("After Sorting: ")
	fmt.Println(arr)
}
