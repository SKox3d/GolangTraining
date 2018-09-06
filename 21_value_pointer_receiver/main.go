package main

import (
	"fmt"
)

type Shape interface {
	area() float64
}

type Square struct {
	lenght float64
}

func (s Square) area() float64 {
	return s.lenght * s.lenght
}

func info(shape Shape) {
	fmt.Println(shape.area())

}

func main() {
	s := Square{10}
	info(&s)
}
