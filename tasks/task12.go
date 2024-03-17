package main

import "fmt"

func task12() {
	array := []string{"cat", "cat", "dog", "cat", "tree"}
	set := map[string]struct{}{}

	for _, val := range array {
		if _, ok := set[val]; !ok {
			set[val] = struct{}{}
		}
	}

	fmt.Println(set)
}
