package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
}

func (rectangle Rectangle) Perimeter() (perimeter float64) {
	perimeter = (rectangle.Width + rectangle.Height) * 2
	return
}

func (rectangle Rectangle) Area() (area float64) {
	area = rectangle.Width * rectangle.Height
	return
}

func (circle Circle) Area() (area float64) {
	area = math.Pi * circle.Radius * circle.Radius
	return
}

func (triangle Triangle) Area() (area float64) {
	area = triangle.Base * triangle.Height / 2
	return
}
