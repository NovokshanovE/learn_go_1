package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	val int64
}

func (c *Counter) Increment() { // добавление 1 к счетчику
	atomic.AddInt64(&c.val, 1) // с помощью sync.atomic можно безопасно работать с базовыми типами
}

func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&c.val) // чтение данных из счетчика
}

// для реализации конкурентного увеличения значения счетчика я использую атомарные операции
func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	for i := 0; i < 10100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
		fmt.Println(counter.Value())
	}

	wg.Wait()
	fmt.Println(counter.Value())
}
