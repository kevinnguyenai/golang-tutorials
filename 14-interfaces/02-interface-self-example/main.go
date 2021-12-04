package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Structure to implement Shape for Circle
type Circle struct {
	Radius float64
}

func (r *Circle) Area() float64 {
	return math.Pi*r.Radius*r.Radius
}

func (r *Circle) Perimeter() float64 {
	return 2*math.Pi*r.Radius
}

// Structure to implement Shape for Rectangular
type Rectangular struct {
	X, Y float64
}

func (r *Rectangular) Area() float64 {
	return r.X*r.Y
}

func (r *Rectangular) Perimeter() float64 {
	return 2*(r.X+r.Y)
}


// Main Func

func main() {
	// Circle implement Shape interface
	r := Circle{10}
	fmt.Println(r.Area())
	fmt.Println(r.Perimeter())

	// Rectangular implement Shape interface
	rec := Rectangular{3, 4}
	fmt.Println(rec.Area())
	fmt.Println(rec.Perimeter())
}