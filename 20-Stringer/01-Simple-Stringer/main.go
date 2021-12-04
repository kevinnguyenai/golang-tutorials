package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Authur", 21}
	b := Person{"Zap Zaboi", 21}
	fmt.Println(a, b)
}