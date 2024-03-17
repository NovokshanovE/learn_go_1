package main

import (
	"fmt"
	"sync"
)

func work(c chan int, res sync.Map, wg sync.WaitGroup) {
	for {
		i := <-c
		if val, ok := res.Load(i); ok {

			res.Store(i, val.(int)+1)
		} else {
			res.Store(i, 1)
		}
	}
	// wg.Done()
}

func task7() {

	var m sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(i, i*i) // Запись данных в map
		}(i)
	}

	wg.Wait()

	m.Range(func(key, value interface{}) bool {
		fmt.Printf("Ключ: %v, Значение: %v\n", key, value)
		return true
	})

}
