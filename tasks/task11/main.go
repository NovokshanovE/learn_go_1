package main

import "fmt"

func main() {

	// множества реализованы с помощью мапы, так как в множестве значения не должны повторяться
	set1 := map[int]struct{}{}

	set2 := map[int]struct{}{}

	intersection := map[int]struct{}{} // множество пересечений
	// для нахождения пересечения необходимо пройтись по одному из множеств и проверить, есть ли соответствующие ключи в другом множестве
	// если есть то добавляем в пересечение
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			set1[i] = struct{}{}
		}
		if i%3 == 0 {
			set2[i] = struct{}{}
		}
	}

	for key, _ := range set1 {
		if _, ok := set2[key]; ok {
			intersection[key] = struct{}{}
		}
	}
	fmt.Println(set1)
	fmt.Println(set2)
	fmt.Println(intersection)

}
