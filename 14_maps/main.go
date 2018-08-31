package main

import "fmt"

func main() {
	myMap := map[int]string{
		0: "Constantin",
		1: "Doinita",
		2: "Ion",
	}

	for key, value := range myMap {
		fmt.Println("Key:", key, "Value:", value)
	}
}
