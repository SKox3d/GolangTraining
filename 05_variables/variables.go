package main

import "fmt"

func main() {
	fmt.Println("Variables:")
	a := 10
	b := "Constantin"
	c := 3.14
	d := true

	var e int
	var f string
	var g float64
	var h bool

	fmt.Printf("%v -- %T \n", a, a)
	fmt.Printf("%v -- %T \n", b, b)
	fmt.Printf("%v -- %T \n", c, c)
	fmt.Printf("%v -- %T \n", d, d)

	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Printf("%v \n", h)

	if f == "" {
		fmt.Println("true")
	}
}
