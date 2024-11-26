package main

import "fmt"

// Define an interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a struct that satisfies the interface
type Rectangle struct {
	Width, Height float64
}

// Implement the methods for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Another struct satisfying the interface
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// Function using the Shape interface
func PrintShapeDetails(s Shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

func main() {
	r := Rectangle{Width: 4, Height: 5}
	c := Circle{Radius: 3}

	fmt.Println("Rectangle:")
	PrintShapeDetails(r)

	fmt.Println("\nCircle:")
	PrintShapeDetails(c)
}
