package main

import (
	"fmt"
	"math"
)

const PI float64 = 3.14

type Shape interface {
	area() float64
}
type Area interface {
	area() float64
}
type Circle struct {
	radius float64
}

type Square struct {
	lenght float64
}
type Lot struct {
	lenght float64
	width  float64
}

func (s Square) area() float64 {
	return s.lenght * s.lenght
}

func (c Circle) area() float64 {
	return PI * math.Pow(c.radius, 2)
}

func (l Lot) area() float64 {
	return l.lenght * l.width
}
func info(shape Shape) {
	fmt.Println(shape)
	fmt.Println(shape.area())

}
func infoLand(land Area) {
	fmt.Println(land)
	fmt.Println(land.area())
}
func main() {
	s := Square{10}
	c := Circle{20}
	l := Lot{15, 15}
	info(s)
	info(c)
	infoLand(l)
}
