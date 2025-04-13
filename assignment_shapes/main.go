package main

import "fmt"

type Shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func main() {
	//// My solution
	//s := square{sideLength: 5}
	//t := triangle{height: 3, base: 10}
	//fmt.Println(printArea(s))
	//fmt.Println(printArea(t))

	// Model solution
	t := triangle{base: 10, height: 10}
	s := square{sideLength: 10}
	printArea(t)
	printArea(s)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

//// My solution
//func printArea(s Shape) string {
//	area := s.getArea()
//	return "The area is: " + fmt.Sprintf("%f", area)
//}

// Model solution
func printArea(s Shape) {
	fmt.Println(s.getArea())
}
