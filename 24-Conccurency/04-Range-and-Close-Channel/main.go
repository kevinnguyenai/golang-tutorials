package main

import (
	"fmt"
)

func Fibbonaci(n int, c chan int64) {
	var x int64 = 0
	var y int64 = 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int64, 10)
	go Fibbonaci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
