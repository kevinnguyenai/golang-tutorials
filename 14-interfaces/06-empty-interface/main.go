package main

import (
	"fmt"
)


func main() {
	var I interface{}
	describe(I)

	I = 42
	describe(I)

	I = "Hello World!"
	describe(I)

}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}