package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circel struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circel) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	var circle, rectangle Shape

	circle = Circel{radius: 2}
	rectangle = Rectangle{width: 5, height: 10}

	fmt.Println(circle.area())
	fmt.Println(rectangle.area())

}
