package shapes

import (
	"math"
)

type Shape interface {
	Area() float64
}

// Rectangle struct, A `struct` is just a named collection of fields where you can store data.
type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area () float64 {
	return r.Width * r.Height
}

// Circle struct, A `struct` is just a named collection of fields where you can store data.
type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area () float64 {
	return (t.Base * t.Height) * 0.5
}

func (c Circle) Area () float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter ...
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area ...
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}