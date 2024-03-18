package main

import (
	"fmt"
	"sync"
)

func main() {

	var m sync.Map // для конкурентной записи в map следует использовать либо самописную структуру, работающую на основе мьютексов, либо sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) { // запуск горутины для конкурентной записи в map
			defer wg.Done()
			m.Store(i, i*i) // Запись данных в map
		}(i)
	}

	wg.Wait()

	m.Range(func(key, value interface{}) bool { // вывод значений
		fmt.Printf("Ключ: %v, Значение: %v\n", key, value)
		return true
	})

}
