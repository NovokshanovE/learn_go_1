package main

import "fmt"

func task10() {
	array := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	res := map[int]([]float32){}

	for _, val := range array {
		fmt.Println(int(val) - int(val)%10)
		if _, ok := res[int(val)-int(val)%10]; ok {
			res[int(val)-int(val)%10] = append(res[int(val)-int(val)%10], val)
		} else {
			res[int(val)-int(val)%10] = []float32{val}
		}
	}

	fmt.Println(res)

}
