package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)

	x := 5
	setZero(&x)
	fmt.Println(x)
}

func setZero(x *int) {
	*x = 0
}
