//golang does not have classes (structs and interfaces)

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}
type Rectangle struct {
	width, height float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}
func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {

	circle1 := Circle{x: 0, y: 0, radius: 10}
	rectangle1 := Rectangle{width: 10, height: 20.9}

	fmt.Printf("Circle area: %f\n", getArea(circle1))
	fmt.Printf("Circle area: %f\n", getArea(rectangle1))

}
