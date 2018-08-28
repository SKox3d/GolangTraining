package main

import "fmt"

func main() {
	for i := 65; i <= 126; i++ {
		fmt.Printf("%d - %q \n", i, i)
	}
}
