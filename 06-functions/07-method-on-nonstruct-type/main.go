package main

import (
	"fmt"
	m "math"
)

type MyFloat float64


func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-m.Sqrt2)
	fmt.Println(f.Abs())
}