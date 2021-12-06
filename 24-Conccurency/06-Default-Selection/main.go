package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(1000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println(" tok.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("tick.")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
