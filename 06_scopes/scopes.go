package main

import "fmt"

var error int = 45

func wrapper() func(x int) int {
	if error >= 5 {
		return func(x int) int {

			return x * x
		}
	} else {
		return func(x int) int {

			return x + x
		}
	}
}

func main() {
	getValue := wrapper()
	fmt.Println(getValue(5))

}
