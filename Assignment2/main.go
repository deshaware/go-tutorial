package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	sqr := square{sideLength: 2.5}
	tri := triangle{height: 1.2, base: 4}

	// fmt.Println(sqr.getArea())
	// fmt.Println(tri.getArea())

	fmt.Println("The area of square is", printArea(sqr))
	fmt.Println("The area of triangle is", printArea(tri))

}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) float64 {
	return s.getArea()
}
