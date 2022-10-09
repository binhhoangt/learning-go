package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triagle struct {
	height float64
	base   float64
}

func main() {
	t := triagle{height: 10, base: 0.5}
	s := square{sideLength: 4}

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triagle) getArea() float64 {
	return 0.5 * t.base * t.height
}
