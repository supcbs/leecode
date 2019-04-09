package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println("1233")
	go p()
	time.Sleep(10*time.Second)
	fmt.Println("123")
}

func p () {
	panic("123")
}