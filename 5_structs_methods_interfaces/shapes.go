// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces

package structs

import "math"

type Shape interface { // can be used by different types (parametric polymorphism: https://en.wikipedia.org/wiki/Parametric_polymorphism)
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 { // This is a function
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() float64 { // This is a method. The only difference between function and method is the syntax of the method receiver 'func (receiverName ReceiverType) MethodName(args)': It is a convention in Go to have the receiver variable be the first letter of the type.
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 { // go does not have 'method overloading'
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}
