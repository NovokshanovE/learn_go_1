package main

import "fmt"

func main() {
	array := []string{"cat", "cat", "dog", "cat", "tree"}
	// множество  реализовано по той же логиеке, как и в задании с пересечением(11)
	set := map[string]struct{}{}

	for _, val := range array {
		if _, ok := set[val]; !ok {
			set[val] = struct{}{}
		}
	}

	fmt.Println(set)
}
