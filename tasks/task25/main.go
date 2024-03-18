package main

import (
	"fmt"
	"time"
)

func sleep(sec int) {
	<-time.After(time.Duration(sec) * time.Second)
}

func main() {
	fmt.Println("Start")
	sleep(5) // sleep на 5 секунд
	fmt.Println("Finish")
}
