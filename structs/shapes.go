package structs

import "math"

// interface declaration
// in go, interface resolution is implicit
// if the type you pass matches interface, it will compile
type Shape interface {
	Area() float64
}

// create simple type using a struct, named collecton of fields to store data
type Rectangle struct {
	Width  float64
	Height float64
}

// method used for specific struct
// given func (receiverName ReceiverType) MethodName(args)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectangle) float64 {

	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
