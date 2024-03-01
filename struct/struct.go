package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radium float64
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
	area = math.Pi * circle.Radium * circle.Radium
	return
}
