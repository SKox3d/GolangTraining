package main

import "fmt"

func main() {
	mySlice := make([][]string, 0)

	fmt.Println(mySlice)
	student1 := make([]string, 4)
	student1[0] = "Constantin"
	student1[1] = "Sipitca"
	student1[2] = "370"
	student1[3] = "250"
	mySlice = append(mySlice, student1)
	fmt.Println(mySlice)

}
