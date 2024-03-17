package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	val int64
}

func (c *Counter) Increment() {
	atomic.AddInt64(&c.val, 1)
}

func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&c.val)
}

func task18() {
	var wg sync.WaitGroup
	counter := Counter{}

	for i := 0; i < 10100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println(counter.Value())
}
