package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() { // горутина отправитель
		for i := 0; ; i++ {
			c <- i
			time.Sleep(time.Second / 10)
		}
	}()

	go func() { // горутина приемник
		for {
			msg := <-c
			fmt.Println(msg)
		}
	}()

	time.Sleep(time.Second * 5) // главная горутина будет работать 5 секунд

}
