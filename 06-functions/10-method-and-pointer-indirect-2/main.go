package main 

import (
	"fmt"
	m "math"
)

type Vertex struct {
	X,Y float64
}

func (v Vertex) Abs() float64 {
	return m.Sqrt(v.X*v.X + v.Y*v.Y);
}

func AbsFunc(v *Vertex) float64 {
	return m.Sqrt(v.X*v.X + v.Y*v.Y);
}


func main() {
	// Assignment 1
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(&v))

	// Assignment 2
	p := &Vertex{3, 4}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(p))

}