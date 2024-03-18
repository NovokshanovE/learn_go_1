package main

import (
	"fmt"
	"sync"
)

func variant1() {

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

type MyMap struct {
	mx sync.RWMutex
	m  map[int]int
}

func NewMyMap() *MyMap {
	m := MyMap{
		m: make(map[int]int),
	}
	return &m
}

func (c *MyMap) Load(key int) (int, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	val, ok := c.m[key]
	return val, ok
}

func (c *MyMap) Print() {
	c.mx.RLock()
	defer c.mx.RUnlock()
	for key, val := range c.m {
		fmt.Printf("Ключ: %v, Значение: %v\n", key, val)
	}
}

func (c *MyMap) Store(key int, value int) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.m[key] = value
}

func variant2() {

	m := NewMyMap() // для конкурентной записи в map следует использовать либо самописную структуру, работающую на основе мьютексов, либо sync.Map
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) { // запуск горутины для конкурентной записи в map
			defer wg.Done()
			m.Store(i, i*i) // Запись данных в map
		}(i)
	}

	wg.Wait()

	m.Print()

}

func main() {
	variant2()
}
