package main

import (
	"fmt"
	"time"
)

func task5() {
	c := make(chan int)
	// start := time.Now()
	// time.Sleep(5)

	go func() {
		for i := 0; ; i++ {
			c <- i
			time.Sleep(time.Second / 10)
		}
	}()

	// Получение значений из канала
	go func() {
		for {
			msg := <-c
			fmt.Println(msg)
		}
	}()

	time.Sleep(time.Second * 5)

}
